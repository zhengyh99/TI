// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package message

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

type OrderRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	TimeStamp            int64    `protobuf:"varint,2,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderRequest) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type OrderInfo struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	OrderName            string   `protobuf:"bytes,2,opt,name=OrderName,proto3" json:"OrderName,omitempty"`
	OrderStatus          string   `protobuf:"bytes,3,opt,name=OrderStatus,proto3" json:"OrderStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderInfo) Reset()         { *m = OrderInfo{} }
func (m *OrderInfo) String() string { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()    {}
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *OrderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfo.Unmarshal(m, b)
}
func (m *OrderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfo.Marshal(b, m, deterministic)
}
func (m *OrderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfo.Merge(m, src)
}
func (m *OrderInfo) XXX_Size() int {
	return xxx_messageInfo_OrderInfo.Size(m)
}
func (m *OrderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfo proto.InternalMessageInfo

func (m *OrderInfo) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderInfo) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *OrderInfo) GetOrderStatus() string {
	if m != nil {
		return m.OrderStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderRequest)(nil), "message.OrderRequest")
	proto.RegisterType((*OrderInfo)(nil), "message.OrderInfo")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2f, 0x4a, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x55, 0x72, 0xe3, 0xe2, 0xf1, 0x07, 0x89, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49,
	0x70, 0xb1, 0x83, 0xf9, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90,
	0x0c, 0x17, 0x67, 0x48, 0x66, 0x6e, 0x6a, 0x70, 0x49, 0x62, 0x6e, 0x81, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x73, 0x10, 0x42, 0x40, 0x29, 0x95, 0x8b, 0x13, 0xa2, 0x30, 0x2f, 0x2d, 0x1f, 0xbf, 0x21,
	0x60, 0xa6, 0x5f, 0x62, 0x6e, 0x2a, 0xd8, 0x10, 0xce, 0x20, 0x84, 0x80, 0x90, 0x02, 0x17, 0x37,
	0x98, 0x13, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x2c, 0xc1, 0x0c, 0x96, 0x47, 0x16, 0x4a, 0x62, 0x03,
	0x3b, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x17, 0x9c, 0x1b, 0xcd, 0x00, 0x00, 0x00,
}