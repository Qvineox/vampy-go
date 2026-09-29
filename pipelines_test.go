package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestScanPipelinesGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.ScanPipelinesGet(ctx, ScanPipelinesGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("ScanPipelinesGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestScanPipelineDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestScanPipelineJobsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}
