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

// ScanIssueCreate Create ScanIssue
func (c *Client) ScanIssueCreate(ctx context.Context, opts ScanIssueCreateOptions) (vampy.ScanIssueCreateResponse, error) {
	path := "/scan_issues/"

	var out vampy.ScanIssueCreateResponse
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueCreateOptions struct {
	Body vampy.ScanIssueCreate
}

// ScanIssuesFilter Filter scan issues by params
func (c *Client) ScanIssuesFilter(ctx context.Context, opts ScanIssuesFilterOptions) (vampy.Page[vampy.ScanIssueWithFPScore], error) {
	path := "/scan_issues/filter/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	var out vampy.Page[vampy.ScanIssueWithFPScore]
	if err := c.doJSON(ctx, http.MethodPost, path, q, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssuesFilterOptions struct {
	Order *string
	Body  vampy.ScanIssueFilter

	vampy.Paging
}

// ScanIssuesAsyncExport Async export scan issues
func (c *Client) ScanIssuesAsyncExport(ctx context.Context, opts ScanIssuesAsyncExportOptions) (vampy.Report, error) {
	path := "/scan_issues/async_export/"
	q := url.Values{}

	if opts.Formatter != nil {
		q.Set("formatter", string(*opts.Formatter))
	}

	if opts.Order != nil {
		q.Set("order", *opts.Order)
	}

	if opts.IncludeComments != nil {
		q.Set("includeComments", strconv.FormatBool(*opts.IncludeComments))
	}

	if opts.TemplateSlug != nil {
		q.Set("templateSlug", string(*opts.TemplateSlug))
	}

	if opts.CSVDialect != nil {
		q.Set("csvDialect", string(*opts.CSVDialect))
	}

	if opts.CSVDelimiter != nil {
		q.Set("csvDelimiter", string(*opts.CSVDelimiter))
	}

	if opts.CSVDoubleQuote != nil {
		q.Set("csvDoubleQuote", strconv.FormatBool(*opts.CSVDoubleQuote))
	}

	if opts.CSVLineTerminator != nil {
		q.Set("csvLineTerminator", *opts.CSVLineTerminator)
	}

	if opts.CSVQuoteChar != nil {
		q.Set("csvQuoteChar", *opts.CSVQuoteChar)
	}

	if opts.CSVEncoding != nil {
		q.Set("csvEncoding", string(*opts.CSVEncoding))
	}

	if opts.CSVEscapeChar != nil {
		q.Set("csvEscapeChar", *opts.CSVEscapeChar)
	}

	if opts.CSVQuoting != nil {
		q.Set("csvQuoting", string(*opts.CSVQuoting))
	}

	var out vampy.Report
	if err := c.doJSON(ctx, http.MethodPost, path, q, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssuesAsyncExportOptions struct {
	Formatter         *vampy.AsyncExportFormatter
	Order             *string
	IncludeComments   *bool
	TemplateSlug      *vampy.TemplateSlug
	CSVDialect        *vampy.CSVExportDialect
	CSVDelimiter      *vampy.CSVExportDelimiter
	CSVDoubleQuote    *bool
	CSVLineTerminator *string
	CSVQuoteChar      *string
	CSVEncoding       *vampy.CSVExportEncoding
	CSVEscapeChar     *string
	CSVQuoting        *vampy.CSVExportQuoting
	Body              vampy.ScanIssueFilter
}

// ScanIssuesExportMeta Get CSV export metadata
func (c *Client) ScanIssuesExportMeta(ctx context.Context, opts ScanIssuesExportMetaOptions) (vampy.CSVExportMeta, error) {
	path := "/scan_issues/export/meta/"

	var out vampy.CSVExportMeta
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssuesExportMetaOptions struct {
	Body vampy.ScanIssueFilter
}

// ScanIssueSearchDuplicates Search ScanIssue's duplicates
func (c *Client) ScanIssueSearchDuplicates(ctx context.Context, opts ScanIssueSearchDuplicatesOptions) (vampy.ScanIssueDuplicatesSearchResult, error) {
	path := "/scan_issues/search_duplicates/"

	var out vampy.ScanIssueDuplicatesSearchResult
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueSearchDuplicatesOptions struct {
	Body vampy.ScanIssueCreate
}

// ScanIssuesDelete ScanIssue: group deletion
func (c *Client) ScanIssuesDelete(ctx context.Context, opts ScanIssuesDeleteOptions) error {
	path := "/scan_issues/delete/"

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ScanIssuesDeleteOptions struct {
	Body vampy.ScanIssueDelete
}

// ScanIssueSchemasGet Get scan_issue schemas by each scan_issue_type
func (c *Client) ScanIssueSchemasGet(ctx context.Context) ([]vampy.ScanIssues, error) {
	path := "/scan_issues/schemas/"

	var out []vampy.ScanIssues
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// ScanIssueDetailsGet Get information about ScanIssue
func (c *Client) ScanIssueDetailsGet(ctx context.Context, opts ScanIssueDetailsGetOptions) (vampy.ScanIssueDetail, error) {
	path := fmt.Sprintf("/scan_issues/%s/", url.PathEscape(opts.ScanIssueID.String()))

	var out vampy.ScanIssueDetail
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueDetailsGetOptions struct {
	ScanIssueID uuid.UUID
}

// ScanIssueStatusChangesGet Get history of status changes
func (c *Client) ScanIssueStatusChangesGet(ctx context.Context, opts ScanIssueStatusChangesGetOptions) ([]vampy.ScanIssueStatusChangeEntry, error) {
	path := fmt.Sprintf("/scan_issues/%s/status_changes/", url.PathEscape(opts.ScanIssueID.String()))

	var out []vampy.ScanIssueStatusChangeEntry
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueStatusChangesGetOptions struct {
	ScanIssueID uuid.UUID
}

// ScanIssueChangesGet Get history of scan issue changes (status, severity, FP score)
func (c *Client) ScanIssueChangesGet(ctx context.Context, opts ScanIssueChangesGetOptions) ([]vampy.ScanIssueChangeEntry, error) {
	path := fmt.Sprintf("/scan_issues/%s/changes/", url.PathEscape(opts.ScanIssueID.String()))
	q := url.Values{}

	for _, v := range opts.ChangeTypes {
		q.Add("changeTypes", string(v))
	}

	var out []vampy.ScanIssueChangeEntry
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueChangesGetOptions struct {
	ScanIssueID uuid.UUID
	ChangeTypes []vampy.IssueChangeType
}

// ScanIssueAIRequestPost Request AI suggestion
func (c *Client) ScanIssueAIRequestPost(ctx context.Context, opts ScanIssueAIRequestPostOptions) (vampy.ScanIssueAIRequestResponse, error) {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/", url.PathEscape(opts.ScanIssueID.String()))

	var out vampy.ScanIssueAIRequestResponse
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueAIRequestPostOptions struct {
	ScanIssueID uuid.UUID
	Body        vampy.ScanIssuesAIRequest
}

// ScanIssueAIRequestCancel Cancel AI requests of specific type
func (c *Client) ScanIssueAIRequestCancel(ctx context.Context, opts ScanIssueAIRequestCancelOptions) error {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/cancel/", url.PathEscape(opts.ScanIssueID.String()))

	return c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, nil)
}

type ScanIssueAIRequestCancelOptions struct {
	ScanIssueID uuid.UUID
	Body        vampy.ScanIssueAIRequestCancel
}

// ScanIssueAIRequestDetailsGet Get AI request details
func (c *Client) ScanIssueAIRequestDetailsGet(ctx context.Context, opts ScanIssueAIRequestDetailsGetOptions) (vampy.ScanIssueAIRequestResponse, error) {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/%s/", url.PathEscape(opts.ScanIssueID.String()), url.PathEscape(opts.AIRequestID.String()))

	var out vampy.ScanIssueAIRequestResponse
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueAIRequestDetailsGetOptions struct {
	ScanIssueID uuid.UUID
	AIRequestID uuid.UUID
}

// ScanIssueAIRequestApply Apply AI request
func (c *Client) ScanIssueAIRequestApply(ctx context.Context, opts ScanIssueAIRequestApplyOptions) (vampy.ScanIssueAIRequestResponse, error) {
	path := fmt.Sprintf("/scan_issues/%s/ai_requests/%s/apply/", url.PathEscape(opts.ScanIssueID.String()), url.PathEscape(opts.AIRequestID.String()))

	var out vampy.ScanIssueAIRequestResponse
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssueAIRequestApplyOptions struct {
	ScanIssueID uuid.UUID
	AIRequestID uuid.UUID
	Body        vampy.ScanIssueAIRequestApply
}

// ScanIssuesUpdate Update batch of scan issues
func (c *Client) ScanIssuesUpdate(ctx context.Context, opts ScanIssuesUpdateOptions) (any, error) {
	path := "/scan_issues/update/"

	var out any
	if err := c.doJSON(ctx, http.MethodPost, path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScanIssuesUpdateOptions struct {
	Body vampy.ScanIssuesUpdate
}
