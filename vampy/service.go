// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type AvailabilityStatus struct {
	Applicable bool                                  `json:"applicable,omitempty"`
	ReasonCode OrchestrationScanUnavailabilityReason `json:"reasonCode,omitempty"`
}

type BackgroundTaskDetails struct {
	BackgroundTask
	SubTasks []BackgroundTaskDetailsSubTasks `json:"subTasks,omitempty"`
}

type BackgroundTaskDetailsSubTasks struct {
	TaskID      UUID                 `json:"taskID,omitempty"`
	Status      BackgroundTaskStatus `json:"status,omitempty"`
	DisplayName string               `json:"displayName,omitempty"`
	Created     string               `json:"created,omitempty"`
	Updated     string               `json:"updated,omitempty"`
}

type BackgroundTaskMetaRulesCheck struct {
	PassedInstances int `json:"passedInstances,omitempty"`
}

type BackgroundTaskMeta struct {
	TaskTracker *BackgroundTaskMetaTaskTracker `json:"taskTracker,omitempty"`
	RulesCheck  *BackgroundTaskMetaRulesCheck  `json:"rulesCheck,omitempty"`
}

type BackgroundTaskMetaTaskTracker struct {
	CompletedIssues []UUID `json:"completedIssues,omitempty"`
	FailedIssues    []UUID `json:"failedIssues,omitempty"`
}

type BackgroundTask struct {
	TaskID            UUID                   `json:"taskID,omitempty"`
	Status            BackgroundTaskStatus   `json:"status,omitempty"`
	DisplayName       string                 `json:"displayName,omitempty"`
	Options           *BackgroundTaskOptions `json:"options,omitempty"`
	TotalSubTasks     int                    `json:"totalSubTasks,omitempty"`
	CompletedSubTasks int                    `json:"completedSubTasks,omitempty"`
	Created           string                 `json:"created,omitempty"`
	Updated           string                 `json:"updated,omitempty"`
	ErrorDetails      *string                `json:"errorDetails,omitempty"`
	Meta              *BackgroundTaskMeta    `json:"meta,omitempty"`
}

type BackgroundTaskOptions struct {
	Parser Parser `json:"parser,omitempty"`
}

type BackgroundTaskStatus string

const (
	BackgroundTaskStatusPENDING  BackgroundTaskStatus = "PENDING"
	BackgroundTaskStatusPROGRESS BackgroundTaskStatus = "PROGRESS"
	BackgroundTaskStatusSUCCESS  BackgroundTaskStatus = "SUCCESS"
	BackgroundTaskStatusFAILED   BackgroundTaskStatus = "FAILED"
	BackgroundTaskStatusRETRYING BackgroundTaskStatus = "RETRYING"
	BackgroundTaskStatusSKIP     BackgroundTaskStatus = "SKIP"
	BackgroundTaskStatusLOCKED   BackgroundTaskStatus = "LOCKED"
)

type CSVExportDelimiter string

const (
	CSVExportDelimiterSEMICOLON CSVExportDelimiter = "SEMICOLON"
	CSVExportDelimiterCOMMA     CSVExportDelimiter = "COMMA"
	CSVExportDelimiterTAB       CSVExportDelimiter = "TAB"
	CSVExportDelimiterSPACE     CSVExportDelimiter = "SPACE"
)

type CSVExportDialect string

const (
	CSVExportDialectEXCEL      CSVExportDialect = "EXCEL"
	CSVExportDialectGOOGLEDOCS CSVExportDialect = "GOOGLE_DOCS"
)

type CSVExportDoubleQuote bool

type CSVExportEncoding string

const (
	CSVExportEncodingUTF8 CSVExportEncoding = "UTF-8"
)

type CSVExportEscapeChar string

type CSVExportLineTerminator string

type CSVExportMeta struct {
	TotalCount         int                 `json:"totalCount,omitempty"`
	MaxExportCount     int                 `json:"maxExportCount,omitempty"`
	SupportedEncodings []CSVExportEncoding `json:"supportedEncodings,omitempty"`
	DialectsSettings   []CSVExportSettings `json:"dialectsSettings,omitempty"`
}

