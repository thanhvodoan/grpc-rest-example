package main

import (
	"context"
	"log"
	"net/http"

	pb "grpc-rest-example/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := runtime.NewServeMux()
	err := pb.RegisterUserServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:50051", // Address of the gRPC server
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	log.Println("REST server listening on :9080")
	http.ListenAndServe(":9080", mux)
}
