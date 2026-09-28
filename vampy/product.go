// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type ProductChangeRelations struct {
	ComponentVersions []UUID `json:"componentVersions,omitempty"`
	Repositories      []UUID `json:"repositories,omitempty"`
}

type ProductCriticality string

const (
	ProductCriticalityLOW      ProductCriticality = "LOW"
	ProductCriticalityMEDIUM   ProductCriticality = "MEDIUM"
	ProductCriticalityHIGH     ProductCriticality = "HIGH"
	ProductCriticalityCRITICAL ProductCriticality = "CRITICAL"
)

type ProductDetail struct {
	Product
	RepositoryIDs         []UUID                    `json:"repositoryIDs,omitempty"`
	Integration           *SpaceIntegration         `json:"integration,omitempty"`
	ScanProfile           *OrchestrationScanProfile `json:"scanProfile,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail   `json:"lastOrchestrationScan,omitempty"`
	LastScanPipeline      *ScanPipeline             `json:"lastScanPipeline,omitempty"`
}

type ProductEdit struct {
	Name               string                 `json:"name,omitempty"`
	Description        string                 `json:"description,omitempty"`
	Criticality        ProductCriticality     `json:"criticality,omitempty"`
	IntegrationID      *UUID                  `json:"integrationID,omitempty"`
	IntegrationService IntegrationServiceType `json:"integrationService,omitempty"`
	QualityGateID      *UUID                  `json:"qualityGateID,omitempty"`
	ScanProfileID      *UUID                  `json:"scanProfileID,omitempty"`
}

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

type Product struct {
	ID                UUID                     `json:"id,omitempty"`
	Name              string                   `json:"name"`
	Slug              string                   `json:"slug,omitempty"`
	Created           string                   `json:"created,omitempty"`
	Description       string                   `json:"description,omitempty"`
	Criticality       ProductCriticality       `json:"criticality,omitempty"`
	RiskScore         *int                     `json:"riskScore,omitempty"`
	RepositoriesCount int                      `json:"repositoriesCount,omitempty"`
	ScanIssuesStat    *ProductScanIssuesStat   `json:"scanIssuesStat,omitempty"`
	LastAggregated    string                   `json:"lastAggregated,omitempty"`
	Users             []DetailPermissionUser   `json:"users,omitempty"`
	Permissions       []ProductPermissionValue `json:"permissions,omitempty"`
	QualityGate       *QualityGateForRelation  `json:"qualityGate,omitempty"`
	TitleTranslateID  *string                  `json:"titleTranslateID,omitempty"`
}

type ProductScanIssuesStat struct {
	BySeverity []ScanIssueBySeverityStat `json:"bySeverity,omitempty"`
	ByStatus   []ScanIssueByStatusStat   `json:"byStatus,omitempty"`
	Counts     *ScanIssueCountsStat      `json:"counts,omitempty"`
}

type ProductShort struct {
	ID          UUID               `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Criticality ProductCriticality `json:"criticality,omitempty"`
	Slug        string             `json:"slug,omitempty"`
	Created     string             `json:"created,omitempty"`
}

type ProductsCountCardData struct {
	SimpleCardValue
}
