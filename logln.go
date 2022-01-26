package logger

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infoln(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Warnln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	log.NewEntry(log.StandardLogger()).Fatalln(args...)
}

// TracelnX logs a message at level Trace on the standard logger.
func TracelnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Traceln(args...)
}

// DebuglnX logs a message at level Debug on the standard logger.
func DebuglnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Debugln(args...)
}

// PrintlnX logs a message at level Info on the standard logger.
func PrintlnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Infoln(args...)
}

// InfolnX logs a message at level Info on the standard logger.
func InfolnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Infoln(args...)
}

// WarnlnX logs a message at level Warn on the standard logger.
func WarnlnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Warnln(args...)
}

// WarninglnX logs a message at level Warn on the standard logger.
func WarninglnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Warnln(args...)
}

// ErrorlnX logs a message at level Error on the standard logger.
func ErrorlnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Errorln(args...)
}

// PaniclnX logs a message at level Panic on the standard logger.
func PaniclnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Panicln(args...)
}

// FatallnX logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func FatallnX(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Fatalln(args...)
}
