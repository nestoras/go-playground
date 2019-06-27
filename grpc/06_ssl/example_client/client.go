package main

import (
	"fmt"
	"go-playground/grpc/01_basic_example/examplepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	fmt.Println("gRPC Client")
	certFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
		return
	}
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	must(err)

	defer cc.Close()

	client := examplepb.ExampleServiceClient(cc)
	fmt.Printf("Client: %f", client)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
