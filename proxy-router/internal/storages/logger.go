package storages

import (
	"strings"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
)

type BadgerLogger struct {
	log lib.ILogger
}

func NewBadgerLogger(log lib.ILogger) *BadgerLogger {
	return &BadgerLogger{
		log: log.Named("BADGER"),
	}
}

func (l *BadgerLogger) Errorf(s string, p ...interface{}) {
	l.log.Errorf(normalize(s), p...)
}
func (l *BadgerLogger) Warningf(s string, p ...interface{}) {
	l.log.Warnf(normalize(s), p...)
}
func (l *BadgerLogger) Infof(s string, p ...interface{}) {
	l.log.Infof(normalize(s), p...)
}
func (l *BadgerLogger) Debugf(s string, p ...interface{}) {
	l.log.Debugf(normalize(s), p...)
}

func normalize(s string) string {
	return strings.TrimRight(s, "\n")
	// trims new line
	// return strings.
}
