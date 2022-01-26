package logger

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infof(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Fatalf(format, args...)
}

// TracefX logs a message at level Trace on the standard logger.
func TracefX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Tracef(format, args...)
}

// DebugfX logs a message at level Debug on the standard logger.
func DebugfX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Debugf(format, args...)
}

// PrintfX logs a message at level Info on the standard logger.
func PrintfX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Infof(format, args...)
}

// InfofX logs a message at level Info on the standard logger.
func InfofX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Infof(format, args...)
}

// WarnfX logs a message at level Warn on the standard logger.
func WarnfX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Warnf(format, args...)
}

// WarningfX logs a message at level Warn on the standard logger.
func WarningfX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Warnf(format, args...)
}

// ErrorfX logs a message at level Error on the standard logger.
func ErrorfX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Errorf(format, args...)
}

// PanicfX logs a message at level Panic on the standard logger.
func PanicfX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Panicf(format, args...)
}

// FatalfX logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalfX(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Fatalf(format, args...)
}
