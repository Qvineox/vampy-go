// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type LicenseInformation struct {
	CommonParameters *LicenseInformationCommonParameters `json:"commonParameters,omitempty"`
}

type LicenseInformationCommonParameters struct {
	User                    *InformationParameter `json:"user,omitempty"`
	BotUser                 *InformationParameter `json:"botUser,omitempty"`
	LDAP                    *InformationParameter `json:"ldap,omitempty"`
	MFA                     *InformationParameter `json:"mfa,omitempty"`
	Repository              *InformationParameter `json:"repository,omitempty"`
	Product                 *InformationParameter `json:"product,omitempty"`
	Jira                    *InformationParameter `json:"jira,omitempty"`
	Ktracker                *InformationParameter `json:"ktracker,omitempty"`
	EvaTeam                 *InformationParameter `json:"evaTeam,omitempty"`
	TaskTrackerIntegration  *InformationParameter `json:"taskTrackerIntegration,omitempty"`
	TaskTrackerIssue        *InformationParameter `json:"taskTrackerIssue,omitempty"`
	Gitlab                  *InformationParameter `json:"gitlab,omitempty"`
	Bitbucket               *InformationParameter `json:"bitbucket,omitempty"`
	Gitea                   *InformationParameter `json:"gitea,omitempty"`
	Gitflic                 *InformationParameter `json:"gitflic,omitempty"`
	Gitflame                *InformationParameter `json:"gitflame,omitempty"`
	VCSIntegration          *InformationParameter `json:"vcsIntegration,omitempty"`
	QualityGate             *InformationParameter `json:"qualityGate,omitempty"`
	Orchestration           *InformationParameter `json:"orchestration,omitempty"`
	Ptai                    *InformationParameter `json:"ptai,omitempty"`
	CodeScoring             *InformationParameter `json:"codeScoring,omitempty"`
	AppScreener             *InformationParameter `json:"appScreener,omitempty"`
	AppsecTrack             *InformationParameter `json:"appsecTrack,omitempty"`
	Sastav                  *InformationParameter `json:"sastav,omitempty"`
	ScannerIntegration      *InformationParameter `json:"scannerIntegration,omitempty"`
	Harbor                  *InformationParameter `json:"harbor,omitempty"`
	Nexus                   *InformationParameter `json:"nexus,omitempty"`
	ArtifactIntegration     *InformationParameter `json:"artifactIntegration,omitempty"`
	OIDC                    *InformationParameter `json:"oidc,omitempty"`
	PDFReport               *InformationParameter `json:"pdfReport,omitempty"`
	ImportIntegration       *InformationParameter `json:"importIntegration,omitempty"`
	DefectDojo              *InformationParameter `json:"defectDojo,omitempty"`
	AIIntegration           *InformationParameter `json:"aiIntegration,omitempty"`
	ScanPipelinesComparison *InformationParameter `json:"scanPipelinesComparison,omitempty"`
	AzureDevOps             *InformationParameter `json:"azureDevOps,omitempty"`
	AzureDevOpsBoards       *InformationParameter `json:"azureDevOpsBoards,omitempty"`
}

type License struct {
	ID           UUID           `json:"id,omitempty"`
	CustomerName string         `json:"customerName,omitempty"`
	StartDate    string         `json:"startDate,omitempty"`
	EndDate      string         `json:"endDate,omitempty"`
	IsNotExpired bool           `json:"isNotExpired,omitempty"`
	IsExpired    bool           `json:"isExpired,omitempty"`
	IsValid      bool           `json:"isValid,omitempty"`
	LicenseType  string         `json:"licenseType,omitempty"`
	ProductName  string         `json:"productName,omitempty"`
	Version      int            `json:"version,omitempty"`
	Limits       *LicenseLimits `json:"limits,omitempty"`
}

type LicenseLimits struct {
	UserCount                   int  `json:"userCount,omitempty"`
	BotUserCount                int  `json:"botUserCount,omitempty"`
	RepositoryCount             int  `json:"repositoryCount,omitempty"`
	ProductCount                int  `json:"productCount,omitempty"`
	TaskTrackerIntegrationCount int  `json:"taskTrackerIntegrationCount,omitempty"`
	TaskTrackerIssueCount       int  `json:"taskTrackerIssueCount,omitempty"`
	ScannerIntegrationCount     int  `json:"scannerIntegrationCount,omitempty"`
	ArtifactIntegrationCount    int  `json:"artifactIntegrationCount,omitempty"`
	ImportIntegrationCount      int  `json:"importIntegrationCount,omitempty"`
	AIIntegrationCount          int  `json:"aiIntegrationCount,omitempty"`
	ScanPipelinesComparison     bool `json:"scanPipelinesComparison,omitempty"`
}

type LicenseBody struct {
	LicenseFile []byte `json:"license_file"`
}
