package logger

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func getLogger() *log.Entry {
	return log.WithFields(log.Fields{
		"service":   os.Getenv("SERVICE_NAME"),
		"timestamp": time.Now(),
	})
}

func Print(msg ...interface{}) {
	getLogger().Print(msg)
}

func Printf(format string, msg ...interface{}) {
	getLogger().Printf(format, msg)
}

func Error(msg ...interface{}) {
	getLogger().Error(msg)
}

func Errorf(format string, msg ...interface{}) {
	getLogger().Errorf(format, msg)
}

func Info(msg ...interface{}) {
	getLogger().Info(msg)
}

func Infof(format string, msg ...interface{}) {
	getLogger().Infof(format, msg)
}

func Debug(msg ...interface{}) {
	getLogger().Debug(msg)
}

func Debugf(format string, msg ...interface{}) {
	getLogger().Debugf(format, msg)
}

func Warn(msg ...interface{}) {
	getLogger().Warn(msg)
}

func Warnf(format string, msg ...interface{}) {
	getLogger().Warnf(format, msg)
}

func Trace(msg ...interface{}) {
	getLogger().Trace(msg)
}

func Tracef(format string, msg ...interface{}) {
	getLogger().Tracef(format, msg)
}

func Fatal(msg ...interface{}) {
	getLogger().Fatal(msg)
}

func Fatalf(format string, msg ...interface{}) {
	getLogger().Fatalf(format, msg)
}

func Panic(msg ...interface{}) {
	getLogger().Panic(msg)
}

func Panicf(format string, msg ...interface{}) {
	getLogger().Panicf(format, msg)
}
