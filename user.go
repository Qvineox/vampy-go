package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/qvineox/vampy-go/vampy"
)

// UsersGet Get all registered users from a server.
func (c *Client) UsersGet(ctx context.Context, opts UsersGetOptions) ([]vampy.UserAdmin, error) {
	path := "/users/"
	q := url.Values{}

	if opts.IncludeBots != nil {
		q.Set("includeBots", strconv.FormatBool(*opts.IncludeBots))
	}

	if opts.RoleID != nil {
		q.Set("roleID", opts.RoleID.String())
	}

	if opts.GroupID != nil {
		q.Set("groupID", opts.GroupID.String())
	}

	var out []vampy.UserAdmin
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UsersGetOptions struct {
	IncludeBots *bool
	RoleID      *uuid.UUID
	GroupID     *uuid.UUID
}

// UserRegistration Post new user on a server. Can be accessed without login
func (c *Client) UserRegistration(ctx context.Context, opts UserRegistrationOptions) error {
	path := "/users/"

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type UserRegistrationOptions struct {
	Body vampy.UserRegistration
}

// SpaceUsersGet Get all registered users from a server
func (c *Client) SpaceUsersGet(ctx context.Context, opts SpaceUsersGetOptions) (vampy.Page[vampy.UserAdmin], error) {
	path := fmt.Sprintf("/spaces/%s/users/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.ReturnUsersType != nil {
		q.Set("returnUsersType", *opts.ReturnUsersType)
	}

	if opts.GroupID != nil {
		q.Set("groupID", opts.GroupID.String())
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.ExcludeGroupID != nil {
		q.Set("excludeGroupID", opts.ExcludeGroupID.String())
	}

	var out vampy.Page[vampy.UserAdmin]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceUsersGetOptions struct {
	ReturnUsersType *string
	GroupID         *uuid.UUID
	Search          *string
	ExcludeGroupID  *uuid.UUID

	vampy.Paging
}

// SpaceUserRegistration Post new user on a server
func (c *Client) SpaceUserRegistration(ctx context.Context, opts SpaceUserRegistrationOptions) error {
	path := fmt.Sprintf("/spaces/%s/users/", url.PathEscape(c.spaceID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type SpaceUserRegistrationOptions struct {
	Body vampy.UserRegistration
}

func (c *Client) UserUserIDGet(ctx context.Context, opts UserUserIDGetOptions) (vampy.UserAdmin, error) {
	path := fmt.Sprintf("/users/%s/", url.PathEscape(opts.UserID.String()))

	var out vampy.UserAdmin
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UserUserIDGetOptions struct {
	UserID uuid.UUID
}

func (c *Client) UsersPatch(ctx context.Context, opts UsersPatchOptions) error {
	path := fmt.Sprintf("/users/%s/", url.PathEscape(opts.UserID.String()))

	return c.doJSON(ctx, "PATCH", path, nil, opts.Body, nil)
}

type UsersPatchOptions struct {
	UserID uuid.UUID
	Body   vampy.UserAdmin
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

func (c *Client) UsersLogoutUser(ctx context.Context, opts UsersLogoutUserOptions) error {
	path := fmt.Sprintf("/users/%s/logout/", url.PathEscape(opts.UserID))

	return c.doJSON(ctx, "POST", path, nil, nil, nil)
}

type UsersLogoutUserOptions struct {
	UserID string
}

func (c *Client) UsersMFAPost(ctx context.Context, opts UsersMFAPostOptions) (vampy.UserMFA, error) {
	path := fmt.Sprintf("/users/%s/mfa/", url.PathEscape(opts.UserID))

	var out vampy.UserMFA
	if err := c.doJSON(ctx, "POST", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UsersMFAPostOptions struct {
	UserID string
}

func (c *Client) UsersMFADelete(ctx context.Context, opts UsersMFADeleteOptions) error {
	path := fmt.Sprintf("/users/%s/mfa/", url.PathEscape(opts.UserID))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type UsersMFADeleteOptions struct {
	UserID string
}

func (c *Client) UsersMFAConfirmPost(ctx context.Context, opts UsersMFAConfirmPostOptions) error {
	path := fmt.Sprintf("/users/%s/mfa/confirm/", url.PathEscape(opts.UserID))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type UsersMFAConfirmPostOptions struct {
	UserID string
	Body   vampy.MFAConfirmation
}
