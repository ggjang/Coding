package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	// f, err := os.Create("traces.txt")
	// if err != nil {
	// 	l.Fatal(err)
	// }
	// defer f.Close()

	// exp, err := newExporter(f)
	// if err != nil {
	// 	l.Fatal(err)
	// }

	// tp := trace.NewTracerProvider(
	// 	trace.WithBatcher(exp),
	// 	trace.WithResource(newResource()),
	// )

	// defer func() {
	// 	if err := tp.Shutdown(context.Background()); err != nil {
	// 		l.Fatal(err)
	// 	}
	// }()

	// otel.SetTracerProvider(tp)

	// sigCh := make(chan os.Signal, 1)
	// signal.Notify(sigCh, os.Interrupt)

	// errCh := make(chan error)
	// app := NewApp(os.Stdin, l)
	// go func() {
	// 	errCh <- app.Run(context.Background())
	// }()

	// select {
	// case <-sigCh:
	// 	l.Println("\ngoodbye")
	// 	return
	// case err := <-errCh:
	// 	if err != nil {
	// 		l.Fatal(err)
	// 	}
	// }
	// this block export to file

	tp, err := tracerProvider("http://192.168.91.134:14268/api/traces")
	if err != nil {
		l.Fatal(err)
	}
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}(ctx)

	tr := tp.Tracer("component-main")

	ctx, span := tr.Start(ctx, "main")
	defer span.End()

	bar(ctx)

}

func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demon
		stdouttrace.WithoutTimestamps(),
	)
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("fib"),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

func tracerProvider(url string) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resurce.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("fib"),
			attribute.String("envirionment", "demo"),
			attribute.Int64("ID", 1),
		)),
	)
	return tp, nil
}

func bar(ctx context.Context) {
	// Use the global TracerProvider.
	tr := otel.Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	// Do bar...
}