type CSVExportQuoteChar string

type CSVExportQuoting string

const (
	CSVExportQuotingQUOTEMINIMAL    CSVExportQuoting = "QUOTE_MINIMAL"
	CSVExportQuotingQUOTEALL        CSVExportQuoting = "QUOTE_ALL"
	CSVExportQuotingQUOTENONNUMERIC CSVExportQuoting = "QUOTE_NONNUMERIC"
	CSVExportQuotingQUOTENONE       CSVExportQuoting = "QUOTE_NONE"
)

type CSVExportSettings struct {
	Dialect            CSVExportDialect        `json:"dialect,omitempty"`
	DialectDisplayName string                  `json:"dialectDisplayName,omitempty"`
	Encoding           CSVExportEncoding       `json:"encoding,omitempty"`
	Delimiter          CSVExportDelimiter      `json:"delimiter,omitempty"`
	DoubleQuote        CSVExportDoubleQuote    `json:"doubleQuote,omitempty"`
	LineTerminator     CSVExportLineTerminator `json:"lineTerminator,omitempty"`
	QuoteChar          CSVExportQuoteChar      `json:"quoteChar,omitempty"`
	EscapeChar         CSVExportEscapeChar     `json:"escapeChar,omitempty"`
	QuotingFields      CSVExportQuoting        `json:"quotingFields,omitempty"`
}

