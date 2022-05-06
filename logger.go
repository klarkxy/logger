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

type Logger struct {
	*logrus.Logger
}

func New(level Level) *Logger {
	logger := logrus.New()
	logger.SetFormatter(NewFormatter())
	logger.SetLevel(level)
	logger.SetReportCaller(true)
	return &Logger{Logger: logger}
}

func (log *Logger) AddOutput(writer io.Writer) {
	log.SetOutput(io.MultiWriter(log.Out, writer))
}

// 增加一个文件输出
// filename 文件名
// maxAgeDays 文件最多保留天数
func (log *Logger) AddFile(filename string, maxAgeDays int) {
	log.AddOutput(&ColorCleaner{Writer: &lumberjack.Logger{
		Filename:  filename,
		MaxAge:    maxAgeDays,
		LocalTime: true,
		Compress:  false,
	}})
}
