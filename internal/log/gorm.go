package log

import (
	"time"

	sloggorm "github.com/orandin/slog-gorm"
	"gorm.io/gorm/logger"
)

func NewGORMLogger(slowQueryThreshold time.Duration) logger.Interface {
	return sloggorm.New(
		sloggorm.WithSlowThreshold(slowQueryThreshold),
		// sloggorm.SetLogLevel(sloggorm.DefaultLogType, slog.LevelDebug),
		// sloggorm.WithTraceAll(),
	)
}
