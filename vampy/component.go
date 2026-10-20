// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import (
	"time"
	"uuid"
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
	ID             uuid.UUID                     `json:"id,omitempty"`
	Name           string                        `json:"name,omitempty"`
	Description    string                        `json:"description,omitempty"`
	Created        time.Time                     `json:"created"`
	Updated        time.Time                     `json:"updated"`
	RiskScore      *int                          `json:"riskScore,omitempty"`
	ScanIssuesStat *ComponentShortScanIssuesStat `json:"scanIssuesStat,omitempty"`
	Versions       []ComponentShortVersions      `json:"versions,omitempty"`
	LastAggregated time.Time                     `json:"lastAggregated,omitempty"`
	Permissions    []ComponentPermissionValue    `json:"permissions,omitempty"`
}

type ComponentShortScanIssuesStat struct {
	BySeverity []ScanIssueBySeverityStat `json:"bySeverity,omitempty"`
}

type ComponentShortVersions struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Version string    `json:"version,omitempty"`
}

type ComponentVersionDetails struct {
	ComponentVersion
	Repositories []RepositoryShort `json:"repositories,omitempty"`
}

type ComponentVersion struct {
	ID             uuid.UUID                     `json:"id,omitempty"`
	Name           string                        `json:"name,omitempty"`
	Version        string                        `json:"version,omitempty"`
	ComponentID    uuid.UUID                     `json:"componentID,omitempty"`
	Description    string                        `json:"description,omitempty"`
	Created        time.Time                     `json:"created"`
	Updated        time.Time                     `json:"updated"`
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
	ID            uuid.UUID `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Version       string    `json:"version,omitempty"`
	ComponentID   uuid.UUID `json:"componentID,omitempty"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	ActualVersion string    `json:"actualVersion,omitempty"`
	PackageURL    string    `json:"packageUrl,omitempty"`
}

type ComponentsCountCardData struct {
	SimpleCardValue
}
