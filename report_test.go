package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestReportsListGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.ReportsListGet(ctx, ReportsListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("ReportsListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestReportGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestReportDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}
