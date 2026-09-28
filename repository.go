package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// RepositoryListGet Get list of available repositories
func (c *Client) RepositoryListGet(ctx context.Context, limit *int, offset *int, spaceID *string, productID *string, order *string, search *string, requiredPermission *string, ids *string) (vampy.Page[vampy.Repository], error) {
	path := "/repositories/"
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
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	if requiredPermission != nil {
		q.Set("requiredPermission", *requiredPermission)
	}
	if ids != nil {
		q.Set("ids", *ids)
	}
	var out vampy.Page[vampy.Repository]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryCreate Create repository
func (c *Client) RepositoryCreate(ctx context.Context, body vampy.RepositoryCreate) (vampy.Repository, error) {
	path := "/repositories/"
	var out vampy.Repository
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryImportFromVCS Import repositories from VCS
func (c *Client) RepositoryImportFromVCS(ctx context.Context, body vampy.RepositoryImportRequest) (vampy.RepositoryImportResponse, error) {
	path := "/repositories/vcs/import/"
	var out vampy.RepositoryImportResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryDetailsGet Get information about repository
func (c *Client) RepositoryDetailsGet(ctx context.Context, repositoryID string) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/", url.PathEscape(repositoryID))
	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryDelete Patch repository
func (c *Client) RepositoryDelete(ctx context.Context, repositoryID string) error {
	path := fmt.Sprintf("/repositories/%s/", url.PathEscape(repositoryID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// RepositoryPatch Patch repository
func (c *Client) RepositoryPatch(ctx context.Context, repositoryID string, body vampy.RepositoryPatch) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/", url.PathEscape(repositoryID))
	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryUserGroupsGet Get user groups attached to repository
func (c *Client) RepositoryUserGroupsGet(ctx context.Context, repositoryID string, offset *int, limit *int, search *string) (vampy.Page[vampy.UserGroupPermission], error) {
	path := fmt.Sprintf("/repositories/%s/user_groups/", url.PathEscape(repositoryID))
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
	var out vampy.Page[vampy.UserGroupPermission]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryUsersSet [This API will be replaced  to user-groups logic in the end of Q2-2026] Add repository users
func (c *Client) RepositoryUsersSet(ctx context.Context, repositoryID string, body vampy.PermissionSetUsers) error {
	path := fmt.Sprintf("/repositories/%s/users/set/", url.PathEscape(repositoryID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// RepositoryUsersUnset [This API will be replaced  to user-groups logic in the end of Q2-2026] Remove repository users
func (c *Client) RepositoryUsersUnset(ctx context.Context, repositoryID string, body vampy.PermissionUnsetUsers) error {
	path := fmt.Sprintf("/repositories/%s/users/unset/", url.PathEscape(repositoryID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// RepositoryUserGroupsSet Add user groups to the repository
func (c *Client) RepositoryUserGroupsSet(ctx context.Context, repositoryID string, body vampy.PermissionSetUserGroups) error {
	path := fmt.Sprintf("/repositories/%s/user_groups/set/", url.PathEscape(repositoryID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// RepositoryUserGroupsUnset Remove user groups from the repository
func (c *Client) RepositoryUserGroupsUnset(ctx context.Context, repositoryID string, body vampy.PermissionUnsetUserGroups) error {
	path := fmt.Sprintf("/repositories/%s/user_groups/unset/", url.PathEscape(repositoryID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// RepositoryVCSLink Link repository with VCS
func (c *Client) RepositoryVCSLink(ctx context.Context, repositoryID string, body vampy.RepositoryLink) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/vcs/link/", url.PathEscape(repositoryID))
	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryVCSUnlink UnLink repository with VCS
func (c *Client) RepositoryVCSUnlink(ctx context.Context, repositoryID string) (vampy.RepositoryDetails, error) {
	path := fmt.Sprintf("/repositories/%s/vcs/link/", url.PathEscape(repositoryID))
	var out vampy.RepositoryDetails
	if err := c.doJSON(ctx, "DELETE", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryBranchListGet Get list of repository branches
func (c *Client) RepositoryBranchListGet(ctx context.Context, spaceID string, limit *int, offset *int, search *string, order *string, ids *string, repositoryID *string, productID *string, isEmpty *bool) (vampy.Page[vampy.RepositoryBranchList], error) {
	path := fmt.Sprintf("/spaces/%s/repository_branches/", url.PathEscape(spaceID))
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if search != nil {
		q.Set("search", *search)
	}
	if order != nil {
		q.Set("order", *order)
	}
	if ids != nil {
		q.Set("ids", *ids)
	}
	if repositoryID != nil {
		q.Set("repositoryID", *repositoryID)
	}
	if productID != nil {
		q.Set("productID", *productID)
	}
	if isEmpty != nil {
		q.Set("isEmpty", strconv.FormatBool(*isEmpty))
	}
	var out vampy.Page[vampy.RepositoryBranchList]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryBranchDelete Delete repository's branch
func (c *Client) RepositoryBranchDelete(ctx context.Context, repositoryID string, branchID string) error {
	path := fmt.Sprintf("/repositories/%s/branches/%s/", url.PathEscape(repositoryID), url.PathEscape(branchID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}
