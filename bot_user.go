package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
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
func (c *Client) BotUserRegistration(ctx context.Context, body vampy.BotUserRegistration) (vampy.BotUserToken, error) {
	path := "/bot_users/"
	var out vampy.BotUserToken
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// BotUserDelete Delete bot user
func (c *Client) BotUserDelete(ctx context.Context, userID string) ([]vampy.BotUser, error) {
	path := fmt.Sprintf("/bot_users/%s/", url.PathEscape(userID))
	var out []vampy.BotUser
	if err := c.doJSON(ctx, "DELETE", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// BotUserPatchToken Generate new token for bot-user
func (c *Client) BotUserPatchToken(ctx context.Context, userID string) (vampy.BotUserToken, error) {
	path := fmt.Sprintf("/bot_users/%s/token/", url.PathEscape(userID))
	var out vampy.BotUserToken
	if err := c.doJSON(ctx, "PATCH", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
