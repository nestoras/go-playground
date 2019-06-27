package main

import (
	"fmt"
	"go-playground/grpc/01_basic_example/examplepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	must(err)
	fmt.Println(err)

	defer cc.Close()

	client := examplepb.ExampleServiceClient(cc)
	fmt.Printf("Client: %f", client)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
