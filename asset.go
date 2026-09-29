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

// AssetListGet Get list of available assets.
func (c *Client) AssetListGet(ctx context.Context, opts AssetListGetOptions) (vampy.Page[vampy.Asset], error) {
	path := fmt.Sprintf("/spaces/%s/assets/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.ParentID != nil {
		q.Set("parentID", opts.ParentID.String())
	}

	var out vampy.Page[vampy.Asset]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AssetListGetOptions struct {
	Order     *string
	Search    *string
	ProductID *uuid.UUID
	ParentID  *uuid.UUID

	vampy.Paging
}

// AssetPost Create asset
func (c *Client) AssetPost(ctx context.Context, opts AssetPostOptions) (vampy.AssetDetails, error) {
	path := fmt.Sprintf("/spaces/%s/assets/", url.PathEscape(c.spaceID.String()))

	var out vampy.AssetDetails
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AssetPostOptions struct {
	Body vampy.AssetCreate
}

// AssetGet Get detailed information about Asset
func (c *Client) AssetGet(ctx context.Context, opts AssetGetOptions) (vampy.AssetDetails, error) {
	path := fmt.Sprintf("/assets/%s/", url.PathEscape(opts.AssetID.String()))

	var out vampy.AssetDetails
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AssetGetOptions struct {
	AssetID uuid.UUID
}

// AssetDelete Delete asset instance and related data.
func (c *Client) AssetDelete(ctx context.Context, opts AssetDeleteOptions) error {
	path := fmt.Sprintf("/assets/%s/", url.PathEscape(opts.AssetID.String()))

	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}

type AssetDeleteOptions struct {
	AssetID uuid.UUID
}

// AssetPatch Change asset data.
func (c *Client) AssetPatch(ctx context.Context, opts AssetPatchOptions) (vampy.AssetDetails, error) {
	path := fmt.Sprintf("/assets/%s/", url.PathEscape(opts.AssetID.String()))

	var out vampy.AssetDetails
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AssetPatchOptions struct {
	AssetID uuid.UUID
	Body    vampy.AssetPatch
}
