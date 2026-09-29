package vampy_go

import (
	"context"
	"net/http"

	"github.com/qvineox/vampy-go/vampy"
)

// SpacesGet Get list of available spaces.
func (c *Client) SpacesGet(ctx context.Context) ([]vampy.Space, error) {
	path := "/spaces/"

	var out []vampy.Space
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// SpaceCreate Create space
func (c *Client) SpaceCreate(ctx context.Context, opts SpaceCreateOptions) (vampy.Space, error) {
	path := "/spaces/"

	var out vampy.Space
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceCreateOptions struct {
	Body vampy.Space
}
