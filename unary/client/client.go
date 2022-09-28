package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smith-golang/grpc-test/unary/unarypb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello i am a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect : %v", err)
	}

	defer conn.Close()

	c := unarypb.NewGreetingServiceClient(conn)
	// fmt.Println("Created client : %f", c)

	doUnary(c)
}

func doUnary(c unarypb.GreetingServiceClient) {
	fmt.Println("Starting to do a Unary RPC ...")
	req := &unarypb.GreetRequest{
		Greeting: &unarypb.Greeting{
			FirstName: "Smith",
			LastName:  "Golang",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("err while ccalling Greet PRC %v", err)
	}
	log.Printf("Response form Greet: %v", res.Result)
}
