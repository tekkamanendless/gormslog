package gormslog

import "gorm.io/gorm/logger"

// Config is the config.
type Config struct {
	Level          logger.LogLevel  // This is the log level.
	ShouldLogError func(error) bool // This should return true if the error should be logged.
}

// Option operates on a Config.
type Option func(*Config)

// WithLogLevel returns an option that sets the log level.
func WithLogLevel(level logger.LogLevel) func(*Config) {
	return func(config *Config) {
		config.Level = level
	}
}

// WithShouldLogError returns an option that sets how errors are logged.
func WithShouldLogError(shouldLogError func(err error) bool) func(*Config) {
	return func(config *Config) {
		config.ShouldLogError = shouldLogError
	}
}
