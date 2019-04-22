package tracing

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/si9ma/KillOJ-common/constants"
	"github.com/si9ma/KillOJ-common/log"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
)

// Create a new instance of Jaeger tracer
func NewTracer(serverName string) opentracing.Tracer {
	// todo why?
	metricsFactory := prometheus.New().Namespace(metrics.NSOptions{Name: constants.ProjectName, Tags: nil})
	cfg, err := config.FromEnv() // init config from environment variable
	if err != nil {
		log.Bg().Fatal("parse Jaeger environment variable fail", zap.Error(err))
	}
	cfg.ServiceName = serverName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1

	jaegerLogger := jaegerLoggerAdapter{log.Bg()}

	metricsFactory = metricsFactory.Namespace(metrics.NSOptions{Name: serverName, Tags: nil})
	tracer, _, err := cfg.NewTracer(
		config.Logger(jaegerLogger),
		config.Metrics(metricsFactory),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		log.Bg().Fatal("initialize Jaeger Tracer fail", zap.Error(err))
	}
	return tracer
}

// adapt our logger to jaeger logger
type jaegerLoggerAdapter struct {
	logger log.Logger
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(msg)
}

func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(msg, args...))
}
