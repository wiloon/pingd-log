package logconfig

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func Test0(t *testing.T) {
	Init()
	logrus.Info("foo")
	logrus.Debugf("bar")
}
