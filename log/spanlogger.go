package log

import (
	"github.com/opentracing/opentracing-go"
	tag "github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// warp span logger and zap logger
// write span log and zap log simultaneously
type spanLogger struct {
	logger *zap.Logger
	span   opentracing.Span
}

func (sl spanLogger) Info(msg string, fields ...zapcore.Field) {
	sl.logToSpan("INFO", msg, fields...)
	sl.logger.Info(msg, fields...)
}

func (sl spanLogger) Error(msg string, fields ...zapcore.Field) {
	sl.logToSpan("ERROR", msg, fields...)
	sl.logger.Info(msg, fields...)
}

func (sl spanLogger) Fatal(msg string, fields ...zapcore.Field) {
	sl.logToSpan("FATAL", msg, fields...)
	tag.Error.Set(sl.span, true) // set error tag
	sl.logger.Info(msg, fields...)
}

func (sl spanLogger) With(fields ...zapcore.Field) Logger {
	return spanLogger{logger: sl.logger.With(fields...), span: sl.span}
}

// write span log
func (sl spanLogger) logToSpan(level, msg string, fields ...zapcore.Field) {
	fa := fieldAdapter(make([]log.Field, 0, 2+len(fields)))
	fa = append(fa, log.String("event", msg))
	fa = append(fa, log.String("level", level))
	for _, field := range fields {
		field.AddTo(&fa)
	}
	sl.span.LogFields(fa...)
}
