//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package backwarder computes backward propagation outputs from trace and gradient data.
package backwarder

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Backwarder computes gradient attribution for one trace step.
type Backwarder interface {
	// Backward derives surface gradients and upstream propagation packets.
	Backward(ctx context.Context, request *Request) (*Result, error)
}

// Request carries a single step trace context for local backwarding.
type Request struct {
	// EvalSetID identifies the evaluation set for trace correlation.
	EvalSetID string
	// EvalCaseID identifies the evaluation case for trace correlation.
	EvalCaseID string
	// Node is the static node that produced this step.
	Node *astructure.Node
	// StepID is the local step identifier in the trace.
	StepID string
	// Input is the concrete input snapshot captured for this step.
	Input *atrace.Snapshot
	// Output is the concrete output snapshot captured for this step.
	Output *atrace.Snapshot
	// Error records runtime failure details when the step did not complete.
	Error string
	// Surfaces are all surfaces whose values affected this step.
	Surfaces []astructure.Surface
	// AllowedGradientSurfaceIDs limits which surfaces may appear in Gradients.
	AllowedGradientSurfaceIDs []string
	// Predecessors stores direct predecessor execution steps.
	Predecessors []Predecessor
	// Incoming carries raw gradients that need to be processed at this step.
	Incoming []GradientPacket
}

// Predecessor captures one direct upstream step for gradient propagation.
type Predecessor struct {
	// StepID identifies the predecessor step id.
	StepID string
	// NodeID identifies the predecessor node id.
	NodeID string
	// Output stores the predecessor output snapshot used by this step.
	Output *atrace.Snapshot
	// Error records predecessor execution error for debugging.
	Error string
}

// GradientPacket carries one scalar gradient unit passed from downstream.
type GradientPacket struct {
	// FromStepID identifies the direct downstream step origin.
	FromStepID string
	// Severity carries failure importance for weighted propagation.
	Severity promptiter.LossSeverity
	// Gradient is the serialized propagated gradient string.
	Gradient string
}

// Result contains gradients attributed to current step and upstream propagation data.
type Result struct {
	// Gradients stores gradients mapped to surfaces affected by this step.
	Gradients []promptiter.SurfaceGradient
	// Upstream carries gradients that still need to propagate to predecessors.
	Upstream []Propagation
}

// Propagation groups packets to be sent to one predecessor step.
type Propagation struct {
	// PredecessorStepID identifies the target upstream step.
	PredecessorStepID string
	// Gradients stores packets forwarded to that predecessor.
	Gradients []GradientPacket
}

// backwarder is the default Backwarder implementation used by the engine.
type backwarder struct {
	// runner executes any external inference used during backward computation.
	runner runner.Runner
	// runOptions are forwarded to the runner on every backward request.
	runOptions []agent.RunOption
	// messageBuilder encodes one request into the runner input message.
	messageBuilder MessageBuilder
	// userIDSupplier provides the request-scoped runner user ID.
	userIDSupplier UserIDSupplier
	// sessionIDSupplier provides the request-scoped runner session ID.
	sessionIDSupplier SessionIDSupplier
}

// New creates a Backwarder instance with injected runner and options.
func New(ctx context.Context, runner runner.Runner, opt ...Option) (Backwarder, error) {
	_ = "STUB: not implemented"
	return *new(Backwarder), nil
}

// Backward computes local gradients and upstream propagation paths for one step.
func (b *backwarder) Backward(ctx context.Context, request *Request) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isNoOpBackwardRequest(request *Request) bool { _ = "STUB: not implemented"; return false }

func backwardStructuredOutput(request *Request) agent.RunOption {
	_ = "STUB: not implemented"
	return *new(agent.RunOption)
}

func backwardResultSchema(request *Request) map[string]any { _ = "STUB: not implemented"; return nil }

func requestSurfaceIDs(request *Request) []string { _ = "STUB: not implemented"; return nil }

func requestAllowedGradientSurfaceIDs(request *Request) []string {
	_ = "STUB: not implemented"
	return nil
}

func requestPredecessorStepIDs(request *Request) []string { _ = "STUB: not implemented"; return nil }

func backwardGradientArraySchema(surfaceIDs []string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func backwardPropagationArraySchema(predecessorStepIDs []string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func normalizeRequest(request *Request) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sanitizeBackwardResult(request *Request, result *Result) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sanitizeSurfaceGradient(
	request *Request,
	surfaceIndex map[string]astructure.Surface,
	gradient promptiter.SurfaceGradient,
) (promptiter.SurfaceGradient, bool, error) {
	_ = "STUB: not implemented"
	return *new(promptiter.SurfaceGradient), false, nil
}

func sanitizePropagation(
	request *Request,
	predecessorIndex map[string]Predecessor,
	propagation Propagation,
) (Propagation, bool, error) {
	_ = "STUB: not implemented"
	return *new(Propagation), false, nil
}

func buildPredecessorIndex(predecessors []Predecessor) (map[string]Predecessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeAllowedGradientSurfaceIDs(
	request *Request,
	surfaceIndex map[string]astructure.Surface,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildAllowedGradientSurfaceIndex(
	request *Request,
	surfaceIndex map[string]astructure.Surface,
) (map[string]astructure.Surface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sanitizeGradientSurfaceID(
	surfaceIndex map[string]astructure.Surface,
	surfaceID string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func sanitizePropagationPredecessorStepID(
	predecessorIndex map[string]Predecessor,
	predecessorStepID string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
