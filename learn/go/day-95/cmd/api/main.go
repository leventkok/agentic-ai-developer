package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc/reflection"

	"learn/go/day-95/internal/app"
	"learn/go/day-95/internal/config"
	"learn/go/day-95/internal/observability/pprof"
	"learn/go/day-95/internal/observability/tracing"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	traceShutdown, err := tracing.Init(context.Background(), "bookmarks-api")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := traceShutdown(context.Background()); err != nil {
			log.Print(err)
		}
	}()

	pprofSrv, err := pprof.Start(pprof.Addr(cfg.PPROFPort))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := pprof.Shutdown(pprofSrv); err != nil {
			log.Print(err)
		}
	}()

	runtime, err := app.Wire(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := runtime.Cleanup(); err != nil {
			log.Print(err)
		}
	}()

	httpAddr := ":" + cfg.Port
	httpServer := &http.Server{
		Addr:         httpAddr,
		Handler:      runtime.HTTPHandler,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSec) * time.Second,
	}

	grpcAddr := ":" + cfg.GRPCPort
	grpcLis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("grpc listen: %v", err)
	}
	reflection.Register(runtime.GRPCServer)

	errCh := make(chan error, 2)
	go func() {
		fmt.Printf("Bookmarks API [%s] HTTP on http://localhost%s\n", cfg.Env, httpAddr)
		errCh <- httpServer.ListenAndServe()
	}()
	go func() {
		fmt.Printf("Bookmarks API [%s] gRPC on localhost%s\n", cfg.Env, grpcAddr)
		errCh <- runtime.GRPCServer.Serve(grpcLis)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeoutSec)*time.Second)
	defer cancel()

	grpcStopped := make(chan struct{})
	go func() {
		runtime.GRPCServer.GracefulStop()
		close(grpcStopped)
	}()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("http shutdown: %v", err)
	}

	select {
	case <-grpcStopped:
	case <-shutdownCtx.Done():
		runtime.GRPCServer.Stop()
	}

	log.Println("servers stopped cleanly")
}
