package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestUsersGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.UsersGet(ctx, UsersGetOptions{})
	if err != nil {
		t.Fatalf("UsersGet: %v", err)
	}
	requireSlice(t, got)
}

func TestUserRegistration(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestSpaceUsersGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.SpaceUsersGet(ctx, SpaceUsersGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("SpaceUsersGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestSpaceUserRegistration(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUserUserIDGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestUsersPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestUserCurrent(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.UserCurrent(ctx)
	if err != nil {
		t.Fatalf("UserCurrent: %v", err)
	}
	requireSchema(t, got)
}

func TestUsersLogoutUser(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUsersMFAPost(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUsersMFADelete(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUsersMFAConfirmPost(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}
