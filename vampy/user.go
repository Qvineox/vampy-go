// Package vampy contains the entity types for the Vampy ASPM API.
package vampy

import "uuid"

type DetailPermissionUser struct {
	UserID   uuid.UUID `json:"userID,omitempty"`
	Level    string    `json:"level,omitempty"`
	IsMapped bool      `json:"isMapped,omitempty"`
}

type LoginRequest struct {
	UserLogin    string  `json:"userLogin"`
	UserPassword string  `json:"userPassword"`
	MFAToken     *string `json:"mfaToken,omitempty"`
}

type MFAConfirmation struct {
	MFAToken *string `json:"mfaToken,omitempty"`
}

type PermissionCategorySlug string

const (
	PermissionCategorySlugREPOSITORY      PermissionCategorySlug = "REPOSITORY"
	PermissionCategorySlugPRODUCT         PermissionCategorySlug = "PRODUCT"
	PermissionCategorySlugCOMPONENT       PermissionCategorySlug = "COMPONENT"
	PermissionCategorySlugASSET           PermissionCategorySlug = "ASSET"
	PermissionCategorySlugSCANISSUE       PermissionCategorySlug = "SCAN_ISSUE"
	PermissionCategorySlugBOTUSER         PermissionCategorySlug = "BOT_USER"
	PermissionCategorySlugQUALITYGATE     PermissionCategorySlug = "QUALITY_GATE"
	PermissionCategorySlugORCHESTRATION   PermissionCategorySlug = "ORCHESTRATION"
	PermissionCategorySlugCOMMENT         PermissionCategorySlug = "COMMENT"
	PermissionCategorySlugARTIFACT        PermissionCategorySlug = "ARTIFACT"
	PermissionCategorySlugRULES           PermissionCategorySlug = "RULES"
	PermissionCategorySlugINTEGRATION     PermissionCategorySlug = "INTEGRATION"
	PermissionCategorySlugBACKGROUNDTASKS PermissionCategorySlug = "BACKGROUND_TASKS"
	PermissionCategorySlugLDAP            PermissionCategorySlug = "LDAP"
	PermissionCategorySlugSMTP            PermissionCategorySlug = "SMTP"
	PermissionCategorySlugISSUESTATUS     PermissionCategorySlug = "ISSUE_STATUS"
	PermissionCategorySlugSYSTEMSETTINGS  PermissionCategorySlug = "SYSTEM_SETTINGS"
	PermissionCategorySlugSLA             PermissionCategorySlug = "SLA"
	PermissionCategorySlugLICENSE         PermissionCategorySlug = "LICENSE"
	PermissionCategorySlugUSERS           PermissionCategorySlug = "USERS"
	PermissionCategorySlugAI              PermissionCategorySlug = "AI"
)

type PermissionSetUserGroups struct {
	Permissions []PermissionUserGroup `json:"permissions,omitempty"`
}

type PermissionSetUsers struct {
	Permissions []PermissionUser `json:"permissions,omitempty"`
}

type PermissionUnsetUserGroups struct {
	UserGroups []uuid.UUID `json:"user_groups,omitempty"`
}

type PermissionUnsetUsers struct {
	Users []uuid.UUID `json:"users,omitempty"`
}

type PermissionUserGroup struct {
	GroupID uuid.UUID `json:"groupID,omitempty"`
	Level   string    `json:"level,omitempty"`
}

type PermissionUser struct {
	UserID uuid.UUID `json:"userID,omitempty"`
	Level  string    `json:"level,omitempty"`
}

