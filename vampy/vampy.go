// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type Action any

type ActionType string

const (
	ActionTypeCHANGESCANISSUESTATUS   ActionType = "CHANGE_SCAN_ISSUE_STATUS"
	ActionTypeCHANGESCANISSUESEVERITY ActionType = "CHANGE_SCAN_ISSUE_SEVERITY"
	ActionTypeADDSCANISSUECOMMENT     ActionType = "ADD_SCAN_ISSUE_COMMENT"
	ActionTypeCREATEREPOSITORY        ActionType = "CREATE_REPOSITORY"
	ActionTypeSENDNOTIFICATION        ActionType = "SEND_NOTIFICATION"
	ActionTypeRUNSCAN                 ActionType = "RUN_SCAN"
)

type AnyValue any

type ArtifactDelete any

type ArtifactDetail struct {
	Artifact
	LastOrchestrationScan OrchestrationScanDetail `json:"lastOrchestrationScan,omitempty"`
}

type ArtifactEdit any

type ArtifactImport any

type ArtifactImportTask any

type Artifact struct {
	ArtifactShort
	ScanIssuesStat        ProductScanIssuesStat   `json:"scanIssuesStat,omitempty"`
	VersionsCount         int                     `json:"versionsCount,omitempty"`
	Integration           SpaceIntegration        `json:"integration,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail `json:"lastOrchestrationScan,omitempty"`
}

type ArtifactShort any

type ArtifactType string

const (
	ArtifactTypeDOCKER ArtifactType = "DOCKER"
)

type ArtifactVersionDelete any

type ArtifactVersion any

type ArtifactVersionShort any

type ArtifactVersionSyncArtifact struct {
	ArtifactID         UUID   `json:"artifactID,omitempty"`
	ArtifactSlug       string `json:"artifactSlug,omitempty"`
	ProductID          UUID   `json:"productID,omitempty"`
	ProductSlug        string `json:"productSlug,omitempty"`
	IntegrationID      UUID   `json:"integrationID,omitempty"`
	IntegrationService string `json:"integrationService,omitempty"`
	VersionsCount      int    `json:"versionsCount,omitempty"`
}

type ArtifactVersionSyncRequest struct {
	Artifact string  `json:"artifact"`
	Product  *string `json:"product,omitempty"`
}

type ArtifactVersionSyncResponse struct {
	Artifact        string                        `json:"artifact,omitempty"`
	Product         *string                       `json:"product,omitempty"`
	TotalArtifacts  int                           `json:"totalArtifacts,omitempty"`
	SyncedArtifacts []ArtifactVersionSyncArtifact `json:"syncedArtifacts,omitempty"`
}

type AssetCreate any

type AssetCriticality string

const (
	AssetCriticalityLOW      AssetCriticality = "LOW"
	AssetCriticalityMEDIUM   AssetCriticality = "MEDIUM"
	AssetCriticalityHIGH     AssetCriticality = "HIGH"
	AssetCriticalityCRITICAL AssetCriticality = "CRITICAL"
)

type AssetDetails struct {
	Asset
	Description           string                       `json:"description,omitempty"`
	Scheme                string                       `json:"scheme,omitempty"`
	Port                  int                          `json:"port,omitempty"`
	Path                  string                       `json:"path,omitempty"`
	IPs                   []string                     `json:"ips,omitempty"`
	ParentID              *UUID                        `json:"parentID,omitempty"`
	Created               string                       `json:"created,omitempty"`
	Updated               string                       `json:"updated,omitempty"`
	DefinedVariables      []OrchestrationScannerEnvVar `json:"definedVariables,omitempty"`
	PredefinedVariables   []OrchestrationScannerEnvVar `json:"predefinedVariables,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail      `json:"lastOrchestrationScan,omitempty"`
	LastScanPipeline      ScanPipeline                 `json:"lastScanPipeline,omitempty"`
}

type AssetForOrchestration struct {
	ID            UUID                         `json:"id,omitempty"`
	ParentID      *string                      `json:"parentID,omitempty"`
	Host          string                       `json:"host,omitempty"`
	ScanVariables []OrchestrationScannerEnvVar `json:"scanVariables,omitempty"`
}

type AssetPatch any

type AssetPermissionValue string

const (
	AssetPermissionValueASSETREAD AssetPermissionValue = "ASSET_READ"
	AssetPermissionValueASSETEDIT AssetPermissionValue = "ASSET_EDIT"
)

type Asset any

type AssetScope string

const (
	AssetScopePROD  AssetScope = "PROD"
	AssetScopeSTAGE AssetScope = "STAGE"
	AssetScopeDEV   AssetScope = "DEV"
)

type AsyncExportFormatter string

const (
	AsyncExportFormatterJSON AsyncExportFormatter = "JSON"
	AsyncExportFormatterPDF  AsyncExportFormatter = "PDF"
	AsyncExportFormatterCSV  AsyncExportFormatter = "CSV"
)

type AvailabilityStatus any

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
	TaskTracker BackgroundTaskMetaTaskTracker `json:"taskTracker,omitempty"`
	RulesCheck  BackgroundTaskMetaRulesCheck  `json:"rulesCheck,omitempty"`
}

type BackgroundTaskMetaTaskTracker struct {
	CompletedIssues []UUID `json:"completedIssues,omitempty"`
	FailedIssues    []UUID `json:"failedIssues,omitempty"`
}

