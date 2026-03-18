# Тест go + grpc

генерация
```sh
protoc   --go_out=.   --go-grpc_out=.   proto/ping.proto
```

старт
```sh
go run cmd/main.go
```

проверка
```
grpcurl -plaintext -d '{"message":"hello"}' localhost:50051 ping.PingService/Ping
```