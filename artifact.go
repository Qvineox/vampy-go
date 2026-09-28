package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// ArtifactsListGet Get artifacts list
func (c *Client) ArtifactsListGet(ctx context.Context, spaceID string, productID *string, order *string, search *string, offset *int, limit *int) (vampy.Page[vampy.Artifact], error) {
	path := fmt.Sprintf("/spaces/%s/artifacts/", url.PathEscape(spaceID))
	q := url.Values{}
	if productID != nil {
		q.Set("productID", *productID)
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	var out vampy.Page[vampy.Artifact]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ArtifactsDeleteBatch Artifact: group deletion
func (c *Client) ArtifactsDeleteBatch(ctx context.Context, spaceID string, body vampy.ArtifactDelete) error {
	path := fmt.Sprintf("/spaces/%s/artifacts/delete/", url.PathEscape(spaceID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ArtifactsDetailsGet Get artifact
func (c *Client) ArtifactsDetailsGet(ctx context.Context, artifactID string) (vampy.Artifact, error) {
	path := fmt.Sprintf("/artifacts/%s/", url.PathEscape(artifactID))
	var out vampy.Artifact
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ArtifactsPatch Get artifact
func (c *Client) ArtifactsPatch(ctx context.Context, artifactID string, body vampy.ArtifactEdit) (vampy.ArtifactDetail, error) {
	path := fmt.Sprintf("/artifacts/%s/", url.PathEscape(artifactID))
	var out vampy.ArtifactDetail
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ArtifactVersionsSync Synchronize artifact versions from registry integrations
func (c *Client) ArtifactVersionsSync(ctx context.Context, spaceID string, body vampy.ArtifactVersionSyncRequest) (vampy.ArtifactVersionSyncResponse, error) {
	path := fmt.Sprintf("/spaces/%s/artifacts/versions/sync/", url.PathEscape(spaceID))
	var out vampy.ArtifactVersionSyncResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ArtifactsVersionsGet Get artifact versions
func (c *Client) ArtifactsVersionsGet(ctx context.Context, artifactID string, search *string, offset *int, limit *int, order *string) (vampy.Page[vampy.ArtifactVersion], error) {
	path := fmt.Sprintf("/artifacts/%s/versions/", url.PathEscape(artifactID))
	q := url.Values{}
	if search != nil {
		q.Set("search", *search)
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if order != nil {
		q.Set("order", *order)
	}
	var out vampy.Page[vampy.ArtifactVersion]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ArtifactVersionsDeleteBatch ArtifactVersion: group deletion
func (c *Client) ArtifactVersionsDeleteBatch(ctx context.Context, artifactID string, body vampy.ArtifactVersionDelete) error {
	path := fmt.Sprintf("/artifacts/%s/versions/delete/", url.PathEscape(artifactID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ArtifactsImport Import artifacts from repository (integration)
func (c *Client) ArtifactsImport(ctx context.Context, spaceID string, body vampy.ArtifactImport) (vampy.ArtifactImportTask, error) {
	path := fmt.Sprintf("/spaces/%s/artifacts/import/", url.PathEscape(spaceID))
	var out vampy.ArtifactImportTask
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}