type Email struct {
	ToAddress string `json:"toAddress"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

type ExportFormatter string

const (
	ExportFormatterJSON ExportFormatter = "JSON"
)

type ExportPayload string

const (
	ExportPayloadINLINE     ExportPayload = "INLINE"
	ExportPayloadATTACHMENT ExportPayload = "ATTACHMENT"
	ExportPayloadARCHIVE    ExportPayload = "ARCHIVE"
)

type FPClassifierDatasetCounters struct {
	TotalCount         int `json:"totalCount,omitempty"`
	FPCount            int `json:"fpCount,omitempty"`
	TpCount            int `json:"tpCount,omitempty"`
	RequiredTotalCount int `json:"requiredTotalCount,omitempty"`
	RequiredFPCount    int `json:"requiredFPCount,omitempty"`
	RequiredTPCount    int `json:"requiredTPCount,omitempty"`
}

type FPClassifierDatasetStats struct {
	ScannerType  ScannerType                     `json:"scannerType,omitempty"`
	Dataset      *FPClassifierDatasetCounters    `json:"dataset,omitempty"`
	RetrainBatch *FPClassifierRetrainBatch       `json:"retrainBatch,omitempty"`
	IconState    FPClassifierScannerIconState    `json:"iconState,omitempty"`
	CounterState FPClassifierScannerCounterState `json:"counterState,omitempty"`
}

type FPClassifierModelSnapshotSource string

const (
	FPClassifierModelSnapshotSourceRETRAIN FPClassifierModelSnapshotSource = "RETRAIN"
	FPClassifierModelSnapshotSourceIMPORT  FPClassifierModelSnapshotSource = "IMPORT"
)

type FPClassifierModelSnapshotTrainingStatus string

const (
	FPClassifierModelSnapshotTrainingStatusSCHEDULED  FPClassifierModelSnapshotTrainingStatus = "SCHEDULED"
	FPClassifierModelSnapshotTrainingStatusINPROGRESS FPClassifierModelSnapshotTrainingStatus = "IN_PROGRESS"
	FPClassifierModelSnapshotTrainingStatusCOMPLETED  FPClassifierModelSnapshotTrainingStatus = "COMPLETED"
	FPClassifierModelSnapshotTrainingStatusFAILED     FPClassifierModelSnapshotTrainingStatus = "FAILED"
)

type FPClassifierRetrainBatch struct {
	TotalCount         int `json:"totalCount,omitempty"`
	RequiredTotalCount int `json:"requiredTotalCount,omitempty"`
}

type FPClassifierScannerCounterState string

const (
	FPClassifierScannerCounterStateCLASSIFIERDISABLED                  FPClassifierScannerCounterState = "CLASSIFIER_DISABLED"
	FPClassifierScannerCounterStateNOLICENSE                           FPClassifierScannerCounterState = "NO_LICENSE"
	FPClassifierScannerCounterStateINITIALDATASETCOLLECTINGACTIVE      FPClassifierScannerCounterState = "INITIAL_DATASET_COLLECTING_ACTIVE"
	FPClassifierScannerCounterStateINITIALDATASETCOLLECTINGCOLLECTONLY FPClassifierScannerCounterState = "INITIAL_DATASET_COLLECTING_COLLECT_ONLY"
	FPClassifierScannerCounterStateINITIALDATASETCOLLECTINGDISABLED    FPClassifierScannerCounterState = "INITIAL_DATASET_COLLECTING_DISABLED"
	FPClassifierScannerCounterStateINITIALDATASETREADYACTIVE           FPClassifierScannerCounterState = "INITIAL_DATASET_READY_ACTIVE"
	FPClassifierScannerCounterStateINITIALDATASETREADYCOLLECTONLY      FPClassifierScannerCounterState = "INITIAL_DATASET_READY_COLLECT_ONLY"
	FPClassifierScannerCounterStateINITIALDATASETREADYDISABLED         FPClassifierScannerCounterState = "INITIAL_DATASET_READY_DISABLED"
	FPClassifierScannerCounterStateTRAININGINPROGRESS                  FPClassifierScannerCounterState = "TRAINING_IN_PROGRESS"
	FPClassifierScannerCounterStateRETRAINBATCHCOLLECTINGACTIVE        FPClassifierScannerCounterState = "RETRAIN_BATCH_COLLECTING_ACTIVE"
	FPClassifierScannerCounterStateRETRAINBATCHCOLLECTINGCOLLECTONLY   FPClassifierScannerCounterState = "RETRAIN_BATCH_COLLECTING_COLLECT_ONLY"
	FPClassifierScannerCounterStateRETRAINBATCHCOLLECTINGDISABLED      FPClassifierScannerCounterState = "RETRAIN_BATCH_COLLECTING_DISABLED"
	FPClassifierScannerCounterStateRETRAINBATCHREADYACTIVE             FPClassifierScannerCounterState = "RETRAIN_BATCH_READY_ACTIVE"
	FPClassifierScannerCounterStateRETRAINBATCHREADYCOLLECTONLY        FPClassifierScannerCounterState = "RETRAIN_BATCH_READY_COLLECT_ONLY"
	FPClassifierScannerCounterStateRETRAINBATCHREADYDISABLED           FPClassifierScannerCounterState = "RETRAIN_BATCH_READY_DISABLED"
	FPClassifierScannerCounterStateERRORNOSNAPSHOTSACTIVE              FPClassifierScannerCounterState = "ERROR_NO_SNAPSHOTS_ACTIVE"
	FPClassifierScannerCounterStateERRORNOSNAPSHOTSCOLLECTONLY         FPClassifierScannerCounterState = "ERROR_NO_SNAPSHOTS_COLLECT_ONLY"
	FPClassifierScannerCounterStateERRORNOSNAPSHOTSDISABLED            FPClassifierScannerCounterState = "ERROR_NO_SNAPSHOTS_DISABLED"
	FPClassifierScannerCounterStateERRORWITHSNAPSHOTSACTIVE            FPClassifierScannerCounterState = "ERROR_WITH_SNAPSHOTS_ACTIVE"
	FPClassifierScannerCounterStateERRORWITHSNAPSHOTSCOLLECTONLY       FPClassifierScannerCounterState = "ERROR_WITH_SNAPSHOTS_COLLECT_ONLY"
	FPClassifierScannerCounterStateERRORWITHSNAPSHOTSDISABLED          FPClassifierScannerCounterState = "ERROR_WITH_SNAPSHOTS_DISABLED"
)

type FPClassifierScannerIconState string

const (
	FPClassifierScannerIconStateACTIVEOK      FPClassifierScannerIconState = "ACTIVE_OK"
	FPClassifierScannerIconStateACTIVEERROR   FPClassifierScannerIconState = "ACTIVE_ERROR"
	FPClassifierScannerIconStateINACTIVEERROR FPClassifierScannerIconState = "INACTIVE_ERROR"
	FPClassifierScannerIconStateINACTIVEEMPTY FPClassifierScannerIconState = "INACTIVE_EMPTY"
)

type FPClassifierSettings struct {
	IsActive     bool                     `json:"isActive,omitempty"`
	TrainingMode FPClassifierTrainingMode `json:"trainingMode,omitempty"`
	Initialized  bool                     `json:"initialized,omitempty"`
	State        FPClassifierState        `json:"state,omitempty"`
}

type FPClassifierSnapshot struct {
	ID                 UUID                                    `json:"id,omitempty"`
	Name               string                                  `json:"name,omitempty"`
	Created            string                                  `json:"created,omitempty"`
	Updated            string                                  `json:"updated,omitempty"`
	SnapshotStats      *FPClassifierSnapshotStats              `json:"snapshotStats,omitempty"`
	TrainingStatus     FPClassifierModelSnapshotTrainingStatus `json:"trainingStatus,omitempty"`
	ScheduledAt        *string                                 `json:"scheduledAt,omitempty"`
	TrainingStartedAt  *string                                 `json:"trainingStartedAt,omitempty"`
	TrainingFinishedAt *string                                 `json:"trainingFinishedAt,omitempty"`
	IsCurrent          bool                                    `json:"isCurrent,omitempty"`
	ScannerType        ScannerType                             `json:"scannerType,omitempty"`
	Source             FPClassifierModelSnapshotSource         `json:"source,omitempty"`
	ActivatedAt        *string                                 `json:"activatedAt,omitempty"`
}

type FPClassifierSnapshotShort struct {
	ID          UUID        `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	IsCurrent   bool        `json:"isCurrent,omitempty"`
	ScannerType ScannerType `json:"scannerType,omitempty"`
}

