package main

import (
	"context"
	"fmt"
	"go-playground/grpc/07_errors/newsfeedpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			Author: "",
		},
	}
	result, err := c.Newsfeed(context.Background(), request)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Printf("Error from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("No author!")
				//retry logic ....
				return
			}
		} else {
			log.Fatalf("Error: %v", err)
			return
		}
	}
	log.Printf("Unary Response: %v", result.Result)

}
