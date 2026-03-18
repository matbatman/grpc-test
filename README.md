# Тест go + grpc

генерация
```sh
protoc \
  -I api/proto \
  --go_out=api/gen --go_opt=paths=source_relative \
  --go-grpc_out=api/gen --go-grpc_opt=paths=source_relative \
  api/proto/hello.proto

```