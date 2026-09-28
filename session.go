package vampy_go

import (
	"context"
	"github.com/qvineox/vampy-go/vampy"
)

// SessionPost Logs in and returns the authentication cookie.
func (c *Client) SessionPost(ctx context.Context, body vampy.LoginRequest) (vampy.InlineResponse200, error) {
	path := "/session/"
	var out vampy.InlineResponse200
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SessionDelete Logs current user out
func (c *Client) SessionDelete(ctx context.Context) error {
	path := "/session/"
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}
