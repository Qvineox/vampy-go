package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// SpaceIntegrationsListGet Get list spaces integrations with external services
func (c *Client) SpaceIntegrationsListGet(ctx context.Context, spaceID string, service *vampy.IntegrationServiceType, role *vampy.IntegrationRole) ([]vampy.SpaceIntegrationWithMeta, error) {
	path := fmt.Sprintf("/spaces/%s/integrations/", url.PathEscape(spaceID))
	q := url.Values{}
	if service != nil {
		q.Set("service", string(*service))
	}
	if role != nil {
		q.Set("role", string(*role))
	}
	var out []vampy.SpaceIntegrationWithMeta
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SpaceIntegrationsCreate Add new space's  integration
func (c *Client) SpaceIntegrationsCreate(ctx context.Context, spaceID string, body vampy.SpaceIntegration) (vampy.SpaceIntegration, error) {
	path := fmt.Sprintf("/spaces/%s/integrations/", url.PathEscape(spaceID))
	var out vampy.SpaceIntegration
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetBitbucketApplicationInfo Get Bitbucket application version/authentication method by hostname
func (c *Client) GetBitbucketApplicationInfo(ctx context.Context, body vampy.BitbucketApplicationInfoRequest) (vampy.BitbucketApplicationInfoResponse, error) {
	path := "/integrations/bitbucket/info/"
	var out vampy.BitbucketApplicationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationsVCSGet Get VCS integration
func (c *Client) IntegrationsVCSGet(ctx context.Context, vcsIntegration vampy.VCSIntegration, integrationID string) (vampy.IntegrationVCS, error) {
	path := fmt.Sprintf("/integrations/%s/%s/", url.PathEscape(string(vcsIntegration)), url.PathEscape(integrationID))
	var out vampy.IntegrationVCS
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationsVCSDelete Delete VCS integration
func (c *Client) IntegrationsVCSDelete(ctx context.Context, vcsIntegration vampy.VCSIntegration, integrationID string) error {
	path := fmt.Sprintf("/integrations/%s/%s/", url.PathEscape(string(vcsIntegration)), url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// IntegrationsVCSPatch Update VCS integration
func (c *Client) IntegrationsVCSPatch(ctx context.Context, vcsIntegration vampy.VCSIntegration, integrationID string, body vampy.IntegrationVCS) (vampy.IntegrationVCS, error) {
	path := fmt.Sprintf("/integrations/%s/%s/", url.PathEscape(string(vcsIntegration)), url.PathEscape(integrationID))
	var out vampy.IntegrationVCS
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationsVCSRepositoriesGet Get VCS integration's repositories
func (c *Client) IntegrationsVCSRepositoriesGet(ctx context.Context, vcsIntegration vampy.VCSIntegration, integrationID string, offset *int, limit *int, query *string) (vampy.Page[vampy.ExternalRepository], error) {
	path := fmt.Sprintf("/integrations/%s/%s/repositories/", url.PathEscape(string(vcsIntegration)), url.PathEscape(integrationID))
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if query != nil {
		q.Set("query", *query)
	}
	var out vampy.Page[vampy.ExternalRepository]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationsVCSInfo Get VCS integration metadata (authentication, unsupported features)
func (c *Client) IntegrationsVCSInfo(ctx context.Context, vcsIntegration vampy.VCSIntegration, body vampy.VCSIntegrationInfoRequest) (vampy.VCSIntegrationInfoResponse, error) {
	path := fmt.Sprintf("/integrations/vcs/%s/info/", url.PathEscape(string(vcsIntegration)))
	var out vampy.VCSIntegrationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationsVCSMeta Get VCS integration metadata using stored credentials with optional overrides
func (c *Client) IntegrationsVCSMeta(ctx context.Context, vcsIntegration vampy.VCSIntegration, integrationID string, body vampy.VCSIntegrationInfoRequest) (vampy.VCSIntegrationInfoResponse, error) {
	path := fmt.Sprintf("/integrations/vcs/%s/%s/meta/", url.PathEscape(string(vcsIntegration)), url.PathEscape(integrationID))
	var out vampy.VCSIntegrationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// GetJiraApplicationVersion Get Jira application version/authentication method by hostname
func (c *Client) GetJiraApplicationVersion(ctx context.Context, body vampy.JiraApplicationInfoRequest) (vampy.JiraApplicationInfoResponse, error) {
	path := "/integrations/jira/info/"
	var out vampy.JiraApplicationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerGetInfo Get Task tracker service info method
func (c *Client) TaskTrackerGetInfo(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, body vampy.TaskTrackerInfoRequest) (vampy.TaskTrackerInfoResponse, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/info/", url.PathEscape(string(taskTrackerIntegration)))
	var out vampy.TaskTrackerInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// JiraIntegrationsMeta Get Jira integration instance meta information about projects and issues
func (c *Client) JiraIntegrationsMeta(ctx context.Context, integrationID string, projectKey *string, initialQuery *bool) (vampy.JiraMetaInfoResponse, error) {
	path := fmt.Sprintf("/integrations/jira/%s/meta/", url.PathEscape(integrationID))
	q := url.Values{}
	if projectKey != nil {
		q.Set("projectKey", *projectKey)
	}
	if initialQuery != nil {
		q.Set("initialQuery", strconv.FormatBool(*initialQuery))
	}
	var out vampy.JiraMetaInfoResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerIntegrationGetMeta Get Task tracker integration instance meta information about projects and issues
func (c *Client) TaskTrackerIntegrationGetMeta(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, integrationID string, project *string, issueType *string, initialQuery *bool) (vampy.TaskTrackerMetaInfoResponse, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/meta/", url.PathEscape(string(taskTrackerIntegration)), url.PathEscape(integrationID))
	q := url.Values{}
	if project != nil {
		q.Set("project", *project)
	}
	if issueType != nil {
		q.Set("issueType", *issueType)
	}
	if initialQuery != nil {
		q.Set("initialQuery", strconv.FormatBool(*initialQuery))
	}
	var out vampy.TaskTrackerMetaInfoResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// JiraIntegrationGet Get Jira integration
func (c *Client) JiraIntegrationGet(ctx context.Context, integrationID string) (vampy.JiraIntegration, error) {
	path := fmt.Sprintf("/integrations/jira/%s/", url.PathEscape(integrationID))
	var out vampy.JiraIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// JiraIntegrationDelete Remove existing Jira integration.
func (c *Client) JiraIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/jira/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// JiraIntegrationPatch Set configuration values for existing Jira integration.
func (c *Client) JiraIntegrationPatch(ctx context.Context, integrationID string, body vampy.JiraIntegration) (vampy.JiraIntegration, error) {
	path := fmt.Sprintf("/integrations/jira/%s/", url.PathEscape(integrationID))
	var out vampy.JiraIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerIntegrationGet Get Task tracker integration
func (c *Client) TaskTrackerIntegrationGet(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, integrationID string) (vampy.TaskTrackerIntegration, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/", url.PathEscape(string(taskTrackerIntegration)), url.PathEscape(integrationID))
	var out vampy.TaskTrackerIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerIntegrationDelete Remove existing Task tracker integration.
func (c *Client) TaskTrackerIntegrationDelete(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, integrationID string) error {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/", url.PathEscape(string(taskTrackerIntegration)), url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// TaskTrackerIntegrationPatch Set configuration values for existing Task tracker integration.
func (c *Client) TaskTrackerIntegrationPatch(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, integrationID string, body vampy.TaskTrackerIntegration) (vampy.TaskTrackerIntegration, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/", url.PathEscape(string(taskTrackerIntegration)), url.PathEscape(integrationID))
	var out vampy.TaskTrackerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SearchItemsJira Preview issues that are about to be sent to jira. Check if we have required default fields
func (c *Client) SearchItemsJira(ctx context.Context, integrationID string, field string, query string, projectKey *string) ([]vampy.JiraSearchResponse, error) {
	path := fmt.Sprintf("/integrations/jira/%s/search/", url.PathEscape(integrationID))
	q := url.Values{}
	q.Set("field", field)
	q.Set("query", query)
	if projectKey != nil {
		q.Set("projectKey", *projectKey)
	}
	var out []vampy.JiraSearchResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerSearchItems Search items in Task tracker
func (c *Client) TaskTrackerSearchItems(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, integrationID string, field string, query string, project *string) ([]vampy.TaskTrackerSearchResponse, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/search/", url.PathEscape(string(taskTrackerIntegration)), url.PathEscape(integrationID))
	q := url.Values{}
	q.Set("field", field)
	q.Set("query", query)
	if project != nil {
		q.Set("project", *project)
	}
	var out []vampy.TaskTrackerSearchResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SearchIssuesJira Search Jira's issues
func (c *Client) SearchIssuesJira(ctx context.Context, integrationID string, query string) ([]vampy.JiraSearchIssuesResponse, error) {
	path := fmt.Sprintf("/integrations/jira/%s/search/issues/", url.PathEscape(integrationID))
	q := url.Values{}
	q.Set("query", query)
	var out []vampy.JiraSearchIssuesResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerSearchIssues Search Task tracker issues
func (c *Client) TaskTrackerSearchIssues(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, integrationID string, query string) ([]vampy.TaskTrackerIssueExport, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/search/issues/", url.PathEscape(string(taskTrackerIntegration)), url.PathEscape(integrationID))
	q := url.Values{}
	q.Set("query", query)
	var out []vampy.TaskTrackerIssueExport
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// KtrackerIntegrationGet Get KTracker integration
func (c *Client) KtrackerIntegrationGet(ctx context.Context, integrationID string) (vampy.IntegrationKTracker, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationKTracker
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// KtrackerIntegrationDelete Delete KTracker integration
func (c *Client) KtrackerIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/ktracker/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// KtrackerIntegrationPatch Update KTracker integration
func (c *Client) KtrackerIntegrationPatch(ctx context.Context, integrationID string, body vampy.IntegrationKTracker) (vampy.IntegrationKTracker, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationKTracker
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// KtrackerIntegrationGetSpaces KTracker integration get Spaces
func (c *Client) KtrackerIntegrationGetSpaces(ctx context.Context, integrationID string, showArchived *bool, limit *int) ([]vampy.IntegrationKTrackerSpace, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/spaces/", url.PathEscape(integrationID))
	q := url.Values{}
	if showArchived != nil {
		q.Set("showArchived", strconv.FormatBool(*showArchived))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	var out []vampy.IntegrationKTrackerSpace
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// KtrackerIntegrationGetCardTypes KTracker integration get Card types
func (c *Client) KtrackerIntegrationGetCardTypes(ctx context.Context, integrationID string, spaceID *int) ([]vampy.IntegrationKTrackerCardType, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/card_types/", url.PathEscape(integrationID))
	q := url.Values{}
	if spaceID != nil {
		q.Set("spaceID", strconv.Itoa(*spaceID))
	}
	var out []vampy.IntegrationKTrackerCardType
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScannerIntegrationGet Get scanner integration
func (c *Client) ScannerIntegrationGet(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, integrationID string) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/", url.PathEscape(string(scannerIntegration)), url.PathEscape(integrationID))
	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScannerIntegrationDelete Delete scanner integration
func (c *Client) ScannerIntegrationDelete(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, integrationID string) error {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/", url.PathEscape(string(scannerIntegration)), url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// ScannerIntegrationPatch Update scanner integration
func (c *Client) ScannerIntegrationPatch(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, integrationID string, body vampy.ScannerIntegrationPatch) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/", url.PathEscape(string(scannerIntegration)), url.PathEscape(integrationID))
	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScannerIntegrationMeta Get scanner integration instance meta information
func (c *Client) ScannerIntegrationMeta(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, integrationID string) (vampy.ScannerIntegrationMetaResponse, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/meta/", url.PathEscape(string(scannerIntegration)), url.PathEscape(integrationID))
	var out vampy.ScannerIntegrationMetaResponse
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScannerIntegrationInfo Get scanner integration info
func (c *Client) ScannerIntegrationInfo(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, body vampy.ScannerIntegrationInfoRequest) (vampy.ScannerIntegrationInfoResponse, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/info/", url.PathEscape(string(scannerIntegration)))
	var out vampy.ScannerIntegrationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScannerIntegrationFetchCLI Fetch CLI from scanner integration API
func (c *Client) ScannerIntegrationFetchCLI(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, integrationID string) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/fetch_cli/", url.PathEscape(string(scannerIntegration)), url.PathEscape(integrationID))
	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScannerIntegrationUploadCLI Upload CLI file
func (c *Client) ScannerIntegrationUploadCLI(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, integrationID string, body vampy.IntegrationIDUploadCLIBody) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/upload_cli/", url.PathEscape(string(scannerIntegration)), url.PathEscape(integrationID))
	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// ScannerIntegrationTestConnection Test scanner integration connection
func (c *Client) ScannerIntegrationTestConnection(ctx context.Context, scannerIntegration vampy.ScannerIntegrationType, integrationID string, body vampy.ScannerIntegrationTestConnection) error {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/test_connection/", url.PathEscape(string(scannerIntegration)), url.PathEscape(integrationID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// PtaiIntegrationGet Get PTAI integration
func (c *Client) PtaiIntegrationGet(ctx context.Context, integrationID string) (vampy.IntegrationPTAI, error) {
	path := fmt.Sprintf("/integrations/ptai/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationPTAI
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// PtaiIntegrationDelete Delete PTAI integration
func (c *Client) PtaiIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/ptai/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// PtaiIntegrationPatch Update PTAI integration
func (c *Client) PtaiIntegrationPatch(ctx context.Context, integrationID string, body vampy.IntegrationPTAI) (vampy.IntegrationPTAI, error) {
	path := fmt.Sprintf("/integrations/ptai/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationPTAI
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// PtaiIntegrationTestConnection Test PTAI connection
func (c *Client) PtaiIntegrationTestConnection(ctx context.Context, integrationID string, body vampy.ScannerIntegrationTestConnection) error {
	path := fmt.Sprintf("/integrations/ptai/%s/test_connection/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// TaskTrackerExportMeta Get metadata for scan issues to be exported to task tracker.
func (c *Client) TaskTrackerExportMeta(ctx context.Context, integrationService vampy.TaskTrackerServiceType, integrationID string, body vampy.TaskTrackerExportMetaRequest) (vampy.ScanIssueForExportResponse, error) {
	path := fmt.Sprintf("/integrations/%s/%s/export/meta/", url.PathEscape(string(integrationService)), url.PathEscape(integrationID))
	var out vampy.ScanIssueForExportResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerExportAsync Export scan issues to task tracker asynchronously
func (c *Client) TaskTrackerExportAsync(ctx context.Context, integrationService vampy.TaskTrackerServiceType, integrationID string, taskState *vampy.BackgroundTaskStatus, failedScanIssueID *string, body vampy.TaskTrackerAsyncExport) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/integrations/%s/%s/export/async/", url.PathEscape(string(integrationService)), url.PathEscape(integrationID))
	q := url.Values{}
	if taskState != nil {
		q.Set("taskState", string(*taskState))
	}
	if failedScanIssueID != nil {
		q.Set("failedScanIssueID", *failedScanIssueID)
	}
	var out vampy.BackgroundTask
	if err := c.doJSON(ctx, "POST", path, q, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// TaskTrackerLinkIssue Link scan issue to an existing external task tracker issue
func (c *Client) TaskTrackerLinkIssue(ctx context.Context, taskTrackerIntegration vampy.TaskTrackerIntegrationType, integrationID string, body vampy.TaskTrackerLinkRequest) ([]vampy.IssueExport, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/link/", url.PathEscape(string(taskTrackerIntegration)), url.PathEscape(integrationID))
	var out []vampy.IssueExport
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CodeScoringIntegrationGet Get CodeScoring integration
func (c *Client) CodeScoringIntegrationGet(ctx context.Context, integrationID string) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/", url.PathEscape(integrationID))
	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CodeScoringIntegrationDelete Delete CodeScoring integration
func (c *Client) CodeScoringIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/code_scoring/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// CodeScoringIntegrationPatch Update CodeScoring integration
func (c *Client) CodeScoringIntegrationPatch(ctx context.Context, integrationID string, body vampy.ScannerIntegrationPatch) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/", url.PathEscape(integrationID))
	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CodeScoringIntegrationMeta Get CodeScoring integration instance meta information
func (c *Client) CodeScoringIntegrationMeta(ctx context.Context, integrationID string) (vampy.ScannerIntegrationMetaResponse, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/meta/", url.PathEscape(integrationID))
	var out vampy.ScannerIntegrationMetaResponse
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CodeScoringIntegrationFetchCLI Fetch CLI from CodeScoring API
func (c *Client) CodeScoringIntegrationFetchCLI(ctx context.Context, integrationID string) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/fetch_cli/", url.PathEscape(integrationID))
	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// CodeScoringIntegrationUploadCLI Upload CLI file
func (c *Client) CodeScoringIntegrationUploadCLI(ctx context.Context, integrationID string, body vampy.IntegrationIDUploadCLIBody1) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/upload_cli/", url.PathEscape(integrationID))
	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AppScreenerIntegrationGet Get AppScreener integration
func (c *Client) AppScreenerIntegrationGet(ctx context.Context, integrationID string) (vampy.IntegrationAppScreener, error) {
	path := fmt.Sprintf("/integrations/app_screener/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationAppScreener
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AppScreenerIntegrationDelete Delete AppScreener integration
func (c *Client) AppScreenerIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/app_screener/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// AppScreenerIntegrationPatch Update AppScreener integration
func (c *Client) AppScreenerIntegrationPatch(ctx context.Context, integrationID string, body vampy.ScannerIntegrationPatch) (vampy.IntegrationAppScreener, error) {
	path := fmt.Sprintf("/integrations/app_screener/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationAppScreener
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AppScreenerIntegrationUploadCLI Upload CLI file
func (c *Client) AppScreenerIntegrationUploadCLI(ctx context.Context, integrationID string, body vampy.IntegrationIDUploadCLIBody2) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/app_screener/%s/upload_cli/", url.PathEscape(integrationID))
	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SastavIntegrationGet Get SASTAV integration
func (c *Client) SastavIntegrationGet(ctx context.Context, integrationID string) (vampy.IntegrationSastav, error) {
	path := fmt.Sprintf("/integrations/sastav/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationSastav
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SastavIntegrationDelete Delete SASTAV integration
func (c *Client) SastavIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/sastav/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// SastavIntegrationPatch Update SASTAV integration
func (c *Client) SastavIntegrationPatch(ctx context.Context, integrationID string, body vampy.ScannerIntegrationPatch) (vampy.IntegrationSastav, error) {
	path := fmt.Sprintf("/integrations/sastav/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationSastav
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DefectDojoIntegrationGet Get Defect Dojo integration
func (c *Client) DefectDojoIntegrationGet(ctx context.Context, integrationID string) (vampy.IntegrationDefectDojo, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationDefectDojo
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DefectDojoIntegrationDelete Delete Defect Dojo integration
func (c *Client) DefectDojoIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// DefectDojoIntegrationPatch Update Defect Dojo integration
func (c *Client) DefectDojoIntegrationPatch(ctx context.Context, integrationID string, body vampy.IntegrationDefectDojo) (vampy.IntegrationDefectDojo, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationDefectDojo
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DefectDojoIntegrationGetProducts Get list of products from Defect Dojo
func (c *Client) DefectDojoIntegrationGetProducts(ctx context.Context, integrationID string, limit *int, offset *int, search *string) (vampy.Page[vampy.DefectDojoProduct], error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/products/", url.PathEscape(integrationID))
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out vampy.Page[vampy.DefectDojoProduct]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DefectDojoIntegrationGetStatuses Get DefectDojo statuses mapping
func (c *Client) DefectDojoIntegrationGetStatuses(ctx context.Context, integrationID string) ([]vampy.DefectDojoStatusMapping, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/statuses/", url.PathEscape(integrationID))
	var out []vampy.DefectDojoStatusMapping
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// DefectDojoIntegrationImport Import data from Defect Dojo: findings, products, repos
func (c *Client) DefectDojoIntegrationImport(ctx context.Context, integrationID string, body vampy.DefectDojoImportRequest) (vampy.DefectDojoImportResponse, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/import/", url.PathEscape(integrationID))
	var out vampy.DefectDojoImportResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AIIntegrationGet Get AI integration
func (c *Client) AIIntegrationGet(ctx context.Context, integrationID string) (vampy.IntegrationAI, error) {
	path := fmt.Sprintf("/integrations/ai/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationAI
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AIIntegrationDelete Delete AI integration
func (c *Client) AIIntegrationDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/ai/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// AIIntegrationPatch Update AI integration
func (c *Client) AIIntegrationPatch(ctx context.Context, integrationID string, body vampy.IntegrationAI) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/ai/%s/", url.PathEscape(integrationID))
	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AIIntegrationFetchModels Update AI models
func (c *Client) AIIntegrationFetchModels(ctx context.Context, integrationID string) (vampy.IntegrationAI, error) {
	path := fmt.Sprintf("/integrations/ai/%s/fetch_models/", url.PathEscape(integrationID))
	var out vampy.IntegrationAI
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AIIntegrationAllModels List of all available AI models
func (c *Client) AIIntegrationAllModels(ctx context.Context) ([]vampy.IntegrationAIModel, error) {
	path := "/integrations/ai/models/"
	var out []vampy.IntegrationAIModel
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AIIntegrationAppliedModels List of historically applied AI models
func (c *Client) AIIntegrationAppliedModels(ctx context.Context, requestType *vampy.IntegrationAIRequestType) ([]vampy.IntegrationAIModel, error) {
	path := "/integrations/ai/models/applied/"
	q := url.Values{}
	if requestType != nil {
		q.Set("requestType", string(*requestType))
	}
	var out []vampy.IntegrationAIModel
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// AIIntegrationVendors List of all available AI vendors
func (c *Client) AIIntegrationVendors(ctx context.Context) ([]vampy.IntegrationAIVendor, error) {
	path := "/integrations/ai/vendors/"
	var out []vampy.IntegrationAIVendor
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationWebhookHandlerPost Webhook JIRA Handler: allows syncing external issue (already exported) with local one
func (c *Client) IntegrationWebhookHandlerPost(ctx context.Context, token string, body vampy.TokenWebhookBody) error {
	path := fmt.Sprintf("/integrations/servers/%s/webhook/", url.PathEscape(token))
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// IntegrationsArtifactDetailsGet Get Artifact integration
func (c *Client) IntegrationsArtifactDetailsGet(ctx context.Context, integrationID string) (vampy.IntegrationArtifact, error) {
	path := fmt.Sprintf("/integrations/artifact/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationArtifact
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationsArtifactDelete Delete Artifact integration
func (c *Client) IntegrationsArtifactDelete(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/integrations/artifact/%s/", url.PathEscape(integrationID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// IntegrationsArtifactPatch Update Artifact integration
func (c *Client) IntegrationsArtifactPatch(ctx context.Context, integrationID string, body vampy.IntegrationArtifact) (vampy.IntegrationArtifact, error) {
	path := fmt.Sprintf("/integrations/artifact/%s/", url.PathEscape(integrationID))
	var out vampy.IntegrationArtifact
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// IntegrationsArtifactExternalArtifactsGet Get integration's artifacts
func (c *Client) IntegrationsArtifactExternalArtifactsGet(ctx context.Context, integrationID string, offset *int, limit *int, query *string) (vampy.Page[vampy.ExternalArtifact], error) {
	path := fmt.Sprintf("/integrations/artifact/%s/items/", url.PathEscape(integrationID))
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if query != nil {
		q.Set("query", *query)
	}
	var out vampy.Page[vampy.ExternalArtifact]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}
