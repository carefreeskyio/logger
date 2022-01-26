package logger

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// TraceFn logs a message from a func at level Trace on the standard logger.
func TraceFn(fn log.LogFunction) {
	log.StandardLogger().TraceFn(fn)
}

// DebugFn logs a message from a func at level Debug on the standard logger.
func DebugFn(fn log.LogFunction) {
	log.StandardLogger().DebugFn(fn)
}

// PrintFn logs a message from a func at level Info on the standard logger.
func PrintFn(fn log.LogFunction) {
	log.StandardLogger().InfoFn(fn)
}

// InfoFn logs a message from a func at level Info on the standard logger.
func InfoFn(fn log.LogFunction) {
	log.StandardLogger().InfoFn(fn)
}

// WarnFn logs a message from a func at level Warn on the standard logger.
func WarnFn(fn log.LogFunction) {
	log.StandardLogger().WarnFn(fn)
}

// WarningFn logs a message from a func at level Warn on the standard logger.
func WarningFn(fn log.LogFunction) {
	log.StandardLogger().WarnFn(fn)
}

// ErrorFn logs a message from a func at level Error on the standard logger.
func ErrorFn(fn log.LogFunction) {
	log.StandardLogger().ErrorFn(fn)
}

// PanicFn logs a message from a func at level Panic on the standard logger.
func PanicFn(fn log.LogFunction) {
	log.StandardLogger().PanicFn(fn)
}

// FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFn(fn log.LogFunction) {
	log.StandardLogger().FatalFn(fn)
}

// TraceFnX logs a message from a func at level Trace on the standard logger.
func TraceFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Traceln(fn()...)
}

// DebugFnX logs a message from a func at level Debug on the standard logger.
func DebugFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Debugln(fn()...)
}

// PrintFnX logs a message from a func at level Info on the standard logger.
func PrintFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Infoln(fn()...)
}

// InfoFnX logs a message from a func at level Info on the standard logger.
func InfoFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Infoln(fn()...)
}

// WarnFnX logs a message from a func at level Warn on the standard logger.
func WarnFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Warnln(fn()...)
}

// WarningFnX logs a message from a func at level Warn on the standard logger.
func WarningFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Warnln(fn()...)
}

// ErrorFnX logs a message from a func at level Error on the standard logger.
func ErrorFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Errorln(fn()...)
}

// PanicFnX logs a message from a func at level Panic on the standard logger.
func PanicFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Panicln(fn()...)
}

// FatalFnX logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFnX(ctx context.Context, fn log.LogFunction) {
	log.NewEntry(log.StandardLogger()).WithContext(ctx).Fatalln(fn()...)
}
