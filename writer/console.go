package writer

import (
	"os"
	"strings"
)

// ConsoleConfig 用户自定义console配置
type ConsoleConfig struct {
	LogLevel string // 日志级别，可选，默认debug
}

// consoleWriter 实现writer接口
type consoleWriter struct {
	opt *ConsoleConfig
}

func (c *consoleWriter) GetLogLevel() string {
	return c.opt.LogLevel
}

func (c *consoleWriter) Check(entry *REntry) bool {
	return checkLogLevel(entry.Loglevel, c.opt.LogLevel)
}

func (c *consoleWriter) Write(message []byte) error {
	_, err := os.Stdout.Write(message)

	return err
}

func NewConsoleWriter(opt *ConsoleConfig) Writer {
	// LogLevel预处理
	opt.LogLevel = strings.ToLower(strings.TrimSpace(opt.LogLevel))

	if opt.LogLevel == "" {
		opt.LogLevel = "debug"
	}

	consoleWriter := consoleWriter{
		opt: opt,
	}
	return &consoleWriter
}
