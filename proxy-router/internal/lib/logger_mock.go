package lib

import (
	"go.uber.org/zap/zapcore"
)

type LoggerMock struct{}

func (l *LoggerMock) Debug(args ...interface{})  {}
func (l *LoggerMock) Info(args ...interface{})   {}
func (l *LoggerMock) Warn(args ...interface{})   {}
func (l *LoggerMock) Error(args ...interface{})  {}
func (l *LoggerMock) DPanic(args ...interface{}) {}
func (l *LoggerMock) Panic(args ...interface{})  {}
func (l *LoggerMock) Fatal(args ...interface{})  {}

func (l *LoggerMock) Debugf(template string, args ...interface{})  {}
func (l *LoggerMock) Infof(template string, args ...interface{})   {}
func (l *LoggerMock) Warnf(template string, args ...interface{})   {}
func (l *LoggerMock) Errorf(template string, args ...interface{})  {}
func (l *LoggerMock) DPanicf(template string, args ...interface{}) {}
func (l *LoggerMock) Panicf(template string, args ...interface{})  {}
func (l *LoggerMock) Fatalf(template string, args ...interface{})  {}

func (l *LoggerMock) Debugw(template string, args ...interface{})  {}
func (l *LoggerMock) Infow(template string, args ...interface{})   {}
func (l *LoggerMock) Warnw(template string, args ...interface{})   {}
func (l *LoggerMock) Errorw(template string, args ...interface{})  {}
func (l *LoggerMock) DPanicw(template string, args ...interface{}) {}
func (l *LoggerMock) Panicw(template string, args ...interface{})  {}
func (l *LoggerMock) Fatalw(template string, args ...interface{})  {}

func (l *LoggerMock) Log(lvl zapcore.Level, msg string, fields ...zapcore.Field) {}

func (l *LoggerMock) Sync() error                      { return nil }
func (l *LoggerMock) Close() error                     { return nil }
func (l *LoggerMock) Named(n string) ILogger           { return l }
func (l *LoggerMock) With(args ...interface{}) ILogger { return l }
