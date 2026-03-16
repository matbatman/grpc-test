package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-test/gen/hello" // Импорт сгенерированного gRPC-кода

	"google.golang.org/grpc"
)

func main() {
	// Устанавливаем соединение с gRPC-сервером.
	// grpc.WithInsecure() — отключает TLS (нормально для локальной разработки).
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Создаём gRPC-клиента для HelloService.
	client := pb.NewHelloServiceClient(conn)

	// Контекст с таймаутом — чтобы запрос не висел бесконечно.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Отправляем RPC-запрос SayHello.
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Давид"})
	if err != nil {
		log.Fatal(err)
	}

	// Выводим ответ сервера.
	fmt.Println("Ответ сервера:", resp.GetMessage())
}
