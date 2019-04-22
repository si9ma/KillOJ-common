package log

import (
	"context"

	"go.uber.org/zap/zapcore"

	"github.com/opentracing/opentracing-go"

	"go.uber.org/zap"
)

// logger Factory
// wrap zap logger, provide normal logger and span logger
type Factory struct {
	logger *zap.Logger
}

func NewFactory(logger *zap.Logger) Factory {
	return Factory{logger: logger}
}

// get the normal Logger without write span log
func (f Factory) Bg() Logger {
	return logger{logger: f.logger}
}

// if the context contains an span,
// return a spanLogger
func (f Factory) For(ctx context.Context) Logger {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		return spanLogger{span: span, logger: f.logger}
	}

	return f.Bg()
}

// create a Factory with child logger
func (f Factory) With(fields ...zapcore.Field) Factory {
	return Factory{logger: f.logger.With(fields...)}
}
