package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/smith-golang/grpc-test/bidir_stream/bidir_streampb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello i am a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect : %v", err)
	}

	defer conn.Close()

	c := bidir_streampb.NewGreetingServiceClient(conn)
	// fmt.Println("Created client : %f", c)

	doBidirStreaming(c)
}

func doBidirStreaming(c bidir_streampb.GreetingServiceClient) {
	fmt.Println("Starting to do a Bi Directional Streaming RPC...")

	// create a stream by invoking the client

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*bidir_streampb.GreetEveryoneRequest{
		&bidir_streampb.GreetEveryoneRequest{
			Greeting: &bidir_streampb.Greeting{
				FirstName: "Smith",
			},
		},
		&bidir_streampb.GreetEveryoneRequest{
			Greeting: &bidir_streampb.Greeting{
				FirstName: "Golang",
			},
		},
		&bidir_streampb.GreetEveryoneRequest{
			Greeting: &bidir_streampb.Greeting{
				FirstName: "Sithu",
			},
		},
		&bidir_streampb.GreetEveryoneRequest{
			Greeting: &bidir_streampb.Greeting{
				FirstName: "Sheinko",
			},
		},
		&bidir_streampb.GreetEveryoneRequest{
			Greeting: &bidir_streampb.Greeting{
				FirstName: "Marsl",
			},
		},
	}

	waitc := make(chan struct{})

	// send a bunch of messages to the client (go routine)

	go func() {
		// function to send a bunch of message
		for _, req := range requests {
			fmt.Println("Sending message %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive a bunch of message from the client (go routine)

	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done

	<-waitc
}
