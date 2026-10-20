// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type UserGroupCounters struct {
	TotalUsers     int `json:"totalUsers,omitempty"`
	TotalInstances int `json:"totalInstances,omitempty"`
}

type UserGroupCreateEdit struct {
	Description string      `json:"description,omitempty"`
	Name        string      `json:"name,omitempty"`
	RoleID      uuid.UUID   `json:"roleID,omitempty"`
	IsDefault   bool        `json:"isDefault,omitempty"`
	UserIDs     []uuid.UUID `json:"userIDs,omitempty"`
}

type UserGroupMoveUsers struct {
	FromGroupID  *uuid.UUID  `json:"fromGroupID,omitempty"`
	ToGroupID    *uuid.UUID  `json:"toGroupID,omitempty"`
	UserIDs      []uuid.UUID `json:"userIDs,omitempty"`
	MoveAllUsers *bool       `json:"moveAllUsers,omitempty"`
}

type UserGroupPermission struct {
	IsMapped      bool       `json:"isMapped,omitempty"`
	Level         string     `json:"level,omitempty"`
	GroupInstance *UserGroup `json:"groupInstance,omitempty"`
}

type UserGroupRelation struct {
	ID          uuid.UUID             `json:"id,omitempty"`
	Type        UserGroupRelationType `json:"type,omitempty"`
	Name        string                `json:"name,omitempty"`
	Slug        string                `json:"slug,omitempty"`
	Description string                `json:"description,omitempty"`
	Level       string                `json:"level,omitempty"`
	IsMapped    bool                  `json:"isMapped,omitempty"`
}

type UserGroupRelationType string

const (
	UserGroupRelationTypeREPOSITORY UserGroupRelationType = "REPOSITORY"
	UserGroupRelationTypePRODUCT    UserGroupRelationType = "PRODUCT"
)

type UserGroup struct {
	ID          uuid.UUID         `json:"id,omitempty"`
	Slug        string            `json:"slug,omitempty"`
	Description string            `json:"description,omitempty"`
	Name        string            `json:"name,omitempty"`
	Created     string            `json:"created,omitempty"`
	Updated     string            `json:"updated,omitempty"`
	IsDefault   bool              `json:"isDefault,omitempty"`
	Role        *UserRoleShort    `json:"role,omitempty"`
	Counters    UserGroupCounters `json:"counters,omitempty"`
}

type UserGroupsSetRole struct {
	UserGroupIDs []uuid.UUID `json:"userGroupIDs,omitempty"`
	RoleID       *uuid.UUID  `json:"roleID,omitempty"`
}

type UserRoleCounters struct {
	RolePermissions  int `json:"rolePermissions,omitempty"`
	TotalPermissions int `json:"totalPermissions,omitempty"`
	TotalUsers       int `json:"totalUsers,omitempty"`
	TotalUserGroups  int `json:"totalUserGroups,omitempty"`
}

type UserRoleEdit struct {
	Name                 string                `json:"name,omitempty"`
	Description          string                `json:"description,omitempty"`
	Permissions          []UserPermissionValue `json:"permissions,omitempty"`
	AllowedIssueStatuses []string              `json:"allowedIssueStatuses,omitempty"`
}

type UserRolePermissionInfo struct {
	Name         string                `json:"name,omitempty"`
	Slug         UserPermissionValue   `json:"slug,omitempty"`
	Dependencies []UserPermissionValue `json:"dependencies,omitempty"`
}

type UserRolePermissionsCategory struct {
	Slug        PermissionCategorySlug   `json:"slug,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Permissions []UserRolePermissionInfo `json:"permissions,omitempty"`
}

type UserRolePermissionsMeta struct {
	TotalCount int                           `json:"totalCount,omitempty"`
	Categories []UserRolePermissionsCategory `json:"categories,omitempty"`
}

type UserRole struct {
	UserRoleShort
	Permissions []UserPermissionValue `json:"permissions,omitempty"`
	Counters    *UserRoleCounters     `json:"counters,omitempty"`
}

type UserRoleShort struct {
	ID                   uuid.UUID             `json:"id,omitempty"`
	Slug                 UserRoleSlug          `json:"slug,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Description          string                `json:"description,omitempty"`
	System               bool                  `json:"system,omitempty"`
	Permissions          []UserPermissionValue `json:"permissions,omitempty"`
	IsDefault            bool                  `json:"isDefault,omitempty"`
	AllowedIssueStatuses []string              `json:"allowedIssueStatuses,omitempty"`
}

type UserRoleSlug string

type UserRoleValidationError struct {
	Title    string                           `json:"title,omitempty"`
	Detail   string                           `json:"detail,omitempty"`
	Status   int                              `json:"status,omitempty"`
	Type     string                           `json:"type,omitempty"`
	Problems *UserRoleValidationErrorProblems `json:"problems,omitempty"`
}

type UserRoleValidationErrorProblems struct {
	MissingPermissions      []UserPermissionValue `json:"missingPermissions,omitempty"`
	UnexpectedPermissions   []UserPermissionValue `json:"unexpectedPermissions,omitempty"`
	UnexpectedIssueStatuses []string              `json:"unexpectedIssueStatuses,omitempty"`
}
