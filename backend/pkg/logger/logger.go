package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

type logger struct {
	*logrus.Entry
}

func New() Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	
	entry := l.WithFields(logrus.Fields{
		"service": "switch-manager",
	})
	
	return &logger{entry}
}

func (l *logger) WithField(key string, value interface{}) Logger {
	return &logger{l.Entry.WithField(key, value)}
}

func (l *logger) WithFields(fields map[string]interface{}) Logger {
	return &logger{l.Entry.WithFields(fields)}
}
