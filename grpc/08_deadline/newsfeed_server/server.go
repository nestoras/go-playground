package main

import (
	"context"
	"fmt"
	"go-playground/grpc/08_deadline/newsfeedpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

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

func (*server) NewsfeedDeadline(ctx context.Context, request *newsfeedpb.NewsfeedDeadlineRequest) (*newsfeedpb.NewsfeedDeadlineResponse, error) {
	fmt.Printf("Send: %v\n", request)
	for i := 0; i < 20; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(2 * time.Second)
	}
	author := request.GetNewsfeed().GetAuthor()
	text := request.GetNewsfeed().GetText()
	message := author + ": " +  text
	res := &newsfeedpb.NewsfeedDeadlineResponse{
		Result: message,
	}
	return res, nil
}
