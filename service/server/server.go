package main

import (
	"context"
	pb "grpc_service/service/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.Operands) (*pb.Result, error) {

	result := in.GetFirstOperand() + in.GetSecondOperand()

	return &pb.Result{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.Operands) (*pb.Result, error) {
	if in.GetSecondOperand() == 0 {
		return &pb.Result{Message: "Cannot divide by zero"}, nil
	}

	result := in.GetFirstOperand() / in.GetSecondOperand()

	return &pb.Result{Result: result}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":12000")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	svr := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(svr, &server{})
	log.Printf("Listening on port %d", 12000)
	if err := svr.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
