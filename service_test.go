package vampy_go

import (
	"context"
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestSpaceIssueStatusList(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.SpaceIssueStatusList(ctx, SpaceIssueStatusListOptions{})
	if err != nil {
		t.Fatalf("SpaceIssueStatusList: %v", err)
	}
	requireSlice(t, got)
}

func TestSpaceIssueStatusCreate(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	got, err := c.SpaceIssueStatusCreate(ctx, SpaceIssueStatusCreateOptions{Body: vampy.IssueStatusCreate{DisplayName: name}})
	if err != nil {
		t.Fatalf("SpaceIssueStatusCreate: %v", err)
	}
	requireSchema(t, got)
	t.Cleanup(func() {
		_ = c.SpaceIssueStatusDelete(context.Background(), SpaceIssueStatusDeleteOptions{IssueStatusName: got.Name})
	})
}

func TestSpaceIssueStatusGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestSpaceIssueStatusDelete(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	created, err := c.SpaceIssueStatusCreate(ctx, SpaceIssueStatusCreateOptions{Body: vampy.IssueStatusCreate{DisplayName: name}})
	if err != nil {
		t.Fatalf("SpaceIssueStatusCreate: %v", err)
	}
	if err := c.SpaceIssueStatusDelete(ctx, SpaceIssueStatusDeleteOptions{IssueStatusName: created.Name}); err != nil {
		t.Fatalf("SpaceIssueStatusDelete: %v", err)
	}
}

func TestSpaceIssueStatusPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestSystemSettingsGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.SystemSettingsGet(ctx)
	if err != nil {
		t.Fatalf("SystemSettingsGet: %v", err)
	}
	requireSchema(t, got)
}

func TestSystemSettingsPatch(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestFilterValuesGet(t *testing.T) {
	// requires a query parameter that cannot be defaulted safely
	t.Skip("requires a query parameter that cannot be defaulted safely")
}

func TestPingGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	err := c.PingGet(ctx)
	if err != nil {
		t.Fatalf("PingGet: %v", err)
	}
}

func TestHealthGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.HealthGet(ctx)
	if err != nil {
		t.Fatalf("HealthGet: %v", err)
	}
	requireSchema(t, got)
}

func TestServerInfo(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.ServerInfo(ctx)
	if err != nil {
		t.Fatalf("ServerInfo: %v", err)
	}
	requireSchema(t, got)
}

func TestLDAPConfigGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.LDAPConfigGet(ctx)
	if err != nil {
		t.Fatalf("LDAPConfigGet: %v", err)
	}
	requireSchema(t, got)
}

func TestLDAPConfigPost(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestLDAPConfigDelete(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestLDAPConfigPatch(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestLDAPConfigTest(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestOIDCConfigList(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.OIDCConfigList(ctx)
	if err != nil {
		t.Fatalf("OIDCConfigList: %v", err)
	}
	requireSlice(t, got)
}

func TestOIDCConfigPost(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestOIDCConfigGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestOIDCConfigDelete(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestOIDCConfigPatch(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestOIDCConfigDefaults(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestOIDCTestConnect(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestSMTPConfigGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.SMTPConfigGet(ctx)
	if err != nil {
		t.Fatalf("SMTPConfigGet: %v", err)
	}
	requireSchema(t, got)
}

func TestSMTPConfigPost(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestSMTPConfigDelete(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestSMTPConfigPatch(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestSendTestEmail(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUserRolesGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.UserRolesGet(ctx, UserRolesGetOptions{})
	if err != nil {
		t.Fatalf("UserRolesGet: %v", err)
	}
	requireSlice(t, got)
}

func TestUserRolesCreate(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUserRolesPermissionsMeta(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.UserRolesPermissionsMeta(ctx)
	if err != nil {
		t.Fatalf("UserRolesPermissionsMeta: %v", err)
	}
	requireSchema(t, got)
}

func TestUserRolesDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestUserRolesDelete(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUserRolesPatch(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestGetParsers(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.GetParsers(ctx)
	if err != nil {
		t.Fatalf("GetParsers: %v", err)
	}
	requireSlice(t, got)
}

func TestBackgroundTasksListGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.BackgroundTasksListGet(ctx, BackgroundTasksListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("BackgroundTasksListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestBackgroundTasksDetailsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestGetGlobalSLASetting(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.GetGlobalSLASetting(ctx)
	if err != nil {
		t.Fatalf("GetGlobalSLASetting: %v", err)
	}
	requireSchema(t, got)
}

func TestPostGlobalSLASetting(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestFPClassifierDatasetDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}

func TestFPClassifierDatasetUpload(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestFPClassifierDatasetClear(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestFPClassifierTrainingDataStats(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.FPClassifierTrainingDataStats(ctx)
	if err != nil {
		t.Fatalf("FPClassifierTrainingDataStats: %v", err)
	}
	requireSlice(t, got)
}

func TestFPClassifierSnapshotList(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.FPClassifierSnapshotList(ctx, FPClassifierSnapshotListOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("FPClassifierSnapshotList: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestFPClassifierSnapshotsDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}

func TestFPClassifierSnapshotUpload(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestFPClassifierSnapshotGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestFPClassifierSnapshotDelete(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestFPClassifierSnapshotPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestFPClassifierSnapshotTrain(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestFPClassifierSnapshotDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}

func TestFPClassifierSnapshotDatasetDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}

func TestFPClassifierSettingsGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.FPClassifierSettingsGet(ctx)
	if err != nil {
		t.Fatalf("FPClassifierSettingsGet: %v", err)
	}
	requireSchema(t, got)
}

func TestFPClassifierSettingsPatch(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestFPClassifierMetricsDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}
