// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

// Paging holds offset/limit pagination parameters shared by list endpoints.
type Paging struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
