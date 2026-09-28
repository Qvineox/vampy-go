// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type BitBucketWebhookRequest struct {
	EventKey string `json:"eventKey"`
}

type BitbucketApplicationInfoRequest struct {
	Hostname       string `json:"hostname,omitempty"`
	IgnoreSSLCheck *bool  `json:"ignoreSSLCheck,omitempty"`
}

type BitbucketApplicationInfoResponse struct {
	Authentication string `json:"authentication,omitempty"`
}

type DefectDojoFindingStatus string

const (
	DefectDojoFindingStatusActive       DefectDojoFindingStatus = "active"
	DefectDojoFindingStatusVerified     DefectDojoFindingStatus = "verified"
	DefectDojoFindingStatusFalseP       DefectDojoFindingStatus = "false_p"
	DefectDojoFindingStatusIsMitigated  DefectDojoFindingStatus = "is_mitigated"
	DefectDojoFindingStatusOutOfScope   DefectDojoFindingStatus = "out_of_scope"
	DefectDojoFindingStatusRiskAccepted DefectDojoFindingStatus = "risk_accepted"
	DefectDojoFindingStatusUnderReview  DefectDojoFindingStatus = "under_review"
)

type DefectDojoImportRequest struct {
	SpaceID         UUID                      `json:"spaceID,omitempty"`
	Timeout         int                       `json:"timeout,omitempty"`
	BatchSize       int                       `json:"batchSize,omitempty"`
	ProductIDs      []int                     `json:"productIDs,omitempty"`
	FindingStatuses []DefectDojoFindingStatus `json:"findingStatuses,omitempty"`
	CreateProducts  bool                      `json:"createProducts,omitempty"`
	CreateComments  bool                      `json:"createComments,omitempty"`
}

type DefectDojoImportResponse struct {
	TaskID           UUID `json:"taskID,omitempty"`
	ProductsCount    int  `json:"productsCount,omitempty"`
	EngagementsCount int  `json:"engagementsCount,omitempty"`
	FindingsCount    int  `json:"findingsCount,omitempty"`
}

type DefectDojoProduct struct {
	ID                  int              `json:"id"`
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	Platform            *string          `json:"platform,omitempty"`
	BusinessCriticality *string          `json:"businessCriticality,omitempty"`
	Tags                []string         `json:"tags"`
	ProductMeta         []map[string]any `json:"productMeta"`
}

type DefectDojoStatusMapping struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	VampyStatus string `json:"vampyStatus"`
}

