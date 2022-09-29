package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/smith-golang/grpc-test/bidir_stream/bidir_streampb"
	"google.golang.org/grpc"
)

type server struct {
	bidir_streampb.UnimplementedGreetingServiceServer
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
// func (*server) GreetManyTimes(req *bidir_streampb.GreetManyTimeRequest, stream bidir_streampb.GreetingService_GreetManyTimesServer, error) {
func (*server) GreetEveryone(stream bidir_streampb.GreetingService_GreetEveryoneServer) error {
	fmt.Println("LongGreet function was invoked with a Bi Directional streaming request \n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Errror while reading stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		sendErr := stream.Send(&bidir_streampb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", err)
		}
	}
}

func main() {
	fmt.Println("------Bi Directional Sreaming ... Testing--------")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") //default port for gRPC
	if err != nil {
		log.Fatal("Failed to listen :%v", err)
	}

	s := grpc.NewServer()
	bidir_streampb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to served %v", err)
	}
}
