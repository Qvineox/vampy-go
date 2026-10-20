// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type Action struct {
	Type       ActionType  `json:"type,omitempty"`
	Enabled    bool        `json:"enabled,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
}

type ActionType string

const (
	ActionTypeCHANGESCANISSUESTATUS   ActionType = "CHANGE_SCAN_ISSUE_STATUS"
	ActionTypeCHANGESCANISSUESEVERITY ActionType = "CHANGE_SCAN_ISSUE_SEVERITY"
	ActionTypeADDSCANISSUECOMMENT     ActionType = "ADD_SCAN_ISSUE_COMMENT"
	ActionTypeCREATEREPOSITORY        ActionType = "CREATE_REPOSITORY"
	ActionTypeSENDNOTIFICATION        ActionType = "SEND_NOTIFICATION"
	ActionTypeRUNSCAN                 ActionType = "RUN_SCAN"
)

type BaseRule struct {
	ID          uuid.UUID   `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Enabled     bool        `json:"enabled,omitempty"`
	Actions     []Action    `json:"actions,omitempty"`
	Conditions  []Condition `json:"conditions,omitempty"`
}

type ComparisonOperator string

const (
	ComparisonOperatorGT  ComparisonOperator = "GT"
	ComparisonOperatorLT  ComparisonOperator = "LT"
	ComparisonOperatorGTE ComparisonOperator = "GTE"
	ComparisonOperatorLTE ComparisonOperator = "LTE"
	ComparisonOperatorEQ  ComparisonOperator = "EQ"
	ComparisonOperatorNEQ ComparisonOperator = "NEQ"
)

type Condition struct {
	Type       ConditionType `json:"type"`
	Conditions []Condition   `json:"conditions"`
	Parameters *Parameter    `json:"parameters,omitempty"`
}

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

type Parameter struct {
	Attribute      string             `json:"attribute,omitempty"`
	AttributeValue any                `json:"attributeValue,omitempty"`
	Operator       ComparisonOperator `json:"operator,omitempty"`
	Warnings       []string           `json:"warnings,omitempty"`
}

type RuleCreate struct {
	BaseRule
	EventType   RulesEngineEventType `json:"eventType,omitempty"`
	ScanIssueID *uuid.UUID           `json:"scanIssueID,omitempty"`
	RunRule     bool                 `json:"runRule,omitempty"`
}

type RulePassedInstances struct {
	PassedInstances int `json:"passedInstances,omitempty"`
}

type RulePatch struct {
	BaseRule
	RunRule bool `json:"runRule,omitempty"`
}

type RuleRead struct {
	BaseRule
	EventType   RulesEngineEventType `json:"eventType,omitempty"`
	ScanIssueID *uuid.UUID           `json:"scanIssueID,omitempty"`
	Created     string               `json:"created,omitempty"`
	Updated     string               `json:"updated,omitempty"`
	Warnings    []RuleWarning        `json:"warnings,omitempty"`
}

type RuleShort struct {
	ID          uuid.UUID     `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Enabled     bool          `json:"enabled,omitempty"`
	Created     string        `json:"created,omitempty"`
	Updated     string        `json:"updated,omitempty"`
	ScanIssueID *uuid.UUID    `json:"scanIssueID,omitempty"`
	Warnings    []RuleWarning `json:"warnings,omitempty"`
}

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

type RulesSupportedEvents struct {
	EventType        RulesEngineEventType `json:"eventType,omitempty"`
	ManuallyRunnable bool                 `json:"manuallyRunnable,omitempty"`
	Actions          []SupportedAction    `json:"actions,omitempty"`
	Conditions       []SupportedCondition `json:"conditions,omitempty"`
}

type SupportedAction struct {
	Type       ActionType           `json:"type,omitempty"`
	Parameters []SupportedParameter `json:"parameters,omitempty"`
}

type SupportedCondition struct {
	Type               ConditionType `json:"type,omitempty"`
	Attribute          string        `json:"attribute,omitempty"`
	AttributeValueType FieldType     `json:"attributeValueType,omitempty"`
}

type SupportedParameter struct {
	Type      FieldType `json:"type,omitempty"`
	Attribute string    `json:"attribute,omitempty"`
	Required  bool      `json:"required,omitempty"`
}
