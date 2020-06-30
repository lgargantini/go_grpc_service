package main

import (
	"context"
	pb "grpc_service/protos"
	"testing"
)

func TestAdd(t *testing.T) {
	// Generate Test cases
	testCases := []struct {
		name      string
		req       *pb.Operands
		resp      float32
		expectErr bool
	}{
		{
			name: "Add two integers",
			req:  &pb.Operands{FirstOperand: 1, SecondOperand: 2},
			resp: 3,
		},
	}

	s := server{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.req
			resp, err := s.Add(context.Background(), req)

			if !tc.expectErr {
				if resp.Result != tc.resp {
					t.Errorf("Expected response Result to be 3 but got %v", resp.Result)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error %v", err)
				}
			}
		})
	}

}
