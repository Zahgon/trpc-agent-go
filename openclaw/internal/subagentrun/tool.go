//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package subagentrun

import (
	"context"

	openclawsubagent "trpc.group/trpc-go/trpc-agent-go/openclaw/subagent"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	toolSubagentsSpawn  = "subagents_spawn"
	toolSubagentsList   = "subagents_list"
	toolSubagentsGet    = "subagents_get"
	toolSubagentsCancel = "subagents_cancel"
	toolSubagentsWait   = "subagents_wait"

	toolSessionsSpawn  = "sessions_spawn"
	toolSessionsList   = "sessions_list"
	toolSessionsGet    = "sessions_get"
	toolSessionsCancel = "sessions_cancel"

	argID             = "id"
	argMode           = "mode"
	argTask           = "task"
	argTimeoutSeconds = "timeout_seconds"
	argWaitSeconds    = "wait_timeout_seconds"

	spawnModeAsync  = "async"
	spawnModeSync   = "sync"
	spawnModeReview = "review"

	schemaTypeInteger = "integer"
	schemaTypeObject  = "object"
	schemaTypeString  = "string"
)

type Tools struct {
	spawn  *spawnTool
	list   *listTool
	get    *getTool
	cancel *cancelTool
	wait   *waitTool

	spawnAlias  *spawnTool
	listAlias   *listTool
	getAlias    *getTool
	cancelAlias *cancelTool
}

func NewTools(svc *Service) Tools { _ = "STUB: not implemented"; return *new(Tools) }

func (t *Tools) SetService(svc *Service) { _ = "STUB: not implemented"; return }

func (t *Tools) All() []tool.Tool { _ = "STUB: not implemented"; return nil }

type serviceAwareTool interface {
	setService(svc *Service)
}

type spawnTool struct {
	name  string
	alias bool
	svc   *Service
}

type listTool struct {
	name  string
	alias bool
	svc   *Service
}

type getTool struct {
	name  string
	alias bool
	svc   *Service
}

type cancelTool struct {
	name  string
	alias bool
	svc   *Service
}

type waitTool struct {
	name string
	svc  *Service
}

type spawnInput struct {
	Task               string `json:"task"`
	Mode               string `json:"mode"`
	TimeoutSeconds     int    `json:"timeout_seconds"`
	WaitTimeoutSeconds int    `json:"wait_timeout_seconds"`
}

type runIDInput struct {
	ID             string `json:"id"`
	TimeoutSeconds int    `json:"timeout_seconds"`
}

type listResult struct {
	Runs []openclawsubagent.Run `json:"runs,omitempty"`
}

func (t *spawnTool) setService(svc *Service) { _ = "STUB: not implemented"; return }

func (t *listTool) setService(svc *Service) { _ = "STUB: not implemented"; return }

func (t *getTool) setService(svc *Service) { _ = "STUB: not implemented"; return }

func (t *cancelTool) setService(svc *Service) { _ = "STUB: not implemented"; return }

func (t *waitTool) setService(svc *Service) { _ = "STUB: not implemented"; return }

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

func markAwaitingReview(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

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

func currentContext(
	ctx context.Context,
) (string, *session.Session, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func isNestedSubagent(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
