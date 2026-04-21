package Logger

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

// Logger is a gorm logger that does absolutely nothing.
type Logger struct{}

var _ logger.Interface = (*Logger)(nil)

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *Logger) Info(ctx context.Context, s string, args ...interface{}) {
	// Do nothing.
}

func (l *Logger) Warn(ctx context.Context, s string, args ...interface{}) {
	// Do nothing.
}

func (l *Logger) Error(ctx context.Context, s string, args ...interface{}) {
	// Do nothing.
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// Do nothing.
}
