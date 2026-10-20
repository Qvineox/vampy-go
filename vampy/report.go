// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type ReportFileFormat string

const (
	ReportFileFormatJSON ReportFileFormat = "JSON"
	ReportFileFormatPDF  ReportFileFormat = "PDF"
)

type Report struct {
	ID               uuid.UUID `json:"id,omitempty"`
	FileFormat       any       `json:"fileFormat,omitempty"`
	Status           any       `json:"status,omitempty"`
	CreatedByID      uuid.UUID `json:"createdById,omitempty"`
	BackgroundTaskID uuid.UUID `json:"backgroundTaskId,omitempty"`
	ErrorMessage     string    `json:"errorMessage,omitempty"`
	Created          string    `json:"created,omitempty"`
	Updated          string    `json:"updated,omitempty"`
}

type ReportStatus string

const (
	ReportStatusSUCCESS  ReportStatus = "SUCCESS"
	ReportStatusPENDING  ReportStatus = "PENDING"
	ReportStatusPROGRESS ReportStatus = "PROGRESS"
	ReportStatusFAILED   ReportStatus = "FAILED"
)
