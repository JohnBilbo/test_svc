package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"test_svc/internal/constants"
)

func InitLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	return log
}

func FillLoggerFields(log *logrus.Logger, uuid, class, methodType string) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		constants.LoggerFieldUuid:   uuid,
		constants.LoggerFieldClass:  class,
		constants.LoggerFieldMethod: methodType,
	})
}
