package domain_test

import (
	"testing"

	"learn/go/day-71/internal/domain"
)

func TestCanModifyBookmark_Owner(t *testing.T) {
	ownerID := 1
	actor := domain.User{ID: 1, Role: domain.RoleMember}
	bookmark := domain.Bookmark{UserID: &ownerID}
	if !domain.CanModifyBookmark(actor, bookmark) {
		t.Fatal("owner should modify own bookmark")
	}
}

func TestCanModifyBookmark_Admin(t *testing.T) {
	ownerID := 2
	actor := domain.User{ID: 1, Role: domain.RoleAdmin}
	bookmark := domain.Bookmark{UserID: &ownerID}
	if !domain.CanModifyBookmark(actor, bookmark) {
		t.Fatal("admin should modify any bookmark")
	}
}

func TestCanModifyBookmark_Forbidden(t *testing.T) {
	ownerID := 1
	actor := domain.User{ID: 2, Role: domain.RoleMember}
	bookmark := domain.Bookmark{UserID: &ownerID}
	if domain.CanModifyBookmark(actor, bookmark) {
		t.Fatal("non-owner member should not modify bookmark")
	}
}

func TestCanBulkCreate_AdminOnly(t *testing.T) {
	if !domain.CanBulkCreate(domain.User{Role: domain.RoleAdmin}) {
		t.Fatal("admin should bulk create")
	}
	if domain.CanBulkCreate(domain.User{Role: domain.RoleMember}) {
		t.Fatal("member should not bulk create")
	}
}
