package main

import "github.com/vvfly/rlog"

const (
	debugLogContent = "这是一个调试日志"
	infoLogContent  = "这是一个信息日志"
	warnLogContent  = "这是一个告警日志"
	errLogContent   = "这是一个错误日志"

	name        = "123456"
	logTemplate = "name=%s, content=%s"
)

func main() {
	// 日志打印
	logPrint()
}

func logPrint() {
	// debug日志输出
	rlog.Debug(debugLogContent)
	rlog.Debugf(logTemplate, name, debugLogContent)

	// info日志输出
	rlog.Info(infoLogContent)
	rlog.Infof(logTemplate, name, infoLogContent)

	// warn日志输出
	rlog.Warn(warnLogContent)
	rlog.Warnf(logTemplate, name, warnLogContent)

	// error日志输出
	rlog.Error(errLogContent)
	rlog.Errorf(logTemplate, name, errLogContent)
}
