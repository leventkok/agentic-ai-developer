package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc/reflection"

	"learn/go/day-81/internal/app"
	"learn/go/day-81/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	grpcServer, cleanup, err := app.WireGRPC(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := cleanup(); err != nil {
			log.Print(err)
		}
	}()

	addr := ":" + cfg.GRPCPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	reflection.Register(grpcServer)

	go func() {
		fmt.Printf("gRPC bookmarks server on localhost%s\n", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	stopped := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
	case <-time.After(time.Duration(cfg.ShutdownTimeoutSec) * time.Second):
		grpcServer.Stop()
	}
	log.Println("gRPC server stopped cleanly")
}
