package vampy_go

import (
	"context"
	"github.com/qvineox/vampy-go/vampy"
)

// LicenseGet Get current license information
func (c *Client) LicenseGet(ctx context.Context) (vampy.License, error) {
	path := "/license/"
	var out vampy.License
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// LicensePost Post new license
func (c *Client) LicensePost(ctx context.Context, body vampy.LicenseBody) error {
	path := "/license/"
	return c.doMultipart(ctx, "POST", path, nil, body, nil)
}

// LicenseDelete Delete license
func (c *Client) LicenseDelete(ctx context.Context) error {
	path := "/license/"
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// LicenseInformation Get information about all restrictions
func (c *Client) LicenseInformation(ctx context.Context) (vampy.LicenseInformation, error) {
	path := "/license/information/"
	var out vampy.LicenseInformation
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
