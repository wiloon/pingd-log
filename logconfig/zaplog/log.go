package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.SugaredLogger

func Init() {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stderr), zapcore.DebugLevel),
	)
	logger = zap.New(core).Sugar()

}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}
func Info(args ...interface{}) {
	logger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}
func Warn(args ...interface{}) {
	logger.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}
func Error(args ...interface{}) {
	logger.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
