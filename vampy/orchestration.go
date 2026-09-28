// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type OrchestrationRunBatchResult struct {
	Items []OrchestrationRunResult `json:"items,omitempty"`
}

type OrchestrationRunMode string

const (
	OrchestrationRunModeCONCURRENT OrchestrationRunMode = "CONCURRENT"
	OrchestrationRunModeCHAIN      OrchestrationRunMode = "CHAIN"
)

type OrchestrationRunResult struct {
	ID                UUID                      `json:"id,omitempty"`
	Start             string                    `json:"start,omitempty"`
	TaskID            UUID                      `json:"taskID,omitempty"`
	PipelineID        *UUID                     `json:"pipelineID,omitempty"`
	Repository        any                       `json:"repository,omitempty"`
	Product           any                       `json:"product,omitempty"`
	ArtifactVersion   *ScanUploadRelation       `json:"artifactVersion,omitempty"`
	Asset             *AssetForOrchestration    `json:"asset,omitempty"`
	ScanProfile       *OrchestrationScanProfile `json:"scanProfile,omitempty"`
	OrchestrationScan *OrchestrationScan        `json:"orchestrationScan,omitempty"`
}

type OrchestrationScanDetail struct {
	OrchestrationScan
	Profile *OrchestrationScanProfile `json:"profile,omitempty"`
}

type OrchestrationScanProfileRelation struct {
	ID   UUID                                 `json:"id,omitempty"`
	Name string                               `json:"name,omitempty"`
	Slug string                               `json:"slug,omitempty"`
	Type OrchestrationScanProfileRelationType `json:"type,omitempty"`
}

type OrchestrationScanProfileRelationType string

const (
	OrchestrationScanProfileRelationTypeREPOSITORY OrchestrationScanProfileRelationType = "REPOSITORY"
	OrchestrationScanProfileRelationTypePRODUCT    OrchestrationScanProfileRelationType = "PRODUCT"
)

type OrchestrationScanProfile struct {
	ID                 UUID                   `json:"id,omitempty"`
	Name               string                 `json:"name"`
	Slug               string                 `json:"slug,omitempty"`
	Persistence        bool                   `json:"persistence,omitempty"`
	Scanners           []OrchestrationScanner `json:"scanners"`
	RunMode            OrchestrationRunMode   `json:"runMode"`
	AvailabilityStatus *AvailabilityStatus    `json:"availabilityStatus,omitempty"`
}

type OrchestrationScan struct {
	ID               UUID                           `json:"id,omitempty"`
	Status           OrchestrationScanStatus        `json:"status,omitempty"`
	TargetType       OrchestrationScanTargetType    `json:"targetType,omitempty"`
	TargetID         UUID                           `json:"targetID,omitempty"`
	SubTargetType    OrchestrationScanSubTargetType `json:"subTargetType,omitempty"`
	Created          string                         `json:"created,omitempty"`
	Updated          string                         `json:"updated,omitempty"`
	BackgroundTaskID UUID                           `json:"backgroundTaskID,omitempty"`
	ErrorMessage     string                         `json:"error_message,omitempty"`
}

type OrchestrationScanStatus string

const (
	OrchestrationScanStatusPENDING  OrchestrationScanStatus = "PENDING"
	OrchestrationScanStatusPROGRESS OrchestrationScanStatus = "PROGRESS"
	OrchestrationScanStatusSUCCESS  OrchestrationScanStatus = "SUCCESS"
	OrchestrationScanStatusFAILED   OrchestrationScanStatus = "FAILED"
)

type OrchestrationScanSubTargetType string

const (
	OrchestrationScanSubTargetTypeREPOSITORYBRANCH OrchestrationScanSubTargetType = "REPOSITORY_BRANCH"
)

type OrchestrationScanTargetType string

const (
	OrchestrationScanTargetTypeREPOSITORY      OrchestrationScanTargetType = "REPOSITORY"
	OrchestrationScanTargetTypePRODUCT         OrchestrationScanTargetType = "PRODUCT"
	OrchestrationScanTargetTypeARTIFACTVERSION OrchestrationScanTargetType = "ARTIFACT_VERSION"
	OrchestrationScanTargetTypeASSET           OrchestrationScanTargetType = "ASSET"
)

