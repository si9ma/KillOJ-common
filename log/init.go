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
	"go.uber.org/zap/zapcore"
)

type LogEncoding string

const Json LogEncoding = "json"
const Console LogEncoding = "console"

var zapLogger *zap.Logger

// auto init
func init() {
	cfg := zap.NewProductionConfig()
	cfg.Encoding = string(Json) // default json
	zapLogger, _ = cfg.Build()
}

// manual init
func Init(outputPaths []string, encode LogEncoding) (err error) {
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

	// config format
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
	if encode == Console {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg := zap.NewProductionConfig()
	cfg.Encoding = string(encode)
	cfg.EncoderConfig = encoderCfg
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
