package vampy_go

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// QualityGateListGet Get list of available Quality Gates
func (c *Client) QualityGateListGet(ctx context.Context, opts QualityGateListGetOptions) ([]vampy.QualityGate, error) {
	path := fmt.Sprintf("/spaces/%s/quality_gates/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out []vampy.QualityGate
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type QualityGateListGetOptions struct {
	Order  *string
	Search *string
}

// QualityGatePost Create Quality Gate.
func (c *Client) QualityGatePost(ctx context.Context, opts QualityGatePostOptions) (vampy.QualityGate, error) {
	path := fmt.Sprintf("/spaces/%s/quality_gates/", url.PathEscape(c.spaceID.String()))

	var out vampy.QualityGate
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type QualityGatePostOptions struct {
	Body vampy.QualityGate
}

// QualityGateDetailsGet Get detailed information about Quality Gate.
func (c *Client) QualityGateDetailsGet(ctx context.Context, opts QualityGateDetailsGetOptions) (vampy.QualityGate, error) {
	path := fmt.Sprintf("/quality_gates/%s/", url.PathEscape(opts.QualityGateID.String()))

	var out vampy.QualityGate
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type QualityGateDetailsGetOptions struct {
	QualityGateID uuid.UUID
}

// QualityGateDelete Delete Quality Gate instance and related results.
func (c *Client) QualityGateDelete(ctx context.Context, opts QualityGateDeleteOptions) error {
	path := fmt.Sprintf("/quality_gates/%s/", url.PathEscape(opts.QualityGateID.String()))

	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}

type QualityGateDeleteOptions struct {
	QualityGateID uuid.UUID
}

// QualityGatePatch Change Quality Gate data.
func (c *Client) QualityGatePatch(ctx context.Context, opts QualityGatePatchOptions) (vampy.QualityGate, error) {
	path := fmt.Sprintf("/quality_gates/%s/", url.PathEscape(opts.QualityGateID.String()))

	var out vampy.QualityGate
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type QualityGatePatchOptions struct {
	QualityGateID uuid.UUID
	Body          vampy.QualityGateEdit
}

// QualityGateResultListGet Get list of available Quality Gate Results.
func (c *Client) QualityGateResultListGet(ctx context.Context, opts QualityGateResultListGetOptions) (vampy.Page[vampy.QualityGateResult], error) {
	path := fmt.Sprintf("/quality_gates/%s/results/", url.PathEscape(opts.QualityGateID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.RelationType != nil {
		q.Set("relationType", string(*opts.RelationType))
	}

	if opts.Statuses != nil {
		q.Set("statuses", *opts.Statuses)
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.QualityGateResult]
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type QualityGateResultListGetOptions struct {
	QualityGateID uuid.UUID
	Order         *string
	RelationType  *vampy.QualityGateRelationType
	Statuses      *string
	Search        *string

	vampy.Paging
}

// QualityGateResultGetDetails Get Quality Gate Result details.
func (c *Client) QualityGateResultGetDetails(ctx context.Context, opts QualityGateResultGetDetailsOptions) (vampy.QualityGateResult, error) {
	path := fmt.Sprintf("/quality_gates/%s/results/%s/", url.PathEscape(opts.QualityGateID.String()), url.PathEscape(opts.QualityGateResultID.String()))

	var out vampy.QualityGateResult
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type QualityGateResultGetDetailsOptions struct {
	QualityGateID       uuid.UUID
	QualityGateResultID uuid.UUID
}
