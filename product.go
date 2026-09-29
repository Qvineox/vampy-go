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

// ProductListGet Get list of available products.
func (c *Client) ProductListGet(ctx context.Context, opts ProductListGetOptions) (vampy.Page[vampy.Product], error) {
	path := fmt.Sprintf("/spaces/%s/products/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.RepositoryID != nil {
		q.Set("repositoryID", opts.RepositoryID.String())
	}

	if opts.IDs != nil {
		q.Set("ids", *opts.IDs)
	}

	var out vampy.Page[vampy.Product]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ProductListGetOptions struct {
	Order        *string
	Search       *string
	RepositoryID *uuid.UUID
	IDs          *string

	vampy.Paging
}

// ProductPost Create product
func (c *Client) ProductPost(ctx context.Context, opts ProductPostOptions) (vampy.Product, error) {
	path := fmt.Sprintf("/spaces/%s/products/", url.PathEscape(c.spaceID.String()))

	var out vampy.Product
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ProductPostOptions struct {
	Body vampy.Product
}

// ProductDetailsGet Get detailed information about Product
func (c *Client) ProductDetailsGet(ctx context.Context, opts ProductDetailsGetOptions) (vampy.ProductDetail, error) {
	path := fmt.Sprintf("/products/%s/", url.PathEscape(opts.ProductID.String()))

	var out vampy.ProductDetail
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ProductDetailsGetOptions struct {
	ProductID uuid.UUID
}

// ProductDelete Delete product instance and related data.
func (c *Client) ProductDelete(ctx context.Context, opts ProductDeleteOptions) error {
	path := fmt.Sprintf("/products/%s/", url.PathEscape(opts.ProductID.String()))

	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}

type ProductDeleteOptions struct {
	ProductID uuid.UUID
}

// ProductPatch Change product data.
func (c *Client) ProductPatch(ctx context.Context, opts ProductPatchOptions) (vampy.Product, error) {
	path := fmt.Sprintf("/products/%s/", url.PathEscape(opts.ProductID.String()))

	var out vampy.Product
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ProductPatchOptions struct {
	ProductID uuid.UUID
	Body      vampy.ProductEdit
}

// ProductUserGroupsGet Get user groups attached to product
func (c *Client) ProductUserGroupsGet(ctx context.Context, opts ProductUserGroupsGetOptions) (vampy.Page[vampy.UserGroupPermission], error) {
	path := fmt.Sprintf("/products/%s/user_groups/", url.PathEscape(opts.ProductID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.UserGroupPermission]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ProductUserGroupsGetOptions struct {
	ProductID uuid.UUID
	Search    *string

	vampy.Paging
}

// ProductUsersSet [This API will be replaced  to user-groups logic in the end of Q2-2026] Add product users
func (c *Client) ProductUsersSet(ctx context.Context, opts ProductUsersSetOptions) error {
	path := fmt.Sprintf("/products/%s/users/set/", url.PathEscape(opts.ProductID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ProductUsersSetOptions struct {
	ProductID uuid.UUID
	Body      vampy.PermissionSetUsers
}

// ProductUsersUnset [This API will be replaced  to user-groups logic in the end of Q2-2026] Remove product users
func (c *Client) ProductUsersUnset(ctx context.Context, opts ProductUsersUnsetOptions) error {
	path := fmt.Sprintf("/products/%s/users/unset/", url.PathEscape(opts.ProductID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ProductUsersUnsetOptions struct {
	ProductID uuid.UUID
	Body      vampy.PermissionUnsetUsers
}

// ProductUserGroupsSet Add user groups to the product
func (c *Client) ProductUserGroupsSet(ctx context.Context, opts ProductUserGroupsSetOptions) error {
	path := fmt.Sprintf("/products/%s/user_groups/set/", url.PathEscape(opts.ProductID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ProductUserGroupsSetOptions struct {
	ProductID uuid.UUID
	Body      vampy.PermissionSetUserGroups
}

// ProductUserGroupsUnset Remove user groups from the product
func (c *Client) ProductUserGroupsUnset(ctx context.Context, opts ProductUserGroupsUnsetOptions) error {
	path := fmt.Sprintf("/products/%s/user_groups/unset/", url.PathEscape(opts.ProductID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ProductUserGroupsUnsetOptions struct {
	ProductID uuid.UUID
	Body      vampy.PermissionUnsetUserGroups
}

// ProductRelationsCreate Create requested products relations
func (c *Client) ProductRelationsCreate(ctx context.Context, opts ProductRelationsCreateOptions) error {
	path := fmt.Sprintf("/products/%s/relations/create/", url.PathEscape(opts.ProductID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ProductRelationsCreateOptions struct {
	ProductID uuid.UUID
	Body      vampy.ProductChangeRelations
}

// ProductRelationsDelete Delete requested products relations
func (c *Client) ProductRelationsDelete(ctx context.Context, opts ProductRelationsDeleteOptions) error {
	path := fmt.Sprintf("/products/%s/relations/delete/", url.PathEscape(opts.ProductID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ProductRelationsDeleteOptions struct {
	ProductID uuid.UUID
	Body      vampy.ProductChangeRelations
}
