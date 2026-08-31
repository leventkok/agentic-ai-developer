package service_test

import (
	"context"
	"testing"

	"learn/go/day-99/internal/db/testutil"
	"learn/go/day-99/internal/domain"
	"learn/go/day-99/internal/messaging/outbox"
	"learn/go/day-99/internal/service"
	"learn/go/day-99/internal/service/testing/fake"
)

func TestBookmarkService_CreateEnqueuesOutbox(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	outboxStore := outbox.NewStore(tdb.DB)

	svc := service.NewBookmarkService(fake.NewBookmarks(), 0, outboxStore, nil)
	user := domain.User{ID: 1, Role: domain.RoleMember}
	if _, err := svc.Create(context.Background(), user, domain.CreateBookmarkInput{
		Title: "Go", URL: "https://go.dev",
	}); err != nil {
		t.Fatal(err)
	}

	records, err := outboxStore.FetchUnpublished(context.Background(), 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 1 {
		t.Fatalf("outbox records = %d, want 1", len(records))
	}
}
