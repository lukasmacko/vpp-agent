// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: value_status.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ValueState int32

const (
	// ValueState_NONEXISTENT is assigned to value that was deleted or has never
	// existed.
	ValueState_NONEXISTENT ValueState = 0
	// ValueState_MISSING is assigned to NB value that was configured but refresh
	// found it to be missing.
	ValueState_MISSING ValueState = 1
	// ValueState_UNIMPLEMENTED marks value received from NB that cannot
	// be configured because there is no registered descriptor associated
	// with it.
	ValueState_UNIMPLEMENTED ValueState = 2
	// ValueState_REMOVED is assigned to NB value after it was removed or when
	// it is being re-created. The state is only temporary: for re-create, the
	// value transits to whatever state the following Create operation produces,
	// and delete values are removed from the graph (go to the NONEXISTENT state)
	// immediately after the notification about the state change is sent.
	ValueState_REMOVED ValueState = 3
	// ValueState_CONFIGURED marks value defined by NB and successfully configured.
	ValueState_CONFIGURED ValueState = 4
	// ValueState_RETRIEVED marks value not managed by NB, instead created
	// automatically or externally in SB. The KVScheduler learns about the value
	// either using Retrieve() or through a SB notification.
	ValueState_RETRIEVED ValueState = 5
	// ValueState_FOUND marks NB value that was found (=retrieved) by refresh
	// but not actually configured by the agent.
	ValueState_FOUND ValueState = 6
	// ValueState_PENDING represents (NB) value that cannot be configured yet
	// due to missing dependencies.
	ValueState_PENDING ValueState = 7
	// ValueState_INVALID represents (NB) value that will not be configured
	// because it has a logically invalid content as declared by the Validate
	// method of the associated descriptor.
	// The corresponding error and the list of affected fields are stored
	// in the <InvalidValueDetails> structure available via <details> for invalid
	// value.
	ValueState_INVALID ValueState = 8
	// ValueState_FAILED marks (NB) value for which the last executed operation
	// returned an error.
	// The error and the type of the operation which caused the error are stored
	// in the <FailedValueDetails> structure available via <details> for failed
	// value.
	ValueState_FAILED ValueState = 9
	// ValueState_RETRYING marks unsucessfully applied (NB) value, for which,
	// however, one or more attempts to fix the error by repeating the last
	// operation are planned, and only if all the retries fail, the value will
	// then transit to the FAILED state.
	ValueState_RETRYING ValueState = 10
)

var ValueState_name = map[int32]string{
	0:  "NONEXISTENT",
	1:  "MISSING",
	2:  "UNIMPLEMENTED",
	3:  "REMOVED",
	4:  "CONFIGURED",
	5:  "RETRIEVED",
	6:  "FOUND",
	7:  "PENDING",
	8:  "INVALID",
	9:  "FAILED",
	10: "RETRYING",
}
var ValueState_value = map[string]int32{
	"NONEXISTENT":   0,
	"MISSING":       1,
	"UNIMPLEMENTED": 2,
	"REMOVED":       3,
	"CONFIGURED":    4,
	"RETRIEVED":     5,
	"FOUND":         6,
	"PENDING":       7,
	"INVALID":       8,
	"FAILED":        9,
	"RETRYING":      10,
}

func (x ValueState) String() string {
	return proto.EnumName(ValueState_name, int32(x))
}
func (ValueState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_value_status_5c67a50a9223cffc, []int{0}
}

type TxnOperation int32

const (
	TxnOperation_UNDEFINED TxnOperation = 0
	TxnOperation_VALIDATE  TxnOperation = 1
	TxnOperation_CREATE    TxnOperation = 2
	TxnOperation_UPDATE    TxnOperation = 3
	TxnOperation_DELETE    TxnOperation = 4
)

var TxnOperation_name = map[int32]string{
	0: "UNDEFINED",
	1: "VALIDATE",
	2: "CREATE",
	3: "UPDATE",
	4: "DELETE",
}
var TxnOperation_value = map[string]int32{
	"UNDEFINED": 0,
	"VALIDATE":  1,
	"CREATE":    2,
	"UPDATE":    3,
	"DELETE":    4,
}

func (x TxnOperation) String() string {
	return proto.EnumName(TxnOperation_name, int32(x))
}
func (TxnOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_value_status_5c67a50a9223cffc, []int{1}
}