type BackgroundTask struct {
	TaskID            UUID                  `json:"taskID,omitempty"`
	Status            BackgroundTaskStatus  `json:"status,omitempty"`
	DisplayName       string                `json:"displayName,omitempty"`
	Options           BackgroundTaskOptions `json:"options,omitempty"`
	TotalSubTasks     int                   `json:"totalSubTasks,omitempty"`
	CompletedSubTasks int                   `json:"completedSubTasks,omitempty"`
	Created           string                `json:"created,omitempty"`
	Updated           string                `json:"updated,omitempty"`
	ErrorDetails      *string               `json:"errorDetails,omitempty"`
	Meta              BackgroundTaskMeta    `json:"meta,omitempty"`
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

type BaseIssueStatus any

type BaseRule any

type BinaryString []byte

type BitBucketWebhookRequest struct {
	EventKey string `json:"eventKey"`
}

type BitbucketApplicationInfoRequest any

type BitbucketApplicationInfoResponse any

type BotUserRegistration any

type BotUser any

type BotUserToken any

type ByScannerType any

type BySeverity any

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

type CommentDetails any

type CommentPermissionValue string

const (
	CommentPermissionValueCOMMENTREAD CommentPermissionValue = "COMMENT_READ"
	CommentPermissionValueCOMMENTEDIT CommentPermissionValue = "COMMENT_EDIT"
)

type CommentRule any

type Comment any

type CommentStatus string

const (
	CommentStatusCREATED   CommentStatus = "CREATED"
	CommentStatusUPDATED   CommentStatus = "UPDATED"
	CommentStatusSYNCERROR CommentStatus = "SYNC_ERROR"
	CommentStatusSYNCED    CommentStatus = "SYNCED"
)

type CommentType string

const (
	CommentTypeLOCAL    CommentType = "LOCAL"
	CommentTypeEXTERNAL CommentType = "EXTERNAL"
)

type CommentUser any

type CommitSHA string

type ComparisonOperator string

const (
	ComparisonOperatorGT  ComparisonOperator = "GT"
	ComparisonOperatorLT  ComparisonOperator = "LT"
	ComparisonOperatorGTE ComparisonOperator = "GTE"
	ComparisonOperatorLTE ComparisonOperator = "LTE"
	ComparisonOperatorEQ  ComparisonOperator = "EQ"
	ComparisonOperatorNEQ ComparisonOperator = "NEQ"
)

type ComponentDetails struct {
	ComponentShort
	Repositories []RepositoryShort `json:"repositories,omitempty"`
}

type ComponentPatch struct {
	Description string `json:"description,omitempty"`
}

type ComponentPermissionValue string

const (
	ComponentPermissionValueCOMPONENTREAD ComponentPermissionValue = "COMPONENT_READ"
	ComponentPermissionValueCOMPONENTEDIT ComponentPermissionValue = "COMPONENT_EDIT"
)

type ComponentShort struct {
	ID             UUID                         `json:"id,omitempty"`
	Name           string                       `json:"name,omitempty"`
	Description    string                       `json:"description,omitempty"`
	Created        string                       `json:"created,omitempty"`
	Updated        string                       `json:"updated,omitempty"`
	RiskScore      *int                         `json:"riskScore,omitempty"`
	ScanIssuesStat ComponentShortScanIssuesStat `json:"scanIssuesStat,omitempty"`
	Versions       []ComponentShortVersions     `json:"versions,omitempty"`
	LastAggregated string                       `json:"lastAggregated,omitempty"`
	Permissions    []ComponentPermissionValue   `json:"permissions,omitempty"`
}

type ComponentShortScanIssuesStat struct {
	BySeverity []ScanIssueBySeverityStat `json:"bySeverity,omitempty"`
}

type ComponentShortVersions struct {
	ID      UUID   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type ComponentVersionDetails struct {
	ComponentVersion
	Repositories []RepositoryShort `json:"repositories,omitempty"`
}

type ComponentVersion struct {
	ID             UUID                         `json:"id,omitempty"`
	Name           string                       `json:"name,omitempty"`
	Version        string                       `json:"version,omitempty"`
	ComponentID    UUID                         `json:"componentID,omitempty"`
	Description    string                       `json:"description,omitempty"`
	Created        string                       `json:"created,omitempty"`
	Updated        string                       `json:"updated,omitempty"`
	RiskScore      *int                         `json:"riskScore,omitempty"`
	ScanIssuesStat ComponentShortScanIssuesStat `json:"scanIssuesStat,omitempty"`
	ActualVersion  string                       `json:"actualVersion,omitempty"`
	PackageURL     string                       `json:"packageUrl,omitempty"`
	PlatformEnum   string                       `json:"platformEnum,omitempty"`
	SwidTagID      string                       `json:"swidTagID,omitempty"`
	Classifier     string                       `json:"classifier,omitempty"`
	Filename       string                       `json:"filename,omitempty"`
	License        string                       `json:"license,omitempty"`
	Author         string                       `json:"author,omitempty"`
	Permissions    []ComponentPermissionValue   `json:"permissions,omitempty"`
}

type ComponentVersionShort struct {
	ID            UUID   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Version       string `json:"version,omitempty"`
	ComponentID   UUID   `json:"componentID,omitempty"`
	Created       string `json:"created,omitempty"`
	Updated       string `json:"updated,omitempty"`
	ActualVersion string `json:"actualVersion,omitempty"`
	PackageURL    string `json:"packageUrl,omitempty"`
}

type ComponentsCountCardData struct {
	SimpleCardValue
}

type Condition any

type ConditionType string

const (
	ConditionTypeGROUPAND                        ConditionType = "GROUP_AND"
	ConditionTypeGROUPOR                         ConditionType = "GROUP_OR"
	ConditionTypeSCANISSUEATTRIBUTEEQUALS        ConditionType = "SCAN_ISSUE_ATTRIBUTE_EQUALS"
	ConditionTypeSCANISSUEATTRIBUTECOMPARE       ConditionType = "SCAN_ISSUE_ATTRIBUTE_COMPARE"
	ConditionTypeSCANISSUEREPOSITORYBRANCHEQUALS ConditionType = "SCAN_ISSUE_REPOSITORY_BRANCH_EQUALS"
	ConditionTypeREPOSITORYCHOICE                ConditionType = "REPOSITORY_CHOICE"
	ConditionTypePRODUCTCHOICE                   ConditionType = "PRODUCT_CHOICE"
	ConditionTypeVCSREPOSITORYPATH               ConditionType = "VCS_REPOSITORY_PATH"
	ConditionTypeHASSOLUTION                     ConditionType = "HAS_SOLUTION"
	ConditionTypeVCSREPOSITORYBRANCH             ConditionType = "VCS_REPOSITORY_BRANCH"
)

type DashboardCardAggregationPeriod string

const (
	DashboardCardAggregationPeriodWEEK        DashboardCardAggregationPeriod = "WEEK"
	DashboardCardAggregationPeriodMONTH       DashboardCardAggregationPeriod = "MONTH"
	DashboardCardAggregationPeriodTWOMONTHS   DashboardCardAggregationPeriod = "TWO_MONTHS"
	DashboardCardAggregationPeriodTHREEMONTHS DashboardCardAggregationPeriod = "THREE_MONTHS"
	DashboardCardAggregationPeriodYEAR        DashboardCardAggregationPeriod = "YEAR"
)

type DashboardCardEntityRelationType string

const (
	DashboardCardEntityRelationTypeREPOSITORY DashboardCardEntityRelationType = "REPOSITORY"
	DashboardCardEntityRelationTypeVCSUSER    DashboardCardEntityRelationType = "VCS_USER"
)

type DashboardCardGraph any

type DashboardCardMeta any

type DashboardCard any

type DashboardCardType string

const (
	DashboardCardTypeBYCATEGORY DashboardCardType = "BY_CATEGORY"
	DashboardCardTypeENTITIES   DashboardCardType = "ENTITIES"
	DashboardCardTypeTABLE      DashboardCardType = "TABLE"
	DashboardCardTypeGRAPH      DashboardCardType = "GRAPH"
	DashboardCardTypeTOPN       DashboardCardType = "TOP_N"
)

type Dashboard any

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

type DefectDojoImportResponse any

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

type DependencyRelationFilter string

const (
	DependencyRelationFilterDIRECT     DependencyRelationFilter = "DIRECT"
	DependencyRelationFilterTRANSITIVE DependencyRelationFilter = "TRANSITIVE"
	DependencyRelationFilterUNKNOWN    DependencyRelationFilter = "UNKNOWN"
)

type DependencyRelation string

const (
	DependencyRelationDIRECT     DependencyRelation = "DIRECT"
	DependencyRelationTRANSITIVE DependencyRelation = "TRANSITIVE"
)

type DetailPermissionUser struct {
	UserID   UUID   `json:"userID,omitempty"`
	Level    string `json:"level,omitempty"`
	IsMapped bool   `json:"isMapped,omitempty"`
}

type Email any

type Engagement struct {
	ID                   UUID   `json:"id,omitempty"`
	ExternalScanPipeline string `json:"externalScanPipeline,omitempty"`
}

type Entity any

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

type ExternalArtifact any

type ExternalRepository any

type FPClassifierDatasetCounters any

type FPClassifierDatasetStats any

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

type FPClassifierRetrainBatch any

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

type FPClassifierSettings any

type FPClassifierSnapshot any

type FPClassifierSnapshotShort any

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
	Integration SpaceIntegrationWithMeta `json:"integration,omitempty"`
	Model       *string                  `json:"model,omitempty"`
	AppliedAt   *string                  `json:"appliedAt,omitempty"`
	Explanation *string                  `json:"explanation,omitempty"`
}

type FPScoreClassifierAppliedDetail struct {
	Snapshot   FPClassifierSnapshotShort `json:"snapshot,omitempty"`
	AppliedAt  *string                   `json:"appliedAt,omitempty"`
	ShapValues []FPScoreShapValue        `json:"shapValues,omitempty"`
}

type FPScoreDetail struct {
	Value             *int                           `json:"value,omitempty"`
	Level             *string                        `json:"level,omitempty"`
	Type              *FPScoreSourceType             `json:"type,omitempty"`
	ClassifierApplied FPScoreClassifierAppliedDetail `json:"classifierApplied,omitempty"`
	AIApplied         FPScoreAIAppliedDetail         `json:"aiApplied,omitempty"`
}

type FPScore struct {
	Value             *int                     `json:"value,omitempty"`
	Level             *string                  `json:"level,omitempty"`
	Type              *FPScoreSourceType       `json:"type,omitempty"`
	ClassifierApplied FPScoreClassifierApplied `json:"classifierApplied,omitempty"`
	AIApplied         FPScoreAIApplied         `json:"aiApplied,omitempty"`
}

type FPScoreAIApplied struct {
	Integration SpaceIntegrationWithMeta `json:"integration,omitempty"`
	Model       *string                  `json:"model,omitempty"`
	AppliedAt   *string                  `json:"appliedAt,omitempty"`
}

type FPScoreClassifierApplied struct {
	Snapshot  FPClassifierSnapshotShort `json:"snapshot,omitempty"`
	AppliedAt *string                   `json:"appliedAt,omitempty"`
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

type FieldType string

const (
	FieldTypeTEXT                             FieldType = "TEXT"
	FieldTypeTEXTAREA                         FieldType = "TEXTAREA"
	FieldTypeDATE                             FieldType = "DATE"
	FieldTypeBOOL                             FieldType = "BOOL"
	FieldTypeINTEGER                          FieldType = "INTEGER"
	FieldTypeSTATUSCHOICE                     FieldType = "STATUS_CHOICE"
	FieldTypeREPOSITORYCHOICE                 FieldType = "REPOSITORY_CHOICE"
	FieldTypePRODUCTCHOICE                    FieldType = "PRODUCT_CHOICE"
	FieldTypeSEVERITYCHOICE                   FieldType = "SEVERITY_CHOICE"
	FieldTypeCRITICALITYCHOICE                FieldType = "CRITICALITY_CHOICE"
	FieldTypePARSERCHOICE                     FieldType = "PARSER_CHOICE"
	FieldTypeSCANNERTYPECHOICE                FieldType = "SCANNER_TYPE_CHOICE"
	FieldTypeMULTIPLEISSUETYPECHOICE          FieldType = "MULTIPLE_ISSUE_TYPE_CHOICE"
	FieldTypeMULTIPLESTATUSCHOICE             FieldType = "MULTIPLE_STATUS_CHOICE"
	FieldTypeMULTIPLEPRODUCTCHOICE            FieldType = "MULTIPLE_PRODUCT_CHOICE"
	FieldTypeMULTIPLECRITICALITYCHOICE        FieldType = "MULTIPLE_CRITICALITY_CHOICE"
	FieldTypeMULTIPLESEVERITYCHOICE           FieldType = "MULTIPLE_SEVERITY_CHOICE"
	FieldTypeMULTIPLEPARSERCHOICE             FieldType = "MULTIPLE_PARSER_CHOICE"
	FieldTypeMULTIPLESCANNERTYPECHOICE        FieldType = "MULTIPLE_SCANNER_TYPE_CHOICE"
	FieldTypeMULTIPLEREPOSITORYCHOICE         FieldType = "MULTIPLE_REPOSITORY_CHOICE"
	FieldTypeMULTIPLEDEPENDENCYRELATIONCHOICE FieldType = "MULTIPLE_DEPENDENCY_RELATION_CHOICE"
	FieldTypeINTEGERCOMPARE                   FieldType = "INTEGER_COMPARE"
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

type FindingCard struct {
	Finding
	DescriptionMD    string           `json:"descriptionMD,omitempty"`
	Solution         string           `json:"solution,omitempty"`
	SolutionMD       string           `json:"solutionMD,omitempty"`
	Data             map[string]any   `json:"data,omitempty"`
	Assets           []Asset          `json:"assets,omitempty"`
	Repository       Repository       `json:"repository,omitempty"`
	RepositoryBranch RepositoryBranch `json:"repositoryBranch,omitempty"`
	ScanIssue        ScanIssue        `json:"scanIssue,omitempty"`
	ScanIssueType    ScanIssueType    `json:"scanIssueType,omitempty"`
}

type Finding any

type FindingState string

const (
	FindingStateDETECTED   FindingState = "DETECTED"
	FindingStateRESOLVED   FindingState = "RESOLVED"
	FindingStateDUPLICATED FindingState = "DUPLICATED"
)

type FindingStatus string

const (
	FindingStatusINBOX     FindingStatus = "INBOX"
	FindingStatusPERFORMED FindingStatus = "PERFORMED"
)

type GitLabWebhookRequest struct {
	EventName string `json:"event_name"`
}

type HealthResponse any

type InformationParameter any

type IntegrationAILanguage string

const (
	IntegrationAILanguageEN IntegrationAILanguage = "EN"
	IntegrationAILanguageRU IntegrationAILanguage = "RU"
)

type IntegrationAIModel struct {
	ID          string                        `json:"id,omitempty"`
	IsDefault   bool                          `json:"isDefault,omitempty"`
	Integration IntegrationAIModelIntegration `json:"integration,omitempty"`
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

type IntegrationAI any

type IntegrationAIModels struct {
	Updated string   `json:"updated,omitempty"`
	Items   []string `json:"items,omitempty"`
}

type IntegrationAIVendorParamInfo any

type IntegrationAIVendorParam any

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

type IntegrationArtifact any

type IntegrationArtifactServiceType string

const (
	IntegrationArtifactServiceTypeHARBOR IntegrationArtifactServiceType = "HARBOR"
	IntegrationArtifactServiceTypeNEXUS  IntegrationArtifactServiceType = "NEXUS"
	IntegrationArtifactServiceTypeOTHER  IntegrationArtifactServiceType = "OTHER"
)

type IntegrationDefectDojo any

type IntegrationKTrackerBoard any

type IntegrationKTrackerCardType any

type IntegrationKTracker any

type IntegrationKTrackerSpace any

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

type IntegrationVCS any

type IntegrationVCSServiceType string

const (
	IntegrationVCSServiceTypeGITLAB      IntegrationVCSServiceType = "GITLAB"
	IntegrationVCSServiceTypeBITBUCKET   IntegrationVCSServiceType = "BITBUCKET"
	IntegrationVCSServiceTypeAZUREDEVOPS IntegrationVCSServiceType = "AZURE_DEVOPS"
	IntegrationVCSServiceTypeGITEA       IntegrationVCSServiceType = "GITEA"
	IntegrationVCSServiceTypeGITFLIC     IntegrationVCSServiceType = "GITFLIC"
	IntegrationVCSServiceTypeGITFLAME    IntegrationVCSServiceType = "GITFLAME"
)

type IssueChangeSource string

const (
	IssueChangeSourceUSER            IssueChangeSource = "USER"
	IssueChangeSourceINTEGRATION     IssueChangeSource = "INTEGRATION"
	IssueChangeSourceSCANUPLOAD      IssueChangeSource = "SCAN_UPLOAD"
	IssueChangeSourceRULESENGINE     IssueChangeSource = "RULES_ENGINE"
	IssueChangeSourceSCHEDULER       IssueChangeSource = "SCHEDULER"
	IssueChangeSourceASSOCIATEDISSUE IssueChangeSource = "ASSOCIATED_ISSUE"
	IssueChangeSourceFPSCORING       IssueChangeSource = "FP_SCORING"
)

type IssueChangeType string

const (
	IssueChangeTypeSTATUS   IssueChangeType = "STATUS"
	IssueChangeTypeFPSCORE  IssueChangeType = "FP_SCORE"
	IssueChangeTypeSEVERITY IssueChangeType = "SEVERITY"
)

type IssueExport any

type IssueStatusChangeIntegration any

type IssueStatusChangeIntegrationMeta struct {
	AI IssueStatusChangeIntegrationMetaAI `json:"ai,omitempty"`
}

type IssueStatusChangeIntegrationMetaAI struct {
	Vendor                string  `json:"vendor,omitempty"`
	IsDefault             bool    `json:"isDefault,omitempty"`
	DefaultModel          *string `json:"defaultModel,omitempty"`
	DefaultModelAvailable bool    `json:"defaultModelAvailable,omitempty"`
}

type IssueStatusChangeRule any

type IssueStatusChangeScanUpload any

type IssueStatusChangeSource string

const (
	IssueStatusChangeSourceUSER            IssueStatusChangeSource = "USER"
	IssueStatusChangeSourceINTEGRATION     IssueStatusChangeSource = "INTEGRATION"
	IssueStatusChangeSourceSCANUPLOAD      IssueStatusChangeSource = "SCAN_UPLOAD"
	IssueStatusChangeSourceRULESENGINE     IssueStatusChangeSource = "RULES_ENGINE"
	IssueStatusChangeSourceSCHEDULER       IssueStatusChangeSource = "SCHEDULER"
	IssueStatusChangeSourceASSOCIATEDISSUE IssueStatusChangeSource = "ASSOCIATED_ISSUE"
)

type IssueStatusChangeUser any

type IssueStatusCreate any

type IssueStatus struct {
	BaseIssueStatus
	Name      string `json:"name,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}

type IssueStatusUpdate struct {
	BaseIssueStatus
}

type IssuesByScannerTypeCardData struct {
	SimpleCardValue
	ByScannerType
}

type IssuesByStatusCardData struct {
	SimpleCardValue
	ByStatus NameBasedBySeverity `json:"by_status,omitempty"`
}

type JiraApplicationInfoRequest any

type JiraApplicationInfoResponse any

type JiraIntegration any

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

type JiraMetaFields any

type JiraMetaInfoResponse any

type JiraMetaInfoResponseExternalStatuses struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type JiraMetaInfoResponseIssuetypes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
}

type JiraMetaInfoResponseProjects struct {
	Key        string                           `json:"key"`
	Name       string                           `json:"name"`
	Issuetypes []JiraMetaInfoResponseIssuetypes `json:"issuetypes,omitempty"`
}

type JiraSearchIssuesResponse any

type JiraSearchResponse any

type JiraWebhookRequest any

type JiraWebhookRequestComment struct {
	ID     string                          `json:"id,omitempty"`
	Body   string                          `json:"body,omitempty"`
	Author JiraWebhookRequestCommentAuthor `json:"author,omitempty"`
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
	ID     string                        `json:"id,omitempty"`
	Key    string                        `json:"key"`
	Fields JiraWebhookRequestIssueFields `json:"fields"`
}

type JiraWebhookRequestIssueFields struct {
	Status JiraWebhookRequestIssueFieldsStatus `json:"status,omitempty"`
}

type JiraWebhookRequestIssueFieldsStatus struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type LDAPAuthenticationConfig any

type LDAPGroupMapping any

type LDAPUser any

type LicenseInformation any

type LicenseInformationCommonParameters struct {
	User                    InformationParameter `json:"user,omitempty"`
	BotUser                 InformationParameter `json:"botUser,omitempty"`
	LDAP                    InformationParameter `json:"ldap,omitempty"`
	MFA                     InformationParameter `json:"mfa,omitempty"`
	Repository              InformationParameter `json:"repository,omitempty"`
	Product                 InformationParameter `json:"product,omitempty"`
	Jira                    InformationParameter `json:"jira,omitempty"`
	Ktracker                InformationParameter `json:"ktracker,omitempty"`
	EvaTeam                 InformationParameter `json:"evaTeam,omitempty"`
	TaskTrackerIntegration  InformationParameter `json:"taskTrackerIntegration,omitempty"`
	TaskTrackerIssue        InformationParameter `json:"taskTrackerIssue,omitempty"`
	Gitlab                  InformationParameter `json:"gitlab,omitempty"`
	Bitbucket               InformationParameter `json:"bitbucket,omitempty"`
	Gitea                   InformationParameter `json:"gitea,omitempty"`
	Gitflic                 InformationParameter `json:"gitflic,omitempty"`
	Gitflame                InformationParameter `json:"gitflame,omitempty"`
	VCSIntegration          InformationParameter `json:"vcsIntegration,omitempty"`
	QualityGate             InformationParameter `json:"qualityGate,omitempty"`
	Orchestration           InformationParameter `json:"orchestration,omitempty"`
	Ptai                    InformationParameter `json:"ptai,omitempty"`
	CodeScoring             InformationParameter `json:"codeScoring,omitempty"`
	AppScreener             InformationParameter `json:"appScreener,omitempty"`
	AppsecTrack             InformationParameter `json:"appsecTrack,omitempty"`
	Sastav                  InformationParameter `json:"sastav,omitempty"`
	ScannerIntegration      InformationParameter `json:"scannerIntegration,omitempty"`
	Harbor                  InformationParameter `json:"harbor,omitempty"`
	Nexus                   InformationParameter `json:"nexus,omitempty"`
	ArtifactIntegration     InformationParameter `json:"artifactIntegration,omitempty"`
	OIDC                    InformationParameter `json:"oidc,omitempty"`
	PDFReport               InformationParameter `json:"pdfReport,omitempty"`
	ImportIntegration       InformationParameter `json:"importIntegration,omitempty"`
	DefectDojo              InformationParameter `json:"defectDojo,omitempty"`
	AIIntegration           InformationParameter `json:"aiIntegration,omitempty"`
	ScanPipelinesComparison InformationParameter `json:"scanPipelinesComparison,omitempty"`
	AzureDevOps             InformationParameter `json:"azureDevOps,omitempty"`
	AzureDevOpsBoards       InformationParameter `json:"azureDevOpsBoards,omitempty"`
}

type License any

type LicenseLimits struct {
	UserCount                   int  `json:"userCount,omitempty"`
	BotUserCount                int  `json:"botUserCount,omitempty"`
	RepositoryCount             int  `json:"repositoryCount,omitempty"`
	ProductCount                int  `json:"productCount,omitempty"`
	TaskTrackerIntegrationCount int  `json:"taskTrackerIntegrationCount,omitempty"`
	TaskTrackerIssueCount       int  `json:"taskTrackerIssueCount,omitempty"`
	ScannerIntegrationCount     int  `json:"scannerIntegrationCount,omitempty"`
	ArtifactIntegrationCount    int  `json:"artifactIntegrationCount,omitempty"`
	ImportIntegrationCount      int  `json:"importIntegrationCount,omitempty"`
	AIIntegrationCount          int  `json:"aiIntegrationCount,omitempty"`
	ScanPipelinesComparison     bool `json:"scanPipelinesComparison,omitempty"`
}

type ListUUID []UUID

type LoginRequest any

type MFAConfirmation any

type NameBasedBySeverity struct {
	NameBasedCardValue
	BySeverity
}

type NameBasedCardValue any

type NameWithEntityBasedCardValue struct {
	NameBasedCardValue
	Entity Entity `json:"entity,omitempty"`
}

type NotAccessibleByLicense any

type OIDCConfig any

type OIDCDefaults any

type OIDCGroupMapping any

type OIDCProvider string

const (
	OIDCProviderKeycloak OIDCProvider = "keycloak"
)

type OIDCRedirect any

type OIDCTestConnectRequest any

type OIDCTestConnectResponse any

type OIDCTestConnectURL any

type OldIssuesCardData struct {
	SimpleCardValue
	BySeverity
}

type OpenedIssuesCardData struct {
	SimpleCardValue
	BySeverity
}

type OrchestrationRunBatchResult any

type OrchestrationRunMode string

const (
	OrchestrationRunModeCONCURRENT OrchestrationRunMode = "CONCURRENT"
	OrchestrationRunModeCHAIN      OrchestrationRunMode = "CHAIN"
)

type OrchestrationRunResult any

type OrchestrationScanDetail struct {
	OrchestrationScan
	Profile OrchestrationScanProfile `json:"profile,omitempty"`
}

type OrchestrationScanProfileRelation any

type OrchestrationScanProfileRelationType string

const (
	OrchestrationScanProfileRelationTypeREPOSITORY OrchestrationScanProfileRelationType = "REPOSITORY"
	OrchestrationScanProfileRelationTypePRODUCT    OrchestrationScanProfileRelationType = "PRODUCT"
)

type OrchestrationScanProfile any

type OrchestrationScan struct {
	ID               UUID                           `json:"id,omitempty"`
	Status           OrchestrationScanStatus        `json:"status,omitempty"`
	TargetType       OrchestrationScanTargetType    `json:"targetType,omitempty"`
	TargetID         UUID                           `json:"targetID,omitempty"`
	SubTargetType    OrchestrationScanSubTargetType `json:"subTargetType,omitempty"`
	Created          string                         `json:"created,omitempty"`
	Updated          string                         `json:"updated,omitempty"`
	BackgroundTaskID UUID                           `json:"backgroundTaskID,omitempty"`
	ErrorMessage     string                         `json:"error_message,omitempty"`
}

type OrchestrationScanStatus string

const (
	OrchestrationScanStatusPENDING  OrchestrationScanStatus = "PENDING"
	OrchestrationScanStatusPROGRESS OrchestrationScanStatus = "PROGRESS"
	OrchestrationScanStatusSUCCESS  OrchestrationScanStatus = "SUCCESS"
	OrchestrationScanStatusFAILED   OrchestrationScanStatus = "FAILED"
)

type OrchestrationScanSubTargetType string

const (
	OrchestrationScanSubTargetTypeREPOSITORYBRANCH OrchestrationScanSubTargetType = "REPOSITORY_BRANCH"
)

type OrchestrationScanTargetType string

const (
	OrchestrationScanTargetTypeREPOSITORY      OrchestrationScanTargetType = "REPOSITORY"
	OrchestrationScanTargetTypePRODUCT         OrchestrationScanTargetType = "PRODUCT"
	OrchestrationScanTargetTypeARTIFACTVERSION OrchestrationScanTargetType = "ARTIFACT_VERSION"
	OrchestrationScanTargetTypeASSET           OrchestrationScanTargetType = "ASSET"
)

type OrchestrationScanUnavailabilityReason string

const (
	OrchestrationScanUnavailabilityReasonSCANNERINTEGRATIONUNAVAILABLE     OrchestrationScanUnavailabilityReason = "SCANNER_INTEGRATION_UNAVAILABLE"
	OrchestrationScanUnavailabilityReasonSCANNERINTEGRATIONLICENSEERROR    OrchestrationScanUnavailabilityReason = "SCANNER_INTEGRATION_LICENSE_ERROR"
	OrchestrationScanUnavailabilityReasonSCANNERINTEGRATIONCLIREQUIRED     OrchestrationScanUnavailabilityReason = "SCANNER_INTEGRATION_CLI_REQUIRED"
	OrchestrationScanUnavailabilityReasonSCANNERINAPPLICABLEFORCONTEXT     OrchestrationScanUnavailabilityReason = "SCANNER_INAPPLICABLE_FOR_CONTEXT"
	OrchestrationScanUnavailabilityReasonSCANNERSETTINGSERROR              OrchestrationScanUnavailabilityReason = "SCANNER_SETTINGS_ERROR"
	OrchestrationScanUnavailabilityReasonSCANNERDISABLED                   OrchestrationScanUnavailabilityReason = "SCANNER_DISABLED"
	OrchestrationScanUnavailabilityReasonSCANTARGETINTEGRATIONLICENSEERROR OrchestrationScanUnavailabilityReason = "SCAN_TARGET_INTEGRATION_LICENSE_ERROR"
	OrchestrationScanUnavailabilityReasonSCANTARGETINTEGRATIONNOTFOUND     OrchestrationScanUnavailabilityReason = "SCAN_TARGET_INTEGRATION_NOT_FOUND"
	OrchestrationScanUnavailabilityReasonSCANTARGETINTEGRATIONUNAVAILABLE  OrchestrationScanUnavailabilityReason = "SCAN_TARGET_INTEGRATION_UNAVAILABLE"
	OrchestrationScanUnavailabilityReasonNOPERMISSIONSACCESS               OrchestrationScanUnavailabilityReason = "NO_PERMISSIONS_ACCESS"
	OrchestrationScanUnavailabilityReasonNOAVAILABLESCANNERS               OrchestrationScanUnavailabilityReason = "NO_AVAILABLE_SCANNERS"
	OrchestrationScanUnavailabilityReasonNOAVAILABLETARGET                 OrchestrationScanUnavailabilityReason = "NO_AVAILABLE_TARGET"
)

type OrchestrationScannerConfigDelete any

type OrchestrationScannerConfigFileInfo any

type OrchestrationScannerConfigFileParams any

type OrchestrationScannerEnvVar any

type OrchestrationScannerInfoEnvVar struct {
	OrchestrationScannerEnvVar
	Required bool `json:"required,omitempty"`
}

type OrchestrationScannerInfo any

type OrchestrationScannerParam any

type OrchestrationScanner any

type Parameter any

type ParserInList any

type Parser string

const (
	ParserAPPSCREENERDAST         Parser = "APP_SCREENER_DAST"
	ParserAPPSCREENERSAST         Parser = "APP_SCREENER_SAST"
	ParserAPPSCREENERSCA          Parser = "APP_SCREENER_SCA"
	ParserSEMGREP                 Parser = "SEMGREP"
	ParserTRIVYFS                 Parser = "TRIVY_FS"
	ParserTRIVYIMAGE              Parser = "TRIVY_IMAGE"
	ParserTRUFFLEHOG              Parser = "TRUFFLEHOG"
	ParserTRUFFLEHOG3             Parser = "TRUFFLEHOG3"
	ParserESLINT                  Parser = "ESLINT"
	ParserGRYPE                   Parser = "GRYPE"
	ParserLUNTRY                  Parser = "LUNTRY"
	ParserNUCLEI                  Parser = "NUCLEI"
	ParserDEPENDENCYCHECK         Parser = "DEPENDENCY_CHECK"
	ParserGITLABCONTAINERSECURITY Parser = "GITLAB_CONTAINER_SECURITY"
	ParserGITLABDAST              Parser = "GITLAB_DAST"
	ParserGITLABDASTZAP           Parser = "GITLAB_DAST_ZAP"
	ParserGITLABDEPENDENCY        Parser = "GITLAB_DEPENDENCY"
	ParserGITLABSAST              Parser = "GITLAB_SAST"
	ParserGITLABSASTBRAKEMAN      Parser = "GITLAB_SAST_BRAKEMAN"
	ParserGITLABSASTFLAWFINDER    Parser = "GITLAB_SAST_FLAWFINDER"
	ParserGITLABSASTKUBESEC       Parser = "GITLAB_SAST_KUBESEC"
	ParserGITLABSASTMOBSF         Parser = "GITLAB_SAST_MOBSF"
	ParserGITLABSASTNODEJS        Parser = "GITLAB_SAST_NODEJS"
	ParserGITLABSASTPHPCS         Parser = "GITLAB_SAST_PHPCS"
	ParserGITLABSASTPMD           Parser = "GITLAB_SAST_PMD"
	ParserGITLABSASTSEMGREP       Parser = "GITLAB_SAST_SEMGREP"
	ParserGITLABSASTSOBELOW       Parser = "GITLAB_SAST_SOBELOW"
	ParserGITLABSASTSPOTBUGS      Parser = "GITLAB_SAST_SPOTBUGS"
	ParserGITLABSECRETDETECT      Parser = "GITLAB_SECRET_DETECT"
	ParserGITLEAKS                Parser = "GITLEAKS"
	ParserCHECKMARX               Parser = "CHECKMARX"
	ParserSARIF                   Parser = "SARIF"
	ParserSARIFCHECKOV            Parser = "SARIF_CHECKOV"
	ParserSARIFCODEQL             Parser = "SARIF_CODEQL"
	ParserSARIFCODESCORING        Parser = "SARIF_CODE_SCORING"
	ParserSARIFGITLEAKS           Parser = "SARIF_GITLEAKS"
	ParserSARIFGRYPE              Parser = "SARIF_GRYPE"
	ParserSARIFKICS               Parser = "SARIF_KICS"
	ParserSARIFPVSSTUDIO          Parser = "SARIF_PVS_STUDIO"
	ParserSARIFSASTAV             Parser = "SARIF_SASTAV"
	ParserSARIFSVACER             Parser = "SARIF_SVACER"
	ParserSBOM                    Parser = "SBOM"
	ParserSBOMCODESCORING         Parser = "SBOM_CODE_SCORING"
	ParserSBOMAPPSECTRACK         Parser = "SBOM_APPSEC_TRACK"
	ParserSNYKCODE                Parser = "SNYK_CODE"
	ParserSNYKSCA                 Parser = "SNYK_SCA"
	ParserSPOTBUGS                Parser = "SPOTBUGS"
	ParserPTAI                    Parser = "PT_AI"
	ParserPVSSTUDIO               Parser = "PVS_STUDIO"
	ParserSONARQUBE               Parser = "SONAR_QUBE"
	ParserZAP                     Parser = "ZAP"
	ParserACUNETIX                Parser = "ACUNETIX"
	ParserBURPSUITE               Parser = "BURP_SUITE"
	ParserVAMPY                   Parser = "VAMPY"
	ParserMANUAL                  Parser = "MANUAL"
)

type ParserSource string

const (
	ParserSourceSCANNER   ParserSource = "SCANNER"
	ParserSourceMANUAL    ParserSource = "MANUAL"
	ParserSourceSUBPARSER ParserSource = "SUB_PARSER"
)

type PermissionCategorySlug string

const (
	PermissionCategorySlugREPOSITORY      PermissionCategorySlug = "REPOSITORY"
	PermissionCategorySlugPRODUCT         PermissionCategorySlug = "PRODUCT"
	PermissionCategorySlugCOMPONENT       PermissionCategorySlug = "COMPONENT"
	PermissionCategorySlugASSET           PermissionCategorySlug = "ASSET"
	PermissionCategorySlugSCANISSUE       PermissionCategorySlug = "SCAN_ISSUE"
	PermissionCategorySlugBOTUSER         PermissionCategorySlug = "BOT_USER"
	PermissionCategorySlugQUALITYGATE     PermissionCategorySlug = "QUALITY_GATE"
	PermissionCategorySlugORCHESTRATION   PermissionCategorySlug = "ORCHESTRATION"
	PermissionCategorySlugCOMMENT         PermissionCategorySlug = "COMMENT"
	PermissionCategorySlugARTIFACT        PermissionCategorySlug = "ARTIFACT"
	PermissionCategorySlugRULES           PermissionCategorySlug = "RULES"
	PermissionCategorySlugINTEGRATION     PermissionCategorySlug = "INTEGRATION"
	PermissionCategorySlugBACKGROUNDTASKS PermissionCategorySlug = "BACKGROUND_TASKS"
	PermissionCategorySlugLDAP            PermissionCategorySlug = "LDAP"
	PermissionCategorySlugSMTP            PermissionCategorySlug = "SMTP"
	PermissionCategorySlugISSUESTATUS     PermissionCategorySlug = "ISSUE_STATUS"
	PermissionCategorySlugSYSTEMSETTINGS  PermissionCategorySlug = "SYSTEM_SETTINGS"
	PermissionCategorySlugSLA             PermissionCategorySlug = "SLA"
	PermissionCategorySlugLICENSE         PermissionCategorySlug = "LICENSE"
	PermissionCategorySlugUSERS           PermissionCategorySlug = "USERS"
	PermissionCategorySlugAI              PermissionCategorySlug = "AI"
)

type PermissionSetUserGroups struct {
	Permissions []PermissionUserGroup `json:"permissions,omitempty"`
}

type PermissionSetUsers struct {
	Permissions []PermissionUser `json:"permissions,omitempty"`
}

type PermissionUnsetUserGroups struct {
	UserGroups []UUID `json:"user_groups,omitempty"`
}

type PermissionUnsetUsers struct {
	Users []UUID `json:"users,omitempty"`
}

type PermissionUserGroup struct {
	GroupID UUID   `json:"groupID,omitempty"`
	Level   string `json:"level,omitempty"`
}

type PermissionUser struct {
	UserID UUID   `json:"userID,omitempty"`
	Level  string `json:"level,omitempty"`
}

type ProductChangeRelations any

type ProductCriticality string

const (
	ProductCriticalityLOW      ProductCriticality = "LOW"
	ProductCriticalityMEDIUM   ProductCriticality = "MEDIUM"
	ProductCriticalityHIGH     ProductCriticality = "HIGH"
	ProductCriticalityCRITICAL ProductCriticality = "CRITICAL"
)

type ProductDetail struct {
	Product
	RepositoryIDs         []UUID                   `json:"repositoryIDs,omitempty"`
	Integration           SpaceIntegration         `json:"integration,omitempty"`
	ScanProfile           OrchestrationScanProfile `json:"scanProfile,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail  `json:"lastOrchestrationScan,omitempty"`
	LastScanPipeline      ScanPipeline             `json:"lastScanPipeline,omitempty"`
}

type ProductEdit any

type ProductPermissionValue string

const (
	ProductPermissionValuePRODUCTEDIT              ProductPermissionValue = "PRODUCT_EDIT"
	ProductPermissionValuePRODUCTREAD              ProductPermissionValue = "PRODUCT_READ"
	ProductPermissionValuePRODUCTSETTINGSCHANGE    ProductPermissionValue = "PRODUCT_SETTINGS_CHANGE"
	ProductPermissionValuePRODUCTSTATSREAD         ProductPermissionValue = "PRODUCT_STATS_READ"
	ProductPermissionValuePRODUCTUSERSREAD         ProductPermissionValue = "PRODUCT_USERS_READ"
	ProductPermissionValueSCANUPLOAD               ProductPermissionValue = "SCAN_UPLOAD"
	ProductPermissionValueSCANISSUEREAD            ProductPermissionValue = "SCAN_ISSUE_READ"
	ProductPermissionValueSCANISSUECREATE          ProductPermissionValue = "SCAN_ISSUE_CREATE"
	ProductPermissionValueSCANISSUEEDIT            ProductPermissionValue = "SCAN_ISSUE_EDIT"
	ProductPermissionValueASSETREAD                ProductPermissionValue = "ASSET_READ"
	ProductPermissionValueASSETCREATE              ProductPermissionValue = "ASSET_CREATE"
	ProductPermissionValueASSETEDIT                ProductPermissionValue = "ASSET_EDIT"
	ProductPermissionValueORCHESTRATIONRUN         ProductPermissionValue = "ORCHESTRATION_RUN"
	ProductPermissionValueORCHESTRATIONPROFILEREAD ProductPermissionValue = "ORCHESTRATION_PROFILE_READ"
	ProductPermissionValueARTIFACTREAD             ProductPermissionValue = "ARTIFACT_READ"
	ProductPermissionValueARTIFACTEDIT             ProductPermissionValue = "ARTIFACT_EDIT"
	ProductPermissionValueSCANPIPELINEREAD         ProductPermissionValue = "SCAN_PIPELINE_READ"
)

type Product any

type ProductScanIssuesStat struct {
	BySeverity []ScanIssueBySeverityStat `json:"bySeverity,omitempty"`
	ByStatus   []ScanIssueByStatusStat   `json:"byStatus,omitempty"`
	Counts     ScanIssueCountsStat       `json:"counts,omitempty"`
}

type ProductShort any

type ProductsCountCardData struct {
	SimpleCardValue
}

type QualityGateCondition any

type QualityGateConditionResult struct {
	QualityGateCondition
	CurrentValue int                        `json:"currentValue,omitempty"`
	Status       QualityGateConditionStatus `json:"status,omitempty"`
}

type QualityGateConditionStatus string

const (
	QualityGateConditionStatusSKIPPED QualityGateConditionStatus = "SKIPPED"
	QualityGateConditionStatusFAILED  QualityGateConditionStatus = "FAILED"
	QualityGateConditionStatusPASSED  QualityGateConditionStatus = "PASSED"
)

type QualityGateCriticality string

const (
	QualityGateCriticalityLOW      QualityGateCriticality = "LOW"
	QualityGateCriticalityMEDIUM   QualityGateCriticality = "MEDIUM"
	QualityGateCriticalityHIGH     QualityGateCriticality = "HIGH"
	QualityGateCriticalityCRITICAL QualityGateCriticality = "CRITICAL"
)

type QualityGateEdit any

type QualityGateForRelation struct {
	ID         UUID                             `json:"id,omitempty"`
	Name       string                           `json:"name,omitempty"`
	Slug       string                           `json:"slug,omitempty"`
	LastResult QualityGateForRelationLastResult `json:"lastResult,omitempty"`
}

type QualityGateForRelationLastResult struct {
	ID             *UUID                   `json:"id,omitempty"`
	Status         QualityGateResultStatus `json:"status,omitempty"`
	LastCalculated *string                 `json:"lastCalculated,omitempty"`
}

type QualityGateRelation any

type QualityGateRelationType string

const (
	QualityGateRelationTypePRODUCT    QualityGateRelationType = "PRODUCT"
	QualityGateRelationTypeREPOSITORY QualityGateRelationType = "REPOSITORY"
)

type QualityGateResultDetails struct {
	ID             UUID                            `json:"id,omitempty"`
	Status         QualityGateResultStatus         `json:"status,omitempty"`
	LastCalculated *string                         `json:"lastCalculated,omitempty"`
	Relation       QualityGateRelation             `json:"relation,omitempty"`
	Details        QualityGateResultDetailsDetails `json:"details,omitempty"`
}

type QualityGateResultDetailsDetails struct {
	Slug         string                       `json:"slug,omitempty"`
	Name         string                       `json:"name,omitempty"`
	Criticality  QualityGateCriticality       `json:"criticality,omitempty"`
	Conditions   []QualityGateConditionResult `json:"conditions,omitempty"`
	ErrorMessage string                       `json:"errorMessage,omitempty"`
}

type QualityGateResult any

type QualityGateResultStatus string

const (
	QualityGateResultStatusINPROGRESS QualityGateResultStatus = "IN_PROGRESS"
	QualityGateResultStatusTIMEDOUT   QualityGateResultStatus = "TIMED_OUT"
	QualityGateResultStatusSKIPPED    QualityGateResultStatus = "SKIPPED"
	QualityGateResultStatusFAILED     QualityGateResultStatus = "FAILED"
	QualityGateResultStatusPASSED     QualityGateResultStatus = "PASSED"
	QualityGateResultStatusERROR      QualityGateResultStatus = "ERROR"
)

type QualityGate any

type QualityGateSubRelationType string

const (
	QualityGateSubRelationTypeREPOSITORYBRANCH QualityGateSubRelationType = "REPOSITORY_BRANCH"
)

type ReportFileFormat string

const (
	ReportFileFormatJSON ReportFileFormat = "JSON"
	ReportFileFormatPDF  ReportFileFormat = "PDF"
)

type Report struct {
	ID               UUID   `json:"id,omitempty"`
	FileFormat       any    `json:"fileFormat,omitempty"`
	Status           any    `json:"status,omitempty"`
	CreatedByID      UUID   `json:"createdById,omitempty"`
	BackgroundTaskID UUID   `json:"backgroundTaskId,omitempty"`
	ErrorMessage     string `json:"errorMessage,omitempty"`
	Created          string `json:"created,omitempty"`
	Updated          string `json:"updated,omitempty"`
}

type ReportStatus string

const (
	ReportStatusSUCCESS  ReportStatus = "SUCCESS"
	ReportStatusPENDING  ReportStatus = "PENDING"
	ReportStatusPROGRESS ReportStatus = "PROGRESS"
	ReportStatusFAILED   ReportStatus = "FAILED"
)

type RepositoriesCountCardData struct {
	SimpleCardValue
}

type RepositoryBranchList any

type RepositoryBranchListRepository struct {
	ID   UUID   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type RepositoryBranch any

type RepositoryBranchSyncRequest struct {
	Repository string `json:"repository"`
}

type RepositoryBranchSyncResponse struct {
	Repository            string                        `json:"repository,omitempty"`
	AddedCount            int                           `json:"addedCount,omitempty"`
	RemovedCount          int                           `json:"removedCount,omitempty"`
	SkippedCount          int                           `json:"skippedCount,omitempty"`
	ExternalBranchesCount int                           `json:"externalBranchesCount,omitempty"`
	LocalBranchesCount    int                           `json:"localBranchesCount,omitempty"`
	AddedBranches         []string                      `json:"addedBranches,omitempty"`
	RemovedBranches       []string                      `json:"removedBranches,omitempty"`
	SkippedBranches       []RepositoryBranchSyncSkipped `json:"skippedBranches,omitempty"`
}

type RepositoryBranchSyncSkipped struct {
	Slug   string `json:"slug,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type RepositoryCreate struct {
	Repository
	ProductID *UUID `json:"productID,omitempty"`
}

type RepositoryCriticality string

const (
	RepositoryCriticalityLOW      RepositoryCriticality = "LOW"
	RepositoryCriticalityMEDIUM   RepositoryCriticality = "MEDIUM"
	RepositoryCriticalityHIGH     RepositoryCriticality = "HIGH"
	RepositoryCriticalityCRITICAL RepositoryCriticality = "CRITICAL"
)

type RepositoryDetails struct {
	Repository
	Languages             map[string]any                `json:"languages,omitempty"`
	ExternalLink          RepositoryDetailsExternalLink `json:"externalLink,omitempty"`
	ScanProfile           OrchestrationScanProfile      `json:"scanProfile,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail       `json:"lastOrchestrationScan,omitempty"`
	LastScanPipeline      ScanPipeline                  `json:"lastScanPipeline,omitempty"`
}

type RepositoryDetailsExternalLink struct {
	Name               string                    `json:"name,omitempty"`
	URL                string                    `json:"url,omitempty"`
	IntegrationService IntegrationVCSServiceType `json:"integrationService,omitempty"`
	IntegrationState   IntegrationState          `json:"integrationState,omitempty"`
}

type RepositoryForExternalLink any

type RepositoryImportRequest struct {
	SpaceID                UUID     `json:"spaceID"`
	ProductID              UUID     `json:"productID,omitempty"`
	IntegrationID          string   `json:"integrationID"`
	PopulateRepositoryData bool     `json:"populateRepositoryData,omitempty"`
	ExternalIDs            []string `json:"externalIDs"`
}

type RepositoryImportResponse struct {
	Repositories []RepositoryImportResponseRepositories `json:"repositories,omitempty"`
}

type RepositoryImportResponseRepositories struct {
	ID   UUID   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type RepositoryLink struct {
	SpaceID                UUID   `json:"spaceID"`
	IntegrationID          string `json:"integrationID"`
	PopulateRepositoryData bool   `json:"populateRepositoryData,omitempty"`
	ExternalID             string `json:"externalID"`
}

type RepositoryPatch struct {
	Criticality   RepositoryCriticality `json:"criticality,omitempty"`
	Name          string                `json:"name,omitempty"`
	URL           string                `json:"url,omitempty"`
	Description   string                `json:"description,omitempty"`
	QualityGateID *UUID                 `json:"qualityGateID,omitempty"`
	ScanProfileID *UUID                 `json:"scanProfileID,omitempty"`
	DefaultBranch *string               `json:"defaultBranch,omitempty"`
}

type RepositoryPermissionValue string

const (
	RepositoryPermissionValueREPOSITORYEDIT           RepositoryPermissionValue = "REPOSITORY_EDIT"
	RepositoryPermissionValueREPOSITORYREAD           RepositoryPermissionValue = "REPOSITORY_READ"
	RepositoryPermissionValueREPOSITORYSETTINGSCHANGE RepositoryPermissionValue = "REPOSITORY_SETTINGS_CHANGE"
	RepositoryPermissionValueREPOSITORYSTATSREAD      RepositoryPermissionValue = "REPOSITORY_STATS_READ"
	RepositoryPermissionValueREPOSITORYUSERSREAD      RepositoryPermissionValue = "REPOSITORY_USERS_READ"
	RepositoryPermissionValueSCANISSUEREAD            RepositoryPermissionValue = "SCAN_ISSUE_READ"
	RepositoryPermissionValueSCANISSUECREATE          RepositoryPermissionValue = "SCAN_ISSUE_CREATE"
	RepositoryPermissionValueSCANISSUEEDIT            RepositoryPermissionValue = "SCAN_ISSUE_EDIT"
	RepositoryPermissionValueSCANUPLOAD               RepositoryPermissionValue = "SCAN_UPLOAD"
	RepositoryPermissionValueORCHESTRATIONRUN         RepositoryPermissionValue = "ORCHESTRATION_RUN"
	RepositoryPermissionValueORCHESTRATIONPROFILEREAD RepositoryPermissionValue = "ORCHESTRATION_PROFILE_READ"
	RepositoryPermissionValueSCANPIPELINEREAD         RepositoryPermissionValue = "SCAN_PIPELINE_READ"
)

type Repository struct {
	ID               UUID                        `json:"id,omitempty"`
	Type             RepositoryType              `json:"type,omitempty"`
	Slug             string                      `json:"slug"`
	Criticality      RepositoryCriticality       `json:"criticality,omitempty"`
	Name             string                      `json:"name"`
	Created          string                      `json:"created,omitempty"`
	LastScan         string                      `json:"lastScan,omitempty"`
	Description      string                      `json:"description,omitempty"`
	URL              string                      `json:"url"`
	RiskScore        *int                        `json:"riskScore,omitempty"`
	ScanIssuesStat   ProductScanIssuesStat       `json:"scanIssuesStat,omitempty"`
	LastAggregated   string                      `json:"lastAggregated,omitempty"`
	Users            []DetailPermissionUser      `json:"users,omitempty"`
	Permissions      []RepositoryPermissionValue `json:"permissions,omitempty"`
	Branches         []RepositoryBranch          `json:"branches,omitempty"`
	DefaultBranch    string                      `json:"defaultBranch,omitempty"`
	QualityGate      QualityGateForRelation      `json:"qualityGate,omitempty"`
	TitleTranslateID *string                     `json:"titleTranslateID,omitempty"`
}

type RepositoryShort any

type RepositoryType string

const (
	RepositoryTypeSOURCECODE RepositoryType = "SOURCE_CODE"
	RepositoryTypeBINARY     RepositoryType = "BINARY"
)

type RuleCreate struct {
	BaseRule
	EventType   RulesEngineEventType `json:"eventType,omitempty"`
	ScanIssueID *UUID                `json:"scanIssueID,omitempty"`
	RunRule     bool                 `json:"runRule,omitempty"`
}

type RulePassedInstances any

type RulePatch struct {
	BaseRule
	RunRule bool `json:"runRule,omitempty"`
}

type RuleRead struct {
	BaseRule
	EventType   RulesEngineEventType `json:"eventType,omitempty"`
	ScanIssueID *UUID                `json:"scanIssueID,omitempty"`
	Created     string               `json:"created,omitempty"`
	Updated     string               `json:"updated,omitempty"`
	Warnings    []RuleWarning        `json:"warnings,omitempty"`
}

type RuleShort any

type RuleWarning struct {
	Type    RuleWarningType `json:"type"`
	Message string          `json:"message"`
}

type RuleWarningType string

const (
	RuleWarningTypeAction RuleWarningType = "action"
)

type RulesEngineCreationSource string

const (
	RulesEngineCreationSourceMANUAL    RulesEngineCreationSource = "MANUAL"
	RulesEngineCreationSourceSCANISSUE RulesEngineCreationSource = "SCAN_ISSUE"
)

type RulesEngineEventType string

const (
	RulesEngineEventTypeSCANISSUECHANGED       RulesEngineEventType = "SCAN_ISSUE_CHANGED"
	RulesEngineEventTypeVCSREPOSITORYCHANGED   RulesEngineEventType = "VCS_REPOSITORY_CHANGED"
	RulesEngineEventTypeVCSMERGEREQUESTCHANGED RulesEngineEventType = "VCS_MERGE_REQUEST_CHANGED"
)

type RulesSupportedEvents any

type SBOMExport any

type SBOMImport any

type SLASetting any

type SLASettingWithDefault any

type ScanIssueAIRequestApply any

type ScanIssueAIRequestCancel any

type ScanIssueAIRequestResponse any

type ScanIssueAIRequestStatus string

const (
	ScanIssueAIRequestStatusPENDING   ScanIssueAIRequestStatus = "PENDING"
	ScanIssueAIRequestStatusCOMPLETED ScanIssueAIRequestStatus = "COMPLETED"
	ScanIssueAIRequestStatusFAILED    ScanIssueAIRequestStatus = "FAILED"
	ScanIssueAIRequestStatusCANCELLED ScanIssueAIRequestStatus = "CANCELLED"
	ScanIssueAIRequestStatusAPPLIED   ScanIssueAIRequestStatus = "APPLIED"
)

type ScanIssueAIRequestType string

const (
	ScanIssueAIRequestTypeSCANISSUEFALSEPOSITIVERATE ScanIssueAIRequestType = "SCAN_ISSUE_FALSE_POSITIVE_RATE"
	ScanIssueAIRequestTypeSCANISSUEDESCRIPTION       ScanIssueAIRequestType = "SCAN_ISSUE_DESCRIPTION"
	ScanIssueAIRequestTypeSCANISSUESOLUTION          ScanIssueAIRequestType = "SCAN_ISSUE_SOLUTION"
	ScanIssueAIRequestTypeSCANISSUECODEEXPLAIN       ScanIssueAIRequestType = "SCAN_ISSUE_CODE_EXPLAIN"
	ScanIssueAIRequestTypeSCANISSUECODEVULN          ScanIssueAIRequestType = "SCAN_ISSUE_CODE_VULN"
	ScanIssueAIRequestTypeAICHECKCONNECTION          ScanIssueAIRequestType = "AI_CHECK_CONNECTION"
	ScanIssueAIRequestTypeAIGETMODELS                ScanIssueAIRequestType = "AI_GET_MODELS"
)

type ScanIssueBySeverityStat any

type ScanIssueByStatusStat any

type ScanIssueChangeEntry any

type ScanIssueChangeEntrySource struct {
	Type                      IssueChangeSource            `json:"type,omitempty"`
	ScanUpload                IssueStatusChangeScanUpload  `json:"scanUpload,omitempty"`
	User                      IssueStatusChangeUser        `json:"user,omitempty"`
	Integration               IssueStatusChangeIntegration `json:"integration,omitempty"`
	Rule                      IssueStatusChangeRule        `json:"rule,omitempty"`
	ScheduledStatusChangeInfo ScheduledStatus              `json:"scheduledStatusChangeInfo,omitempty"`
	SourceScanIssue           ScanIssueShortWithBranch     `json:"sourceScanIssue,omitempty"`
	AIModelName               *string                      `json:"aiModelName,omitempty"`
	FPClassifierSnapshot      FPClassifierSnapshotShort    `json:"fpClassifierSnapshot,omitempty"`
}

type ScanIssueCountsStat any

type ScanIssueCreateResponse any

type ScanIssueCreate any

type ScanIssueDelete any

type ScanIssueDetail struct {
	ScanIssue
	Data                 map[string]any                      `json:"data,omitempty"`
	Assets               []Asset                             `json:"assets,omitempty"`
	Description          string                              `json:"description,omitempty"`
	CodeSnippet          string                              `json:"codeSnippet,omitempty"`
	Solution             string                              `json:"solution,omitempty"`
	References           []string                            `json:"references,omitempty"`
	ArtifactVersion      ArtifactVersionShort                `json:"artifactVersion,omitempty"`
	Product              ProductShort                        `json:"product,omitempty"`
	ComponentVersion     ComponentVersionShort               `json:"componentVersion,omitempty"`
	FilePathVCSLink      *string                             `json:"filePathVCSLink,omitempty"`
	CommitSHA            ScanIssueDetailCommitSHA            `json:"commitSha,omitempty"`
	GitBlame             ScanIssueGitBlame                   `json:"gitBlame,omitempty"`
	Duplicates           ScanIssueDetailDuplicates           `json:"duplicates,omitempty"`
	Comments             ScanIssueDetailDuplicates           `json:"comments,omitempty"`
	AssociatedScanIssues ScanIssueDetailAssociatedScanIssues `json:"associatedScanIssues,omitempty"`
	AIRequests           []ScanIssueAIRequestResponse        `json:"aiRequests,omitempty"`
	AIApplied            ScanIssueDetailAIApplied            `json:"aiApplied,omitempty"`
	FPScore              FPScoreDetail                       `json:"fpScore,omitempty"`
}

type ScanIssueDetailAIApplied struct {
	FalsePositiveRate ScanIssueAIRequestResponse `json:"falsePositiveRate,omitempty"`
	Description       ScanIssueAIRequestResponse `json:"description,omitempty"`
	Solution          ScanIssueAIRequestResponse `json:"solution,omitempty"`
	CodeExplain       ScanIssueAIRequestResponse `json:"codeExplain,omitempty"`
	CodeVuln          ScanIssueAIRequestResponse `json:"codeVuln,omitempty"`
}

type ScanIssueDetailAssociatedScanIssues struct {
	Count int `json:"count,omitempty"`
}

type ScanIssueDetailCommitSHA struct {
	ByScanner *string  `json:"byScanner,omitempty"`
	LastSeen  *string  `json:"lastSeen,omitempty"`
	AllSeen   []string `json:"allSeen,omitempty"`
}

type ScanIssueDetailDuplicates struct {
	Count  int    `json:"count,omitempty"`
	Latest string `json:"latest,omitempty"`
}

type ScanIssueDuplicatesSearchResult any

type ScanIssueDuplicatesSearchResultScanIssues struct {
	ID     UUID   `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Status string `json:"status,omitempty"`
}

type ScanIssueExport any

type ScanIssueFPScoreAIModelFilter struct {
	IntegrationID UUID   `json:"integrationID"`
	AIModel       string `json:"aiModel"`
}

type ScanIssueFPScoreFilter struct {
	ValueFrom *int  `json:"valueFrom,omitempty"`
	ValueTo   *int  `json:"valueTo,omitempty"`
	NoValue   *bool `json:"noValue,omitempty"`
}

type ScanIssueFilter any

type ScanIssueForExportResponse struct {
	SummarizedIssue UnitedScanIssueForExport   `json:"summarizedIssue"`
	ExportingIssues []SingleScanIssueForExport `json:"exportingIssues"`
}

type ScanIssueForExport struct {
	Name    string  `json:"name"`
	Summary string  `json:"summary"`
	SLA     *string `json:"sla,omitempty"`
}

type ScanIssueGitBlame any

type ScanIssueGitBlameCommitAuthor struct {
	ID    UUID   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type ScanIssuePermissionValue string

const (
	ScanIssuePermissionValueSCANISSUEREAD               ScanIssuePermissionValue = "SCAN_ISSUE_READ"
	ScanIssuePermissionValueSCANISSUEEDIT               ScanIssuePermissionValue = "SCAN_ISSUE_EDIT"
	ScanIssuePermissionValueSCANUPLOADDOWNLOAD          ScanIssuePermissionValue = "SCAN_UPLOAD_DOWNLOAD"
	ScanIssuePermissionValueSCANISSUESENDTOINTEGRATIONS ScanIssuePermissionValue = "SCAN_ISSUE_SEND_TO_INTEGRATIONS"
	ScanIssuePermissionValueCOMMENTCREATE               ScanIssuePermissionValue = "COMMENT_CREATE"
	ScanIssuePermissionValueSCANISSUESEVERITYCHANGE     ScanIssuePermissionValue = "SCAN_ISSUE_SEVERITY_CHANGE"
	ScanIssuePermissionValueSCANISSUEDELETE             ScanIssuePermissionValue = "SCAN_ISSUE_DELETE"
)

type ScanIssue any

type ScanIssues struct {
	ScanIssueType ScanIssueType    `json:"scanIssueType,omitempty"`
	Schema        ScanIssuesSchema `json:"schema,omitempty"`
}

type ScanIssuesSchema struct {
	Name          string                          `json:"name,omitempty"`
	DisplayName   string                          `json:"displayName,omitempty"`
	Type          string                          `json:"type,omitempty"`
	Hidden        bool                            `json:"hidden,omitempty"`
	IsList        bool                            `json:"isList,omitempty"`
	Metadata      ScanIssuesSchemaMetadata        `json:"metadata,omitempty"`
	AllowedValues []ScanIssuesSchemaAllowedValues `json:"allowedValues,omitempty"`
	Required      bool                            `json:"required,omitempty"`
}

type ScanIssuesSchemaAllowedValues struct {
	Value        string `json:"value,omitempty"`
	DisplayValue string `json:"displayValue,omitempty"`
}

type ScanIssuesSchemaMetadata struct {
	Location   string `json:"location,omitempty"`
	Filterable bool   `json:"filterable,omitempty"`
}

type ScanIssueShort any

type ScanIssueShortWithBranch struct {
	ScanIssueShort
	RepositoryBranch RepositoryBranch `json:"repositoryBranch,omitempty"`
}

type ScanIssueStatusChangeEntry any

type ScanIssueStatusChangeEntrySource struct {
	Type                      IssueStatusChangeSource      `json:"type,omitempty"`
	ScanUpload                IssueStatusChangeScanUpload  `json:"scanUpload,omitempty"`
	User                      IssueStatusChangeUser        `json:"user,omitempty"`
	Integration               IssueStatusChangeIntegration `json:"integration,omitempty"`
	Rule                      IssueStatusChangeRule        `json:"rule,omitempty"`
	ScheduledStatusChangeInfo ScheduledStatus              `json:"scheduledStatusChangeInfo,omitempty"`
	SourceScanIssue           ScanIssueShortWithBranch     `json:"sourceScanIssue,omitempty"`
}

type ScanIssueType string

const (
	ScanIssueTypeVULNERABILITY   ScanIssueType = "VULNERABILITY"
	ScanIssueTypeSOURCECODE      ScanIssueType = "SOURCE_CODE"
	ScanIssueTypeDEPENDENCYCHECK ScanIssueType = "DEPENDENCY_CHECK"
)

type ScanIssueWithFPScore struct {
	ScanIssue
	FPScore FPScore `json:"fpScore,omitempty"`
}

type ScanIssuesAIRequest any

type ScanIssuesUpdate struct {
	Issues              []UUID   `json:"issues,omitempty"`
	Status              string   `json:"status,omitempty"`
	Severity            Severity `json:"severity,omitempty"`
	StatusExpiryTime    string   `json:"statusExpiryTime,omitempty"`
	StatusChangeComment *string  `json:"statusChangeComment,omitempty"`
}

type ScanPipelineAggregatedData struct {
	ByFindingState NameBasedBySeverity `json:"by_finding_state,omitempty"`
}

type ScanPipelineCreateExtResponse struct {
	ID         UUID               `json:"id,omitempty"`
	ExternalID string             `json:"externalID,omitempty"`
	Status     ScanPipelineStatus `json:"status,omitempty"`
	Created    string             `json:"created,omitempty"`
}

type ScanPipelineCreateExt struct {
	Repository       string  `json:"repository,omitempty"`
	Product          string  `json:"product,omitempty"`
	RepositoryBranch string  `json:"repositoryBranch,omitempty"`
	Artifact         string  `json:"artifact,omitempty"`
	Asset            string  `json:"asset,omitempty"`
	SpaceID          *UUID   `json:"spaceID,omitempty"`
	Name             string  `json:"name,omitempty"`
	ExternalID       string  `json:"externalID,omitempty"`
	CommitSHA        *string `json:"commitSha,omitempty"`
}

type ScanPipelineDetailsExt struct {
	ID                   UUID                       `json:"id,omitempty"`
	Created              string                     `json:"created,omitempty"`
	Updated              string                     `json:"updated,omitempty"`
	Finished             *string                    `json:"finished,omitempty"`
	Error                *string                    `json:"error,omitempty"`
	Duration             int                        `json:"duration,omitempty"`
	ScanningTool         string                     `json:"scanningTool,omitempty"`
	AggregatedData       ScanPipelineAggregatedData `json:"aggregatedData,omitempty"`
	AggregationCompleted bool                       `json:"aggregationCompleted,omitempty"`
	Status               ScanPipelineStatus         `json:"status,omitempty"`
	Relation             ScanPipelineRelation       `json:"relation,omitempty"`
	ScanProfile          OrchestrationScanProfile   `json:"scanProfile,omitempty"`
	ExternalID           *string                    `json:"externalID,omitempty"`
	CommitSHA            CommitSHA                  `json:"commitSha,omitempty"`
}

type ScanPipelineDetails struct {
	ScanPipeline
	CommitSHA CommitSHA `json:"commitSha,omitempty"`
}

type ScanPipelineInitiator struct {
	ID          UUID                      `json:"id,omitempty"`
	Type        ScanPipelineInitiatorType `json:"type,omitempty"`
	Slug        string                    `json:"slug,omitempty"`
	DisplayName string                    `json:"displayName,omitempty"`
}

type ScanPipelineInitiatorType string

const (
	ScanPipelineInitiatorTypeUSER        ScanPipelineInitiatorType = "USER"
	ScanPipelineInitiatorTypeBOT         ScanPipelineInitiatorType = "BOT"
	ScanPipelineInitiatorTypeRULESENGINE ScanPipelineInitiatorType = "RULES_ENGINE"
	ScanPipelineInitiatorTypeSCHEDULER   ScanPipelineInitiatorType = "SCHEDULER"
)

type ScanPipelineJob struct {
	ID             UUID                       `json:"id,omitempty"`
	SpaceID        UUID                       `json:"spaceID,omitempty"`
	ScanPipelineID UUID                       `json:"scanPipelineID,omitempty"`
	Created        string                     `json:"created,omitempty"`
	Updated        string                     `json:"updated,omitempty"`
	Finished       string                     `json:"finished,omitempty"`
	Duration       int                        `json:"duration,omitempty"`
	Status         ScanPipelineJobStatus      `json:"status,omitempty"`
	Error          string                     `json:"error,omitempty"`
	ScanningTool   string                     `json:"scanningTool,omitempty"`
	CommitSHA      CommitSHA                  `json:"commitSha,omitempty"`
	Relation       ScanPipelineRelation       `json:"relation,omitempty"`
	ScanUpload     ScanUpload                 `json:"scanUpload,omitempty"`
	AggregatedData ScanPipelineAggregatedData `json:"aggregatedData,omitempty"`
}

type ScanPipelineJobStatus string

const (
	ScanPipelineJobStatusPENDING    ScanPipelineJobStatus = "PENDING"
	ScanPipelineJobStatusSCANNING   ScanPipelineJobStatus = "SCANNING"
	ScanPipelineJobStatusPROCESSING ScanPipelineJobStatus = "PROCESSING"
	ScanPipelineJobStatusERROR      ScanPipelineJobStatus = "ERROR"
	ScanPipelineJobStatusPARSING    ScanPipelineJobStatus = "PARSING"
	ScanPipelineJobStatusSUCCESS    ScanPipelineJobStatus = "SUCCESS"
	ScanPipelineJobStatusCANCELED   ScanPipelineJobStatus = "CANCELED"
	ScanPipelineJobStatusRETRYING   ScanPipelineJobStatus = "RETRYING"
)

type ScanPipelineListExtItem struct {
	ID                   UUID                     `json:"id,omitempty"`
	ExternalID           *string                  `json:"externalID,omitempty"`
	Status               ScanPipelineStatus       `json:"status,omitempty"`
	ScanningTool         string                   `json:"scanningTool,omitempty"`
	RelationType         ScanPipelineRelationType `json:"relationType,omitempty"`
	RelationID           UUID                     `json:"relationId,omitempty"`
	Created              string                   `json:"created,omitempty"`
	Finished             *string                  `json:"finished,omitempty"`
	AggregationCompleted bool                     `json:"aggregationCompleted,omitempty"`
}

type ScanPipelineListExtResponse struct {
	Items      []ScanPipelineListExtItem `json:"items,omitempty"`
	HasNext    bool                      `json:"hasNext,omitempty"`
	Offset     int                       `json:"offset,omitempty"`
	TotalCount int                       `json:"totalCount,omitempty"`
}

type ScanPipelineRelation any

type ScanPipelineRelationSubType string

const (
	ScanPipelineRelationSubTypeREPOSITORYBRANCH ScanPipelineRelationSubType = "REPOSITORY_BRANCH"
	ScanPipelineRelationSubTypeNONE             ScanPipelineRelationSubType = "NONE_"
)

type ScanPipelineRelationType string

const (
	ScanPipelineRelationTypeREPOSITORY      ScanPipelineRelationType = "REPOSITORY"
	ScanPipelineRelationTypePRODUCT         ScanPipelineRelationType = "PRODUCT"
	ScanPipelineRelationTypeARTIFACTVERSION ScanPipelineRelationType = "ARTIFACT_VERSION"
	ScanPipelineRelationTypeASSET           ScanPipelineRelationType = "ASSET"
)

type ScanPipeline struct {
	ID                   UUID                       `json:"id,omitempty"`
	SpaceID              UUID                       `json:"spaceID,omitempty"`
	ScanPipelineType     ScanPipelineType           `json:"scanPipelineType,omitempty"`
	Created              string                     `json:"created,omitempty"`
	Updated              string                     `json:"updated,omitempty"`
	Finished             string                     `json:"finished,omitempty"`
	Duration             int                        `json:"duration,omitempty"`
	Status               ScanPipelineStatus         `json:"status,omitempty"`
	Error                string                     `json:"error,omitempty"`
	ScanningTool         string                     `json:"scanningTool,omitempty"`
	Initiator            ScanPipelineInitiator      `json:"initiator,omitempty"`
	Relation             ScanPipelineRelation       `json:"relation,omitempty"`
	Engagement           Engagement                 `json:"engagement,omitempty"`
	OrchestrationScan    OrchestrationScan          `json:"orchestrationScan,omitempty"`
	AggregatedData       ScanPipelineAggregatedData `json:"aggregatedData,omitempty"`
	AggregationCompleted bool                       `json:"aggregationCompleted,omitempty"`
	Profile              OrchestrationScanProfile   `json:"profile,omitempty"`
}

type ScanPipelineStatus string

const (
	ScanPipelineStatusPENDING        ScanPipelineStatus = "PENDING"
	ScanPipelineStatusINPROGRESS     ScanPipelineStatus = "IN_PROGRESS"
	ScanPipelineStatusERROR          ScanPipelineStatus = "ERROR"
	ScanPipelineStatusSUCCESS        ScanPipelineStatus = "SUCCESS"
	ScanPipelineStatusPARTIALSUCCESS ScanPipelineStatus = "PARTIAL_SUCCESS"
)

type ScanPipelineStatusExt struct {
	ID                   UUID               `json:"id,omitempty"`
	ExternalID           *string            `json:"externalID,omitempty"`
	Status               ScanPipelineStatus `json:"status,omitempty"`
	ScanningTool         string             `json:"scanningTool,omitempty"`
	Created              string             `json:"created,omitempty"`
	Finished             *string            `json:"finished,omitempty"`
	AggregationCompleted bool               `json:"aggregationCompleted,omitempty"`
}

type ScanPipelineType string

const (
	ScanPipelineTypeINTERNAL ScanPipelineType = "INTERNAL"
	ScanPipelineTypeEXTERNAL ScanPipelineType = "EXTERNAL"
)

type ScanUploadRelation any

type ScanUpload any

type ScannerCategory string

const (
	ScannerCategoryDEFAULT     ScannerCategory = "DEFAULT"
	ScannerCategoryCUSTOM      ScannerCategory = "CUSTOM"
	ScannerCategoryINTEGRATION ScannerCategory = "INTEGRATION"
	ScannerCategoryEXTERNAL    ScannerCategory = "EXTERNAL"
)

type ScannerContentType string

const (
	ScannerContentTypeApplicationJSON ScannerContentType = "application/json"
	ScannerContentTypeTextXML         ScannerContentType = "text/xml"
)

type ScannerExternal any

type ScannerIntegrationCLIFileInfo any

type ScannerIntegrationCLI any

type ScannerIntegrationCLISource string

const (
	ScannerIntegrationCLISourceMANUAL        ScannerIntegrationCLISource = "MANUAL"
	ScannerIntegrationCLISourceFETCHONDEMAND ScannerIntegrationCLISource = "FETCH_ON_DEMAND"
)

type ScannerIntegrationInfoRequest any

type ScannerIntegrationInfoResponse any

type ScannerIntegrationMetaResponse any

type ScannerIntegrationPatch struct {
	ScannerIntegration
	CLISource     ScannerIntegrationCLISource `json:"cliSource,omitempty"`
	CLIFilePath   *string                     `json:"cliFilePath,omitempty"`
	CLIFileName   *string                     `json:"cliFileName,omitempty"`
	CLIVersion    *string                     `json:"cliVersion,omitempty"`
	CLIJDKVersion *string                     `json:"cliJDKVersion,omitempty"`
}

type ScannerIntegration any

type ScannerIntegrationTestConnection any

type ScannerIntegrationType string

const (
	ScannerIntegrationTypePtai        ScannerIntegrationType = "ptai"
	ScannerIntegrationTypeCodeScoring ScannerIntegrationType = "code_scoring"
	ScannerIntegrationTypeAppScreener ScannerIntegrationType = "app_screener"
	ScannerIntegrationTypeSastav      ScannerIntegrationType = "sastav"
	ScannerIntegrationTypeAppsecTrack ScannerIntegrationType = "appsec_track"
)

type Scanner string

const (
	ScannerSEMGREP                  Scanner = "SEMGREP"
	ScannerTRIVYFS                  Scanner = "TRIVY_FS"
	ScannerTRIVYIMAGE               Scanner = "TRIVY_IMAGE"
	ScannerTRUFFLEHOG               Scanner = "TRUFFLEHOG"
	ScannerGITLEAKS                 Scanner = "GITLEAKS"
	ScannerPTAI                     Scanner = "PT_AI"
	ScannerCODESCORINGREPOSITORY    Scanner = "CODE_SCORING_REPOSITORY"
	ScannerCODESCORINGARTIFACT      Scanner = "CODE_SCORING_ARTIFACT"
	ScannerAPPSCREENERSAST          Scanner = "APP_SCREENER_SAST"
	ScannerAPPSCREENERSCAREPOSITORY Scanner = "APP_SCREENER_SCA_REPOSITORY"
	ScannerSASTAV                   Scanner = "SASTAV"
	ScannerZAP                      Scanner = "ZAP"
	ScannerAPPSECTRACKREPOSITORY    Scanner = "APPSEC_TRACK_REPOSITORY"
	ScannerAPPSECTRACKARTIFACT      Scanner = "APPSEC_TRACK_ARTIFACT"
)

type ScannerType string

const (
	ScannerTypeSAST ScannerType = "SAST"
	ScannerTypeDAST ScannerType = "DAST"
	ScannerTypeSCA  ScannerType = "SCA"
)

type ScannerWithConfigFiles string

const (
	ScannerWithConfigFilesZAP ScannerWithConfigFiles = "ZAP"
)

type ScannerWithCustom any

type ScheduledStatus struct {
	StatusFromName           string  `json:"statusFromName,omitempty"`
	StatusToName             string  `json:"statusToName,omitempty"`
	StatusExpiryTime         string  `json:"statusExpiryTime,omitempty"`
	Source                   *string `json:"source,omitempty"`
	InheritedFromScanIssueID *UUID   `json:"inheritedFromScanIssueId,omitempty"`
}

type ServerInfo any

type SettingsConstants struct {
	ScanIssueStatusMaxCustomStatuses int `json:"scanIssueStatusMaxCustomStatuses,omitempty"`
}

type Severity string

const (
	SeverityNONE     Severity = "NONE"
	SeverityLOW      Severity = "LOW"
	SeverityMEDIUM   Severity = "MEDIUM"
	SeverityHIGH     Severity = "HIGH"
	SeverityCRITICAL Severity = "CRITICAL"
	SeverityINFO     Severity = "INFO"
)

type SimpleCardValue any

type SingleScanIssueForExport struct {
	ScanIssueForExport
	ID UUID `json:"id"`
}

type SLAOverdueCardData struct {
	SimpleCardValue
	BySeverity
}

type SMTPConfig any

type SpaceIntegration any

type SpaceIntegrationWithMeta struct {
	SpaceIntegration
	Meta SpaceIntegrationWithMetaMeta `json:"meta,omitempty"`
}

type SpaceIntegrationWithMetaMeta struct {
	AI SpaceIntegrationWithMetaMetaAI `json:"ai,omitempty"`
}

type SpaceIntegrationWithMetaMetaAI struct {
	Vendor                string `json:"vendor,omitempty"`
	IsDefault             bool   `json:"isDefault,omitempty"`
	DefaultModel          string `json:"defaultModel,omitempty"`
	DefaultModelAvailable bool   `json:"defaultModelAvailable,omitempty"`
}

type Space any

type SupportedAction any

type SupportedCondition any

type SupportedParameter any

type SystemWideSettings any

type TaskTrackerAsyncExport struct {
	ExportMode          string                                `json:"exportMode"`
	SummarizedIssue     TaskTrackerAsyncExportSummarizedIssue `json:"summarizedIssue,omitempty"`
	Issues              []TaskTrackerAsyncExportIssues        `json:"issues,omitempty"`
	FieldsDefaultValues []JiraIntegrationFieldsDefaultValues  `json:"fieldsDefaultValues,omitempty"`
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

type TaskTrackerInfoRequest any

type TaskTrackerInfoResponse any

type TaskTrackerIntegration any

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

type TaskTrackerMetaInfoResponse any

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

type TaskTrackerSearchResponse any

type TaskTrackerServiceType string

const (
	TaskTrackerServiceTypeJira              TaskTrackerServiceType = "jira"
	TaskTrackerServiceTypeKtracker          TaskTrackerServiceType = "ktracker"
	TaskTrackerServiceTypeAzureDevopsBoards TaskTrackerServiceType = "azure_devops_boards"
	TaskTrackerServiceTypeEvaTeam           TaskTrackerServiceType = "eva_team"
)

type TemplateSlug string

const (
	TemplateSlugDETAILED TemplateSlug = "DETAILED"
)

type TimeBasedCardValue any

type TopScanIssuesByField struct {
	NameBasedCardValue
}

type TopVCSUsersByIssuesCard struct {
	NameWithEntityBasedCardValue
	Email string `json:"email,omitempty"`
}

type UUID string

type UnitedScanIssueForExport struct {
	ScanIssueForExport
	IDs []UUID `json:"ids"`
}

type UserAdmin any

type UserGroupCounters any

type UserGroupCreateEdit any

type UserGroupMoveUsers any

type UserGroupPermission any

type UserGroupRelation any

type UserGroupRelationType string

const (
	UserGroupRelationTypeREPOSITORY UserGroupRelationType = "REPOSITORY"
	UserGroupRelationTypePRODUCT    UserGroupRelationType = "PRODUCT"
)

type UserGroup any

type UserGroupsSetRole any

type UserMFA any

type UserPermissionValue string

const (
	UserPermissionValueALL                            UserPermissionValue = "ALL"
	UserPermissionValueREPOSITORYREAD                 UserPermissionValue = "REPOSITORY_READ"
	UserPermissionValueREPOSITORYCREATE               UserPermissionValue = "REPOSITORY_CREATE"
	UserPermissionValueREPOSITORYEDIT                 UserPermissionValue = "REPOSITORY_EDIT"
	UserPermissionValueREPOSITORYSETTINGSCHANGE       UserPermissionValue = "REPOSITORY_SETTINGS_CHANGE"
	UserPermissionValuePRODUCTREAD                    UserPermissionValue = "PRODUCT_READ"
	UserPermissionValuePRODUCTCREATE                  UserPermissionValue = "PRODUCT_CREATE"
	UserPermissionValuePRODUCTEDIT                    UserPermissionValue = "PRODUCT_EDIT"
	UserPermissionValuePRODUCTSETTINGSCHANGE          UserPermissionValue = "PRODUCT_SETTINGS_CHANGE"
	UserPermissionValueCOMPONENTREAD                  UserPermissionValue = "COMPONENT_READ"
	UserPermissionValueCOMPONENTEDIT                  UserPermissionValue = "COMPONENT_EDIT"
	UserPermissionValueASSETREAD                      UserPermissionValue = "ASSET_READ"
	UserPermissionValueASSETCREATE                    UserPermissionValue = "ASSET_CREATE"
	UserPermissionValueASSETEDIT                      UserPermissionValue = "ASSET_EDIT"
	UserPermissionValueSCANISSUEREAD                  UserPermissionValue = "SCAN_ISSUE_READ"
	UserPermissionValueSCANISSUECREATE                UserPermissionValue = "SCAN_ISSUE_CREATE"
	UserPermissionValueSCANISSUEDELETE                UserPermissionValue = "SCAN_ISSUE_DELETE"
	UserPermissionValueSCANISSUEEDIT                  UserPermissionValue = "SCAN_ISSUE_EDIT"
	UserPermissionValueSCANISSUESENDTOINTEGRATIONS    UserPermissionValue = "SCAN_ISSUE_SEND_TO_INTEGRATIONS"
	UserPermissionValueSCANISSUESTATUSCHANGE          UserPermissionValue = "SCAN_ISSUE_STATUS_CHANGE"
	UserPermissionValueSCANISSUESEVERITYCHANGE        UserPermissionValue = "SCAN_ISSUE_SEVERITY_CHANGE"
	UserPermissionValueBOTUSERREAD                    UserPermissionValue = "BOT_USER_READ"
	UserPermissionValueBOTUSERCREATE                  UserPermissionValue = "BOT_USER_CREATE"
	UserPermissionValueBOTUSEREDIT                    UserPermissionValue = "BOT_USER_EDIT"
	UserPermissionValueSCANPIPELINEREAD               UserPermissionValue = "SCAN_PIPELINE_READ"
	UserPermissionValueSCANPIPELINEWRITE              UserPermissionValue = "SCAN_PIPELINE_WRITE"
	UserPermissionValueQUALITYGATEREAD                UserPermissionValue = "QUALITY_GATE_READ"
	UserPermissionValueQUALITYGATECREATE              UserPermissionValue = "QUALITY_GATE_CREATE"
	UserPermissionValueQUALITYGATEEDIT                UserPermissionValue = "QUALITY_GATE_EDIT"
	UserPermissionValueRULESREAD                      UserPermissionValue = "RULES_READ"
	UserPermissionValueRULESWRITE                     UserPermissionValue = "RULES_WRITE"
	UserPermissionValueADMINSPACEWRITE                UserPermissionValue = "ADMIN_SPACE_WRITE"
	UserPermissionValueADMININTEGRATIONREAD           UserPermissionValue = "ADMIN_INTEGRATION_READ"
	UserPermissionValueADMININTEGRATIONWRITE          UserPermissionValue = "ADMIN_INTEGRATION_WRITE"
	UserPermissionValueADMININTEGRATIONIMPORTFINDINGS UserPermissionValue = "ADMIN_INTEGRATION_IMPORT_FINDINGS"
	UserPermissionValueADMINBACKGROUNDTASKSREAD       UserPermissionValue = "ADMIN_BACKGROUND_TASKS_READ"
	UserPermissionValueADMINLDAPREAD                  UserPermissionValue = "ADMIN_LDAP_READ"
	UserPermissionValueADMINLDAPWRITE                 UserPermissionValue = "ADMIN_LDAP_WRITE"
	UserPermissionValueADMINOIDCREAD                  UserPermissionValue = "ADMIN_OIDC_READ"
	UserPermissionValueADMINOIDCWRITE                 UserPermissionValue = "ADMIN_OIDC_WRITE"
	UserPermissionValueADMINSMTPREAD                  UserPermissionValue = "ADMIN_SMTP_READ"
	UserPermissionValueADMINSMTPWRITE                 UserPermissionValue = "ADMIN_SMTP_WRITE"
	UserPermissionValueADMINISSUESTATUSWRITE          UserPermissionValue = "ADMIN_ISSUE_STATUS_WRITE"
	UserPermissionValueADMINSYSTEMSETTINGSWRITE       UserPermissionValue = "ADMIN_SYSTEM_SETTINGS_WRITE"
	UserPermissionValueADMINSLAREAD                   UserPermissionValue = "ADMIN_SLA_READ"
	UserPermissionValueADMINSLAWRITE                  UserPermissionValue = "ADMIN_SLA_WRITE"
	UserPermissionValueADMINLICENSEREAD               UserPermissionValue = "ADMIN_LICENSE_READ"
	UserPermissionValueADMINLICENSEWRITE              UserPermissionValue = "ADMIN_LICENSE_WRITE"
	UserPermissionValueADMINUSERREAD                  UserPermissionValue = "ADMIN_USER_READ"
	UserPermissionValueADMINUSERWRITE                 UserPermissionValue = "ADMIN_USER_WRITE"
	UserPermissionValueADMINUSERROLEREAD              UserPermissionValue = "ADMIN_USER_ROLE_READ"
	UserPermissionValueADMINUSERROLEWRITE             UserPermissionValue = "ADMIN_USER_ROLE_WRITE"
	UserPermissionValueADMINHEALTHCHECK               UserPermissionValue = "ADMIN_HEALTH_CHECK"
	UserPermissionValueADMINFPCLASSIFIERREAD          UserPermissionValue = "ADMIN_FP_CLASSIFIER_READ"
	UserPermissionValueADMINFPCLASSIFIERWRITE         UserPermissionValue = "ADMIN_FP_CLASSIFIER_WRITE"
	UserPermissionValueORCHESTRATIONRUN               UserPermissionValue = "ORCHESTRATION_RUN"
	UserPermissionValueORCHESTRATIONPROFILEREAD       UserPermissionValue = "ORCHESTRATION_PROFILE_READ"
	UserPermissionValueORCHESTRATIONPROFILEEDIT       UserPermissionValue = "ORCHESTRATION_PROFILE_EDIT"
	UserPermissionValueORCHESTRATIONPROFILECREATE     UserPermissionValue = "ORCHESTRATION_PROFILE_CREATE"
	UserPermissionValueCOMMENTCREATE                  UserPermissionValue = "COMMENT_CREATE"
	UserPermissionValueCOMMENTREAD                    UserPermissionValue = "COMMENT_READ"
	UserPermissionValueCOMMENTEDIT                    UserPermissionValue = "COMMENT_EDIT"
	UserPermissionValueCOMMENTEDITOWN                 UserPermissionValue = "COMMENT_EDIT_OWN"
	UserPermissionValueARTIFACTREAD                   UserPermissionValue = "ARTIFACT_READ"
	UserPermissionValueARTIFACTEDIT                   UserPermissionValue = "ARTIFACT_EDIT"
	UserPermissionValueAIMODELSACCESS                 UserPermissionValue = "AI_MODELS_ACCESS"
	UserPermissionValueAISCANISSUEAPPLY               UserPermissionValue = "AI_SCAN_ISSUE_APPLY"
	UserPermissionValueAISCANISSUECHANGEANSWER        UserPermissionValue = "AI_SCAN_ISSUE_CHANGE_ANSWER"
	UserPermissionValueADMINUSERGROUPREAD             UserPermissionValue = "ADMIN_USER_GROUP_READ"
	UserPermissionValueADMINUSERGROUPWRITE            UserPermissionValue = "ADMIN_USER_GROUP_WRITE"
)

type UserRegistration any

type UserRoleCounters any

type UserRoleEdit any

type UserRolePermissionInfo any

type UserRolePermissionsCategory any

type UserRolePermissionsMeta any

type UserRole struct {
	UserRoleShort
	Permissions []UserPermissionValue `json:"permissions,omitempty"`
	Counters    UserRoleCounters      `json:"counters,omitempty"`
}

type UserRoleShort any

type UserRoleSlug string

type UserRoleValidationError any

type UserRoleValidationErrorProblems struct {
	MissingPermissions      []UserPermissionValue `json:"missingPermissions,omitempty"`
	UnexpectedPermissions   []UserPermissionValue `json:"unexpectedPermissions,omitempty"`
	UnexpectedIssueStatuses []string              `json:"unexpectedIssueStatuses,omitempty"`
}

type UserShort any

type UserSystemRoleValue string

const (
	UserSystemRoleValueADMIN       UserSystemRoleValue = "ADMIN"
	UserSystemRoleValueSYSENGINEER UserSystemRoleValue = "SYS_ENGINEER"
	UserSystemRoleValueSECENGINEER UserSystemRoleValue = "SEC_ENGINEER"
	UserSystemRoleValueOWNER       UserSystemRoleValue = "OWNER"
	UserSystemRoleValueDEVELOPER   UserSystemRoleValue = "DEVELOPER"
)

type UserTiny any

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

type DatasetUploadBody struct {
	File []byte `json:"file"`
}

type ImportDefectDojoBody struct {
	SpaceID         *UUID   `json:"spaceID,omitempty"`
	Hostname        string  `json:"hostname"`
	Token           string  `json:"token"`
	IgnoreSSLCheck  bool    `json:"ignoreSSLCheck,omitempty"`
	Timeout         *int    `json:"timeout,omitempty"`
	BatchSize       *int    `json:"batchSize,omitempty"`
	ProductIDs      *string `json:"productIDs,omitempty"`
	FindingStatuses *string `json:"findingStatuses,omitempty"`
	CreateProducts  bool    `json:"createProducts,omitempty"`
	CreateComments  bool    `json:"createComments,omitempty"`
}

type InlineResponse200 any

type InlineResponse2001 any

type InlineResponse20010 any

type InlineResponse20011 any

type InlineResponse20012 struct {
	Items      []Repository `json:"items,omitempty"`
	HasNext    bool         `json:"hasNext,omitempty"`
	Offset     int          `json:"offset,omitempty"`
	TotalCount int          `json:"totalCount,omitempty"`
}

type InlineResponse20013 struct {
	Items      []RepositoryBranchList `json:"items,omitempty"`
	HasNext    bool                   `json:"hasNext,omitempty"`
	Offset     int                    `json:"offset,omitempty"`
	TotalCount int                    `json:"totalCount,omitempty"`
}

type InlineResponse20014 struct {
	Items      []ComponentShort `json:"items,omitempty"`
	HasNext    bool             `json:"hasNext,omitempty"`
	Offset     int              `json:"offset,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
}

type InlineResponse20015 struct {
	Items      []ComponentVersion `json:"items,omitempty"`
	HasNext    bool               `json:"hasNext,omitempty"`
	Offset     int                `json:"offset,omitempty"`
	TotalCount int                `json:"totalCount,omitempty"`
}

type InlineResponse20016 any

type InlineResponse20017 any

type InlineResponse20018 struct {
	Items      []OrchestrationScanDetail `json:"items,omitempty"`
	HasNext    bool                      `json:"hasNext,omitempty"`
	Offset     int                       `json:"offset,omitempty"`
	TotalCount int                       `json:"totalCount,omitempty"`
}

type InlineResponse20019 struct {
	Items      []OrchestrationScanProfile `json:"items,omitempty"`
	HasNext    bool                       `json:"hasNext,omitempty"`
	Offset     int                        `json:"offset,omitempty"`
	TotalCount int                        `json:"totalCount,omitempty"`
}

type InlineResponse2002 any

type InlineResponse20020 struct {
	Items      []OrchestrationScanProfileRelation `json:"items,omitempty"`
	HasNext    bool                               `json:"hasNext,omitempty"`
	Offset     int                                `json:"offset,omitempty"`
	TotalCount int                                `json:"totalCount,omitempty"`
}

type InlineResponse20021 struct {
	Items      []CommentDetails `json:"items,omitempty"`
	HasNext    bool             `json:"hasNext,omitempty"`
	Offset     int              `json:"offset,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
}

type InlineResponse20022 struct {
	Items      []Artifact `json:"items,omitempty"`
	HasNext    bool       `json:"hasNext,omitempty"`
	Offset     int        `json:"offset,omitempty"`
	TotalCount int        `json:"totalCount,omitempty"`
}

type InlineResponse20023 struct {
	Items      []ArtifactVersion `json:"items,omitempty"`
	HasNext    bool              `json:"hasNext,omitempty"`
	Offset     int               `json:"offset,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
}

type InlineResponse20024 any

type InlineResponse20025 any

type InlineResponse20026 struct {
	Items      []Report `json:"items,omitempty"`
	HasNext    bool     `json:"hasNext,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	TotalCount int      `json:"totalCount,omitempty"`
}

type InlineResponse20027 struct {
	Items      []ScanPipeline `json:"items,omitempty"`
	HasNext    bool           `json:"hasNext,omitempty"`
	Offset     int            `json:"offset,omitempty"`
	TotalCount int            `json:"totalCount,omitempty"`
}

type InlineResponse20028 struct {
	Items      []ScanPipelineJob `json:"items,omitempty"`
	HasNext    bool              `json:"hasNext,omitempty"`
	Offset     int               `json:"offset,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
}

type InlineResponse20029 any

type InlineResponse2003 struct {
	Items      []DefectDojoProduct `json:"items,omitempty"`
	HasNext    bool                `json:"hasNext,omitempty"`
	Offset     int                 `json:"offset,omitempty"`
	TotalCount int                 `json:"totalCount,omitempty"`
}

type InlineResponse20030 any

type InlineResponse20031 struct {
	Items      []QualityGate `json:"items,omitempty"`
	HasNext    bool          `json:"hasNext,omitempty"`
	Offset     int           `json:"offset,omitempty"`
	TotalCount int           `json:"totalCount,omitempty"`
}

type InlineResponse2004 any

type InlineResponse2005 any

type InlineResponse2006 any

type InlineResponse2007 struct {
	Items      []Product `json:"items,omitempty"`
	HasNext    bool      `json:"hasNext,omitempty"`
	Offset     int       `json:"offset,omitempty"`
	TotalCount int       `json:"totalCount,omitempty"`
}

type InlineResponse2008 struct {
	Items      []Asset `json:"items,omitempty"`
	HasNext    bool    `json:"hasNext,omitempty"`
	Offset     int     `json:"offset,omitempty"`
	TotalCount int     `json:"totalCount,omitempty"`
}

type InlineResponse2009 struct {
	Items      []QualityGateResult `json:"items,omitempty"`
	HasNext    bool                `json:"hasNext,omitempty"`
	Offset     int                 `json:"offset,omitempty"`
	TotalCount int                 `json:"totalCount,omitempty"`
}

type IntegrationIDUploadCLIBody struct {
	File []byte `json:"file"`
}

type IntegrationIDUploadCLIBody1 struct {
	File []byte `json:"file"`
}

type IntegrationIDUploadCLIBody2 struct {
	File []byte `json:"file"`
}

type LicenseBody struct {
	LicenseFile []byte `json:"license_file"`
}

type OrchestrationRunBatchBody struct {
	SpaceID            *UUID   `json:"spaceID,omitempty"`
	ArtifactIDs        *string `json:"artifactIDs,omitempty"`
	ArtifactVersionIDs *string `json:"artifactVersionIDs,omitempty"`
	AssetIDs           *string `json:"assetIDs,omitempty"`
	Scanner            Scanner `json:"scanner,omitempty"`
	ScannerOptions     string  `json:"scannerOptions,omitempty"`
	Profile            *string `json:"profile,omitempty"`
}

type OrchestrationRunBody struct {
	SpaceID                    *UUID   `json:"spaceID,omitempty"`
	Repository                 *string `json:"repository,omitempty"`
	Product                    *string `json:"product,omitempty"`
	RepositoryBranch           *string `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool   `json:"repositoryBranchSetDefault,omitempty"`
	RepositoryArchive          []byte  `json:"repositoryArchive,omitempty"`
	ArtifactArchive            []byte  `json:"artifactArchive,omitempty"`
	CreateMissing              *bool   `json:"createMissing,omitempty"`
	Artifact                   *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool   `json:"artifactVersionSetDefault,omitempty"`
	Asset                      *string `json:"asset,omitempty"`
	Endpoint                   *string `json:"endpoint,omitempty"`
	Scanner                    Scanner `json:"scanner,omitempty"`
	ScannerOptions             string  `json:"scannerOptions,omitempty"`
	Profile                    *string `json:"profile,omitempty"`
}

type OrchestrationRunBody1 struct {
	SpaceID                    *UUID   `json:"spaceID,omitempty"`
	Repository                 *string `json:"repository,omitempty"`
	Product                    *string `json:"product,omitempty"`
	RepositoryBranch           *string `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool   `json:"repositoryBranchSetDefault,omitempty"`
	RepositoryArchive          []byte  `json:"repositoryArchive,omitempty"`
	ArtifactArchive            []byte  `json:"artifactArchive,omitempty"`
	CreateMissing              *bool   `json:"createMissing,omitempty"`
	Artifact                   *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool   `json:"artifactVersionSetDefault,omitempty"`
	Asset                      *string `json:"asset,omitempty"`
	Scanner                    Scanner `json:"scanner,omitempty"`
	ScannerOptions             string  `json:"scannerOptions,omitempty"`
	DryRun                     bool    `json:"dryRun,omitempty"`
	Profile                    *string `json:"profile,omitempty"`
}

type OrchestrationUploadScannerConfigBody struct {
	Scanner ScannerWithConfigFiles `json:"scanner"`
	File    []byte                 `json:"file"`
}

type SBOMImportBody struct {
	SBOMResultFile []byte `json:"sbomResultFile"`
	RepositoryID   *UUID  `json:"repositoryID,omitempty"`
}

type ScanUploadsBody struct {
	ScannerResultFile          []byte  `json:"scannerResultFile"`
	ScannerType                Parser  `json:"scannerType"`
	SpaceID                    *UUID   `json:"spaceID,omitempty"`
	Repository                 *string `json:"repository,omitempty"`
	RepositoryBranch           *string `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool   `json:"repositoryBranchSetDefault,omitempty"`
	Product                    *string `json:"product,omitempty"`
	Artifact                   *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool   `json:"artifactVersionSetDefault,omitempty"`
	VCSURL                     *string `json:"vcsURL,omitempty"`
	ScanPipelineJobID          UUID    `json:"scanPipelineJobID,omitempty"`
	ExternalScanPipeline       string  `json:"externalScanPipeline,omitempty"`
	ScanPipelineProduct        *string `json:"scanPipelineProduct,omitempty"`
	FinishScanPipeline         *bool   `json:"finishScanPipeline,omitempty"`
	ScanPipelineName           *string `json:"scanPipelineName,omitempty"`
	SkipAutoResolve            *bool   `json:"skipAutoResolve,omitempty"`
	CommitSHA                  *string `json:"commitSha,omitempty"`
}

type ScanUploadsBulkDownloadBody struct {
	ScanUploads []UUID `json:"scanUploads,omitempty"`
}

type SnapshotsSnapshotIDBody struct {
	IsCurrent bool   `json:"isCurrent,omitempty"`
	Name      string `json:"name,omitempty"`
}

type SnapshotsUploadBody struct {
	File []byte `json:"file"`
}

type TokenWebhookBody any

type V1ScanBody struct {
	SpaceID                   *UUID   `json:"spaceID,omitempty"`
	Repository                *string `json:"repository,omitempty"`
	Product                   *string `json:"product,omitempty"`
	RepositoryBranch          *string `json:"repositoryBranch,omitempty"`
	RepositoryArchive         []byte  `json:"repositoryArchive,omitempty"`
	ArtifactArchive           []byte  `json:"artifactArchive,omitempty"`
	CreateMissing             *bool   `json:"createMissing,omitempty"`
	Scanner                   Scanner `json:"scanner,omitempty"`
	ScannerOptions            string  `json:"scannerOptions,omitempty"`
	DryRun                    bool    `json:"dryRun,omitempty"`
	Profile                   *string `json:"profile,omitempty"`
	Artifact                  *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault *bool   `json:"artifactVersionSetDefault,omitempty"`
	Asset                     *string `json:"asset,omitempty"`
	Endpoint                  *string `json:"endpoint,omitempty"`
}

type V1ScanUploadsBody struct {
	ScannerResultFile          []byte  `json:"scannerResultFile"`
	ScannerType                Parser  `json:"scannerType"`
	SpaceID                    *UUID   `json:"spaceID,omitempty"`
	Artifact                   *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool   `json:"artifactVersionSetDefault,omitempty"`
	Repository                 *string `json:"repository,omitempty"`
	RepositoryBranch           *string `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool   `json:"repositoryBranchSetDefault,omitempty"`
	Product                    *string `json:"product,omitempty"`
	VCSURL                     *string `json:"vcsURL,omitempty"`
	ExternalScanPipeline       string  `json:"externalScanPipeline,omitempty"`
	ScanPipelineProduct        *string `json:"scanPipelineProduct,omitempty"`
	FinishScanPipeline         *bool   `json:"finishScanPipeline,omitempty"`
	ScanPipelineName           *string `json:"scanPipelineName,omitempty"`
	SkipAutoResolve            *bool   `json:"skipAutoResolve,omitempty"`
	CommitSHA                  *string `json:"commitSha,omitempty"`
}
