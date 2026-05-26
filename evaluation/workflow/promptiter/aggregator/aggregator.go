//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package aggregator normalizes and aggregates per-surface gradients before optimization.
package aggregator

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Aggregator merges sample-level gradients into surface-level aggregate signals.
type Aggregator interface {
	// Aggregate computes one normalized surface gradient from one request.
	Aggregate(ctx context.Context, request *Request) (*Result, error)
}

// Request describes all information needed to aggregate one surface.
type Request struct {
	// SurfaceID is the surface for which aggregation is executed.
	SurfaceID string
	// NodeID binds the request to the surface owner in the snapshot.
	NodeID string
	// Type ensures aggregator logic uses correct semantics per surface type.
	Type astructure.SurfaceType
	// Gradients contains gradients from all samples contributing to this surface.
	Gradients []promptiter.SurfaceGradient
}

// Result carries a single surface-level aggregated gradient.
type Result struct {
	// Gradient is the normalized result that can be optimized by next stage.
	Gradient *promptiter.AggregatedSurfaceGradient
}

type aggregatedGradientProposal struct {
	Gradients []gradientProposal
}

type gradientProposal struct {
	Severity promptiter.LossSeverity
	Gradient string
}

// aggregator is the default Aggregator implementation used by engine.
type aggregator struct {
	// runner executes model-assisted aggregation workflows when required.
	runner runner.Runner
	// runOptions are forwarded to the runner on every aggregation request.
	runOptions []agent.RunOption
	// messageBuilder encodes one request into the runner input message.
	messageBuilder MessageBuilder
	// userIDSupplier provides the request-scoped runner user ID.
	userIDSupplier UserIDSupplier
	// sessionIDSupplier provides the request-scoped runner session ID.
	sessionIDSupplier SessionIDSupplier
}

// New creates an Aggregator implementation bound to the provided runner.
func New(ctx context.Context, runner runner.Runner, opt ...Option) (Aggregator, error) {
	_ = "STUB: not implemented"
	return *new(Aggregator), nil
}

// Aggregate runs surface aggregation and returns the merged gradient output.
func (a *aggregator) Aggregate(ctx context.Context, request *Request) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeRequest(request *Request) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeGradients(surfaceID string, gradients []promptiter.SurfaceGradient) ([]promptiter.SurfaceGradient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compareGradients(left promptiter.SurfaceGradient, right promptiter.SurfaceGradient) int {
	_ = "STUB: not implemented"
	return 0
}

func sanitizeAggregatedGradientProposal(
	request *Request,
	proposal *aggregatedGradientProposal,
) (*promptiter.AggregatedSurfaceGradient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
