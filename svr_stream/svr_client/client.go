package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/smith-golang/grpc-test/svr_stream/svr_streampb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello i am a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect : %v", err)
	}

	defer conn.Close()

	c := svr_streampb.NewGreetingServiceClient(conn)
	// fmt.Println("Created client : %f", c)

	doServerStreaming(c)
}

func doServerStreaming(c svr_streampb.GreetingServiceClient) {
	fmt.Println("Starting to do a serverStreaming RPC ...")
	req := &svr_streampb.GreetManyTimeRequest{
		Greeting: &svr_streampb.Greeting{
			FirstName: "Smith",
			LastName:  "Golang",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("err while calling GreetManyTimes PRC %v", err)
	}
	for {

		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading stream: ", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