type FPClassifierSnapshotStats struct {
	FPPercentage          *int `json:"fpPercentage,omitempty"`
	TpPercentage          *int `json:"tpPercentage,omitempty"`
	DatasetRecordsCount   *int `json:"datasetRecordsCount,omitempty"`
	TrainingExamplesCount *int `json:"trainingExamplesCount,omitempty"`
}

type FPClassifierState string

const (
	FPClassifierStateTURNEDOFF      FPClassifierState = "TURNED_OFF"
	FPClassifierStateLOCKED         FPClassifierState = "LOCKED"
	FPClassifierStateNOTSCORING     FPClassifierState = "NOT_SCORING"
	FPClassifierStatePARTIALSCORING FPClassifierState = "PARTIAL_SCORING"
	FPClassifierStateSCORING        FPClassifierState = "SCORING"
)

type FPClassifierTrainingMode string

const (
	FPClassifierTrainingModeACTIVE      FPClassifierTrainingMode = "ACTIVE"
	FPClassifierTrainingModeCOLLECTONLY FPClassifierTrainingMode = "COLLECT_ONLY"
	FPClassifierTrainingModeDISABLED    FPClassifierTrainingMode = "DISABLED"
)

type FPScoreAIAppliedDetail struct {
	Integration *SpaceIntegrationWithMeta `json:"integration,omitempty"`
	Model       *string                   `json:"model,omitempty"`
	AppliedAt   *string                   `json:"appliedAt,omitempty"`
	Explanation *string                   `json:"explanation,omitempty"`
}

