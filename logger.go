package logger

import (
	"context"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/logger/hook"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"time"
)

type BaseOptions struct {
	Path         string
	RotationTime time.Duration
	WithMaxAge   time.Duration
	Level        string
}

type CustomOptions struct {
	CtxFiles    []string
	FieldPrefix string
}

var (
	baseOptions          *BaseOptions
	DefaultCustomOptions = &CustomOptions{
		CtxFiles:    []string{},
		FieldPrefix: "x_",
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

	writer, err := rotatelogs.New(
		baseOptions.Path+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(baseOptions.Path),
		rotatelogs.WithRotationTime(baseOptions.RotationTime),
		rotatelogs.WithMaxAge(baseOptions.WithMaxAge),
	)
	if err != nil {
		log.Fatalf("init logger failed: err=%v", err)
	}
	log.SetOutput(writer)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if level, err := log.ParseLevel(baseOptions.Level); err != nil {
		log.Fatalf("init log failed: err=%v", err)
		log.SetLevel(level)
	}
	log.SetReportCaller(true)

	log.AddHook(hook.CtxHook{
		Prefix: DefaultCustomOptions.FieldPrefix,
		Fields: DefaultCustomOptions.CtxFiles,
	})
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

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	log.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	log.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	log.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	log.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// TraceFn logs a message from a func at level Trace on the standard logger.
func TraceFn(fn log.LogFunction) {
	log.TraceFn(fn)
}

// DebugFn logs a message from a func at level Debug on the standard logger.
func DebugFn(fn log.LogFunction) {
	log.DebugFn(fn)
}

// PrintFn logs a message from a func at level Info on the standard logger.
func PrintFn(fn log.LogFunction) {
	log.PrintFn(fn)
}

// InfoFn logs a message from a func at level Info on the standard logger.
func InfoFn(fn log.LogFunction) {
	log.InfoFn(fn)
}

// WarnFn logs a message from a func at level Warn on the standard logger.
func WarnFn(fn log.LogFunction) {
	log.WarnFn(fn)
}

// WarningFn logs a message from a func at level Warn on the standard logger.
func WarningFn(fn log.LogFunction) {
	log.WarningFn(fn)
}

// ErrorFn logs a message from a func at level Error on the standard logger.
func ErrorFn(fn log.LogFunction) {
	log.ErrorFn(fn)
}

// PanicFn logs a message from a func at level Panic on the standard logger.
func PanicFn(fn log.LogFunction) {
	log.PanicFn(fn)
}

// FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFn(fn log.LogFunction) {
	log.FatalFn(fn)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	log.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	log.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	log.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	log.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	log.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

// Tracex logs a message at level Trace on the standard logger.
func Tracex(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Trace(args...)
}

// Debugx logs a message at level Debug on the standard logger.
func Debugx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Debug(args...)
}

// Printx logs a message at level Info on the standard logger.
func Printx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Print(args...)
}

// Infox logs a message at level Info on the standard logger.
func Infox(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Info(args...)
}

// Warnx logs a message at level Warn on the standard logger.
func Warnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Warn(args...)
}

// Warningx logs a message at level Warn on the standard logger.
func Warningx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).WithContext(ctx).Warning(args...)
}

// Errorx logs a message at level Error on the standard logger.
func Errorx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Error(args...)
}

// Panicx logs a message at level Panic on the standard logger.
func Panicx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Panic(args...)
}

// Fatalx logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Fatal(args...)
}

// TraceFnx logs a message from a func at level Trace on the standard logger.
func TraceFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.TraceFn(fn)
}

// DebugFnx logs a message from a func at level Debug on the standard logger.
func DebugFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.DebugFn(fn)
}

// PrintFnx logs a message from a func at level Info on the standard logger.
func PrintFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.PrintFn(fn)
}

// InfoFnx logs a message from a func at level Info on the standard logger.
func InfoFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.InfoFn(fn)
}

// WarnFnx logs a message from a func at level Warn on the standard logger.
func WarnFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.WarnFn(fn)
}

// WarningFnx logs a message from a func at level Warn on the standard logger.
func WarningFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.WarningFn(fn)
}

// ErrorFnx logs a message from a func at level Error on the standard logger.
func ErrorFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.ErrorFn(fn)
}

// PanicFnx logs a message from a func at level Panic on the standard logger.
func PanicFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.PanicFn(fn)
}

// FatalFnx logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFnx(ctx context.Context, tag string, fn log.LogFunction) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag)
	log.FatalFn(fn)
}

// Tracefx logs a message at level Trace on the standard logger.
func Tracefx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Tracef(format, args...)
}

// Debugfx logs a message at level Debug on the standard logger.
func Debugfx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Debugf(format, args...)
}

// Printfx logs a message at level Info on the standard logger.
func Printfx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Printf(format, args...)
}

// Infofx logs a message at level Info on the standard logger.
func Infofx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Infof(format, args...)
}

// Warnfx logs a message at level Warn on the standard logger.
func Warnfx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Warnf(format, args...)
}

// Warningfx logs a message at level Warn on the standard logger.
func Warningfx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Warningf(format, args...)
}

// Errorfx logs a message at level Error on the standard logger.
func Errorfx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Errorf(format, args...)
}

// Panicfx logs a message at level Panic on the standard logger.
func Panicfx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Panicf(format, args...)
}

// Fatalfx logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalfx(ctx context.Context, tag string, format string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Fatalf(format, args...)
}

// Tracelnx logs a message at level Trace on the standard logger.
func Tracelnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Traceln(args...)
}

// Debuglnx logs a message at level Debug on the standard logger.
func Debuglnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Debugln(args...)
}

// Printlnx logs a message at level Info on the standard logger.
func Printlnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Println(args...)
}

// Infolnx logs a message at level Info on the standard logger.
func Infolnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Infoln(args...)
}

// Warnlnx logs a message at level Warn on the standard logger.
func Warnlnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Warnln(args...)
}

// Warninglnx logs a message at level Warn on the standard logger.
func Warninglnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Warningln(args...)
}

// Errorlnx logs a message at level Error on the standard logger.
func Errorlnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Errorln(args...)
}

// Paniclnx logs a message at level Panic on the standard logger.
func Paniclnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Panicln(args...)
}

// Fatallnx logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatallnx(ctx context.Context, tag string, args ...interface{}) {
	log.WithContext(ctx).WithField(DefaultCustomOptions.FieldPrefix+"tag", tag).Fatalln(args...)
}
