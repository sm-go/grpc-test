package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smith-golang/grpc-test/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello i am a Calculator client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect : %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	// fmt.Println("Created client : %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum RPC ...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("err while ccalling Sum PRC %v", err)
	}
	log.Printf("Response form Sum: %v", res.SumResult)
}
