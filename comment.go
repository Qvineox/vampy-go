package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// CommentListGet Paginated list of comments
func (c *Client) CommentListGet(ctx context.Context, opts CommentListGetOptions) (vampy.Page[vampy.CommentDetails], error) {
	path := fmt.Sprintf("/scan_issues/%s/comments/", url.PathEscape(opts.ScanIssueID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.UserID != nil {
		q.Set("userID", opts.UserID.String())
	}

	if opts.ExportRecordID != nil {
		q.Set("exportRecordID", *opts.ExportRecordID)
	}

	if opts.CommentType != nil {
		q.Set("commentType", string(*opts.CommentType))
	}

	if opts.IncludeAssociatedIssues != nil {
		q.Set("includeAssociatedIssues", strconv.FormatBool(*opts.IncludeAssociatedIssues))
	}

	var out vampy.Page[vampy.CommentDetails]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CommentListGetOptions struct {
	ScanIssueID             uuid.UUID
	UserID                  *uuid.UUID
	ExportRecordID          *string
	CommentType             *vampy.CommentType
	IncludeAssociatedIssues *bool

	vampy.Paging
}

// CommentPost Create comment
func (c *Client) CommentPost(ctx context.Context, opts CommentPostOptions) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/scan_issues/%s/comments/", url.PathEscape(opts.ScanIssueID.String()))

	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CommentPostOptions struct {
	ScanIssueID uuid.UUID
	Body        vampy.Comment
}

// CommentDetailsGet Get information about comment
func (c *Client) CommentDetailsGet(ctx context.Context, opts CommentDetailsGetOptions) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/comments/%s/", url.PathEscape(opts.CommentID.String()))

	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CommentDetailsGetOptions struct {
	CommentID uuid.UUID
}

// CommentDelete Delete comment
func (c *Client) CommentDelete(ctx context.Context, opts CommentDeleteOptions) error {
	path := fmt.Sprintf("/comments/%s/", url.PathEscape(opts.CommentID.String()))
	q := url.Values{}

	if opts.Force != nil {
		q.Set("force", strconv.FormatBool(*opts.Force))
	}

	return c.doJSON(ctx, "DELETE", path, q, nil, nil)
}

type CommentDeleteOptions struct {
	CommentID uuid.UUID
	Force     *bool
}

// CommentPatch Patch comment
func (c *Client) CommentPatch(ctx context.Context, opts CommentPatchOptions) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/comments/%s/", url.PathEscape(opts.CommentID.String()))

	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CommentPatchOptions struct {
	CommentID uuid.UUID
	Body      vampy.Comment
}

// CommentSync Sync comment with task tracker
func (c *Client) CommentSync(ctx context.Context, opts CommentSyncOptions) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/comments/%s/sync/", url.PathEscape(opts.CommentID.String()))

	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "POST", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CommentSyncOptions struct {
	CommentID uuid.UUID
}
