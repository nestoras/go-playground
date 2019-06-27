package main

import (
	"context"
	"fmt"
	"go-playground/grpc/08_deadline/newsfeedpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	fmt.Println("Unary with Deadline RPC...")
	request := &newsfeedpb.NewsfeedDeadlineRequest{
		Newsfeed: &newsfeedpb.Newsfeed{
			Text: "My Paris photo tour",
			Author:  "Nestoras Stefanou",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

	defer cancel()

	result, err := c.NewsfeedDeadline(ctx, request)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline was exceeded")
			} else {
				fmt.Printf("Error: %v", statusErr)
			}
		} else {
			log.Fatalf("Error: %v", err)
		}



		log.Fatalf("Error: %v", err)
	}
	log.Printf("Unary Response: %v", result.Result)

}
