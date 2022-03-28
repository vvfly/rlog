package writer

import "github.com/sirupsen/logrus"

// REntry 日志内容
type REntry struct {
	Loglevel logrus.Level // 日志级别
	Message  []byte       // 日志内容
}

type Writer interface {
	// GetLogLevel 获取日志级别
	GetLogLevel() string
	// Check 日志检查，过滤掉不必要的日志输入
	Check(*REntry) bool
	// Write 日志输出内容
	Write([]byte) error
}
