package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/qvineox/vampy-go/vampy"
)

// ExtV1ProductListGet Get Products List
func (c *Client) ExtV1ProductListGet(ctx context.Context, opts ExtV1ProductListGetOptions) (vampy.Page[vampy.Product], error) {
	path := "/ext/v1/products/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.Repository != nil {
		q.Set("repository", *opts.Repository)
	}

	if opts.IDs != nil {
		q.Set("ids", *opts.IDs)
	}

	var out vampy.Page[vampy.Product]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ProductListGetOptions struct {
	Order      *string
	Search     *string
	Repository *string
	IDs        *string

	vampy.Paging
}

// ExtV1ArtifactVersionsSync Synchronize artifact versions from registry integrations
func (c *Client) ExtV1ArtifactVersionsSync(ctx context.Context, opts ExtV1ArtifactVersionsSyncOptions) (vampy.ArtifactVersionSyncResponse, error) {
	path := "/ext/v1/artifacts/versions/sync/"

	var out vampy.ArtifactVersionSyncResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ArtifactVersionsSyncOptions struct {
	Body vampy.ArtifactVersionSyncRequest
}

// ExtV1RepositoryBranchesSync Synchronize repository branches from VCS integrations
func (c *Client) ExtV1RepositoryBranchesSync(ctx context.Context, opts ExtV1RepositoryBranchesSyncOptions) (vampy.RepositoryBranchSyncResponse, error) {
	path := "/ext/v1/repositories/branches/sync/"

	var out vampy.RepositoryBranchSyncResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1RepositoryBranchesSyncOptions struct {
	Body vampy.RepositoryBranchSyncRequest
}

// ExtV1RepositoryListGet Get Repositories List
func (c *Client) ExtV1RepositoryListGet(ctx context.Context, opts ExtV1RepositoryListGetOptions) (vampy.Page[vampy.Repository], error) {
	path := "/ext/v1/repositories/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Product != nil {
		q.Set("product", *opts.Product)
	}

	if opts.Artifact != nil {
		q.Set("artifact", *opts.Artifact)
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.IDs != nil {
		q.Set("ids", *opts.IDs)
	}

	var out vampy.Page[vampy.Repository]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1RepositoryListGetOptions struct {
	Product  *string
	Artifact *string
	Order    *string
	Search   *string
	IDs      *string

	vampy.Paging
}

// ExtV1ScanUploadPost External post new scanner result
func (c *Client) ExtV1ScanUploadPost(ctx context.Context, opts ExtV1ScanUploadPostOptions) (vampy.ScanUpload, error) {
	path := "/ext/v1/scan_uploads/"

	var out vampy.ScanUpload
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ScanUploadPostOptions struct {
	Body vampy.V1ScanUploadsBody
}

// ExtV1QualityGateGet Get Quality Gate result for product or repository.
func (c *Client) ExtV1QualityGateGet(ctx context.Context, opts ExtV1QualityGateGetOptions) (vampy.QualityGateResultDetails, error) {
	path := "/ext/v1/quality_gate/"
	q := url.Values{}

	if opts.SpaceID != nil {
		q.Set("spaceID", opts.SpaceID.String())
	}

	if opts.Product != nil {
		q.Set("product", *opts.Product)
	}

	if opts.Repository != nil {
		q.Set("repository", *opts.Repository)
	}

	if opts.RepositoryBranch != nil {
		q.Set("repositoryBranch", *opts.RepositoryBranch)
	}

	if opts.Slug != nil {
		q.Set("slug", *opts.Slug)
	}

	var out vampy.QualityGateResultDetails
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1QualityGateGetOptions struct {
	SpaceID          *uuid.UUID
	Product          *string
	Repository       *string
	RepositoryBranch *string
	Slug             *string
}

// ExtV1RunOrchestration Run orchestration
func (c *Client) ExtV1RunOrchestration(ctx context.Context, opts ExtV1RunOrchestrationOptions) (vampy.OrchestrationRunResult, error) {
	path := "/ext/v1/orchestration/run/"

	var out vampy.OrchestrationRunResult
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1RunOrchestrationOptions struct {
	Body vampy.OrchestrationRunBody1
}

// ExtV1RunScan Start Scanning for relation
func (c *Client) ExtV1RunScan(ctx context.Context, opts ExtV1RunScanOptions) (vampy.OrchestrationRunResult, error) {
	path := "/ext/v1/scan/"

	var out vampy.OrchestrationRunResult
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1RunScanOptions struct {
	Body vampy.V1ScanBody
}

// ExtV1ScanPipelineListGet Get scan pipelines list
func (c *Client) ExtV1ScanPipelineListGet(ctx context.Context, opts ExtV1ScanPipelineListGetOptions) (vampy.ScanPipelineListExtResponse, error) {
	path := "/ext/v1/scan/pipelines/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Repository != nil {
		q.Set("repository", *opts.Repository)
	}

	if opts.Product != nil {
		q.Set("product", *opts.Product)
	}

	if opts.RepositoryBranch != nil {
		q.Set("repositoryBranch", *opts.RepositoryBranch)
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	var out vampy.ScanPipelineListExtResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ScanPipelineListGetOptions struct {
	Repository       *string
	Product          *string
	RepositoryBranch *string
	Order            *string

	vampy.Paging
}

// ExtV1ScanPipelinePost Create scan pipeline
func (c *Client) ExtV1ScanPipelinePost(ctx context.Context, opts ExtV1ScanPipelinePostOptions) (vampy.ScanPipelineCreateExtResponse, error) {
	path := "/ext/v1/scan/pipelines/"

	var out vampy.ScanPipelineCreateExtResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ScanPipelinePostOptions struct {
	Body vampy.ScanPipelineCreateExt
}

// ExtV1ScanPipelineDetailsGet Get scan pipeline details
func (c *Client) ExtV1ScanPipelineDetailsGet(ctx context.Context, opts ExtV1ScanPipelineDetailsGetOptions) (vampy.ScanPipelineDetailsExt, error) {
	path := fmt.Sprintf("/ext/v1/scan/pipelines/%s/", url.PathEscape(opts.PipelineID))

	var out vampy.ScanPipelineDetailsExt
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ScanPipelineDetailsGetOptions struct {
	PipelineID string
}

// ExtV1ScanPipelineStatusGet Get scan pipeline status
func (c *Client) ExtV1ScanPipelineStatusGet(ctx context.Context, opts ExtV1ScanPipelineStatusGetOptions) (vampy.ScanPipelineStatusExt, error) {
	path := fmt.Sprintf("/ext/v1/scan/pipelines/%s/status/", url.PathEscape(opts.PipelineID))

	var out vampy.ScanPipelineStatusExt
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ScanPipelineStatusGetOptions struct {
	PipelineID string
}

// ExtV1ScanPipelineFinish Finish scan pipeline
func (c *Client) ExtV1ScanPipelineFinish(ctx context.Context, opts ExtV1ScanPipelineFinishOptions) error {
	path := fmt.Sprintf("/ext/v1/scan/pipelines/%s/finish/", url.PathEscape(opts.PipelineID))

	return c.doJSON(ctx, "POST", path, nil, nil, nil)
}

type ExtV1ScanPipelineFinishOptions struct {
	PipelineID string
}

// ExtV1GetParsers Get list of parsers (scanner types) for result upload
func (c *Client) ExtV1GetParsers(ctx context.Context) ([]vampy.ParserInList, error) {
	path := "/ext/v1/scanners/"

	var out []vampy.ParserInList
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// ExtV1GetScannersUpload Get list of scanners supporting result upload
func (c *Client) ExtV1GetScannersUpload(ctx context.Context) ([]vampy.ParserInList, error) {
	path := "/ext/v1/info/scanners/upload/"

	var out []vampy.ParserInList
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// ExtV1GetScannersOrchestration Get list of scanners available for orchestration
func (c *Client) ExtV1GetScannersOrchestration(ctx context.Context) ([]vampy.ScannerExternal, error) {
	path := "/ext/v1/info/scanners/orchestration/"

	var out []vampy.ScannerExternal
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// ExtV1GetBackgroundTasksDetails Detailed info about background task
func (c *Client) ExtV1GetBackgroundTasksDetails(ctx context.Context, opts ExtV1GetBackgroundTasksDetailsOptions) (vampy.BackgroundTaskDetails, error) {
	path := fmt.Sprintf("/ext/v1/background_tasks/%s/", url.PathEscape(opts.TaskID.String()))

	var out vampy.BackgroundTaskDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1GetBackgroundTasksDetailsOptions struct {
	TaskID uuid.UUID
}

// ExtV1ScanIssuesExport Export scan issues
func (c *Client) ExtV1ScanIssuesExport(ctx context.Context, opts ExtV1ScanIssuesExportOptions) (vampy.Page[vampy.ScanIssueExport], error) {
	path := "/ext/v1/scan_issues/export/"
	q := url.Values{}

	if opts.Payload != nil {
		q.Set("payload", string(*opts.Payload))
	}

	if opts.Formatter != nil {
		q.Set("formatter", string(*opts.Formatter))
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.SpaceID != nil {
		q.Set("spaceID", opts.SpaceID.String())
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.Statuses != nil {
		q.Set("statuses", *opts.Statuses)
	}

	if opts.Severities != nil {
		q.Set("severities", *opts.Severities)
	}

	if opts.Products != nil {
		q.Set("products", *opts.Products)
	}

	if opts.Repositories != nil {
		q.Set("repositories", *opts.Repositories)
	}

	if opts.RepositoryBranches != nil {
		q.Set("repositoryBranches", *opts.RepositoryBranches)
	}

	if opts.Scanners != nil {
		q.Set("scanners", *opts.Scanners)
	}

	if opts.ScannerTypes != nil {
		q.Set("scannerTypes", *opts.ScannerTypes)
	}

	if opts.HasSolution != nil {
		q.Set("hasSolution", strconv.FormatBool(*opts.HasSolution))
	}

	if opts.IncludeComments != nil {
		q.Set("includeComments", strconv.FormatBool(*opts.IncludeComments))
	}

	var out vampy.Page[vampy.ScanIssueExport]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ScanIssuesExportOptions struct {
	Payload            *vampy.ExportPayload
	Formatter          *vampy.ExportFormatter
	Order              *string
	SpaceID            *uuid.UUID
	Search             *string
	Statuses           *string
	Severities         *string
	Products           *string
	Repositories       *string
	RepositoryBranches *string
	Scanners           *string
	ScannerTypes       *string
	HasSolution        *bool
	IncludeComments    *bool
}

// ExtV1ScanProfilesListGet Get Scan Profiles List
func (c *Client) ExtV1ScanProfilesListGet(ctx context.Context, opts ExtV1ScanProfilesListGetOptions) (vampy.Page[vampy.OrchestrationScanProfile], error) {
	path := "/ext/v1/info/scan_profiles/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.OrchestrationScanProfile]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ScanProfilesListGetOptions struct {
	Order  *string
	Search *string

	vampy.Paging
}

// ExtV1QualityGatesListGet Get list of available Quality Gates
func (c *Client) ExtV1QualityGatesListGet(ctx context.Context, opts ExtV1QualityGatesListGetOptions) (vampy.Page[vampy.QualityGate], error) {
	path := "/ext/v1/info/quality_gates/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.QualityGate]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1QualityGatesListGetOptions struct {
	Order  *string
	Search *string

	vampy.Paging
}

// ExtV1ImportDefectDojo Import from Defect Dojo
func (c *Client) ExtV1ImportDefectDojo(ctx context.Context, opts ExtV1ImportDefectDojoOptions) (vampy.DefectDojoImportResponse, error) {
	path := "/ext/v1/import/defect_dojo/"

	var out vampy.DefectDojoImportResponse
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ExtV1ImportDefectDojoOptions struct {
	Body vampy.ImportDefectDojoBody
}