type ExternalArtifact struct {
	ID           string `json:"id,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	Description  string `json:"description,omitempty"`
	VersionCount int    `json:"versionCount,omitempty"`
}

type ExternalRepository struct {
	ID                string                     `json:"id,omitempty"`
	Name              string                     `json:"name,omitempty"`
	FullName          string                     `json:"fullName,omitempty"`
	Description       string                     `json:"description,omitempty"`
	Slug              string                     `json:"slug,omitempty"`
	URL               string                     `json:"url,omitempty"`
	DefaultBranch     string                     `json:"defaultBranch,omitempty"`
	AvatarURL         string                     `json:"avatarUrl,omitempty"`
	AlreadyLinkedRepo *RepositoryForExternalLink `json:"alreadyLinkedRepo,omitempty"`
	WillBeLinkedRepo  *RepositoryForExternalLink `json:"willBeLinkedRepo,omitempty"`
	CanRead           bool                       `json:"canRead,omitempty"`
	CanImport         bool                       `json:"canImport,omitempty"`
	RelationStatus    string                     `json:"relationStatus,omitempty"`
}

type GitLabWebhookRequest struct {
	EventName string `json:"event_name"`
}

type IntegrationAILanguage string

const (
	IntegrationAILanguageEN IntegrationAILanguage = "EN"
	IntegrationAILanguageRU IntegrationAILanguage = "RU"
)

type IntegrationAIModel struct {
	ID          string                         `json:"id,omitempty"`
	IsDefault   bool                           `json:"isDefault,omitempty"`
	Integration *IntegrationAIModelIntegration `json:"integration,omitempty"`
}

type IntegrationAIModelIntegration struct {
	ID                    UUID             `json:"id,omitempty"`
	Name                  string           `json:"name,omitempty"`
	State                 IntegrationState `json:"state,omitempty"`
	Vendor                string           `json:"vendor,omitempty"`
	IsDefault             bool             `json:"isDefault,omitempty"`
	DefaultModel          string           `json:"defaultModel,omitempty"`
	DefaultModelAvailable bool             `json:"defaultModelAvailable,omitempty"`
}

type IntegrationAIRequestType string

const (
	IntegrationAIRequestTypeSCANISSUEFALSEPOSITIVERATE IntegrationAIRequestType = "SCAN_ISSUE_FALSE_POSITIVE_RATE"
	IntegrationAIRequestTypeSCANISSUEDESCRIPTION       IntegrationAIRequestType = "SCAN_ISSUE_DESCRIPTION"
	IntegrationAIRequestTypeSCANISSUESOLUTION          IntegrationAIRequestType = "SCAN_ISSUE_SOLUTION"
	IntegrationAIRequestTypeSCANISSUECODEEXPLAIN       IntegrationAIRequestType = "SCAN_ISSUE_CODE_EXPLAIN"
	IntegrationAIRequestTypeSCANISSUECODEVULN          IntegrationAIRequestType = "SCAN_ISSUE_CODE_VULN"
)

type IntegrationAI struct {
	ID              UUID                       `json:"id,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Hostname        string                     `json:"hostname,omitempty"`
	Vendor          string                     `json:"vendor,omitempty"`
	Token           string                     `json:"token,omitempty"`
	Service         IntegrationServiceType     `json:"service,omitempty"`
	Roles           []IntegrationRole          `json:"roles,omitempty"`
	State           IntegrationState           `json:"state,omitempty"`
	Created         string                     `json:"created,omitempty"`
	IgnoreSSLCheck  *bool                      `json:"ignoreSSLCheck,omitempty"`
	ServerCert      *string                    `json:"serverCert,omitempty"`
	IsDefault       bool                       `json:"isDefault,omitempty"`
	DefaultModel    string                     `json:"defaultModel,omitempty"`
	DefaultLanguage IntegrationAILanguage      `json:"defaultLanguage,omitempty"`
	VendorParams    []IntegrationAIVendorParam `json:"vendorParams,omitempty"`
	Models          *IntegrationAIModels       `json:"models,omitempty"`
}

type IntegrationAIModels struct {
	Updated string   `json:"updated,omitempty"`
	Items   []string `json:"items,omitempty"`
}

type IntegrationAIVendorParamInfo struct {
	Type        IntegrationAIVendorParamType `json:"type,omitempty"`
	Name        string                       `json:"name,omitempty"`
	Description string                       `json:"description,omitempty"`
	Default     string                       `json:"default,omitempty"`
	Required    bool                         `json:"required,omitempty"`
	HelpURL     string                       `json:"helpURL,omitempty"`
}

type IntegrationAIVendorParam struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type IntegrationAIVendorParamType string

const (
	IntegrationAIVendorParamTypeSTRING IntegrationAIVendorParamType = "STRING"
)

type IntegrationAIVendor struct {
	Slug   string                         `json:"slug,omitempty"`
	Name   string                         `json:"name,omitempty"`
	URL    *string                        `json:"url,omitempty"`
	Params []IntegrationAIVendorParamInfo `json:"params,omitempty"`
}

type IntegrationAppScreener struct {
	ScannerIntegration
}

type IntegrationArtifact struct {
	ID             UUID                           `json:"id,omitempty"`
	Name           string                         `json:"name,omitempty"`
	Hostname       string                         `json:"hostname,omitempty"`
	Username       string                         `json:"username,omitempty"`
	Token          string                         `json:"token,omitempty"`
	Webhook        string                         `json:"webhook,omitempty"`
	Service        IntegrationArtifactServiceType `json:"service,omitempty"`
	Roles          []IntegrationRole              `json:"roles,omitempty"`
	State          IntegrationState               `json:"state,omitempty"`
	ServerID       UUID                           `json:"serverID,omitempty"`
	Created        string                         `json:"created,omitempty"`
	IgnoreSSLCheck bool                           `json:"ignoreSSLCheck,omitempty"`
}

type IntegrationArtifactServiceType string

