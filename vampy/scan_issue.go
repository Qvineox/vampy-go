// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type BaseIssueStatus struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	System      bool   `json:"system,omitempty"`
	Author      string `json:"author,omitempty"`
	Color       string `json:"color,omitempty"`
	ListOrder   int    `json:"listOrder,omitempty"`
	NeedComment bool   `json:"needComment,omitempty"`
	Deleted     bool   `json:"deleted,omitempty"`
}

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

type IssueExport struct {
	ID                 uuid.UUID               `json:"id,omitempty"`
	ExternalID         string                  `json:"externalID,omitempty"`
	IntegrationService string                  `json:"integrationService,omitempty"`
	IntegrationID      uuid.UUID               `json:"integrationID,omitempty"`
	IntegrationExists  bool                    `json:"integrationExists,omitempty"`
	MetaData           *TaskTrackerIssueExport `json:"metaData,omitempty"`
	Created            string                  `json:"created,omitempty"`
}

type IssueStatusChangeIntegration struct {
	ID       uuid.UUID                         `json:"id,omitempty"`
	Name     string                            `json:"name,omitempty"`
	Hostname string                            `json:"hostname,omitempty"`
	Meta     *IssueStatusChangeIntegrationMeta `json:"meta,omitempty"`
}

type IssueStatusChangeIntegrationMeta struct {
	AI *IssueStatusChangeIntegrationMetaAI `json:"ai,omitempty"`
}

type IssueStatusChangeIntegrationMetaAI struct {
	Vendor                string  `json:"vendor,omitempty"`
	IsDefault             bool    `json:"isDefault,omitempty"`
	DefaultModel          *string `json:"defaultModel,omitempty"`
	DefaultModelAvailable bool    `json:"defaultModelAvailable,omitempty"`
}

type IssueStatusChangeRule struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

type IssueStatusChangeScanUpload struct {
	ID         uuid.UUID              `json:"id,omitempty"`
	Created    string                 `json:"created,omitempty"`
	Parser     Parser                 `json:"parser,omitempty"`
	FileExists bool                   `json:"fileExists,omitempty"`
	Uploader   *IssueStatusChangeUser `json:"uploader,omitempty"`
}

type IssueStatusChangeSource string

const (
	IssueStatusChangeSourceUSER            IssueStatusChangeSource = "USER"
	IssueStatusChangeSourceINTEGRATION     IssueStatusChangeSource = "INTEGRATION"
	IssueStatusChangeSourceSCANUPLOAD      IssueStatusChangeSource = "SCAN_UPLOAD"
	IssueStatusChangeSourceRULESENGINE     IssueStatusChangeSource = "RULES_ENGINE"
	IssueStatusChangeSourceSCHEDULER       IssueStatusChangeSource = "SCHEDULER"
	IssueStatusChangeSourceASSOCIATEDISSUE IssueStatusChangeSource = "ASSOCIATED_ISSUE"
)

type IssueStatusChangeUser struct {
	ID           uuid.UUID      `json:"id,omitempty"`
	UserLogin    string         `json:"userLogin,omitempty"`
	UserEmail    string         `json:"userEmail,omitempty"`
	FirstName    string         `json:"firstName,omitempty"`
	LastName     string         `json:"lastName,omitempty"`
	Role         UserRoleSlug   `json:"role,omitempty"`
	RoleInstance *UserRoleShort `json:"roleInstance,omitempty"`
	IsBot        bool           `json:"isBot,omitempty"`
}

type IssueStatusCreate struct {
	DisplayName string `json:"displayName,omitempty"`
	System      bool   `json:"system,omitempty"`
	Author      string `json:"author,omitempty"`
	Color       string `json:"color,omitempty"`
	ListOrder   *int   `json:"listOrder,omitempty"`
	Completed   bool   `json:"completed,omitempty"`
	NeedComment bool   `json:"needComment,omitempty"`
}

