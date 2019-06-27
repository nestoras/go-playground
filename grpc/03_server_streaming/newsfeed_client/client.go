package main

import (
	"context"
	"fmt"
	"go-playground/grpc/03_server_streaming/newsfeedpb"
	"google.golang.org/grpc"
	"io"
	"log"
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

	fmt.Println("Server Streaming RPC...")

	request := &newsfeedpb.NewsfeedServerRequest{
		User: &newsfeedpb.User{
			UserId: 1,
		},
	}

	resStream, err := c.NewsfeedServer(context.Background(),request)

	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error : %v", err)
		}
		log.Printf("Response : %v", msg.GetResult())
	}


}