const (
	IntegrationArtifactServiceTypeHARBOR IntegrationArtifactServiceType = "HARBOR"
	IntegrationArtifactServiceTypeNEXUS  IntegrationArtifactServiceType = "NEXUS"
	IntegrationArtifactServiceTypeOTHER  IntegrationArtifactServiceType = "OTHER"
)

type IntegrationDefectDojo struct {
	ID             UUID                      `json:"id,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Hostname       string                    `json:"hostname,omitempty"`
	Token          string                    `json:"token,omitempty"`
	Service        IntegrationVCSServiceType `json:"service,omitempty"`
	Roles          []IntegrationRole         `json:"roles,omitempty"`
	State          IntegrationState          `json:"state,omitempty"`
	ServerID       UUID                      `json:"serverID,omitempty"`
	Created        string                    `json:"created,omitempty"`
	IgnoreSSLCheck bool                      `json:"ignoreSSLCheck,omitempty"`
}

type IntegrationKTrackerBoard struct {
	ID                int    `json:"id,omitempty"`
	Title             string `json:"title,omitempty"`
	DefaultCardTypeID int    `json:"defaultCardTypeID,omitempty"`
}

type IntegrationKTrackerCardType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type IntegrationKTracker struct {
	ID             UUID                         `json:"id,omitempty"`
	Name           string                       `json:"name,omitempty"`
	Hostname       string                       `json:"hostname,omitempty"`
	Token          string                       `json:"token,omitempty"`
	Webhook        string                       `json:"webhook,omitempty"`
	Service        IntegrationVCSServiceType    `json:"service,omitempty"`
	Roles          []IntegrationRole            `json:"roles,omitempty"`
	State          IntegrationState             `json:"state,omitempty"`
	ServerID       UUID                         `json:"serverID,omitempty"`
	Created        string                       `json:"created,omitempty"`
	IgnoreSSLCheck *bool                        `json:"ignoreSSLCheck,omitempty"`
	TargetSpace    *IntegrationKTrackerSpace    `json:"targetSpace,omitempty"`
	TargetBoard    *IntegrationKTrackerBoard    `json:"targetBoard,omitempty"`
	TargetCardType *IntegrationKTrackerCardType `json:"targetCardType,omitempty"`
}

type IntegrationKTrackerSpace struct {
	ID               int                        `json:"id,omitempty"`
	Title            string                     `json:"title,omitempty"`
	Archived         bool                       `json:"archived,omitempty"`
	Boards           []IntegrationKTrackerBoard `json:"boards,omitempty"`
	AllowedCardTypes []int                      `json:"allowedCardTypes,omitempty"`
}

type IntegrationPTAI struct {
	ScannerIntegration
}

type IntegrationRole string

const (
	IntegrationRoleTASKTRACKER      IntegrationRole = "TASK_TRACKER"
	IntegrationRoleVCS              IntegrationRole = "VCS"
	IntegrationRoleARTIFACTREGISTRY IntegrationRole = "ARTIFACT_REGISTRY"
	IntegrationRoleSCANNER          IntegrationRole = "SCANNER"
	IntegrationRoleIMPORT           IntegrationRole = "IMPORT"
	IntegrationRoleJIRA             IntegrationRole = "JIRA"
	IntegrationRoleAI               IntegrationRole = "AI"
)

type IntegrationSastav struct {
	ScannerIntegration
	Username string `json:"username"`
}

type IntegrationScannerServiceType string

const (
	IntegrationScannerServiceTypePTAI        IntegrationScannerServiceType = "PTAI"
	IntegrationScannerServiceTypeCODESCORING IntegrationScannerServiceType = "CODE_SCORING"
	IntegrationScannerServiceTypeAPPSCREENER IntegrationScannerServiceType = "APP_SCREENER"
	IntegrationScannerServiceTypeSASTAV      IntegrationScannerServiceType = "SASTAV"
	IntegrationScannerServiceTypeAPPSECTRACK IntegrationScannerServiceType = "APPSEC_TRACK"
)

type IntegrationServiceType string

const (
	IntegrationServiceTypeJIRA              IntegrationServiceType = "JIRA"
	IntegrationServiceTypeGITLAB            IntegrationServiceType = "GITLAB"
	IntegrationServiceTypeBITBUCKET         IntegrationServiceType = "BITBUCKET"
	IntegrationServiceTypeAZUREDEVOPS       IntegrationServiceType = "AZURE_DEVOPS"
	IntegrationServiceTypeAZUREDEVOPSBOARDS IntegrationServiceType = "AZURE_DEVOPS_BOARDS"
	IntegrationServiceTypeGITEA             IntegrationServiceType = "GITEA"
	IntegrationServiceTypeGITFLIC           IntegrationServiceType = "GITFLIC"
	IntegrationServiceTypeGITFLAME          IntegrationServiceType = "GITFLAME"
	IntegrationServiceTypeEVATEAM           IntegrationServiceType = "EVA_TEAM"
	IntegrationServiceTypeKTRACKER          IntegrationServiceType = "KTRACKER"
	IntegrationServiceTypePTAI              IntegrationServiceType = "PTAI"
	IntegrationServiceTypeHARBOR            IntegrationServiceType = "HARBOR"
	IntegrationServiceTypeNEXUS             IntegrationServiceType = "NEXUS"
	IntegrationServiceTypeOTHER             IntegrationServiceType = "OTHER"
	IntegrationServiceTypeDEFECTDOJO        IntegrationServiceType = "DEFECT_DOJO"
	IntegrationServiceTypeCODESCORING       IntegrationServiceType = "CODE_SCORING"
	IntegrationServiceTypeAPPSCREENER       IntegrationServiceType = "APP_SCREENER"
	IntegrationServiceTypeSASTAV            IntegrationServiceType = "SASTAV"
	IntegrationServiceTypeAPPSECTRACK       IntegrationServiceType = "APPSEC_TRACK"
	IntegrationServiceTypeAI                IntegrationServiceType = "AI"
)

type IntegrationState string

const (
	IntegrationStateCREATED       IntegrationState = "CREATED"
	IntegrationStateAUTHENTICATED IntegrationState = "AUTHENTICATED"
)

type IntegrationVCS struct {
	ID             UUID                      `json:"id,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Hostname       string                    `json:"hostname,omitempty"`
	Token          string                    `json:"token,omitempty"`
	Username       *string                   `json:"username,omitempty"`
	Webhook        string                    `json:"webhook,omitempty"`
	Service        IntegrationVCSServiceType `json:"service,omitempty"`
	Roles          []IntegrationRole         `json:"roles,omitempty"`
	State          IntegrationState          `json:"state,omitempty"`
	ServerID       UUID                      `json:"serverID,omitempty"`
	Created        string                    `json:"created,omitempty"`
	IgnoreSSLCheck *bool                     `json:"ignoreSSLCheck,omitempty"`
	ServerCert     *string                   `json:"serverCert,omitempty"`
}

