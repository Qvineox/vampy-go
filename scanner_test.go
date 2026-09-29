package vampy_go

import (
	"testing"
)

func TestScanUploadPost(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestScanUploadDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}

func TestScanUploadBulkDownload(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}