type IssueStatus struct {
	BaseIssueStatus
	Name      string `json:"name,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}

type IssueStatusUpdate struct {
	BaseIssueStatus
}

type ScanIssueAIRequestApply struct {
	ValueNumber *float64 `json:"valueNumber,omitempty"`
	ValueText   *string  `json:"valueText,omitempty"`
}

type ScanIssueAIRequestCancel struct {
	Request ScanIssueAIRequestType `json:"request"`
}

type ScanIssueAIRequestResponse struct {
	ID          uuid.UUID                 `json:"id,omitempty"`
	Status      ScanIssueAIRequestStatus  `json:"status,omitempty"`
	Request     ScanIssueAIRequestType    `json:"request,omitempty"`
	Error       *string                   `json:"error,omitempty"`
	ScanIssueID uuid.UUID                 `json:"scanIssueID,omitempty"`
	Integration *SpaceIntegrationWithMeta `json:"integration,omitempty"`
	AIModel     string                    `json:"aiModel,omitempty"`
	ValueNumber *float64                  `json:"valueNumber,omitempty"`
	ValueText   *string                   `json:"valueText,omitempty"`
	Explanation *string                   `json:"explanation,omitempty"`
	Created     string                    `json:"created,omitempty"`
	Updated     *string                   `json:"updated,omitempty"`
	CreatedBy   *UserTiny                 `json:"createdBy,omitempty"`
	UpdatedBy   *UserTiny                 `json:"updatedBy,omitempty"`
}

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

type ScanIssueBySeverityStat struct {
	Severity Severity `json:"severity,omitempty"`
	Count    int      `json:"count,omitempty"`
}

type ScanIssueByStatusStat struct {
	Status string `json:"status,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type ScanIssueChangeEntry struct {
	ID         uuid.UUID                   `json:"id,omitempty"`
	Created    string                      `json:"created,omitempty"`
	ChangeType IssueChangeType             `json:"changeType,omitempty"`
	OldValue   *string                     `json:"oldValue,omitempty"`
	NewValue   string                      `json:"newValue,omitempty"`
	Automatic  bool                        `json:"automatic,omitempty"`
	Source     *ScanIssueChangeEntrySource `json:"source,omitempty"`
}

type ScanIssueChangeEntrySource struct {
	Type                      IssueChangeSource             `json:"type,omitempty"`
	ScanUpload                *IssueStatusChangeScanUpload  `json:"scanUpload,omitempty"`
	User                      *IssueStatusChangeUser        `json:"user,omitempty"`
	Integration               *IssueStatusChangeIntegration `json:"integration,omitempty"`
	Rule                      *IssueStatusChangeRule        `json:"rule,omitempty"`
	ScheduledStatusChangeInfo *ScheduledStatus              `json:"scheduledStatusChangeInfo,omitempty"`
	SourceScanIssue           *ScanIssueShortWithBranch     `json:"sourceScanIssue,omitempty"`
	AIModelName               *string                       `json:"aiModelName,omitempty"`
	FPClassifierSnapshot      *FPClassifierSnapshotShort    `json:"fpClassifierSnapshot,omitempty"`
}

type ScanIssueCountsStat struct {
	Total      int `json:"total,omitempty"`
	Incomplete int `json:"incomplete,omitempty"`
}

type ScanIssueCreateResponse struct {
	Status    string           `json:"status,omitempty"`
	ScanIssue *ScanIssueDetail `json:"scanIssue,omitempty"`
}

type ScanIssueCreate struct {
	Title            string     `json:"title"`
	Status           string     `json:"status,omitempty"`
	Severity         Severity   `json:"severity,omitempty"`
	Description      string     `json:"description,omitempty"`
	CVEID            string     `json:"cveID,omitempty"`
	CWEID            int        `json:"cweID,omitempty"`
	CodeSnippet      string     `json:"codeSnippet,omitempty"`
	LibraryName      string     `json:"libraryName,omitempty"`
	LibraryVersion   string     `json:"libraryVersion,omitempty"`
	FilePath         string     `json:"filePath,omitempty"`
	FileLineno       *int       `json:"fileLineno,omitempty"`
	Solution         string     `json:"solution,omitempty"`
	RepositoryID     *uuid.UUID `json:"repositoryID,omitempty"`
	RepositoryBranch string     `json:"repositoryBranch,omitempty"`
	References       []string   `json:"references,omitempty"`
}

type ScanIssueDelete struct {
	Issues []uuid.UUID `json:"issues,omitempty"`
}