type FPScoreClassifierAppliedDetail struct {
	Snapshot   *FPClassifierSnapshotShort `json:"snapshot,omitempty"`
	AppliedAt  *string                    `json:"appliedAt,omitempty"`
	ShapValues []FPScoreShapValue         `json:"shapValues,omitempty"`
}

type FPScoreDetail struct {
	Value             *int                            `json:"value,omitempty"`
	Level             *string                         `json:"level,omitempty"`
	Type              *FPScoreSourceType              `json:"type,omitempty"`
	ClassifierApplied *FPScoreClassifierAppliedDetail `json:"classifierApplied,omitempty"`
	AIApplied         *FPScoreAIAppliedDetail         `json:"aiApplied,omitempty"`
}

type FPScore struct {
	Value             *int                      `json:"value,omitempty"`
	Level             *string                   `json:"level,omitempty"`
	Type              *FPScoreSourceType        `json:"type,omitempty"`
	ClassifierApplied *FPScoreClassifierApplied `json:"classifierApplied,omitempty"`
	AIApplied         *FPScoreAIApplied         `json:"aiApplied,omitempty"`
}

type FPScoreAIApplied struct {
	Integration *SpaceIntegrationWithMeta `json:"integration,omitempty"`
	Model       *string                   `json:"model,omitempty"`
	AppliedAt   *string                   `json:"appliedAt,omitempty"`
}

type FPScoreClassifierApplied struct {
	Snapshot  *FPClassifierSnapshotShort `json:"snapshot,omitempty"`
	AppliedAt *string                    `json:"appliedAt,omitempty"`
}

type FPScoreShapValue struct {
	FeatureSlug  string  `json:"featureSlug,omitempty"`
	FeatureValue string  `json:"featureValue,omitempty"`
	ShapValue    float64 `json:"shapValue,omitempty"`
}

type FPScoreSourceType string

const (
	FPScoreSourceTypeAIINTEGRATION FPScoreSourceType = "AI_INTEGRATION"
	FPScoreSourceTypeFPCLASSIFIER  FPScoreSourceType = "FP_CLASSIFIER"
)

type FilterMetaRelationType string

const (
	FilterMetaRelationTypeREPOSITORY      FilterMetaRelationType = "REPOSITORY"
	FilterMetaRelationTypePRODUCT         FilterMetaRelationType = "PRODUCT"
	FilterMetaRelationTypeARTIFACTVERSION FilterMetaRelationType = "ARTIFACT_VERSION"
	FilterMetaRelationTypeASSET           FilterMetaRelationType = "ASSET"
	FilterMetaRelationTypeSCANPIPELINE    FilterMetaRelationType = "SCAN_PIPELINE"
)

type FilterMetaResponse []FilterMetaResponseInner

type FilterMetaResponseInner struct {
	Target FilterMetaTarget `json:"target,omitempty"`
	Values []string         `json:"values,omitempty"`
}

type FilterMetaTarget string

const (
	FilterMetaTargetSCANPIPELINESCANNINGTOOL    FilterMetaTarget = "SCAN_PIPELINE__SCANNING_TOOL"
	FilterMetaTargetSCANPIPELINEJOBSCANNINGTOOL FilterMetaTarget = "SCAN_PIPELINE_JOB__SCANNING_TOOL"
	FilterMetaTargetSCANPIPELINEINITIATORTYPE   FilterMetaTarget = "SCAN_PIPELINE__INITIATOR_TYPE"
)

type HealthResponse struct {
	DataStorageReady  bool  `json:"dataStorageReady,omitempty"`
	VaultServiceReady *bool `json:"vaultServiceReady,omitempty"`
}

