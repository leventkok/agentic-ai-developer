package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"learn/go/day-63/internal/app"
	"learn/go/day-63/internal/config"
	bookmarksv1 "learn/go/day-63/internal/gen/bookmarksv1"
	"learn/go/day-63/internal/grpcapi"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	deps, cleanup, err := app.WireGRPCDeps(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := cleanup(); err != nil {
			log.Print(err)
		}
	}()

	addr := ":9090"
	if port := os.Getenv("GRPC_PORT"); port != "" {
		addr = ":" + port
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	server := grpc.NewServer()
	bookmarksv1.RegisterBookmarkServiceServer(server, grpcapi.NewServer(deps.Bookmarks, deps.Auth))
	reflection.Register(server)

	go func() {
		fmt.Printf("gRPC bookmarks server on localhost%s\n", addr)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	stopped := make(chan struct{})
	go func() {
		server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
	case <-time.After(time.Duration(cfg.ShutdownTimeoutSec) * time.Second):
		server.Stop()
	}
	log.Println("gRPC server stopped cleanly")
}
