// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type RepositoriesCountCardData struct {
	SimpleCardValue
}

type RepositoryBranchList struct {
	ID         uuid.UUID                       `json:"id,omitempty"`
	Slug       string                          `json:"slug,omitempty"`
	Created    string                          `json:"created,omitempty"`
	IsDefault  bool                            `json:"isDefault,omitempty"`
	IsEmpty    bool                            `json:"isEmpty,omitempty"`
	Repository *RepositoryBranchListRepository `json:"repository,omitempty"`
}

type RepositoryBranchListRepository struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Slug string    `json:"slug,omitempty"`
}

type RepositoryBranch struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	Created   string    `json:"created,omitempty"`
	IsDefault bool      `json:"isDefault,omitempty"`
}

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
	ProductID *uuid.UUID `json:"productID,omitempty"`
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
	Languages             map[string]any                 `json:"languages,omitempty"`
	ExternalLink          *RepositoryDetailsExternalLink `json:"externalLink,omitempty"`
	ScanProfile           *OrchestrationScanProfile      `json:"scanProfile,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail        `json:"lastOrchestrationScan,omitempty"`
	LastScanPipeline      *ScanPipeline                  `json:"lastScanPipeline,omitempty"`
}

type RepositoryDetailsExternalLink struct {
	Name               string                    `json:"name,omitempty"`
	URL                string                    `json:"url,omitempty"`
	IntegrationService IntegrationVCSServiceType `json:"integrationService,omitempty"`
	IntegrationState   IntegrationState          `json:"integrationState,omitempty"`
}

type RepositoryForExternalLink struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
}

type RepositoryImportRequest struct {
	SpaceID                uuid.UUID `json:"spaceID"`
	ProductID              uuid.UUID `json:"productID,omitempty"`
	IntegrationID          string    `json:"integrationID"`
	PopulateRepositoryData bool      `json:"populateRepositoryData,omitempty"`
	ExternalIDs            []string  `json:"externalIDs"`
}

type RepositoryImportResponse struct {
	Repositories []RepositoryImportResponseRepositories `json:"repositories,omitempty"`
}

type RepositoryImportResponseRepositories struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Slug string    `json:"slug,omitempty"`
}

type RepositoryLink struct {
	SpaceID                uuid.UUID `json:"spaceID"`
	IntegrationID          string    `json:"integrationID"`
	PopulateRepositoryData bool      `json:"populateRepositoryData,omitempty"`
	ExternalID             string    `json:"externalID"`
}

type RepositoryPatch struct {
	Criticality   RepositoryCriticality `json:"criticality,omitempty"`
	Name          string                `json:"name,omitempty"`
	URL           string                `json:"url,omitempty"`
	Description   string                `json:"description,omitempty"`
	QualityGateID *uuid.UUID            `json:"qualityGateID,omitempty"`
	ScanProfileID *uuid.UUID            `json:"scanProfileID,omitempty"`
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
	ID               uuid.UUID                   `json:"id,omitempty"`
	Type             RepositoryType              `json:"type,omitempty"`
	Slug             string                      `json:"slug"`
	Criticality      RepositoryCriticality       `json:"criticality,omitempty"`
	Name             string                      `json:"name"`
	Created          string                      `json:"created,omitempty"`
	LastScan         string                      `json:"lastScan,omitempty"`
	Description      string                      `json:"description,omitempty"`
	URL              string                      `json:"url"`
	RiskScore        *int                        `json:"riskScore,omitempty"`
	ScanIssuesStat   *ProductScanIssuesStat      `json:"scanIssuesStat,omitempty"`
	LastAggregated   string                      `json:"lastAggregated,omitempty"`
	Users            []DetailPermissionUser      `json:"users,omitempty"`
	Permissions      []RepositoryPermissionValue `json:"permissions,omitempty"`
	Branches         []RepositoryBranch          `json:"branches,omitempty"`
	DefaultBranch    string                      `json:"defaultBranch,omitempty"`
	QualityGate      *QualityGateForRelation     `json:"qualityGate,omitempty"`
	TitleTranslateID *string                     `json:"titleTranslateID,omitempty"`
}

type RepositoryShort struct {
	ID          uuid.UUID             `json:"id,omitempty"`
	Type        RepositoryType        `json:"type,omitempty"`
	Slug        string                `json:"slug,omitempty"`
	Criticality RepositoryCriticality `json:"criticality,omitempty"`
	Name        string                `json:"name,omitempty"`
	Created     string                `json:"created,omitempty"`
}

type RepositoryType string

const (
	RepositoryTypeSOURCECODE RepositoryType = "SOURCE_CODE"
	RepositoryTypeBINARY     RepositoryType = "BINARY"
)
