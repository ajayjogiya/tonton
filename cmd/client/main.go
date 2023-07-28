package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ajayjogiya/tonton/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

	ctx := context.Background()
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	signupRequest := &proto.SignupRequest{}
	fmt.Println("enter the username: ")
	if sc.Scan() {
		signupRequest.Username = sc.Text()
	}
	fmt.Println("and the password please: ")
	if sc.Scan() {
		signupRequest.Password = sc.Text()
	}

	token, err := c.Signup(ctx, signupRequest)
	if err != nil {
		log.Fatalf("error in signup: %+v\n", err)
		return
	}
	fmt.Printf("signup response: %+v\n", token)

	// add token instead of user
	md := metadata.New(map[string]string{
		"user": signupRequest.Username,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	client.streamChat(ctx)
}