type UserAdmin struct {
	ID                     uuid.UUID             `json:"id,omitempty"`
	UserLogin              string                `json:"userLogin,omitempty"`
	UserEmail              *string               `json:"userEmail,omitempty"`
	FirstName              string                `json:"firstName,omitempty"`
	LastName               string                `json:"lastName,omitempty"`
	Registered             string                `json:"registered,omitempty"`
	State                  string                `json:"state,omitempty"`
	Permissions            []UserPermissionValue `json:"permissions,omitempty"`
	PasswordChangeRequired bool                  `json:"passwordChangeRequired,omitempty"`
	UserPassword           string                `json:"userPassword,omitempty"`
	UserType               string                `json:"userType,omitempty"`
	MFAConfirmed           *bool                 `json:"mfaConfirmed,omitempty"`
	IsBot                  bool                  `json:"isBot,omitempty"`
	IsAdmin                bool                  `json:"isAdmin,omitempty"`
	IsStaff                bool                  `json:"isStaff,omitempty"`
	Role                   UserRoleSlug          `json:"role,omitempty"`
	GroupID                uuid.UUID             `json:"groupID,omitempty"`
	GroupInstance          UserGroup             `json:"groupInstance,omitempty"`
	RoleInstance           *UserRoleShort        `json:"roleInstance,omitempty"`
}

type UserMFA struct {
	Base32     string `json:"base32,omitempty"`
	OtpauthURL string `json:"otpauthUrl,omitempty"`
}

type UserPermissionValue string

