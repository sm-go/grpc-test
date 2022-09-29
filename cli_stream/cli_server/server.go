package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/smith-golang/grpc-test/cli_stream/cli_streampb"
	"google.golang.org/grpc"
)

type server struct {
	cli_streampb.UnimplementedGreetingServiceServer
}

// func (*server) Greet(ctx context.Context, req *unarypb.GreetRequest) (*unarypb.GreetResponse, error) {
// 	fmt.Println("Greet functions was invoked with %v", req)
// 	firstName := req.GetGreeting().GetFirstName()
// 	result := "Hello " + firstName
// 	res := &unarypb.GreetResponse{
// 		Result: result,
// 	}
// 	return res, nil
// }

// Server Straming
// func (*server) GreetManyTimes(req *cli_streampb.GreetManyTimeRequest, stream cli_streampb.GreetingService_GreetManyTimesServer, error) {
func (*server) LongGreet(stream cli_streampb.GreetingService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked with a streaming request \n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//we have finished reading the client stream
			return stream.SendAndClose(&cli_streampb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream : %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "hello" + firstName + "!"
	}
}

func main() {
	fmt.Println("------Client Sreaming ... Testing--------")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") //default port for gRPC
	if err != nil {
		log.Fatal("Failed to listen :%v", err)
	}

	s := grpc.NewServer()
	cli_streampb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to served %v", err)
	}
}
