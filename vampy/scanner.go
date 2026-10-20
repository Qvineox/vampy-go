// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type ParserInList struct {
	Slug                 Parser             `json:"slug,omitempty"`
	DisplayName          string             `json:"displayName,omitempty"`
	Source               ParserSource       `json:"source,omitempty"`
	ScannerType          ScannerType        `json:"scannerType,omitempty"`
	NeedRepository       bool               `json:"needRepository,omitempty"`
	NeedProduct          bool               `json:"needProduct,omitempty"`
	SupportOrchestration bool               `json:"supportOrchestration,omitempty"`
	ContentType          ScannerContentType `json:"contentType,omitempty"`
}

type Parser string

const (
	ParserAPPSCREENERDAST         Parser = "APP_SCREENER_DAST"
	ParserAPPSCREENERSAST         Parser = "APP_SCREENER_SAST"
	ParserAPPSCREENERSCA          Parser = "APP_SCREENER_SCA"
	ParserSEMGREP                 Parser = "SEMGREP"
	ParserTRIVYFS                 Parser = "TRIVY_FS"
	ParserTRIVYIMAGE              Parser = "TRIVY_IMAGE"
	ParserTRUFFLEHOG              Parser = "TRUFFLEHOG"
	ParserTRUFFLEHOG3             Parser = "TRUFFLEHOG3"
	ParserESLINT                  Parser = "ESLINT"
	ParserGRYPE                   Parser = "GRYPE"
	ParserLUNTRY                  Parser = "LUNTRY"
	ParserNUCLEI                  Parser = "NUCLEI"
	ParserDEPENDENCYCHECK         Parser = "DEPENDENCY_CHECK"
	ParserGITLABCONTAINERSECURITY Parser = "GITLAB_CONTAINER_SECURITY"
	ParserGITLABDAST              Parser = "GITLAB_DAST"
	ParserGITLABDASTZAP           Parser = "GITLAB_DAST_ZAP"
	ParserGITLABDEPENDENCY        Parser = "GITLAB_DEPENDENCY"
	ParserGITLABSAST              Parser = "GITLAB_SAST"
	ParserGITLABSASTBRAKEMAN      Parser = "GITLAB_SAST_BRAKEMAN"
	ParserGITLABSASTFLAWFINDER    Parser = "GITLAB_SAST_FLAWFINDER"
	ParserGITLABSASTKUBESEC       Parser = "GITLAB_SAST_KUBESEC"
	ParserGITLABSASTMOBSF         Parser = "GITLAB_SAST_MOBSF"
	ParserGITLABSASTNODEJS        Parser = "GITLAB_SAST_NODEJS"
	ParserGITLABSASTPHPCS         Parser = "GITLAB_SAST_PHPCS"
	ParserGITLABSASTPMD           Parser = "GITLAB_SAST_PMD"
	ParserGITLABSASTSEMGREP       Parser = "GITLAB_SAST_SEMGREP"
	ParserGITLABSASTSOBELOW       Parser = "GITLAB_SAST_SOBELOW"
	ParserGITLABSASTSPOTBUGS      Parser = "GITLAB_SAST_SPOTBUGS"
	ParserGITLABSECRETDETECT      Parser = "GITLAB_SECRET_DETECT"
	ParserGITLEAKS                Parser = "GITLEAKS"
	ParserCHECKMARX               Parser = "CHECKMARX"
	ParserSARIF                   Parser = "SARIF"
	ParserSARIFCHECKOV            Parser = "SARIF_CHECKOV"
	ParserSARIFCODEQL             Parser = "SARIF_CODEQL"
	ParserSARIFCODESCORING        Parser = "SARIF_CODE_SCORING"
	ParserSARIFGITLEAKS           Parser = "SARIF_GITLEAKS"
	ParserSARIFGRYPE              Parser = "SARIF_GRYPE"
	ParserSARIFKICS               Parser = "SARIF_KICS"
	ParserSARIFPVSSTUDIO          Parser = "SARIF_PVS_STUDIO"
	ParserSARIFSASTAV             Parser = "SARIF_SASTAV"
	ParserSARIFSVACER             Parser = "SARIF_SVACER"
	ParserSBOM                    Parser = "SBOM"
	ParserSBOMCODESCORING         Parser = "SBOM_CODE_SCORING"
	ParserSBOMAPPSECTRACK         Parser = "SBOM_APPSEC_TRACK"
	ParserSNYKCODE                Parser = "SNYK_CODE"
	ParserSNYKSCA                 Parser = "SNYK_SCA"
	ParserSPOTBUGS                Parser = "SPOTBUGS"
	ParserPTAI                    Parser = "PT_AI"
	ParserPVSSTUDIO               Parser = "PVS_STUDIO"
	ParserSONARQUBE               Parser = "SONAR_QUBE"
	ParserZAP                     Parser = "ZAP"
	ParserACUNETIX                Parser = "ACUNETIX"
	ParserBURPSUITE               Parser = "BURP_SUITE"
	ParserVAMPY                   Parser = "VAMPY"
	ParserMANUAL                  Parser = "MANUAL"
)

