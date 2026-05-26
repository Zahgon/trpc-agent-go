//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"go.opentelemetry.io/otel/metric"
)

// initTelemetry initializes OpenTelemetry trace and metric exporters.
func initTelemetry() (metric.Meter, error) {
	_ = "STUB: not implemented"
	// Start trace telemetry.
	return *new(metric.Meter), nil
}

// Start metric telemetry.

// Register cleanup functions.
// Note: In a real application, you would want to handle cleanup more gracefully.

// Wait for the application to finish, then cleanup.

// initMetrics initializes OpenTelemetry metrics.
func (e *toolTimerExample) initMetrics() error {
	_ = "STUB: not implemented"

	// Initialize histograms for duration measurements.
	return nil
}

// Initialize counters for execution counts.
