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
	if !utils.IsDirExists(path) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	file, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {

		logrus.SetOutput(file)
	}else {
		logrus.Info("Failed to log to file, using default stderr,error:", err)
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
