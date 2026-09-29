package vampy_go

import (
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestRunOrchestration(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestRunOrchestrationBatch(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestOrchestrationScanHistory(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.OrchestrationScanHistory(ctx, OrchestrationScanHistoryOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("OrchestrationScanHistory: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestGetOrchestrationProfiles(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.GetOrchestrationProfiles(ctx, GetOrchestrationProfilesOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("GetOrchestrationProfiles: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestCreateOrchestrationProfile(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestGetOrchestrationProfile(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestDeleteOrchestrationProfile(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestUpdateOrchestrationProfile(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestGetOrchestrationProfileRelations(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestDeleteScannerConfig(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestGetOrchestrationScannersInfo(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.GetOrchestrationScannersInfo(ctx, GetOrchestrationScannersInfoOptions{})
	if err != nil {
		t.Fatalf("GetOrchestrationScannersInfo: %v", err)
	}
	requireSlice(t, got)
}

func TestUploadScannerConfig(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestDownloadScannerConfig(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}
