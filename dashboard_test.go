package vampy_go

import (
	"testing"
)

func TestGlobalDashboardGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.GlobalDashboardGet(ctx, GlobalDashboardGetOptions{})
	if err != nil {
		t.Fatalf("GlobalDashboardGet: %v", err)
	}
	requireSchema(t, got)
}

func TestGlobalDashboardCardGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestProductDashboardGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestProductDashboardCardGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestRepositoryDashboardGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestRepositoryDashboardCardGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}
