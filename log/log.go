package log

import (
	_ "golang.org/x/sys/unix"
	"os"
	"log"
	"github.com/sirupsen/logrus"
	"github.com/wiloon/app-config"
)

var fileLogger *logrus.Logger
var stdLogger *logrus.Logger

var file *os.File
var logConfig *Conf

func init() {
	logConfig = (&Conf{
		Level:       config.GetString("log.level", "debug"),
		To:          config.GetString("log.to", "file"),
		ProjectName: config.GetString(config.KeyProjectName, ""),
	}).Init()

	fileLogger = getFileLogger()
	stdLogger = getStdLogger()
}

func Info(args ...interface{}) {
	fileLogger.Info(args...)

	if logConfig.ToConsole {
		stdLogger.Info(args...)
	}
}

func Debug(args ...interface{}) {
	fileLogger.Debug(args...)

	if logConfig.ToConsole {
		stdLogger.Debug(args...)
	}
}

func getStdLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logConfig.GetLogrusLevel())
	logger.Out = os.Stdout
	logger.Formatter = &logrus.TextFormatter{}

	logger.Info("logger init done, logrus")
	return logger
}

func getFileLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logConfig.GetLogrusLevel())

	os.MkdirAll(logConfig.GetParentPath(), 0744)
	logFile := logConfig.GetFullPath()

	var err error
	file, err = os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		log.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr,error:", err)
	}
	logger.Out = file
	logger.Formatter = &logrus.TextFormatter{}

	logger.Info("logger init done, logrus")
	return logger
}
