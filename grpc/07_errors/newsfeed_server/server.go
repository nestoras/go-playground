package main

import (
	"context"
	"fmt"
	"go-playground/grpc/07_errors/newsfeedpb"
	"google.golang.org/grpc"
	"net"
	"log"

	// error handling
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if author == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received an empty author"),
		)
	}
	message := author + ": " +  text
	result := &newsfeedpb.NewsfeedResponse{
		Result: message,
	}
	return result, nil
}
