package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// ReportsListGet Get paginated list of reports
func (c *Client) ReportsListGet(ctx context.Context, spaceID string, offset *int, limit *int, order *string, search *string, fileFormat *vampy.ReportFileFormat, status *vampy.ReportStatus) (vampy.Page[vampy.Report], error) {
	path := fmt.Sprintf("/spaces/%s/reports/", url.PathEscape(spaceID))
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	if fileFormat != nil {
		q.Set("fileFormat", string(*fileFormat))
	}
	if status != nil {
		q.Set("status", string(*status))
	}
	var out vampy.Page[vampy.Report]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ReportGet Get report details
func (c *Client) ReportGet(ctx context.Context, reportID string) (vampy.Report, error) {
	path := fmt.Sprintf("/reports/%s/", url.PathEscape(reportID))
	var out vampy.Report
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ReportDownload Get report file
func (c *Client) ReportDownload(ctx context.Context, reportID string) ([]byte, error) {
	path := fmt.Sprintf("/reports/%s/download/", url.PathEscape(reportID))
	return c.doBinary(ctx, "GET", path, nil, nil)
}
