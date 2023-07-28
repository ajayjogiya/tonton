package grpc

import (
	"context"
	"crypto/rand"
	"log"

	"github.com/ajayjogiya/tonton/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) Signup(ctx context.Context, in *proto.SignupRequest) (*proto.SignupResponse, error) {
	switch {
	case in.Username == "":
		log.Fatal("username should not be empty")
		return nil, status.Error(codes.Unauthenticated, "username should not be empty")
	case in.Password == "":
		log.Fatal("password should not be empty")
		return nil, status.Error(codes.Unauthenticated, "password should not be empty")
	}

	token := s.genToken(in.Username)
	s.setClientName(in.Username, token)

	return &proto.SignupResponse{Token: token}, nil
}

func (s *Server) openStream(stream proto.ChatService_StreamChatServer) {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Print("failed to read value from context")
		return
	}
	username := md.Get("user")

	if len(username) > 0 {
		s.mutex.Lock()
		s.ClientStreams[username[0]] = stream
		s.mutex.Unlock()
	} else {
		log.Print("there is no metadata")
	}

	log.Printf("stream opened for client: %s\n", username[0])
}

func (s *Server) setClientName(username, token string) {
	s.mutex.Lock()
	s.ClientNames[token] = username
	s.mutex.Unlock()

	log.Printf("saved token %s and username %s\n", token, username)
}

func (s *Server) genToken(username string) string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("failed to generate token: %+v\n", err)
		return ""
	}
	return username
}
