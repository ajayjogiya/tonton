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
	s.openStream(stream)
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

			s.send(stream, resp)
		}
	}()
	wg.Wait()

	return nil
}

func (s *Server) send(stream proto.ChatService_StreamChatServer, chat *proto.Chat) {
	to := chat.To
	if val, ok := s.ClientStreams[to]; ok {
		err := val.Send(chat)
		if err != nil {
			log.Fatalf("server failed send message: %+v\n", err)
			return
		}
	} else {
		log.Print("---------------------------------")
		log.Print("user not active")
		log.Print("---------------------------------")
		// Add logic to store messages
	}
}
