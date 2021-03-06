

package xapp

import (
	"fmt"
	mdclog "nRIC/internal/xapp/golog"
	"time"
)

type Log struct {
	logger *mdclog.MdcLogger
}

func NewLogger(name string) *Log {
	l, _ := mdclog.InitLogger(name)
	return &Log{
		logger: l,
	}
}

func (l *Log) SetFormat(logMonitor int) {
	l.logger.Mdclog_format_initialize(logMonitor)
}

func (l *Log) SetLevel(level int) {
	l.logger.LevelSet(mdclog.Level(level))
}

func (l *Log) SetMdc(key string, value string) {
	l.logger.MdcAdd(key, value)
}

func (l *Log) GetLevel() mdclog.Level {
	return l.logger.LevelGet()
}

func (l *Log) Error(pattern string, args ...interface{}) {
	if l.logger.LevelGet() < mdclog.ERR {
		return
	}
	l.SetMdc("time", timeFormat())
	l.logger.Error(pattern, args...)
}

func (l *Log) Warn(pattern string, args ...interface{}) {
	if l.logger.LevelGet() < mdclog.WARN {
		return
	}
	l.SetMdc("time", timeFormat())
	l.logger.Warning(pattern, args...)
}

func (l *Log) Info(pattern string, args ...interface{}) {
	if l.logger.LevelGet() < mdclog.INFO {
		return
	}
	l.SetMdc("time", timeFormat())
	l.logger.Info(pattern, args...)
}

func (l *Log) Debug(pattern string, args ...interface{}) {
	if l.logger.LevelGet() < mdclog.DEBUG {
		return
	}
	l.SetMdc("time", timeFormat())
	l.logger.Debug(pattern, args...)
}

func timeFormat() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
