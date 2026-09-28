package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// GlobalDashboardGet Get global dashboard
func (c *Client) GlobalDashboardGet(ctx context.Context, spaceID string, scannerType *vampy.ScannerType) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/spaces/%s/dashboard/", url.PathEscape(spaceID))
	q := url.Values{}
	if scannerType != nil {
		q.Set("scannerType", string(*scannerType))
	}
	var out vampy.Dashboard
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GlobalDashboardCardGet Get global dashboard card
func (c *Client) GlobalDashboardCardGet(ctx context.Context, spaceID string, cardID string, period *vampy.DashboardCardAggregationPeriod, severity *vampy.Severity, scannerType *vampy.ScannerType) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/spaces/%s/dashboard/cards/%s/", url.PathEscape(spaceID), url.PathEscape(cardID))
	q := url.Values{}
	if period != nil {
		q.Set("period", string(*period))
	}
	if severity != nil {
		q.Set("severity", string(*severity))
	}
	if scannerType != nil {
		q.Set("scannerType", string(*scannerType))
	}
	var out vampy.Dashboard
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ProductDashboardGet Get product dashboard
func (c *Client) ProductDashboardGet(ctx context.Context, productID string, scannerType *vampy.ScannerType) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/products/%s/dashboard/", url.PathEscape(productID))
	q := url.Values{}
	if scannerType != nil {
		q.Set("scannerType", string(*scannerType))
	}
	var out vampy.Dashboard
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ProductDashboardCardGet Get product dashboard card
func (c *Client) ProductDashboardCardGet(ctx context.Context, productID string, cardID string, period *vampy.DashboardCardAggregationPeriod, severity *vampy.Severity, scannerType *vampy.ScannerType) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/products/%s/dashboard/cards/%s/", url.PathEscape(productID), url.PathEscape(cardID))
	q := url.Values{}
	if period != nil {
		q.Set("period", string(*period))
	}
	if severity != nil {
		q.Set("severity", string(*severity))
	}
	if scannerType != nil {
		q.Set("scannerType", string(*scannerType))
	}
	var out vampy.Dashboard
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryDashboardGet Get repository dashboard
func (c *Client) RepositoryDashboardGet(ctx context.Context, repositoryID string, scannerType *vampy.ScannerType, repositoryBranchID *string, allBranches *bool) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/repositories/%s/dashboard/", url.PathEscape(repositoryID))
	q := url.Values{}
	if scannerType != nil {
		q.Set("scannerType", string(*scannerType))
	}
	if repositoryBranchID != nil {
		q.Set("repositoryBranchID", *repositoryBranchID)
	}
	if allBranches != nil {
		q.Set("allBranches", strconv.FormatBool(*allBranches))
	}
	var out vampy.Dashboard
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// RepositoryDashboardCardGet Get repository dashboard card
func (c *Client) RepositoryDashboardCardGet(ctx context.Context, repositoryID string, cardID string, period *vampy.DashboardCardAggregationPeriod, scannerType *vampy.ScannerType, repositoryBranchID *string, allBranches *bool, severity *vampy.Severity) (vampy.Dashboard, error) {
	path := fmt.Sprintf("/repositories/%s/dashboard/cards/%s/", url.PathEscape(repositoryID), url.PathEscape(cardID))
	q := url.Values{}
	if period != nil {
		q.Set("period", string(*period))
	}
	if scannerType != nil {
		q.Set("scannerType", string(*scannerType))
	}
	if repositoryBranchID != nil {
		q.Set("repositoryBranchID", *repositoryBranchID)
	}
	if allBranches != nil {
		q.Set("allBranches", strconv.FormatBool(*allBranches))
	}
	if severity != nil {
		q.Set("severity", string(*severity))
	}
	var out vampy.Dashboard
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
