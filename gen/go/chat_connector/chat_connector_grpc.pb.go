// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: chat_connector/chat_connector.proto

package chatv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatConnectorClient is the client API for ChatConnector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatConnectorClient interface {
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	ConnectChat(ctx context.Context, in *ConnectChatRequest, opts ...grpc.CallOption) (ChatConnector_ConnectChatClient, error)
	LeaveChat(ctx context.Context, in *LeaveChatRequest, opts ...grpc.CallOption) (*LeaveChatResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	CheckMember(ctx context.Context, in *CheckMemberRequest, opts ...grpc.CallOption) (*CheckMemberResponse, error)
}

type chatConnectorClient struct {
	cc grpc.ClientConnInterface
}

func NewChatConnectorClient(cc grpc.ClientConnInterface) ChatConnectorClient {
	return &chatConnectorClient{cc}
}

func (c *chatConnectorClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, "/chat_connector.ChatConnector/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatConnectorClient) ConnectChat(ctx context.Context, in *ConnectChatRequest, opts ...grpc.CallOption) (ChatConnector_ConnectChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatConnector_ServiceDesc.Streams[0], "/chat_connector.ChatConnector/ConnectChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatConnectorConnectChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatConnector_ConnectChatClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatConnectorConnectChatClient struct {
	grpc.ClientStream
}

func (x *chatConnectorConnectChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatConnectorClient) LeaveChat(ctx context.Context, in *LeaveChatRequest, opts ...grpc.CallOption) (*LeaveChatResponse, error) {
	out := new(LeaveChatResponse)
	err := c.cc.Invoke(ctx, "/chat_connector.ChatConnector/LeaveChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatConnectorClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/chat_connector.ChatConnector/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatConnectorClient) CheckMember(ctx context.Context, in *CheckMemberRequest, opts ...grpc.CallOption) (*CheckMemberResponse, error) {
	out := new(CheckMemberResponse)
	err := c.cc.Invoke(ctx, "/chat_connector.ChatConnector/CheckMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatConnectorServer is the server API for ChatConnector service.
// All implementations must embed UnimplementedChatConnectorServer
// for forward compatibility
type ChatConnectorServer interface {
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	ConnectChat(*ConnectChatRequest, ChatConnector_ConnectChatServer) error
	LeaveChat(context.Context, *LeaveChatRequest) (*LeaveChatResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	CheckMember(context.Context, *CheckMemberRequest) (*CheckMemberResponse, error)
	mustEmbedUnimplementedChatConnectorServer()
}

// UnimplementedChatConnectorServer must be embedded to have forward compatible implementations.
type UnimplementedChatConnectorServer struct {
}

func (UnimplementedChatConnectorServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatConnectorServer) ConnectChat(*ConnectChatRequest, ChatConnector_ConnectChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectChat not implemented")
}
func (UnimplementedChatConnectorServer) LeaveChat(context.Context, *LeaveChatRequest) (*LeaveChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChat not implemented")
}
func (UnimplementedChatConnectorServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatConnectorServer) CheckMember(context.Context, *CheckMemberRequest) (*CheckMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckMember not implemented")
}
func (UnimplementedChatConnectorServer) mustEmbedUnimplementedChatConnectorServer() {}

// UnsafeChatConnectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatConnectorServer will
// result in compilation errors.
type UnsafeChatConnectorServer interface {
	mustEmbedUnimplementedChatConnectorServer()
}

func RegisterChatConnectorServer(s grpc.ServiceRegistrar, srv ChatConnectorServer) {
	s.RegisterService(&ChatConnector_ServiceDesc, srv)
}

func _ChatConnector_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatConnectorServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_connector.ChatConnector/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatConnectorServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatConnector_ConnectChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatConnectorServer).ConnectChat(m, &chatConnectorConnectChatServer{stream})
}

type ChatConnector_ConnectChatServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatConnectorConnectChatServer struct {
	grpc.ServerStream
}

func (x *chatConnectorConnectChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatConnector_LeaveChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatConnectorServer).LeaveChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_connector.ChatConnector/LeaveChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatConnectorServer).LeaveChat(ctx, req.(*LeaveChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatConnector_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatConnectorServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_connector.ChatConnector/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatConnectorServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatConnector_CheckMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatConnectorServer).CheckMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_connector.ChatConnector/CheckMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatConnectorServer).CheckMember(ctx, req.(*CheckMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatConnector_ServiceDesc is the grpc.ServiceDesc for ChatConnector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatConnector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_connector.ChatConnector",
	HandlerType: (*ChatConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatConnector_CreateChat_Handler,
		},
		{
			MethodName: "LeaveChat",
			Handler:    _ChatConnector_LeaveChat_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatConnector_SendMessage_Handler,
		},
		{
			MethodName: "CheckMember",
			Handler:    _ChatConnector_CheckMember_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectChat",
			Handler:       _ChatConnector_ConnectChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat_connector/chat_connector.proto",
}
