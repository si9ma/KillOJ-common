package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

// wrap zap logger
type logger struct {
	logger *zap.Logger
}

func (l logger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l logger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

func (l logger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}
