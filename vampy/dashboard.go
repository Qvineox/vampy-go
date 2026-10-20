// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import (
	"encoding/json"
	"slices"
	"strings"
)

type ByScannerType struct {
	ByScannerType []NameBasedCardValue `json:"by_scanner_type,omitempty"`
}

type BySeverity struct {
	BySeverity []NameBasedCardValue `json:"by_severity,omitempty"`
}

type DashboardCardAggregationPeriod string

const (
	DashboardCardAggregationPeriodWEEK        DashboardCardAggregationPeriod = "WEEK"
	DashboardCardAggregationPeriodMONTH       DashboardCardAggregationPeriod = "MONTH"
	DashboardCardAggregationPeriodTWOMONTHS   DashboardCardAggregationPeriod = "TWO_MONTHS"
	DashboardCardAggregationPeriodTHREEMONTHS DashboardCardAggregationPeriod = "THREE_MONTHS"
	DashboardCardAggregationPeriodYEAR        DashboardCardAggregationPeriod = "YEAR"
)

type DashboardCardEntityRelationType string

const (
	DashboardCardEntityRelationTypeREPOSITORY DashboardCardEntityRelationType = "REPOSITORY"
	DashboardCardEntityRelationTypeVCSUSER    DashboardCardEntityRelationType = "VCS_USER"
)

type DashboardCardGraph struct {
	Name   string               `json:"name,omitempty"`
	Values []TimeBasedCardValue `json:"values,omitempty"`
}

type DashboardCardMeta struct {
	ScanIssueField           string `json:"scanIssueField,omitempty"`
	GraphDisplayMode         string `json:"graphDisplayMode,omitempty"`
	ScanIssueOldBoundaryDays int    `json:"scanIssueOldBoundaryDays,omitempty"`
}

type DashboardCard struct {
	ID              UUID               `json:"id,omitempty"`
	Slug            string             `json:"slug,omitempty"`
	Name            string             `json:"name,omitempty"`
	NameTranslateID string             `json:"nameTranslateID,omitempty"`
	CardType        DashboardCardType  `json:"cardType,omitempty"`
	LazyAggregate   bool               `json:"lazyAggregate,omitempty"`
	Meta            *DashboardCardMeta `json:"meta,omitempty"`
	AggregatedData  any                `json:"aggregatedData,omitempty"`
}

type DashboardCardType string

const (
	DashboardCardTypeBYCATEGORY DashboardCardType = "BY_CATEGORY"
	DashboardCardTypeENTITIES   DashboardCardType = "ENTITIES"
	DashboardCardTypeTABLE      DashboardCardType = "TABLE"
	DashboardCardTypeGRAPH      DashboardCardType = "GRAPH"
	DashboardCardTypeTOPN       DashboardCardType = "TOP_N"
)

type Dashboard struct {
	ID    UUID            `json:"id,omitempty"`
	Slug  string          `json:"slug,omitempty"`
	Name  string          `json:"name,omitempty"`
	Cards []DashboardCard `json:"cards,omitempty"`
}

func (d *Dashboard) GetOpenedIssues() (info, low, medium, high, critical, total int64) {
	index := slices.IndexFunc(d.Cards, func(card DashboardCard) bool {
		return card.Slug == "opened_issues_card"
	})

	if index == -1 {
		return
	}

	// todo: switch to Unmarshal in Dashboard
	marshal, err := json.Marshal(d.Cards[index].AggregatedData)
	if err != nil {
		return
	}

	var card OpenedIssuesCardData
	err = json.Unmarshal(marshal, &card)
	if err != nil {
		return
	}

	total = int64(card.SimpleCardValue.Value)
	for _, v := range card.BySeverity.BySeverity {
		switch strings.ToLower(v.Name) {
		case "info":
			info = int64(v.Value)
		case "low":
			low = int64(v.Value)
		case "medium":
			medium = int64(v.Value)
		case "high":
			high = int64(v.Value)
		case "critical":
			critical = int64(v.Value)
		}
	}

	return
}

type Engagement struct {
	ID                   UUID   `json:"id,omitempty"`
	ExternalScanPipeline string `json:"externalScanPipeline,omitempty"`
}

type Entity struct {
	RelationID   UUID                            `json:"relationID,omitempty"`
	RelationType DashboardCardEntityRelationType `json:"relationType,omitempty"`
}

type IssuesByScannerTypeCardData struct {
	SimpleCardValue
	ByScannerType
}

type IssuesByStatusCardData struct {
	SimpleCardValue
	ByStatus *NameBasedBySeverity `json:"by_status,omitempty"`
}

type NameBasedBySeverity struct {
	NameBasedCardValue
	BySeverity
}

type NameBasedCardValue struct {
	Value        float64 `json:"value,omitempty"`
	DisplayValue float64 `json:"displayValue,omitempty"`
	Name         string  `json:"name,omitempty"`
}

type NameWithEntityBasedCardValue struct {
	NameBasedCardValue
	Entity *Entity `json:"entity,omitempty"`
}

type OldIssuesCardData struct {
	SimpleCardValue
	BySeverity
}

type OpenedIssuesCardData struct {
	SimpleCardValue
	BySeverity
}

type SimpleCardValue struct {
	Value float64 `json:"value,omitempty"`
}

type TimeBasedCardValue struct {
	Value        float64 `json:"value,omitempty"`
	DisplayValue float64 `json:"displayValue,omitempty"`
	Time         string  `json:"time,omitempty"`
}

type TopScanIssuesByField struct {
	NameBasedCardValue
}

type TopVCSUsersByIssuesCard struct {
	NameWithEntityBasedCardValue
	Email string `json:"email,omitempty"`
}
