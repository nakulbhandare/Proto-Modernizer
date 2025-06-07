package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpc_migration/proto/modern"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUser(ctx, &pb.UserRequest{Id: "123"})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User: ID=%s, Name=%s, Description=%s", r.Id, r.Name, r.Description)
}
