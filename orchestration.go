package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// RunOrchestration Run orchestration
func (c *Client) RunOrchestration(ctx context.Context, spaceID string, body vampy.OrchestrationRunBody) (vampy.OrchestrationRunResult, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/run/", url.PathEscape(spaceID))
	var out vampy.OrchestrationRunResult
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RunOrchestrationBatch Run orchestration batch
func (c *Client) RunOrchestrationBatch(ctx context.Context, spaceID string, body vampy.OrchestrationRunBatchBody) (vampy.OrchestrationRunBatchResult, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/run_batch/", url.PathEscape(spaceID))
	var out vampy.OrchestrationRunBatchResult
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// OrchestrationScanHistory Get history of scans
func (c *Client) OrchestrationScanHistory(ctx context.Context, spaceID string, offset *int, limit *int, repositoryID *string, repositoryBranchID *string, productID *string, artifactID *string, artifactVersionID *string, assetID *string) (vampy.InlineResponse20018, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/scan_history/", url.PathEscape(spaceID))
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if repositoryID != nil {
		q.Set("repositoryID", *repositoryID)
	}
	if repositoryBranchID != nil {
		q.Set("repositoryBranchID", *repositoryBranchID)
	}
	if productID != nil {
		q.Set("productID", *productID)
	}
	if artifactID != nil {
		q.Set("artifactID", *artifactID)
	}
	if artifactVersionID != nil {
		q.Set("artifactVersionID", *artifactVersionID)
	}
	if assetID != nil {
		q.Set("assetID", *assetID)
	}
	var out vampy.InlineResponse20018
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetOrchestrationProfiles List of orchestration profiles
func (c *Client) GetOrchestrationProfiles(ctx context.Context, spaceID string, availableForRepositoryID *string, availableForProductID *string, availableForArtifactVersionID *string, availableForArtifactID *string, availableForAssetID *string, offset *int, limit *int) (vampy.InlineResponse20019, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/profiles/", url.PathEscape(spaceID))
	q := url.Values{}
	if availableForRepositoryID != nil {
		q.Set("availableForRepositoryID", *availableForRepositoryID)
	}
	if availableForProductID != nil {
		q.Set("availableForProductID", *availableForProductID)
	}
	if availableForArtifactVersionID != nil {
		q.Set("availableForArtifactVersionID", *availableForArtifactVersionID)
	}
	if availableForArtifactID != nil {
		q.Set("availableForArtifactID", *availableForArtifactID)
	}
	if availableForAssetID != nil {
		q.Set("availableForAssetID", *availableForAssetID)
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	var out vampy.InlineResponse20019
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CreateOrchestrationProfile Create orchestration profile
func (c *Client) CreateOrchestrationProfile(ctx context.Context, spaceID string, body vampy.OrchestrationScanProfile) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/profiles/", url.PathEscape(spaceID))
	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetOrchestrationProfile Get information about orchestration profile
func (c *Client) GetOrchestrationProfile(ctx context.Context, scanProfileID string, availableForRepositoryID *string, availableForProductID *string, availableForArtifactVersionID *string, availableForArtifactID *string, availableForAssetID *string) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/", url.PathEscape(scanProfileID))
	q := url.Values{}
	if availableForRepositoryID != nil {
		q.Set("availableForRepositoryID", *availableForRepositoryID)
	}
	if availableForProductID != nil {
		q.Set("availableForProductID", *availableForProductID)
	}
	if availableForArtifactVersionID != nil {
		q.Set("availableForArtifactVersionID", *availableForArtifactVersionID)
	}
	if availableForArtifactID != nil {
		q.Set("availableForArtifactID", *availableForArtifactID)
	}
	if availableForAssetID != nil {
		q.Set("availableForAssetID", *availableForAssetID)
	}
	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DeleteOrchestrationProfile Delete orchestration profile
func (c *Client) DeleteOrchestrationProfile(ctx context.Context, scanProfileID string) error {
	path := fmt.Sprintf("/orchestration/profiles/%s/", url.PathEscape(scanProfileID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// UpdateOrchestrationProfile Patch orchestration profile
func (c *Client) UpdateOrchestrationProfile(ctx context.Context, scanProfileID string, body vampy.OrchestrationScanProfile) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/", url.PathEscape(scanProfileID))
	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetOrchestrationProfileRelations Get information about orchestration profile relations
func (c *Client) GetOrchestrationProfileRelations(ctx context.Context, scanProfileID string, relationType vampy.OrchestrationScanProfileRelationType, offset *int, limit *int) (vampy.InlineResponse20020, error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/relations/", url.PathEscape(scanProfileID))
	q := url.Values{}
	q.Set("relationType", string(relationType))
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	var out vampy.InlineResponse20020
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DeleteScannerConfig Delete scanner config from profile
func (c *Client) DeleteScannerConfig(ctx context.Context, scanProfileID string, body vampy.OrchestrationScannerConfigDelete) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/delete_scanner_config/", url.PathEscape(scanProfileID))
	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetOrchestrationScannersInfo Get a list of available scanners with their parameters
func (c *Client) GetOrchestrationScannersInfo(ctx context.Context, availableForRepositoryID *string, availableForProductID *string, availableForArtifactVersionID *string, availableForArtifactID *string, availableForAssetID *string) ([]vampy.OrchestrationScannerInfo, error) {
	path := "/orchestration/scanners/"
	q := url.Values{}
	if availableForRepositoryID != nil {
		q.Set("availableForRepositoryID", *availableForRepositoryID)
	}
	if availableForProductID != nil {
		q.Set("availableForProductID", *availableForProductID)
	}
	if availableForArtifactVersionID != nil {
		q.Set("availableForArtifactVersionID", *availableForArtifactVersionID)
	}
	if availableForArtifactID != nil {
		q.Set("availableForArtifactID", *availableForArtifactID)
	}
	if availableForAssetID != nil {
		q.Set("availableForAssetID", *availableForAssetID)
	}
	var out []vampy.OrchestrationScannerInfo
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// UploadScannerConfig Upload scanner config file
func (c *Client) UploadScannerConfig(ctx context.Context, body vampy.OrchestrationUploadScannerConfigBody) (vampy.OrchestrationScannerConfigFileInfo, error) {
	path := "/orchestration/upload_scanner_config/"
	var out vampy.OrchestrationScannerConfigFileInfo
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DownloadScannerConfig Download scanner config file
func (c *Client) DownloadScannerConfig(ctx context.Context, scanner vampy.ScannerWithConfigFiles, filePath *string, fileName *string, default_ *bool) ([]byte, error) {
	path := "/orchestration/download_scanner_config/"
	q := url.Values{}
	q.Set("scanner", string(scanner))
	if filePath != nil {
		q.Set("filePath", *filePath)
	}
	if fileName != nil {
		q.Set("fileName", *fileName)
	}
	if default_ != nil {
		q.Set("default", strconv.FormatBool(*default_))
	}
	return c.doBinary(ctx, "GET", path, q, nil)
}
