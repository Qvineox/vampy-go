// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

type BotUserRegistration struct {
	UserLogin string `json:"userLogin"`
	GroupID   *UUID  `json:"groupID,omitempty"`
}

type BotUser struct {
	ID            string                `json:"id,omitempty"`
	UserLogin     string                `json:"userLogin"`
	Role          UserRoleSlug          `json:"role,omitempty"`
	GroupID       UUID                  `json:"groupID,omitempty"`
	GroupInstance *UserGroup            `json:"groupInstance,omitempty"`
	RoleInstance  *UserRoleShort        `json:"roleInstance,omitempty"`
	Permissions   []UserPermissionValue `json:"permissions,omitempty"`
}

type BotUserToken struct {
	ID    string `json:"id,omitempty"`
	Token string `json:"token,omitempty"`
}
