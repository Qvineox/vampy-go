package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/qvineox/vampy-go/vampy"
)

// RulesListGet Rule: List of rules
func (c *Client) RulesListGet(ctx context.Context, opts RulesListGetOptions) (vampy.Page[vampy.RuleShort], error) {
	path := fmt.Sprintf("/spaces/%s/rules/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Source != nil {
		q.Set("source", string(*opts.Source))
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	if opts.HasWarnings != nil {
		q.Set("hasWarnings", strconv.FormatBool(*opts.HasWarnings))
	}

	var out vampy.Page[vampy.RuleShort]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RulesListGetOptions struct {
	Source      *vampy.RulesEngineCreationSource
	Search      *string
	HasWarnings *bool

	vampy.Paging
}

// RulesPost Rule: Create a new one
func (c *Client) RulesPost(ctx context.Context, opts RulesPostOptions) (vampy.RuleRead, error) {
	path := fmt.Sprintf("/spaces/%s/rules/", url.PathEscape(c.spaceID.String()))

	var out vampy.RuleRead
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RulesPostOptions struct {
	Body vampy.RuleCreate
}

// RulesCheck Rule: applicable instance count
func (c *Client) RulesCheck(ctx context.Context, opts RulesCheckOptions) (vampy.RulePassedInstances, error) {
	path := fmt.Sprintf("/spaces/%s/rules/check/", url.PathEscape(c.spaceID.String()))

	var out vampy.RulePassedInstances
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RulesCheckOptions struct {
	Body vampy.RuleCreate
}

// RulesCheckAsync Rule: applicable instance count (async)
func (c *Client) RulesCheckAsync(ctx context.Context, opts RulesCheckAsyncOptions) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/spaces/%s/rules/check/async/", url.PathEscape(c.spaceID.String()))

	var out vampy.BackgroundTask
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RulesCheckAsyncOptions struct {
	Body vampy.RuleCreate
}

// RulesEventsGet Rule: available events
func (c *Client) RulesEventsGet(ctx context.Context) ([]vampy.RulesSupportedEvents, error) {
	path := "/rules/events/"

	var out []vampy.RulesSupportedEvents
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// RulesDetailsGet Rule: get details
func (c *Client) RulesDetailsGet(ctx context.Context, opts RulesDetailsGetOptions) (vampy.RuleRead, error) {
	path := fmt.Sprintf("/rules/%s/", url.PathEscape(opts.RuleID.String()))

	var out vampy.RuleRead
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RulesDetailsGetOptions struct {
	RuleID uuid.UUID
}

// RulesDelete Rule: delete
func (c *Client) RulesDelete(ctx context.Context, opts RulesDeleteOptions) error {
	path := fmt.Sprintf("/rules/%s/", url.PathEscape(opts.RuleID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type RulesDeleteOptions struct {
	RuleID uuid.UUID
}

// RulesPatch Rule: patch specific one
func (c *Client) RulesPatch(ctx context.Context, opts RulesPatchOptions) (vampy.RuleRead, error) {
	path := fmt.Sprintf("/rules/%s/", url.PathEscape(opts.RuleID.String()))

	var out vampy.RuleRead
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RulesPatchOptions struct {
	RuleID uuid.UUID
	Body   vampy.RulePatch
}
