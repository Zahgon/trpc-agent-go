//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package processor provides request and response processing functionality.
package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// BasicRequestProcessor implements the basic request processing logic.
type BasicRequestProcessor struct {
	// GenerationConfig contains the default generation configuration.
	GenerationConfig model.GenerationConfig
}

// NewBasicRequestProcessor creates a new basic request processor with default settings.
func NewBasicRequestProcessor(opts ...BasicOption) *BasicRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// Apply options.

// BasicOption is a functional option for configuring the BasicRequestProcessor.
type BasicOption func(*BasicRequestProcessor)

// WithGenerationConfig sets the default generation configuration.
func WithGenerationConfig(config model.GenerationConfig) BasicOption {
	_ = "STUB: not implemented"
	return *new(BasicOption)
}

// ProcessRequest implements the flow.RequestProcessor interface.
// It handles setting generation parameters.
func (p *BasicRequestProcessor) ProcessRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Set generation configuration.

// Propagate structured output from invocation to request if present.

// Send a preprocessing event.
