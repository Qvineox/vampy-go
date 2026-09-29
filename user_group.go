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

// GetUserGroups Get a list of user groups
func (c *Client) GetUserGroups(ctx context.Context, opts GetUserGroupsOptions) (vampy.Page[vampy.UserGroup], error) {
	path := fmt.Sprintf("/spaces/%s/user_groups/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.RoleID != nil {
		q.Set("roleID", opts.RoleID.String())
	}

	if opts.LDAPConfigID != nil {
		q.Set("ldapConfigID", opts.LDAPConfigID.String())
	}

	if opts.ExcludeLDAPConfigID != nil {
		q.Set("excludeLDAPConfigID", opts.ExcludeLDAPConfigID.String())
	}

	if opts.OIDCConfigID != nil {
		q.Set("oidcConfigID", opts.OIDCConfigID.String())
	}

	if opts.ExcludeOIDCConfigID != nil {
		q.Set("excludeOIDCConfigID", opts.ExcludeOIDCConfigID.String())
	}

	var out vampy.Page[vampy.UserGroup]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetUserGroupsOptions struct {
	Order               *string
	Search              *string
	RoleID              *uuid.UUID
	LDAPConfigID        *uuid.UUID
	ExcludeLDAPConfigID *uuid.UUID
	OIDCConfigID        *uuid.UUID
	ExcludeOIDCConfigID *uuid.UUID

	vampy.Paging
}

// PostUserGroup Create new user group
func (c *Client) PostUserGroup(ctx context.Context, opts PostUserGroupOptions) (vampy.UserGroup, error) {
	path := fmt.Sprintf("/spaces/%s/user_groups/", url.PathEscape(c.spaceID.String()))

	var out vampy.UserGroup
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type PostUserGroupOptions struct {
	Body vampy.UserGroupCreateEdit
}

// GetUserGroup Get details about a user group.
func (c *Client) GetUserGroup(ctx context.Context, opts GetUserGroupOptions) (vampy.UserGroup, error) {
	path := fmt.Sprintf("/user_groups/%s/", url.PathEscape(opts.UserGroupID.String()))

	var out vampy.UserGroup
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetUserGroupOptions struct {
	UserGroupID uuid.UUID
}

// DeleteUserGroup [This API will be replaced  to user-groups logic in the end of Q2-2026] Delete the user group(only admin)
func (c *Client) DeleteUserGroup(ctx context.Context, opts DeleteUserGroupOptions) error {
	path := fmt.Sprintf("/user_groups/%s/", url.PathEscape(opts.UserGroupID.String()))

	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}

type DeleteUserGroupOptions struct {
	UserGroupID uuid.UUID
}

// PatchUserGroup [This API will be replaced  to user-groups logic in the end of Q2-2026] Update the user group.
func (c *Client) PatchUserGroup(ctx context.Context, opts PatchUserGroupOptions) (vampy.UserGroup, error) {
	path := fmt.Sprintf("/user_groups/%s/", url.PathEscape(opts.UserGroupID.String()))

	var out vampy.UserGroup
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type PatchUserGroupOptions struct {
	UserGroupID uuid.UUID
	Body        vampy.UserGroupCreateEdit
}

// GetUserGroupRelations Get relations attached to user group
func (c *Client) GetUserGroupRelations(ctx context.Context, opts GetUserGroupRelationsOptions) (vampy.Page[vampy.UserGroupRelation], error) {
	path := fmt.Sprintf("/user_groups/%s/relations/", url.PathEscape(opts.UserGroupID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	q.Set("relationType", string(opts.RelationType))

	if opts.IsMapped != nil {
		q.Set("isMapped", strconv.FormatBool(*opts.IsMapped))
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.UserGroupRelation]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetUserGroupRelationsOptions struct {
	UserGroupID  uuid.UUID
	RelationType vampy.UserGroupRelationType
	IsMapped     *bool
	Search       *string

	vampy.Paging
}

// UserGroupsSetRole Set role for multiple user groups
func (c *Client) UserGroupsSetRole(ctx context.Context, opts UserGroupsSetRoleOptions) error {
	path := fmt.Sprintf("/spaces/%s/user_groups/set_role/", url.PathEscape(c.spaceID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type UserGroupsSetRoleOptions struct {
	Body vampy.UserGroupsSetRole
}

// UserGroupMoveUsers Move users between user groups
func (c *Client) UserGroupMoveUsers(ctx context.Context, opts UserGroupMoveUsersOptions) error {
	path := fmt.Sprintf("/spaces/%s/user_groups/move_users/", url.PathEscape(c.spaceID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type UserGroupMoveUsersOptions struct {
	Body vampy.UserGroupMoveUsers
}
