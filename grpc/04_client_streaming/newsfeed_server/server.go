package main

import (
	"fmt"
	"go-playground/grpc/04_client_streaming/newsfeedpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
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

func (*server) NewsfeedClient(stream newsfeedpb.NewsfeedService_NewsfeedClientServer) error {
	fmt.Printf("Received client request\n")
	result := ""
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&newsfeedpb.NewsfeedClientResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		userId := request.GetUser().GetUserId()
		result += "Data for user:"  + strconv.Itoa(int(userId)) + " "
	}
}
