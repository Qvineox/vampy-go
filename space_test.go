package vampy_go

import (
	"testing"
)

func TestSpacesGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.SpacesGet(ctx)
	if err != nil {
		t.Fatalf("SpacesGet: %v", err)
	}
	requireSlice(t, got)
}

func TestSpaceCreate(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}