type ParserSource string

const (
	ParserSourceSCANNER   ParserSource = "SCANNER"
	ParserSourceMANUAL    ParserSource = "MANUAL"
	ParserSourceSUBPARSER ParserSource = "SUB_PARSER"
)

type ScannerCategory string

const (
	ScannerCategoryDEFAULT     ScannerCategory = "DEFAULT"
	ScannerCategoryCUSTOM      ScannerCategory = "CUSTOM"
	ScannerCategoryINTEGRATION ScannerCategory = "INTEGRATION"
	ScannerCategoryEXTERNAL    ScannerCategory = "EXTERNAL"
)

type ScannerContentType string

const (
	ScannerContentTypeApplicationJSON ScannerContentType = "application/json"
	ScannerContentTypeTextXML         ScannerContentType = "text/xml"
)

type ScannerExternal struct {
	Slug        Scanner                       `json:"slug,omitempty"`
	DisplayName string                        `json:"displayName,omitempty"`
	Targets     []OrchestrationScanTargetType `json:"targets,omitempty"`
	Type        ScannerType                   `json:"type,omitempty"`
}

type ScannerIntegrationCLIFileInfo struct {
	FileName   string  `json:"fileName,omitempty"`
	Version    *string `json:"version,omitempty"`
	JDKVersion *string `json:"JDKVersion,omitempty"`
	Path       string  `json:"path,omitempty"`
}

type ScannerIntegrationCLI struct {
	Source           ScannerIntegrationCLISource `json:"source,omitempty"`
	FileName         *string                     `json:"fileName,omitempty"`
	Version          *string                     `json:"version,omitempty"`
	JDKVersion       *string                     `json:"JDKVersion,omitempty"`
	Updated          *string                     `json:"updated,omitempty"`
	ManualFileName   *string                     `json:"manualFileName,omitempty"`
	ManualVersion    *string                     `json:"manualVersion,omitempty"`
	ManualJDKVersion *string                     `json:"manualJDKVersion,omitempty"`
	ManualUpdated    *string                     `json:"manualUpdated,omitempty"`
}

type ScannerIntegrationCLISource string

const (
	ScannerIntegrationCLISourceMANUAL        ScannerIntegrationCLISource = "MANUAL"
	ScannerIntegrationCLISourceFETCHONDEMAND ScannerIntegrationCLISource = "FETCH_ON_DEMAND"
)

type ScannerIntegrationInfoRequest struct {
	Hostname       string  `json:"hostname,omitempty"`
	IgnoreSSLCheck bool    `json:"ignoreSSLCheck,omitempty"`
	ServerCert     *string `json:"serverCert,omitempty"`
}

type ScannerIntegrationInfoResponse struct {
	Authentication      string `json:"authentication"`
	IsCLIUploadRequired bool   `json:"isCLIUploadRequired"`
}

