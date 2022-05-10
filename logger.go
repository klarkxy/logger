package logger

import (
	"io"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level = logrus.Level

const (
	TraceLevel = logrus.TraceLevel
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
	FatalLevel = logrus.FatalLevel
	PanicLevel = logrus.PanicLevel
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(NewFormatter())
	Logger.SetLevel(DebugLevel)
	Logger.SetReportCaller(true)
}

func AddOutput(writer io.Writer) {
	Logger.SetOutput(io.MultiWriter(Logger.Out, writer))
}

// 增加一个文件输出
// filename 文件名
// maxAgeDays 文件最多保留天数
func AddFile(filename string, maxAgeDays int) {
	AddOutput(&ColorCleaner{Writer: &lumberjack.Logger{
		Filename:  filename,
		MaxAge:    maxAgeDays,
		LocalTime: true,
		Compress:  false,
	}})
}
