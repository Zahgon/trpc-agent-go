//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package optimizer transforms aggregated gradients into patch suggestions for the target prompt.
package optimizer

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Optimizer converts aggregated gradients into patch updates for one surface.
type Optimizer interface {
	// Optimize generates one patch proposal for one surface request.
	Optimize(ctx context.Context, request *Request) (*Result, error)
}

// Request carries gradient and baseline surface context for optimization.
type Request struct {
	// Surface is the source surface baseline that may be changed.
	Surface *astructure.Surface
	// Gradient is the merged signal that drives optimization decisions.
	Gradient *promptiter.AggregatedSurfaceGradient
}

// Result carries the patch suggestion for one optimized surface.
type Result struct {
	// Patch is the proposed change for the requested surface.
	Patch *promptiter.SurfacePatch
}

type surfacePatchProposal struct {
	Value  astructure.SurfaceValue
	Reason string
}

// optimizer is the default Optimizer implementation used by the engine.
type optimizer struct {
	// runner executes external inference needed to draft patch proposals.
	runner runner.Runner
	// runOptions are forwarded to the runner on every optimization request.
	runOptions []agent.RunOption
	// messageBuilder encodes one request into the runner input message.
	messageBuilder MessageBuilder
	// userIDSupplier provides the request-scoped runner user ID.
	userIDSupplier UserIDSupplier
	// sessionIDSupplier provides the request-scoped runner session ID.
	sessionIDSupplier SessionIDSupplier
}

// New creates an Optimizer instance bound to the provided runner.
func New(ctx context.Context, runner runner.Runner, opt ...Option) (Optimizer, error) {
	_ = "STUB: not implemented"
	return *new(Optimizer), nil
}

// Optimize runs optimization logic and returns one patch proposal.
func (o *optimizer) Optimize(ctx context.Context, request *Request) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeRequest(request *Request) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sanitizePatchProposal(request *Request, proposal *surfacePatchProposal) (*promptiter.SurfacePatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
