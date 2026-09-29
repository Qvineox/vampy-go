package vampy_go

import (
	"testing"
)

func TestLicenseGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.LicenseGet(ctx)
	if err != nil {
		t.Fatalf("LicenseGet: %v", err)
	}
	requireSchema(t, got)
}

func TestLicensePost(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestLicenseDelete(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestLicenseInformation(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.LicenseInformation(ctx)
	if err != nil {
		t.Fatalf("LicenseInformation: %v", err)
	}
	requireSchema(t, got)
}
