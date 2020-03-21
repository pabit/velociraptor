// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notebooks.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NotebookExportRequest struct {
	NotebookId           string   `protobuf:"bytes,1,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookExportRequest) Reset()         { *m = NotebookExportRequest{} }
func (m *NotebookExportRequest) String() string { return proto.CompactTextString(m) }
func (*NotebookExportRequest) ProtoMessage()    {}
func (*NotebookExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{0}
}

func (m *NotebookExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookExportRequest.Unmarshal(m, b)
}
func (m *NotebookExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookExportRequest.Marshal(b, m, deterministic)
}
func (m *NotebookExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookExportRequest.Merge(m, src)
}
func (m *NotebookExportRequest) XXX_Size() int {
	return xxx_messageInfo_NotebookExportRequest.Size(m)
}
func (m *NotebookExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookExportRequest proto.InternalMessageInfo

func (m *NotebookExportRequest) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

func (m *NotebookExportRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type NotebookCellRequest struct {
	NotebookId           string   `protobuf:"bytes,1,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	CellId               string   `protobuf:"bytes,2,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	Input                string   `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Offset               uint64   `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                uint64   `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	CurrentlyEditing     bool     `protobuf:"varint,8,opt,name=currently_editing,json=currentlyEditing,proto3" json:"currently_editing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookCellRequest) Reset()         { *m = NotebookCellRequest{} }
func (m *NotebookCellRequest) String() string { return proto.CompactTextString(m) }
func (*NotebookCellRequest) ProtoMessage()    {}
func (*NotebookCellRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{1}
}

