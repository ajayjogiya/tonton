package grpc

import (
	"sync"

	"github.com/ajayjogiya/tonton/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.ChatServiceServer
	ClientNames   map[string]string
	ClientStreams map[string]proto.ChatService_StreamChatServer
	mutex         sync.Mutex
}

func NewChatServer() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterChatServiceServer(s, &Server{
		ClientNames:   make(map[string]string),
		ClientStreams: make(map[string]proto.ChatService_StreamChatServer),
	})
	return s
}
