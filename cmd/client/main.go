package main

import (
	"log"

	"github.com/ajayjogiya/tonton/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = "0.0.0.0:8000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %+v\n", err)
		return
	}
	defer conn.Close()

	c := proto.NewChatServiceClient(conn)
	client := NewClient(c)

	client.streamChat()
}
