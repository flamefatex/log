package rotation

import (
	"fmt"
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func NewRotationWriter(serviceName string, c *RotationConfig) io.Writer {
	if c.Filename == "" {
		c.Filename = fmt.Sprintf("/var/log/%s/error.log", serviceName)
	}
	if c.MaxSize == 0 {
		c.MaxSize = 300
	}
	if c.MaxBackups == 0 {
		c.MaxBackups = 2
	}
	return &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize, // megabytes
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,   //days
		Compress:   c.Compress, // disabled by default
		LocalTime:  c.LocalTime,
	}
}
