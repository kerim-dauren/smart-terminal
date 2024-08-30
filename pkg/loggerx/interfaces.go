package loggerx

import (
	"go.uber.org/zap"
)

type Zapper interface {
	Logger
	Zap() *zap.Logger
}

type Logger interface {
	Error(message string, fields ...zap.Field)
	Errorf(message string, args ...interface{})
	Warn(message string, fields ...zap.Field)
	Warnf(message string, args ...interface{})
	Fatal(message string, fields ...zap.Field)
	Fatalf(message string, args ...interface{})
	Info(message string, fields ...zap.Field)
	Infof(message string, args ...interface{})
	Debug(message string, fields ...zap.Field)
	Debugf(message string, args ...interface{})
	Panic(message string, fields ...zap.Field)
	Panicf(message string, args ...interface{})
}
