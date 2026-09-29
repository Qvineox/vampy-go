// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

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
	ID             UUID                          `json:"id,omitempty"`
	Name           string                        `json:"name,omitempty"`
	Description    string                        `json:"description,omitempty"`
	Created        string                        `json:"created,omitempty"`
	Updated        string                        `json:"updated,omitempty"`
	RiskScore      *int                          `json:"riskScore,omitempty"`
	ScanIssuesStat *ComponentShortScanIssuesStat `json:"scanIssuesStat,omitempty"`
	Versions       []ComponentShortVersions      `json:"versions,omitempty"`
	LastAggregated string                        `json:"lastAggregated,omitempty"`
	Permissions    []ComponentPermissionValue    `json:"permissions,omitempty"`
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
	ID             UUID                          `json:"id,omitempty"`
	Name           string                        `json:"name,omitempty"`
	Version        string                        `json:"version,omitempty"`
	ComponentID    UUID                          `json:"componentID,omitempty"`
	Description    string                        `json:"description,omitempty"`
	Created        string                        `json:"created,omitempty"`
	Updated        string                        `json:"updated,omitempty"`
	RiskScore      *int                          `json:"riskScore,omitempty"`
	ScanIssuesStat *ComponentShortScanIssuesStat `json:"scanIssuesStat,omitempty"`
	ActualVersion  string                        `json:"actualVersion,omitempty"`
	PackageURL     string                        `json:"packageUrl,omitempty"`
	PlatformEnum   string                        `json:"platformEnum,omitempty"`
	SwidTagID      string                        `json:"swidTagID,omitempty"`
	Classifier     string                        `json:"classifier,omitempty"`
	Filename       string                        `json:"filename,omitempty"`
	License        string                        `json:"license,omitempty"`
	Author         string                        `json:"author,omitempty"`
	Permissions    []ComponentPermissionValue    `json:"permissions,omitempty"`
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
