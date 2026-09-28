package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// AssetListGet Get list of available assets.
func (c *Client) AssetListGet(ctx context.Context, spaceID string, limit *int, offset *int, order *string, search *string, productID *string, parentID *string) (vampy.InlineResponse2008, error) {
	path := fmt.Sprintf("/spaces/%s/assets/", url.PathEscape(spaceID))
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
	if productID != nil {
		q.Set("productID", *productID)
	}
	if parentID != nil {
		q.Set("parentID", *parentID)
	}
	var out vampy.InlineResponse2008
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AssetPost Create asset
func (c *Client) AssetPost(ctx context.Context, spaceID string, body vampy.AssetCreate) (vampy.AssetDetails, error) {
	path := fmt.Sprintf("/spaces/%s/assets/", url.PathEscape(spaceID))
	var out vampy.AssetDetails
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AssetGet Get detailed information about Asset
func (c *Client) AssetGet(ctx context.Context, assetID string) (vampy.AssetDetails, error) {
	path := fmt.Sprintf("/assets/%s/", url.PathEscape(assetID))
	var out vampy.AssetDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AssetDelete Delete asset instance and related data.
func (c *Client) AssetDelete(ctx context.Context, assetID string) error {
	path := fmt.Sprintf("/assets/%s/", url.PathEscape(assetID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// AssetPatch Change asset data.
func (c *Client) AssetPatch(ctx context.Context, assetID string, body vampy.AssetPatch) (vampy.AssetDetails, error) {
	path := fmt.Sprintf("/assets/%s/", url.PathEscape(assetID))
	var out vampy.AssetDetails
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}
