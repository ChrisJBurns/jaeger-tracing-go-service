package tracer

import (
	"log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
)

func InitJaeger() func() {
	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "jaeger-tracing-go-service",
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.Float64("float", 312.23),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
	}
}