type LDAPAuthenticationConfig struct {
	ID                       UUID               `json:"id,omitempty"`
	Protocol                 string             `json:"protocol,omitempty"`
	Host                     string             `json:"host,omitempty"`
	Port                     int                `json:"port,omitempty"`
	BaseDN                   string             `json:"baseDN,omitempty"`
	AdminLogin               string             `json:"adminLogin,omitempty"`
	AdminPassword            string             `json:"adminPassword,omitempty"`
	AdminGroupDN             *string            `json:"adminGroupDN,omitempty"`
	EditorGroupDN            *string            `json:"editorGroupDN,omitempty"`
	ReadonlyGroupDN          *string            `json:"readonlyGroupDN,omitempty"`
	BlockedGroupDN           *string            `json:"blockedGroupDN,omitempty"`
	UserFilter               string             `json:"userFilter,omitempty"`
	ObjectClass              string             `json:"objectClass,omitempty"`
	LoginAttribute           string             `json:"loginAttribute,omitempty"`
	EmailAttribute           string             `json:"emailAttribute,omitempty"`
	FirstNameAttribute       string             `json:"firstNameAttribute,omitempty"`
	LastNameAttribute        string             `json:"lastNameAttribute,omitempty"`
	Enabled                  bool               `json:"enabled,omitempty"`
	UserGroupsMappingEnabled bool               `json:"userGroupsMappingEnabled,omitempty"`
	UserGroupsMapping        []LDAPGroupMapping `json:"userGroupsMapping,omitempty"`
}

type LDAPGroupMapping struct {
	LDAPGroupDN      string `json:"ldapGroupDN,omitempty"`
	VampyUserGroupID UUID   `json:"vampyUserGroupID,omitempty"`
}

