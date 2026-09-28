package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// ComponentsListGet Get list of available components
func (c *Client) ComponentsListGet(ctx context.Context, limit *int, offset *int, spaceID *string, productID *string, ids *string, order *string, search *string) (vampy.Page[vampy.ComponentShort], error) {
	path := "/components/"
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if spaceID != nil {
		q.Set("spaceID", *spaceID)
	}
	if productID != nil {
		q.Set("productID", *productID)
	}
	if ids != nil {
		q.Set("ids", *ids)
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out vampy.Page[vampy.ComponentShort]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ComponentDetailsGet Get detailed information about component
func (c *Client) ComponentDetailsGet(ctx context.Context, componentID string) (vampy.ComponentDetails, error) {
	path := fmt.Sprintf("/components/%s/", url.PathEscape(componentID))
	var out vampy.ComponentDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ComponentPatch Patch component
func (c *Client) ComponentPatch(ctx context.Context, componentID string, body vampy.ComponentPatch) error {
	path := fmt.Sprintf("/components/%s/", url.PathEscape(componentID))
	return c.doJSON(ctx, "PATCH", path, nil, body, nil)
}

// ComponentVersionsListGet Get list of available components' versions
func (c *Client) ComponentVersionsListGet(ctx context.Context, limit *int, offset *int, spaceID *string, repositoryID *string, componentID *string, productID *string, ids *string, order *string, search *string, skipAggregations *bool) (vampy.Page[vampy.ComponentVersion], error) {
	path := "/component_versions/"
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if spaceID != nil {
		q.Set("spaceID", *spaceID)
	}
	if repositoryID != nil {
		q.Set("repositoryID", *repositoryID)
	}
	if componentID != nil {
		q.Set("componentID", *componentID)
	}
	if productID != nil {
		q.Set("productID", *productID)
	}
	if ids != nil {
		q.Set("ids", *ids)
	}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	if skipAggregations != nil {
		q.Set("skipAggregations", strconv.FormatBool(*skipAggregations))
	}
	var out vampy.Page[vampy.ComponentVersion]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ComponentVersionDetailsGet Get detailed information about component version
func (c *Client) ComponentVersionDetailsGet(ctx context.Context, componentVersionID string, repositoryID *string, productID *string) (vampy.ComponentVersionDetails, error) {
	path := fmt.Sprintf("/component_versions/%s/", url.PathEscape(componentVersionID))
	q := url.Values{}
	if repositoryID != nil {
		q.Set("repositoryID", *repositoryID)
	}
	if productID != nil {
		q.Set("productID", *productID)
	}
	var out vampy.ComponentVersionDetails
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ComponentVersionPatch Patch component version
func (c *Client) ComponentVersionPatch(ctx context.Context, componentVersionID string, body vampy.ComponentPatch) error {
	path := fmt.Sprintf("/component_versions/%s/", url.PathEscape(componentVersionID))
	return c.doJSON(ctx, "PATCH", path, nil, body, nil)
}
