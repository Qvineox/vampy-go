package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// CommentListGet Paginated list of comments
func (c *Client) CommentListGet(ctx context.Context, scanIssueID string, offset *int, limit *int, userID *string, exportRecordID *string, commentType *vampy.CommentType, includeAssociatedIssues *bool) (vampy.Page[vampy.CommentDetails], error) {
	path := fmt.Sprintf("/scan_issues/%s/comments/", url.PathEscape(scanIssueID))
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if userID != nil {
		q.Set("userID", *userID)
	}
	if exportRecordID != nil {
		q.Set("exportRecordID", *exportRecordID)
	}
	if commentType != nil {
		q.Set("commentType", string(*commentType))
	}
	if includeAssociatedIssues != nil {
		q.Set("includeAssociatedIssues", strconv.FormatBool(*includeAssociatedIssues))
	}
	var out vampy.Page[vampy.CommentDetails]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
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