const (
	UserPermissionValueALL                            UserPermissionValue = "ALL"
	UserPermissionValueREPOSITORYREAD                 UserPermissionValue = "REPOSITORY_READ"
	UserPermissionValueREPOSITORYCREATE               UserPermissionValue = "REPOSITORY_CREATE"
	UserPermissionValueREPOSITORYEDIT                 UserPermissionValue = "REPOSITORY_EDIT"
	UserPermissionValueREPOSITORYSETTINGSCHANGE       UserPermissionValue = "REPOSITORY_SETTINGS_CHANGE"
	UserPermissionValuePRODUCTREAD                    UserPermissionValue = "PRODUCT_READ"
	UserPermissionValuePRODUCTCREATE                  UserPermissionValue = "PRODUCT_CREATE"
	UserPermissionValuePRODUCTEDIT                    UserPermissionValue = "PRODUCT_EDIT"
	UserPermissionValuePRODUCTSETTINGSCHANGE          UserPermissionValue = "PRODUCT_SETTINGS_CHANGE"
	UserPermissionValueCOMPONENTREAD                  UserPermissionValue = "COMPONENT_READ"
	UserPermissionValueCOMPONENTEDIT                  UserPermissionValue = "COMPONENT_EDIT"
	UserPermissionValueASSETREAD                      UserPermissionValue = "ASSET_READ"
	UserPermissionValueASSETCREATE                    UserPermissionValue = "ASSET_CREATE"
	UserPermissionValueASSETEDIT                      UserPermissionValue = "ASSET_EDIT"
	UserPermissionValueSCANISSUEREAD                  UserPermissionValue = "SCAN_ISSUE_READ"
	UserPermissionValueSCANISSUECREATE                UserPermissionValue = "SCAN_ISSUE_CREATE"
	UserPermissionValueSCANISSUEDELETE                UserPermissionValue = "SCAN_ISSUE_DELETE"
	UserPermissionValueSCANISSUEEDIT                  UserPermissionValue = "SCAN_ISSUE_EDIT"
	UserPermissionValueSCANISSUESENDTOINTEGRATIONS    UserPermissionValue = "SCAN_ISSUE_SEND_TO_INTEGRATIONS"
	UserPermissionValueSCANISSUESTATUSCHANGE          UserPermissionValue = "SCAN_ISSUE_STATUS_CHANGE"
	UserPermissionValueSCANISSUESEVERITYCHANGE        UserPermissionValue = "SCAN_ISSUE_SEVERITY_CHANGE"
	UserPermissionValueBOTUSERREAD                    UserPermissionValue = "BOT_USER_READ"
	UserPermissionValueBOTUSERCREATE                  UserPermissionValue = "BOT_USER_CREATE"
	UserPermissionValueBOTUSEREDIT                    UserPermissionValue = "BOT_USER_EDIT"
	UserPermissionValueSCANPIPELINEREAD               UserPermissionValue = "SCAN_PIPELINE_READ"
	UserPermissionValueSCANPIPELINEWRITE              UserPermissionValue = "SCAN_PIPELINE_WRITE"
	UserPermissionValueQUALITYGATEREAD                UserPermissionValue = "QUALITY_GATE_READ"
	UserPermissionValueQUALITYGATECREATE              UserPermissionValue = "QUALITY_GATE_CREATE"
	UserPermissionValueQUALITYGATEEDIT                UserPermissionValue = "QUALITY_GATE_EDIT"
	UserPermissionValueRULESREAD                      UserPermissionValue = "RULES_READ"
	UserPermissionValueRULESWRITE                     UserPermissionValue = "RULES_WRITE"
	UserPermissionValueADMINSPACEWRITE                UserPermissionValue = "ADMIN_SPACE_WRITE"
	UserPermissionValueADMININTEGRATIONREAD           UserPermissionValue = "ADMIN_INTEGRATION_READ"
	UserPermissionValueADMININTEGRATIONWRITE          UserPermissionValue = "ADMIN_INTEGRATION_WRITE"
	UserPermissionValueADMININTEGRATIONIMPORTFINDINGS UserPermissionValue = "ADMIN_INTEGRATION_IMPORT_FINDINGS"
	UserPermissionValueADMINBACKGROUNDTASKSREAD       UserPermissionValue = "ADMIN_BACKGROUND_TASKS_READ"
	UserPermissionValueADMINLDAPREAD                  UserPermissionValue = "ADMIN_LDAP_READ"
	UserPermissionValueADMINLDAPWRITE                 UserPermissionValue = "ADMIN_LDAP_WRITE"
	UserPermissionValueADMINOIDCREAD                  UserPermissionValue = "ADMIN_OIDC_READ"
	UserPermissionValueADMINOIDCWRITE                 UserPermissionValue = "ADMIN_OIDC_WRITE"
	UserPermissionValueADMINSMTPREAD                  UserPermissionValue = "ADMIN_SMTP_READ"
	UserPermissionValueADMINSMTPWRITE                 UserPermissionValue = "ADMIN_SMTP_WRITE"
	UserPermissionValueADMINISSUESTATUSWRITE          UserPermissionValue = "ADMIN_ISSUE_STATUS_WRITE"
	UserPermissionValueADMINSYSTEMSETTINGSWRITE       UserPermissionValue = "ADMIN_SYSTEM_SETTINGS_WRITE"
	UserPermissionValueADMINSLAREAD                   UserPermissionValue = "ADMIN_SLA_READ"
	UserPermissionValueADMINSLAWRITE                  UserPermissionValue = "ADMIN_SLA_WRITE"
	UserPermissionValueADMINLICENSEREAD               UserPermissionValue = "ADMIN_LICENSE_READ"
	UserPermissionValueADMINLICENSEWRITE              UserPermissionValue = "ADMIN_LICENSE_WRITE"
	UserPermissionValueADMINUSERREAD                  UserPermissionValue = "ADMIN_USER_READ"
	UserPermissionValueADMINUSERWRITE                 UserPermissionValue = "ADMIN_USER_WRITE"
	UserPermissionValueADMINUSERROLEREAD              UserPermissionValue = "ADMIN_USER_ROLE_READ"
	UserPermissionValueADMINUSERROLEWRITE             UserPermissionValue = "ADMIN_USER_ROLE_WRITE"
	UserPermissionValueADMINHEALTHCHECK               UserPermissionValue = "ADMIN_HEALTH_CHECK"
	UserPermissionValueADMINFPCLASSIFIERREAD          UserPermissionValue = "ADMIN_FP_CLASSIFIER_READ"
	UserPermissionValueADMINFPCLASSIFIERWRITE         UserPermissionValue = "ADMIN_FP_CLASSIFIER_WRITE"
	UserPermissionValueORCHESTRATIONRUN               UserPermissionValue = "ORCHESTRATION_RUN"
	UserPermissionValueORCHESTRATIONPROFILEREAD       UserPermissionValue = "ORCHESTRATION_PROFILE_READ"
	UserPermissionValueORCHESTRATIONPROFILEEDIT       UserPermissionValue = "ORCHESTRATION_PROFILE_EDIT"
	UserPermissionValueORCHESTRATIONPROFILECREATE     UserPermissionValue = "ORCHESTRATION_PROFILE_CREATE"
	UserPermissionValueCOMMENTCREATE                  UserPermissionValue = "COMMENT_CREATE"
	UserPermissionValueCOMMENTREAD                    UserPermissionValue = "COMMENT_READ"
	UserPermissionValueCOMMENTEDIT                    UserPermissionValue = "COMMENT_EDIT"
	UserPermissionValueCOMMENTEDITOWN                 UserPermissionValue = "COMMENT_EDIT_OWN"
	UserPermissionValueARTIFACTREAD                   UserPermissionValue = "ARTIFACT_READ"
	UserPermissionValueARTIFACTEDIT                   UserPermissionValue = "ARTIFACT_EDIT"
	UserPermissionValueAIMODELSACCESS                 UserPermissionValue = "AI_MODELS_ACCESS"
	UserPermissionValueAISCANISSUEAPPLY               UserPermissionValue = "AI_SCAN_ISSUE_APPLY"
	UserPermissionValueAISCANISSUECHANGEANSWER        UserPermissionValue = "AI_SCAN_ISSUE_CHANGE_ANSWER"
	UserPermissionValueADMINUSERGROUPREAD             UserPermissionValue = "ADMIN_USER_GROUP_READ"
	UserPermissionValueADMINUSERGROUPWRITE            UserPermissionValue = "ADMIN_USER_GROUP_WRITE"
)