type IntegrationVCSServiceType string

const (
	IntegrationVCSServiceTypeGITLAB      IntegrationVCSServiceType = "GITLAB"
	IntegrationVCSServiceTypeBITBUCKET   IntegrationVCSServiceType = "BITBUCKET"
	IntegrationVCSServiceTypeAZUREDEVOPS IntegrationVCSServiceType = "AZURE_DEVOPS"
	IntegrationVCSServiceTypeGITEA       IntegrationVCSServiceType = "GITEA"
	IntegrationVCSServiceTypeGITFLIC     IntegrationVCSServiceType = "GITFLIC"
	IntegrationVCSServiceTypeGITFLAME    IntegrationVCSServiceType = "GITFLAME"
)

type JiraApplicationInfoRequest struct {
	Hostname       string `json:"hostname,omitempty"`
	IgnoreSSLCheck *bool  `json:"ignoreSSLCheck,omitempty"`
}

type JiraApplicationInfoResponse struct {
	Authentication string `json:"authentication,omitempty"`
}

type JiraIntegration struct {
	ServerID              UUID                                   `json:"serverID,omitempty"`
	Name                  string                                 `json:"name,omitempty"`
	Service               IntegrationServiceType                 `json:"service,omitempty"`
	Roles                 []IntegrationRole                      `json:"roles,omitempty"`
	Username              string                                 `json:"username,omitempty"`
	Token                 *string                                `json:"token,omitempty"`
	Hostname              string                                 `json:"hostname,omitempty"`
	JiraProjectKey        string                                 `json:"jiraProjectKey,omitempty"`
	JiraIssueType         *JiraIntegrationJiraIssueType          `json:"jiraIssueType,omitempty"`
	BackwardSyncEnabled   *bool                                  `json:"backwardSyncEnabled,omitempty"`
	SyncCommentsEnabled   bool                                   `json:"syncCommentsEnabled,omitempty"`
	BackwardStatusMapping []JiraIntegrationBackwardStatusMapping `json:"backwardStatusMapping,omitempty"`
	FieldsDefaultValues   []JiraIntegrationFieldsDefaultValues   `json:"fieldsDefaultValues,omitempty"`
	BotUserID             *UUID                                  `json:"botUserID,omitempty"`
	Webhook               *string                                `json:"webhook,omitempty"`
	IgnoreSSLCheck        *bool                                  `json:"ignoreSSLCheck,omitempty"`
}

