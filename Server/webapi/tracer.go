package webapi

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// createTracer creates a new instance of Jaeger tracer.
func createTracer(serviceName string) opentracing.Tracer {
	cfg, err := config.FromEnv()
	failOnErr("%v", err)
	cfg.ServiceName = serviceName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1

	// a quick hack to ensure random generators get different seeds, which are based on current time.
	time.Sleep(100 * time.Millisecond)
	tracer, _, err := cfg.NewTracer()
	failOnErr("%v", err)
	return tracer
}
