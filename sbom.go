package vampy_go

import (
	"context"
	"fmt"
	"net/url"

	"github.com/qvineox/vampy-go/vampy"
)

// SBOMImport Import SBOM result file
func (c *Client) SBOMImport(ctx context.Context, opts SBOMImportOptions) (vampy.SBOMImport, error) {
	path := fmt.Sprintf("/space/%s/sbom/import/", url.PathEscape(c.spaceID.String()))

	var out vampy.SBOMImport
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SBOMImportOptions struct {
	Body vampy.SBOMImportBody
}

// SBOMExport Export SBOM to file
func (c *Client) SBOMExport(ctx context.Context, opts SBOMExportOptions) ([]byte, error) {
	path := fmt.Sprintf("/space/%s/sbom/export/", url.PathEscape(c.spaceID.String()))

	return c.doBinary(ctx, "POST", path, nil, opts.Body)
}

type SBOMExportOptions struct {
	Body vampy.SBOMExport
}
