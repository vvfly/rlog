package writer

import (
	"strings"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// FileConfig 用户自定义File配置
type FileConfig struct {
	Filename    string // 日志文件, 默认"./log/rlog.log",可选
	MaxFileSize int    // 日志文件单个文件最大大小，单位为MB，默认100,可选
	MaxBackups  int    // 日志文件最大历史保留份数，默认5,可选
	MaxAge      int    // 日志文件最长存活时间，单位为天， 默认30,可选
	Compress    bool   // 历史日志压缩保存，默认false,可选
	LogLevel    string // 日志级别，默认debug,可选
}

// fileWriter 实现writer接口
type fileWriter struct {
	log *lumberjack.Logger
	opt *FileConfig
}

func (f *fileWriter) GetLogLevel() string {
	return f.opt.LogLevel
}

func (f *fileWriter) Check(entry *REntry) bool {
	return checkLogLevel(entry.Loglevel, f.opt.LogLevel)
}

func (f *fileWriter) Write(message []byte) error {
	_, err := f.log.Write(message)
	return err
}

func NewFileWriter(opt *FileConfig) Writer {
	// LogLevel预处理
	opt.LogLevel = strings.ToLower(strings.TrimSpace(opt.LogLevel))

	if opt.Filename == "" {
		opt.Filename = "./log/rlog.log"
	}
	if opt.MaxFileSize == 0 {
		opt.MaxFileSize = 100
	}
	if opt.MaxAge == 0 {
		opt.MaxAge = 30
	}
	if opt.MaxBackups == 0 {
		opt.MaxBackups = 5
	}
	if opt.LogLevel == "" {
		opt.LogLevel = "debug"
	}

	// 日志输出控制
	lumberJackLogger := &lumberjack.Logger{
		Filename:   opt.Filename,
		MaxSize:    opt.MaxFileSize,
		MaxBackups: opt.MaxBackups,
		MaxAge:     opt.MaxAge,
		Compress:   opt.Compress,
	}

	fileWriter := fileWriter{
		log: lumberJackLogger,
		opt: opt,
	}
	return &fileWriter
}
