package grpc

import (
	"github.com/ajayjogiya/tonton/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.ChatServiceServer
}

func NewChatServer() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterChatServiceServer(s, &Server{})
	return s
}
