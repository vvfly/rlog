package rlog

import (
	"github.com/sirupsen/logrus"
	"github.com/vvfly/rlog/writer"
	"io/ioutil"
)

// RHook 实现logrus Hook接口
type RHook struct {
	writers   []writer.Writer
	Formatter logrus.Formatter
}

func (hook *RHook) Levels() []logrus.Level {

	var levels []logrus.Level
	levels = append(levels, logrus.AllLevels...)

	return levels
}

func (hook *RHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}

	REntry := writer.REntry{
		Loglevel: entry.Level,
		Message:  b,
	}

	for i := range hook.writers {
		if hook.writers[i].Check(&REntry) {
			hook.writers[i].Write(b)
		}
	}

	entry.Logger.Out = ioutil.Discard

	return nil
}
