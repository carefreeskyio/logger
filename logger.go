package logger

import (
	"context"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/logger/hook"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type BaseOptions struct {
	Path         string
	RotationTime time.Duration
	WithMaxAge   time.Duration
	Level        string
}

type CustomOptions struct {
	CtxFiles            []string
	InternalFilesPrefix string
}

var (
	baseOptions          *BaseOptions
	DefaultCustomOptions = &CustomOptions{
		CtxFiles:            []string{},
		InternalFilesPrefix: "x_",
	}
)

func initBaseOptions() {
	baseOptions = &BaseOptions{
		Path:         config.GetString("Log.Path"),
		RotationTime: config.GetDuration("Log.RotationTime") * time.Hour,
		WithMaxAge:   time.Duration(config.GetInt("Log.WithMaxAge")*24) * time.Hour,
		Level:        config.GetString("Log.Level"),
	}
}

func InitLogger() {
	initBaseOptions()

	logFile, err := rotatelogs.New(
		baseOptions.Path+".%Y%m%d_%H",
		rotatelogs.WithLinkName(baseOptions.Path),
		rotatelogs.WithRotationTime(baseOptions.RotationTime),
		rotatelogs.WithMaxAge(baseOptions.WithMaxAge),
	)
	if err != nil {
		log.Fatalf("init logger failed: err=%v", err)
	}
	stdout := os.Stdout
	log.SetOutput(io.MultiWriter(logFile, stdout))
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if level, err := log.ParseLevel(baseOptions.Level); err != nil {
		log.Fatalf("init log failed: err=%v", err)
		log.SetLevel(level)
	}

	log.AddHook(hook.CtxHook{
		Prefix: DefaultCustomOptions.InternalFilesPrefix,
		Fields: DefaultCustomOptions.CtxFiles,
	})
	log.AddHook(hook.CallerHook{
		Prefix: DefaultCustomOptions.InternalFilesPrefix,
	})
}

func AddCtxFields(fields []string) {
	DefaultCustomOptions.CtxFiles = append(DefaultCustomOptions.CtxFiles, fields...)
}

func AddHook(hook log.Hook) {
	log.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *log.Entry {
	return log.WithField(log.ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func WithContext(ctx context.Context) *log.Entry {
	return log.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *log.Entry {
	return log.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields log.Fields) *log.Entry {
	return log.WithFields(fields)
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *log.Entry {
	return log.WithTime(t)
}
