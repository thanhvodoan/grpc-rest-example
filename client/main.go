package main

import (
	"context"
	"log"
	"time"

	pb "grpc-rest-example/proto"

	"google.golang.org/grpc"
)

func getUser(client pb.UserServiceClient) {
	req := &pb.GetUserRequest{
		Id: "123",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	log.Printf("User: %+v", res.User)
}

func createUser(client pb.UserServiceClient) {
	createReq := &pb.CreateUserRequest{
		Name: "Jane Doe",
		Age:  28,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createRes, err := client.CreateUser(ctx, createReq)
	if err != nil {
		log.Fatalf("Error calling CreateUser: %v", err)
	}

	log.Printf("Created User ID: %s", createRes.Id)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	getUser(client)
	createUser(client)
}
