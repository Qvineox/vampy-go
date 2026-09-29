package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// ComponentsListGet Get list of available components
func (c *Client) ComponentsListGet(ctx context.Context, opts ComponentsListGetOptions) (vampy.Page[vampy.ComponentShort], error) {
	path := "/components/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.SpaceID != nil {
		q.Set("spaceID", opts.SpaceID.String())
	}

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.IDs != nil {
		q.Set("ids", *opts.IDs)
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.ComponentShort]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ComponentsListGetOptions struct {
	SpaceID   *uuid.UUID
	ProductID *uuid.UUID
	IDs       *string
	Order     *string
	Search    *string

	vampy.Paging
}

// ComponentDetailsGet Get detailed information about component
func (c *Client) ComponentDetailsGet(ctx context.Context, opts ComponentDetailsGetOptions) (vampy.ComponentDetails, error) {
	path := fmt.Sprintf("/components/%s/", url.PathEscape(opts.ComponentID.String()))

	var out vampy.ComponentDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ComponentDetailsGetOptions struct {
	ComponentID uuid.UUID
}

// ComponentPatch Patch component
func (c *Client) ComponentPatch(ctx context.Context, opts ComponentPatchOptions) error {
	path := fmt.Sprintf("/components/%s/", url.PathEscape(opts.ComponentID.String()))

	return c.doJSON(ctx, "PATCH", path, nil, opts.Body, nil)
}

type ComponentPatchOptions struct {
	ComponentID uuid.UUID
	Body        vampy.ComponentPatch
}

// ComponentVersionsListGet Get list of available components' versions
func (c *Client) ComponentVersionsListGet(ctx context.Context, opts ComponentVersionsListGetOptions) (vampy.Page[vampy.ComponentVersion], error) {
	path := "/component_versions/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.SpaceID != nil {
		q.Set("spaceID", opts.SpaceID.String())
	}

	if opts.RepositoryID != nil {
		q.Set("repositoryID", opts.RepositoryID.String())
	}

	if opts.ComponentID != nil {
		q.Set("componentID", opts.ComponentID.String())
	}

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.IDs != nil {
		q.Set("ids", *opts.IDs)
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.SkipAggregations != nil {
		q.Set("skipAggregations", strconv.FormatBool(*opts.SkipAggregations))
	}

	var out vampy.Page[vampy.ComponentVersion]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ComponentVersionsListGetOptions struct {
	SpaceID          *uuid.UUID
	RepositoryID     *uuid.UUID
	ComponentID      *uuid.UUID
	ProductID        *uuid.UUID
	IDs              *string
	Order            *string
	Search           *string
	SkipAggregations *bool

	vampy.Paging
}

// ComponentVersionDetailsGet Get detailed information about component version
func (c *Client) ComponentVersionDetailsGet(ctx context.Context, opts ComponentVersionDetailsGetOptions) (vampy.ComponentVersionDetails, error) {
	path := fmt.Sprintf("/component_versions/%s/", url.PathEscape(opts.ComponentVersionID.String()))
	q := url.Values{}

	if opts.RepositoryID != nil {
		q.Set("repositoryID", opts.RepositoryID.String())
	}

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	var out vampy.ComponentVersionDetails
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ComponentVersionDetailsGetOptions struct {
	ComponentVersionID uuid.UUID
	RepositoryID       *uuid.UUID
	ProductID          *uuid.UUID
}

// ComponentVersionPatch Patch component version
func (c *Client) ComponentVersionPatch(ctx context.Context, opts ComponentVersionPatchOptions) error {
	path := fmt.Sprintf("/component_versions/%s/", url.PathEscape(opts.ComponentVersionID.String()))

	return c.doJSON(ctx, "PATCH", path, nil, opts.Body, nil)
}

type ComponentVersionPatchOptions struct {
	ComponentVersionID uuid.UUID
	Body               vampy.ComponentPatch
}
