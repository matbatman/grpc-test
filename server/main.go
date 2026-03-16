package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-test/gen/hello" // Импорт сгенерированного protobuf/gRPC-кода

	"google.golang.org/grpc"
)

// server — реализация gRPC-сервиса HelloService.
// Встраиваем UnimplementedHelloServiceServer, чтобы автоматически
// получать заглушки для будущих методов и совместимость с gRPC.
type server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello — реализация RPC-метода из .proto.
// Принимает контекст и запрос, возвращает ответ или ошибку.
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Привет, " + req.GetName(),
	}, nil
}

func main() {
	// Создаём TCP‑листенер на порту 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Создаём новый gRPC‑сервер.
	s := grpc.NewServer()

	// Регистрируем нашу реализацию HelloService на сервере.
	pb.RegisterHelloServiceServer(s, &server{})

	fmt.Println("gRPC server running on :50051")

	// Запускаем сервер и начинаем принимать соединения.
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