type LDAPUser struct {
	Login     string    `json:"login,omitempty"`
	Email     string    `json:"email,omitempty"`
	Dn        string    `json:"dn,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Groups    []string  `json:"groups,omitempty"`
	IsAdmin   bool      `json:"isAdmin,omitempty"`
	IsBlocked bool      `json:"isBlocked,omitempty"`
	UserRole  *string   `json:"userRole,omitempty"`
	UserGroup UserGroup `json:"userGroup,omitempty"`
}

type OIDCConfig struct {
	ID                       UUID               `json:"id,omitempty"`
	Enabled                  bool               `json:"enabled,omitempty"`
	UserGroupsMappingEnabled bool               `json:"userGroupsMappingEnabled,omitempty"`
	Provider                 OIDCProvider       `json:"provider,omitempty"`
	URL                      string             `json:"url,omitempty"`
	Realm                    string             `json:"realm,omitempty"`
	Scope                    string             `json:"scope,omitempty"`
	ClientID                 string             `json:"clientID,omitempty"`
	ClientName               string             `json:"clientName,omitempty"`
	ClientSecret             string             `json:"clientSecret,omitempty"`
	EmailAttribute           string             `json:"emailAttribute,omitempty"`
	LoginAttribute           string             `json:"loginAttribute,omitempty"`
	FirstNameAttribute       string             `json:"firstNameAttribute,omitempty"`
	LastNameAttribute        string             `json:"lastNameAttribute,omitempty"`
	GroupsAttribute          string             `json:"groupsAttribute,omitempty"`
	MappedGroups             map[string]any     `json:"mappedGroups,omitempty"`
	UserGroupsMapping        []OIDCGroupMapping `json:"userGroupsMapping,omitempty"`
	BlockedGroups            []string           `json:"blockedGroups,omitempty"`
	AdminGroups              []string           `json:"adminGroups,omitempty"`
	EditorGroups             []string           `json:"editorGroups,omitempty"`
	ReadonlyGroups           []string           `json:"readonlyGroups,omitempty"`
	CallbackURL              string             `json:"callbackURL,omitempty"`
	AuthorizeURL             string             `json:"authorizeURL,omitempty"`
	AccessTokenURL           string             `json:"accessTokenURL,omitempty"`
	UserInfoURL              string             `json:"userInfoURL,omitempty"`
	ServerMetadataURL        string             `json:"serverMetadataURL,omitempty"`
}

type OIDCDefaults struct {
	Scope              string `json:"scope,omitempty"`
	EmailAttribute     string `json:"emailAttribute,omitempty"`
	LoginAttribute     string `json:"loginAttribute,omitempty"`
	FirstNameAttribute string `json:"firstNameAttribute,omitempty"`
	LastNameAttribute  string `json:"lastNameAttribute,omitempty"`
	AuthorizeURL       string `json:"authorizeURL,omitempty"`
	AccessTokenURL     string `json:"accessTokenURL,omitempty"`
	UserInfoURL        string `json:"userInfoURL,omitempty"`
	ServerMetadataURL  string `json:"serverMetadataURL,omitempty"`
	CallbackURL        string `json:"callbackURL,omitempty"`
	ClientsURL         string `json:"clientsURL,omitempty"`
}

type OIDCGroupMapping struct {
	VampyUserGroupID UUID     `json:"vampyUserGroupID,omitempty"`
	OIDCUserGroups   []string `json:"oidcUserGroups,omitempty"`
}

type OIDCProvider string

const (
	OIDCProviderKeycloak OIDCProvider = "keycloak"
)

type OIDCRedirect struct {
	RedirectURL string `json:"redirectURL,omitempty"`
}

type OIDCTestConnectRequest struct {
	URL               string `json:"url,omitempty"`
	Realm             string `json:"realm,omitempty"`
	AuthorizeURL      string `json:"authorizeURL,omitempty"`
	AccessTokenURL    string `json:"accessTokenURL,omitempty"`
	UserInfoURL       string `json:"userInfoURL,omitempty"`
	ServerMetadataURL string `json:"serverMetadataURL,omitempty"`
}

type OIDCTestConnectResponse struct {
	AuthorizeURL      *OIDCTestConnectURL `json:"authorizeURL,omitempty"`
	AccessTokenURL    *OIDCTestConnectURL `json:"accessTokenURL,omitempty"`
	UserInfoURL       *OIDCTestConnectURL `json:"userInfoURL,omitempty"`
	ServerMetadataURL *OIDCTestConnectURL `json:"serverMetadataURL,omitempty"`
}

type OIDCTestConnectURL struct {
	Available bool   `json:"available,omitempty"`
	Message   string `json:"message,omitempty"`
}

type SLASetting struct {
	SLADaysInfo     *int `json:"SLADaysInfo,omitempty"`
	SLADaysLow      *int `json:"SLADaysLow,omitempty"`
	SLADaysMedium   *int `json:"SLADaysMedium,omitempty"`
	SLADaysHigh     *int `json:"SLADaysHigh,omitempty"`
	SLADaysCritical *int `json:"SLADaysCritical,omitempty"`
}

type SLASettingWithDefault struct {
	Actual  *SLASetting `json:"actual,omitempty"`
	Default *SLASetting `json:"default,omitempty"`
}

type ScheduledStatus struct {
	StatusFromName           string  `json:"statusFromName,omitempty"`
	StatusToName             string  `json:"statusToName,omitempty"`
	StatusExpiryTime         string  `json:"statusExpiryTime,omitempty"`
	Source                   *string `json:"source,omitempty"`
	InheritedFromScanIssueID *UUID   `json:"inheritedFromScanIssueId,omitempty"`
}

type ServerInfo struct {
	AuthenticationMethods []string `json:"authenticationMethods,omitempty"`
}

type SMTPConfig struct {
	Host               string  `json:"host,omitempty"`
	Port               int     `json:"port,omitempty"`
	Username           *string `json:"username,omitempty"`
	Password           *string `json:"password,omitempty"`
	FromAddress        string  `json:"fromAddress,omitempty"`
	ConnectionSecurity string  `json:"connectionSecurity,omitempty"`
	Enabled            bool    `json:"enabled,omitempty"`
	SkipAuth           bool    `json:"skipAuth,omitempty"`
}

type SystemWideSettings struct {
	MFARequired bool               `json:"mfaRequired,omitempty"`
	Constants   *SettingsConstants `json:"constants,omitempty"`
}

type TemplateSlug string

const (
	TemplateSlugDETAILED TemplateSlug = "DETAILED"
)
