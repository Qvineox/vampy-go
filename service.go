package vampy_go

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"uuid"

	"github.com/qvineox/vampy-go/vampy"
)

// SpaceIssueStatusList Get space issue statuses list
func (c *Client) SpaceIssueStatusList(ctx context.Context, opts SpaceIssueStatusListOptions) ([]vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	if opts.Deleted != nil {
		q.Set("deleted", strconv.FormatBool(*opts.Deleted))
	}

	var out []vampy.IssueStatus
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceIssueStatusListOptions struct {
	Deleted *bool
}

// SpaceIssueStatusCreate Create new issue status
func (c *Client) SpaceIssueStatusCreate(ctx context.Context, opts SpaceIssueStatusCreateOptions) (vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/", url.PathEscape(c.spaceID.String()))

	var out vampy.IssueStatus
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceIssueStatusCreateOptions struct {
	Body vampy.IssueStatusCreate
}

// SpaceIssueStatusGet Get issue status
func (c *Client) SpaceIssueStatusGet(ctx context.Context, opts SpaceIssueStatusGetOptions) (vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/%s/", url.PathEscape(c.spaceID.String()), url.PathEscape(opts.IssueStatusName))

	var out vampy.IssueStatus
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceIssueStatusGetOptions struct {
	IssueStatusName string
}

// SpaceIssueStatusDelete Delete issue status
func (c *Client) SpaceIssueStatusDelete(ctx context.Context, opts SpaceIssueStatusDeleteOptions) error {
	path := fmt.Sprintf("/spaces/%s/statuses/%s/", url.PathEscape(c.spaceID.String()), url.PathEscape(opts.IssueStatusName))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type SpaceIssueStatusDeleteOptions struct {
	IssueStatusName string
}

// SpaceIssueStatusPatch Update existing issue status
func (c *Client) SpaceIssueStatusPatch(ctx context.Context, opts SpaceIssueStatusPatchOptions) (vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/%s/", url.PathEscape(c.spaceID.String()), url.PathEscape(opts.IssueStatusName))

	var out vampy.IssueStatus
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SpaceIssueStatusPatchOptions struct {
	IssueStatusName string
	Body            vampy.IssueStatusUpdate
}

// SystemSettingsGet Get system-wide settings and public engine constants.
func (c *Client) SystemSettingsGet(ctx context.Context) (vampy.SystemWideSettings, error) {
	path := "/service/settings/"

	var out vampy.SystemWideSettings
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// SystemSettingsPatch Update system settings.
func (c *Client) SystemSettingsPatch(ctx context.Context, opts SystemSettingsPatchOptions) (vampy.SystemWideSettings, error) {
	path := "/service/settings/"

	var out vampy.SystemWideSettings
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SystemSettingsPatchOptions struct {
	Body vampy.SystemWideSettings
}

// FilterValuesGet Get possible values for given filter targets
func (c *Client) FilterValuesGet(ctx context.Context, opts FilterValuesGetOptions) (vampy.FilterMetaResponse, error) {
	path := fmt.Sprintf("/service/meta/%s/filter_values/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	for _, v := range opts.Target {
		q.Add("target", string(v))
	}

	if opts.RelationType != nil {
		q.Set("relationType", string(*opts.RelationType))
	}

	if opts.RelationID != nil {
		q.Set("relationID", opts.RelationID.String())
	}

	if opts.Search != nil {
		q.Set("search", *opts.Search)
	}

	var out vampy.FilterMetaResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FilterValuesGetOptions struct {
	Target       []vampy.FilterMetaTarget
	RelationType *vampy.FilterMetaRelationType
	RelationID   *uuid.UUID
	Search       *string
}

// PingGet Check if Vampy service is up
func (c *Client) PingGet(ctx context.Context) error {
	path := "/service/ping/"

	return c.doJSON(ctx, "GET", path, nil, nil, nil)
}

// HealthGet Check if Vampy-dependant services is up
func (c *Client) HealthGet(ctx context.Context) (vampy.HealthResponse, error) {
	path := "/service/health/"

	var out vampy.HealthResponse
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// ServerInfo Get various server-related info for client
func (c *Client) ServerInfo(ctx context.Context) (vampy.ServerInfo, error) {
	path := "/service/info/"

	var out vampy.ServerInfo
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// LDAPConfigGet Get existing LDAP authentication config
func (c *Client) LDAPConfigGet(ctx context.Context) (vampy.LDAPAuthenticationConfig, error) {
	path := "/service/ldap/"

	var out vampy.LDAPAuthenticationConfig
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// LDAPConfigPost Create new LDAP authentication config
func (c *Client) LDAPConfigPost(ctx context.Context, opts LDAPConfigPostOptions) (vampy.LDAPAuthenticationConfig, error) {
	path := "/service/ldap/"

	var out vampy.LDAPAuthenticationConfig
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type LDAPConfigPostOptions struct {
	Body vampy.LDAPAuthenticationConfig
}

// LDAPConfigDelete Delete LDAP authentication config
func (c *Client) LDAPConfigDelete(ctx context.Context) error {
	path := "/service/ldap/"

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// LDAPConfigPatch Update LDAP authentication config
func (c *Client) LDAPConfigPatch(ctx context.Context, opts LDAPConfigPatchOptions) (vampy.LDAPAuthenticationConfig, error) {
	path := "/service/ldap/"

	var out vampy.LDAPAuthenticationConfig
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type LDAPConfigPatchOptions struct {
	Body vampy.LDAPAuthenticationConfig
}

// LDAPConfigTest Test LDAP authentication config without creation
func (c *Client) LDAPConfigTest(ctx context.Context, opts LDAPConfigTestOptions) (vampy.Page[vampy.LDAPUser], error) {
	path := "/service/ldap/test/"

	var out vampy.Page[vampy.LDAPUser]
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type LDAPConfigTestOptions struct {
	Body vampy.LDAPAuthenticationConfig
}

// OIDCConfigList Get existing OpenID connect configs
func (c *Client) OIDCConfigList(ctx context.Context) ([]vampy.OIDCConfig, error) {
	path := "/service/oidc/"

	var out []vampy.OIDCConfig
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// OIDCConfigPost Create new OpenID connect config
func (c *Client) OIDCConfigPost(ctx context.Context, opts OIDCConfigPostOptions) (vampy.OIDCConfig, error) {
	path := "/service/oidc/"

	var out vampy.OIDCConfig
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type OIDCConfigPostOptions struct {
	Body vampy.OIDCConfig
}

// OIDCConfigGet Get existing OpenID connect configs
func (c *Client) OIDCConfigGet(ctx context.Context, opts OIDCConfigGetOptions) (vampy.OIDCConfig, error) {
	path := fmt.Sprintf("/service/oidc/%s/", url.PathEscape(opts.ConfigID.String()))

	var out vampy.OIDCConfig
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type OIDCConfigGetOptions struct {
	ConfigID uuid.UUID
}

// OIDCConfigDelete Delete existing OpenID connect config
func (c *Client) OIDCConfigDelete(ctx context.Context, opts OIDCConfigDeleteOptions) error {
	path := fmt.Sprintf("/service/oidc/%s/", url.PathEscape(opts.ConfigID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type OIDCConfigDeleteOptions struct {
	ConfigID uuid.UUID
}

// OIDCConfigPatch Patch existing OpenID connect config
func (c *Client) OIDCConfigPatch(ctx context.Context, opts OIDCConfigPatchOptions) (vampy.OIDCConfig, error) {
	path := fmt.Sprintf("/service/oidc/%s/", url.PathEscape(opts.ConfigID.String()))

	var out vampy.OIDCConfig
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type OIDCConfigPatchOptions struct {
	ConfigID uuid.UUID
	Body     vampy.OIDCConfig
}

// OIDCConfigDefaults Get default value for OIDC provider attributes
func (c *Client) OIDCConfigDefaults(ctx context.Context, opts OIDCConfigDefaultsOptions) (vampy.OIDCDefaults, error) {
	path := fmt.Sprintf("/service/oidc/%s/defaults/", url.PathEscape(string(opts.Provider)))
	q := url.Values{}

	if opts.URL != nil {
		q.Set("url", *opts.URL)
	}

	if opts.Realm != nil {
		q.Set("realm", *opts.Realm)
	}

	var out vampy.OIDCDefaults
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type OIDCConfigDefaultsOptions struct {
	Provider vampy.OIDCProvider
	URL      *string
	Realm    *string
}

// OIDCTestConnect Test connection to OIDC provider server
func (c *Client) OIDCTestConnect(ctx context.Context, opts OIDCTestConnectOptions) ([]vampy.OIDCTestConnectResponse, error) {
	path := fmt.Sprintf("/service/oidc/%s/test_connect/", url.PathEscape(string(opts.Provider)))

	var out []vampy.OIDCTestConnectResponse
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type OIDCTestConnectOptions struct {
	Provider vampy.OIDCProvider
	Body     vampy.OIDCTestConnectRequest
}

// SMTPConfigGet Get existing SMTP authentication config
func (c *Client) SMTPConfigGet(ctx context.Context) (vampy.SMTPConfig, error) {
	path := "/service/smtp_config/"

	var out vampy.SMTPConfig
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// SMTPConfigPost Create new SMTP authentication config
func (c *Client) SMTPConfigPost(ctx context.Context, opts SMTPConfigPostOptions) (vampy.SMTPConfig, error) {
	path := "/service/smtp_config/"

	var out vampy.SMTPConfig
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SMTPConfigPostOptions struct {
	Body vampy.SMTPConfig
}

// SMTPConfigDelete Delete SMTP authentication config
func (c *Client) SMTPConfigDelete(ctx context.Context) error {
	path := "/service/smtp_config/"

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// SMTPConfigPatch Update already an existing SMTP authentication config
func (c *Client) SMTPConfigPatch(ctx context.Context, opts SMTPConfigPatchOptions) (vampy.SMTPConfig, error) {
	path := "/service/smtp_config/"

	var out vampy.SMTPConfig
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type SMTPConfigPatchOptions struct {
	Body vampy.SMTPConfig
}

func (c *Client) SendTestEmail(ctx context.Context, opts SendTestEmailOptions) error {
	path := "/service/smtp_config/test_email/"

	return c.doJSON(ctx, "POST", path, nil, opts.Body, nil)
}

type SendTestEmailOptions struct {
	Body vampy.Email
}

// UserRolesGet Get all roles
func (c *Client) UserRolesGet(ctx context.Context, opts UserRolesGetOptions) ([]vampy.UserRole, error) {
	path := "/service/user_roles/"
	q := url.Values{}

	if opts.System != nil {
		q.Set("system", strconv.FormatBool(*opts.System))
	}

	var out []vampy.UserRole
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UserRolesGetOptions struct {
	System *bool
}

// UserRolesCreate Create a new role
func (c *Client) UserRolesCreate(ctx context.Context, opts UserRolesCreateOptions) (vampy.UserRole, error) {
	path := "/service/user_roles/"

	var out vampy.UserRole
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UserRolesCreateOptions struct {
	Body vampy.UserRoleEdit
}

// UserRolesPermissionsMeta Get meta information about permissions
func (c *Client) UserRolesPermissionsMeta(ctx context.Context) (vampy.UserRolePermissionsMeta, error) {
	path := "/service/user_roles/permissions/meta/"

	var out vampy.UserRolePermissionsMeta
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// UserRolesDetailsGet Get details of a user role
func (c *Client) UserRolesDetailsGet(ctx context.Context, opts UserRolesDetailsGetOptions) (vampy.UserRole, error) {
	path := fmt.Sprintf("/service/user_roles/%s/", url.PathEscape(opts.UserRoleID.String()))

	var out vampy.UserRole
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type UserRolesDetailsGetOptions struct {
	UserRoleID uuid.UUID
}

// UserRolesDelete Delete user role
func (c *Client) UserRolesDelete(ctx context.Context, opts UserRolesDeleteOptions) error {
	path := fmt.Sprintf("/service/user_roles/%s/", url.PathEscape(opts.UserRoleID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type UserRolesDeleteOptions struct {
	UserRoleID uuid.UUID
}

// UserRolesPatch Edit existing user role
func (c *Client) UserRolesPatch(ctx context.Context, opts UserRolesPatchOptions) error {
	path := fmt.Sprintf("/service/user_roles/%s/", url.PathEscape(opts.UserRoleID.String()))

	return c.doJSON(ctx, "PATCH", path, nil, opts.Body, nil)
}

type UserRolesPatchOptions struct {
	UserRoleID uuid.UUID
	Body       vampy.UserRoleEdit
}

func (c *Client) GetParsers(ctx context.Context) ([]vampy.ParserInList, error) {
	path := "/service/parsers/"

	var out []vampy.ParserInList
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// BackgroundTasksListGet Paginated list of current tasks execution
func (c *Client) BackgroundTasksListGet(ctx context.Context, opts BackgroundTasksListGetOptions) (vampy.Page[vampy.BackgroundTask], error) {
	path := "/service/background_tasks/"
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	var out vampy.Page[vampy.BackgroundTask]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type BackgroundTasksListGetOptions struct {
	vampy.Paging
}

// BackgroundTasksDetailsGet Detailed info about background task
func (c *Client) BackgroundTasksDetailsGet(ctx context.Context, opts BackgroundTasksDetailsGetOptions) (vampy.BackgroundTaskDetails, error) {
	path := fmt.Sprintf("/service/background_tasks/%s/", url.PathEscape(opts.TaskID.String()))

	var out vampy.BackgroundTaskDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type BackgroundTasksDetailsGetOptions struct {
	TaskID uuid.UUID
}

// GetGlobalSLASetting Get information about sla global setting
func (c *Client) GetGlobalSLASetting(ctx context.Context) (vampy.SLASettingWithDefault, error) {
	path := "/service/sla_config/"

	var out vampy.SLASettingWithDefault
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// PostGlobalSLASetting Create or update global sla setting
func (c *Client) PostGlobalSLASetting(ctx context.Context, opts PostGlobalSLASettingOptions) (vampy.SLASetting, error) {
	path := "/service/sla_config/"

	var out vampy.SLASetting
	if err := c.doJSON(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type PostGlobalSLASettingOptions struct {
	Body vampy.SLASetting
}

// FPClassifierDatasetDownload Download scan issues dataset
func (c *Client) FPClassifierDatasetDownload(ctx context.Context) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/download/", url.PathEscape(c.spaceID.String()))

	return c.doBinary(ctx, "GET", path, nil, nil)
}

// FPClassifierDatasetUpload Upload scan issues dataset
func (c *Client) FPClassifierDatasetUpload(ctx context.Context, opts FPClassifierDatasetUploadOptions) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/upload/", url.PathEscape(c.spaceID.String()))

	var out vampy.BackgroundTask
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FPClassifierDatasetUploadOptions struct {
	Body vampy.DatasetUploadBody
}

// FPClassifierDatasetClear Clear scan issues dataset
func (c *Client) FPClassifierDatasetClear(ctx context.Context) error {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/clear/", url.PathEscape(c.spaceID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// FPClassifierTrainingDataStats Get accumulated training data statistics
func (c *Client) FPClassifierTrainingDataStats(ctx context.Context) ([]vampy.FPClassifierDatasetStats, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/stats/", url.PathEscape(c.spaceID.String()))

	var out []vampy.FPClassifierDatasetStats
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// FPClassifierSnapshotList Get FP classifier snapshots list
func (c *Client) FPClassifierSnapshotList(ctx context.Context, opts FPClassifierSnapshotListOptions) (vampy.Page[vampy.FPClassifierSnapshot], error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/snapshots/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("offset", strconv.Itoa(opts.Offset))
	q.Set("limit", strconv.Itoa(opts.Limit))

	if opts.ScannerType != nil {
		q.Set("scannerType", string(*opts.ScannerType))
	}

	if opts.TestTotalCount != nil {
		q.Set("_test_total_count", strconv.Itoa(*opts.TestTotalCount))
	}

	var out vampy.Page[vampy.FPClassifierSnapshot]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FPClassifierSnapshotListOptions struct {
	ScannerType    *vampy.ScannerType
	TestTotalCount *int

	vampy.Paging
}

// FPClassifierSnapshotsDownload Download active FP classifier model files archive
func (c *Client) FPClassifierSnapshotsDownload(ctx context.Context) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/snapshots/download/", url.PathEscape(c.spaceID.String()))

	return c.doBinary(ctx, "GET", path, nil, nil)
}

// FPClassifierSnapshotUpload Upload FP classifier model archive
func (c *Client) FPClassifierSnapshotUpload(ctx context.Context, opts FPClassifierSnapshotUploadOptions) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/snapshots/upload/", url.PathEscape(c.spaceID.String()))

	var out vampy.BackgroundTask
	if err := c.doMultipart(ctx, "POST", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FPClassifierSnapshotUploadOptions struct {
	Body vampy.SnapshotsUploadBody
}

// FPClassifierSnapshotGet Get FP classifier snapshot
func (c *Client) FPClassifierSnapshotGet(ctx context.Context, opts FPClassifierSnapshotGetOptions) (vampy.FPClassifierSnapshot, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/", url.PathEscape(opts.SnapshotID.String()))

	var out vampy.FPClassifierSnapshot
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FPClassifierSnapshotGetOptions struct {
	SnapshotID uuid.UUID
}

// FPClassifierSnapshotDelete Delete FP classifier snapshot
func (c *Client) FPClassifierSnapshotDelete(ctx context.Context, opts FPClassifierSnapshotDeleteOptions) error {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/", url.PathEscape(opts.SnapshotID.String()))

	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

type FPClassifierSnapshotDeleteOptions struct {
	SnapshotID uuid.UUID
}

// FPClassifierSnapshotPatch Update FP classifier snapshot
func (c *Client) FPClassifierSnapshotPatch(ctx context.Context, opts FPClassifierSnapshotPatchOptions) (vampy.FPClassifierSnapshot, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/", url.PathEscape(opts.SnapshotID.String()))

	var out vampy.FPClassifierSnapshot
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FPClassifierSnapshotPatchOptions struct {
	SnapshotID uuid.UUID
	Body       vampy.SnapshotsSnapshotIDBody
}

// FPClassifierSnapshotTrain Train FP classifier snapshot
func (c *Client) FPClassifierSnapshotTrain(ctx context.Context, opts FPClassifierSnapshotTrainOptions) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/train/", url.PathEscape(opts.SnapshotID.String()))

	var out vampy.BackgroundTask
	if err := c.doJSON(ctx, "POST", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FPClassifierSnapshotTrainOptions struct {
	SnapshotID uuid.UUID
}

// FPClassifierSnapshotDownload Download FP classifier model archive
func (c *Client) FPClassifierSnapshotDownload(ctx context.Context, opts FPClassifierSnapshotDownloadOptions) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/download/", url.PathEscape(opts.SnapshotID.String()))

	return c.doBinary(ctx, "GET", path, nil, nil)
}

type FPClassifierSnapshotDownloadOptions struct {
	SnapshotID uuid.UUID
}

// FPClassifierSnapshotDatasetDownload Download FP classifier snapshot dataset
func (c *Client) FPClassifierSnapshotDatasetDownload(ctx context.Context, opts FPClassifierSnapshotDatasetDownloadOptions) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/dataset/download/", url.PathEscape(opts.SnapshotID.String()))

	return c.doBinary(ctx, "GET", path, nil, nil)
}

type FPClassifierSnapshotDatasetDownloadOptions struct {
	SnapshotID uuid.UUID
}

// FPClassifierSettingsGet Get FP classifier settings
func (c *Client) FPClassifierSettingsGet(ctx context.Context) (vampy.FPClassifierSettings, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/settings/", url.PathEscape(c.spaceID.String()))

	var out vampy.FPClassifierSettings
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}

	return out, nil
}

// FPClassifierSettingsPatch Update FP classifier settings
func (c *Client) FPClassifierSettingsPatch(ctx context.Context, opts FPClassifierSettingsPatchOptions) (vampy.FPClassifierSettings, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/settings/", url.PathEscape(c.spaceID.String()))

	var out vampy.FPClassifierSettings
	if err := c.doJSON(ctx, "PATCH", path, nil, opts.Body, &out); err != nil {
		return out, err
	}

	return out, nil
}

type FPClassifierSettingsPatchOptions struct {
	Body vampy.FPClassifierSettings
}

// FPClassifierMetricsDownload Download FP classifier training metrics
func (c *Client) FPClassifierMetricsDownload(ctx context.Context, opts FPClassifierMetricsDownloadOptions) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/metrics/download/", url.PathEscape(c.spaceID.String()))
	q := url.Values{}

	q.Set("limit", strconv.Itoa(opts.Limit))

	return c.doBinary(ctx, "GET", path, q, nil)
}

type FPClassifierMetricsDownloadOptions struct {
	vampy.Paging
}
