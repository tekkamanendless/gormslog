package nillogger

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

// NilLogger is a gorm logger that does absolutely nothing.
type NilLogger struct{}

var _ logger.Interface = (*NilLogger)(nil)

func (l *NilLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *NilLogger) Info(ctx context.Context, s string, args ...interface{}) {
	// Do nothing.
}

func (l *NilLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	// Do nothing.
}

func (l *NilLogger) Error(ctx context.Context, s string, args ...interface{}) {
	// Do nothing.
}

func (l *NilLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// Do nothing.
}
