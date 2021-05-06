package trace

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	exporttrace "go.opentelemetry.io/otel/sdk/export/trace"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type BaseResourceEventHandler struct {
	Store        *Store
	SpanExporter exporttrace.SpanExporter
	Logger       *logrus.Logger
}

func (h BaseResourceEventHandler) TracerProviderFor(attrs ...attribute.KeyValue) trace.TracerProvider {
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(h.SpanExporter), /*
			sdktrace.WithMaxQueueSize(20),
			sdktrace.WithBatchTimeout(1000),
			sdktrace.WithMaxExportBatchSize(512),
		*/
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(sdkresource.NewWithAttributes(attrs...)),
	)
}
