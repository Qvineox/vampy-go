// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type SBOMExport struct {
	Search       string `json:"search,omitempty"`
	Order        string `json:"order,omitempty"`
	ComponentID  UUID   `json:"componentID,omitempty"`
	RepositoryID UUID   `json:"repositoryID,omitempty"`
}

type SBOMImport struct {
	Created string `json:"created,omitempty"`
	TaskID  UUID   `json:"taskID,omitempty"`
}

type SBOMImportBody struct {
	SBOMResultFile []byte `json:"sbomResultFile"`
	RepositoryID   *UUID  `json:"repositoryID,omitempty"`
}
