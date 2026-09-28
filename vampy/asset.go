// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type AssetCreate struct {
	Name             string                       `json:"name,omitempty"`
	Description      string                       `json:"description,omitempty"`
	URI              string                       `json:"uri"`
	IPs              []string                     `json:"ips,omitempty"`
	ProductID        UUID                         `json:"productID"`
	Scope            AssetScope                   `json:"scope,omitempty"`
	Criticality      AssetCriticality             `json:"criticality,omitempty"`
	DefinedVariables []OrchestrationScannerEnvVar `json:"definedVariables,omitempty"`
}

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
	LastScanPipeline      *ScanPipeline                `json:"lastScanPipeline,omitempty"`
}

type AssetForOrchestration struct {
	ID            UUID                         `json:"id,omitempty"`
	ParentID      *string                      `json:"parentID,omitempty"`
	Host          string                       `json:"host,omitempty"`
	ScanVariables []OrchestrationScannerEnvVar `json:"scanVariables,omitempty"`
}

type AssetPatch struct {
	Name             string                       `json:"name,omitempty"`
	Description      string                       `json:"description,omitempty"`
	Scope            AssetScope                   `json:"scope,omitempty"`
	Criticality      AssetCriticality             `json:"criticality,omitempty"`
	DefinedVariables []OrchestrationScannerEnvVar `json:"definedVariables,omitempty"`
}

type AssetPermissionValue string

const (
	AssetPermissionValueASSETREAD AssetPermissionValue = "ASSET_READ"
	AssetPermissionValueASSETEDIT AssetPermissionValue = "ASSET_EDIT"
)

type Asset struct {
	ID                    UUID                    `json:"id,omitempty"`
	Name                  string                  `json:"name,omitempty"`
	Scope                 AssetScope              `json:"scope,omitempty"`
	Criticality           AssetCriticality        `json:"criticality,omitempty"`
	Host                  string                  `json:"host,omitempty"`
	HostName              string                  `json:"hostName,omitempty"`
	URI                   string                  `json:"uri,omitempty"`
	ProductID             *UUID                   `json:"productID,omitempty"`
	ParentID              *UUID                   `json:"parentID,omitempty"`
	QueryStrings          []string                `json:"queryStrings,omitempty"`
	Permissions           []AssetPermissionValue  `json:"permissions,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail `json:"lastOrchestrationScan,omitempty"`
	LastScanPipeline      *ScanPipeline           `json:"lastScanPipeline,omitempty"`
}

type AssetScope string

const (
	AssetScopePROD  AssetScope = "PROD"
	AssetScopeSTAGE AssetScope = "STAGE"
	AssetScopeDEV   AssetScope = "DEV"
)
