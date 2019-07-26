// Code generated by protoc-gen-go. DO NOT EDIT.
// source: traefikgrpc.proto

package proxy

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ProxyMeRequest struct {
	Req                  string   `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyMeRequest) Reset()         { *m = ProxyMeRequest{} }
func (m *ProxyMeRequest) String() string { return proto.CompactTextString(m) }
func (*ProxyMeRequest) ProtoMessage()    {}
func (*ProxyMeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8104e58e1a4f566, []int{0}
}

func (m *ProxyMeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyMeRequest.Unmarshal(m, b)
}
func (m *ProxyMeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyMeRequest.Marshal(b, m, deterministic)
}
func (m *ProxyMeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyMeRequest.Merge(m, src)
}
func (m *ProxyMeRequest) XXX_Size() int {
	return xxx_messageInfo_ProxyMeRequest.Size(m)
}
func (m *ProxyMeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyMeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyMeRequest proto.InternalMessageInfo

func (m *ProxyMeRequest) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

type ProxyMeResponse struct {
	Resp                 string   `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyMeResponse) Reset()         { *m = ProxyMeResponse{} }
func (m *ProxyMeResponse) String() string { return proto.CompactTextString(m) }
func (*ProxyMeResponse) ProtoMessage()    {}
func (*ProxyMeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8104e58e1a4f566, []int{1}
}

func (m *ProxyMeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyMeResponse.Unmarshal(m, b)
}
func (m *ProxyMeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyMeResponse.Marshal(b, m, deterministic)
}
func (m *ProxyMeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyMeResponse.Merge(m, src)
}
func (m *ProxyMeResponse) XXX_Size() int {
	return xxx_messageInfo_ProxyMeResponse.Size(m)
}
func (m *ProxyMeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyMeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyMeResponse proto.InternalMessageInfo

func (m *ProxyMeResponse) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

func init() {
	proto.RegisterType((*ProxyMeRequest)(nil), "proxy.ProxyMeRequest")
	proto.RegisterType((*ProxyMeResponse)(nil), "proxy.ProxyMeResponse")
}

func init() { proto.RegisterFile("traefikgrpc.proto", fileDescriptor_a8104e58e1a4f566) }

var fileDescriptor_a8104e58e1a4f566 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x29, 0x4a, 0x4c,
	0x4d, 0xcb, 0xcc, 0x4e, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d,
	0x28, 0xca, 0xaf, 0xa8, 0x54, 0x52, 0xe2, 0xe2, 0x0b, 0x00, 0x31, 0x7c, 0x53, 0x83, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0x8b, 0x52, 0x0b, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x40, 0x4c, 0x25, 0x55, 0x2e, 0x7e, 0xb8, 0x9a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x21, 0x2e, 0x96, 0xa2, 0xd4, 0xe2, 0x02, 0x09, 0x26, 0xb0, 0x2a, 0x30, 0xdb, 0xc8, 0x8f,
	0x4b, 0x20, 0x04, 0x62, 0x8d, 0x7b, 0x50, 0x80, 0x33, 0x58, 0x87, 0x90, 0x15, 0x17, 0x3b, 0x54,
	0xab, 0x90, 0xa8, 0x1e, 0xd8, 0x46, 0x3d, 0x54, 0xeb, 0xa4, 0xc4, 0xd0, 0x85, 0x21, 0x36, 0x28,
	0x31, 0x24, 0xb1, 0x81, 0x1d, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xce, 0x3b, 0x54, 0xff,
	0xbd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraefikGRPCProxyClient is the client API for TraefikGRPCProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraefikGRPCProxyClient interface {
	ProxyMe(ctx context.Context, in *ProxyMeRequest, opts ...grpc.CallOption) (*ProxyMeResponse, error)
}

type traefikGRPCProxyClient struct {
	cc *grpc.ClientConn
}

func NewTraefikGRPCProxyClient(cc *grpc.ClientConn) TraefikGRPCProxyClient {
	return &traefikGRPCProxyClient{cc}
}

func (c *traefikGRPCProxyClient) ProxyMe(ctx context.Context, in *ProxyMeRequest, opts ...grpc.CallOption) (*ProxyMeResponse, error) {
	out := new(ProxyMeResponse)
	err := c.cc.Invoke(ctx, "/proxy.TraefikGRPCProxy/ProxyMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraefikGRPCProxyServer is the server API for TraefikGRPCProxy service.
type TraefikGRPCProxyServer interface {
	ProxyMe(context.Context, *ProxyMeRequest) (*ProxyMeResponse, error)
}

// UnimplementedTraefikGRPCProxyServer can be embedded to have forward compatible implementations.
type UnimplementedTraefikGRPCProxyServer struct {
}

func (*UnimplementedTraefikGRPCProxyServer) ProxyMe(ctx context.Context, req *ProxyMeRequest) (*ProxyMeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProxyMe not implemented")
}

func RegisterTraefikGRPCProxyServer(s *grpc.Server, srv TraefikGRPCProxyServer) {
	s.RegisterService(&_TraefikGRPCProxy_serviceDesc, srv)
}

func _TraefikGRPCProxy_ProxyMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraefikGRPCProxyServer).ProxyMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.TraefikGRPCProxy/ProxyMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraefikGRPCProxyServer).ProxyMe(ctx, req.(*ProxyMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraefikGRPCProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.TraefikGRPCProxy",
	HandlerType: (*TraefikGRPCProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProxyMe",
			Handler:    _TraefikGRPCProxy_ProxyMe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traefikgrpc.proto",
}
