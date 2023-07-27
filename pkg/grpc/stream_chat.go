package grpc

import (
	"io"
	"log"
	"sync"

	"github.com/ajayjogiya/tonton/proto"
)

var wg sync.WaitGroup

func (s *Server) StreamChat(stream proto.ChatService_StreamChatServer) error {
	log.Printf("streamChat invoked")
	wg.Add(1)
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Println("EOF received, closing the stream")
					wg.Done()
					return
				}
				log.Fatal("failed to read client stream")
				return
			}
			log.Printf("message received from %s, message: %+v\n", resp.From, resp.Message)

			s.send(stream, &proto.Chat{Message: "Test message", From: "server", To: "client"})
		}
	}()
	wg.Wait()

	return nil
}

func (s *Server) send(stream proto.ChatService_StreamChatServer, chat *proto.Chat) {
	err := stream.Send(chat)
	if err != nil {
		log.Fatalf("server failed send message: %+v\n", err)
		return
	}
}
