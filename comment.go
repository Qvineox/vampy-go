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
		q.Set("userID", *opts.UserID)
	}

	if opts.ExportRecordID != nil {
		q.Set("exportRecordID", opts.ExportRecordID.String())
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
	ExportRecordID          *uuid.UUID
	CommentType             *vampy.CommentType
	IncludeAssociatedIssues *bool
	UserID                  *string

	vampy.Paging
}

// CommentPost Create comment
func (c *Client) CommentPost(ctx context.Context, scanIssueID string, body vampy.Comment) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/scan_issues/%s/comments/", url.PathEscape(scanIssueID))
	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CommentDetailsGet Get information about comment
func (c *Client) CommentDetailsGet(ctx context.Context, commentID string) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/comments/%s/", url.PathEscape(commentID))
	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CommentDelete Delete comment
func (c *Client) CommentDelete(ctx context.Context, commentID string, force *bool) error {
	path := fmt.Sprintf("/comments/%s/", url.PathEscape(commentID))
	q := url.Values{}
	if force != nil {
		q.Set("force", strconv.FormatBool(*force))
	}
	return c.doJSON(ctx, "DELETE", path, q, nil, nil)
}

// CommentPatch Patch comment
func (c *Client) CommentPatch(ctx context.Context, commentID string, body vampy.Comment) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/comments/%s/", url.PathEscape(commentID))
	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CommentSync Sync comment with task tracker
func (c *Client) CommentSync(ctx context.Context, commentID string) (vampy.CommentDetails, error) {
	path := fmt.Sprintf("/comments/%s/sync/", url.PathEscape(commentID))
	var out vampy.CommentDetails
	if err := c.doJSON(ctx, "POST", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
