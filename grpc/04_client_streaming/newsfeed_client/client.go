package main

import (
	"context"
	"fmt"
	"go-playground/grpc/04_client_streaming/newsfeedpb"
	"google.golang.org/grpc"
	"log"
	"time"
)



func main() {

	fmt.Println("Newsfeed client .....")
	opts := grpc.WithInsecure()

	grpc, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer grpc.Close()

	c := newsfeedpb.NewNewsfeedServiceClient(grpc)

	fmt.Println("Client Streaming RPC...")

	requests := []*newsfeedpb.NewsfeedClientRequest{
		&newsfeedpb.NewsfeedClientRequest{
			User: &newsfeedpb.User{
				UserId: 1,
			},
		},
		&newsfeedpb.NewsfeedClientRequest{
			User: &newsfeedpb.User{
				UserId: 12,
			},
		},
		&newsfeedpb.NewsfeedClientRequest{
			User: &newsfeedpb.User{
				UserId: 21,
			},
		},
		&newsfeedpb.NewsfeedClientRequest{
			User: &newsfeedpb.User{
				UserId: 2,
			},
		},
	}

	stream, err := c.NewsfeedClient(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(5000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Client Streaming Response: %v\n", res)
	}



