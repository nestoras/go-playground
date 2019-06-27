package main

import (
	"fmt"
	"go-playground/grpc/05_bidi_streaming/newsfeedpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
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

func (*server) FindMaximumLike(stream newsfeedpb.NewsfeedService_FindMaximumLikeServer) error {
	fmt.Println("Received bidi request")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
			return err
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&newsfeedpb.FindMaximumLikeResponse{
				Maximum: maximum,
			})
			fmt.Println(maximum)
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", err)
				return err
			}
		}
	}
}

