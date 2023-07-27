package main

import (
	"log"
	"net"

	g "github.com/ajayjogiya/tonton/pkg/grpc"
	"github.com/ajayjogiya/tonton/proto"
	"google.golang.org/grpc"
)

var address = "0.0.0.0:8000"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to listen on address: %s", address)
		return
	}
	log.Printf("listening on %s", address)

	s := grpc.NewServer()
	proto.RegisterChatServiceServer(s, &g.Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v\n", err)
	}

}
