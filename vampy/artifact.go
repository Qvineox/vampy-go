// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type ArtifactDelete struct {
	Artifacts []uuid.UUID `json:"artifacts,omitempty"`
	DeleteAll bool        `json:"deleteAll,omitempty"`
	Search    *string     `json:"search,omitempty"`
}

type ArtifactDetail struct {
	Artifact
	LastOrchestrationScan *OrchestrationScanDetail `json:"lastOrchestrationScan,omitempty"`
}

type ArtifactEdit struct {
	DefaultVersion string `json:"defaultVersion,omitempty"`
}

type ArtifactImport struct {
	ProductID          uuid.UUID `json:"productID"`
	IntegrationID      uuid.UUID `json:"integrationID"`
	IntegrationService string    `json:"integrationService"`
	ImportAll          bool      `json:"importAll,omitempty"`
	SearchQuery        string    `json:"searchQuery,omitempty"`
	Items              []string  `json:"items"`
}

type ArtifactImportTask struct {
	ID      string    `json:"id,omitempty"`
	Created string    `json:"created,omitempty"`
	TaskID  uuid.UUID `json:"taskID,omitempty"`
}

type Artifact struct {
	ArtifactShort
	ScanIssuesStat        *ProductScanIssuesStat  `json:"scanIssuesStat,omitempty"`
	VersionsCount         int                     `json:"versionsCount,omitempty"`
	Integration           *SpaceIntegration       `json:"integration,omitempty"`
	LastOrchestrationScan OrchestrationScanDetail `json:"lastOrchestrationScan,omitempty"`
}

type ArtifactShort struct {
	ID             uuid.UUID    `json:"id,omitempty"`
	Slug           string       `json:"slug,omitempty"`
	Description    string       `json:"description,omitempty"`
	DefaultVersion string       `json:"defaultVersion,omitempty"`
	Updated        string       `json:"updated,omitempty"`
	Created        string       `json:"created,omitempty"`
	Type           ArtifactType `json:"type,omitempty"`
}

type ArtifactType string

const (
	ArtifactTypeDOCKER ArtifactType = "DOCKER"
)

type ArtifactVersionDelete struct {
	ArtifactVersions []uuid.UUID `json:"artifactVersions,omitempty"`
	DeleteAll        bool        `json:"deleteAll,omitempty"`
	Search           *string     `json:"search,omitempty"`
}

type ArtifactVersion struct {
	ID             uuid.UUID              `json:"id,omitempty"`
	Slug           uuid.UUID              `json:"slug,omitempty"`
	Digests        []string               `json:"digests,omitempty"`
	Pushed         string                 `json:"pushed,omitempty"`
	ScanIssuesStat *ProductScanIssuesStat `json:"scanIssuesStat,omitempty"`
	LastScan       string                 `json:"lastScan,omitempty"`
	PullsCount     *int                   `json:"pullsCount,omitempty"`
	IsDefault      bool                   `json:"isDefault,omitempty"`
	Artifact       *ArtifactShort         `json:"artifact,omitempty"`
}

type ArtifactVersionShort struct {
	ID        uuid.UUID      `json:"id,omitempty"`
	Slug      string         `json:"slug,omitempty"`
	Artifact  *ArtifactShort `json:"artifact,omitempty"`
	IsDefault bool           `json:"isDefault,omitempty"`
}

type ArtifactVersionSyncArtifact struct {
	ArtifactID         uuid.UUID `json:"artifactID,omitempty"`
	ArtifactSlug       string    `json:"artifactSlug,omitempty"`
	ProductID          uuid.UUID `json:"productID,omitempty"`
	ProductSlug        string    `json:"productSlug,omitempty"`
	IntegrationID      uuid.UUID `json:"integrationID,omitempty"`
	IntegrationService string    `json:"integrationService,omitempty"`
	VersionsCount      int       `json:"versionsCount,omitempty"`
}

type ArtifactVersionSyncRequest struct {
	Artifact string  `json:"artifact"`
	Product  *string `json:"product,omitempty"`
}

type ArtifactVersionSyncResponse struct {
	Artifact        string                        `json:"artifact,omitempty"`
	Product         *string                       `json:"product,omitempty"`
	TotalArtifacts  int                           `json:"totalArtifacts,omitempty"`
	SyncedArtifacts []ArtifactVersionSyncArtifact `json:"syncedArtifacts,omitempty"`
}
