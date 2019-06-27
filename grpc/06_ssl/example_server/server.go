package main

import (
	"fmt"
	"go-playground/grpc/06_ssl/examplepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"

)

type server struct{}

func main() {
	fmt.Println("Hello gRPC world!")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	must(err)
	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	must(sslErr)

	s := grpc.NewServer(grpc.Creds(creds))

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
