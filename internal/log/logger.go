package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var instance *zap.SugaredLogger

func init() {
	stdOutWriterCore := zapcore.NewCore(
		createEncoder(),
		createStdoutWriter(),
		zapcore.DebugLevel,
	)
	instance = zap.New(stdOutWriterCore, zap.AddCaller()).Sugar()
}

func createEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func createStdoutWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func Info(template string, args ...interface{}) {
	instance.Infof(template, args)
}

func Debug(template string, args ...interface{}) {
	instance.Debugf(template, args)
}

func Warn(template string, args ...interface{}) {
	instance.Warnf(template, args)
}

func Error(template string, args ...interface{}) {
	instance.Errorf(template, args)
}
