//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package recorder

import (
	"context"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
)

// Recorder records runner event streams into evalset assets via an injected evalset.Manager.
type Recorder struct {
	manager            evalset.Manager
	name               string
	asyncWriteEnabled  bool
	writeTimeout       time.Duration
	evalSetIDResolver  EvalSetIDResolver
	evalCaseIDResolver EvalCaseIDResolver
	traceModeEnabled   bool
	accumulators       sync.Map // It maps request IDs to in-progress accumulators.
	locker             *keyedLocker
	writeMu            sync.Mutex
	closed             bool
	writesWg           sync.WaitGroup
}

var _ plugin.Plugin = (*Recorder)(nil)
var _ plugin.Closer = (*Recorder)(nil)

// New creates a Recorder plugin.
func New(manager evalset.Manager, opts ...Option) (*Recorder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Name implements plugin.Plugin.
func (r *Recorder) Name() string {
	_ = "STUB: not implemented"

	// Register implements plugin.Plugin.
	return ""
}

func (r *Recorder) Register(reg *plugin.Registry) { _ = "STUB: not implemented"; return }

// Close waits for in-flight async writes to finish.
func (r *Recorder) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *Recorder) onEvent(
	ctx context.Context,
	inv *agent.Invocation,
	e *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Recorder) handleRunError(
	ctx context.Context,
	inv *agent.Invocation,
	requestID string,
	acc *accumulator,
	e *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (r *Recorder) handleResponseEvent(acc *accumulator, rsp *model.Response) {
	_ = "STUB: not implemented"
	return
}

func (r *Recorder) handleRunCompletion(
	ctx context.Context,
	inv *agent.Invocation,
	requestID string,
	acc *accumulator,
	completionTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

func (r *Recorder) buildTurn(
	ctx context.Context,
	inv *agent.Invocation,
	requestID string,
	snapshot turnSnapshot,
	isError bool,
	createdAt time.Time,
) (*turnToPersist, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildFinalResponse(snapshot turnSnapshot, isError bool) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatRunError(err model.ResponseError) string { _ = "STUB: not implemented"; return "" }

func (r *Recorder) resolveEvalSetID(ctx context.Context, inv *agent.Invocation) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *Recorder) resolveEvalCaseID(ctx context.Context, inv *agent.Invocation) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *Recorder) startWrite(ctx context.Context, turn *turnToPersist) {
	_ = "STUB: not implemented"
	return
}

func extractAssistantContentMessage(rsp *model.Response) (model.Message, bool) {
	_ = "STUB: not implemented"
	return *new(model.Message), false
}
