package grpcapi

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"learn/go/day-77/internal/ctxkey"
	bookmarksv1 "learn/go/day-77/internal/gen/bookmarksv1"
	"learn/go/day-77/internal/service"
)

var protectedMethods = map[string]struct{}{
	bookmarksv1.BookmarkService_CreateBookmark_FullMethodName: {},
	bookmarksv1.BookmarkService_DeleteBookmark_FullMethodName: {},
}

// UnaryServerInterceptor logs RPC calls and validates auth metadata on protected methods.
func UnaryServerInterceptor(auth *service.AuthService, logger *log.Logger) grpc.UnaryServerInterceptor {
	if logger == nil {
		logger = log.Default()
	}
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		requestID := RequestIDFromMetadata(ctx)
		if requestID == "" {
			requestID = newRequestID()
		}

		if _, protected := protectedMethods[info.FullMethod]; protected {
			token, err := BearerToken(ctx)
			if err != nil {
				logger.Printf("rpc=%s request_id=%s code=%s duration=%s", info.FullMethod, requestID, codes.Unauthenticated, time.Since(start))
				return nil, err
			}
			user, err := auth.UserFromToken(ctx, token)
			if err != nil {
				st := ToStatus(err)
				logger.Printf("rpc=%s request_id=%s code=%s duration=%s err=%v", info.FullMethod, requestID, statusCode(st), time.Since(start), err)
				return nil, st
			}
			ctx = ctxkey.WithUser(ctx, user)
		}

		resp, err := handler(ctx, req)
		logger.Printf("rpc=%s request_id=%s code=%s duration=%s", info.FullMethod, requestID, statusCode(err), time.Since(start))
		return resp, err
	}
}

// RequestIDClientInterceptor attaches a unique request ID to outgoing RPC metadata.
func RequestIDClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = metadata.AppendToOutgoingContext(ctx, MetadataRequestID, newRequestID())
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func statusCode(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	st, ok := status.FromError(err)
	if !ok {
		return codes.Unknown
	}
	return st.Code()
}

func newRequestID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
