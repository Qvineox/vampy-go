package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// QualityGateListGet Get list of available Quality Gates
func (c *Client) QualityGateListGet(ctx context.Context, spaceID string, order *string, search *string) ([]vampy.QualityGate, error) {
	path := fmt.Sprintf("/spaces/%s/quality_gates/", url.PathEscape(spaceID))
	q := url.Values{}
	if order != nil {
		q.Set("order", *order)
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out []vampy.QualityGate
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// QualityGatePost Create Quality Gate.
func (c *Client) QualityGatePost(ctx context.Context, spaceID string, body vampy.QualityGate) (vampy.QualityGate, error) {
	path := fmt.Sprintf("/spaces/%s/quality_gates/", url.PathEscape(spaceID))
	var out vampy.QualityGate
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// QualityGateDetailsGet Get detailed information about Quality Gate.
func (c *Client) QualityGateDetailsGet(ctx context.Context, qualityGateID string) (vampy.QualityGate, error) {
	path := fmt.Sprintf("/quality_gates/%s/", url.PathEscape(qualityGateID))
	var out vampy.QualityGate
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// QualityGateDelete Delete Quality Gate instance and related results.
func (c *Client) QualityGateDelete(ctx context.Context, qualityGateID string) error {
	path := fmt.Sprintf("/quality_gates/%s/", url.PathEscape(qualityGateID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// QualityGatePatch Change Quality Gate data.
func (c *Client) QualityGatePatch(ctx context.Context, qualityGateID string, body vampy.QualityGateEdit) (vampy.QualityGate, error) {
	path := fmt.Sprintf("/quality_gates/%s/", url.PathEscape(qualityGateID))
	var out vampy.QualityGate
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// QualityGateResultListGet Get list of available Quality Gate Results.
func (c *Client) QualityGateResultListGet(ctx context.Context, qualityGateID string, limit *int, offset *int, order *string, relationType *vampy.QualityGateRelationType, statuses *string, search *string) (vampy.InlineResponse2009, error) {
	path := fmt.Sprintf("/quality_gates/%s/results/", url.PathEscape(qualityGateID))
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if order != nil {
		q.Set("order", *order)
	}
	if relationType != nil {
		q.Set("relationType", string(*relationType))
	}
	if statuses != nil {
		q.Set("statuses", *statuses)
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out vampy.InlineResponse2009
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// QualityGateResultGetDetails Get Quality Gate Result details.
func (c *Client) QualityGateResultGetDetails(ctx context.Context, qualityGateID string, qualityGateResultID string) (vampy.QualityGateResultDetails, error) {
	path := fmt.Sprintf("/quality_gates/%s/results/%s/", url.PathEscape(qualityGateID), url.PathEscape(qualityGateResultID))
	var out vampy.QualityGateResultDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
