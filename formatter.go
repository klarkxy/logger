package logger

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/gookit/color"

	"github.com/sirupsen/logrus"
)

type Sprinter interface {
	Sprint(a ...interface{}) string
}

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	TimestampFormat  string
	CallerPrettyfier func(*runtime.Frame) (ret string)
	LevelColor       map[logrus.Level]Sprinter
	NoColor          bool
}

func NewFormatter() *Formatter {
	return &Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) string {
			if f != nil {
				return fmt.Sprintf("[%s:%d] ", f.File, f.Line)
			}
			return ""
		},
		LevelColor: map[logrus.Level]Sprinter{
			logrus.TraceLevel: color.Gray,
			logrus.DebugLevel: color.Green,
			logrus.InfoLevel:  color.Cyan,
			logrus.WarnLevel:  color.Magenta,
			logrus.ErrorLevel: color.Red,
			logrus.FatalLevel: color.BgRed,
			logrus.PanicLevel: color.BgMagenta,
		},
		NoColor: false,
	}
}

var bufferPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 对error进行特殊处理
	if field, ok := entry.Data[logrus.ErrorKey]; ok {
		if err, ok := field.(error); ok {
			entry.Data[logrus.ErrorKey] = struct {
				Text  string
				Error error
			}{Text: err.Error(), Error: err}
		}
	}
	// 根据颜色选择合适的Spirnt
	Sprint := func(a ...interface{}) string {
		if f.NoColor {
			return fmt.Sprint(a...)
		}

		col, has := f.LevelColor[entry.Level]
		if !has {
			col = color.White
		}
		return col.Sprint(a...)
	}

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)
	// 写入时间
	buf.WriteString(Sprint(entry.Time.Format(f.TimestampFormat)))
	// 写入日志等级
	buf.WriteString(Sprint(" [", strings.ToUpper(entry.Level.String()), "] "))
	// 写入调用者
	buf.WriteString(Sprint(f.CallerPrettyfier(entry.Caller)))
	// 写入日志信息
	buf.WriteString(Sprint(entry.Message))
	// 换行后写入YAML
	buf.WriteString("\n")
	if entry.Data != nil && len(entry.Data) > 0 {
		yaml, _ := yaml.Marshal(entry.Data)
		buf.WriteString(Sprint(string(yaml)))
	}

	// 返回缓冲区的内容
	return buf.Bytes(), nil
}
