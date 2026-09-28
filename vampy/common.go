// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type AnyValue any

type AsyncExportFormatter string

const (
	AsyncExportFormatterJSON AsyncExportFormatter = "JSON"
	AsyncExportFormatterPDF  AsyncExportFormatter = "PDF"
	AsyncExportFormatterCSV  AsyncExportFormatter = "CSV"
)

type BinaryString []byte

type CommitSHA string

type DependencyRelationFilter string

const (
	DependencyRelationFilterDIRECT     DependencyRelationFilter = "DIRECT"
	DependencyRelationFilterTRANSITIVE DependencyRelationFilter = "TRANSITIVE"
	DependencyRelationFilterUNKNOWN    DependencyRelationFilter = "UNKNOWN"
)

type DependencyRelation string

const (
	DependencyRelationDIRECT     DependencyRelation = "DIRECT"
	DependencyRelationTRANSITIVE DependencyRelation = "TRANSITIVE"
)

type InformationParameter struct {
	MaxCount         *int `json:"maxCount,omitempty"`
	CurrentCount     *int `json:"currentCount,omitempty"`
	LeftCount        *int `json:"leftCount,omitempty"`
	IsAvailable      bool `json:"isAvailable,omitempty"`
	DisplayUnlimited bool `json:"displayUnlimited,omitempty"`
}

type ListUUID []UUID

type NotAccessibleByLicense struct {
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Type   string `json:"type,omitempty"`
}

type SettingsConstants struct {
	ScanIssueStatusMaxCustomStatuses int `json:"scanIssueStatusMaxCustomStatuses,omitempty"`
}

type Severity string

const (
	SeverityNONE     Severity = "NONE"
	SeverityLOW      Severity = "LOW"
	SeverityMEDIUM   Severity = "MEDIUM"
	SeverityHIGH     Severity = "HIGH"
	SeverityCRITICAL Severity = "CRITICAL"
	SeverityINFO     Severity = "INFO"
)

type SLAOverdueCardData struct {
	SimpleCardValue
	BySeverity
}

type UUID string

type DatasetUploadBody struct {
	File []byte `json:"file"`
}

type ImportDefectDojoBody struct {
	SpaceID         *UUID   `json:"spaceID,omitempty"`
	Hostname        string  `json:"hostname"`
	Token           string  `json:"token"`
	IgnoreSSLCheck  bool    `json:"ignoreSSLCheck,omitempty"`
	Timeout         *int    `json:"timeout,omitempty"`
	BatchSize       *int    `json:"batchSize,omitempty"`
	ProductIDs      *string `json:"productIDs,omitempty"`
	FindingStatuses *string `json:"findingStatuses,omitempty"`
	CreateProducts  bool    `json:"createProducts,omitempty"`
	CreateComments  bool    `json:"createComments,omitempty"`
}

type SnapshotsSnapshotIDBody struct {
	IsCurrent bool   `json:"isCurrent,omitempty"`
	Name      string `json:"name,omitempty"`
}

type SnapshotsUploadBody struct {
	File []byte `json:"file"`
}

type TokenWebhookBody any

type V1ScanBody struct {
	SpaceID                   *UUID   `json:"spaceID,omitempty"`
	Repository                *string `json:"repository,omitempty"`
	Product                   *string `json:"product,omitempty"`
	RepositoryBranch          *string `json:"repositoryBranch,omitempty"`
	RepositoryArchive         []byte  `json:"repositoryArchive,omitempty"`
	ArtifactArchive           []byte  `json:"artifactArchive,omitempty"`
	CreateMissing             *bool   `json:"createMissing,omitempty"`
	Scanner                   Scanner `json:"scanner,omitempty"`
	ScannerOptions            string  `json:"scannerOptions,omitempty"`
	DryRun                    bool    `json:"dryRun,omitempty"`
	Profile                   *string `json:"profile,omitempty"`
	Artifact                  *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault *bool   `json:"artifactVersionSetDefault,omitempty"`
	Asset                     *string `json:"asset,omitempty"`
	Endpoint                  *string `json:"endpoint,omitempty"`
}

type V1ScanUploadsBody struct {
	ScannerResultFile          []byte  `json:"scannerResultFile"`
	ScannerType                Parser  `json:"scannerType"`
	SpaceID                    *UUID   `json:"spaceID,omitempty"`
	Artifact                   *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool   `json:"artifactVersionSetDefault,omitempty"`
	Repository                 *string `json:"repository,omitempty"`
	RepositoryBranch           *string `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool   `json:"repositoryBranchSetDefault,omitempty"`
	Product                    *string `json:"product,omitempty"`
	VCSURL                     *string `json:"vcsURL,omitempty"`
	ExternalScanPipeline       string  `json:"externalScanPipeline,omitempty"`
	ScanPipelineProduct        *string `json:"scanPipelineProduct,omitempty"`
	FinishScanPipeline         *bool   `json:"finishScanPipeline,omitempty"`
	ScanPipelineName           *string `json:"scanPipelineName,omitempty"`
	SkipAutoResolve            *bool   `json:"skipAutoResolve,omitempty"`
	CommitSHA                  *string `json:"commitSha,omitempty"`
}
