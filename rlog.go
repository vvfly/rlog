package rlog

import (
	"github.com/sirupsen/logrus"
	"github.com/vvfly/rlog/formatter"
	"github.com/vvfly/rlog/writer"
)

// global logger, can be reset by Use()
var _logger *Logger

func init() {
	initLogger("DEBUG")
}

// initLogger init a default logger
func initLogger(level string) {
	// console
	consoleWriter := writer.NewConsoleWriter(&writer.ConsoleConfig{})
	// file
	fileWriter := writer.NewFileWriter(&writer.FileConfig{})
	Use(consoleWriter, fileWriter)
}

// Use 注册writer
func Use(writers ...writer.Writer) error {
	hook := &RHook{
		writers: writers,
		Formatter: &formatter.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.000",
		},
	}

	_logger = &Logger{
		logger: logrus.New(),
	}
	_logger.logger.AddHook(hook)
	_logger.logger.SetLevel(writer.GetLowLogLevel(writers))
	return nil
}

func Debug(args ...interface{}) {
	_logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	_logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	_logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	_logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	_logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	_logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	_logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	_logger.Errorf(format, args...)
}