type ScanIssueDetail struct {
	ScanIssue
	Data                 map[string]any                       `json:"data,omitempty"`
	Assets               []Asset                              `json:"assets,omitempty"`
	Description          string                               `json:"description,omitempty"`
	CodeSnippet          string                               `json:"codeSnippet,omitempty"`
	Solution             string                               `json:"solution,omitempty"`
	References           []string                             `json:"references,omitempty"`
	ArtifactVersion      *ArtifactVersionShort                `json:"artifactVersion,omitempty"`
	Product              *ProductShort                        `json:"product,omitempty"`
	ComponentVersion     *ComponentVersionShort               `json:"componentVersion,omitempty"`
	FilePathVCSLink      *string                              `json:"filePathVCSLink,omitempty"`
	CommitSHA            *ScanIssueDetailCommitSHA            `json:"commitSha,omitempty"`
	GitBlame             *ScanIssueGitBlame                   `json:"gitBlame,omitempty"`
	Duplicates           *ScanIssueDetailDuplicates           `json:"duplicates,omitempty"`
	Comments             *ScanIssueDetailDuplicates           `json:"comments,omitempty"`
	AssociatedScanIssues *ScanIssueDetailAssociatedScanIssues `json:"associatedScanIssues,omitempty"`
	AIRequests           []ScanIssueAIRequestResponse         `json:"aiRequests,omitempty"`
	AIApplied            *ScanIssueDetailAIApplied            `json:"aiApplied,omitempty"`
	FPScore              *FPScoreDetail                       `json:"fpScore,omitempty"`
}

