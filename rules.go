package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// RulesListGet Rule: List of rules
func (c *Client) RulesListGet(ctx context.Context, spaceID string, offset *int, limit *int, source *vampy.RulesEngineCreationSource, search *string, hasWarnings *bool) (vampy.InlineResponse20017, error) {
	path := fmt.Sprintf("/spaces/%s/rules/", url.PathEscape(spaceID))
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if source != nil {
		q.Set("source", string(*source))
	}
	if search != nil {
		q.Set("search", *search)
	}
	if hasWarnings != nil {
		q.Set("hasWarnings", strconv.FormatBool(*hasWarnings))
	}
	var out vampy.InlineResponse20017
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RulesPost Rule: Create a new one
func (c *Client) RulesPost(ctx context.Context, spaceID string, body vampy.RuleCreate) (vampy.RuleRead, error) {
	path := fmt.Sprintf("/spaces/%s/rules/", url.PathEscape(spaceID))
	var out vampy.RuleRead
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RulesCheck Rule: applicable instance count
func (c *Client) RulesCheck(ctx context.Context, spaceID string, body vampy.RuleCreate) (vampy.RulePassedInstances, error) {
	path := fmt.Sprintf("/spaces/%s/rules/check/", url.PathEscape(spaceID))
	var out vampy.RulePassedInstances
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RulesCheckAsync Rule: applicable instance count (async)
func (c *Client) RulesCheckAsync(ctx context.Context, spaceID string, body vampy.RuleCreate) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/spaces/%s/rules/check/async/", url.PathEscape(spaceID))
	var out vampy.BackgroundTask
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
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
func (c *Client) RulesDetailsGet(ctx context.Context, ruleID string) (vampy.RuleRead, error) {
	path := fmt.Sprintf("/rules/%s/", url.PathEscape(ruleID))
	var out vampy.RuleRead
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RulesDelete Rule: delete
func (c *Client) RulesDelete(ctx context.Context, ruleID string) error {
	path := fmt.Sprintf("/rules/%s/", url.PathEscape(ruleID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// RulesPatch Rule: patch specific one
func (c *Client) RulesPatch(ctx context.Context, ruleID string, body vampy.RulePatch) (vampy.RuleRead, error) {
	path := fmt.Sprintf("/rules/%s/", url.PathEscape(ruleID))
	var out vampy.RuleRead
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}
