package vampy_go

import (
	"context"
	"fmt"
	"net/url"

	"github.com/qvineox/vampy-go/vampy"
)

// OIDCRedirect Redirect to OIDC authorize URL
func (c *Client) OIDCRedirect(ctx context.Context, opts OIDCRedirectOptions) (vampy.OIDCRedirect, error) {
	path := fmt.Sprintf("/oidc/%s/redirect/", url.PathEscape(string(opts.Provider)))
	q := url.Values{}

	if opts.PostAuthURL != nil {
		q.Set("postAuthURL", *opts.PostAuthURL)
	}

	var out vampy.OIDCRedirect
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type OIDCRedirectOptions struct {
	Provider    vampy.OIDCProvider
	PostAuthURL *string
}

// OIDCCallback OIDC authorize user and redirect to post auth URL
func (c *Client) OIDCCallback(ctx context.Context, opts OIDCCallbackOptions) error {
	path := fmt.Sprintf("/oidc/%s/callback/", url.PathEscape(string(opts.Provider)))
	q := url.Values{}

	if opts.PostAuthURL != nil {
		q.Set("postAuthURL", *opts.PostAuthURL)
	}

	if opts.State != nil {
		q.Set("state", *opts.State)
	}

	if opts.SessionState != nil {
		q.Set("session_state", *opts.SessionState)
	}

	if opts.Iss != nil {
		q.Set("iss", *opts.Iss)
	}

	if opts.Code != nil {
		q.Set("code", *opts.Code)
	}

	return c.doJSON(ctx, "GET", path, q, nil, nil)
}

type OIDCCallbackOptions struct {
	Provider     vampy.OIDCProvider
	PostAuthURL  *string
	State        *string
	SessionState *string
	Iss          *string
	Code         *string
}