func (m *NotebookCellRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookCellRequest.Unmarshal(m, b)
}
func (m *NotebookCellRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookCellRequest.Marshal(b, m, deterministic)
}
func (m *NotebookCellRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookCellRequest.Merge(m, src)
}
func (m *NotebookCellRequest) XXX_Size() int {
	return xxx_messageInfo_NotebookCellRequest.Size(m)
}
func (m *NotebookCellRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookCellRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookCellRequest proto.InternalMessageInfo

func (m *NotebookCellRequest) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

func (m *NotebookCellRequest) GetCellId() string {
	if m != nil {
		return m.CellId
	}
	return ""
}

func (m *NotebookCellRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *NotebookCellRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *NotebookCellRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *NotebookCellRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NotebookCellRequest) GetCurrentlyEditing() bool {
	if m != nil {
		return m.CurrentlyEditing
	}
	return false
}

type NotebookMetadata struct {
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Creator      string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	CreatedTime  int64  `protobuf:"varint,4,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	ModifiedTime int64  `protobuf:"varint,5,opt,name=modified_time,json=modifiedTime,proto3" json:"modified_time,omitempty"`
	NotebookId   string `protobuf:"bytes,7,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	// Deprecated
	Cells                []string            `protobuf:"bytes,6,rep,name=cells,proto3" json:"cells,omitempty"`
	CellMetadata         []*NotebookCell     `protobuf:"bytes,11,rep,name=cell_metadata,json=cellMetadata,proto3" json:"cell_metadata,omitempty"`
	LatestCellId         string              `protobuf:"bytes,8,opt,name=latest_cell_id,json=latestCellId,proto3" json:"latest_cell_id,omitempty"`
	Hidden               bool                `protobuf:"varint,9,opt,name=hidden,proto3" json:"hidden,omitempty"`
	AvailableDownloads   *AvailableDownloads `protobuf:"bytes,10,opt,name=available_downloads,json=availableDownloads,proto3" json:"available_downloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NotebookMetadata) Reset()         { *m = NotebookMetadata{} }
func (m *NotebookMetadata) String() string { return proto.CompactTextString(m) }
func (*NotebookMetadata) ProtoMessage()    {}
func (*NotebookMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{2}
}

func (m *NotebookMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookMetadata.Unmarshal(m, b)
}
func (m *NotebookMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookMetadata.Marshal(b, m, deterministic)
}
func (m *NotebookMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookMetadata.Merge(m, src)
}
func (m *NotebookMetadata) XXX_Size() int {
	return xxx_messageInfo_NotebookMetadata.Size(m)
}
func (m *NotebookMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookMetadata proto.InternalMessageInfo

func (m *NotebookMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NotebookMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NotebookMetadata) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *NotebookMetadata) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *NotebookMetadata) GetModifiedTime() int64 {
	if m != nil {
		return m.ModifiedTime
	}
	return 0
}

func (m *NotebookMetadata) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

func (m *NotebookMetadata) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *NotebookMetadata) GetCellMetadata() []*NotebookCell {
	if m != nil {
		return m.CellMetadata
	}
	return nil
}

func (m *NotebookMetadata) GetLatestCellId() string {
	if m != nil {
		return m.LatestCellId
	}
	return ""
}

func (m *NotebookMetadata) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *NotebookMetadata) GetAvailableDownloads() *AvailableDownloads {
	if m != nil {
		return m.AvailableDownloads
	}
	return nil
}

type Notebooks struct {
	Items                []*NotebookMetadata `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Notebooks) Reset()         { *m = Notebooks{} }
func (m *Notebooks) String() string { return proto.CompactTextString(m) }
func (*Notebooks) ProtoMessage()    {}
func (*Notebooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{3}
}

func (m *Notebooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notebooks.Unmarshal(m, b)
}
func (m *Notebooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notebooks.Marshal(b, m, deterministic)
}
func (m *Notebooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notebooks.Merge(m, src)
}
func (m *Notebooks) XXX_Size() int {
	return xxx_messageInfo_Notebooks.Size(m)
}
func (m *Notebooks) XXX_DiscardUnknown() {
	xxx_messageInfo_Notebooks.DiscardUnknown(m)
}

var xxx_messageInfo_Notebooks proto.InternalMessageInfo

func (m *Notebooks) GetItems() []*NotebookMetadata {
	if m != nil {
		return m.Items
	}
	return nil
}

type NotebookCell struct {
	Input     string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output    string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Data      string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	CellId    string   `protobuf:"bytes,4,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	Messages  []string `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty"`
	Timestamp int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The type of this cell.
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	CurrentlyEditing     bool     `protobuf:"varint,8,opt,name=currently_editing,json=currentlyEditing,proto3" json:"currently_editing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookCell) Reset()         { *m = NotebookCell{} }
func (m *NotebookCell) String() string { return proto.CompactTextString(m) }
func (*NotebookCell) ProtoMessage()    {}
func (*NotebookCell) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{4}
}

func (m *NotebookCell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookCell.Unmarshal(m, b)
}
func (m *NotebookCell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookCell.Marshal(b, m, deterministic)
}
func (m *NotebookCell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookCell.Merge(m, src)
}
func (m *NotebookCell) XXX_Size() int {
	return xxx_messageInfo_NotebookCell.Size(m)
}
func (m *NotebookCell) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookCell.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookCell proto.InternalMessageInfo

func (m *NotebookCell) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *NotebookCell) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *NotebookCell) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *NotebookCell) GetCellId() string {
	if m != nil {
		return m.CellId
	}
	return ""
}

func (m *NotebookCell) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *NotebookCell) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *NotebookCell) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NotebookCell) GetCurrentlyEditing() bool {
	if m != nil {
		return m.CurrentlyEditing
	}
	return false
}

func init() {
	proto.RegisterType((*NotebookExportRequest)(nil), "proto.NotebookExportRequest")
	proto.RegisterType((*NotebookCellRequest)(nil), "proto.NotebookCellRequest")
	proto.RegisterType((*NotebookMetadata)(nil), "proto.NotebookMetadata")
	proto.RegisterType((*Notebooks)(nil), "proto.Notebooks")
	proto.RegisterType((*NotebookCell)(nil), "proto.NotebookCell")
}

func init() { proto.RegisterFile("notebooks.proto", fileDescriptor_6a2fb83cea3dee99) }

var fileDescriptor_6a2fb83cea3dee99 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x6f, 0xd3, 0x4c,
	0x14, 0x94, 0x3f, 0xc7, 0x4e, 0xf2, 0x9c, 0x7e, 0x94, 0x0d, 0xb4, 0x4b, 0x85, 0x84, 0x09, 0x1c,
	0x22, 0x21, 0x7a, 0x28, 0x17, 0xc4, 0x0d, 0x95, 0x1e, 0x8a, 0x80, 0x83, 0xc5, 0x3d, 0xda, 0x64,
	0x5f, 0xca, 0x8a, 0xb5, 0xd7, 0x78, 0x37, 0x94, 0xfe, 0x20, 0xfe, 0x11, 0x3f, 0x81, 0x1f, 0x82,
	0xfc, 0xbc, 0xeb, 0xa4, 0xe9, 0x05, 0x71, 0xca, 0xce, 0xec, 0x64, 0xfd, 0x66, 0xde, 0xc0, 0xbd,
	0xca, 0x38, 0x5c, 0x1a, 0xf3, 0xd5, 0x9e, 0xd6, 0x8d, 0x71, 0x86, 0x25, 0xf4, 0x73, 0x92, 0xad,
	0xb5, 0xb9, 0xf6, 0xdc, 0xec, 0x03, 0x3c, 0xfc, 0xe4, 0x65, 0x17, 0x3f, 0x6a, 0xd3, 0xb8, 0x02,
	0xbf, 0x6d, 0xd0, 0x3a, 0xf6, 0x04, 0xb2, 0xf0, 0xff, 0x85, 0x92, 0x3c, 0xca, 0xa3, 0xf9, 0xb8,
	0x80, 0x40, 0x5d, 0x4a, 0xc6, 0x60, 0xe0, 0x6e, 0x6a, 0xe4, 0xff, 0xd1, 0x0d, 0x9d, 0x67, 0xbf,
	0x22, 0x98, 0x86, 0xe7, 0xce, 0x51, 0xeb, 0xbf, 0x7e, 0xec, 0x18, 0x86, 0x2b, 0xd4, 0xba, 0xbd,
	0xec, 0xde, 0x4b, 0x5b, 0x78, 0x29, 0xd9, 0x03, 0x48, 0x54, 0x55, 0x6f, 0x1c, 0x8f, 0x89, 0xee,
	0x00, 0x3b, 0x82, 0xd4, 0xac, 0xd7, 0x16, 0x1d, 0x1f, 0xe4, 0xd1, 0x7c, 0x50, 0x78, 0xd4, 0xaa,
	0x57, 0x66, 0x53, 0x39, 0x9e, 0x10, 0xdd, 0x81, 0x7e, 0xd2, 0x74, 0x3b, 0x29, 0x7b, 0x01, 0xf7,
	0x57, 0x9b, 0xa6, 0xc1, 0xca, 0xe9, 0x9b, 0x05, 0x4a, 0xe5, 0x54, 0x75, 0xc5, 0x47, 0x79, 0x34,
	0x1f, 0x15, 0x87, 0xfd, 0xc5, 0x45, 0xc7, 0xcf, 0x7e, 0xc6, 0x70, 0x18, 0x6c, 0x7d, 0x44, 0x27,
	0xa4, 0x70, 0xa2, 0x7d, 0xb5, 0x12, 0x25, 0x7a, 0x33, 0x74, 0x66, 0x39, 0x64, 0x12, 0xed, 0xaa,
	0x51, 0xb5, 0x53, 0xa6, 0xf2, 0x56, 0x76, 0x29, 0xc6, 0x61, 0xb8, 0x6a, 0x50, 0x38, 0xd3, 0x78,
	0x47, 0x01, 0xb2, 0xa7, 0x30, 0xa1, 0x23, 0xca, 0x85, 0x53, 0x25, 0x92, 0xb3, 0xb8, 0xc8, 0x3c,
	0xf7, 0x59, 0x95, 0xc8, 0x9e, 0xc1, 0x41, 0x69, 0xa4, 0x5a, 0xab, 0xa0, 0x49, 0x48, 0x33, 0x09,
	0x24, 0x89, 0xf6, 0xb2, 0x1e, 0xde, 0xc9, 0xba, 0x0d, 0x09, 0xb5, 0xb6, 0x3c, 0xcd, 0xe3, 0x36,
	0x52, 0x02, 0xec, 0x35, 0x1c, 0xd0, 0x06, 0x4a, 0xef, 0x8f, 0x67, 0x79, 0x3c, 0xcf, 0xce, 0xa6,
	0x5d, 0x4f, 0x4e, 0x6f, 0x6d, 0x75, 0xd2, 0x2a, 0xfb, 0x20, 0x9e, 0xc3, 0xff, 0x5a, 0x38, 0xb4,
	0x6e, 0x11, 0x56, 0x38, 0xa2, 0x6f, 0x4e, 0x3a, 0xf6, 0xbc, 0x5b, 0xe4, 0x11, 0xa4, 0x5f, 0x94,
	0x94, 0x58, 0xf1, 0x31, 0xa5, 0xec, 0x11, 0x7b, 0x0f, 0x53, 0xf1, 0x5d, 0x28, 0x2d, 0x96, 0x1a,
	0x17, 0xd2, 0x5c, 0x57, 0xda, 0x08, 0x69, 0x39, 0xe4, 0xd1, 0x3c, 0x3b, 0x7b, 0xe4, 0xbf, 0xfe,
	0x36, 0x28, 0xde, 0x05, 0x41, 0xc1, 0xc4, 0x1d, 0x6e, 0xf6, 0x06, 0xc6, 0x61, 0x4e, 0xcb, 0x5e,
	0x42, 0xa2, 0x1c, 0x96, 0x96, 0x47, 0x64, 0xe4, 0x78, 0xcf, 0x48, 0x18, 0xbf, 0xe8, 0x54, 0xb3,
	0xdf, 0x11, 0x4c, 0x76, 0x4d, 0x6e, 0x9b, 0x17, 0xed, 0x37, 0x6f, 0xe3, 0x5a, 0xda, 0xf7, 0xb4,
	0x43, 0x6d, 0x1b, 0x28, 0xb5, 0x6e, 0xa9, 0x74, 0xde, 0x2d, 0xf5, 0xe0, 0x56, 0xa9, 0x4f, 0x60,
	0x54, 0xa2, 0xb5, 0xe2, 0x0a, 0x2d, 0x4f, 0x68, 0x09, 0x3d, 0x66, 0x8f, 0x61, 0xdc, 0xae, 0xd6,
	0x3a, 0x51, 0xd6, 0xd4, 0xd8, 0xb8, 0xd8, 0x12, 0x7d, 0x95, 0x87, 0xff, 0x58, 0xe5, 0x65, 0x4a,
	0x29, 0xbc, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x5b, 0x98, 0x0e, 0x1d, 0x04, 0x00, 0x00,
}