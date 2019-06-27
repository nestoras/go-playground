package main

import (
	"context"
	"fmt"
	"go-playground/grpc/02_unary/newsfeedpb"
	"google.golang.org/grpc"
	"net"
	"log"

)

type server struct{}

func main() {
	fmt.Println("Newsfeed server is ready...")

	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	options := []grpc.ServerOption{}

	s := grpc.NewServer(options...)
	newsfeedpb.RegisterNewsfeedServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (*server) Newsfeed(ctx context.Context, request *newsfeedpb.NewsfeedRequest) (*newsfeedpb.NewsfeedResponse, error) {
	fmt.Printf("Send: %v\n", request)
	author := request.GetNewsfeed().GetAuthor()
	text := request.GetNewsfeed().GetText()
	message := author + ": " +  text
	result := &newsfeedpb.NewsfeedResponse{
		Result: message,
	}
	return result, nil
}