type OrchestrationScanUnavailabilityReason string

const (
	OrchestrationScanUnavailabilityReasonSCANNERINTEGRATIONUNAVAILABLE     OrchestrationScanUnavailabilityReason = "SCANNER_INTEGRATION_UNAVAILABLE"
	OrchestrationScanUnavailabilityReasonSCANNERINTEGRATIONLICENSEERROR    OrchestrationScanUnavailabilityReason = "SCANNER_INTEGRATION_LICENSE_ERROR"
	OrchestrationScanUnavailabilityReasonSCANNERINTEGRATIONCLIREQUIRED     OrchestrationScanUnavailabilityReason = "SCANNER_INTEGRATION_CLI_REQUIRED"
	OrchestrationScanUnavailabilityReasonSCANNERINAPPLICABLEFORCONTEXT     OrchestrationScanUnavailabilityReason = "SCANNER_INAPPLICABLE_FOR_CONTEXT"
	OrchestrationScanUnavailabilityReasonSCANNERSETTINGSERROR              OrchestrationScanUnavailabilityReason = "SCANNER_SETTINGS_ERROR"
	OrchestrationScanUnavailabilityReasonSCANNERDISABLED                   OrchestrationScanUnavailabilityReason = "SCANNER_DISABLED"
	OrchestrationScanUnavailabilityReasonSCANTARGETINTEGRATIONLICENSEERROR OrchestrationScanUnavailabilityReason = "SCAN_TARGET_INTEGRATION_LICENSE_ERROR"
	OrchestrationScanUnavailabilityReasonSCANTARGETINTEGRATIONNOTFOUND     OrchestrationScanUnavailabilityReason = "SCAN_TARGET_INTEGRATION_NOT_FOUND"
	OrchestrationScanUnavailabilityReasonSCANTARGETINTEGRATIONUNAVAILABLE  OrchestrationScanUnavailabilityReason = "SCAN_TARGET_INTEGRATION_UNAVAILABLE"
	OrchestrationScanUnavailabilityReasonNOPERMISSIONSACCESS               OrchestrationScanUnavailabilityReason = "NO_PERMISSIONS_ACCESS"
	OrchestrationScanUnavailabilityReasonNOAVAILABLESCANNERS               OrchestrationScanUnavailabilityReason = "NO_AVAILABLE_SCANNERS"
	OrchestrationScanUnavailabilityReasonNOAVAILABLETARGET                 OrchestrationScanUnavailabilityReason = "NO_AVAILABLE_TARGET"
)

type OrchestrationScannerConfigDelete struct {
	FilePath string `json:"filePath"`
}

type OrchestrationScannerConfigFileInfo struct {
	FileName string `json:"fileName,omitempty"`
	FilePath string `json:"filePath,omitempty"`
}

type OrchestrationScannerConfigFileParams struct {
	UseDefault *bool   `json:"useDefault,omitempty"`
	FileName   *string `json:"fileName,omitempty"`
	FilePath   *string `json:"filePath,omitempty"`
}

type OrchestrationScannerEnvVar struct {
	Name        string  `json:"name,omitempty"`
	Value       string  `json:"value,omitempty"`
	Description *string `json:"description,omitempty"`
}

type OrchestrationScannerInfoEnvVar struct {
	OrchestrationScannerEnvVar
	Required bool `json:"required,omitempty"`
}

type OrchestrationScannerInfo struct {
	Category             ScannerCategory                  `json:"category,omitempty"`
	Scanner              Scanner                          `json:"scanner,omitempty"`
	Name                 *string                          `json:"name,omitempty"`
	DockerImage          *string                          `json:"dockerImage,omitempty"`
	Params               []OrchestrationScannerParam      `json:"params,omitempty"`
	ExtraEnvVars         []OrchestrationScannerInfoEnvVar `json:"extraEnvVars,omitempty"`
	IntegrationService   IntegrationScannerServiceType    `json:"integrationService,omitempty"`
	Enabled              bool                             `json:"enabled,omitempty"`
	AvailabilityStatus   *AvailabilityStatus              `json:"availabilityStatus,omitempty"`
	ConfigFileRequired   bool                             `json:"configFileRequired,omitempty"`
	HasDefaultConfigFile bool                             `json:"hasDefaultConfigFile,omitempty"`
}

