package main

import (
	"context"
	"log"
	"time"

	pb "grpc.go/service/protos"

	"google.golang.org/grpc"
)

func initialize() (pb.CalculatorServiceClient, context.Context) {
	conn, err := grpc.Dial("localhost:12000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}

	defer conn.Close()

	// Instantiate a new client
	client := pb.NewCalculatorServiceClient(conn)

	// Generate the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	return client, ctx
}

func main() {
	client, ctx := initialize()

	response, err := client.Add(ctx, &pb.Operands{FirstOperand: 10, SecondOperand: 20})
	if err != nil {
		log.Fatalf("couldn't call method: %v", err)
	}

	log.Printf("Response: %v", response.GetResult())
}
