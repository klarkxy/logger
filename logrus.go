package logger

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

func WithField(key string, value interface{}) *logrus.Entry {
	return Logger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return Logger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return Logger.WithError(err)
}

func WithContext(ctx context.Context) *logrus.Entry {
	return Logger.WithContext(ctx)
}

func WithTime(t time.Time) *logrus.Entry {
	return Logger.WithTime(t)
}

func Logf(level Level, format string, args ...interface{}) {
	Logger.Logf(level, format, args...)
}

func Tracef(format string, args ...interface{}) {
	Logger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	Logger.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	Logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	Logger.Panicf(format, args...)
}

func Log(level Level, args ...interface{}) {
	Logger.Log(level, args...)
}

func LogFn(level Level, fn logrus.LogFunction) {
	Logger.LogFn(level, fn)
}

func Trace(args ...interface{}) {
	Logger.Trace(args...)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Print(args ...interface{}) {
	Logger.Print(args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

func Warning(args ...interface{}) {
	Logger.Warning(args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

func TraceFn(fn logrus.LogFunction) {
	Logger.LogFn(TraceLevel, fn)
}

func DebugFn(fn logrus.LogFunction) {
	Logger.LogFn(DebugLevel, fn)
}

func InfoFn(fn logrus.LogFunction) {
	Logger.InfoFn(fn)
}

func PrintFn(fn logrus.LogFunction) {
	Logger.PrintFn(fn)
}

func WarnFn(fn logrus.LogFunction) {
	Logger.WarnFn(fn)
}

func WarningFn(fn logrus.LogFunction) {
	Logger.WarningFn(fn)
}

func ErrorFn(fn logrus.LogFunction) {
	Logger.ErrorFn(fn)
}

func FatalFn(fn logrus.LogFunction) {
	Logger.FatalFn(fn)
}

func PanicFn(fn logrus.LogFunction) {
	Logger.PanicFn(fn)
}

func Logln(level Level, args ...interface{}) {
	Logger.Logln(level, args...)
}

func Traceln(args ...interface{}) {
	Logger.Traceln(args...)
}

func Debugln(args ...interface{}) {
	Logger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	Logger.Infoln(args...)
}

func Println(args ...interface{}) {
	Logger.Println(args...)
}

func Warnln(args ...interface{}) {
	Logger.Warnln(args...)
}

func Warningln(args ...interface{}) {
	Logger.Warningln(args...)
}

func Errorln(args ...interface{}) {
	Logger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	Logger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	Logger.Panicln(args...)
}

func Exit(code int) {
	Logger.Exit(code)
}

func SetNoLock() {
	Logger.SetNoLock()
}

func SetLevel(level Level) {
	Logger.SetLevel(level)
}

func GetLevel() Level {
	return Logger.GetLevel()
}

func AddHook(hook logrus.Hook) {
	Logger.AddHook(hook)
}

func IsLevelEnabled(level Level) bool {
	return Logger.IsLevelEnabled(level)
}

func SetFormatter(formatter logrus.Formatter) {
	Logger.SetFormatter(formatter)
}

func SetOutput(output io.Writer) {
	Logger.SetOutput(output)
}

func SetReportCaller(reportCaller bool) {
	Logger.SetReportCaller(reportCaller)
}

func ReplaceHooks(hooks logrus.LevelHooks) logrus.LevelHooks {
	return Logger.ReplaceHooks(hooks)
}
