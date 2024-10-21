package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid; default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	AccountType  AccountType `gorm:"size:32; check:account_type <> ''"`
	RoleMappings []UserApplicationRole
}

type AccountType string

const (
	AccountTypeLocal  AccountType = "LOCAL"
	AccountTypeGoogle AccountType = "GOOGLE"
)

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type UserApplicationRole struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time

	Role Role `gorm:"size:32; check:role <> ''"`

	UserID        uuid.UUID `gorm:"type:uuid; index:idx_user_id_application_id,unique,priority:1"`
	ApplicationID int       `gorm:"index:idx_user_id_application_id,unique,priority:2"`
}

type GoogleAccount struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	GoogleID   string  `gorm:"size:256; check:google_id <> ''"`
	Email      *string `gorm:"size:512"`
	FullName   *string `gorm:"size:512"`
	FamilyName *string `gorm:"size:512"`
	GivenName  *string `gorm:"size:512"`
	PictureURL *string `gorm:"size:2048"`

	UserID uuid.UUID `gorm:"type:uuid; uniqueIndex"`
}
