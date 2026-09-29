package vampy_go

import (
	"context"
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestBotUserList(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.BotUserList(ctx)
	if err != nil {
		t.Fatalf("BotUserList: %v", err)
	}
	requireSlice(t, got)
}

func TestBotUserRegistration(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	got, err := c.BotUserRegistration(ctx, BotUserRegistrationOptions{Body: vampy.BotUserRegistration{UserLogin: name}})
	if err != nil {
		t.Fatalf("BotUserRegistration: %v", err)
	}
	requireSchema(t, got)
	t.Cleanup(func() {
		_, _ = c.BotUserDelete(context.Background(), BotUserDeleteOptions{UserID: uuidFromString(t, got.ID)})
	})
}

func TestBotUserDelete(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	created, err := c.BotUserRegistration(ctx, BotUserRegistrationOptions{Body: vampy.BotUserRegistration{UserLogin: name}})
	if err != nil {
		t.Fatalf("BotUserRegistration: %v", err)
	}
	if _, err := c.BotUserDelete(ctx, BotUserDeleteOptions{UserID: uuidFromString(t, created.ID)}); err != nil {
		t.Fatalf("BotUserDelete: %v", err)
	}
}

func TestBotUserPatchToken(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}
