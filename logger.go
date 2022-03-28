package rlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"runtime/debug"
)

type Logger struct {
	logger *logrus.Logger
}

func (log *Logger) Debug(args ...interface{}) {
	log.logger.Debug(appendArgsWithLineNumber(args...)...)
}
func (log *Logger) Debugf(format string, args ...interface{}) {
	log.logger.Debugf(format+"%v", appendArgsWithLineNumber(args...)...)
}

func (log *Logger) Info(args ...interface{}) {
	log.logger.Info(appendArgsWithLineNumber(args...)...)
}
func (log *Logger) Infof(format string, args ...interface{}) {
	log.logger.Infof(format+"%v", appendArgsWithLineNumber(args...)...)
}

func (log *Logger) Warn(args ...interface{}) {
	log.logger.Warn(appendArgsWithLineNumber(args...)...)
}
func (log *Logger) Warnf(format string, args ...interface{}) {
	log.logger.Warnf(format+"%v", appendArgsWithLineNumber(args...)...)
}

func (log *Logger) Error(args ...interface{}) {
	log.logger.Error(appendArgsWithStack(args...)...)
}
func (log *Logger) Errorf(format string, args ...interface{}) {
	log.logger.Errorf(format+"%v%v", appendArgsWithStack(args...)...)
}

func getLineNumber(skip int) string {
	if pc, file, line, ok := runtime.Caller(skip); ok {
		funcName := runtime.FuncForPC(pc).Name()
		return fmt.Sprintf(" (%v:%v:%v)", path.Base(funcName), path.Base(file), line)
	}
	return " (no line number)"
}

func appendArgsWithLineNumber(args ...interface{}) []interface{} {
	lineNum := getLineNumber(4)
	var arr []interface{}
	arr = append(arr, args...)
	arr = append(arr, lineNum)
	return arr
}

func appendArgsWithStack(args ...interface{}) []interface{} {
	lineNum := getLineNumber(4)
	var arr []interface{}
	arr = append(arr, args...)
	arr = append(arr, lineNum)
	arr = append(arr, "\n"+string(debug.Stack()))
	return arr
}
