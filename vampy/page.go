// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

// Page is a generic paginated response envelope.
type Page[T any] struct {
	Items      []T  `json:"items,omitempty"`
	HasNext    bool `json:"hasNext,omitempty"`
	Offset     int  `json:"offset,omitempty"`
	TotalCount int  `json:"totalCount,omitempty"`
}

// LoginResult is the response of POST /session/. It is either a user or an MFA challenge.
type LoginResult = any
