package grpcapi_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	bookmarksv1 "learn/go/day-77/internal/gen/bookmarksv1"
	"learn/go/day-77/internal/grpcapi"
	"learn/go/day-77/internal/service"
	"learn/go/day-77/internal/service/testing/fake"
)

func TestUnaryServerInterceptor_RejectsMissingAuth(t *testing.T) {
	bookmarkRepo := fake.NewBookmarks()
	authRepo := fake.NewAuth()
	bookmarkSvc := service.NewBookmarkService(bookmarkRepo, time.Second)
	authSvc := service.NewAuthService(authRepo)
	serverImpl := grpcapi.NewServer(bookmarkSvc)

	lis := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcapi.UnaryServerInterceptor(authSvc, log.Default())))
	bookmarksv1.RegisterBookmarkServiceServer(grpcServer, serverImpl)

	go func() {
		_ = grpcServer.Serve(lis)
	}()
	t.Cleanup(func() {
		grpcServer.Stop()
		_ = lis.Close()
	})

	conn, err := grpc.NewClient("passthrough:///bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcapi.RequestIDClientInterceptor()),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })

	client := bookmarksv1.NewBookmarkServiceClient(conn)
	_, err = client.CreateBookmark(context.Background(), &bookmarksv1.CreateBookmarkRequest{
		Title: "blocked",
		Url:   "https://example.com",
	})
	if err == nil {
		t.Fatal("expected unauthenticated error")
	}
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status, got %v", err)
	}
	if st.Code() != codes.Unauthenticated {
		t.Fatalf("code = %v, want Unauthenticated", st.Code())
	}
}

func TestUnaryServerInterceptor_AllowsPublicList(t *testing.T) {
	bookmarkRepo := fake.NewBookmarks()
	authRepo := fake.NewAuth()
	bookmarkSvc := service.NewBookmarkService(bookmarkRepo, time.Second)
	authSvc := service.NewAuthService(authRepo)
	serverImpl := grpcapi.NewServer(bookmarkSvc)

	lis := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcapi.UnaryServerInterceptor(authSvc, log.Default())))
	bookmarksv1.RegisterBookmarkServiceServer(grpcServer, serverImpl)

	go func() {
		_ = grpcServer.Serve(lis)
	}()
	t.Cleanup(func() {
		grpcServer.Stop()
		_ = lis.Close()
	})

	conn, err := grpc.NewClient("passthrough:///bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcapi.RequestIDClientInterceptor()),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })

	client := bookmarksv1.NewBookmarkServiceClient(conn)
	resp, err := client.ListBookmarks(context.Background(), &bookmarksv1.ListBookmarksRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.GetItems()) != 0 {
		t.Fatalf("items len = %d, want 0", len(resp.GetItems()))
	}
}
