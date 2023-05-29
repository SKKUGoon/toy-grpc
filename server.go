package main

import (
	"log"
	"net"
	"srv-grpc/chat"

	"google.golang.org/grpc"
)

func main() {
	// Listener
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	// Define chat server. Recv, and send message back
	s := chat.Server{}

	// gRPC
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s) // Now exposing `SayHello` function
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
