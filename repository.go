package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/qvineox/vampy-go/vampy"
)

// RepositoryListGet Get list of available repositories
func (c *Client) RepositoryListGet(ctx context.Context, opts RepositoryListGetOptions) (vampy.Page[vampy.Repository], error) {
	path := "/repositories/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.SpaceID != nil {
		q.Set("spaceID", opts.SpaceID.String())
	}

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.RequiredPermission != nil {
		q.Set("requiredPermission", *opts.RequiredPermission)
	}

	if opts.IDs != nil {
		q.Set("ids", *opts.IDs)
	}

	var out vampy.Page[vampy.Repository]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryListGetOptions struct {
	SpaceID            *uuid.UUID
	ProductID          *uuid.UUID
	Order              *string
	Search             *string
	RequiredPermission *string
	IDs                *string

	vampy.Paging
}

// RepositoryCreate Create repository
func (c *Client) RepositoryCreate(ctx context.Context, opts RepositoryCreateOptions) (vampy.Repository, error) {
	path := "/repositories/"

	var out vampy.Repository
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryCreateOptions struct {
	Body vampy.RepositoryCreate
}

// RepositoryImportFromVCS Import repositories from VCS
func (c *Client) RepositoryImportFromVCS(ctx context.Context, opts RepositoryImportFromVCSOptions) (vampy.RepositoryImportResponse, error) {
	path := "/repositories/vcs/import/"

	var out vampy.RepositoryImportResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryImportFromVCSOptions struct {
	Body vampy.RepositoryImportRequest
}

// RepositoryDetailsGet Get information about repository
func (c *Client) RepositoryDetailsGet(ctx context.Context, opts RepositoryDetailsGetOptions) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/", url.PathEscape(opts.RepositoryID.String()))

	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryDetailsGetOptions struct {
	RepositoryID uuid.UUID
}

// RepositoryDelete Patch repository
func (c *Client) RepositoryDelete(ctx context.Context, opts RepositoryDeleteOptions) error {
	path := fmt.Sprintf("/repositories/%s/", url.PathEscape(opts.RepositoryID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type RepositoryDeleteOptions struct {
	RepositoryID uuid.UUID
}

// RepositoryPatch Patch repository
func (c *Client) RepositoryPatch(ctx context.Context, opts RepositoryPatchOptions) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/", url.PathEscape(opts.RepositoryID.String()))

	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryPatchOptions struct {
	RepositoryID uuid.UUID
	Body         vampy.RepositoryPatch
}

// RepositoryUserGroupsGet Get user groups attached to repository
func (c *Client) RepositoryUserGroupsGet(ctx context.Context, opts RepositoryUserGroupsGetOptions) (vampy.Page[vampy.UserGroupPermission], error) {
	path := fmt.Sprintf("/repositories/%s/user_groups/", url.PathEscape(opts.RepositoryID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.UserGroupPermission]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryUserGroupsGetOptions struct {
	RepositoryID uuid.UUID
	Search       *string

	vampy.Paging
}

// RepositoryUsersSet [This API will be replaced  to user-groups logic in the end of Q2-2026] Add repository users
func (c *Client) RepositoryUsersSet(ctx context.Context, opts RepositoryUsersSetOptions) error {
	path := fmt.Sprintf("/repositories/%s/users/set/", url.PathEscape(opts.RepositoryID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type RepositoryUsersSetOptions struct {
	RepositoryID uuid.UUID
	Body         vampy.PermissionSetUsers
}

// RepositoryUsersUnset [This API will be replaced  to user-groups logic in the end of Q2-2026] Remove repository users
func (c *Client) RepositoryUsersUnset(ctx context.Context, opts RepositoryUsersUnsetOptions) error {
	path := fmt.Sprintf("/repositories/%s/users/unset/", url.PathEscape(opts.RepositoryID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type RepositoryUsersUnsetOptions struct {
	RepositoryID uuid.UUID
	Body         vampy.PermissionUnsetUsers
}

// RepositoryUserGroupsSet Add user groups to the repository
func (c *Client) RepositoryUserGroupsSet(ctx context.Context, opts RepositoryUserGroupsSetOptions) error {
	path := fmt.Sprintf("/repositories/%s/user_groups/set/", url.PathEscape(opts.RepositoryID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type RepositoryUserGroupsSetOptions struct {
	RepositoryID uuid.UUID
	Body         vampy.PermissionSetUserGroups
}

// RepositoryUserGroupsUnset Remove user groups from the repository
func (c *Client) RepositoryUserGroupsUnset(ctx context.Context, opts RepositoryUserGroupsUnsetOptions) error {
	path := fmt.Sprintf("/repositories/%s/user_groups/unset/", url.PathEscape(opts.RepositoryID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type RepositoryUserGroupsUnsetOptions struct {
	RepositoryID uuid.UUID
	Body         vampy.PermissionUnsetUserGroups
}

// RepositoryVCSLink Link repository with VCS
func (c *Client) RepositoryVCSLink(ctx context.Context, opts RepositoryVCSLinkOptions) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/vcs/link/", url.PathEscape(opts.RepositoryID.String()))

	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryVCSLinkOptions struct {
	RepositoryID uuid.UUID
	Body         vampy.RepositoryLink
}

// RepositoryVCSUnlink UnLink repository with VCS
func (c *Client) RepositoryVCSUnlink(ctx context.Context, opts RepositoryVCSUnlinkOptions) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/vcs/link/", url.PathEscape(opts.RepositoryID.String()))

	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "DELETE", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryVCSUnlinkOptions struct {
	RepositoryID uuid.UUID
}

// RepositoryBranchListGet Get list of repository branches
func (c *Client) RepositoryBranchListGet(ctx context.Context, opts RepositoryBranchListGetOptions) (vampy.Page[vampy.RepositoryBranchList], error) {
	path := fmt.Sprintf("/spaces/%s/repository_branches/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.IDs != nil {
		q.Set("ids", *opts.IDs)
	}

	if opts.RepositoryID != nil {
		q.Set("repositoryID", opts.RepositoryID.String())
	}

	if opts.ProductID != nil {
		q.Set("productID", opts.ProductID.String())
	}

	if opts.IsEmpty != nil {
		q.Set("isEmpty", strconv.FormatBool(*opts.IsEmpty))
	}

	var out vampy.Page[vampy.RepositoryBranchList]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryBranchListGetOptions struct {
	Search       *string
	Order        *string
	IDs          *string
	RepositoryID *uuid.UUID
	ProductID    *uuid.UUID
	IsEmpty      *bool

	vampy.Paging
}

// RepositoryBranchDelete Delete repository's branch
func (c *Client) RepositoryBranchDelete(ctx context.Context, opts RepositoryBranchDeleteOptions) error {
	path := fmt.Sprintf("/repositories/%s/branches/%s/", url.PathEscape(opts.RepositoryID.String()), url.PathEscape(opts.BranchID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type RepositoryBranchDeleteOptions struct {
	RepositoryID uuid.UUID
	BranchID     uuid.UUID
}
