package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"git.inke.cn/BackendPlatform/jaeger-client-go/config"
	"git.inke.cn/BackendPlatform/jaeger-client-go"
	//"github.com/uber/jaeger-client-go/config"
	//"github.com/uber/jaeger-client-go"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "probabilistic",
			Param: 0.1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
