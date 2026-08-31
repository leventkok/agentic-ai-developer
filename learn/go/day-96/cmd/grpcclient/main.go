package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	bookmarksv1 "learn/go/day-96/internal/gen/bookmarksv1"
	"learn/go/day-96/internal/grpcapi"
)

func main() {
	addr := flag.String("addr", "localhost:9090", "gRPC server address")
	httpBase := flag.String("http", "http://localhost:8080", "HTTP API base URL for login")
	email := flag.String("email", "demo@example.com", "login email when fetching token")
	password := flag.String("password", "password123", "login password when fetching token")
	tokenFlag := flag.String("token", "", "Bearer token (skips HTTP login when set)")
	flag.Parse()

	token := strings.TrimSpace(*tokenFlag)
	if token == "" {
		var err error
		token, err = loginToken(*httpBase, *email, *password)
		if err != nil {
			log.Printf("HTTP login failed (%v); continuing without token for list-only demo", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcapi.RequestIDClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	client := bookmarksv1.NewBookmarkServiceClient(conn)

	listResp, err := client.ListBookmarks(ctx, &bookmarksv1.ListBookmarksRequest{})
	if err != nil {
		log.Fatalf("ListBookmarks: %v", err)
	}
	fmt.Printf("Listed %d bookmark(s)\n", len(listResp.GetItems()))
	for _, item := range listResp.GetItems() {
		fmt.Printf("- [%d] %s (%s)\n", item.GetId(), item.GetTitle(), item.GetUrl())
	}

	if token == "" {
		return
	}

	authCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	createResp, err := client.CreateBookmark(authCtx, &bookmarksv1.CreateBookmarkRequest{
		Title: "gRPC demo bookmark",
		Url:   "https://grpc.io",
		Tags:  []string{"grpc", "demo"},
	})
	if err != nil {
		log.Fatalf("CreateBookmark: %v", err)
	}
	fmt.Printf("Created bookmark id=%d title=%q\n", createResp.GetId(), createResp.GetTitle())
}

func loginToken(baseURL, email, password string) (string, error) {
	body := fmt.Sprintf(`{"email":%q,"password":%q}`, email, password)
	req, err := http.NewRequest(http.MethodPost, strings.TrimRight(baseURL, "/")+"/auth/login", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login status %d: %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}

	var parsed struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return "", err
	}
	if strings.TrimSpace(parsed.Token) == "" {
		return "", fmt.Errorf("login response missing token")
	}
	return parsed.Token, nil
}
