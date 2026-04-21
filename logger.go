package gormslog

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

// Logger is a Gorm-compatible logger that uses log/slog.
type Logger struct {
	level          logger.LogLevel  // This is the log level.
	shouldLogError func(error) bool // This should return true if the error should be logged.
}

var _ logger.Interface = (*Logger)(nil)

// New returns a new Logger for use with Gorm.
func New(options ...Option) *Logger {
	config := Config{
		Level:          logger.Warn,
		ShouldLogError: DefaultShouldLogError,
	}
	for _, option := range options {
		option(&config)
	}
	return &Logger{
		level:          config.Level,
		shouldLogError: config.ShouldLogError,
	}
}

// LogMode returns a new logger with the desired log mode.
//
// This is required as part of the logger.Interface interface.
func (l Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := l
	newLogger.level = level
	return &newLogger
}

// determineProgramCounter returns the appropriate program counter for the caller.
// This will be the source location outside of the gorm package.
//
// If, for whatever reason, we couldn't find the program counter, then this will return 0.
func determineProgramCounter() uintptr {
	skips := 2 // 1 for this function; 1 for the previous function in `Logger`.
	programCounters := make([]uintptr, 20)
	n := runtime.Callers(skips, programCounters)
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for {
			frame, more := frames.Next()

			// NOTE: At some point you will need to debug this.
			// NOTE: Uncomment this block to print the various strings that you can use to perform the check.
			/*
				fmt.Printf("ENTRY: File: %s\n", frame.File)
				fmt.Printf("ENTRY: Function: %s\n", frame.Function)
				fmt.Printf("ENTRY: Line: %d\n", frame.Line)
				//*/
			if strings.Contains(frame.Function, "/gormslog.Logger.") {
				// This is one of our own functions, so we can skip it.
			} else if strings.Contains(frame.File, "/gorm@") {
				// This is a Gorm call, so we can skip it.
			} else {
				return frame.PC
			}

			if !more {
				break
			}
		}
	}
	return 0
}

// Info logs an info message.
func (l Logger) Info(ctx context.Context, line string, variables ...any) {
	if l.level >= logger.Info {
		programCounter := determineProgramCounter()
		if slog.Default().Handler().Enabled(ctx, slog.LevelInfo) {
			slog.Default().Handler().Handle(ctx, slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(line, variables...), programCounter))
		}
	}
}

// Warn logs a warning message.
func (l Logger) Warn(ctx context.Context, line string, variables ...any) {
	if l.level >= logger.Warn {
		programCounter := determineProgramCounter()
		if slog.Default().Handler().Enabled(ctx, slog.LevelWarn) {
			slog.Default().Handler().Handle(ctx, slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprintf(line, variables...), programCounter))
		}
	}
}

// Error logs an error message.
func (l Logger) Error(ctx context.Context, line string, variables ...any) {
	if l.level >= logger.Error {
		programCounter := determineProgramCounter()
		if slog.Default().Handler().Enabled(ctx, slog.LevelError) {
			slog.Default().Handler().Handle(ctx, slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(line, variables...), programCounter))
		}
	}
}

// Trace logs a trace message.
func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	rowsAffectedString := "-"
	if rowsAffected >= 0 {
		rowsAffectedString = fmt.Sprintf("%d", rowsAffected)
	}
	if l.shouldLogError(err) {
		programCounter := determineProgramCounter()
		if slog.Default().Handler().Enabled(ctx, slog.LevelError) {
			slog.Default().Handler().Handle(ctx, slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf("[%3.5fs] [rows:%s] [error] %s", time.Since(begin).Seconds(), rowsAffectedString, sql), programCounter))
			slog.Default().Handler().Handle(ctx, slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf("Database error was: [%T] %v", err, err), programCounter))
		}
	} else {
		programCounter := determineProgramCounter()
		if slog.Default().Handler().Enabled(ctx, slog.LevelDebug) {
			slog.Default().Handler().Handle(ctx, slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprintf("[%3.5fs] [rows:%s] [okay] %s", time.Since(begin).Seconds(), rowsAffectedString, sql), programCounter))
		}
	}
}
