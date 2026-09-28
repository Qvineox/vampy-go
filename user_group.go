package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// GetUserGroups Get a list of user groups
func (c *Client) GetUserGroups(ctx context.Context, spaceID string, limit *int, offset *int, order *string, search *string, roleID *string, ldapConfigID *string, excludeLDAPConfigID *string, oidcConfigID *string, excludeOIDCConfigID *string) (vampy.InlineResponse2005, error) {
	path := fmt.Sprintf("/spaces/%s/user_groups/", url.PathEscape(spaceID))
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
	if roleID != nil {
		q.Set("roleID", *roleID)
	}
	if ldapConfigID != nil {
		q.Set("ldapConfigID", *ldapConfigID)
	}
	if excludeLDAPConfigID != nil {
		q.Set("excludeLDAPConfigID", *excludeLDAPConfigID)
	}
	if oidcConfigID != nil {
		q.Set("oidcConfigID", *oidcConfigID)
	}
	if excludeOIDCConfigID != nil {
		q.Set("excludeOIDCConfigID", *excludeOIDCConfigID)
	}
	var out vampy.InlineResponse2005
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// PostUserGroup Create new user group
func (c *Client) PostUserGroup(ctx context.Context, spaceID string, body vampy.UserGroupCreateEdit) (vampy.UserGroup, error) {
	path := fmt.Sprintf("/spaces/%s/user_groups/", url.PathEscape(spaceID))
	var out vampy.UserGroup
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetUserGroup Get details about a user group.
func (c *Client) GetUserGroup(ctx context.Context, userGroupID string) (vampy.UserGroup, error) {
	path := fmt.Sprintf("/user_groups/%s/", url.PathEscape(userGroupID))
	var out vampy.UserGroup
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DeleteUserGroup [This API will be replaced  to user-groups logic in the end of Q2-2026] Delete the user group(only admin)
func (c *Client) DeleteUserGroup(ctx context.Context, userGroupID string) error {
	path := fmt.Sprintf("/user_groups/%s/", url.PathEscape(userGroupID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// PatchUserGroup [This API will be replaced  to user-groups logic in the end of Q2-2026] Update the user group.
func (c *Client) PatchUserGroup(ctx context.Context, userGroupID string, body vampy.UserGroupCreateEdit) (vampy.UserGroup, error) {
	path := fmt.Sprintf("/user_groups/%s/", url.PathEscape(userGroupID))
	var out vampy.UserGroup
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetUserGroupRelations Get relations attached to user group
func (c *Client) GetUserGroupRelations(ctx context.Context, userGroupID string, relationType vampy.UserGroupRelationType, isMapped *bool, search *string, limit *int, offset *int) (vampy.InlineResponse2006, error) {
	path := fmt.Sprintf("/user_groups/%s/relations/", url.PathEscape(userGroupID))
	q := url.Values{}
	q.Set("relationType", string(relationType))
	if isMapped != nil {
		q.Set("isMapped", strconv.FormatBool(*isMapped))
	}
	if search != nil {
		q.Set("search", *search)
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	var out vampy.InlineResponse2006
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// UserGroupsSetRole Set role for multiple user groups
func (c *Client) UserGroupsSetRole(ctx context.Context, spaceID string, body vampy.UserGroupsSetRole) error {
	path := fmt.Sprintf("/spaces/%s/user_groups/set_role/", url.PathEscape(spaceID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// UserGroupMoveUsers Move users between user groups
func (c *Client) UserGroupMoveUsers(ctx context.Context, spaceID string, body vampy.UserGroupMoveUsers) error {
	path := fmt.Sprintf("/spaces/%s/user_groups/move_users/", url.PathEscape(spaceID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}
