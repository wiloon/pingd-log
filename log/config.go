package log

import (
	"strings"
	"path/filepath"
	"github.com/sirupsen/logrus"
)

type Conf struct {
	Level       string
	To          string
	LogFilePath string
	ToConsole   bool
	ProjectName string
}

func (config *Conf) GetParentPath() string {
	return "/data/" + config.ProjectName + "/logs"
}

func (config *Conf) GetFullPath() string {
	return config.GetParentPath() + string(filepath.Separator) + config.Level + ".log"
}

func (config *Conf) Init() *Conf {
	if strings.Contains(config.To, "console") {
		config.ToConsole = true
	}
	return config
}

func (config *Conf) GetLogrusLevel() logrus.Level {
	var level logrus.Level
	if strings.EqualFold(config.Level, "debug") {
		level = logrus.DebugLevel
	}
	if strings.EqualFold(config.Level, "info") {
		level = logrus.InfoLevel
	}
	if strings.EqualFold(config.Level, "warn") {
		level = logrus.WarnLevel
	}
	return level
}
