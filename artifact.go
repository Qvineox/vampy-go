package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/qvineox/vampy-go/vampy"
)

// ArtifactsListGet Get artifacts list
func (c *Client) ArtifactsListGet(ctx context.Context, opts ArtifactsListGetOptions) (vampy.Page[vampy.Artifact], error) {
	path := fmt.Sprintf("/spaces/%s/artifacts/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.Artifact]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ArtifactsListGetOptions struct {
	ProductID *uuid.UUID
	Order     *string
	Search    *string

	vampy.Paging
}

// ArtifactsDeleteBatch Artifact: group deletion
func (c *Client) ArtifactsDeleteBatch(ctx context.Context, opts ArtifactsDeleteBatchOptions) error {
	path := fmt.Sprintf("/spaces/%s/artifacts/delete/", url.PathEscape(c.spaceID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type ArtifactsDeleteBatchOptions struct {
	Body vampy.ArtifactDelete
}

// ArtifactsDetailsGet Get artifact
func (c *Client) ArtifactsDetailsGet(ctx context.Context, opts ArtifactsDetailsGetOptions) (vampy.Artifact, error) {
	path := fmt.Sprintf("/artifacts/%s/", url.PathEscape(opts.ArtifactID.String()))

	var out vampy.Artifact
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ArtifactsDetailsGetOptions struct {
	ArtifactID uuid.UUID
}

// ArtifactsPatch Get artifact
func (c *Client) ArtifactsPatch(ctx context.Context, opts ArtifactsPatchOptions) (vampy.ArtifactDetail, error) {
	path := fmt.Sprintf("/artifacts/%s/", url.PathEscape(opts.ArtifactID.String()))

	var out vampy.ArtifactDetail
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ArtifactsPatchOptions struct {
	ArtifactID uuid.UUID
	Body       vampy.ArtifactEdit
}

// ArtifactVersionsSync Synchronize artifact versions from registry integrations
func (c *Client) ArtifactVersionsSync(ctx context.Context, opts ArtifactVersionsSyncOptions) (vampy.ArtifactVersionSyncResponse, error) {
	path := fmt.Sprintf("/spaces/%s/artifacts/versions/sync/", url.PathEscape(c.spaceID.String()))

	var out vampy.ArtifactVersionSyncResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ArtifactVersionsSyncOptions struct {
	Body vampy.ArtifactVersionSyncRequest
}

// ArtifactsVersionsGet Get artifact versions
func (c *Client) ArtifactsVersionsGet(ctx context.Context, opts ArtifactsVersionsGetOptions) (vampy.Page[vampy.ArtifactVersion], error) {
	path := fmt.Sprintf("/artifacts/%s/versions/", url.PathEscape(opts.ArtifactID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	var out vampy.Page[vampy.ArtifactVersion]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ArtifactsVersionsGetOptions struct {
	ArtifactID uuid.UUID
	Search     *string
	Order      *string

	vampy.Paging
}

// ArtifactVersionsDeleteBatch ArtifactVersion: group deletion
func (c *Client) ArtifactVersionsDeleteBatch(ctx context.Context, opts ArtifactVersionsDeleteBatchOptions) error {
	path := fmt.Sprintf("/artifacts/%s/versions/delete/", url.PathEscape(opts.ArtifactID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type ArtifactVersionsDeleteBatchOptions struct {
	ArtifactID uuid.UUID
	Body       vampy.ArtifactVersionDelete
}

// ArtifactsImport Import artifacts from repository (integration)
func (c *Client) ArtifactsImport(ctx context.Context, opts ArtifactsImportOptions) (vampy.ArtifactImportTask, error) {
	path := fmt.Sprintf("/spaces/%s/artifacts/import/", url.PathEscape(c.spaceID.String()))

	var out vampy.ArtifactImportTask
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ArtifactsImportOptions struct {
	Body vampy.ArtifactImport
}
