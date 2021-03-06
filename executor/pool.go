/*
   The pool client pretends to be a large number of clients in order
   to exert a large load on the server.  In reality each client is
   running in an go routine in parallel.

   Therefore when we do a hunt, each pool client goroutine will
   receive the same VQL query and run the same code. This reduces the
   load the pool client can impart since it is busy running the same
   query multiple times.

   This pool executor memoizes the results from each query in memory
   so each query is run only once but the results are returned from
   each goroutine fake client as it was unique. This increases the
   total number of pool clients we can support since most of the work
   is pushed out to the comms.
*/

package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
	actions_proto "www.velocidex.com/golang/velociraptor/actions/proto"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	crypto_proto "www.velocidex.com/golang/velociraptor/crypto/proto"
	"www.velocidex.com/golang/velociraptor/json"
	"www.velocidex.com/golang/velociraptor/utils"
)

var (
	pool_mu sync.Mutex
	// Get transactions by session id
	session_id_cache = make(map[string]*transaction)

	// Get transactions by query name
	query_cache = make(map[string]*transaction)
	ts          = time.Now().UnixNano()
)

type transaction struct {
	Request   *crypto_proto.GrrMessage
	Responses []*crypto_proto.GrrMessage

	// While the transaction is running we need to make other
	// threads wait until it is done.
	IsDone chan bool
	Done   bool
}

func getInc() int64 {
	pool_mu.Lock()
	defer pool_mu.Unlock()

	ts++
	return ts
}

type PoolClientExecutor struct {
	*ClientExecutor
	Outbound chan *crypto_proto.GrrMessage
	id       int
}

func (self *PoolClientExecutor) ReadResponse() <-chan *crypto_proto.GrrMessage {
	return self.Outbound
}

// Inspect the request and decide if we will cache it under a query.
func getQueryName(message *crypto_proto.GrrMessage) string {
	if message.VQLClientAction != nil {
		for _, query := range message.VQLClientAction.Query {
			if query.Name != "" {
				return query.Name
			}
		}

	}

	return ""
}

func getCompletedTransaction(message *crypto_proto.GrrMessage) *transaction {
	pool_mu.Lock()
	defer pool_mu.Unlock()

	query_name := getQueryName(message)
	// Do not cache empty queries.
	if query_name == "" {
		return nil
	}

	result, pres := query_cache[query_name]

	// Transaction fully cached and completed.
	if pres {
		return result
	}

	// There is no transaction there yet so build one ready for
	// the results.
	trans := &transaction{
		Request: message,
		IsDone:  make(chan bool),
	}
	query_cache[query_name] = trans
	session_id_cache[message.SessionId] = trans

	fmt.Printf("Starting transaction for %v\n", message.SessionId)
	return nil
}

// Feed a server request to the executor for execution.
func (self *PoolClientExecutor) ProcessRequest(
	ctx context.Context,
	message *crypto_proto.GrrMessage) {

	tran := getCompletedTransaction(message)
	if tran != nil {
		// Wait until the transaction is done.
		<-tran.IsDone

		fmt.Printf("Getting %v responses from cache\n", len(tran.Responses))
		for _, resp := range tran.Responses {
			response := &crypto_proto.GrrMessage{
				SessionId:   message.SessionId,
				RequestId:   message.RequestId,
				ResponseId:  uint64(getInc()),
				TaskId:      message.TaskId,
				VQLResponse: self.maybeTransformResponse(resp.VQLResponse),
				LogMessage:  resp.LogMessage,
				Status:      resp.Status,
			}
			self.Outbound <- response
		}
		return
	}

	self.ClientExecutor.ProcessRequest(ctx, message)
}

func (self *PoolClientExecutor) maybeTransformResponse(
	response *actions_proto.VQLResponse) *actions_proto.VQLResponse {

	if response != nil {

		// We need to make the Hostname unique so if the response contains
		// a Hostname we need to transform it.
		if utils.InString(response.Columns, "Hostname") {
			rows, err := utils.ParseJsonToDicts([]byte(response.JSONLResponse))
			if err != nil || len(rows) == 0 {
				utils.Debug(err)
				return response
			}

			// Replace the Hostname
			hostname, pres := rows[0].Get("Hostname")
			if !pres {
				return response
			}

			new_hostname := fmt.Sprintf("%s-%d", hostname, self.id)
			rows[0].Set("Fqdn", new_hostname)
			rows[0].Set("Hostname", new_hostname)

			new_rows, err := json.MarshalJsonl(rows)
			if err != nil {
				return response
			}
			result := proto.Clone(response).(*actions_proto.VQLResponse)
			result.JSONLResponse = string(new_rows)

			return result
		}
	}
	return response
}

func NewPoolClientExecutor(
	ctx context.Context,
	config_obj *config_proto.Config, id int) (*PoolClientExecutor, error) {
	exe, err := NewClientExecutor(ctx, config_obj)
	if err != nil {
		return nil, err
	}

	output := make(chan *crypto_proto.GrrMessage, 10)

	go func() {
		delegate_messages := exe.ReadResponse()
		for {
			select {
			case <-ctx.Done():
				return

			case message := <-delegate_messages:
				if message.SessionId != "F.Monitoring" {
					maybeCacheResult(message)
				}
				output <- message
			}
		}
	}()

	return &PoolClientExecutor{
		ClientExecutor: exe,
		Outbound:       output,
		id:             id,
	}, nil
}

func maybeCacheResult(response *crypto_proto.GrrMessage) {
	pool_mu.Lock()
	defer pool_mu.Unlock()

	session_id := response.SessionId

	// Check if the transaction is tracked
	tran, pres := session_id_cache[session_id]
	if pres {
		fmt.Printf("%v\n", response)
		tran.Responses = append(tran.Responses, response)
		if response.Status != nil && !tran.Done {
			fmt.Printf("Completing transaction for session_id %v\n", response.SessionId)
			// The transaction is now done.
			close(tran.IsDone)
			tran.Done = true
		}
	}

}
