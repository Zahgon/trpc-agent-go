//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package metric provides metrics collection functionality for the trpc-agent-go framework.
// It integrates with OpenTelemetry to provide comprehensive metrics capabilities.
package metric

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

// InitMeterProvider initializes the meter provider and default meters.
func InitMeterProvider(mp metric.MeterProvider) error { _ = "STUB: not implemented"; return nil }

// GetMeterProvider returns the meter provider.
func GetMeterProvider() metric.MeterProvider {
	_ = "STUB: not implemented"
	return *new(metric.MeterProvider)
}

// SetHistogramBuckets updates bucket boundaries for a specific histogram metric.
// The metricName should be one of the defined metric names in the metrics package.
// Note: This creates a new histogram instrument; old data is not migrated.
func SetHistogramBuckets(meterName string, metricName string, boundaries []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func setChatHistogramBuckets(metricName string, boundaries []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func setExecuteToolHistogramBuckets(metricName string, boundaries []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func setInvokeAgentHistogramBuckets(metricName string, boundaries []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func setWorkflowHistogramBuckets(metricName string, boundaries []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func initInvokeAgentMetrics(mp metric.MeterProvider) error { _ = "STUB: not implemented"; return nil }

func initWorkflowMetrics(mp metric.MeterProvider) error { _ = "STUB: not implemented"; return nil }

// NewMeterProvider creates a new meter provider with optional configuration.
// The environment variables described below can be used for Endpoint configuration.
// OTEL_EXPORTER_OTLP_ENDPOINT, OTEL_EXPORTER_OTLP_METRICS_ENDPOINT (default: "https://localhost:4317")
// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc
func NewMeterProvider(ctx context.Context, opts ...Option) (*sdkmetric.MeterProvider, error) {
	_ = "STUB: not implemented"
	// Set default options
	return nil, nil
}

// Default to gRPC

// Set endpoint based on protocol if not explicitly set

func metricsEndpoint(protocol string) string { _ = "STUB: not implemented"; return "" }

// Return different default endpoints based on protocol

// HTTP endpoint base URL (otlpmetrichttp will add /v1/metrics automatically)

// gRPC endpoint (host:port)

// Initializes an OTLP HTTP exporter, and configures the corresponding meter provider.
func newHTTPMeterProvider(ctx context.Context, res *resource.Resource, endpoint string) (*sdkmetric.MeterProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initializes an OTLP gRPC exporter, and configures the corresponding meter provider.
func newGRPCMeterProvider(ctx context.Context, res *resource.Resource, endpoint string) (*sdkmetric.MeterProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Option is a function that configures meter options.
type Option func(*options)

// options holds the configuration options for meter.
type options struct {
	metricsEndpoint    string
	serviceName        string
	serviceVersion     string
	serviceNamespace   string
	protocol           string // Protocol to use (grpc or http)
	resourceAttributes *[]attribute.KeyValue
}

// WithEndpoint sets the metrics endpoint(host and port) the Exporter will connect to.
// The provided endpoint should resemble "example.com:4317" (no scheme or path).
// If the OTEL_EXPORTER_OTLP_ENDPOINT or OTEL_EXPORTER_OTLP_METRICS_ENDPOINT environment variable is set,
// and this option is not passed, that variable value will be used.
// If both environment variables are set, OTEL_EXPORTER_OTLP_METRICS_ENDPOINT will take precedence.
// If an environment variable is set, and this option is passed, this option will take precedence.
func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithProtocol sets the protocol to use for metrics export.
// Supported protocols are "grpc" (default) and "http".
func WithProtocol(protocol string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithServiceName overrides the service.name resource attribute.
func WithServiceName(serviceName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithServiceNamespace overrides the service.namespace resource attribute.
func WithServiceNamespace(serviceNamespace string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithServiceVersion overrides the service.version resource attribute.
func WithServiceVersion(serviceVersion string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithResourceAttributes appends custom resource attributes.
func WithResourceAttributes(attrs ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func buildResource(ctx context.Context, options *options) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	// Build resource with options values
	return nil, nil
}

// Adds host.name
// Adds telemetry.sdk.{name,language,version}

// Append custom resource attributes
