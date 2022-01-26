package logger

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	log.StandardLogger().Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.StandardLogger().Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	log.StandardLogger().Info(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	log.StandardLogger().Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	log.StandardLogger().Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	log.StandardLogger().Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	log.StandardLogger().Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	log.StandardLogger().Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	log.StandardLogger().Fatal(args...)
}

// TraceX logs a message at level Trace on the standard logger.
func TraceX(ctx context.Context, args ...interface{}) {
	logX(log.TraceLevel, ctx, args...)
}

// DebugX logs a message at level Debug on the standard logger.
func DebugX(ctx context.Context, args ...interface{}) {
	logX(log.DebugLevel, ctx, args...)
}

// PrintX logs a message at level Info on the standard logger.
func PrintX(ctx context.Context, args ...interface{}) {
	logX(log.InfoLevel, ctx, args...)
}

// InfoX logs a message at level Info on the standard logger.
func InfoX(ctx context.Context, args ...interface{}) {
	logX(log.InfoLevel, ctx, args...)
}

// WarnX logs a message at level Warn on the standard logger.
func WarnX(ctx context.Context, args ...interface{}) {
	logX(log.WarnLevel, ctx, args...)
}

// WarningX logs a message at level Warn on the standard logger.
func WarningX(ctx context.Context, args ...interface{}) {
	logX(log.WarnLevel, ctx, args...)
}

// ErrorX logs a message at level Error on the standard logger.
func ErrorX(ctx context.Context, args ...interface{}) {
	logX(log.ErrorLevel, ctx, args...)
}

// PanicX logs a message at level Panic on the standard logger.
func PanicX(ctx context.Context, args ...interface{}) {
	logX(log.PanicLevel, ctx, args...)
}

// FatalX logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalX(ctx context.Context, args ...interface{}) {
	logX(log.FatalLevel, ctx, args...)
}

func logX(level log.Level, ctx context.Context, args ...interface{}) {
	switch level {
	case log.InfoLevel:
		log.WithContext(ctx).Info(args...)
	case log.PanicLevel:
		log.WithContext(ctx).Panic(args...)
	case log.FatalLevel:
		log.WithContext(ctx).Fatal(args...)
	case log.ErrorLevel:
		log.WithContext(ctx).Error(args...)
	case log.WarnLevel:
		log.WithContext(ctx).Warn(args...)
	case log.DebugLevel:
		log.WithContext(ctx).Debug(args...)
	case log.TraceLevel:
		log.WithContext(ctx).Trace(args...)
	}
}
