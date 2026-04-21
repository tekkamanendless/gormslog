package gormslog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestLogger(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		l := New()
		assert.Equal(t, logger.Warn, l.level)
		assert.Equal(t, fmt.Sprintf("%p", DefaultShouldLogError), fmt.Sprintf("%p", l.shouldLogError))

		t.Run("LogMode", func(t *testing.T) {
			l2logger := l.LogMode(logger.Silent)
			l2 := l2logger.(*Logger)
			assert.Equal(t, logger.Silent, l2.level)
			assert.Equal(t, fmt.Sprintf("%p", DefaultShouldLogError), fmt.Sprintf("%p", l.shouldLogError))
		})
	})
	t.Run("WithLogLevel", func(t *testing.T) {
		l := New(WithLogLevel(logger.Error))
		assert.Equal(t, logger.Error, l.level)
		assert.Equal(t, fmt.Sprintf("%p", DefaultShouldLogError), fmt.Sprintf("%p", l.shouldLogError))

		t.Run("LogMode", func(t *testing.T) {
			l2logger := l.LogMode(logger.Silent)
			l2 := l2logger.(*Logger)
			assert.Equal(t, logger.Silent, l2.level)
			assert.Equal(t, fmt.Sprintf("%p", DefaultShouldLogError), fmt.Sprintf("%p", l.shouldLogError))
		})
	})
	t.Run("WithShouldLogError", func(t *testing.T) {
		l := New(WithShouldLogError(NeverLogError))
		assert.Equal(t, logger.Warn, l.level)
		assert.Equal(t, fmt.Sprintf("%p", NeverLogError), fmt.Sprintf("%p", l.shouldLogError))

		t.Run("LogMode", func(t *testing.T) {
			l2logger := l.LogMode(logger.Silent)
			l2 := l2logger.(*Logger)
			assert.Equal(t, logger.Silent, l2.level)
			assert.Equal(t, fmt.Sprintf("%p", NeverLogError), fmt.Sprintf("%p", l.shouldLogError))
		})
	})
}
