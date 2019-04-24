package asyncjob

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/si9ma/KillOJ-common/constants"

	"github.com/RichardKnop/machinery/v1"

	"github.com/si9ma/KillOJ-common/log"
	"go.uber.org/zap"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/google/uuid"

	"github.com/opentracing/opentracing-go"
	"github.com/si9ma/KillOJ-common/tracing"
)

func TestConfig_String(t *testing.T) {
	cfg := defaultConfig()
	cfg.Broker = "amqp://si9ma:rabbitmq@localhost:5672/"
	if _, err := Init(cfg); err != nil {
		t.Fatal("Init server fail", err)
	}
}

func TestSender(t *testing.T) {
	cfg := defaultConfig()
	cfg.Broker = "amqp://si9ma:rabbitmq@localhost:5672/"
	cfg.BindingKey = strings.Join([]string{constants.ProjectName, "Judge", "Queue"}, "_")
	var server *machinery.Server
	var err error
	if server, err = Init(cfg); err != nil {
		t.Fatal("Init server fail", err)
	}

	tracer, closer := tracing.NewTracer("test_sender")
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	for i := 0; i < 100; i++ {
		span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")

		batchID := uuid.New().String()
		span.SetBaggageItem("batch.id", batchID)
		log.For(ctx).Info("batch", zap.String("batch.id", batchID))

		judgeTask := tasks.Signature{
			Name: "judge",
		}

		_, err := server.SendTaskWithContext(ctx, &judgeTask)
		if err != nil {
			t.Fatal("send task fail", err)
		}
		span.Finish()
	}
}

func TestWorker(t *testing.T) {
	cfg := defaultConfig()
	cfg.Broker = "amqp://si9ma:rabbitmq@localhost:5672/"
	var server *machinery.Server
	var err error
	if server, err = Init(cfg); err != nil {
		t.Fatal("Init server fail", err)
	}
	tracer, closer := tracing.NewTracer("test_worker")
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	if err := server.RegisterTask("test", func() error {
		fmt.Println("I a test worker!")
		return nil
	}); err != nil {
		t.Fatal()
	}

	if err := server.NewWorker("test", 1).Launch(); err != nil {
		t.Fatal("launch worker fail", err)
	}
}
