// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: chat.proto

package proto

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

const (
	ChatService_StreamChat_FullMethodName = "/tonton.ChatService/StreamChat"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	StreamChat(ctx context.Context, opts ...grpc.CallOption) (ChatService_StreamChatClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) StreamChat(ctx context.Context, opts ...grpc.CallOption) (ChatService_StreamChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_StreamChat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceStreamChatClient{stream}
	return x, nil
}

type ChatService_StreamChatClient interface {
	Send(*Chat) error
	Recv() (*Chat, error)
	grpc.ClientStream
}

type chatServiceStreamChatClient struct {
	grpc.ClientStream
}

func (x *chatServiceStreamChatClient) Send(m *Chat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceStreamChatClient) Recv() (*Chat, error) {
	m := new(Chat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	StreamChat(ChatService_StreamChatServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) StreamChat(ChatService_StreamChatServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamChat not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_StreamChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).StreamChat(&chatServiceStreamChatServer{stream})
}

type ChatService_StreamChatServer interface {
	Send(*Chat) error
	Recv() (*Chat, error)
	grpc.ServerStream
}

type chatServiceStreamChatServer struct {
	grpc.ServerStream
}

func (x *chatServiceStreamChatServer) Send(m *Chat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceStreamChatServer) Recv() (*Chat, error) {
	m := new(Chat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tonton.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChat",
			Handler:       _ChatService_StreamChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}