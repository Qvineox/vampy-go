// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

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

type QualityGateCriticality string

const (
	QualityGateCriticalityLOW      QualityGateCriticality = "LOW"
	QualityGateCriticalityMEDIUM   QualityGateCriticality = "MEDIUM"
	QualityGateCriticalityHIGH     QualityGateCriticality = "HIGH"
	QualityGateCriticalityCRITICAL QualityGateCriticality = "CRITICAL"
)

type QualityGateEdit struct {
	Name       string                 `json:"name,omitempty"`
	Enabled    bool                   `json:"enabled,omitempty"`
	Conditions []QualityGateCondition `json:"conditions,omitempty"`
}

type QualityGateForRelation struct {
	ID         UUID                              `json:"id,omitempty"`
	Name       string                            `json:"name,omitempty"`
	Slug       string                            `json:"slug,omitempty"`
	LastResult *QualityGateForRelationLastResult `json:"lastResult,omitempty"`
}

type QualityGateForRelationLastResult struct {
	ID             *UUID                   `json:"id,omitempty"`
	Status         QualityGateResultStatus `json:"status,omitempty"`
	LastCalculated *string                 `json:"lastCalculated,omitempty"`
}

type QualityGateRelation struct {
	ID            UUID                       `json:"id,omitempty"`
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

type QualityGateResultDetails struct {
	ID             UUID                             `json:"id,omitempty"`
	Status         QualityGateResultStatus          `json:"status,omitempty"`
	LastCalculated *string                          `json:"lastCalculated,omitempty"`
	Relation       *QualityGateRelation             `json:"relation,omitempty"`
	Details        *QualityGateResultDetailsDetails `json:"details,omitempty"`
}

type QualityGateResultDetailsDetails struct {
	Slug         string                       `json:"slug,omitempty"`
	Name         string                       `json:"name,omitempty"`
	Criticality  QualityGateCriticality       `json:"criticality,omitempty"`
	Conditions   []QualityGateConditionResult `json:"conditions,omitempty"`
	ErrorMessage string                       `json:"errorMessage,omitempty"`
}

type QualityGateResult struct {
	ID             UUID                    `json:"id,omitempty"`
	Status         QualityGateResultStatus `json:"status,omitempty"`
	LastCalculated string                  `json:"lastCalculated,omitempty"`
	Relation       *QualityGateRelation    `json:"relation,omitempty"`
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
	ID          UUID                   `json:"id,omitempty"`
	Slug        string                 `json:"slug"`
	Name        string                 `json:"name"`
	IsDefault   bool                   `json:"isDefault,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Created     string                 `json:"created,omitempty"`
	Updated     *string                `json:"updated,omitempty"`
	Criticality QualityGateCriticality `json:"criticality,omitempty"`
	Conditions  []QualityGateCondition `json:"conditions,omitempty"`
}

type QualityGateSubRelationType string

const (
	QualityGateSubRelationTypeREPOSITORYBRANCH QualityGateSubRelationType = "REPOSITORY_BRANCH"
)
