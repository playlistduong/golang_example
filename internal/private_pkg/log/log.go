package log

import (
	log "github.com/sirupsen/logrus"
)

func Warn(format string) {
	log.Warn(format)
}

func Info(format string) {
	log.Info(format)
}

func Debug(format string) {
	log.Debug(format)
}

func Panic(format string) {
	log.Panic(format)
}

func Fatal(format string) {
	log.Fatal(format)
}
