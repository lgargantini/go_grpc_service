package main

import (
	"context"
	"fmt"
	"log"
	"time"

	flag "github.com/spf13/pflag"

	pb "grpc.go/service/protos"

	"google.golang.org/grpc"
)

func initialize() (pb.CalculatorServiceClient, context.Context, context.CancelFunc, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:12000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}

	// Instantiate a new client
	client := pb.NewCalculatorServiceClient(conn)

	// Generate the context
	ctx, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)

	return client, ctx, cancelContext, conn
}

func main() {
	action := flag.String("o", "add", "Operation to perform")
	var defaultFloats []float32
	operands := flag.Float32Slice("n", defaultFloats, "Operands for the operation")

	flag.Parse()

	if len(*operands) != 2 {
		log.Fatal("operands must be two")
	}

	fmt.Println((*operands)[0])

	client, ctx, cancelContext, connection := initialize()
	defer connection.Close()
	defer cancelContext()

	var response *pb.Result
	var err error

	switch *action {
	case "add":
		response, err = client.Add(ctx, &pb.Operands{FirstOperand: (*operands)[0], SecondOperand: (*operands)[1]})
	case "division":
		response, err = client.Divide(ctx, &pb.Operands{FirstOperand: (*operands)[0], SecondOperand: (*operands)[1]})
	default:
		log.Fatalf("don't know operation %v\n", *action)
	}

	if err != nil {
		log.Fatalf("couldn't call method: %v", err)
	}

	log.Printf("Response: %v", response.GetResult())
}
