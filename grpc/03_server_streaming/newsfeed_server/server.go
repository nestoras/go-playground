package main

import (
	"fmt"
	"go-playground/grpc/03_server_streaming/newsfeedpb"
	"google.golang.org/grpc"
	"net"
	"log"
	"strconv"
	"time"
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

func (*server) NewsfeedServer(request *newsfeedpb.NewsfeedServerRequest, stream newsfeedpb.NewsfeedService_NewsfeedServerServer) error {
	fmt.Printf("Received server request : %v\n", request)
	userId := request.GetUser().GetUserId()
	fmt.Println(userId)
	for i := 0; i < 5; i++ {
		now := time.Now()
		message :=  "Messages for user: " +  strconv.Itoa(int(userId)) + " created at: " + now.String()
		result := &newsfeedpb.NewsfeedServerResponse{
			Result: message,
		}
		stream.Send(result)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}
