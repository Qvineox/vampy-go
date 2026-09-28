package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
)

// SBOMImport Import SBOM result file
func (c *Client) SBOMImport(ctx context.Context, spaceID string, body vampy.SBOMImportBody) (vampy.SBOMImport, error) {
	path := fmt.Sprintf("/space/%s/sbom/import/", url.PathEscape(spaceID))
	var out vampy.SBOMImport
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SBOMExport Export SBOM to file
func (c *Client) SBOMExport(ctx context.Context, spaceID string, body vampy.SBOMExport) ([]byte, error) {
	path := fmt.Sprintf("/space/%s/sbom/export/", url.PathEscape(spaceID))
	return c.doBinary(ctx, "POST", path, nil, body)
}
