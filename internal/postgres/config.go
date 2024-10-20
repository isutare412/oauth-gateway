package postgres

import "time"

type Config struct {
	Host               string
	Port               int
	Database           string
	User               string
	Password           string
	SlowQueryThreshold time.Duration
}
