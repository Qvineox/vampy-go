package vampy_go

import (
	"context"
	"fmt"
	"net/url"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// BotUserList Get list of bot users
func (c *Client) BotUserList(ctx context.Context) ([]vampy.BotUser, error) {
	path := "/bot_users/"

	var out []vampy.BotUser
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// BotUserRegistration Post new bot user on a server
func (c *Client) BotUserRegistration(ctx context.Context, opts BotUserRegistrationOptions) (vampy.BotUserToken, error) {
	path := "/bot_users/"

	var out vampy.BotUserToken
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type BotUserRegistrationOptions struct {
	Body vampy.BotUserRegistration
}

// BotUserDelete Delete bot user
func (c *Client) BotUserDelete(ctx context.Context, opts BotUserDeleteOptions) ([]vampy.BotUser, error) {
	path := fmt.Sprintf("/bot_users/%s/", url.PathEscape(opts.UserID.String()))

	var out []vampy.BotUser
	if err := c.doJSON(ctx, "DELETE", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type BotUserDeleteOptions struct {
	UserID uuid.UUID
}

// BotUserPatchToken Generate new token for bot-user
func (c *Client) BotUserPatchToken(ctx context.Context, opts BotUserPatchTokenOptions) (vampy.BotUserToken, error) {
	path := fmt.Sprintf("/bot_users/%s/token/", url.PathEscape(opts.UserID.String()))

	var out vampy.BotUserToken
	if err := c.doJSON(ctx, "PATCH", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type BotUserPatchTokenOptions struct {
	UserID uuid.UUID
}
