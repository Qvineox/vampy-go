package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// ExtV1ProductListGet Get Products List
func (c *Client) ExtV1ProductListGet(ctx context.Context, limit *int, offset *int, order *string, search *string, repository *string, ids *string) (vampy.InlineResponse2007, error) {
	path := "/ext/v1/products/"
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	if repository != nil {
		q.Set("repository", *repository)
	}
	if ids != nil {
		q.Set("ids", *ids)
	}
	var out vampy.InlineResponse2007
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ArtifactVersionsSync Synchronize artifact versions from registry integrations
func (c *Client) ExtV1ArtifactVersionsSync(ctx context.Context, body vampy.ArtifactVersionSyncRequest) (vampy.ArtifactVersionSyncResponse, error) {
	path := "/ext/v1/artifacts/versions/sync/"
	var out vampy.ArtifactVersionSyncResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1RepositoryBranchesSync Synchronize repository branches from VCS integrations
func (c *Client) ExtV1RepositoryBranchesSync(ctx context.Context, body vampy.RepositoryBranchSyncRequest) (vampy.RepositoryBranchSyncResponse, error) {
	path := "/ext/v1/repositories/branches/sync/"
	var out vampy.RepositoryBranchSyncResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1RepositoryListGet Get Repositories List
func (c *Client) ExtV1RepositoryListGet(ctx context.Context, limit *int, offset *int, product *string, artifact *string, order *string, search *string, ids *string) (vampy.InlineResponse20012, error) {
	path := "/ext/v1/repositories/"
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if product != nil {
		q.Set("product", *product)
	}
	if artifact != nil {
		q.Set("artifact", *artifact)
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	if ids != nil {
		q.Set("ids", *ids)
	}
	var out vampy.InlineResponse20012
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanUploadPost External post new scanner result
func (c *Client) ExtV1ScanUploadPost(ctx context.Context, body vampy.V1ScanUploadsBody) (vampy.ScanUpload, error) {
	path := "/ext/v1/scan_uploads/"
	var out vampy.ScanUpload
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1QualityGateGet Get Quality Gate result for product or repository.
func (c *Client) ExtV1QualityGateGet(ctx context.Context, spaceID *string, product *string, repository *string, repositoryBranch *string, slug *string) (vampy.QualityGateResultDetails, error) {
	path := "/ext/v1/quality_gate/"
	q := url.Values{}
	if spaceID != nil {
		q.Set("spaceID", *spaceID)
	}
	if product != nil {
		q.Set("product", *product)
	}
	if repository != nil {
		q.Set("repository", *repository)
	}
	if repositoryBranch != nil {
		q.Set("repositoryBranch", *repositoryBranch)
	}
	if slug != nil {
		q.Set("slug", *slug)
	}
	var out vampy.QualityGateResultDetails
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1RunOrchestration Run orchestration
func (c *Client) ExtV1RunOrchestration(ctx context.Context, body vampy.OrchestrationRunBody1) (vampy.OrchestrationRunResult, error) {
	path := "/ext/v1/orchestration/run/"
	var out vampy.OrchestrationRunResult
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1RunScan Start Scanning for relation
func (c *Client) ExtV1RunScan(ctx context.Context, body vampy.V1ScanBody) (vampy.OrchestrationRunResult, error) {
	path := "/ext/v1/scan/"
	var out vampy.OrchestrationRunResult
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanPipelineListGet Get scan pipelines list
func (c *Client) ExtV1ScanPipelineListGet(ctx context.Context, repository *string, product *string, repositoryBranch *string, limit *int, offset *int, order *string) (vampy.ScanPipelineListExtResponse, error) {
	path := "/ext/v1/scan/pipelines/"
	q := url.Values{}
	if repository != nil {
		q.Set("repository", *repository)
	}
	if product != nil {
		q.Set("product", *product)
	}
	if repositoryBranch != nil {
		q.Set("repositoryBranch", *repositoryBranch)
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if order != nil {
		q.Set("order", *order)
	}
	var out vampy.ScanPipelineListExtResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanPipelinePost Create scan pipeline
func (c *Client) ExtV1ScanPipelinePost(ctx context.Context, body vampy.ScanPipelineCreateExt) (vampy.ScanPipelineCreateExtResponse, error) {
	path := "/ext/v1/scan/pipelines/"
	var out vampy.ScanPipelineCreateExtResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanPipelineDetailsGet Get scan pipeline details
func (c *Client) ExtV1ScanPipelineDetailsGet(ctx context.Context, pipelineID string) (vampy.ScanPipelineDetailsExt, error) {
	path := fmt.Sprintf("/ext/v1/scan/pipelines/%s/", url.PathEscape(pipelineID))
	var out vampy.ScanPipelineDetailsExt
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanPipelineStatusGet Get scan pipeline status
func (c *Client) ExtV1ScanPipelineStatusGet(ctx context.Context, pipelineID string) (vampy.ScanPipelineStatusExt, error) {
	path := fmt.Sprintf("/ext/v1/scan/pipelines/%s/status/", url.PathEscape(pipelineID))
	var out vampy.ScanPipelineStatusExt
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanPipelineFinish Finish scan pipeline
func (c *Client) ExtV1ScanPipelineFinish(ctx context.Context, pipelineID string) error {
	path := fmt.Sprintf("/ext/v1/scan/pipelines/%s/finish/", url.PathEscape(pipelineID))
	return c.doJSON(ctx, "POST", path, nil, nil, nil)
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
func (c *Client) ExtV1GetBackgroundTasksDetails(ctx context.Context, taskID string) (vampy.BackgroundTaskDetails, error) {
	path := fmt.Sprintf("/ext/v1/background_tasks/%s/", url.PathEscape(taskID))
	var out vampy.BackgroundTaskDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanIssuesExport Export scan issues
func (c *Client) ExtV1ScanIssuesExport(ctx context.Context, payload *vampy.ExportPayload, formatter *vampy.ExportFormatter, order *string, spaceID *string, search *string, statuses *string, severities *string, products *string, repositories *string, repositoryBranches *string, scanners *string, scannerTypes *string, hasSolution *bool, includeComments *bool) (vampy.InlineResponse20030, error) {
	path := "/ext/v1/scan_issues/export/"
	q := url.Values{}
	if payload != nil {
		q.Set("payload", string(*payload))
	}
	if formatter != nil {
		q.Set("formatter", string(*formatter))
	}
	if order != nil {
		q.Set("order", *order)
	}
	if spaceID != nil {
		q.Set("spaceID", *spaceID)
	}
	if search != nil {
		q.Set("search", *search)
	}
	if statuses != nil {
		q.Set("statuses", *statuses)
	}
	if severities != nil {
		q.Set("severities", *severities)
	}
	if products != nil {
		q.Set("products", *products)
	}
	if repositories != nil {
		q.Set("repositories", *repositories)
	}
	if repositoryBranches != nil {
		q.Set("repositoryBranches", *repositoryBranches)
	}
	if scanners != nil {
		q.Set("scanners", *scanners)
	}
	if scannerTypes != nil {
		q.Set("scannerTypes", *scannerTypes)
	}
	if hasSolution != nil {
		q.Set("hasSolution", strconv.FormatBool(*hasSolution))
	}
	if includeComments != nil {
		q.Set("includeComments", strconv.FormatBool(*includeComments))
	}
	var out vampy.InlineResponse20030
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ScanProfilesListGet Get Scan Profiles List
func (c *Client) ExtV1ScanProfilesListGet(ctx context.Context, limit *int, offset *int, order *string, search *string) (vampy.InlineResponse20019, error) {
	path := "/ext/v1/info/scan_profiles/"
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out vampy.InlineResponse20019
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1QualityGatesListGet Get list of available Quality Gates
func (c *Client) ExtV1QualityGatesListGet(ctx context.Context, limit *int, offset *int, order *string, search *string) (vampy.InlineResponse20031, error) {
	path := "/ext/v1/info/quality_gates/"
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out vampy.InlineResponse20031
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ExtV1ImportDefectDojo Import from Defect Dojo
func (c *Client) ExtV1ImportDefectDojo(ctx context.Context, body vampy.ImportDefectDojoBody) (vampy.DefectDojoImportResponse, error) {
	path := "/ext/v1/import/defect_dojo/"
	var out vampy.DefectDojoImportResponse
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}
