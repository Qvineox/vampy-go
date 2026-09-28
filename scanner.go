package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
)

// ScanUploadPost Post new scanner result
func (c *Client) ScanUploadPost(ctx context.Context, body vampy.ScanUploadsBody) (vampy.ScanUpload, error) {
	path := "/scan_uploads/"
	var out vampy.ScanUpload
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanUploadDownload Download scan uploads file
func (c *Client) ScanUploadDownload(ctx context.Context, scanUploadID string) ([]byte, error) {
	path := fmt.Sprintf("/scan_uploads/%s/download/", url.PathEscape(scanUploadID))
	return c.doBinary(ctx, "GET", path, nil, nil)
}

// ScanUploadBulkDownload Download archive with scan upload files
func (c *Client) ScanUploadBulkDownload(ctx context.Context, body vampy.ScanUploadsBulkDownloadBody) ([]byte, error) {
	path := "/scan_uploads/bulk_download/"
	return c.doBinary(ctx, "POST", path, nil, body)
}