type ScanIssueDetailAIApplied struct {
	FalsePositiveRate *ScanIssueAIRequestResponse `json:"falsePositiveRate,omitempty"`
	Description       *ScanIssueAIRequestResponse `json:"description,omitempty"`
	Solution          *ScanIssueAIRequestResponse `json:"solution,omitempty"`
	CodeExplain       *ScanIssueAIRequestResponse `json:"codeExplain,omitempty"`
	CodeVuln          *ScanIssueAIRequestResponse `json:"codeVuln,omitempty"`
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

type ScanIssueDuplicatesSearchResult struct {
	ScanIssues []ScanIssueDuplicatesSearchResultScanIssues `json:"scanIssues,omitempty"`
}

type ScanIssueDuplicatesSearchResultScanIssues struct {
	ID     uuid.UUID `json:"id,omitempty"`
	Title  string    `json:"title,omitempty"`
	Status string    `json:"status,omitempty"`
}

type ScanIssueExport struct {
	ID                 uuid.UUID          `json:"id,omitempty"`
	Status             string             `json:"status,omitempty"`
	Title              string             `json:"title,omitempty"`
	Created            string             `json:"created,omitempty"`
	Completed          string             `json:"completed,omitempty"`
	FilePath           string             `json:"filePath,omitempty"`
	FileLineno         *int               `json:"fileLineno,omitempty"`
	LibraryName        string             `json:"libraryName,omitempty"`
	LibraryVersion     string             `json:"libraryVersion,omitempty"`
	CVEID              string             `json:"cveID,omitempty"`
	CWEID              *int               `json:"cweID,omitempty"`
	DependencyRelation DependencyRelation `json:"dependencyRelation,omitempty"`
	Severity           Severity           `json:"severity,omitempty"`
	Parser             Parser             `json:"parser,omitempty"`
	FoundInParsers     []Parser           `json:"foundInParsers,omitempty"`
	SLA                *string            `json:"sla,omitempty"`
	Product            *string            `json:"product,omitempty"`
	Repository         *string            `json:"repository,omitempty"`
	RepositoryBranch   *string            `json:"repositoryBranch,omitempty"`
	IssueLink          *string            `json:"issueLink,omitempty"`
	Solution           *string            `json:"solution,omitempty"`
	CommitSHALastSeen  *string            `json:"commitShaLastSeen,omitempty"`
}

type ScanIssueFPScoreAIModelFilter struct {
	IntegrationID uuid.UUID `json:"integrationID"`
	AIModel       string    `json:"aiModel"`
}

type ScanIssueFPScoreFilter struct {
	ValueFrom *int  `json:"valueFrom,omitempty"`
	ValueTo   *int  `json:"valueTo,omitempty"`
	NoValue   *bool `json:"noValue,omitempty"`
}

type ScanIssueFilter struct {
	SpaceID               uuid.UUID                       `json:"spaceID,omitempty"`
	Search                string                          `json:"search,omitempty"`
	Statuses              []string                        `json:"statuses,omitempty"`
	Severities            []Severity                      `json:"severities,omitempty"`
	Repositories          []uuid.UUID                     `json:"repositories,omitempty"`
	RepositoryBranches    []uuid.UUID                     `json:"repositoryBranches,omitempty"`
	Products              []uuid.UUID                     `json:"products,omitempty"`
	Engagements           []uuid.UUID                     `json:"engagements,omitempty"`
	Assets                []uuid.UUID                     `json:"assets,omitempty"`
	Components            []uuid.UUID                     `json:"components,omitempty"`
	ComponentVersions     []uuid.UUID                     `json:"componentVersions,omitempty"`
	Parsers               []string                        `json:"parsers,omitempty"`
	HasSolution           bool                            `json:"hasSolution,omitempty"`
	ScannerTypes          []ScannerType                   `json:"scannerTypes,omitempty"`
	DependencyRelations   []DependencyRelationFilter      `json:"dependencyRelations,omitempty"`
	Overdue               []string                        `json:"overdue,omitempty"`
	VCSUsers              []uuid.UUID                     `json:"vcsUsers,omitempty"`
	DefaultRelations      bool                            `json:"defaultRelations,omitempty"`
	CWEIDs                []int                           `json:"cweIDs,omitempty"`
	CVEIDs                []string                        `json:"cveIDs,omitempty"`
	FilePaths             []string                        `json:"filePaths,omitempty"`
	Titles                []string                        `json:"titles,omitempty"`
	ArtifactVersionIDs    []string                        `json:"artifactVersionIDs,omitempty"`
	ArtifactIDs           []string                        `json:"artifactIDs,omitempty"`
	OldIssues             *bool                           `json:"oldIssues,omitempty"`
	AssociatedWithIssueID uuid.UUID                       `json:"associatedWithIssueID,omitempty"`
	ScanPipelineID        uuid.UUID                       `json:"scanPipelineID,omitempty"`
	ScanPipelineJobID     uuid.UUID                       `json:"scanPipelineJobID,omitempty"`
	FindingState          FindingState                    `json:"findingState,omitempty"`
	FindingSeverity       Severity                        `json:"findingSeverity,omitempty"`
	FPScore               *ScanIssueFPScoreFilter         `json:"fpScore,omitempty"`
	FPScoreAIModels       []ScanIssueFPScoreAIModelFilter `json:"fpScoreAIModels,omitempty"`
	FPScoreSourceType     []FPScoreSourceType             `json:"fpScoreSourceType,omitempty"`
}

type ScanIssueForExportResponse struct {
	SummarizedIssue *UnitedScanIssueForExport  `json:"summarizedIssue,omitempty"`
	ExportingIssues []SingleScanIssueForExport `json:"exportingIssues"`
}

type ScanIssueForExport struct {
	Name    string  `json:"name"`
	Summary string  `json:"summary"`
	SLA     *string `json:"sla,omitempty"`
}

type ScanIssueGitBlame struct {
	ID            uuid.UUID                      `json:"id,omitempty"`
	CommitTime    string                         `json:"commitTime,omitempty"`
	CommitID      string                         `json:"commitID,omitempty"`
	CommitURL     string                         `json:"commitURL,omitempty"`
	CommitMessage string                         `json:"commitMessage,omitempty"`
	CommitAuthor  *ScanIssueGitBlameCommitAuthor `json:"commitAuthor,omitempty"`
}

type ScanIssueGitBlameCommitAuthor struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Email string    `json:"email,omitempty"`
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

type ScanIssue struct {
	ID                 uuid.UUID                  `json:"id,omitempty"`
	Title              string                     `json:"title,omitempty"`
	Created            string                     `json:"created,omitempty"`
	FilePath           string                     `json:"filePath,omitempty"`
	FileLineno         *int                       `json:"fileLineno,omitempty"`
	LibraryName        string                     `json:"libraryName,omitempty"`
	LibraryVersion     string                     `json:"libraryVersion,omitempty"`
	CVEID              string                     `json:"cveID,omitempty"`
	CWEID              *int                       `json:"cweID,omitempty"`
	DependencyRelation DependencyRelation         `json:"dependencyRelation,omitempty"`
	Severity           Severity                   `json:"severity,omitempty"`
	SLA                *string                    `json:"sla,omitempty"`
	Status             string                     `json:"status,omitempty"`
	Parser             Parser                     `json:"parser,omitempty"`
	FoundInParsers     []Parser                   `json:"foundInParsers,omitempty"`
	Repository         *RepositoryShort           `json:"repository,omitempty"`
	RepositoryBranches []RepositoryBranch         `json:"repositoryBranches,omitempty"`
	RepositoryBranch   *RepositoryBranch          `json:"repositoryBranch,omitempty"`
	ExportRecords      []IssueExport              `json:"exportRecords,omitempty"`
	Permissions        []ScanIssuePermissionValue `json:"permissions,omitempty"`
	ScheduledStatus    *ScheduledStatus           `json:"scheduledStatus,omitempty"`
}

type ScanIssues struct {
	ScanIssueType ScanIssueType     `json:"scanIssueType,omitempty"`
	Schema        *ScanIssuesSchema `json:"schema,omitempty"`
}

type ScanIssuesSchema struct {
	Name          string                          `json:"name,omitempty"`
	DisplayName   string                          `json:"displayName,omitempty"`
	Type          string                          `json:"type,omitempty"`
	Hidden        bool                            `json:"hidden,omitempty"`
	IsList        bool                            `json:"isList,omitempty"`
	Metadata      *ScanIssuesSchemaMetadata       `json:"metadata,omitempty"`
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

type ScanIssueShort struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Created  string    `json:"created,omitempty"`
	Status   string    `json:"status,omitempty"`
	Severity Severity  `json:"severity,omitempty"`
}

type ScanIssueShortWithBranch struct {
	ScanIssueShort
	RepositoryBranch *RepositoryBranch `json:"repositoryBranch,omitempty"`
}

type ScanIssueStatusChangeEntry struct {
	ID        uuid.UUID                         `json:"id,omitempty"`
	Created   string                            `json:"created,omitempty"`
	OldValue  *string                           `json:"oldValue,omitempty"`
	NewValue  string                            `json:"newValue,omitempty"`
	Automatic bool                              `json:"automatic,omitempty"`
	Source    *ScanIssueStatusChangeEntrySource `json:"source,omitempty"`
}

type ScanIssueStatusChangeEntrySource struct {
	Type                      IssueStatusChangeSource       `json:"type,omitempty"`
	ScanUpload                *IssueStatusChangeScanUpload  `json:"scanUpload,omitempty"`
	User                      *IssueStatusChangeUser        `json:"user,omitempty"`
	Integration               *IssueStatusChangeIntegration `json:"integration,omitempty"`
	Rule                      *IssueStatusChangeRule        `json:"rule,omitempty"`
	ScheduledStatusChangeInfo *ScheduledStatus              `json:"scheduledStatusChangeInfo,omitempty"`
	SourceScanIssue           *ScanIssueShortWithBranch     `json:"sourceScanIssue,omitempty"`
}

type ScanIssueType string

const (
	ScanIssueTypeVULNERABILITY   ScanIssueType = "VULNERABILITY"
	ScanIssueTypeSOURCECODE      ScanIssueType = "SOURCE_CODE"
	ScanIssueTypeDEPENDENCYCHECK ScanIssueType = "DEPENDENCY_CHECK"
)

type ScanIssueWithFPScore struct {
	ScanIssue
	FPScore *FPScore `json:"fpScore,omitempty"`
}

type ScanIssuesAIRequest struct {
	Request       ScanIssueAIRequestType `json:"request"`
	AIModel       *string                `json:"aiModel,omitempty"`
	IntegrationID *uuid.UUID             `json:"integrationID,omitempty"`
	AutoApply     *bool                  `json:"autoApply,omitempty"`
}

type ScanIssuesUpdate struct {
	Issues              []uuid.UUID `json:"issues,omitempty"`
	Status              string      `json:"status,omitempty"`
	Severity            Severity    `json:"severity,omitempty"`
	StatusExpiryTime    string      `json:"statusExpiryTime,omitempty"`
	StatusChangeComment *string     `json:"statusChangeComment,omitempty"`
}

type ScanPipelineAggregatedData struct {
	ByFindingState *NameBasedBySeverity `json:"by_finding_state,omitempty"`
}

type ScanPipelineCreateExtResponse struct {
	ID         uuid.UUID          `json:"id,omitempty"`
	ExternalID string             `json:"externalID,omitempty"`
	Status     ScanPipelineStatus `json:"status,omitempty"`
	Created    string             `json:"created,omitempty"`
}

type ScanPipelineCreateExt struct {
	Repository       string     `json:"repository,omitempty"`
	Product          string     `json:"product,omitempty"`
	RepositoryBranch string     `json:"repositoryBranch,omitempty"`
	Artifact         string     `json:"artifact,omitempty"`
	Asset            string     `json:"asset,omitempty"`
	SpaceID          *uuid.UUID `json:"spaceID,omitempty"`
	Name             string     `json:"name,omitempty"`
	ExternalID       string     `json:"externalID,omitempty"`
	CommitSHA        *string    `json:"commitSha,omitempty"`
}

type ScanPipelineDetailsExt struct {
	ID                   uuid.UUID                   `json:"id,omitempty"`
	Created              string                      `json:"created,omitempty"`
	Updated              string                      `json:"updated,omitempty"`
	Finished             *string                     `json:"finished,omitempty"`
	Error                *string                     `json:"error,omitempty"`
	Duration             int                         `json:"duration,omitempty"`
	ScanningTool         string                      `json:"scanningTool,omitempty"`
	AggregatedData       *ScanPipelineAggregatedData `json:"aggregatedData,omitempty"`
	AggregationCompleted bool                        `json:"aggregationCompleted,omitempty"`
	Status               ScanPipelineStatus          `json:"status,omitempty"`
	Relation             *ScanPipelineRelation       `json:"relation,omitempty"`
	ScanProfile          *OrchestrationScanProfile   `json:"scanProfile,omitempty"`
	ExternalID           *string                     `json:"externalID,omitempty"`
	CommitSHA            CommitSHA                   `json:"commitSha,omitempty"`
}

type ScanPipelineDetails struct {
	ScanPipeline
	CommitSHA CommitSHA `json:"commitSha,omitempty"`
}

type ScanPipelineInitiator struct {
	ID          uuid.UUID                 `json:"id,omitempty"`
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
	ID             uuid.UUID                   `json:"id,omitempty"`
	SpaceID        uuid.UUID                   `json:"spaceID,omitempty"`
	ScanPipelineID uuid.UUID                   `json:"scanPipelineID,omitempty"`
	Created        string                      `json:"created,omitempty"`
	Updated        string                      `json:"updated,omitempty"`
	Finished       string                      `json:"finished,omitempty"`
	Duration       int                         `json:"duration,omitempty"`
	Status         ScanPipelineJobStatus       `json:"status,omitempty"`
	Error          string                      `json:"error,omitempty"`
	ScanningTool   string                      `json:"scanningTool,omitempty"`
	CommitSHA      CommitSHA                   `json:"commitSha,omitempty"`
	Relation       *ScanPipelineRelation       `json:"relation,omitempty"`
	ScanUpload     *ScanUpload                 `json:"scanUpload,omitempty"`
	AggregatedData *ScanPipelineAggregatedData `json:"aggregatedData,omitempty"`
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
	ID                   uuid.UUID                `json:"id,omitempty"`
	ExternalID           *string                  `json:"externalID,omitempty"`
	Status               ScanPipelineStatus       `json:"status,omitempty"`
	ScanningTool         string                   `json:"scanningTool,omitempty"`
	RelationType         ScanPipelineRelationType `json:"relationType,omitempty"`
	RelationID           uuid.UUID                `json:"relationId,omitempty"`
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

type ScanPipelineRelation struct {
	ID                uuid.UUID                   `json:"id,omitempty"`
	Name              string                      `json:"name,omitempty"`
	Slug              string                      `json:"slug,omitempty"`
	Type              ScanPipelineRelationType    `json:"type,omitempty"`
	SubType           ScanPipelineRelationSubType `json:"subType,omitempty"`
	SubRelationID     string                      `json:"subRelationID,omitempty"`
	IsDefaultRelation bool                        `json:"isDefaultRelation,omitempty"`
	SubRelationSlug   string                      `json:"subRelationSlug,omitempty"`
}

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
	ID                   uuid.UUID                   `json:"id,omitempty"`
	SpaceID              uuid.UUID                   `json:"spaceID,omitempty"`
	ScanPipelineType     ScanPipelineType            `json:"scanPipelineType,omitempty"`
	Created              string                      `json:"created,omitempty"`
	Updated              string                      `json:"updated,omitempty"`
	Finished             string                      `json:"finished,omitempty"`
	Duration             int                         `json:"duration,omitempty"`
	Status               ScanPipelineStatus          `json:"status,omitempty"`
	Error                string                      `json:"error,omitempty"`
	ScanningTool         string                      `json:"scanningTool,omitempty"`
	Initiator            *ScanPipelineInitiator      `json:"initiator,omitempty"`
	Relation             *ScanPipelineRelation       `json:"relation,omitempty"`
	Engagement           *Engagement                 `json:"engagement,omitempty"`
	OrchestrationScan    *OrchestrationScan          `json:"orchestrationScan,omitempty"`
	AggregatedData       *ScanPipelineAggregatedData `json:"aggregatedData,omitempty"`
	AggregationCompleted bool                        `json:"aggregationCompleted,omitempty"`
	Profile              *OrchestrationScanProfile   `json:"profile,omitempty"`
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
	ID                   uuid.UUID          `json:"id,omitempty"`
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

type ScanUploadRelation struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Slug string    `json:"slug,omitempty"`
}

type ScanUpload struct {
	ID              uuid.UUID           `json:"id,omitempty"`
	Created         string              `json:"created,omitempty"`
	Parser          Parser              `json:"parser,omitempty"`
	TaskID          uuid.UUID           `json:"taskID,omitempty"`
	Repository      any                 `json:"repository,omitempty"`
	Artifact        *ScanUploadRelation `json:"artifact,omitempty"`
	ArtifactVersion *ScanUploadRelation `json:"artifactVersion,omitempty"`
	Product         any                 `json:"product,omitempty"`
	ScanPipelineID  uuid.UUID           `json:"scanPipelineID,omitempty"`
	ScanJobID       uuid.UUID           `json:"scanJobID,omitempty"`
}

type SingleScanIssueForExport struct {
	ScanIssueForExport
	ID uuid.UUID `json:"id"`
}

type UnitedScanIssueForExport struct {
	ScanIssueForExport
	IDs []uuid.UUID `json:"ids"`
}

type ScanUploadsBody struct {
	ScannerResultFile          []byte     `json:"scannerResultFile"`
	ScannerType                Parser     `json:"scannerType"`
	SpaceID                    *uuid.UUID `json:"spaceID,omitempty"`
	Repository                 *string    `json:"repository,omitempty"`
	RepositoryBranch           *string    `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool      `json:"repositoryBranchSetDefault,omitempty"`
	Product                    *string    `json:"product,omitempty"`
	Artifact                   *string    `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool      `json:"artifactVersionSetDefault,omitempty"`
	VCSURL                     *string    `json:"vcsURL,omitempty"`
	ScanPipelineJobID          uuid.UUID  `json:"scanPipelineJobID,omitempty"`
	ExternalScanPipeline       string     `json:"externalScanPipeline,omitempty"`
	ScanPipelineProduct        *string    `json:"scanPipelineProduct,omitempty"`
	FinishScanPipeline         *bool      `json:"finishScanPipeline,omitempty"`
	ScanPipelineName           *string    `json:"scanPipelineName,omitempty"`
	SkipAutoResolve            *bool      `json:"skipAutoResolve,omitempty"`
	CommitSHA                  *string    `json:"commitSha,omitempty"`
}

type ScanUploadsBulkDownloadBody struct {
	ScanUploads []uuid.UUID `json:"scanUploads,omitempty"`
}
