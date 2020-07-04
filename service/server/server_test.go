package main

import (
	"context"
	"testing"

	pb "grpc.go/service/protos"
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

			if resp.GetResult() != tc.resp {
				t.Errorf("Expected response Result to be %v but got %v", tc.resp, resp.GetResult())
			}
		})
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		name    string
		req     *pb.Operands
		resp    float32
		message string
	}{
		{
			name: "Divide two integers",
			req:  &pb.Operands{FirstOperand: 10, SecondOperand: 5},
			resp: 2,
		},
		{
			name:    "Divide by zero",
			req:     &pb.Operands{FirstOperand: 10, SecondOperand: 0},
			message: "Cannot divide by zero",
		},
	}

	s := server{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.req
			resp, _ := s.Divide(context.Background(), req)

			if tc.message != "" {
				if resp.GetResult() != tc.resp {
					t.Errorf("Expected response Result to be %v but got %v", tc.resp, resp.GetResult())
				}
			} else {
				if resp.GetMessage() != tc.message {
					t.Errorf("Expected response Message to be %v but got %v", tc.message, resp.GetMessage())
				}
			}
		})
	}

}
