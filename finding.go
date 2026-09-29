package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// FindingsGet Paginated entrypoint for getting Findings
func (c *Client) FindingsGet(ctx context.Context, opts FindingsGetOptions) (vampy.Page[vampy.Finding], error) {
	path := "/findings/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.SpaceID != nil {
		q.Set("spaceID", opts.SpaceID.String())
	}

	if opts.EngagementID != nil {
		q.Set("engagementID", opts.EngagementID.String())
	}

	if opts.RepositoryID != nil {
		q.Set("repositoryID", opts.RepositoryID.String())
	}

	if opts.RepositoryBranchID != nil {
		q.Set("repositoryBranchID", opts.RepositoryBranchID.String())
	}

	if opts.ScanIssueID != nil {
		q.Set("scanIssueID", opts.ScanIssueID.String())
	}

	for _, v := range opts.Statuses {
		q.Add("statuses", string(v))
	}

	for _, v := range opts.Parsers {
		q.Add("parsers", v)
	}

	var out vampy.Page[vampy.Finding]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FindingsGetOptions struct {
	SpaceID            *uuid.UUID
	EngagementID       *uuid.UUID
	RepositoryID       *uuid.UUID
	RepositoryBranchID *uuid.UUID
	ScanIssueID        *uuid.UUID
	Statuses           []vampy.FindingStatus
	Parsers            []string

	vampy.Paging
}

// FindingCardGet Get information about Finding
func (c *Client) FindingCardGet(ctx context.Context, opts FindingCardGetOptions) (vampy.FindingCard, error) {
	path := fmt.Sprintf("/findings/%s/", url.PathEscape(opts.FindingID.String()))

	var out vampy.FindingCard
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FindingCardGetOptions struct {
	FindingID uuid.UUID
}
