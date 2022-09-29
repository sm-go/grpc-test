package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/smith-golang/grpc-test/cli_stream/cli_streampb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello i am a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect : %v", err)
	}

	defer conn.Close()

	c := cli_streampb.NewGreetingServiceClient(conn)
	// fmt.Println("Created client : %f", c)

	doClientStreaming(c)
}

func doClientStreaming(c cli_streampb.GreetingServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*cli_streampb.LongGreetRequest{
		&cli_streampb.LongGreetRequest{
			Greeting: &cli_streampb.Greeting{
				FirstName: "Smith",
			},
		},
		&cli_streampb.LongGreetRequest{
			Greeting: &cli_streampb.Greeting{
				FirstName: "Golang",
			},
		},
		&cli_streampb.LongGreetRequest{
			Greeting: &cli_streampb.Greeting{
				FirstName: "Sithu",
			},
		},
		&cli_streampb.LongGreetRequest{
			Greeting: &cli_streampb.Greeting{
				FirstName: "Sheinko",
			},
		},
		&cli_streampb.LongGreetRequest{
			Greeting: &cli_streampb.Greeting{
				FirstName: "Marsl",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("errror while calling LongCreet %v", err)
	}

	for _, req := range requests {
		fmt.Println("Sending req : %v", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response for LongGreet : %v", err)
	}
	fmt.Printf("LongGreet Response: %v \n", res)
}
