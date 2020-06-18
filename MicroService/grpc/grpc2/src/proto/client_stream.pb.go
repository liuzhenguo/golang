// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_stream.proto

package proto

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

// 定义发送请求信息
type SimpleRequest struct {
	// 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
	// 参数类型 参数名 标识号(不可重复)
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_729fd3a5e8f9e1ff, []int{0}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

func (m *SimpleRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// 定义响应信息
type SimpleResponse struct {
	// 定义接收的参数
	// 参数类型 参数名 标识号(不可重复)
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_729fd3a5e8f9e1ff, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SimpleResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// 定义流式请求信息
type StreamRequest struct {
	//流式请求参数
	StreamData           string   `protobuf:"bytes,1,opt,name=stream_data,json=streamData,proto3" json:"stream_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_729fd3a5e8f9e1ff, []int{2}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetStreamData() string {
	if m != nil {
		return m.StreamData
	}
	return ""
}

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "proto.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "proto.SimpleResponse")
	proto.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
}

func init() { proto.RegisterFile("client_stream.proto", fileDescriptor_729fd3a5e8f9e1ff) }

var fileDescriptor_729fd3a5e8f9e1ff = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x89, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x53, 0x4a, 0xca, 0x5c, 0xbc, 0xc1, 0x99, 0xb9, 0x05, 0x39, 0xa9, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x15, 0x17, 0x1f, 0x4c, 0x51, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0x2a, 0x48, 0x55, 0x72, 0x7e, 0x4a, 0x2a, 0x58, 0x15, 0x6b, 0x10, 0x98, 0x2d, 0x24, 0xc2, 0xc5,
	0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0xd6, 0x0a, 0xe1, 0x28, 0x19, 0x70, 0xf1, 0x06,
	0x83, 0xed, 0x85, 0x59, 0x20, 0xcf, 0xc5, 0x0d, 0x71, 0x48, 0x3c, 0x92, 0x3d, 0x5c, 0x10, 0x21,
	0x97, 0xc4, 0x92, 0x44, 0xa3, 0x16, 0x46, 0x2e, 0x1e, 0x88, 0x16, 0x67, 0xb0, 0xbb, 0x85, 0xcc,
	0xb8, 0x58, 0x83, 0xf2, 0x4b, 0x4b, 0x52, 0x85, 0x44, 0x20, 0x6e, 0xd7, 0x43, 0x71, 0xb1, 0x94,
	0x28, 0x9a, 0x28, 0xc4, 0x89, 0x4a, 0x0c, 0x42, 0x36, 0x5c, 0x9c, 0x60, 0x7d, 0x3e, 0x99, 0xc5,
	0x25, 0x08, 0xbd, 0xc8, 0x8e, 0xc1, 0xa9, 0x57, 0x83, 0x31, 0x89, 0x0d, 0x2c, 0x63, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x03, 0x69, 0x4b, 0xd8, 0x3e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamClientClient is the client API for StreamClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClientClient interface {
	Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// 客户端流式rpc，在请求的参数前添加stream
	RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error)
}

type streamClientClient struct {
	cc *grpc.ClientConn
}

func NewStreamClientClient(cc *grpc.ClientConn) StreamClientClient {
	return &streamClientClient{cc}
}

func (c *streamClientClient) Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/proto.StreamClient/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClientClient) RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamClient_serviceDesc.Streams[0], "/proto.StreamClient/RouteList", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamClientRouteListClient{stream}
	return x, nil
}

type StreamClient_RouteListClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type streamClientRouteListClient struct {
	grpc.ClientStream
}

func (x *streamClientRouteListClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamClientRouteListClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamClientServer is the server API for StreamClient service.
type StreamClientServer interface {
	Route(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// 客户端流式rpc，在请求的参数前添加stream
	RouteList(StreamClient_RouteListServer) error
}

// UnimplementedStreamClientServer can be embedded to have forward compatible implementations.
type UnimplementedStreamClientServer struct {
}

func (*UnimplementedStreamClientServer) Route(ctx context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (*UnimplementedStreamClientServer) RouteList(srv StreamClient_RouteListServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteList not implemented")
}

func RegisterStreamClientServer(s *grpc.Server, srv StreamClientServer) {
	s.RegisterService(&_StreamClient_serviceDesc, srv)
}

func _StreamClient_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamClientServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StreamClient/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamClientServer).Route(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamClient_RouteList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamClientServer).RouteList(&streamClientRouteListServer{stream})
}

type StreamClient_RouteListServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamClientRouteListServer struct {
	grpc.ServerStream
}

func (x *streamClientRouteListServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamClientRouteListServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamClient",
	HandlerType: (*StreamClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _StreamClient_Route_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteList",
			Handler:       _StreamClient_RouteList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client_stream.proto",
}
