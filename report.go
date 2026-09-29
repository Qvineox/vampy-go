package vampy_go

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// ReportsListGet Get paginated list of reports
func (c *Client) ReportsListGet(ctx context.Context, opts ReportsListGetOptions) (vampy.Page[vampy.Report], error) {
	path := fmt.Sprintf("/spaces/%s/reports/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.FileFormat != nil {
		q.Set("fileFormat", string(*opts.FileFormat))
	}

	if opts.Status != nil {
		q.Set("status", string(*opts.Status))
	}

	var out vampy.Page[vampy.Report]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ReportsListGetOptions struct {
	Order      *string
	Search     *string
	FileFormat *vampy.ReportFileFormat
	Status     *vampy.ReportStatus

	vampy.Paging
}

// ReportGet Get report details
func (c *Client) ReportGet(ctx context.Context, opts ReportGetOptions) (vampy.Report, error) {
	path := fmt.Sprintf("/reports/%s/", url.PathEscape(opts.ReportID.String()))

	var out vampy.Report
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ReportGetOptions struct {
	ReportID uuid.UUID
}

// ReportDownload Get report file
func (c *Client) ReportDownload(ctx context.Context, opts ReportDownloadOptions) ([]byte, error) {
	path := fmt.Sprintf("/reports/%s/download/", url.PathEscape(opts.ReportID.String()))

	return c.doBinary(ctx, http.MethodGet, path, nil, nil)
}

type ReportDownloadOptions struct {
	ReportID uuid.UUID
}
