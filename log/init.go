package log

import (
	"context"
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/si9ma/KillOJ-common/utils"

	"github.com/opentracing/opentracing-go"
	"github.com/si9ma/KillOJ-common/constants"

	"go.uber.org/zap"
)

var zapLogger *zap.Logger

// auto init
func init() {
	zapLogger, _ = zap.NewProduction()
}

// manual init
func Init(outputPaths []string) (err error) {
	// write to stdout when output path is empty
	if len(outputPaths) == 0 {
		outputPaths = []string{"stdout"}
	}

	// on debug mode, write log to stdout at the same time
	if e := os.Getenv(constants.EnvDebug); e != "" {
		debug := false
		if debug, err = strconv.ParseBool(e); err != nil {
			return errors.Wrapf(err, "parse env var %s=%s fail", constants.EnvDebug, e)
		}

		if debug && !utils.ContainsString(outputPaths, "stdout") {
			outputPaths = append(outputPaths, "stdout")
		}
	}

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
