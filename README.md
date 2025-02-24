**GRPC Rest Example**

1. Generate Go code from `user.proto`

```
protoc -I proto -I ./googleapis \
  --go_out=proto --go_opt=paths=source_relative \
  --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative \
  proto/user.proto
```

2. Start the gRPC Server

```
go run server/main.go
```

3. Start the REST Gateway (gRPC-Gateway)

```
go run gateway/main.go
```

4. Test gRPC with grpcurl

```
grpcurl -plaintext -d '{"id":"1"}' localhost:50051 grpc.rest.example.UserService/GetUser
```

```
grpcurl -plaintext -d '{"name":"Dori"}' localhost:50051 grpc.rest.example.UserService/CreateUser
```

5. Test RESET with curl

```
curl -X GET localhost:9080/v1/users/1
```

```
curl -X POST -d '{"name":"Alice"}' localhost:9080/v1/users
```

6. Run client to make calls to gRPC Server (optional)

```
go run client/main.go
```