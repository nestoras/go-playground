package main

import (
	"context"
	"fmt"
	"go-playground/grpc/05_bidi_streaming/newsfeedpb"
	"google.golang.org/grpc"
	"io"
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

	fmt.Println("BiDi Streaming RPC...")

	stream, err := c.FindMaximumLike(context.Background())

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int32{8,2,1,5,7,22,35,64,302,4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			fmt.Printf("Sending number: %v\n", number)
			stream.Send(&newsfeedpb.FindMaximumLikeRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error: %v", err)
				break
			}
			maximum := result.GetMaximum()
			fmt.Printf("Received a new maximum number of...: %v\n", maximum)
		}
		close(waitc)
	}()
	<-waitc

	}



