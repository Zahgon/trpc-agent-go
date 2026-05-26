//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package skill

import (
	"context"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/internal/programsession"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultExecYieldMS = programsession.DefaultExecYieldMS
	defaultIOYieldMS   = programsession.DefaultIOYieldMS
	defaultPollLines   = programsession.DefaultPollLines
	defaultSessionTTL  = programsession.DefaultSessionTTL
	defaultSessionKill = programsession.DefaultSessionKill
)

const (
	interactionKindPrompt    = "prompt"
	interactionKindSelection = "selection"
)

type execInput struct {
	runInput
	TTY       bool `json:"tty,omitempty"`
	YieldMS   int  `json:"yield_ms,omitempty"`
	PollLines int  `json:"poll_lines,omitempty"`
}

type sessionWriteInput struct {
	SessionID string `json:"session_id"`
	Chars     string `json:"chars,omitempty"`
	Submit    bool   `json:"submit,omitempty"`
	YieldMS   int    `json:"yield_ms,omitempty"`
	PollLines int    `json:"poll_lines,omitempty"`
}

type sessionPollInput struct {
	SessionID string `json:"session_id"`
	YieldMS   int    `json:"yield_ms,omitempty"`
	PollLines int    `json:"poll_lines,omitempty"`
}

type sessionKillInput struct {
	SessionID string `json:"session_id"`
}

type sessionInteraction struct {
	NeedsInput bool   `json:"needs_input"`
	Kind       string `json:"kind,omitempty"`
	Hint       string `json:"hint,omitempty"`
}

type execOutput struct {
	Status      string              `json:"status"`
	SessionID   string              `json:"session_id"`
	Output      string              `json:"output,omitempty"`
	Offset      int                 `json:"offset"`
	NextOffset  int                 `json:"next_offset"`
	ExitCode    *int                `json:"exit_code,omitempty"`
	Interaction *sessionInteraction `json:"interaction,omitempty"`
	Result      *runOutput          `json:"result,omitempty"`
}

type sessionKillOutput struct {
	OK        bool   `json:"ok"`
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
}

type execSession struct {
	mu sync.Mutex

	proc codeexecutor.ProgramSession
	eng  codeexecutor.Engine
	ws   codeexecutor.Workspace

	in                    runInput
	staged                []stagedInput
	stageWarnings         []string
	saveRequested         bool
	outputsSaveSkipReason string
	final                 *runOutput
	exitedAt              time.Time
	finalized             bool
	finalizedAt           time.Time
}

// ExecTool starts interactive commands inside skill workspaces.
type ExecTool struct {
	run *RunTool

	mu       sync.Mutex
	sessions map[string]*execSession
	ttl      time.Duration
	clock    func() time.Time
}

// WriteStdinTool writes stdin to running skill sessions.
type WriteStdinTool struct {
	exec *ExecTool
}

// PollSessionTool polls running skill sessions for new output.
type PollSessionTool struct {
	exec *ExecTool
}

// KillSessionTool terminates running skill sessions.
type KillSessionTool struct {
	exec *ExecTool
}

// NewExecTool creates the interactive skill execution tool.
func NewExecTool(run *RunTool) *ExecTool { _ = "STUB: not implemented"; return nil }

// NewWriteStdinTool creates the stdin write tool for skill_exec sessions.
func NewWriteStdinTool(exec *ExecTool) *WriteStdinTool { _ = "STUB: not implemented"; return nil }

// NewPollSessionTool creates the poll tool for skill_exec sessions.
func NewPollSessionTool(exec *ExecTool) *PollSessionTool { _ = "STUB: not implemented"; return nil }

// NewKillSessionTool creates the kill tool for skill_exec sessions.
func NewKillSessionTool(exec *ExecTool) *KillSessionTool { _ = "STUB: not implemented"; return nil }

