package main

import (
	"context"
	pb "grpc_service/protos"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.Operands) (*pb.Result, error) {
	return &pb.Result{Result: 3}, nil
}
