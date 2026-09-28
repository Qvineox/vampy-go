package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// ScanIssueCreate Create ScanIssue
func (c *Client) ScanIssueCreate(ctx context.Context, body vampy.ScanIssueCreate) (vampy.ScanIssueCreateResponse, error) {
	path := "/scan_issues/"
	var out vampy.ScanIssueCreateResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssuesFilter Filter scan issues by params
func (c *Client) ScanIssuesFilter(ctx context.Context, offset *int, limit *int, order *string, body vampy.ScanIssueFilter) (vampy.Page[vampy.ScanIssueWithFPScore], error) {
	path := "/scan_issues/filter/"
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if order != nil {
		q.Set("order", *order)
	}
	var out vampy.Page[vampy.ScanIssueWithFPScore]
	if err := c.doJSON(ctx, "POST", path, q, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssuesAsyncExport Async export scan issues
func (c *Client) ScanIssuesAsyncExport(ctx context.Context, formatter *vampy.AsyncExportFormatter, order *string, includeComments *bool, templateSlug *vampy.TemplateSlug, csvDialect *vampy.CSVExportDialect, csvDelimiter *vampy.CSVExportDelimiter, csvDoubleQuote *string, csvLineTerminator *string, csvQuoteChar *string, csvEncoding *vampy.CSVExportEncoding, csvEscapeChar *string, csvQuoting *vampy.CSVExportQuoting, body vampy.ScanIssueFilter) (vampy.Report, error) {
	path := "/scan_issues/async_export/"
	q := url.Values{}
	if formatter != nil {
		q.Set("formatter", string(*formatter))
	}
	if order != nil {
		q.Set("order", *order)
	}
	if includeComments != nil {
		q.Set("includeComments", strconv.FormatBool(*includeComments))
	}
	if templateSlug != nil {
		q.Set("templateSlug", string(*templateSlug))
	}
	if csvDialect != nil {
		q.Set("csvDialect", string(*csvDialect))
	}
	if csvDelimiter != nil {
		q.Set("csvDelimiter", string(*csvDelimiter))
	}
	if csvDoubleQuote != nil {
		q.Set("csvDoubleQuote", *csvDoubleQuote)
	}
	if csvLineTerminator != nil {
		q.Set("csvLineTerminator", *csvLineTerminator)
	}
	if csvQuoteChar != nil {
		q.Set("csvQuoteChar", *csvQuoteChar)
	}
	if csvEncoding != nil {
		q.Set("csvEncoding", string(*csvEncoding))
	}
	if csvEscapeChar != nil {
		q.Set("csvEscapeChar", *csvEscapeChar)
	}
	if csvQuoting != nil {
		q.Set("csvQuoting", string(*csvQuoting))
	}
	var out vampy.Report
	if err := c.doJSON(ctx, "POST", path, q, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssuesExportMeta Get CSV export metadata
func (c *Client) ScanIssuesExportMeta(ctx context.Context, body vampy.ScanIssueFilter) (vampy.CSVExportMeta, error) {
	path := "/scan_issues/export/meta/"
	var out vampy.CSVExportMeta
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssueSearchDuplicates Search ScanIssue's duplicates
func (c *Client) ScanIssueSearchDuplicates(ctx context.Context, body vampy.ScanIssueCreate) (vampy.ScanIssueDuplicatesSearchResult, error) {
	path := "/scan_issues/search_duplicates/"
	var out vampy.ScanIssueDuplicatesSearchResult
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssuesDelete ScanIssue: group deletion
func (c *Client) ScanIssuesDelete(ctx context.Context, body vampy.ScanIssueDelete) error {
	path := "/scan_issues/delete/"
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ScanIssueSchemasGet Get scan_issue schemas by each scan_issue_type
func (c *Client) ScanIssueSchemasGet(ctx context.Context) ([]vampy.ScanIssues, error) {
	path := "/scan_issues/schemas/"
	var out []vampy.ScanIssues
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssueDetailsGet Get information about ScanIssue
func (c *Client) ScanIssueDetailsGet(ctx context.Context, scanIssueID string) (vampy.ScanIssueDetail, error) {
	path := fmt.Sprintf("/scan_issues/%s/", url.PathEscape(scanIssueID))
	var out vampy.ScanIssueDetail
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssueStatusChangesGet Get history of status changes
func (c *Client) ScanIssueStatusChangesGet(ctx context.Context, scanIssueID string) ([]vampy.ScanIssueStatusChangeEntry, error) {
	path := fmt.Sprintf("/scan_issues/%s/status_changes/", url.PathEscape(scanIssueID))
	var out []vampy.ScanIssueStatusChangeEntry
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssueChangesGet Get history of scan issue changes (status, severity, FP score)
func (c *Client) ScanIssueChangesGet(ctx context.Context, scanIssueID string, changeTypes []string) ([]vampy.ScanIssueChangeEntry, error) {
	path := fmt.Sprintf("/scan_issues/%s/changes/", url.PathEscape(scanIssueID))
	q := url.Values{}
	for _, v := range changeTypes {
		q.Add("changeTypes", v)
	}
	var out []vampy.ScanIssueChangeEntry
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssueAIRequestPost Request AI suggestion
func (c *Client) ScanIssueAIRequestPost(ctx context.Context, scanIssueID string, body vampy.ScanIssuesAIRequest) (vampy.ScanIssueAIRequestResponse, error) {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/", url.PathEscape(scanIssueID))
	var out vampy.ScanIssueAIRequestResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssueAIRequestCancel Cancel AI requests of specific type
func (c *Client) ScanIssueAIRequestCancel(ctx context.Context, scanIssueID string, body vampy.ScanIssueAIRequestCancel) error {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/cancel/", url.PathEscape(scanIssueID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// ScanIssueAIRequestDetailsGet Get AI request details
func (c *Client) ScanIssueAIRequestDetailsGet(ctx context.Context, scanIssueID string, aiRequestID string) (vampy.ScanIssueAIRequestResponse, error) {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/%s/", url.PathEscape(scanIssueID), url.PathEscape(aiRequestID))
	var out vampy.ScanIssueAIRequestResponse
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssueAIRequestApply Apply AI request
func (c *Client) ScanIssueAIRequestApply(ctx context.Context, scanIssueID string, aiRequestID string, body vampy.ScanIssueAIRequestApply) (vampy.ScanIssueAIRequestResponse, error) {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/%s/apply/", url.PathEscape(scanIssueID), url.PathEscape(aiRequestID))
	var out vampy.ScanIssueAIRequestResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScanIssuesUpdate Update batch of scan issues
func (c *Client) ScanIssuesUpdate(ctx context.Context, body vampy.ScanIssuesUpdate) (any, error) {
	path := "/scan_issues/update/"
	var out any
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}