type JiraIntegrationBackwardStatusMapping struct {
	LocalStatus      string `json:"localStatus"`
	ExternalStatusID string `json:"externalStatusID"`
}

type JiraIntegrationFieldsDefaultValues struct {
	Field string   `json:"field"`
	Value AnyValue `json:"value"`
}

type JiraIntegrationJiraIssueType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type JiraMetaFields struct {
	Name          string           `json:"name,omitempty"`
	DisplayName   string           `json:"displayName,omitempty"`
	Type          string           `json:"type,omitempty"`
	Items         string           `json:"items,omitempty"`
	AllowedValues []map[string]any `json:"allowedValues,omitempty"`
	DefaultValue  map[string]any   `json:"defaultValue,omitempty"`
	FillFrom      string           `json:"fillFrom,omitempty"`
	HelpText      *string          `json:"helpText,omitempty"`
	SearchField   *string          `json:"searchField,omitempty"`
}

type JiraMetaInfoResponse struct {
	Projects          []JiraMetaInfoResponseProjects         `json:"projects,omitempty"`
	ExternalStatuses  []JiraMetaInfoResponseExternalStatuses `json:"externalStatuses,omitempty"`
	LocalStatuses     []IssueStatus                          `json:"localStatuses,omitempty"`
	RequiredFields    []JiraMetaFields                       `json:"requiredFields,omitempty"`
	NonRequiredFields []JiraMetaFields                       `json:"nonRequiredFields,omitempty"`
}

type JiraMetaInfoResponseExternalStatuses struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type JiraMetaInfoResponseIssueTypes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
}

type JiraMetaInfoResponseProjects struct {
	Key        string                           `json:"key"`
	Name       string                           `json:"name"`
	IssueTypes []JiraMetaInfoResponseIssueTypes `json:"issuetypes,omitempty"`
}

