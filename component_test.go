package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestComponentsListGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.ComponentsListGet(ctx, ComponentsListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("ComponentsListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestComponentDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestComponentPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestComponentVersionsListGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.ComponentVersionsListGet(ctx, ComponentVersionsListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("ComponentVersionsListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestComponentVersionDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestComponentVersionPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}
