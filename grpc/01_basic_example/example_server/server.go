package main

import (
	"fmt"
	"go-playground/grpc/01_basic_example/examplepb"
	"google.golang.org/grpc"
	"net"

)

type server struct{}

func main() {
	fmt.Println("Hello gRPC world!")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	must(err)

	s := grpc.NewServer()
	examplepb.RegisterExampleServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		must(err)
	}
}



func must(err error) {
	if err != nil {
		panic(err)
	}
}
