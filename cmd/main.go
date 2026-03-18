package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	proto "grpc-test/proto/ping"
	"grpc-test/server"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterPingServiceServer(s, &server.PingServer{})

	// Включаем reflection
	reflection.Register(s)

	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
