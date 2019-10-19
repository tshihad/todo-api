package core

import (
	"io"

	"github.com/sirupsen/logrus"
)

// NewLogger returns logrus.logger
func NewLogger(level logrus.Level, out io.Writer) *logrus.Logger {
	l := &logrus.Logger{
		Out:   out,
		Level: level,
		Formatter: &logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		},
	}
	return l
}