type OrchestrationScannerParam struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Default     string `json:"default,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

type OrchestrationScanner struct {
	Category             ScannerCategory                       `json:"category,omitempty"`
	Scanner              ScannerWithCustom                     `json:"scanner,omitempty"`
	Name                 string                                `json:"name,omitempty"`
	Enabled              bool                                  `json:"enabled,omitempty"`
	Command              string                                `json:"command,omitempty"`
	DockerImage          *string                               `json:"dockerImage,omitempty"`
	ExtraEnvVars         []OrchestrationScannerEnvVar          `json:"extraEnvVars,omitempty"`
	IntegrationID        *UUID                                 `json:"integrationID,omitempty"`
	IntegrationService   IntegrationScannerServiceType         `json:"integrationService,omitempty"`
	AvailabilityStatus   *AvailabilityStatus                   `json:"availabilityStatus,omitempty"`
	ConfigFileParams     *OrchestrationScannerConfigFileParams `json:"configFileParams,omitempty"`
	ConfigFileRequired   bool                                  `json:"configFileRequired,omitempty"`
	HasDefaultConfigFile bool                                  `json:"hasDefaultConfigFile,omitempty"`
}

type OrchestrationRunBatchBody struct {
	SpaceID            *UUID   `json:"spaceID,omitempty"`
	ArtifactIDs        *string `json:"artifactIDs,omitempty"`
	ArtifactVersionIDs *string `json:"artifactVersionIDs,omitempty"`
	AssetIDs           *string `json:"assetIDs,omitempty"`
	Scanner            Scanner `json:"scanner,omitempty"`
	ScannerOptions     string  `json:"scannerOptions,omitempty"`
	Profile            *string `json:"profile,omitempty"`
}

type OrchestrationRunBody struct {
	SpaceID                    *UUID   `json:"spaceID,omitempty"`
	Repository                 *string `json:"repository,omitempty"`
	Product                    *string `json:"product,omitempty"`
	RepositoryBranch           *string `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool   `json:"repositoryBranchSetDefault,omitempty"`
	RepositoryArchive          []byte  `json:"repositoryArchive,omitempty"`
	ArtifactArchive            []byte  `json:"artifactArchive,omitempty"`
	CreateMissing              *bool   `json:"createMissing,omitempty"`
	Artifact                   *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool   `json:"artifactVersionSetDefault,omitempty"`
	Asset                      *string `json:"asset,omitempty"`
	Endpoint                   *string `json:"endpoint,omitempty"`
	Scanner                    Scanner `json:"scanner,omitempty"`
	ScannerOptions             string  `json:"scannerOptions,omitempty"`
	Profile                    *string `json:"profile,omitempty"`
}

type OrchestrationRunBody1 struct {
	SpaceID                    *UUID   `json:"spaceID,omitempty"`
	Repository                 *string `json:"repository,omitempty"`
	Product                    *string `json:"product,omitempty"`
	RepositoryBranch           *string `json:"repositoryBranch,omitempty"`
	RepositoryBranchSetDefault *bool   `json:"repositoryBranchSetDefault,omitempty"`
	RepositoryArchive          []byte  `json:"repositoryArchive,omitempty"`
	ArtifactArchive            []byte  `json:"artifactArchive,omitempty"`
	CreateMissing              *bool   `json:"createMissing,omitempty"`
	Artifact                   *string `json:"artifact,omitempty"`
	ArtifactVersionSetDefault  *bool   `json:"artifactVersionSetDefault,omitempty"`
	Asset                      *string `json:"asset,omitempty"`
	Scanner                    Scanner `json:"scanner,omitempty"`
	ScannerOptions             string  `json:"scannerOptions,omitempty"`
	DryRun                     bool    `json:"dryRun,omitempty"`
	Profile                    *string `json:"profile,omitempty"`
}

type OrchestrationUploadScannerConfigBody struct {
	Scanner ScannerWithConfigFiles `json:"scanner"`
	File    []byte                 `json:"file"`
}
