package writer

import "github.com/sirupsen/logrus"

var levelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

func checkLogLevel(entryLv logrus.Level, checkLv string) bool {
	return GetLoggerLevel(checkLv) >= entryLv
}

func GetLoggerLevel(lvl string) logrus.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return logrus.InfoLevel
}

// GetLowLogLevel 获取writer中最低日志级别
func GetLowLogLevel(writers []Writer) logrus.Level {
	var level = logrus.ErrorLevel
	for i := range writers {
		logLevel := writers[i].GetLogLevel()
		rLogLevel := GetLoggerLevel(logLevel)
		if rLogLevel >= level {
			level = rLogLevel
		}
	}
	return level
}
