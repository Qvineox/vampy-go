package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
)

// OIDCRedirect Redirect to OIDC authorize URL
func (c *Client) OIDCRedirect(ctx context.Context, provider vampy.OIDCProvider, postAuthURL *string) (vampy.OIDCRedirect, error) {
	path := fmt.Sprintf("/oidc/%s/redirect/", url.PathEscape(string(provider)))
	q := url.Values{}
	if postAuthURL != nil {
		q.Set("postAuthURL", *postAuthURL)
	}
	var out vampy.OIDCRedirect
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// OIDCCallback OIDC authorize user and redirect to post auth URL
func (c *Client) OIDCCallback(ctx context.Context, provider vampy.OIDCProvider, postAuthURL *string, state *string, sessionState *string, iss *string, code *string) error {
	path := fmt.Sprintf("/oidc/%s/callback/", url.PathEscape(string(provider)))
	q := url.Values{}
	if postAuthURL != nil {
		q.Set("postAuthURL", *postAuthURL)
	}
	if state != nil {
		q.Set("state", *state)
	}
	if sessionState != nil {
		q.Set("session_state", *sessionState)
	}
	if iss != nil {
		q.Set("iss", *iss)
	}
	if code != nil {
		q.Set("code", *code)
	}
	return c.doJSON(ctx, "GET", path, q, nil, nil)
}
