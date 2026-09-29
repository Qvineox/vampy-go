package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// ScanPipelinesGet Get scan pipelines
func (c *Client) ScanPipelinesGet(ctx context.Context, opts ScanPipelinesGetOptions) (vampy.Page[vampy.ScanPipeline], error) {
	path := fmt.Sprintf("/spaces/%s/scan_pipelines/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.RepositoryID != nil {
		q.Set("repositoryID", opts.RepositoryID.String())
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	for _, v := range opts.Status {
		q.Add("status", string(v))
	}

	for _, v := range opts.ScanningTool {
		q.Add("scanningTool", v)
	}

	for _, v := range opts.InitiatorType {
		q.Add("initiatorType", string(v))
	}

	for _, v := range opts.ScanPipelineType {
		q.Add("scanPipelineType", string(v))
	}

	for _, v := range opts.RelationType {
		q.Add("relationType", string(v))
	}

	if opts.CreatedLte != nil {
		q.Set("createdLte", *opts.CreatedLte)
	}

	if opts.CreatedGte != nil {
		q.Set("createdGte", *opts.CreatedGte)
	}

	var out vampy.Page[vampy.ScanPipeline]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanPipelinesGetOptions struct {
	ProductID        *uuid.UUID
	RepositoryID     *uuid.UUID
	Order            *string
	Status           []vampy.ScanPipelineStatus
	ScanningTool     []string
	InitiatorType    []vampy.ScanPipelineInitiatorType
	ScanPipelineType []vampy.ScanPipelineType
	RelationType     []vampy.ScanPipelineRelationType
	CreatedLte       *string
	CreatedGte       *string

	vampy.Paging
}

// ScanPipelineDetailsGet Get scan pipelines
func (c *Client) ScanPipelineDetailsGet(ctx context.Context, opts ScanPipelineDetailsGetOptions) (vampy.ScanPipelineDetails, error) {
	path := fmt.Sprintf("/scan_pipelines/%s/", url.PathEscape(opts.ScanPipelineID.String()))

	var out vampy.ScanPipelineDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanPipelineDetailsGetOptions struct {
	ScanPipelineID uuid.UUID
}

// ScanPipelineJobsGet Get scan pipeline jobs
func (c *Client) ScanPipelineJobsGet(ctx context.Context, opts ScanPipelineJobsGetOptions) (vampy.Page[vampy.ScanPipelineJob], error) {
	path := fmt.Sprintf("/scan_pipelines/%s/jobs/", url.PathEscape(opts.ScanPipelineID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	for _, v := range opts.Status {
		q.Add("status", string(v))
	}

	for _, v := range opts.ScanningTool {
		q.Add("scanningTool", v)
	}

	for _, v := range opts.RelationType {
		q.Add("relationType", string(v))
	}

	var out vampy.Page[vampy.ScanPipelineJob]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanPipelineJobsGetOptions struct {
	ScanPipelineID uuid.UUID
	Order          *string
	Status         []vampy.ScanPipelineJobStatus
	ScanningTool   []string
	RelationType   []vampy.ScanPipelineRelationType

	vampy.Paging
}
