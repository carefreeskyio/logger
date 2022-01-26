package logger

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// TraceF logs a message at level Trace on the standard logger.
func TraceF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Tracef(format, args...)
}

// DebugF logs a message at level Debug on the standard logger.
func DebugF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Debugf(format, args...)
}

// PrintF logs a message at level Info on the standard logger.
func PrintF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infof(format, args...)
}

// InfoF logs a message at level Info on the standard logger.
func InfoF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infof(format, args...)
}

// WarnF logs a message at level Warn on the standard logger.
func WarnF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnf(format, args...)
}

// WarningF logs a message at level Warn on the standard logger.
func WarningF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnf(format, args...)
}

// ErrorF logs a message at level Error on the standard logger.
func ErrorF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Errorf(format, args...)
}

// PanicF logs a message at level Panic on the standard logger.
func PanicF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Panicf(format, args...)
}

// FatalF logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalF(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Fatalf(format, args...)
}

// TraceFX logs a message at level Trace on the standard logger.
func TraceFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Tracef(format, args...)
}

// DebugFX logs a message at level Debug on the standard logger.
func DebugFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Debugf(format, args...)
}

// PrintFX logs a message at level Info on the standard logger.
func PrintFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Infof(format, args...)
}

// InfoFX logs a message at level Info on the standard logger.
func InfoFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Infof(format, args...)
}

// WarnFX logs a message at level Warn on the standard logger.
func WarnFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Warnf(format, args...)
}

// WarningFX logs a message at level Warn on the standard logger.
func WarningFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Warnf(format, args...)
}

// ErrorFX logs a message at level Error on the standard logger.
func ErrorFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Errorf(format, args...)
}

// PanicFX logs a message at level Panic on the standard logger.
func PanicFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Panicf(format, args...)
}

// FatalFX logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Fatalf(format, args...)
}
