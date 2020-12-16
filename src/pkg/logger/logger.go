package logger

import (
	"github.com/sirupsen/logrus"
)

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}

func Logf(level logrus.Level, format string, args ...interface{}) {
	if log.IsLevelEnabled(level) {
		log.Logf(level, format, args...)
	}
}

func Tracef(format string, args ...interface{}) {
	log.Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Logf(logrus.InfoLevel, format, args...)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Logf(logrus.FatalLevel, format, args...)
	log.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	log.Logf(logrus.PanicLevel, format, args...)
}

func Log(level logrus.Level, args ...interface{}) {
	if log.IsLevelEnabled(level) {
		log.Log(level, args...)
	}
}

func Trace(args ...interface{}) {
	log.Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	log.Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	log.Log(logrus.InfoLevel, args...)
}

func Print(args ...interface{}) {
	log.Print(args...)
}

func Warn(args ...interface{}) {
	log.Log(logrus.WarnLevel, args...)
}

func Warning(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	log.Log(logrus.FatalLevel, args...)
	log.Exit(1)
}

func Panic(args ...interface{}) {
	log.Log(logrus.PanicLevel, args...)
}

func Logln(level logrus.Level, args ...interface{}) {
	if log.IsLevelEnabled(level) {
		log.Logln(level, args...)
	}
}