// Declaration returns the tool schema for skill_exec.
func (t *ExecTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Declaration returns the tool schema for skill_write_stdin.
func (t *WriteStdinTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Declaration returns the tool schema for skill_poll_session.
func (t *PollSessionTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Declaration returns the tool schema for skill_kill_session.
func (t *KillSessionTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// StreamableCall starts an interactive skill session and streams output.
func (t *ExecTool) StreamableCall(
	ctx context.Context,
	args []byte,
) (*tool.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamableCall writes stdin to a running skill session.
func (t *WriteStdinTool) StreamableCall(
	ctx context.Context,
	args []byte,
) (*tool.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamableCall polls a running skill session for new output.
func (t *PollSessionTool) StreamableCall(
	ctx context.Context,
	args []byte,
) (*tool.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call terminates a running skill session and removes it.
func (t *KillSessionTool) Call(
	_ context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *ExecTool) putSession(id string, sess *execSession) { _ = "STUB: not implemented"; return }

func (t *ExecTool) getSession(id string) (*execSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ExecTool) removeSession(id string) (*execSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ExecTool) cleanupExpiredLocked() { _ = "STUB: not implemented"; return }

func (t *ExecTool) buildExecOutput(
	ctx context.Context,
	sessionID string,
	sess *execSession,
	poll codeexecutor.ProgramPoll,
) (execOutput, error) {
	_ = "STUB: not implemented"
	return *new(execOutput), nil
}

func (t *ExecTool) captureFinalResult(
	ctx context.Context,
	sess *execSession,
	poll codeexecutor.ProgramPoll,
) (*runOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sessionRunResult(
	proc codeexecutor.ProgramSession,
	poll codeexecutor.ProgramPoll,
) codeexecutor.RunResult {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult)
}

func buildExecStream(
	output string,
	result execOutput,
) *tool.StreamReader {
	_ = "STUB: not implemented"
	return nil
}

func waitForProgramOutput(
	proc codeexecutor.ProgramSession,
	yield time.Duration,
	limit *int,
) codeexecutor.ProgramPoll {
	_ = "STUB: not implemented"
	return *new(codeexecutor.ProgramPoll)
}

func yieldDuration(ms int, fallback int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func pollLineLimit(lines int) *int { _ = "STUB: not implemented"; return nil }

func detectInteraction(
	poll codeexecutor.ProgramPoll,
) *sessionInteraction {
	_ = "STUB: not implemented"
	return nil
}

func lastNonEmptyLine(text string) string { _ = "STUB: not implemented"; return "" }

func hasSelectionItems(text string) bool { _ = "STUB: not implemented"; return false }

func parseExecArgs(args []byte) (execInput, error) {
	_ = "STUB: not implemented"
	return *new(execInput), nil
}

func parseSessionWriteArgs(
	args []byte,
) (sessionWriteInput, error) {
	_ = "STUB: not implemented"
	return *new(sessionWriteInput), nil
}

func parseSessionPollArgs(
	args []byte,
) (sessionPollInput, error) {
	_ = "STUB: not implemented"
	return *new(sessionPollInput), nil
}

func execOutputSchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func execArtifactsStateDelta(
	toolCallID string,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// StateDelta stores artifact references from a completed session result.
func (t *ExecTool) StateDelta(
	toolCallID string,
	_ []byte,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// StateDelta stores artifact references from a completed session result.
func (t *WriteStdinTool) StateDelta(
	toolCallID string,
	_ []byte,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// StateDelta stores artifact references from a completed session result.
func (t *PollSessionTool) StateDelta(
	toolCallID string,
	_ []byte,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

var _ tool.Tool = (*ExecTool)(nil)
var _ tool.StreamableTool = (*ExecTool)(nil)
var _ stateDeltaProvider = (*ExecTool)(nil)
var _ tool.Tool = (*WriteStdinTool)(nil)
var _ tool.StreamableTool = (*WriteStdinTool)(nil)
var _ stateDeltaProvider = (*WriteStdinTool)(nil)
var _ tool.Tool = (*PollSessionTool)(nil)
var _ tool.StreamableTool = (*PollSessionTool)(nil)
var _ stateDeltaProvider = (*PollSessionTool)(nil)
var _ tool.Tool = (*KillSessionTool)(nil)
var _ tool.CallableTool = (*KillSessionTool)(nil)
