package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-rest-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"math/rand/v2"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received GetUser request: %v", req) // Log incoming request

	age := int32(rand.IntN(100))
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:   req.GetId(),
			Name: "John Doe",
			Age:  &age,
		},
	}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Received CreateUser request: %+v", req) // Log incoming request

	// Simulate creating a user and returning an ID
	newID := rand.IntN(10000) // This would typically come from a database
	return &pb.CreateUserResponse{
		Id:   fmt.Sprintf("%v", newID),
		Name: req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	// Enable reflection
	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
