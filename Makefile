init:
	go mod download

server:
	go run ./cmd/server

client: 
	go run ./cmd/client

fmt:
	go fmt ./...