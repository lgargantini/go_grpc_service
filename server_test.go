package main

import (
	"context"
	pb "grpc_service/protos"
	"testing"
)

func TestAdd(t *testing.T) {
	// Generate Test cases
	testCases := []struct {
		name string
		req  *pb.Operands
		resp float32
	}{
		{
			name: "Add two integers",
			req:  &pb.Operands{FirstOperand: 1, SecondOperand: 2},
			resp: 3,
		},
		{
			name: "Add two floats",
			req:  &pb.Operands{FirstOperand: 2.5, SecondOperand: 10.5},
			resp: 13.0,
		},
	}

	s := server{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.req
			resp, _ := s.Add(context.Background(), req)

			if resp.Result != tc.resp {
				t.Errorf("Expected response Result to be %v but got %v", tc.resp, resp.Result)
			}
		})
	}

}
