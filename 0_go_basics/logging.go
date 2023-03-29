package main

import (
	"log"
	"os"
)

type LogLevel int

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func (le *LogExtended) SetLogLevel(logLvl LogLevel) {
	le.logLevel = logLvl
}

func (le *LogExtended) Infoln(s string) {
	if le.logLevel >= LogLevelInfo {
		le.Logger.Println(s)
	}
}

func (le *LogExtended) Warnln(s string) {
	if le.logLevel >= LogLevelWarning {
		le.Logger.Println(s)
	}
}

func (le *LogExtended) Errorln(s string) {
	if le.logLevel >= LogLevelError {
		le.Logger.Println(s)
	}
}

func NewLogExtended() LogExtended {
	return LogExtended{logLevel: 0, Logger: log.New(os.Stdout, "", log.LstdFlags)}
}

func LogTest() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
