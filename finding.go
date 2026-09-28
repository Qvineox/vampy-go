package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// FindingsGet Paginated entrypoint for getting Findings
func (c *Client) FindingsGet(ctx context.Context, offset *int, limit *int, spaceID *string, engagementID *string, repositoryID *string, repositoryBranchID *string, scanIssueID *string, statuses []string, parsers []string) (vampy.InlineResponse20016, error) {
	path := "/findings/"
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if spaceID != nil {
		q.Set("spaceID", *spaceID)
	}
	if engagementID != nil {
		q.Set("engagementID", *engagementID)
	}
	if repositoryID != nil {
		q.Set("repositoryID", *repositoryID)
	}
	if repositoryBranchID != nil {
		q.Set("repositoryBranchID", *repositoryBranchID)
	}
	if scanIssueID != nil {
		q.Set("scanIssueID", *scanIssueID)
	}
	for _, v := range statuses {
		q.Add("statuses", v)
	}
	for _, v := range parsers {
		q.Add("parsers", v)
	}
	var out vampy.InlineResponse20016
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FindingCardGet Get information about Finding
func (c *Client) FindingCardGet(ctx context.Context, findingID string) (vampy.FindingCard, error) {
	path := fmt.Sprintf("/findings/%s/", url.PathEscape(findingID))
	var out vampy.FindingCard
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
