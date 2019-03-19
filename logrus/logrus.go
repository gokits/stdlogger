package logrus

import (
	log "github.com/sirupsen/logrus"
)

type entryLogger struct {
	entry *log.Entry
}

func FromEntry(entry *log.Entry) *entryLogger {
	return &entryLogger{
		entry: entry,
	}
}

func (ent *entryLogger) Debugf(format string, v ...interface{}) {
	ent.entry.Debugf(format, v...)
}

func (ent *entryLogger) Debug(v ...interface{}) {
	ent.entry.Debug(v...)
}

func (ent *entryLogger) Debugln(v ...interface{}) {
	ent.entry.Debugln(v...)
}

func (ent *entryLogger) Infof(format string, v ...interface{}) {
	ent.entry.Infof(format, v...)
}

func (ent *entryLogger) Info(v ...interface{}) {
	ent.entry.Info(v...)
}

func (ent *entryLogger) Infoln(v ...interface{}) {
	ent.entry.Infoln(v...)
}

func (ent *entryLogger) Warnf(format string, v ...interface{}) {
	ent.entry.Warnf(format, v...)
}

func (ent *entryLogger) Warn(v ...interface{}) {
	ent.entry.Warn(v...)
}

func (ent *entryLogger) Warnln(v ...interface{}) {
	ent.entry.Warnln(v...)
}

func (ent *entryLogger) Errorf(format string, v ...interface{}) {
	ent.entry.Errorf(format, v...)
}

func (ent *entryLogger) Error(v ...interface{}) {
	ent.entry.Error(v...)
}

func (ent *entryLogger) Errorln(v ...interface{}) {
	ent.entry.Errorln(v...)
}

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

func (ent *globalLogger) Debugln(v ...interface{}) {
	log.Debugln(v...)
}

func (ent *globalLogger) Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func (ent *globalLogger) Info(v ...interface{}) {
	log.Info(v...)
}

func (ent *globalLogger) Infoln(v ...interface{}) {
	log.Infoln(v...)
}

func (ent *globalLogger) Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

func (ent *globalLogger) Warn(v ...interface{}) {
	log.Warn(v...)
}

func (ent *globalLogger) Warnln(v ...interface{}) {
	log.Warnln(v...)
}

func (ent *globalLogger) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

func (ent *globalLogger) Error(v ...interface{}) {
	log.Error(v...)
}

func (ent *globalLogger) Errorln(v ...interface{}) {
	log.Errorln(v...)
}
