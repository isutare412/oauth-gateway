package model

import (
	"time"
)

type Application struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time

	Name              string `gorm:"size:128; check:name <> ''; uniqueIndex"`
	AuthorizedOrigins []AuthorizedOrigin
	AuthorizedURIs    []AuthorizedRedirectionURI
}

type AuthorizedOrigin struct {
	ID     int
	Origin string `gorm:"size:1024; check:origin <> ''; index:idx_application_id_origin,unique,priority:2"`

	ApplicationID int `gorm:"index:idx_application_id_origin,unique,priority:1"`
}

type AuthorizedRedirectionURI struct {
	ID  int
	URI string `gorm:"size:1024; check:uri <> ''; index:idx_application_id_uri,unique,priority:2"`

	ApplicationID int `gorm:"index:idx_application_id_uri,unique,priority:1"`
}
