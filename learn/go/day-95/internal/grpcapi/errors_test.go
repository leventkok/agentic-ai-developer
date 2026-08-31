package grpcapi_test

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"learn/go/day-95/internal/domain"
	"learn/go/day-95/internal/grpcapi"
)

func TestToStatus_DomainErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		code codes.Code
	}{
		{"not found", domain.ErrNotFound, codes.NotFound},
		{"forbidden", domain.ErrForbidden, codes.PermissionDenied},
		{"unauthorized", domain.ErrUnauthorized, codes.Unauthenticated},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			st, ok := status.FromError(grpcapi.ToStatus(tc.err))
			if !ok {
				t.Fatal("expected gRPC status")
			}
			if st.Code() != tc.code {
				t.Fatalf("code = %v, want %v", st.Code(), tc.code)
			}
		})
	}
}

func TestToStatus_Passthrough(t *testing.T) {
	original := status.Error(codes.Aborted, "aborted")
	got := grpcapi.ToStatus(original)
	if !errors.Is(got, original) {
		t.Fatal("expected passthrough of existing status error")
	}
}
