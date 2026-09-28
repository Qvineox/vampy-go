package vampy_go

import (
	"context"
	"github.com/qvineox/vampy-go/vampy"
)

// SpacesGet Get list of available spaces.
func (c *Client) SpacesGet(ctx context.Context) ([]vampy.Space, error) {
	path := "/spaces/"
	var out []vampy.Space
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SpaceCreate Create space
func (c *Client) SpaceCreate(ctx context.Context, body vampy.Space) (vampy.Space, error) {
	path := "/spaces/"
	var out vampy.Space
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}
