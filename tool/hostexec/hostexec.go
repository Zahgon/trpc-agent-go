//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package hostexec provides direct host command execution tools for
// personal-agent style workflows.
package hostexec

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultBaseDir      = "."
	defaultToolSetName  = "hostexec"
	defaultWriteYieldMS = 200

	toolExecCommand = "exec_command"
	toolWriteStdin  = "write_stdin"
	toolKillSession = "kill_session"
)

const (
	errExecToolNotConfigured  = "exec tool is not configured"
	errWriteToolNotConfigured = "write_stdin tool is not configured"
	errKillToolNotConfigured  = "kill_session tool is not configured"
	errCommandRequired        = "command is required"
	errSessionIDRequired      = "session id is required"
)

type config struct {
	baseDir  string
	name     string
	maxLines int
	jobTTL   time.Duration
	baseEnv  map[string]string
}

// Option configures the hostexec tool set.
type Option func(*config)

// WithBaseDir sets the default directory used by exec_command.
func WithBaseDir(baseDir string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithName sets the tool set name.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxLines sets the maximum retained output lines per session.
func WithMaxLines(lines int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJobTTL sets how long finished sessions are retained.
func WithJobTTL(ttl time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseEnv sets environment variables applied to all commands.
func WithBaseEnv(env map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func defaultConfig() config { _ = "STUB: not implemented"; return *new(config) }

// NewToolSet creates a host command execution tool set.
func NewToolSet(opts ...Option) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

type toolSet struct {
	name    string
	baseDir string
	mgr     *manager
	tools   []tool.Tool
}

func (s *toolSet) Tools(context.Context) []tool.Tool { _ = "STUB: not implemented"; return nil }

func (s *toolSet) Close() error { _ = "STUB: not implemented"; return nil }

func (s *toolSet) Name() string { _ = "STUB: not implemented"; return "" }

type execCommandTool struct {
	mgr     *manager
	baseDir string
}

func (t *execCommandTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type execInput struct {
	Command       string            `json:"command"`
	Workdir       string            `json:"workdir,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	YieldTimeMS   *int              `json:"yield_time_ms,omitempty"`
	YieldMs       *int              `json:"yieldMs,omitempty"`
	Background    bool              `json:"background,omitempty"`
	TimeoutSec    *int              `json:"timeout_sec,omitempty"`
	TimeoutSecOld *int              `json:"timeoutSec,omitempty"`
	TTY           *bool             `json:"tty,omitempty"`
	PTY           *bool             `json:"pty,omitempty"`
}

func (t *execCommandTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type writeStdinTool struct {
	mgr *manager
}

func (t *writeStdinTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type writeInput struct {
	SessionID     string `json:"session_id,omitempty"`
	SessionIDOld  string `json:"sessionId,omitempty"`
	Chars         string `json:"chars,omitempty"`
	YieldTimeMS   *int   `json:"yield_time_ms,omitempty"`
	YieldMs       *int   `json:"yieldMs,omitempty"`
	AppendNewline *bool  `json:"append_newline,omitempty"`
	Submit        *bool  `json:"submit,omitempty"`
}

func (t *writeStdinTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type killSessionTool struct {
	mgr *manager
}

func (t *killSessionTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type killInput struct {
	SessionID    string `json:"session_id,omitempty"`
	SessionIDOld string `json:"sessionId,omitempty"`
}

func (t *killSessionTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func pollOutputSchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func mapExecResult(res execResult) map[string]any { _ = "STUB: not implemented"; return nil }

func mapPollResult(
	sessionID string,
	poll processPoll,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func firstInt(values ...*int) *int { _ = "STUB: not implemented"; return nil }

func firstBool(values ...*bool) bool { _ = "STUB: not implemented"; return false }

func firstNonEmpty(values ...string) string { _ = "STUB: not implemented"; return "" }

func resolveBaseDir(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func resolveWorkdir(
	raw string,
	baseDir string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func cloneEnvMap(env map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

var _ tool.ToolSet = (*toolSet)(nil)
var _ tool.CallableTool = (*execCommandTool)(nil)
var _ tool.CallableTool = (*writeStdinTool)(nil)
var _ tool.CallableTool = (*killSessionTool)(nil)