type JiraSearchIssuesResponse struct {
	ID            string `json:"id,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
	Key           string `json:"key,omitempty"`
	Summary       string `json:"summary,omitempty"`
	AlreadyLinked bool   `json:"alreadyLinked,omitempty"`
}

type JiraSearchResponse struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type JiraWebhookRequest struct {
	Issue              *JiraWebhookRequestIssue   `json:"issue,omitempty"`
	WebhookEvent       string                     `json:"webhookEvent,omitempty"`
	IssueEventTypeName string                     `json:"issue_event_type_name,omitempty"`
	Comment            *JiraWebhookRequestComment `json:"comment,omitempty"`
}

type JiraWebhookRequestComment struct {
	ID     string                           `json:"id,omitempty"`
	Body   string                           `json:"body,omitempty"`
	Author *JiraWebhookRequestCommentAuthor `json:"author,omitempty"`
}

type JiraWebhookRequestCommentAuthor struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Key          string `json:"key,omitempty"`
	Active       bool   `json:"active,omitempty"`
	AccountID    string `json:"accountId,omitempty"`
}

type JiraWebhookRequestIssue struct {
	ID     string                         `json:"id,omitempty"`
	Key    string                         `json:"key"`
	Fields *JiraWebhookRequestIssueFields `json:"fields,omitempty"`
}

type JiraWebhookRequestIssueFields struct {
	Status *JiraWebhookRequestIssueFieldsStatus `json:"status,omitempty"`
}

type JiraWebhookRequestIssueFieldsStatus struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type TaskTrackerAsyncExport struct {
	ExportMode          string                                 `json:"exportMode"`
	SummarizedIssue     *TaskTrackerAsyncExportSummarizedIssue `json:"summarizedIssue,omitempty"`
	Issues              []TaskTrackerAsyncExportIssues         `json:"issues,omitempty"`
	FieldsDefaultValues []JiraIntegrationFieldsDefaultValues   `json:"fieldsDefaultValues,omitempty"`
}

type TaskTrackerAsyncExportIssues struct {
	ID           UUID                                                `json:"id"`
	Summary      string                                              `json:"summary"`
	FieldsValues []TaskTrackerAsyncExportSummarizedIssueFieldsValues `json:"fieldsValues,omitempty"`
}

type TaskTrackerAsyncExportSummarizedIssue struct {
	IDs          []UUID                                              `json:"ids,omitempty"`
	Summary      string                                              `json:"summary,omitempty"`
	FieldsValues []TaskTrackerAsyncExportSummarizedIssueFieldsValues `json:"fieldsValues,omitempty"`
}

type TaskTrackerAsyncExportSummarizedIssueFieldsValues struct {
	Field string   `json:"field"`
	Value AnyValue `json:"value"`
}

type TaskTrackerExportMetaRequest struct {
	ScanIssueIDs []UUID `json:"scanIssueIDs"`
}

type TaskTrackerInfoRequest struct {
	Hostname       string  `json:"hostname,omitempty"`
	IgnoreSSLCheck *bool   `json:"ignoreSSLCheck,omitempty"`
	ServerCert     *string `json:"serverCert,omitempty"`
}

type TaskTrackerInfoResponse struct {
	Authentication string `json:"authentication,omitempty"`
	Version        string `json:"version,omitempty"`
}

type TaskTrackerIntegration struct {
	ServerID              UUID                                          `json:"serverID,omitempty"`
	Name                  string                                        `json:"name,omitempty"`
	Service               IntegrationServiceType                        `json:"service,omitempty"`
	Roles                 []IntegrationRole                             `json:"roles,omitempty"`
	Username              string                                        `json:"username,omitempty"`
	Token                 *string                                       `json:"token,omitempty"`
	Hostname              string                                        `json:"hostname,omitempty"`
	Project               *TaskTrackerIntegrationProject                `json:"project,omitempty"`
	Board                 *TaskTrackerIntegrationBoard                  `json:"board,omitempty"`
	IssueType             *TaskTrackerIntegrationIssueType              `json:"issueType,omitempty"`
	BackwardSyncEnabled   *bool                                         `json:"backwardSyncEnabled,omitempty"`
	SyncCommentsEnabled   bool                                          `json:"syncCommentsEnabled,omitempty"`
	BackwardStatusMapping []TaskTrackerIntegrationBackwardStatusMapping `json:"backwardStatusMapping,omitempty"`
	FieldsDefaultValues   []JiraIntegrationFieldsDefaultValues          `json:"fieldsDefaultValues,omitempty"`
	BotUserID             *UUID                                         `json:"botUserID,omitempty"`
	Webhook               *string                                       `json:"webhook,omitempty"`
	IgnoreSSLCheck        *bool                                         `json:"ignoreSSLCheck,omitempty"`
	ServerCert            *string                                       `json:"serverCert,omitempty"`
}

type TaskTrackerIntegrationBackwardStatusMapping struct {
	LocalStatus      string `json:"localStatus"`
	ExternalStatusID string `json:"externalStatusID"`
}

type TaskTrackerIntegrationBoard struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type TaskTrackerIntegrationIssueType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TaskTrackerIntegrationProject struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type TaskTrackerIntegrationType string

const (
	TaskTrackerIntegrationTypeAzureDevopsBoards TaskTrackerIntegrationType = "azure_devops_boards"
	TaskTrackerIntegrationTypeEvaTeam           TaskTrackerIntegrationType = "eva_team"
	TaskTrackerIntegrationTypeJira              TaskTrackerIntegrationType = "jira"
	TaskTrackerIntegrationTypeKtracker          TaskTrackerIntegrationType = "ktracker"
)

type TaskTrackerIssueExport struct {
	Key  string  `json:"key"`
	URL  *string `json:"url,omitempty"`
	Name *string `json:"name,omitempty"`
}

type TaskTrackerLinkRequest struct {
	ScanIssueIDs     []UUID `json:"scanIssueIDs"`
	ExternalIssueKey string `json:"externalIssueKey"`
}

type TaskTrackerMetaInfoResponse struct {
	Projects          []TaskTrackerMetaInfoResponseProjects         `json:"projects,omitempty"`
	ExternalStatuses  []TaskTrackerMetaInfoResponseExternalStatuses `json:"externalStatuses,omitempty"`
	LocalStatuses     []IssueStatus                                 `json:"localStatuses,omitempty"`
	RequiredFields    []JiraMetaFields                              `json:"requiredFields,omitempty"`
	NonRequiredFields []JiraMetaFields                              `json:"nonRequiredFields,omitempty"`
}

type TaskTrackerMetaInfoResponseExternalStatuses struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Color    *string `json:"color,omitempty"`
	Category *string `json:"category,omitempty"`
}

type TaskTrackerMetaInfoResponseIssueTypes struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	IconURL     *string `json:"iconUrl,omitempty"`
}

type TaskTrackerMetaInfoResponseProjects struct {
	ID         string                                  `json:"id,omitempty"`
	Key        *string                                 `json:"key,omitempty"`
	Name       string                                  `json:"name"`
	IssueTypes []TaskTrackerMetaInfoResponseIssueTypes `json:"issueTypes,omitempty"`
}

type TaskTrackerSearchResponse struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type TaskTrackerServiceType string

const (
	TaskTrackerServiceTypeJira              TaskTrackerServiceType = "jira"
	TaskTrackerServiceTypeKtracker          TaskTrackerServiceType = "ktracker"
	TaskTrackerServiceTypeAzureDevopsBoards TaskTrackerServiceType = "azure_devops_boards"
	TaskTrackerServiceTypeEvaTeam           TaskTrackerServiceType = "eva_team"
)

type VCSAuthenticationType string

const (
	VCSAuthenticationTypeTOKEN VCSAuthenticationType = "TOKEN"
	VCSAuthenticationTypeCLOUD VCSAuthenticationType = "CLOUD"
	VCSAuthenticationTypeBASIC VCSAuthenticationType = "BASIC"
)

type VCSIntegrationFeature string

const (
	VCSIntegrationFeatureORCHESTRATIONSCAN   VCSIntegrationFeature = "ORCHESTRATION_SCAN"
	VCSIntegrationFeatureGITBLAME            VCSIntegrationFeature = "GIT_BLAME"
	VCSIntegrationFeatureWHREPOSITORYCHANGES VCSIntegrationFeature = "WH_REPOSITORY_CHANGES"
)

type VCSIntegrationInfoRequest struct {
	Hostname       string  `json:"hostname,omitempty"`
	Token          string  `json:"token,omitempty"`
	IgnoreSSLCheck bool    `json:"ignoreSSLCheck,omitempty"`
	ServerCert     *string `json:"serverCert,omitempty"`
}

type VCSIntegrationInfoResponse struct {
	Authentication      VCSAuthenticationType   `json:"authentication"`
	UnsupportedFeatures []VCSIntegrationFeature `json:"unsupportedFeatures"`
}

type VCSIntegration string

const (
	VCSIntegrationGitlab      VCSIntegration = "gitlab"
	VCSIntegrationBitbucket   VCSIntegration = "bitbucket"
	VCSIntegrationAzureDevops VCSIntegration = "azure_devops"
	VCSIntegrationGitea       VCSIntegration = "gitea"
	VCSIntegrationGitflic     VCSIntegration = "gitflic"
	VCSIntegrationGitflame    VCSIntegration = "gitflame"
)

type IntegrationIDUploadCLIBody struct {
	File []byte `json:"file"`
}

type IntegrationIDUploadCLIBody1 struct {
	File []byte `json:"file"`
}

type IntegrationIDUploadCLIBody2 struct {
	File []byte `json:"file"`
}
