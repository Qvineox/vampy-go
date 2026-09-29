package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestArtifactsListGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.ArtifactsListGet(ctx, ArtifactsListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("ArtifactsListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestArtifactsDeleteBatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestArtifactsDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestArtifactsPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestArtifactVersionsSync(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestArtifactsVersionsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestArtifactVersionsDeleteBatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestArtifactsImport(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}
