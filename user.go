package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// UsersGet Get all registered users from a server.
func (c *Client) UsersGet(ctx context.Context, includeBots *bool, roleID *string, groupID *string) ([]vampy.UserAdmin, error) {
	path := "/users/"
	q := url.Values{}
	if includeBots != nil {
		q.Set("includeBots", strconv.FormatBool(*includeBots))
	}
	if roleID != nil {
		q.Set("roleID", *roleID)
	}
	if groupID != nil {
		q.Set("groupID", *groupID)
	}
	var out []vampy.UserAdmin
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// UserRegistration Post new user on a server. Can be accessed without login
func (c *Client) UserRegistration(ctx context.Context, body vampy.UserRegistration) error {
	path := "/users/"
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// SpaceUsersGet Get all registered users from a server
func (c *Client) SpaceUsersGet(ctx context.Context, spaceID string, returnUsersType *string, groupID *string, offset *int, limit *int, search *string, excludeGroupID *string) (vampy.InlineResponse2001, error) {
	path := fmt.Sprintf("/spaces/%s/users/", url.PathEscape(spaceID))
	q := url.Values{}
	if returnUsersType != nil {
		q.Set("returnUsersType", *returnUsersType)
	}
	if groupID != nil {
		q.Set("groupID", *groupID)
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if search != nil {
		q.Set("search", *search)
	}
	if excludeGroupID != nil {
		q.Set("excludeGroupID", *excludeGroupID)
	}
	var out vampy.InlineResponse2001
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SpaceUserRegistration Post new user on a server
func (c *Client) SpaceUserRegistration(ctx context.Context, spaceID string, body vampy.UserRegistration) error {
	path := fmt.Sprintf("/spaces/%s/users/", url.PathEscape(spaceID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

func (c *Client) UserUserIDGet(ctx context.Context, userID string) (vampy.UserAdmin, error) {
	path := fmt.Sprintf("/users/%s/", url.PathEscape(userID))
	var out vampy.UserAdmin
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

func (c *Client) UsersPatch(ctx context.Context, userID string, body vampy.UserAdmin) error {
	path := fmt.Sprintf("/users/%s/", url.PathEscape(userID))
	return c.doJSON(ctx, "PATCH", path, nil, body, nil)
}

// UserCurrent Get current user info.
func (c *Client) UserCurrent(ctx context.Context) (vampy.UserShort, error) {
	path := "/users/me/"
	var out vampy.UserShort
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

func (c *Client) UsersLogoutUser(ctx context.Context, userID string) error {
	path := fmt.Sprintf("/users/%s/logout/", url.PathEscape(userID))
	return c.doJSON(ctx, "POST", path, nil, nil, nil)
}

func (c *Client) UsersMFAPost(ctx context.Context, userID string) (vampy.UserMFA, error) {
	path := fmt.Sprintf("/users/%s/mfa/", url.PathEscape(userID))
	var out vampy.UserMFA
	if err := c.doJSON(ctx, "POST", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

func (c *Client) UsersMFADelete(ctx context.Context, userID string) error {
	path := fmt.Sprintf("/users/%s/mfa/", url.PathEscape(userID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

func (c *Client) UsersMFAConfirmPost(ctx context.Context, userID string, body vampy.MFAConfirmation) error {
	path := fmt.Sprintf("/users/%s/mfa/confirm/", url.PathEscape(userID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}
