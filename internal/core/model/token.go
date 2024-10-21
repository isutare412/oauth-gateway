package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type APIToken struct {
	ID        uuid.UUID `gorm:"type:uuid; default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Subject string `gorm:"size:256; check:subject <> ''"`
}
