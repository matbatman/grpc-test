package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-test/gen/hello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Давид"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ответ сервера:", resp.GetMessage())
}
