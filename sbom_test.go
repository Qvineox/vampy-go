package vampy_go

import (
	"testing"
)

func TestSBOMImport(t *testing.T) {
	// multipart upload requires a file fixture
	t.Skip("multipart upload requires a file fixture")
}

func TestSBOMExport(t *testing.T) {
	// binary download requires an existing artifact
	t.Skip("binary download requires an existing artifact")
}
