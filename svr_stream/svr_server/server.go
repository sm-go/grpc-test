package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/smith-golang/grpc-test/svr_stream/svr_streampb"
	"google.golang.org/grpc"
)

type server struct {
	svr_streampb.UnimplementedGreetingServiceServer
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
// func (*server) GreetManyTimes(req *svr_streampb.GreetManyTimeRequest, stream svr_streampb.GreetingService_GreetManyTimesServer, error) {
func (s *server) GreetManyTimes(req *svr_streampb.GreetManyTimeRequest, stream svr_streampb.GreetingService_GreetManyTimesServer) error {

	fmt.Println("GreetManyTimes function invoke with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &svr_streampb.GreetManyTimeResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil

}

func main() {
	fmt.Println("------Server Sreaming ... Testing--------")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") //default port for gRPC
	if err != nil {
		log.Fatal("Failed to listen :%v", err)
	}

	s := grpc.NewServer()
	svr_streampb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to served %v", err)
	}
}
