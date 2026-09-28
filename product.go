package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// ProductListGet Get list of available products.
func (c *Client) ProductListGet(ctx context.Context, spaceID string, limit *int, offset *int, order *string, search *string, repositoryID *string, ids *string) (vampy.InlineResponse2007, error) {
	path := fmt.Sprintf("/spaces/%s/products/", url.PathEscape(spaceID))
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
	if repositoryID != nil {
		q.Set("repositoryID", *repositoryID)
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

// ProductPost Create product
func (c *Client) ProductPost(ctx context.Context, spaceID string, body vampy.Product) (vampy.Product, error) {
	path := fmt.Sprintf("/spaces/%s/products/", url.PathEscape(spaceID))
	var out vampy.Product
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ProductDetailsGet Get detailed information about Product
func (c *Client) ProductDetailsGet(ctx context.Context, productID string) (vampy.ProductDetail, error) {
	path := fmt.Sprintf("/products/%s/", url.PathEscape(productID))
	var out vampy.ProductDetail
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ProductDelete Delete product instance and related data.
func (c *Client) ProductDelete(ctx context.Context, productID string) error {
	path := fmt.Sprintf("/products/%s/", url.PathEscape(productID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// ProductPatch Change product data.
func (c *Client) ProductPatch(ctx context.Context, productID string, body vampy.ProductEdit) (vampy.Product, error) {
	path := fmt.Sprintf("/products/%s/", url.PathEscape(productID))
	var out vampy.Product
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ProductUserGroupsGet Get user groups attached to product
func (c *Client) ProductUserGroupsGet(ctx context.Context, productID string, offset *int, limit *int, search *string) (vampy.InlineResponse20010, error) {
	path := fmt.Sprintf("/products/%s/user_groups/", url.PathEscape(productID))
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out vampy.InlineResponse20010
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ProductUsersSet [This API will be replaced  to user-groups logic in the end of Q2-2026] Add product users
func (c *Client) ProductUsersSet(ctx context.Context, productID string, body vampy.PermissionSetUsers) error {
	path := fmt.Sprintf("/products/%s/users/set/", url.PathEscape(productID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ProductUsersUnset [This API will be replaced  to user-groups logic in the end of Q2-2026] Remove product users
func (c *Client) ProductUsersUnset(ctx context.Context, productID string, body vampy.PermissionUnsetUsers) error {
	path := fmt.Sprintf("/products/%s/users/unset/", url.PathEscape(productID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ProductUserGroupsSet Add user groups to the product
func (c *Client) ProductUserGroupsSet(ctx context.Context, productID string, body vampy.PermissionSetUserGroups) error {
	path := fmt.Sprintf("/products/%s/user_groups/set/", url.PathEscape(productID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ProductUserGroupsUnset Remove user groups from the product
func (c *Client) ProductUserGroupsUnset(ctx context.Context, productID string, body vampy.PermissionUnsetUserGroups) error {
	path := fmt.Sprintf("/products/%s/user_groups/unset/", url.PathEscape(productID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ProductRelationsCreate Create requested products relations
func (c *Client) ProductRelationsCreate(ctx context.Context, productID string, body vampy.ProductChangeRelations) error {
	path := fmt.Sprintf("/products/%s/relations/create/", url.PathEscape(productID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ProductRelationsDelete Delete requested products relations
func (c *Client) ProductRelationsDelete(ctx context.Context, productID string, body vampy.ProductChangeRelations) error {
	path := fmt.Sprintf("/products/%s/relations/delete/", url.PathEscape(productID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}
