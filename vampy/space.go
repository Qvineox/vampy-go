// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import (
	"time"
	"uuid"
)

type SpaceIntegration struct {
	ID             uuid.UUID              `json:"id,omitempty"`
	Name           string                 `json:"name"`
	Hostname       string                 `json:"hostname"`
	Token          string                 `json:"token,omitempty"`
	Webhook        string                 `json:"webhook,omitempty"`
	Service        IntegrationServiceType `json:"service"`
	State          IntegrationState       `json:"state,omitempty"`
	ServerID       uuid.UUID              `json:"serverID,omitempty"`
	Created        time.Time              `json:"created,omitempty"`
	Roles          []IntegrationRole      `json:"roles,omitempty"`
	IgnoreSSLCheck *bool                  `json:"ignoreSSLCheck,omitempty"`
}

type SpaceIntegrationWithMeta struct {
	SpaceIntegration
	Meta *SpaceIntegrationWithMetaMeta `json:"meta,omitempty"`
}

type SpaceIntegrationWithMetaMeta struct {
	AI *SpaceIntegrationWithMetaMetaAI `json:"ai,omitempty"`
}

type SpaceIntegrationWithMetaMetaAI struct {
	Vendor                string `json:"vendor,omitempty"`
	IsDefault             bool   `json:"isDefault,omitempty"`
	DefaultModel          string `json:"defaultModel,omitempty"`
	DefaultModelAvailable bool   `json:"defaultModelAvailable,omitempty"`
}

type Space struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name"`
	Created time.Time `json:"created,omitempty"`
}
