package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestFindingsGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.FindingsGet(ctx, FindingsGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("FindingsGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestFindingCardGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}