type ScannerIntegrationMetaResponse struct {
	Version             string `json:"version,omitempty"`
	IsAvailableFetchCLI bool   `json:"isAvailableFetchCLI,omitempty"`
}

type ScannerIntegrationPatch struct {
	ScannerIntegration
	CLISource     ScannerIntegrationCLISource `json:"cliSource,omitempty"`
	CLIFilePath   *string                     `json:"cliFilePath,omitempty"`
	CLIFileName   *string                     `json:"cliFileName,omitempty"`
	CLIVersion    *string                     `json:"cliVersion,omitempty"`
	CLIJDKVersion *string                     `json:"cliJDKVersion,omitempty"`
}

type ScannerIntegration struct {
	ID             uuid.UUID              `json:"id,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Hostname       string                 `json:"hostname,omitempty"`
	Token          string                 `json:"token,omitempty"`
	Username       string                 `json:"username,omitempty"`
	Service        IntegrationServiceType `json:"service,omitempty"`
	Roles          []IntegrationRole      `json:"roles,omitempty"`
	State          IntegrationState       `json:"state,omitempty"`
	ServerID       uuid.UUID              `json:"serverID,omitempty"`
	Created        string                 `json:"created,omitempty"`
	IgnoreSSLCheck bool                   `json:"ignoreSSLCheck,omitempty"`
	ServerCert     string                 `json:"serverCert,omitempty"`
	CLI            ScannerIntegrationCLI  `json:"cli,omitempty"`
}

type ScannerIntegrationTestConnection struct {
	Hostname       *string `json:"hostname,omitempty"`
	Token          *string `json:"token,omitempty"`
	IgnoreSSLCheck *bool   `json:"ignoreSSLCheck,omitempty"`
	ServerCert     *string `json:"serverCert,omitempty"`
	RaiseException *bool   `json:"raiseException,omitempty"`
}

type ScannerIntegrationType string

const (
	ScannerIntegrationTypePtai        ScannerIntegrationType = "ptai"
	ScannerIntegrationTypeCodeScoring ScannerIntegrationType = "code_scoring"
	ScannerIntegrationTypeAppScreener ScannerIntegrationType = "app_screener"
	ScannerIntegrationTypeSastav      ScannerIntegrationType = "sastav"
	ScannerIntegrationTypeAppsecTrack ScannerIntegrationType = "appsec_track"
)

type Scanner string

const (
	ScannerSEMGREP                  Scanner = "SEMGREP"
	ScannerTRIVYFS                  Scanner = "TRIVY_FS"
	ScannerTRIVYIMAGE               Scanner = "TRIVY_IMAGE"
	ScannerTRUFFLEHOG               Scanner = "TRUFFLEHOG"
	ScannerGITLEAKS                 Scanner = "GITLEAKS"
	ScannerPTAI                     Scanner = "PT_AI"
	ScannerCODESCORINGREPOSITORY    Scanner = "CODE_SCORING_REPOSITORY"
	ScannerCODESCORINGARTIFACT      Scanner = "CODE_SCORING_ARTIFACT"
	ScannerAPPSCREENERSAST          Scanner = "APP_SCREENER_SAST"
	ScannerAPPSCREENERSCAREPOSITORY Scanner = "APP_SCREENER_SCA_REPOSITORY"
	ScannerSASTAV                   Scanner = "SASTAV"
	ScannerZAP                      Scanner = "ZAP"
	ScannerAPPSECTRACKREPOSITORY    Scanner = "APPSEC_TRACK_REPOSITORY"
	ScannerAPPSECTRACKARTIFACT      Scanner = "APPSEC_TRACK_ARTIFACT"
)

type ScannerType string

const (
	ScannerTypeSAST ScannerType = "SAST"
	ScannerTypeDAST ScannerType = "DAST"
	ScannerTypeSCA  ScannerType = "SCA"
)

type ScannerWithConfigFiles string

const (
	ScannerWithConfigFilesZAP ScannerWithConfigFiles = "ZAP"
)

type ScannerWithCustom any
