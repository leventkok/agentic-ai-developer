package grpcapi

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"learn/go/day-65/internal/domain"
)

const (
	MetadataAuthorization = "authorization"
	MetadataRequestID     = "x-request-id"
)

// ToStatus maps domain and transport errors to gRPC status errors.
func ToStatus(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := status.FromError(err); ok {
		return err
	}
	if errors.Is(err, context.Canceled) {
		return status.Error(codes.Canceled, err.Error())
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error())
	}
	if errors.Is(err, domain.ErrNotFound) {
		return status.Error(codes.NotFound, "bookmark not found")
	}
	if errors.Is(err, domain.ErrForbidden) {
		return status.Error(codes.PermissionDenied, "forbidden")
	}
	if errors.Is(err, domain.ErrUnauthorized) {
		return status.Error(codes.Unauthenticated, "unauthorized")
	}
	if errors.Is(err, domain.ErrInvalidCredentials) {
		return status.Error(codes.Unauthenticated, "invalid credentials")
	}
	if errors.Is(err, domain.ErrDuplicateEmail) {
		return status.Error(codes.AlreadyExists, "email already registered")
	}
	if domain.IsValidation(err) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return status.Error(codes.Internal, "internal error")
}

// BearerToken extracts a Bearer JWT from incoming gRPC metadata.
func BearerToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "missing metadata")
	}
	vals := md.Get(MetadataAuthorization)
	if len(vals) == 0 {
		return "", status.Error(codes.Unauthenticated, "missing authorization metadata")
	}
	raw := strings.TrimSpace(vals[0])
	const prefix = "Bearer "
	if !strings.HasPrefix(raw, prefix) {
		return "", status.Error(codes.Unauthenticated, "authorization must use Bearer scheme")
	}
	token := strings.TrimSpace(strings.TrimPrefix(raw, prefix))
	if token == "" {
		return "", status.Error(codes.Unauthenticated, "empty bearer token")
	}
	return token, nil
}

// RequestIDFromMetadata reads the request identifier attached by clients.
func RequestIDFromMetadata(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	vals := md.Get(MetadataRequestID)
	if len(vals) == 0 {
		return ""
	}
	return strings.TrimSpace(vals[0])
}