type ValueStatus struct {
	Key           string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	State         ValueState   `protobuf:"varint,2,opt,name=state,proto3,enum=api.ValueState" json:"state,omitempty"`
	Error         string       `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	LastOperation TxnOperation `protobuf:"varint,4,opt,name=last_operation,json=lastOperation,proto3,enum=api.TxnOperation" json:"last_operation,omitempty"`
	// - for invalid value, details is a list of invalid fields
	// - for pending value, details is a list of missing dependencies (labels)
	Details              []string `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueStatus) Reset()         { *m = ValueStatus{} }
func (m *ValueStatus) String() string { return proto.CompactTextString(m) }
func (*ValueStatus) ProtoMessage()    {}
func (*ValueStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_status_5c67a50a9223cffc, []int{0}
}
func (m *ValueStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueStatus.Unmarshal(m, b)
}
func (m *ValueStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueStatus.Marshal(b, m, deterministic)
}
func (dst *ValueStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueStatus.Merge(dst, src)
}
func (m *ValueStatus) XXX_Size() int {
	return xxx_messageInfo_ValueStatus.Size(m)
}
func (m *ValueStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ValueStatus proto.InternalMessageInfo

func (m *ValueStatus) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ValueStatus) GetState() ValueState {
	if m != nil {
		return m.State
	}
	return ValueState_NONEXISTENT
}

func (m *ValueStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ValueStatus) GetLastOperation() TxnOperation {
	if m != nil {
		return m.LastOperation
	}
	return TxnOperation_UNDEFINED
}

func (m *ValueStatus) GetDetails() []string {
	if m != nil {
		return m.Details
	}
	return nil
}

