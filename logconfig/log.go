package logconfig

import (
	"github.com/sirupsen/logrus"
	"github.com/wiloon/pingd-config"
	"github.com/wiloon/pingd-utils/utils"
	"os"
	"strings"
)

func Init() {
	level := config.GetString("log.level", "debug")
	path := config.GetString("log.path", "/tmp/")
	fileName := config.GetString("log.file-name", "foo.log") //todo project-name.log
	if !utils.IsFileOrDirExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logrus.Errorf("failed to create dir: %v, e: %v", path, err)
		} else {
			logrus.Infof("create new dir: %v", path)
		}

	}
	fullPath := path + fileName
	logrus.Infof("open log file: %v", fullPath)
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Infof("Failed to log to file, using default stderr,error: %v", err)
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(stringLevelToLogrusLevel(level))

}

func stringLevelToLogrusLevel(str string) logrus.Level {
	var level logrus.Level
	if strings.EqualFold(str, "debug") {
		level = logrus.DebugLevel
	}
	if strings.EqualFold(str, "info") {
		level = logrus.InfoLevel
	}
	if strings.EqualFold(str, "warn") {
		level = logrus.WarnLevel
	}
	return level
}
