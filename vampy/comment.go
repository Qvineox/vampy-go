// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type CommentDetails struct {
	ID               UUID                     `json:"id,omitempty"`
	Body             string                   `json:"body,omitempty"`
	Status           CommentStatus            `json:"status,omitempty"`
	Type             CommentType              `json:"type,omitempty"`
	Created          string                   `json:"created,omitempty"`
	Updated          string                   `json:"updated,omitempty"`
	User             *CommentUser             `json:"user,omitempty"`
	Rule             *CommentRule             `json:"rule,omitempty"`
	ScanIssue        *ScanIssueShort          `json:"scanIssue,omitempty"`
	ExternalUserName *string                  `json:"externalUserName,omitempty"`
	External         *bool                    `json:"external,omitempty"`
	ExportRecordID   *UUID                    `json:"exportRecordID,omitempty"`
	Permissions      []CommentPermissionValue `json:"permissions,omitempty"`
}

type CommentPermissionValue string

const (
	CommentPermissionValueCOMMENTREAD CommentPermissionValue = "COMMENT_READ"
	CommentPermissionValueCOMMENTEDIT CommentPermissionValue = "COMMENT_EDIT"
)

type CommentRule struct {
	ID   UUID   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Comment struct {
	Body           string `json:"body"`
	External       bool   `json:"external,omitempty"`
	ExportRecordID UUID   `json:"exportRecordID,omitempty"`
}

type CommentStatus string

const (
	CommentStatusCREATED   CommentStatus = "CREATED"
	CommentStatusUPDATED   CommentStatus = "UPDATED"
	CommentStatusSYNCERROR CommentStatus = "SYNC_ERROR"
	CommentStatusSYNCED    CommentStatus = "SYNCED"
)

type CommentType string

const (
	CommentTypeLOCAL    CommentType = "LOCAL"
	CommentTypeEXTERNAL CommentType = "EXTERNAL"
)

type CommentUser struct {
	ID        UUID    `json:"id,omitempty"`
	Login     string  `json:"login,omitempty"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}
