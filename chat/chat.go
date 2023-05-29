package chat

import (
	"context"
	"log"
)

type Server struct {
	// From new protoc-version, one must implement this code to ensure Forward compatibility.
	// https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s\n", message.Body)
	return &Message{Body: "Hello from the server!"}, nil
}
