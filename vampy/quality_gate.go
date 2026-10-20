// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import (
	"time"
	"uuid"
)

type QualityGateCondition struct {
	Enabled      bool          `json:"enabled,omitempty"`
	ScannerTypes []ScannerType `json:"scannerTypes,omitempty"`
	Severities   []Severity    `json:"severities,omitempty"`
	MaxValue     int           `json:"maxValue,omitempty"`
}

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

type QualityGateEdit struct {
	Name       string                 `json:"name,omitempty"`
	Enabled    bool                   `json:"enabled,omitempty"`
	Conditions []QualityGateCondition `json:"conditions,omitempty"`
}

type QualityGateForRelation struct {
	ID         uuid.UUID                         `json:"id,omitempty"`
	Name       string                            `json:"name,omitempty"`
	Slug       string                            `json:"slug,omitempty"`
	LastResult *QualityGateForRelationLastResult `json:"lastResult,omitempty"`
}

type QualityGateForRelationLastResult struct {
	ID             *uuid.UUID              `json:"id,omitempty"`
	Status         QualityGateResultStatus `json:"status,omitempty"`
	LastCalculated *time.Time              `json:"lastCalculated,omitempty"`
}

type QualityGateRelation struct {
	ID            uuid.UUID                  `json:"id,omitempty"`
	Name          string                     `json:"name,omitempty"`
	Slug          string                     `json:"slug,omitempty"`
	Type          QualityGateRelationType    `json:"type,omitempty"`
	SubType       QualityGateSubRelationType `json:"subType,omitempty"`
	SubRelationID string                     `json:"subRelationID,omitempty"`
}

type QualityGateRelationType string

const (
	QualityGateRelationTypePRODUCT    QualityGateRelationType = "PRODUCT"
	QualityGateRelationTypeREPOSITORY QualityGateRelationType = "REPOSITORY"
)

type QualityGateResult struct {
	ID             uuid.UUID                 `json:"id,omitempty"`
	Status         QualityGateResultStatus   `json:"status,omitempty"`
	LastCalculated *time.Time                `json:"lastCalculated,omitempty"`
	Relation       *QualityGateRelation      `json:"relation,omitempty"`
	Details        *QualityGateResultDetails `json:"details,omitempty"`
}

type QualityGateResultDetails struct {
	Slug         string                       `json:"slug,omitempty"`
	Name         string                       `json:"name,omitempty"`
	Criticality  Criticality                  `json:"criticality,omitempty"`
	Conditions   []QualityGateConditionResult `json:"conditions,omitempty"`
	ErrorMessage string                       `json:"errorMessage,omitempty"`
}

type QualityGateResultStatus string

const (
	QualityGateResultStatusINPROGRESS QualityGateResultStatus = "IN_PROGRESS"
	QualityGateResultStatusTIMEDOUT   QualityGateResultStatus = "TIMED_OUT"
	QualityGateResultStatusSKIPPED    QualityGateResultStatus = "SKIPPED"
	QualityGateResultStatusFAILED     QualityGateResultStatus = "FAILED"
	QualityGateResultStatusPASSED     QualityGateResultStatus = "PASSED"
	QualityGateResultStatusERROR      QualityGateResultStatus = "ERROR"
)

type QualityGate struct {
	ID          uuid.UUID              `json:"id,omitempty"`
	Slug        string                 `json:"slug"`
	Name        string                 `json:"name"`
	IsDefault   bool                   `json:"isDefault,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Created     time.Time              `json:"created"`
	Updated     *time.Time             `json:"updated,omitempty"`
	Criticality Criticality            `json:"criticality,omitempty"`
	Conditions  []QualityGateCondition `json:"conditions,omitempty"`
}

type QualityGateSubRelationType string

const (
	QualityGateSubRelationTypeREPOSITORYBRANCH QualityGateSubRelationType = "REPOSITORY_BRANCH"
)