type UserRegistration struct {
	UserLogin    string       `json:"userLogin"`
	UserPassword string       `json:"userPassword"`
	UserEmail    *string      `json:"userEmail,omitempty"`
	FirstName    string       `json:"firstName,omitempty"`
	LastName     string       `json:"lastName,omitempty"`
	Role         UserRoleSlug `json:"role,omitempty"`
	GroupID      *uuid.UUID   `json:"groupID,omitempty"`
}

type UserShort struct {
	ID                     uuid.UUID             `json:"id,omitempty"`
	UserLogin              string                `json:"userLogin,omitempty"`
	UserEmail              *string               `json:"userEmail,omitempty"`
	FirstName              string                `json:"firstName,omitempty"`
	LastName               string                `json:"lastName,omitempty"`
	IsBot                  bool                  `json:"isBot,omitempty"`
	IsAdmin                bool                  `json:"isAdmin,omitempty"`
	IsStaff                bool                  `json:"isStaff,omitempty"`
	PasswordChangeRequired bool                  `json:"passwordChangeRequired,omitempty"`
	Role                   UserRoleSlug          `json:"role,omitempty"`
	RoleInstance           *UserRoleShort        `json:"roleInstance,omitempty"`
	Permissions            []UserPermissionValue `json:"permissions,omitempty"`
}

type UserSystemRoleValue string

const (
	UserSystemRoleValueADMIN       UserSystemRoleValue = "ADMIN"
	UserSystemRoleValueSYSENGINEER UserSystemRoleValue = "SYS_ENGINEER"
	UserSystemRoleValueSECENGINEER UserSystemRoleValue = "SEC_ENGINEER"
	UserSystemRoleValueOWNER       UserSystemRoleValue = "OWNER"
	UserSystemRoleValueDEVELOPER   UserSystemRoleValue = "DEVELOPER"
)

type UserTiny struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserLogin string    `json:"userLogin,omitempty"`
	UserEmail *string   `json:"userEmail,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
}
