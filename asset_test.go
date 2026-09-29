package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestAssetListGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.AssetListGet(ctx, AssetListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("AssetListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestAssetPost(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestAssetGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestAssetDelete(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestAssetPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}
