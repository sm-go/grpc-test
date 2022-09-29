package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/smith-golang/grpc-test/unary/unarypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type server struct {
	unarypb.UnimplementedGreetingServiceServer
}

func (*server) Greet(ctx context.Context, req *unarypb.GreetRequest) (*unarypb.GreetResponse, error) {
	fmt.Println("Greet functions was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &unarypb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("------Unary Testing--------")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") //default port for gRPC
	if err != nil {
		log.Fatal("Failed to listen :%v", err)
	}

	opts := []grpc.ServerOption{}
	tls := false
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	unarypb.RegisterGreetingServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to served %v", err)
	}
}