type BaseValueStatus struct {
	Value                *ValueStatus   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	DerivedValues        []*ValueStatus `protobuf:"bytes,2,rep,name=derived_values,json=derivedValues,proto3" json:"derived_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BaseValueStatus) Reset()         { *m = BaseValueStatus{} }
func (m *BaseValueStatus) String() string { return proto.CompactTextString(m) }
func (*BaseValueStatus) ProtoMessage()    {}
func (*BaseValueStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_status_5c67a50a9223cffc, []int{1}
}
func (m *BaseValueStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseValueStatus.Unmarshal(m, b)
}
func (m *BaseValueStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseValueStatus.Marshal(b, m, deterministic)
}
func (dst *BaseValueStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseValueStatus.Merge(dst, src)
}
func (m *BaseValueStatus) XXX_Size() int {
	return xxx_messageInfo_BaseValueStatus.Size(m)
}
func (m *BaseValueStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseValueStatus.DiscardUnknown(m)
}

var xxx_messageInfo_BaseValueStatus proto.InternalMessageInfo

func (m *BaseValueStatus) GetValue() *ValueStatus {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *BaseValueStatus) GetDerivedValues() []*ValueStatus {
	if m != nil {
		return m.DerivedValues
	}
	return nil
}

func init() {
	proto.RegisterType((*ValueStatus)(nil), "api.ValueStatus")
	proto.RegisterType((*BaseValueStatus)(nil), "api.BaseValueStatus")
	proto.RegisterEnum("api.ValueState", ValueState_name, ValueState_value)
	proto.RegisterEnum("api.TxnOperation", TxnOperation_name, TxnOperation_value)
}

func init() { proto.RegisterFile("value_status.proto", fileDescriptor_value_status_5c67a50a9223cffc) }

var fileDescriptor_value_status_5c67a50a9223cffc = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x8d, 0x31, 0x86, 0x78, 0x08, 0xb0, 0x19, 0xf5, 0xe0, 0x23, 0x8a, 0xd4, 0x0a, 0xe5, 0xc0,
	0x21, 0x3d, 0xb4, 0x57, 0x9a, 0x1d, 0xa2, 0x95, 0x60, 0x8d, 0x16, 0x1b, 0xb5, 0x27, 0xb4, 0x15,
	0x7b, 0xb0, 0x8a, 0x62, 0xe4, 0x35, 0x51, 0xfb, 0x8f, 0x7a, 0xec, 0x4f, 0xac, 0x66, 0x69, 0xd2,
	0x54, 0xea, 0xed, 0xbd, 0x79, 0x1f, 0x9a, 0x59, 0x2d, 0xe0, 0x93, 0x3d, 0x9c, 0xdc, 0xce, 0xb7,
	0xb6, 0x3d, 0xf9, 0xd9, 0xb1, 0xa9, 0xdb, 0x1a, 0x63, 0x7b, 0xac, 0x6e, 0x7e, 0x45, 0x30, 0xd8,
	0xb2, 0xb6, 0x09, 0x12, 0x0a, 0x88, 0xbf, 0xb9, 0x1f, 0x59, 0x34, 0x89, 0xa6, 0xa9, 0x61, 0x88,
	0x6f, 0x21, 0xe1, 0x98, 0xcb, 0x3a, 0x93, 0x68, 0x3a, 0xba, 0x1b, 0xcf, 0xec, 0xb1, 0x9a, 0xbd,
	0x44, 0x9c, 0x39, 0xab, 0xf8, 0x06, 0x12, 0xd7, 0x34, 0x75, 0x93, 0xc5, 0x21, 0x7a, 0x26, 0xf8,
	0x11, 0x46, 0x07, 0xeb, 0xdb, 0x5d, 0x7d, 0x74, 0x8d, 0x6d, 0xab, 0xfa, 0x31, 0xeb, 0x86, 0x96,
	0xeb, 0xd0, 0x52, 0x7c, 0x7f, 0xcc, 0x9f, 0x05, 0x33, 0x64, 0xe3, 0x0b, 0xc5, 0x0c, 0xfa, 0x7b,
	0xd7, 0xda, 0xea, 0xe0, 0xb3, 0x64, 0x12, 0x4f, 0x53, 0xf3, 0x4c, 0x6f, 0x1a, 0x18, 0x7f, 0xb2,
	0xde, 0xbd, 0xde, 0xfa, 0x1d, 0x24, 0xe1, 0xc0, 0xb0, 0xf7, 0xe0, 0x4e, 0xfc, 0xbb, 0xe3, 0xc9,
	0x9b, 0xb3, 0x8c, 0x1f, 0x60, 0xb4, 0x77, 0x4d, 0xf5, 0xe4, 0xf6, 0xbb, 0x30, 0xf0, 0x59, 0x67,
	0x12, 0xff, 0x37, 0x30, 0xfc, 0xe3, 0x0b, 0x33, 0x7f, 0xfb, 0x33, 0x02, 0xf8, 0x7b, 0x33, 0x8e,
	0x61, 0xa0, 0x73, 0x4d, 0x9f, 0xd5, 0xa6, 0x20, 0x5d, 0x88, 0x0b, 0x1c, 0x40, 0x7f, 0xa5, 0x36,
	0x1b, 0xa5, 0x1f, 0x44, 0x84, 0xd7, 0x30, 0x2c, 0xb5, 0x5a, 0xad, 0x97, 0xb4, 0x22, 0x5d, 0x90,
	0x14, 0x1d, 0xd6, 0x0d, 0xad, 0xf2, 0x2d, 0x49, 0x11, 0xe3, 0x08, 0xe0, 0x3e, 0xd7, 0x0b, 0xf5,
	0x50, 0x1a, 0x92, 0xa2, 0x8b, 0x43, 0x48, 0x0d, 0x15, 0x46, 0x11, 0xcb, 0x09, 0xa6, 0x90, 0x2c,
	0xf2, 0x52, 0x4b, 0xd1, 0xe3, 0xd8, 0x9a, 0xb4, 0xe4, 0xda, 0x3e, 0x13, 0xa5, 0xb7, 0xf3, 0xa5,
	0x92, 0xe2, 0x12, 0x01, 0x7a, 0x8b, 0xb9, 0x5a, 0x92, 0x14, 0x29, 0x5e, 0xc1, 0x25, 0xe7, 0xbf,
	0xb0, 0x0d, 0x6e, 0x73, 0xb8, 0x7a, 0xfd, 0xae, 0xdc, 0x5e, 0x6a, 0x49, 0x0b, 0xa5, 0x49, 0x8a,
	0x0b, 0x36, 0x87, 0x8e, 0x79, 0x41, 0x22, 0xe2, 0x9a, 0x7b, 0x43, 0x8c, 0x3b, 0x8c, 0xcb, 0x75,
	0x98, 0xc7, 0x8c, 0x25, 0x2d, 0xa9, 0x20, 0xd1, 0xfd, 0xda, 0x0b, 0xdf, 0xe5, 0xfd, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x82, 0x3c, 0xf4, 0x44, 0x02, 0x00, 0x00,
}