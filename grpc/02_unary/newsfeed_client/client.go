package main

import (
	"context"
	"fmt"
	"go-playground/grpc/02_unary/newsfeedpb"
	"google.golang.org/grpc"
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

	fmt.Println("Unary RPC...")
	request := &newsfeedpb.NewsfeedRequest{
		Newsfeed: &newsfeedpb.Newsfeed{
			Text: "My Paris photo tour",
			Author:  "Nestoras Stefanou",
		},
	}
	result, err := c.Newsfeed(context.Background(), request)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Unary Response: %v", result.Result)

}
