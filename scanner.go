package vampy_go

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// ScanUploadPost Post new scanner result
func (c *Client) ScanUploadPost(ctx context.Context, opts ScanUploadPostOptions) (vampy.ScanUpload, error) {
	path := "/scan_uploads/"

	var out vampy.ScanUpload
	if err := c.doMultipart(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanUploadPostOptions struct {
	Body vampy.ScanUploadsBody
}

// ScanUploadDownload Download scan uploads file
func (c *Client) ScanUploadDownload(ctx context.Context, opts ScanUploadDownloadOptions) ([]byte, error) {
	path := fmt.Sprintf("/scan_uploads/%s/download/", url.PathEscape(opts.ScanUploadID.String()))

	return c.doBinary(ctx, http.MethodGet, path, nil, nil)
}

type ScanUploadDownloadOptions struct {
	ScanUploadID uuid.UUID
}

// ScanUploadBulkDownload Download archive with scan upload files
func (c *Client) ScanUploadBulkDownload(ctx context.Context, opts ScanUploadBulkDownloadOptions) ([]byte, error) {
	path := "/scan_uploads/bulk_download/"

	return c.doBinary(ctx, http.MethodPost, path, nil, opts.Body)
}

type ScanUploadBulkDownloadOptions struct {
	Body vampy.ScanUploadsBulkDownloadBody
}
