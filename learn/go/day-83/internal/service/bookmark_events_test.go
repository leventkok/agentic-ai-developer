package service_test

import (
	"context"
	"sync/atomic"
	"testing"

	"learn/go/day-83/internal/domain"
	"learn/go/day-83/internal/messaging"
	"learn/go/day-83/internal/messaging/memory"
	"learn/go/day-83/internal/service"
	"learn/go/day-83/internal/service/testing/fake"
)

func TestBookmarkService_CreatePublishesEvent(t *testing.T) {
	bus := memory.NewBus()
	var count atomic.Int32
	if err := bus.Subscribe(messaging.EventBookmarkCreated, func(ctx context.Context, evt messaging.Event) error {
		count.Add(1)
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	svc := service.NewBookmarkService(fake.NewBookmarks(), 0, bus)
	user := domain.User{ID: 1, Role: domain.RoleMember}
	if _, err := svc.Create(context.Background(), user, domain.CreateBookmarkInput{
		Title: "Go", URL: "https://go.dev",
	}); err != nil {
		t.Fatal(err)
	}
	if count.Load() != 1 {
		t.Fatalf("events published = %d", count.Load())
	}
}
