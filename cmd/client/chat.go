package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ajayjogiya/tonton/proto"
)

type client struct {
	proto.ChatServiceClient
}

func NewClient(c proto.ChatServiceClient) *client {
	return &client{c}
}

var wg sync.WaitGroup

func (c *client) streamChat(ctx context.Context) {
	log.Print("mayChat invoked")

	stream, err := c.StreamChat(ctx)
	if err != nil {
		log.Fatalf("fail to create client stream: %+v\n", err)
		return
	}
	wg.Add(2)

	go c.send(ctx, stream)
	go c.receive(ctx, stream)

	wg.Wait()
}

func (c *client) send(ctx context.Context, stream proto.ChatService_StreamChatClient) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	for {
		cm := &proto.Chat{}
		// should be print in better manner
		fmt.Println("Message: ")
		if sc.Scan() {
			cm.Message = sc.Text()
		}
		fmt.Println("To: ")
		if sc.Scan() {
			cm.To = sc.Text()
		}
		// Hacky way to close the stream
		if cm.Message == "exit" {
			break
		}
		err := stream.Send(cm)
		if err != nil {
			log.Fatalf("client failed to send message: %+v\n", err)
			return
		}
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("client failed to close the send stream: %+v\n", err)
		return
	}
	// close send and receive
	wg.Done()
	wg.Done()
}

func (c *client) receive(ctx context.Context, stream proto.ChatService_StreamChatClient) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("client failed to receive message: %+v\n", err)
			return
		}
		log.Printf("message received from %s, message: %s", resp.From, resp.Message)
	}

}
