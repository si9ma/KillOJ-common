package log

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func Init(outputPaths []string) (err error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = outputPaths // set output paths
	zapLogger, err = cfg.Build()
	return
}

// get the normal Logger without write span log
func Bg() Logger {
	return logger{logger: zapLogger}
}

// if the context contains an span,
// return a spanLogger
func For(ctx context.Context) Logger {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		return spanLogger{span: span, logger: zapLogger}
	}

	return Bg()
}
