package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smith-golang/grpc-test/unary/unarypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("hello i am a client")

	certFile := "ssl/ca.crt" //Certificate Authority Trust Certificate

	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
		return
	}

	opts := grpc.WithTransportCredentials(creds)

	conn, err := grpc.Dial("localhost:50051", opts)

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
		log.Fatal("err while calling Greet PRC %v", err)
	}
	log.Printf("Response form Greet: %v", res.Result)
}
