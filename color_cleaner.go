package logger

import (
	"io"

	"github.com/gookit/color"
)

type ColorCleaner struct {
	Writer io.Writer
}

func (c *ColorCleaner) Write(p []byte) (n int, err error) {
	str := string(p)
	data := color.ClearCode(str)
	n, err = c.Writer.Write([]byte(data))
	if err != nil {
		return n, err
	}
	return len(p), nil
}
