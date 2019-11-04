package utils

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

/*
 * User related types are defined here.
 * Also validation added to the types
 */

// JwtClaim : payload for jwt
type JwtClaim struct {
	UserID         uint
	UserName       string
	OrgCode        string
	Permissions    pq.StringArray
	RoleCode       string
	AppPermissions postgres.Jsonb
	jwt.StandardClaims
}

// UserData : create user struct
type UserData struct {
	Username       string         `json:"username" validate:"required"`
	OrgCode        string         `json:"org_code" validate:"required"`
	Name           string         `json:"name" validate:"required"`
	Email          string         `json:"email" validate:"required,email"`
	Status         bool           `json:"status"`
	RoleID         int            `json:"role_id" validate:"required"`
	Msisdn         string         `json:"msisdn" validate:"len=13"`
	CountryCode    string         `json:"country_code" validate:"required,max=5"`
	AdditionalInfo postgres.Jsonb `json:"additional_info" validate:"required"`
	CreatedBy      int            `json:"created_by"`
	UpdatedBy      int            `json:"updated_by"`
}

// UserUpdate : update user struct
type UserUpdate struct {
	ID             uint           `json:"id"`
	OrgCode        string         `json:"org_code"`
	Name           string         `json:"name"`
	Address        string         `json:"address"`
	RoleID         int            `json:"role_id"`
	CountryCode    string         `json:"country_code"`
	AdditionalInfo postgres.Jsonb `json:"additional_info"`
	UpdatedBy      int            `json:"updated_by"`
}

// UserPermissions : all user permissions list
type UserPermissions struct {
	Post    string
	Put     string
	Get     string
	GetList string
	Delete  string
}

// ResponseWriterWithLog ...
type ResponseWriterWithLog struct {
	http.ResponseWriter
	Status int
	Length int
	Body   []byte
}

// OrganizationPermissions : permission for organization apis
type OrganizationPermissions struct {
	Get string
	Put string
}

// RoleUpdate : role update object
type RoleUpdate struct {
	RoleName       string         `json:"role_name"`
	Permissions    []string       `json:"permissions"`
	UpdatedBy      int            `json:"updated_by"`
	AppPermissions postgres.Jsonb `json:"app_permissions"`
}

// RoleData : ..
type RoleData struct {
	RoleName       string         `json:"role_name"`
	Permissions    []string       `json:"permissions"`
	RoleCode       string         `json:"role_code"`
	CreatedBy      int            `json:"created_by"`
	UpdatedBy      int            `json:"updated_by"`
	AppPermissions postgres.Jsonb `json:"app_permissions"`
}

// CategoryData ..
type CategoryData struct {
	Code      string `json:"code" validate:"required,max=10"`
	Name      string `json:"name" validate:"required,max=200"`
	OrgCode   string `json:"org_code" validate:"required,max=10"`
	ParentID  int    `json:"parent_id" validate:"numeric"`
	IsActive  bool   `json:"is_active"`
	Image     string `json:"image"`
	CreatedBy int    `json:"created_by" validate:"numeric"`
}
