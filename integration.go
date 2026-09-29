package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/qvineox/vampy-go/vampy"
)

// SpaceIntegrationsListGet Get list spaces integrations with external services
func (c *Client) SpaceIntegrationsListGet(ctx context.Context, opts SpaceIntegrationsListGetOptions) ([]vampy.SpaceIntegrationWithMeta, error) {
	path := fmt.Sprintf("/spaces/%s/integrations/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	if opts.Service != nil {
		q.Set("service", string(*opts.Service))
	}

	if opts.Role != nil {
		q.Set("role", string(*opts.Role))
	}

	var out []vampy.SpaceIntegrationWithMeta
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceIntegrationsListGetOptions struct {
	Service *vampy.IntegrationServiceType
	Role    *vampy.IntegrationRole
}

// SpaceIntegrationsCreate Add new space's  integration
func (c *Client) SpaceIntegrationsCreate(ctx context.Context, opts SpaceIntegrationsCreateOptions) (vampy.SpaceIntegration, error) {
	path := fmt.Sprintf("/spaces/%s/integrations/", url.PathEscape(c.spaceID.String()))

	var out vampy.SpaceIntegration
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceIntegrationsCreateOptions struct {
	Body vampy.SpaceIntegration
}

// GetBitbucketApplicationInfo Get Bitbucket application version/authentication method by hostname
func (c *Client) GetBitbucketApplicationInfo(ctx context.Context, opts GetBitbucketApplicationInfoOptions) (vampy.BitbucketApplicationInfoResponse, error) {
	path := "/integrations/bitbucket/info/"

	var out vampy.BitbucketApplicationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetBitbucketApplicationInfoOptions struct {
	Body vampy.BitbucketApplicationInfoRequest
}

// IntegrationsVCSGet Get VCS integration
func (c *Client) IntegrationsVCSGet(ctx context.Context, opts IntegrationsVCSGetOptions) (vampy.IntegrationVCS, error) {
	path := fmt.Sprintf("/integrations/%s/%s/", url.PathEscape(string(opts.VCSIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationVCS
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsVCSGetOptions struct {
	VCSIntegration vampy.VCSIntegration
	IntegrationID  uuid.UUID
}

// IntegrationsVCSDelete Delete VCS integration
func (c *Client) IntegrationsVCSDelete(ctx context.Context, opts IntegrationsVCSDeleteOptions) error {
	path := fmt.Sprintf("/integrations/%s/%s/", url.PathEscape(string(opts.VCSIntegration)), url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type IntegrationsVCSDeleteOptions struct {
	VCSIntegration vampy.VCSIntegration
	IntegrationID  uuid.UUID
}

// IntegrationsVCSPatch Update VCS integration
func (c *Client) IntegrationsVCSPatch(ctx context.Context, opts IntegrationsVCSPatchOptions) (vampy.IntegrationVCS, error) {
	path := fmt.Sprintf("/integrations/%s/%s/", url.PathEscape(string(opts.VCSIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationVCS
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsVCSPatchOptions struct {
	VCSIntegration vampy.VCSIntegration
	IntegrationID  uuid.UUID
	Body           vampy.IntegrationVCS
}

// IntegrationsVCSRepositoriesGet Get VCS integration's repositories
func (c *Client) IntegrationsVCSRepositoriesGet(ctx context.Context, opts IntegrationsVCSRepositoriesGetOptions) (vampy.Page[vampy.ExternalRepository], error) {
	path := fmt.Sprintf("/integrations/%s/%s/repositories/", url.PathEscape(string(opts.VCSIntegration)), url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Query != nil {
		q.Set("query", *opts.Query)
	}

	var out vampy.Page[vampy.ExternalRepository]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsVCSRepositoriesGetOptions struct {
	VCSIntegration vampy.VCSIntegration
	IntegrationID  uuid.UUID
	Query          *string

	vampy.Paging
}

// IntegrationsVCSInfo Get VCS integration metadata (authentication, unsupported features)
func (c *Client) IntegrationsVCSInfo(ctx context.Context, opts IntegrationsVCSInfoOptions) (vampy.VCSIntegrationInfoResponse, error) {
	path := fmt.Sprintf("/integrations/vcs/%s/info/", url.PathEscape(string(opts.VCSIntegration)))

	var out vampy.VCSIntegrationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsVCSInfoOptions struct {
	VCSIntegration vampy.VCSIntegration
	Body           vampy.VCSIntegrationInfoRequest
}

// IntegrationsVCSMeta Get VCS integration metadata using stored credentials with optional overrides
func (c *Client) IntegrationsVCSMeta(ctx context.Context, opts IntegrationsVCSMetaOptions) (vampy.VCSIntegrationInfoResponse, error) {
	path := fmt.Sprintf("/integrations/vcs/%s/%s/meta/", url.PathEscape(string(opts.VCSIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.VCSIntegrationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsVCSMetaOptions struct {
	VCSIntegration vampy.VCSIntegration
	IntegrationID  uuid.UUID
	Body           vampy.VCSIntegrationInfoRequest
}

// GetJiraApplicationVersion Get Jira application version/authentication method by hostname
func (c *Client) GetJiraApplicationVersion(ctx context.Context, opts GetJiraApplicationVersionOptions) (vampy.JiraApplicationInfoResponse, error) {
	path := "/integrations/jira/info/"

	var out vampy.JiraApplicationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type GetJiraApplicationVersionOptions struct {
	Body vampy.JiraApplicationInfoRequest
}

// TaskTrackerGetInfo Get Task tracker service info method
func (c *Client) TaskTrackerGetInfo(ctx context.Context, opts TaskTrackerGetInfoOptions) (vampy.TaskTrackerInfoResponse, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/info/", url.PathEscape(string(opts.TaskTrackerIntegration)))

	var out vampy.TaskTrackerInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerGetInfoOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	Body                   vampy.TaskTrackerInfoRequest
}

// JiraIntegrationsMeta Get Jira integration instance meta information about projects and issues
func (c *Client) JiraIntegrationsMeta(ctx context.Context, opts JiraIntegrationsMetaOptions) (vampy.JiraMetaInfoResponse, error) {
	path := fmt.Sprintf("/integrations/jira/%s/meta/", url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	if opts.ProjectKey != nil {
		q.Set("projectKey", *opts.ProjectKey)
	}

	if opts.InitialQuery != nil {
		q.Set("initialQuery", strconv.FormatBool(*opts.InitialQuery))
	}

	var out vampy.JiraMetaInfoResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type JiraIntegrationsMetaOptions struct {
	IntegrationID uuid.UUID
	ProjectKey    *string
	InitialQuery  *bool
}

// TaskTrackerIntegrationGetMeta Get Task tracker integration instance meta information about projects and issues
func (c *Client) TaskTrackerIntegrationGetMeta(ctx context.Context, opts TaskTrackerIntegrationGetMetaOptions) (vampy.TaskTrackerMetaInfoResponse, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/meta/", url.PathEscape(string(opts.TaskTrackerIntegration)), url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	if opts.Project != nil {
		q.Set("project", *opts.Project)
	}

	if opts.IssueType != nil {
		q.Set("issueType", *opts.IssueType)
	}

	if opts.InitialQuery != nil {
		q.Set("initialQuery", strconv.FormatBool(*opts.InitialQuery))
	}

	var out vampy.TaskTrackerMetaInfoResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerIntegrationGetMetaOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	IntegrationID          uuid.UUID
	Project                *string
	IssueType              *string
	InitialQuery           *bool
}

// JiraIntegrationGet Get Jira integration
func (c *Client) JiraIntegrationGet(ctx context.Context, opts JiraIntegrationGetOptions) (vampy.JiraIntegration, error) {
	path := fmt.Sprintf("/integrations/jira/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.JiraIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type JiraIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// JiraIntegrationDelete Remove existing Jira integration.
func (c *Client) JiraIntegrationDelete(ctx context.Context, opts JiraIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/jira/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type JiraIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// JiraIntegrationPatch Set configuration values for existing Jira integration.
func (c *Client) JiraIntegrationPatch(ctx context.Context, opts JiraIntegrationPatchOptions) (vampy.JiraIntegration, error) {
	path := fmt.Sprintf("/integrations/jira/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.JiraIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type JiraIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.JiraIntegration
}

// TaskTrackerIntegrationGet Get Task tracker integration
func (c *Client) TaskTrackerIntegrationGet(ctx context.Context, opts TaskTrackerIntegrationGetOptions) (vampy.TaskTrackerIntegration, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/", url.PathEscape(string(opts.TaskTrackerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.TaskTrackerIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerIntegrationGetOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	IntegrationID          uuid.UUID
}

// TaskTrackerIntegrationDelete Remove existing Task tracker integration.
func (c *Client) TaskTrackerIntegrationDelete(ctx context.Context, opts TaskTrackerIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/", url.PathEscape(string(opts.TaskTrackerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type TaskTrackerIntegrationDeleteOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	IntegrationID          uuid.UUID
}

// TaskTrackerIntegrationPatch Set configuration values for existing Task tracker integration.
func (c *Client) TaskTrackerIntegrationPatch(ctx context.Context, opts TaskTrackerIntegrationPatchOptions) (vampy.TaskTrackerIntegration, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/", url.PathEscape(string(opts.TaskTrackerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.TaskTrackerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerIntegrationPatchOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	IntegrationID          uuid.UUID
	Body                   vampy.TaskTrackerIntegration
}

// SearchItemsJira Preview issues that are about to be sent to jira. Check if we have required default fields
func (c *Client) SearchItemsJira(ctx context.Context, opts SearchItemsJiraOptions) ([]vampy.JiraSearchResponse, error) {
	path := fmt.Sprintf("/integrations/jira/%s/search/", url.PathEscape(opts.IntegrationID))
	q := url.Values{}

	q.Set("field", opts.Field)

	q.Set("query", opts.Query)

	if opts.ProjectKey != nil {
		q.Set("projectKey", *opts.ProjectKey)
	}

	var out []vampy.JiraSearchResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SearchItemsJiraOptions struct {
	IntegrationID string
	Field         string
	Query         string
	ProjectKey    *string
}

// TaskTrackerSearchItems Search items in Task tracker
func (c *Client) TaskTrackerSearchItems(ctx context.Context, opts TaskTrackerSearchItemsOptions) ([]vampy.TaskTrackerSearchResponse, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/search/", url.PathEscape(string(opts.TaskTrackerIntegration)), url.PathEscape(opts.IntegrationID))
	q := url.Values{}

	q.Set("field", opts.Field)

	q.Set("query", opts.Query)

	if opts.Project != nil {
		q.Set("project", *opts.Project)
	}

	var out []vampy.TaskTrackerSearchResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerSearchItemsOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	IntegrationID          string
	Field                  string
	Query                  string
	Project                *string
}

// SearchIssuesJira Search Jira's issues
func (c *Client) SearchIssuesJira(ctx context.Context, opts SearchIssuesJiraOptions) ([]vampy.JiraSearchIssuesResponse, error) {
	path := fmt.Sprintf("/integrations/jira/%s/search/issues/", url.PathEscape(opts.IntegrationID))
	q := url.Values{}

	q.Set("query", opts.Query)

	var out []vampy.JiraSearchIssuesResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SearchIssuesJiraOptions struct {
	IntegrationID string
	Query         string
}

// TaskTrackerSearchIssues Search Task tracker issues
func (c *Client) TaskTrackerSearchIssues(ctx context.Context, opts TaskTrackerSearchIssuesOptions) ([]vampy.TaskTrackerIssueExport, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/search/issues/", url.PathEscape(string(opts.TaskTrackerIntegration)), url.PathEscape(opts.IntegrationID))
	q := url.Values{}

	q.Set("query", opts.Query)

	var out []vampy.TaskTrackerIssueExport
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerSearchIssuesOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	IntegrationID          string
	Query                  string
}

// KtrackerIntegrationGet Get KTracker integration
func (c *Client) KtrackerIntegrationGet(ctx context.Context, opts KtrackerIntegrationGetOptions) (vampy.IntegrationKTracker, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationKTracker
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type KtrackerIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// KtrackerIntegrationDelete Delete KTracker integration
func (c *Client) KtrackerIntegrationDelete(ctx context.Context, opts KtrackerIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/ktracker/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type KtrackerIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// KtrackerIntegrationPatch Update KTracker integration
func (c *Client) KtrackerIntegrationPatch(ctx context.Context, opts KtrackerIntegrationPatchOptions) (vampy.IntegrationKTracker, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationKTracker
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type KtrackerIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.IntegrationKTracker
}

// KtrackerIntegrationGetSpaces KTracker integration get Spaces
func (c *Client) KtrackerIntegrationGetSpaces(ctx context.Context, opts KtrackerIntegrationGetSpacesOptions) ([]vampy.IntegrationKTrackerSpace, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/spaces/", url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.ShowArchived != nil {
		q.Set("showArchived", strconv.FormatBool(*opts.ShowArchived))
	}

	var out []vampy.IntegrationKTrackerSpace
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type KtrackerIntegrationGetSpacesOptions struct {
	IntegrationID uuid.UUID
	ShowArchived  *bool

	vampy.Paging
}

// KtrackerIntegrationGetCardTypes KTracker integration get Card types
func (c *Client) KtrackerIntegrationGetCardTypes(ctx context.Context, opts KtrackerIntegrationGetCardTypesOptions) ([]vampy.IntegrationKTrackerCardType, error) {
	path := fmt.Sprintf("/integrations/ktracker/%s/card_types/", url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	if opts.SpaceID != nil {
		q.Set("spaceID", strconv.Itoa(*opts.SpaceID))
	}

	var out []vampy.IntegrationKTrackerCardType
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type KtrackerIntegrationGetCardTypesOptions struct {
	IntegrationID uuid.UUID
	SpaceID       *int
}

// ScannerIntegrationGet Get scanner integration
func (c *Client) ScannerIntegrationGet(ctx context.Context, opts ScannerIntegrationGetOptions) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/", url.PathEscape(string(opts.ScannerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScannerIntegrationGetOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	IntegrationID      uuid.UUID
}

// ScannerIntegrationDelete Delete scanner integration
func (c *Client) ScannerIntegrationDelete(ctx context.Context, opts ScannerIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/", url.PathEscape(string(opts.ScannerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type ScannerIntegrationDeleteOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	IntegrationID      uuid.UUID
}

// ScannerIntegrationPatch Update scanner integration
func (c *Client) ScannerIntegrationPatch(ctx context.Context, opts ScannerIntegrationPatchOptions) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/", url.PathEscape(string(opts.ScannerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScannerIntegrationPatchOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	IntegrationID      uuid.UUID
	Body               vampy.ScannerIntegrationPatch
}

// ScannerIntegrationMeta Get scanner integration instance meta information
func (c *Client) ScannerIntegrationMeta(ctx context.Context, opts ScannerIntegrationMetaOptions) (vampy.ScannerIntegrationMetaResponse, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/meta/", url.PathEscape(string(opts.ScannerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegrationMetaResponse
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScannerIntegrationMetaOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	IntegrationID      uuid.UUID
}

// ScannerIntegrationInfo Get scanner integration info
func (c *Client) ScannerIntegrationInfo(ctx context.Context, opts ScannerIntegrationInfoOptions) (vampy.ScannerIntegrationInfoResponse, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/info/", url.PathEscape(string(opts.ScannerIntegration)))

	var out vampy.ScannerIntegrationInfoResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScannerIntegrationInfoOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	Body               vampy.ScannerIntegrationInfoRequest
}

// ScannerIntegrationFetchCLI Fetch CLI from scanner integration API
func (c *Client) ScannerIntegrationFetchCLI(ctx context.Context, opts ScannerIntegrationFetchCLIOptions) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/fetch_cli/", url.PathEscape(string(opts.ScannerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScannerIntegrationFetchCLIOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	IntegrationID      uuid.UUID
}

// ScannerIntegrationUploadCLI Upload CLI file
func (c *Client) ScannerIntegrationUploadCLI(ctx context.Context, opts ScannerIntegrationUploadCLIOptions) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/upload_cli/", url.PathEscape(string(opts.ScannerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type ScannerIntegrationUploadCLIOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	IntegrationID      uuid.UUID
	Body               vampy.IntegrationIDUploadCLIBody
}

// ScannerIntegrationTestConnection Test scanner integration connection
func (c *Client) ScannerIntegrationTestConnection(ctx context.Context, opts ScannerIntegrationTestConnectionOptions) error {
	path := fmt.Sprintf("/integrations/scanner/%s/%s/test_connection/", url.PathEscape(string(opts.ScannerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type ScannerIntegrationTestConnectionOptions struct {
	ScannerIntegration vampy.ScannerIntegrationType
	IntegrationID      uuid.UUID
	Body               vampy.ScannerIntegrationTestConnection
}

// PtaiIntegrationGet Get PTAI integration
func (c *Client) PtaiIntegrationGet(ctx context.Context, opts PtaiIntegrationGetOptions) (vampy.IntegrationPTAI, error) {
	path := fmt.Sprintf("/integrations/ptai/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationPTAI
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type PtaiIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// PtaiIntegrationDelete Delete PTAI integration
func (c *Client) PtaiIntegrationDelete(ctx context.Context, opts PtaiIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/ptai/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type PtaiIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// PtaiIntegrationPatch Update PTAI integration
func (c *Client) PtaiIntegrationPatch(ctx context.Context, opts PtaiIntegrationPatchOptions) (vampy.IntegrationPTAI, error) {
	path := fmt.Sprintf("/integrations/ptai/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationPTAI
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type PtaiIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.IntegrationPTAI
}

// PtaiIntegrationTestConnection Test PTAI connection
func (c *Client) PtaiIntegrationTestConnection(ctx context.Context, opts PtaiIntegrationTestConnectionOptions) error {
	path := fmt.Sprintf("/integrations/ptai/%s/test_connection/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type PtaiIntegrationTestConnectionOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.ScannerIntegrationTestConnection
}

// TaskTrackerExportMeta Get metadata for scan issues to be exported to task tracker.
func (c *Client) TaskTrackerExportMeta(ctx context.Context, opts TaskTrackerExportMetaOptions) (vampy.ScanIssueForExportResponse, error) {
	path := fmt.Sprintf("/integrations/%s/%s/export/meta/", url.PathEscape(string(opts.IntegrationService)), url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScanIssueForExportResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerExportMetaOptions struct {
	IntegrationService vampy.TaskTrackerServiceType
	IntegrationID      uuid.UUID
	Body               vampy.TaskTrackerExportMetaRequest
}

// TaskTrackerExportAsync Export scan issues to task tracker asynchronously
func (c *Client) TaskTrackerExportAsync(ctx context.Context, opts TaskTrackerExportAsyncOptions) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/integrations/%s/%s/export/async/", url.PathEscape(string(opts.IntegrationService)), url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	if opts.TaskState != nil {
		q.Set("taskState", string(*opts.TaskState))
	}

	if opts.FailedScanIssueID != nil {
		q.Set("failedScanIssueID", opts.FailedScanIssueID.String())
	}

	var out vampy.BackgroundTask
	if err := c.doJSON(ctx, "POST", path, q, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerExportAsyncOptions struct {
	IntegrationService vampy.TaskTrackerServiceType
	IntegrationID      uuid.UUID
	TaskState          *vampy.BackgroundTaskStatus
	FailedScanIssueID  *uuid.UUID
	Body               vampy.TaskTrackerAsyncExport
}

// TaskTrackerLinkIssue Link scan issue to an existing external task tracker issue
func (c *Client) TaskTrackerLinkIssue(ctx context.Context, opts TaskTrackerLinkIssueOptions) ([]vampy.IssueExport, error) {
	path := fmt.Sprintf("/integrations/task_tracker/%s/%s/link/", url.PathEscape(string(opts.TaskTrackerIntegration)), url.PathEscape(opts.IntegrationID.String()))

	var out []vampy.IssueExport
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type TaskTrackerLinkIssueOptions struct {
	TaskTrackerIntegration vampy.TaskTrackerIntegrationType
	IntegrationID          uuid.UUID
	Body                   vampy.TaskTrackerLinkRequest
}

// CodeScoringIntegrationGet Get CodeScoring integration
func (c *Client) CodeScoringIntegrationGet(ctx context.Context, opts CodeScoringIntegrationGetOptions) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CodeScoringIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// CodeScoringIntegrationDelete Delete CodeScoring integration
func (c *Client) CodeScoringIntegrationDelete(ctx context.Context, opts CodeScoringIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/code_scoring/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type CodeScoringIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// CodeScoringIntegrationPatch Update CodeScoring integration
func (c *Client) CodeScoringIntegrationPatch(ctx context.Context, opts CodeScoringIntegrationPatchOptions) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CodeScoringIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.ScannerIntegrationPatch
}

// CodeScoringIntegrationMeta Get CodeScoring integration instance meta information
func (c *Client) CodeScoringIntegrationMeta(ctx context.Context, opts CodeScoringIntegrationMetaOptions) (vampy.ScannerIntegrationMetaResponse, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/meta/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegrationMetaResponse
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CodeScoringIntegrationMetaOptions struct {
	IntegrationID uuid.UUID
}

// CodeScoringIntegrationFetchCLI Fetch CLI from CodeScoring API
func (c *Client) CodeScoringIntegrationFetchCLI(ctx context.Context, opts CodeScoringIntegrationFetchCLIOptions) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/fetch_cli/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CodeScoringIntegrationFetchCLIOptions struct {
	IntegrationID uuid.UUID
}

// CodeScoringIntegrationUploadCLI Upload CLI file
func (c *Client) CodeScoringIntegrationUploadCLI(ctx context.Context, opts CodeScoringIntegrationUploadCLIOptions) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/code_scoring/%s/upload_cli/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type CodeScoringIntegrationUploadCLIOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.IntegrationIDUploadCLIBody1
}

// AppScreenerIntegrationGet Get AppScreener integration
func (c *Client) AppScreenerIntegrationGet(ctx context.Context, opts AppScreenerIntegrationGetOptions) (vampy.IntegrationAppScreener, error) {
	path := fmt.Sprintf("/integrations/app_screener/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationAppScreener
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AppScreenerIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// AppScreenerIntegrationDelete Delete AppScreener integration
func (c *Client) AppScreenerIntegrationDelete(ctx context.Context, opts AppScreenerIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/app_screener/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type AppScreenerIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// AppScreenerIntegrationPatch Update AppScreener integration
func (c *Client) AppScreenerIntegrationPatch(ctx context.Context, opts AppScreenerIntegrationPatchOptions) (vampy.IntegrationAppScreener, error) {
	path := fmt.Sprintf("/integrations/app_screener/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationAppScreener
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AppScreenerIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.ScannerIntegrationPatch
}

// AppScreenerIntegrationUploadCLI Upload CLI file
func (c *Client) AppScreenerIntegrationUploadCLI(ctx context.Context, opts AppScreenerIntegrationUploadCLIOptions) (vampy.ScannerIntegrationCLIFileInfo, error) {
	path := fmt.Sprintf("/integrations/app_screener/%s/upload_cli/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegrationCLIFileInfo
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AppScreenerIntegrationUploadCLIOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.IntegrationIDUploadCLIBody2
}

// SastavIntegrationGet Get SASTAV integration
func (c *Client) SastavIntegrationGet(ctx context.Context, opts SastavIntegrationGetOptions) (vampy.IntegrationSastav, error) {
	path := fmt.Sprintf("/integrations/sastav/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationSastav
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SastavIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// SastavIntegrationDelete Delete SASTAV integration
func (c *Client) SastavIntegrationDelete(ctx context.Context, opts SastavIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/sastav/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type SastavIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// SastavIntegrationPatch Update SASTAV integration
func (c *Client) SastavIntegrationPatch(ctx context.Context, opts SastavIntegrationPatchOptions) (vampy.IntegrationSastav, error) {
	path := fmt.Sprintf("/integrations/sastav/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationSastav
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SastavIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.ScannerIntegrationPatch
}

// DefectDojoIntegrationGet Get Defect Dojo integration
func (c *Client) DefectDojoIntegrationGet(ctx context.Context, opts DefectDojoIntegrationGetOptions) (vampy.IntegrationDefectDojo, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationDefectDojo
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type DefectDojoIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// DefectDojoIntegrationDelete Delete Defect Dojo integration
func (c *Client) DefectDojoIntegrationDelete(ctx context.Context, opts DefectDojoIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type DefectDojoIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// DefectDojoIntegrationPatch Update Defect Dojo integration
func (c *Client) DefectDojoIntegrationPatch(ctx context.Context, opts DefectDojoIntegrationPatchOptions) (vampy.IntegrationDefectDojo, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationDefectDojo
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type DefectDojoIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.IntegrationDefectDojo
}

// DefectDojoIntegrationGetProducts Get list of products from Defect Dojo
func (c *Client) DefectDojoIntegrationGetProducts(ctx context.Context, opts DefectDojoIntegrationGetProductsOptions) (vampy.Page[vampy.DefectDojoProduct], error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/products/", url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.Page[vampy.DefectDojoProduct]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type DefectDojoIntegrationGetProductsOptions struct {
	IntegrationID uuid.UUID
	Search        *string

	vampy.Paging
}

// DefectDojoIntegrationGetStatuses Get DefectDojo statuses mapping
func (c *Client) DefectDojoIntegrationGetStatuses(ctx context.Context, opts DefectDojoIntegrationGetStatusesOptions) ([]vampy.DefectDojoStatusMapping, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/statuses/", url.PathEscape(opts.IntegrationID.String()))

	var out []vampy.DefectDojoStatusMapping
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type DefectDojoIntegrationGetStatusesOptions struct {
	IntegrationID uuid.UUID
}

// DefectDojoIntegrationImport Import data from Defect Dojo: findings, products, repos
func (c *Client) DefectDojoIntegrationImport(ctx context.Context, opts DefectDojoIntegrationImportOptions) (vampy.DefectDojoImportResponse, error) {
	path := fmt.Sprintf("/integrations/defect_dojo/%s/import/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.DefectDojoImportResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type DefectDojoIntegrationImportOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.DefectDojoImportRequest
}

// AIIntegrationGet Get AI integration
func (c *Client) AIIntegrationGet(ctx context.Context, opts AIIntegrationGetOptions) (vampy.IntegrationAI, error) {
	path := fmt.Sprintf("/integrations/ai/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationAI
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AIIntegrationGetOptions struct {
	IntegrationID uuid.UUID
}

// AIIntegrationDelete Delete AI integration
func (c *Client) AIIntegrationDelete(ctx context.Context, opts AIIntegrationDeleteOptions) error {
	path := fmt.Sprintf("/integrations/ai/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type AIIntegrationDeleteOptions struct {
	IntegrationID uuid.UUID
}

// AIIntegrationPatch Update AI integration
func (c *Client) AIIntegrationPatch(ctx context.Context, opts AIIntegrationPatchOptions) (vampy.ScannerIntegration, error) {
	path := fmt.Sprintf("/integrations/ai/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.ScannerIntegration
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AIIntegrationPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.IntegrationAI
}

// AIIntegrationFetchModels Update AI models
func (c *Client) AIIntegrationFetchModels(ctx context.Context, opts AIIntegrationFetchModelsOptions) (vampy.IntegrationAI, error) {
	path := fmt.Sprintf("/integrations/ai/%s/fetch_models/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationAI
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AIIntegrationFetchModelsOptions struct {
	IntegrationID uuid.UUID
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
func (c *Client) AIIntegrationAppliedModels(ctx context.Context, opts AIIntegrationAppliedModelsOptions) ([]vampy.IntegrationAIModel, error) {
	path := "/integrations/ai/models/applied/"
	q := url.Values{}

	if opts.RequestType != nil {
		q.Set("requestType", string(*opts.RequestType))
	}

	var out []vampy.IntegrationAIModel
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type AIIntegrationAppliedModelsOptions struct {
	RequestType *vampy.IntegrationAIRequestType
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
func (c *Client) IntegrationWebhookHandlerPost(ctx context.Context, opts IntegrationWebhookHandlerPostOptions) error {
	path := fmt.Sprintf("/integrations/servers/%s/webhook/", url.PathEscape(opts.Token.String()))

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type IntegrationWebhookHandlerPostOptions struct {
	Token uuid.UUID
	Body  vampy.TokenWebhookBody
}

// IntegrationsArtifactDetailsGet Get Artifact integration
func (c *Client) IntegrationsArtifactDetailsGet(ctx context.Context, opts IntegrationsArtifactDetailsGetOptions) (vampy.IntegrationArtifact, error) {
	path := fmt.Sprintf("/integrations/artifact/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationArtifact
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsArtifactDetailsGetOptions struct {
	IntegrationID uuid.UUID
}

// IntegrationsArtifactDelete Delete Artifact integration
func (c *Client) IntegrationsArtifactDelete(ctx context.Context, opts IntegrationsArtifactDeleteOptions) error {
	path := fmt.Sprintf("/integrations/artifact/%s/", url.PathEscape(opts.IntegrationID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type IntegrationsArtifactDeleteOptions struct {
	IntegrationID uuid.UUID
}

// IntegrationsArtifactPatch Update Artifact integration
func (c *Client) IntegrationsArtifactPatch(ctx context.Context, opts IntegrationsArtifactPatchOptions) (vampy.IntegrationArtifact, error) {
	path := fmt.Sprintf("/integrations/artifact/%s/", url.PathEscape(opts.IntegrationID.String()))

	var out vampy.IntegrationArtifact
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsArtifactPatchOptions struct {
	IntegrationID uuid.UUID
	Body          vampy.IntegrationArtifact
}

// IntegrationsArtifactExternalArtifactsGet Get integration's artifacts
func (c *Client) IntegrationsArtifactExternalArtifactsGet(ctx context.Context, opts IntegrationsArtifactExternalArtifactsGetOptions) (vampy.Page[vampy.ExternalArtifact], error) {
	path := fmt.Sprintf("/integrations/artifact/%s/items/", url.PathEscape(opts.IntegrationID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.Query != nil {
		q.Set("query", *opts.Query)
	}

	var out vampy.Page[vampy.ExternalArtifact]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type IntegrationsArtifactExternalArtifactsGetOptions struct {
	IntegrationID uuid.UUID
	Query         *string

	vampy.Paging
}
