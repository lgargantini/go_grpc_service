package main

import (
	"context"
	pb "grpc_service/protos"
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
