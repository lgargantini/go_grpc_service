package main

import (
	"context"
	pb "grpc_service/protos"
	"testing"
)

func TestAdd(t *testing.T) {
	s := server{}
	req := &pb.Operands{FirstOperand: 1, SecondOperand: 2}
	resp, _ := s.Add(context.Background(), req)

	if resp.Result != 3 {
		t.Errorf("Expected response Result to be 3 but got %v", resp.Result)
	}
}
