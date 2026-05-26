//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package taskrun provides tools for controlling background task runs.
package taskrun

import (
	"context"
	"time"

	taskrunruntime "trpc.group/trpc-go/trpc-agent-go/agent/taskrun"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	toolSpawn  = "start_task_run"
	toolList   = "list_task_runs"
	toolGet    = "get_task_run"
	toolCancel = "cancel_task_run"
	toolWait   = "wait_task_run"

	argAgentName      = "agent_name"
	argID             = "id"
	argMode           = "mode"
	argTask           = "task"
	argTimeoutSeconds = "timeout_seconds"
	argWaitSeconds    = "wait_timeout_seconds"

	spawnModeAsync = "async"
	spawnModeSync  = "sync"

	schemaTypeInteger = "integer"
	schemaTypeObject  = "object"
	schemaTypeString  = "string"
)

// Option configures the generated tools.
type Option func(*options)

type options struct {
	defaultAgentName        string
	runtimeState            map[string]any
	injectedContextMessages []model.Message
	allowNested             bool
}

// WithDefaultAgentName configures the agent selected by spawn when the caller
// does not provide one.
func WithDefaultAgentName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRuntimeState merges static runtime state into each spawned run.
func WithRuntimeState(state map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInjectedContextMessages appends non-persisted context messages to each
// spawned run.
func WithInjectedContextMessages(messages []model.Message) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNestedSpawns allows a task run to spawn additional task runs.
func WithNestedSpawns(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// Tools contains all task run control tools.
type Tools struct {
	state *toolState

	spawn  *spawnTool
	list   *listTool
	get    *getTool
	cancel *cancelTool
	wait   *waitTool
}

type toolState struct {
	controller taskrunruntime.Controller
	options    options
}

// NewTools creates task run control tools.
func NewTools(
	controller taskrunruntime.Controller,
	opts ...Option,
) Tools {
	_ = "STUB: not implemented"
	return *new(Tools)
}

// SetController updates the controller used by all tools.
func (t *Tools) SetController(controller taskrunruntime.Controller) {
	_ = "STUB: not implemented"
	return
}

// All returns all tool declarations.
func (t *Tools) All() []tool.Tool { _ = "STUB: not implemented"; return nil }

type spawnTool struct {
	state *toolState
}

type listTool struct {
	state *toolState
}

type getTool struct {
	state *toolState
}

type cancelTool struct {
	state *toolState
}

type waitTool struct {
	state *toolState
}

type spawnInput struct {
	Task               string `json:"task"`
	AgentName          string `json:"agent_name"`
	Mode               string `json:"mode"`
	TimeoutSeconds     int    `json:"timeout_seconds"`
	WaitTimeoutSeconds int    `json:"wait_timeout_seconds"`
}

type runIDInput struct {
	ID             string `json:"id"`
	TimeoutSeconds int    `json:"timeout_seconds"`
}

type listResult struct {
	Runs []taskrunruntime.Run `json:"runs,omitempty"`
}

func (t *spawnTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *spawnTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *listTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *listTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *getTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *getTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *cancelTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *cancelTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *waitTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *waitTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func normalizeSpawnMode(mode string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func waitContext(
	ctx context.Context,
	timeoutSeconds int,
) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func waitTimedOut(
	ctx context.Context,
	waitCtx context.Context,
	err error,
	timeoutSeconds int,
) bool {
	_ = "STUB: not implemented"
	return false
}

func runIDSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func waitSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func requireTools(state *toolState) (*toolState, error) { _ = "STUB: not implemented"; return nil, nil }

func validateEmptyArgs(args []byte) error { _ = "STUB: not implemented"; return nil }

func decodeRunIDArgs(
	ctx context.Context,
	args []byte,
) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func decodeWaitArgs(
	ctx context.Context,
	args []byte,
) (runIDInput, string, error) {
	_ = "STUB: not implemented"
	return *new(runIDInput), "", nil
}

func currentContext(ctx context.Context) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func isNestedTaskRun(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func sameOwner(run *taskrunruntime.Run, userID string) bool {
	_ = "STUB: not implemented"
	return false
}

func secondsDuration(seconds int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func cloneRuntimeState(state map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }
