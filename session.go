package vampy_go

import (
	"context"

	"github.com/qvineox/vampy-go/vampy"
)

// SessionPost Logs in and returns the authentication cookie.
func (c *Client) SessionPost(ctx context.Context, opts SessionPostOptions) (vampy.LoginResult, error) {
	path := "/session/"

	var out vampy.LoginResult
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SessionPostOptions struct {
	Body vampy.LoginRequest
}

// SessionDelete Logs current user out
func (c *Client) SessionDelete(ctx context.Context) error {
	path := "/session/"

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}
