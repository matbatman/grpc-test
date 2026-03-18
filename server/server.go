package server

import (
	"context"
	proto "grpc-test/proto/ping"
)

type PingServer struct {
	proto.UnimplementedPingServiceServer
}

func (s *PingServer) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Message: "pong: " + req.Message,
	}, nil
}

func (s *PingServer) PingReverse(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	r := reverse(req.Message)
	return &proto.PingResponse{
		Message: r,
	}, nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
