// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type SBOMExport struct {
	Search       string    `json:"search,omitempty"`
	Order        string    `json:"order,omitempty"`
	ComponentID  uuid.UUID `json:"componentID,omitempty"`
	RepositoryID uuid.UUID `json:"repositoryID,omitempty"`
}

type SBOMImport struct {
	Created string    `json:"created,omitempty"`
	TaskID  uuid.UUID `json:"taskID,omitempty"`
}

type SBOMImportBody struct {
	SBOMResultFile []byte     `json:"sbomResultFile"`
	RepositoryID   *uuid.UUID `json:"repositoryID,omitempty"`
}
