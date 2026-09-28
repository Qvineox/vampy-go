package vampy_go

import (
	"context"
	"fmt"
	"github.com/qvineox/vampy-go/vampy"
	"net/url"
	"strconv"
)

// SpaceIssueStatusList Get space issue statuses list
func (c *Client) SpaceIssueStatusList(ctx context.Context, spaceID string, deleted *bool) ([]vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/", url.PathEscape(spaceID))
	q := url.Values{}
	if deleted != nil {
		q.Set("deleted", strconv.FormatBool(*deleted))
	}
	var out []vampy.IssueStatus
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SpaceIssueStatusCreate Create new issue status
func (c *Client) SpaceIssueStatusCreate(ctx context.Context, spaceID string, body vampy.IssueStatusCreate) (vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/", url.PathEscape(spaceID))
	var out vampy.IssueStatus
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SpaceIssueStatusGet Get issue status
func (c *Client) SpaceIssueStatusGet(ctx context.Context, spaceID string, issueStatusName string) (vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/%s/", url.PathEscape(spaceID), url.PathEscape(issueStatusName))
	var out vampy.IssueStatus
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SpaceIssueStatusDelete Delete issue status
func (c *Client) SpaceIssueStatusDelete(ctx context.Context, spaceID string, issueStatusName string) error {
	path := fmt.Sprintf("/spaces/%s/statuses/%s/", url.PathEscape(spaceID), url.PathEscape(issueStatusName))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// SpaceIssueStatusPatch Update existing issue status
func (c *Client) SpaceIssueStatusPatch(ctx context.Context, spaceID string, issueStatusName string, body vampy.IssueStatusUpdate) (vampy.IssueStatus, error) {
	path := fmt.Sprintf("/spaces/%s/statuses/%s/", url.PathEscape(spaceID), url.PathEscape(issueStatusName))
	var out vampy.IssueStatus
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
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
func (c *Client) SystemSettingsPatch(ctx context.Context, body vampy.SystemWideSettings) (vampy.SystemWideSettings, error) {
	path := "/service/settings/"
	var out vampy.SystemWideSettings
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FilterValuesGet Get possible values for given filter targets
func (c *Client) FilterValuesGet(ctx context.Context, spaceID string, target []string, relationType *vampy.FilterMetaRelationType, relationID *string, search *string) (vampy.FilterMetaResponse, error) {
	path := fmt.Sprintf("/service/meta/%s/filter_values/", url.PathEscape(spaceID))
	q := url.Values{}
	for _, v := range target {
		q.Add("target", v)
	}
	if relationType != nil {
		q.Set("relationType", string(*relationType))
	}
	if relationID != nil {
		q.Set("relationID", *relationID)
	}
	if search != nil {
		q.Set("search", *search)
	}
	var out vampy.FilterMetaResponse
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
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
func (c *Client) LDAPConfigPost(ctx context.Context, body vampy.LDAPAuthenticationConfig) (vampy.LDAPAuthenticationConfig, error) {
	path := "/service/ldap/"
	var out vampy.LDAPAuthenticationConfig
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// LDAPConfigDelete Delete LDAP authentication config
func (c *Client) LDAPConfigDelete(ctx context.Context) error {
	path := "/service/ldap/"
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// LDAPConfigPatch Update LDAP authentication config
func (c *Client) LDAPConfigPatch(ctx context.Context, body vampy.LDAPAuthenticationConfig) (vampy.LDAPAuthenticationConfig, error) {
	path := "/service/ldap/"
	var out vampy.LDAPAuthenticationConfig
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// LDAPConfigTest Test LDAP authentication config without creation
func (c *Client) LDAPConfigTest(ctx context.Context, body vampy.LDAPAuthenticationConfig) (vampy.Page[vampy.LDAPUser], error) {
	path := "/service/ldap/test/"
	var out vampy.Page[vampy.LDAPUser]
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
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
func (c *Client) OIDCConfigPost(ctx context.Context, body vampy.OIDCConfig) (vampy.OIDCConfig, error) {
	path := "/service/oidc/"
	var out vampy.OIDCConfig
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// OIDCConfigGet Get existing OpenID connect configs
func (c *Client) OIDCConfigGet(ctx context.Context, configID string) (vampy.OIDCConfig, error) {
	path := fmt.Sprintf("/service/oidc/%s/", url.PathEscape(configID))
	var out vampy.OIDCConfig
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// OIDCConfigDelete Delete existing OpenID connect config
func (c *Client) OIDCConfigDelete(ctx context.Context, configID string) error {
	path := fmt.Sprintf("/service/oidc/%s/", url.PathEscape(configID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// OIDCConfigPatch Patch existing OpenID connect config
func (c *Client) OIDCConfigPatch(ctx context.Context, configID string, body vampy.OIDCConfig) (vampy.OIDCConfig, error) {
	path := fmt.Sprintf("/service/oidc/%s/", url.PathEscape(configID))
	var out vampy.OIDCConfig
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// OIDCConfigDefaults Get default value for OIDC provider attributes
func (c *Client) OIDCConfigDefaults(ctx context.Context, provider vampy.OIDCProvider, url_ *string, realm *string) (vampy.OIDCDefaults, error) {
	path := fmt.Sprintf("/service/oidc/%s/defaults/", url.PathEscape(string(provider)))
	q := url.Values{}
	if url_ != nil {
		q.Set("url", *url_)
	}
	if realm != nil {
		q.Set("realm", *realm)
	}
	var out vampy.OIDCDefaults
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// OIDCTestConnect Test connection to OIDC provider server
func (c *Client) OIDCTestConnect(ctx context.Context, provider vampy.OIDCProvider, body vampy.OIDCTestConnectRequest) ([]vampy.OIDCTestConnectResponse, error) {
	path := fmt.Sprintf("/service/oidc/%s/test_connect/", url.PathEscape(string(provider)))
	var out []vampy.OIDCTestConnectResponse
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
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
func (c *Client) SMTPConfigPost(ctx context.Context, body vampy.SMTPConfig) (vampy.SMTPConfig, error) {
	path := "/service/smtp_config/"
	var out vampy.SMTPConfig
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// SMTPConfigDelete Delete SMTP authentication config
func (c *Client) SMTPConfigDelete(ctx context.Context) error {
	path := "/service/smtp_config/"
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// SMTPConfigPatch Update already an existing SMTP authentication config
func (c *Client) SMTPConfigPatch(ctx context.Context, body vampy.SMTPConfig) (vampy.SMTPConfig, error) {
	path := "/service/smtp_config/"
	var out vampy.SMTPConfig
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

func (c *Client) SendTestEmail(ctx context.Context, body vampy.Email) error {
	path := "/service/smtp_config/test_email/"
	return c.doJSON(ctx, "POST", path, nil, body, nil)
}

// UserRolesGet Get all roles
func (c *Client) UserRolesGet(ctx context.Context, system *bool) ([]vampy.UserRole, error) {
	path := "/service/user_roles/"
	q := url.Values{}
	if system != nil {
		q.Set("system", strconv.FormatBool(*system))
	}
	var out []vampy.UserRole
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// UserRolesCreate Create a new role
func (c *Client) UserRolesCreate(ctx context.Context, body vampy.UserRoleEdit) (vampy.UserRole, error) {
	path := "/service/user_roles/"
	var out vampy.UserRole
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
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
func (c *Client) UserRolesDetailsGet(ctx context.Context, userRoleID string) (vampy.UserRole, error) {
	path := fmt.Sprintf("/service/user_roles/%s/", url.PathEscape(userRoleID))
	var out vampy.UserRole
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// UserRolesDelete Delete user role
func (c *Client) UserRolesDelete(ctx context.Context, userRoleID string) error {
	path := fmt.Sprintf("/service/user_roles/%s/", url.PathEscape(userRoleID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// UserRolesPatch Edit existing user role
func (c *Client) UserRolesPatch(ctx context.Context, userRoleID string, body vampy.UserRoleEdit) error {
	path := fmt.Sprintf("/service/user_roles/%s/", url.PathEscape(userRoleID))
	return c.doJSON(ctx, "PATCH", path, nil, body, nil)
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
func (c *Client) BackgroundTasksListGet(ctx context.Context, offset *int, limit *int) (vampy.Page[vampy.BackgroundTask], error) {
	path := "/service/background_tasks/"
	q := url.Values{}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	var out vampy.Page[vampy.BackgroundTask]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// BackgroundTasksDetailsGet Detailed info about background task
func (c *Client) BackgroundTasksDetailsGet(ctx context.Context, taskID string) (vampy.BackgroundTaskDetails, error) {
	path := fmt.Sprintf("/service/background_tasks/%s/", url.PathEscape(taskID))
	var out vampy.BackgroundTaskDetails
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
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
func (c *Client) PostGlobalSLASetting(ctx context.Context, body vampy.SLASetting) (vampy.SLASetting, error) {
	path := "/service/sla_config/"
	var out vampy.SLASetting
	if err := c.doJSON(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierDatasetDownload Download scan issues dataset
func (c *Client) FPClassifierDatasetDownload(ctx context.Context, spaceID string) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/download/", url.PathEscape(spaceID))
	return c.doBinary(ctx, "GET", path, nil, nil)
}

// FPClassifierDatasetUpload Upload scan issues dataset
func (c *Client) FPClassifierDatasetUpload(ctx context.Context, spaceID string, body vampy.DatasetUploadBody) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/upload/", url.PathEscape(spaceID))
	var out vampy.BackgroundTask
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierDatasetClear Clear scan issues dataset
func (c *Client) FPClassifierDatasetClear(ctx context.Context, spaceID string) error {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/clear/", url.PathEscape(spaceID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// FPClassifierTrainingDataStats Get accumulated training data statistics
func (c *Client) FPClassifierTrainingDataStats(ctx context.Context, spaceID string) ([]vampy.FPClassifierDatasetStats, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/dataset/stats/", url.PathEscape(spaceID))
	var out []vampy.FPClassifierDatasetStats
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierSnapshotList Get FP classifier snapshots list
func (c *Client) FPClassifierSnapshotList(ctx context.Context, spaceID string, scannerType *vampy.ScannerType, offset *int, limit *int, TestTotalCount *int) (vampy.Page[vampy.FPClassifierSnapshot], error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/snapshots/", url.PathEscape(spaceID))
	q := url.Values{}
	if scannerType != nil {
		q.Set("scannerType", string(*scannerType))
	}
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if TestTotalCount != nil {
		q.Set("_test_total_count", strconv.Itoa(*TestTotalCount))
	}
	var out vampy.Page[vampy.FPClassifierSnapshot]
	if err := c.doJSON(ctx, "GET", path, q, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierSnapshotsDownload Download active FP classifier model files archive
func (c *Client) FPClassifierSnapshotsDownload(ctx context.Context, spaceID string) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/snapshots/download/", url.PathEscape(spaceID))
	return c.doBinary(ctx, "GET", path, nil, nil)
}

// FPClassifierSnapshotUpload Upload FP classifier model archive
func (c *Client) FPClassifierSnapshotUpload(ctx context.Context, spaceID string, body vampy.SnapshotsUploadBody) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/snapshots/upload/", url.PathEscape(spaceID))
	var out vampy.BackgroundTask
	if err := c.doMultipart(ctx, "POST", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierSnapshotGet Get FP classifier snapshot
func (c *Client) FPClassifierSnapshotGet(ctx context.Context, snapshotID string) (vampy.FPClassifierSnapshot, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/", url.PathEscape(snapshotID))
	var out vampy.FPClassifierSnapshot
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierSnapshotDelete Delete FP classifier snapshot
func (c *Client) FPClassifierSnapshotDelete(ctx context.Context, snapshotID string) error {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/", url.PathEscape(snapshotID))
	return c.doJSON(ctx, "DELETE", path, nil, nil, nil)
}

// FPClassifierSnapshotPatch Update FP classifier snapshot
func (c *Client) FPClassifierSnapshotPatch(ctx context.Context, snapshotID string, body vampy.SnapshotsSnapshotIDBody) (vampy.FPClassifierSnapshot, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/", url.PathEscape(snapshotID))
	var out vampy.FPClassifierSnapshot
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierSnapshotTrain Train FP classifier snapshot
func (c *Client) FPClassifierSnapshotTrain(ctx context.Context, snapshotID string) (vampy.BackgroundTask, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/train/", url.PathEscape(snapshotID))
	var out vampy.BackgroundTask
	if err := c.doJSON(ctx, "POST", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierSnapshotDownload Download FP classifier model archive
func (c *Client) FPClassifierSnapshotDownload(ctx context.Context, snapshotID string) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/download/", url.PathEscape(snapshotID))
	return c.doBinary(ctx, "GET", path, nil, nil)
}

// FPClassifierSnapshotDatasetDownload Download FP classifier snapshot dataset
func (c *Client) FPClassifierSnapshotDatasetDownload(ctx context.Context, snapshotID string) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/snapshots/%s/dataset/download/", url.PathEscape(snapshotID))
	return c.doBinary(ctx, "GET", path, nil, nil)
}

// FPClassifierSettingsGet Get FP classifier settings
func (c *Client) FPClassifierSettingsGet(ctx context.Context, spaceID string) (vampy.FPClassifierSettings, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/settings/", url.PathEscape(spaceID))
	var out vampy.FPClassifierSettings
	if err := c.doJSON(ctx, "GET", path, nil, nil, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierSettingsPatch Update FP classifier settings
func (c *Client) FPClassifierSettingsPatch(ctx context.Context, spaceID string, body vampy.FPClassifierSettings) (vampy.FPClassifierSettings, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/settings/", url.PathEscape(spaceID))
	var out vampy.FPClassifierSettings
	if err := c.doJSON(ctx, "PATCH", path, nil, body, &out); err != nil {
		return out, err
	}
	return out, nil
}

// FPClassifierMetricsDownload Download FP classifier training metrics
func (c *Client) FPClassifierMetricsDownload(ctx context.Context, spaceID string, limit *int) ([]byte, error) {
	path := fmt.Sprintf("/service/fp_classifier/%s/metrics/download/", url.PathEscape(spaceID))
	q := url.Values{}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	return c.doBinary(ctx, "GET", path, q, nil)
}
