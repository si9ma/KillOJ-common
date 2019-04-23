package tracing

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"

	"github.com/opentracing/opentracing-go"
	"github.com/si9ma/KillOJ-common/log"
)

func TestNewTracer(t *testing.T) {
	tracer, closer := NewTracer("test_new_tracer")
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("test")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	defer span.Finish()
	log.For(ctx).Info("info")
	log.For(ctx).Error("error")
}

func TestServeMux(t *testing.T) {
	tracer, _ := NewTracer("test_mux")
	mux := NewServeMux(tracer)
	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello"))
	}))
	if err := http.ListenAndServe("localhost:8888", mux); err != nil {
		t.Fatal(err.Error())
	}
}

func TestGin(t *testing.T) {
	tracer, closer := NewTracer("test_gin")
	defer closer.Close()

	fn := func(c *gin.Context) {
		span := opentracing.SpanFromContext(c.Request.Context())
		if span == nil {
			t.Error("Span is nil")
		}
		_, _ = c.Writer.Write([]byte("hello gin"))
	}

	r := gin.New()
	r.Use(ginhttp.Middleware(tracer))
	group := r.Group("")
	group.GET("", fn)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Error non-nil %v", err)
	}
	r.ServeHTTP(w, req)
}
