package tracing

import (
    "context"
    "log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/attribute"
)

func InitTracer() (*trace.TracerProvider, error) {
    exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
    if err != nil {
        log.Fatal(err)
    }

    res, err := resource.New(
        context.Background(),
        resource.WithAttributes(
            attribute.String("service.name", "web-backend"),
        ),
    )
    if err != nil {
        return nil, err
    }

    tracerProvider := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(res),
    )

    // Set the global trace provider
    otel.SetTracerProvider(tracerProvider)
    return tracerProvider, nil
}
