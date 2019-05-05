package wredis

import (
	"context"
	"testing"
	"time"

	"github.com/si9ma/KillOJ-common/tracing"

	"github.com/opentracing/opentracing-go"
	"github.com/si9ma/KillOJ-common/log"

	"github.com/go-redis/redis"
)

func TestInit(t *testing.T) {
	var client *redis.ClusterClient
	var err error

	cfg := Config{
		Addrs:        []string{"redis:6379"},
		DialTimeout:  2000,
		ReadTimeout:  2000,
		WriteTimeout: 2000,
	}
	if client, err = Init(cfg); err != nil {
		t.Fatal("init redis fail", err)
	}

	client.Set("test", "test", time.Minute)
	cmd := client.Get("test")
	if cmd.Err() != nil {
		t.Fatal(cmd.Err())
	}

	if res, err := cmd.Result(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}

// cmd: gorun -m=TEST -e=JAEGER_AGENT_HOST=jaeger -n=prod_net github.com/si9ma/KillOJ-common/wredis TestRedisOpenTracing
func TestRedisOpenTracing(t *testing.T) {
	var client *redis.ClusterClient
	var err error

	cfg := Config{
		Addrs:        []string{"redis:6379"},
		DialTimeout:  2000,
		ReadTimeout:  2000,
		WriteTimeout: 2000,
	}
	if client, err = Init(cfg); err != nil {
		t.Fatal("init redis fail", err)
	}

	// tracing
	tracer, closer := tracing.NewTracer("TestRedisOpenTracing")
	defer closer.Close()

	span := tracer.StartSpan("test")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	defer span.Finish()
	log.For(ctx).Info("info")
	log.For(ctx).Error("error")

	rclient := WrapRedisClient(ctx, client)
	rclient.Set("test", "test", time.Minute)
	rclient.Get("test")
}
