package main

import (
	"context"
	"log"
	"time"

	pb "github.com/charles-co/go-usermgmt-grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Contact the server and print out its response.
	name := "Charles"
	email := "test@test.com"
	age := int32(30)

	r, err := c.CreateUserRequest(ctx, &pb.CreateUserRequest{Name: name, Email: email, Age: age})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(`User Details:
	ID: %d
	Name: %s
	Email: %s
	`, r.GetId(), r.GetName(), r.GetEmail())
}
