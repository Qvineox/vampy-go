package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// ScanPipelinesGet Get scan pipelines
func (c *Client) ScanPipelinesGet(ctx context.Context, spaceID string, productID *string, repositoryID *string, order *string, status []string, scanningTool []string, initiatorType []string, scanPipelineType []string, relationType []string, createdLte *string, createdGte *string, offset *int, limit *int) (vampy.InlineResponse20027, error) {
	path := fmt.Sprintf("/spaces/%s/scan_pipelines/", url.PathEscape(spaceID))
	q := url.Values{}
	if productID != nil {
		q.Set("productID", *productID)
	}
	if repositoryID != nil {
		q.Set("repositoryID", *repositoryID)
	}
	if order != nil {
		q.Set("order", *order)
	}
	for _, v := range status {
		q.Add("status", v)
	}
	for _, v := range scanningTool {
		q.Add("scanningTool", v)
	}
	for _, v := range initiatorType {
		q.Add("initiatorType", v)
	}
	for _, v := range scanPipelineType {
		q.Add("scanPipelineType", v)
	}
	for _, v := range relationType {
		q.Add("relationType", v)
	}
	if createdLte != nil {
		q.Set("createdLte", *createdLte)
	}
	if createdGte != nil {
		q.Set("createdGte", *createdGte)
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	var out vampy.InlineResponse20027
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanPipelineDetailsGet Get scan pipelines
func (c *Client) ScanPipelineDetailsGet(ctx context.Context, scanPipelineID string) (vampy.ScanPipelineDetails, error) {
	path := fmt.Sprintf("/scan_pipelines/%s/", url.PathEscape(scanPipelineID))
	var out vampy.ScanPipelineDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanPipelineJobsGet Get scan pipeline jobs
func (c *Client) ScanPipelineJobsGet(ctx context.Context, scanPipelineID string, order *string, status []string, scanningTool []string, relationType []string, offset *int, limit *int) (vampy.InlineResponse20028, error) {
	path := fmt.Sprintf("/scan_pipelines/%s/jobs/", url.PathEscape(scanPipelineID))
	q := url.Values{}
	if order != nil {
		q.Set("order", *order)
	}
	for _, v := range status {
		q.Add("status", v)
	}
	for _, v := range scanningTool {
		q.Add("scanningTool", v)
	}
	for _, v := range relationType {
		q.Add("relationType", v)
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	var out vampy.InlineResponse20028
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
