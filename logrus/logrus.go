package logrus

import (
	log "github.com/sirupsen/logrus"
)

type globalLogger int

func FromGlobal() *globalLogger {
	return new(globalLogger)
}

func (ent *globalLogger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

func (ent *globalLogger) Debug(v ...interface{}) {
	log.Debug(v...)
}

func (ent *globalLogger) Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func (ent *globalLogger) Info(v ...interface{}) {
	log.Info(v...)
}

func (ent *globalLogger) Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

func (ent *globalLogger) Warn(v ...interface{}) {
	log.Warn(v...)
}

func (ent *globalLogger) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

func (ent *globalLogger) Error(v ...interface{}) {
	log.Error(v...)
}
