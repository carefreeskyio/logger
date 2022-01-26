package logger

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// TraceLn logs a message at level Trace on the standard logger.
func TraceLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Traceln(args...)
}

// DebugLn logs a message at level Debug on the standard logger.
func DebugLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Debugln(args...)
}

// PrintLn logs a message at level Info on the standard logger.
func PrintLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infoln(args...)
}

// InfoLn logs a message at level Info on the standard logger.
func InfoLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infoln(args...)
}

// WarnLn logs a message at level Warn on the standard logger.
func WarnLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnln(args...)
}

// WarningLn logs a message at level Warn on the standard logger.
func WarningLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnln(args...)
}

// ErrorLn logs a message at level Error on the standard logger.
func ErrorLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Errorln(args...)
}

// PanicLn logs a message at level Panic on the standard logger.
func PanicLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Panicln(args...)
}

// FatalLn logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalLn(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Fatalln(args...)
}

// TraceLnX logs a message at level Trace on the standard logger.
func TraceLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Traceln(args...)
}

// DebugLnX logs a message at level Debug on the standard logger.
func DebugLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Debugln(args...)
}

// PrintLnX logs a message at level Info on the standard logger.
func PrintLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Infoln(args...)
}

// InfoLnX logs a message at level Info on the standard logger.
func InfoLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Infoln(args...)
}

// WarnLnX logs a message at level Warn on the standard logger.
func WarnLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Warnln(args...)
}

// WarningLnX logs a message at level Warn on the standard logger.
func WarningLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Warnln(args...)
}

// ErrorLnX logs a message at level Error on the standard logger.
func ErrorLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Errorln(args...)
}

// PanicLnX logs a message at level Panic on the standard logger.
func PanicLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Panicln(args...)
}

// FatalLnX logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalLnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Fatalln(args...)
}
