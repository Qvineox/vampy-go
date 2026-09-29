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

// RunOrchestration Run orchestration
func (c *Client) RunOrchestration(ctx context.Context, opts RunOrchestrationOptions) (vampy.OrchestrationRunResult, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/run/", url.PathEscape(c.spaceID.String()))

	var out vampy.OrchestrationRunResult
	if err := c.doMultipart(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RunOrchestrationOptions struct {
	Body vampy.OrchestrationRunBody
}

// RunOrchestrationBatch Run orchestration batch
func (c *Client) RunOrchestrationBatch(ctx context.Context, opts RunOrchestrationBatchOptions) (vampy.OrchestrationRunBatchResult, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/run_batch/", url.PathEscape(c.spaceID.String()))

	var out vampy.OrchestrationRunBatchResult
	if err := c.doMultipart(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RunOrchestrationBatchOptions struct {
	Body vampy.OrchestrationRunBatchBody
}

// OrchestrationScanHistory Get history of scans
func (c *Client) OrchestrationScanHistory(ctx context.Context, opts OrchestrationScanHistoryOptions) (vampy.Page[vampy.OrchestrationScanDetail], error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/scan_history/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.RepositoryID != nil {
		q.Set("repositoryID", opts.RepositoryID.String())
	}

	if opts.RepositoryBranchID != nil {
		q.Set("repositoryBranchID", opts.RepositoryBranchID.String())
	}

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.ArtifactID != nil {
		q.Set("artifactID", opts.ArtifactID.String())
	}

	if opts.ArtifactVersionID != nil {
		q.Set("artifactVersionID", opts.ArtifactVersionID.String())
	}

	if opts.AssetID != nil {
		q.Set("assetID", opts.AssetID.String())
	}

	var out vampy.Page[vampy.OrchestrationScanDetail]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type OrchestrationScanHistoryOptions struct {
	RepositoryID       *uuid.UUID
	RepositoryBranchID *uuid.UUID
	ProductID          *uuid.UUID
	ArtifactID         *uuid.UUID
	ArtifactVersionID  *uuid.UUID
	AssetID            *uuid.UUID

	vampy.Paging
}

// GetOrchestrationProfiles List of orchestration profiles
func (c *Client) GetOrchestrationProfiles(ctx context.Context, opts GetOrchestrationProfilesOptions) (vampy.Page[vampy.OrchestrationScanProfile], error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/profiles/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.AvailableForRepositoryID != nil {
		q.Set("availableForRepositoryID", opts.AvailableForRepositoryID.String())
	}

	if opts.AvailableForProductID != nil {
		q.Set("availableForProductID", opts.AvailableForProductID.String())
	}

	if opts.AvailableForArtifactVersionID != nil {
		q.Set("availableForArtifactVersionID", opts.AvailableForArtifactVersionID.String())
	}

	if opts.AvailableForArtifactID != nil {
		q.Set("availableForArtifactID", opts.AvailableForArtifactID.String())
	}

	if opts.AvailableForAssetID != nil {
		q.Set("availableForAssetID", opts.AvailableForAssetID.String())
	}

	var out vampy.Page[vampy.OrchestrationScanProfile]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetOrchestrationProfilesOptions struct {
	AvailableForRepositoryID      *uuid.UUID
	AvailableForProductID         *uuid.UUID
	AvailableForArtifactVersionID *uuid.UUID
	AvailableForArtifactID        *uuid.UUID
	AvailableForAssetID           *uuid.UUID

	vampy.Paging
}

// CreateOrchestrationProfile Create orchestration profile
func (c *Client) CreateOrchestrationProfile(ctx context.Context, opts CreateOrchestrationProfileOptions) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/spaces/%s/orchestration/profiles/", url.PathEscape(c.spaceID.String()))

	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CreateOrchestrationProfileOptions struct {
	Body vampy.OrchestrationScanProfile
}

// GetOrchestrationProfile Get information about orchestration profile
func (c *Client) GetOrchestrationProfile(ctx context.Context, opts GetOrchestrationProfileOptions) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/", url.PathEscape(opts.ScanProfileID.String()))
	q := url.Values{}

	if opts.AvailableForRepositoryID != nil {
		q.Set("availableForRepositoryID", opts.AvailableForRepositoryID.String())
	}

	if opts.AvailableForProductID != nil {
		q.Set("availableForProductID", opts.AvailableForProductID.String())
	}

	if opts.AvailableForArtifactVersionID != nil {
		q.Set("availableForArtifactVersionID", opts.AvailableForArtifactVersionID.String())
	}

	if opts.AvailableForArtifactID != nil {
		q.Set("availableForArtifactID", opts.AvailableForArtifactID.String())
	}

	if opts.AvailableForAssetID != nil {
		q.Set("availableForAssetID", opts.AvailableForAssetID.String())
	}

	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetOrchestrationProfileOptions struct {
	ScanProfileID                 uuid.UUID
	AvailableForRepositoryID      *uuid.UUID
	AvailableForProductID         *uuid.UUID
	AvailableForArtifactVersionID *uuid.UUID
	AvailableForArtifactID        *uuid.UUID
	AvailableForAssetID           *uuid.UUID
}

// DeleteOrchestrationProfile Delete orchestration profile
func (c *Client) DeleteOrchestrationProfile(ctx context.Context, opts DeleteOrchestrationProfileOptions) error {
	path := fmt.Sprintf("/orchestration/profiles/%s/", url.PathEscape(opts.ScanProfileID.String()))

	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}

type DeleteOrchestrationProfileOptions struct {
	ScanProfileID uuid.UUID
}

// UpdateOrchestrationProfile Patch orchestration profile
func (c *Client) UpdateOrchestrationProfile(ctx context.Context, opts UpdateOrchestrationProfileOptions) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/", url.PathEscape(opts.ScanProfileID.String()))

	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UpdateOrchestrationProfileOptions struct {
	ScanProfileID uuid.UUID
	Body          vampy.OrchestrationScanProfile
}

// GetOrchestrationProfileRelations Get information about orchestration profile relations
func (c *Client) GetOrchestrationProfileRelations(ctx context.Context, opts GetOrchestrationProfileRelationsOptions) (vampy.Page[vampy.OrchestrationScanProfileRelation], error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/relations/", url.PathEscape(opts.ScanProfileID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	q.Set("relationType", string(opts.RelationType))

	var out vampy.Page[vampy.OrchestrationScanProfileRelation]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetOrchestrationProfileRelationsOptions struct {
	ScanProfileID uuid.UUID
	RelationType  vampy.OrchestrationScanProfileRelationType

	vampy.Paging
}

// DeleteScannerConfig Delete scanner config from profile
func (c *Client) DeleteScannerConfig(ctx context.Context, opts DeleteScannerConfigOptions) (vampy.OrchestrationScanProfile, error) {
	path := fmt.Sprintf("/orchestration/profiles/%s/delete_scanner_config/", url.PathEscape(opts.ScanProfileID.String()))

	var out vampy.OrchestrationScanProfile
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type DeleteScannerConfigOptions struct {
	ScanProfileID uuid.UUID
	Body          vampy.OrchestrationScannerConfigDelete
}

// GetOrchestrationScannersInfo Get a list of available scanners with their parameters
func (c *Client) GetOrchestrationScannersInfo(ctx context.Context, opts GetOrchestrationScannersInfoOptions) ([]vampy.OrchestrationScannerInfo, error) {
	path := "/orchestration/scanners/"
	q := url.Values{}

	if opts.AvailableForRepositoryID != nil {
		q.Set("availableForRepositoryID", opts.AvailableForRepositoryID.String())
	}

	if opts.AvailableForProductID != nil {
		q.Set("availableForProductID", opts.AvailableForProductID.String())
	}

	if opts.AvailableForArtifactVersionID != nil {
		q.Set("availableForArtifactVersionID", opts.AvailableForArtifactVersionID.String())
	}

	if opts.AvailableForArtifactID != nil {
		q.Set("availableForArtifactID", opts.AvailableForArtifactID.String())
	}

	if opts.AvailableForAssetID != nil {
		q.Set("availableForAssetID", opts.AvailableForAssetID.String())
	}

	var out []vampy.OrchestrationScannerInfo
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetOrchestrationScannersInfoOptions struct {
	AvailableForRepositoryID      *uuid.UUID
	AvailableForProductID         *uuid.UUID
	AvailableForArtifactVersionID *uuid.UUID
	AvailableForArtifactID        *uuid.UUID
	AvailableForAssetID           *uuid.UUID
}

// UploadScannerConfig Upload scanner config file
func (c *Client) UploadScannerConfig(ctx context.Context, opts UploadScannerConfigOptions) (vampy.OrchestrationScannerConfigFileInfo, error) {
	path := "/orchestration/upload_scanner_config/"

	var out vampy.OrchestrationScannerConfigFileInfo
	if err := c.doMultipart(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UploadScannerConfigOptions struct {
	Body vampy.OrchestrationUploadScannerConfigBody
}

// DownloadScannerConfig Download scanner config file
func (c *Client) DownloadScannerConfig(ctx context.Context, opts DownloadScannerConfigOptions) ([]byte, error) {
	path := "/orchestration/download_scanner_config/"
	q := url.Values{}

	q.Set("scanner", string(opts.Scanner))

	if opts.FilePath != nil {
		q.Set("filePath", *opts.FilePath)
	}

	if opts.FileName != nil {
		q.Set("fileName", *opts.FileName)
	}

	if opts.Default != nil {
		q.Set("default", strconv.FormatBool(*opts.Default))
	}

	return c.doBinary(ctx, http.MethodGet, path, q, nil)
}

type DownloadScannerConfigOptions struct {
	Scanner  vampy.ScannerWithConfigFiles
	FilePath *string
	FileName *string
	Default  *bool
}
