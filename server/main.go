package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-test/gen/hello"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Привет, " + req.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	fmt.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
