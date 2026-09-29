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

// GlobalDashboardGet Get global dashboard
func (c *Client) GlobalDashboardGet(ctx context.Context, opts GlobalDashboardGetOptions) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/spaces/%s/dashboard/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	if opts.ScannerType != nil {
		q.Set("scannerType", string(*opts.ScannerType))
	}

	var out vampy.Dashboard
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GlobalDashboardGetOptions struct {
	ScannerType *vampy.ScannerType
}

// GlobalDashboardCardGet Get global dashboard card
func (c *Client) GlobalDashboardCardGet(ctx context.Context, opts GlobalDashboardCardGetOptions) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/spaces/%s/dashboard/cards/%s/", url.PathEscape(c.spaceID.String()), url.PathEscape(opts.CardID.String()))
	q := url.Values{}

	if opts.Period != nil {
		q.Set("period", string(*opts.Period))
	}

	if opts.Severity != nil {
		q.Set("severity", string(*opts.Severity))
	}

	if opts.ScannerType != nil {
		q.Set("scannerType", string(*opts.ScannerType))
	}

	var out vampy.Dashboard
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GlobalDashboardCardGetOptions struct {
	CardID      uuid.UUID
	Period      *vampy.DashboardCardAggregationPeriod
	Severity    *vampy.Severity
	ScannerType *vampy.ScannerType
}

// ProductDashboardGet Get product dashboard
func (c *Client) ProductDashboardGet(ctx context.Context, opts ProductDashboardGetOptions) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/products/%s/dashboard/", url.PathEscape(opts.ProductID.String()))
	q := url.Values{}

	if opts.ScannerType != nil {
		q.Set("scannerType", string(*opts.ScannerType))
	}

	var out vampy.Dashboard
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ProductDashboardGetOptions struct {
	ProductID   uuid.UUID
	ScannerType *vampy.ScannerType
}

// ProductDashboardCardGet Get product dashboard card
func (c *Client) ProductDashboardCardGet(ctx context.Context, opts ProductDashboardCardGetOptions) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/products/%s/dashboard/cards/%s/", url.PathEscape(opts.ProductID.String()), url.PathEscape(opts.CardID.String()))
	q := url.Values{}

	if opts.Period != nil {
		q.Set("period", string(*opts.Period))
	}

	if opts.Severity != nil {
		q.Set("severity", string(*opts.Severity))
	}

	if opts.ScannerType != nil {
		q.Set("scannerType", string(*opts.ScannerType))
	}

	var out vampy.Dashboard
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ProductDashboardCardGetOptions struct {
	ProductID   uuid.UUID
	CardID      uuid.UUID
	Period      *vampy.DashboardCardAggregationPeriod
	Severity    *vampy.Severity
	ScannerType *vampy.ScannerType
}

// RepositoryDashboardGet Get repository dashboard
func (c *Client) RepositoryDashboardGet(ctx context.Context, opts RepositoryDashboardGetOptions) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/repositories/%s/dashboard/", url.PathEscape(opts.RepositoryID.String()))
	q := url.Values{}

	if opts.ScannerType != nil {
		q.Set("scannerType", string(*opts.ScannerType))
	}

	if opts.RepositoryBranchID != nil {
		q.Set("repositoryBranchID", opts.RepositoryBranchID.String())
	}

	if opts.AllBranches != nil {
		q.Set("allBranches", strconv.FormatBool(*opts.AllBranches))
	}

	var out vampy.Dashboard
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryDashboardGetOptions struct {
	RepositoryID       uuid.UUID
	ScannerType        *vampy.ScannerType
	RepositoryBranchID *uuid.UUID
	AllBranches        *bool
}

// RepositoryDashboardCardGet Get repository dashboard card
func (c *Client) RepositoryDashboardCardGet(ctx context.Context, opts RepositoryDashboardCardGetOptions) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/repositories/%s/dashboard/cards/%s/", url.PathEscape(opts.RepositoryID.String()), url.PathEscape(opts.CardID.String()))
	q := url.Values{}

	if opts.Period != nil {
		q.Set("period", string(*opts.Period))
	}

	if opts.ScannerType != nil {
		q.Set("scannerType", string(*opts.ScannerType))
	}

	if opts.RepositoryBranchID != nil {
		q.Set("repositoryBranchID", opts.RepositoryBranchID.String())
	}

	if opts.AllBranches != nil {
		q.Set("allBranches", strconv.FormatBool(*opts.AllBranches))
	}

	if opts.Severity != nil {
		q.Set("severity", string(*opts.Severity))
	}

	var out vampy.Dashboard
	if err := c.doJSON(ctx, http.MethodGet, path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type RepositoryDashboardCardGetOptions struct {
	RepositoryID       uuid.UUID
	CardID             uuid.UUID
	Period             *vampy.DashboardCardAggregationPeriod
	ScannerType        *vampy.ScannerType
	RepositoryBranchID *uuid.UUID
	AllBranches        *bool
	Severity           *vampy.Severity
}
