package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestRulesListGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.RulesListGet(ctx, RulesListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("RulesListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestRulesPost(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRulesCheck(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRulesCheckAsync(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRulesEventsGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.RulesEventsGet(ctx)
	if err != nil {
		t.Fatalf("RulesEventsGet: %v", err)
	}
	requireSlice(t, got)
}

func TestRulesDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestRulesDelete(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRulesPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}
