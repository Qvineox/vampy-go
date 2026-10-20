// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type FindingCard struct {
	Finding
	DescriptionMD    string            `json:"descriptionMD,omitempty"`
	Solution         string            `json:"solution,omitempty"`
	SolutionMD       string            `json:"solutionMD,omitempty"`
	Data             map[string]any    `json:"data,omitempty"`
	Assets           []Asset           `json:"assets,omitempty"`
	Repository       *Repository       `json:"repository,omitempty"`
	RepositoryBranch *RepositoryBranch `json:"repositoryBranch,omitempty"`
	ScanIssue        *ScanIssue        `json:"scanIssue,omitempty"`
	ScanIssueType    ScanIssueType     `json:"scanIssueType,omitempty"`
}

type Finding struct {
	ID             uuid.UUID     `json:"id,omitempty"`
	Created        string        `json:"created,omitempty"`
	Status         FindingStatus `json:"status,omitempty"`
	Parser         Parser        `json:"parser,omitempty"`
	Title          string        `json:"title,omitempty"`
	Description    string        `json:"description,omitempty"`
	FilePath       string        `json:"filePath,omitempty"`
	FileLineno     *int          `json:"fileLineno,omitempty"`
	LibraryName    string        `json:"libraryName,omitempty"`
	LibraryVersion string        `json:"libraryVersion,omitempty"`
	CVEID          string        `json:"cveID,omitempty"`
	CodeSnippet    string        `json:"codeSnippet,omitempty"`
	Duplicate      bool          `json:"duplicate,omitempty"`
	Severity       Severity      `json:"severity,omitempty"`
	EngagementID   uuid.UUID     `json:"engagementID,omitempty"`
	ScanIssueID    uuid.UUID     `json:"scanIssueID,omitempty"`
	CommitSHA      CommitSHA     `json:"commitSha,omitempty"`
}

type FindingState string

const (
	FindingStateDETECTED   FindingState = "DETECTED"
	FindingStateRESOLVED   FindingState = "RESOLVED"
	FindingStateDUPLICATED FindingState = "DUPLICATED"
)

type FindingStatus string

const (
	FindingStatusINBOX     FindingStatus = "INBOX"
	FindingStatusPERFORMED FindingStatus = "PERFORMED"
)
