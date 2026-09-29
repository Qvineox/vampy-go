package vampy_go

import (
	"testing"
)

func TestScanIssueCreate(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestScanIssuesFilter(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestScanIssuesAsyncExport(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestScanIssuesExportMeta(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestScanIssueSearchDuplicates(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestScanIssuesDelete(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestScanIssueSchemasGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.ScanIssueSchemasGet(ctx)
	if err != nil {
		t.Fatalf("ScanIssueSchemasGet: %v", err)
	}
	requireSlice(t, got)
}

func TestScanIssueDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestScanIssueStatusChangesGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestScanIssueChangesGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestScanIssueAIRequestPost(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestScanIssueAIRequestCancel(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestScanIssueAIRequestDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestScanIssueAIRequestApply(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestScanIssuesUpdate(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}
