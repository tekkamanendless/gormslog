package gormslog

import (
	"context"
	"database/sql/driver"
	"errors"
	"strings"
	"syscall"

	"gorm.io/gorm"
)

// AlwaysLogError always returns true if the error is not nil.
func AlwaysLogError(err error) bool {
	return err != nil
}

// NeverLogError never returns true.
func NeverLogError(_ error) bool {
	return false
}

// DefaultShouldLogError returns true if we should log the error and false otherwise.
func DefaultShouldLogError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return false
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	if errors.Is(err, syscall.ECONNREFUSED) {
		return false
	}
	if errors.Is(err, driver.ErrBadConn) {
		return false
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	if errors.Is(err, context.Canceled) {
		return false
	}
	errorString := strings.ToLower(err.Error())
	if strings.Contains(errorString, "deadlock") {
		return false
	}
	if strings.Contains(errorString, "try restarting transaction") {
		return false
	}
	if strings.Contains(errorString, "invalid connection") {
		return false
	}
	if strings.Contains(errorString, "connection timed out") {
		return false
	}
	if strings.Contains(errorString, "context canceled") { // This should be covered by "context.Canceled", but other errors might also return this string.
		return false
	}
	return true
}
