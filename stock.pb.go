// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/stock/stock.proto

package stock

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f076fe61217ff59, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f076fe61217ff59, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f076fe61217ff59, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f076fe61217ff59, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "stock.Consignment")
	proto.RegisterType((*Container)(nil), "stock.Container")
	proto.RegisterType((*GetRequest)(nil), "stock.GetRequest")
	proto.RegisterType((*Response)(nil), "stock.Response")
}

func init() { proto.RegisterFile("proto/stock/stock.proto", fileDescriptor_5f076fe61217ff59) }

var fileDescriptor_5f076fe61217ff59 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4e, 0xe3, 0x30,
	0x14, 0x85, 0x27, 0xfd, 0xcf, 0x4d, 0x35, 0x9d, 0xde, 0xc5, 0x34, 0x9a, 0x59, 0x10, 0x65, 0xd5,
	0x55, 0x41, 0x05, 0x81, 0xc4, 0xb6, 0x8b, 0xaa, 0x5b, 0xf7, 0x01, 0x50, 0xb1, 0xaf, 0x52, 0x0b,
	0x6a, 0x07, 0xdb, 0x2d, 0x6f, 0xc0, 0x03, 0xf0, 0x10, 0x3c, 0x27, 0xaa, 0x9d, 0xd0, 0x20, 0xba,
	0x89, 0x74, 0xce, 0x3d, 0x27, 0xf9, 0x7c, 0x63, 0x98, 0x94, 0x46, 0x3b, 0x7d, 0x69, 0x9d, 0xe6,
	0x4f, 0xe1, 0x39, 0xf3, 0x0e, 0x76, 0xbd, 0xc8, 0x3f, 0x22, 0x48, 0x16, 0x5a, 0x59, 0x59, 0xa8,
	0x1d, 0x29, 0x87, 0xbf, 0xa1, 0x25, 0x45, 0x1a, 0x65, 0xd1, 0x34, 0x66, 0x2d, 0x29, 0x30, 0x83,
	0x44, 0x90, 0xe5, 0x46, 0x96, 0x4e, 0x6a, 0x95, 0xb6, 0xfc, 0xa0, 0x69, 0xe1, 0x5f, 0xe8, 0xbd,
	0x92, 0x2c, 0xb6, 0x2e, 0x6d, 0x67, 0xd1, 0xb4, 0xcb, 0x2a, 0x85, 0x57, 0x00, 0x5c, 0x2b, 0xb7,
	0x91, 0x8a, 0x8c, 0x4d, 0x3b, 0x59, 0x7b, 0x9a, 0xcc, 0xff, 0xcc, 0x02, 0xc2, 0xa2, 0x1e, 0xb0,
	0x46, 0x06, 0xff, 0x43, 0x7c, 0x20, 0x6b, 0xe9, 0xf9, 0x41, 0x8a, 0xb4, 0xeb, 0xbf, 0x34, 0x08,
	0xc6, 0x4a, 0xe4, 0x3b, 0x88, 0xbf, 0x5a, 0x3f, 0x28, 0x2f, 0x20, 0xe1, 0x7b, 0xeb, 0xf4, 0x8e,
	0xcc, 0xb1, 0x1b, 0x28, 0xa1, 0xb6, 0x56, 0xe2, 0x08, 0xa9, 0x8d, 0x2c, 0xa4, 0xf2, 0x90, 0x31,
	0xab, 0x14, 0x4e, 0xa0, 0xbf, 0xb7, 0xa1, 0xd4, 0x09, 0x83, 0xa3, 0x5c, 0x89, 0x7c, 0x08, 0xb0,
	0x24, 0xc7, 0xe8, 0x65, 0x4f, 0xd6, 0xe5, 0xef, 0x11, 0x0c, 0x18, 0xd9, 0x52, 0x2b, 0x4b, 0x98,
	0x42, 0x9f, 0x1b, 0xda, 0x38, 0x0a, 0x04, 0x03, 0x56, 0x4b, 0xbc, 0x81, 0x84, 0x9f, 0x76, 0xe9,
	0x31, 0x92, 0x39, 0x9e, 0xce, 0x5c, 0x4f, 0x58, 0x33, 0x86, 0xb7, 0x30, 0x6c, 0x48, 0x9b, 0xb6,
	0xfd, 0xaa, 0xce, 0xd5, 0xbe, 0xe5, 0xe6, 0x6f, 0x11, 0x8c, 0xd6, 0x5b, 0x59, 0x96, 0x52, 0x15,
	0x6b, 0x32, 0x07, 0xc9, 0x09, 0xef, 0x61, 0xbc, 0xf0, 0x30, 0xcd, 0x7f, 0x7a, 0xe6, 0x55, 0xff,
	0x46, 0x95, 0x57, 0x9f, 0x2a, 0xff, 0x85, 0x77, 0x30, 0x5a, 0x92, 0x6b, 0x84, 0x2c, 0x8e, 0xab,
	0xd4, 0x69, 0x15, 0x67, 0x8a, 0x8f, 0x3d, 0x7f, 0xa3, 0xae, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x37, 0xdc, 0x01, 0x6c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "stock"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}
