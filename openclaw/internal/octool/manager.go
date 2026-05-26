//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package octool

import (
	"context"
	"io"
	"os/exec"
	"sync"
	"time"
)

const (
	defaultYieldMs         = 10_000
	defaultTimeoutS        = 1_800
	defaultLogTail         = 40
	defaultLogLimit        = 200
	defaultMaxLines        = 20_000
	defaultJobTTL          = 30 * time.Minute
	defaultKillGrace       = 2 * time.Second
	defaultIODrain         = 1 * time.Second
	defaultShellEnvTimeout = 5 * time.Second

	shellProgram        = "bash"
	shellLoginFlag      = "-lc"
	shellEnvDumpCommand = "env -0"
)

type Manager struct {
	mu       sync.Mutex
	sessions map[string]*session

	maxLines int
	jobTTL   time.Duration
	baseEnv  map[string]string
	policy   CommandPolicy
	redactor OutputRedactor

	clock func() time.Time

	shellEnvSnapshot func(context.Context, string) map[string]string
}

type Option func(*Manager)

func WithMaxLines(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithJobTTL(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBaseEnv(env map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCommandPolicy(policy CommandPolicy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOutputRedactor(redactor OutputRedactor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewManager(opts ...Option) *Manager { _ = "STUB: not implemented"; return nil }

type execParams struct {
	Command    string
	Workdir    string
	Env        map[string]string
	Pty        bool
	Background bool

	YieldMs  *int
	TimeoutS *int
}

type execResult struct {
	Status     string   `json:"status"`
	Output     string   `json:"output,omitempty"`
	ExitCode   int      `json:"exitCode,omitempty"`
	SessionID  string   `json:"sessionId,omitempty"`
	MediaFiles []string `json:"media_files,omitempty"`
	MediaDirs  []string `json:"media_dirs,omitempty"`
}

func (m *Manager) Exec(
	ctx context.Context,
	params execParams,
) (execResult, error) {
	_ = "STUB: not implemented"
	return *new(execResult), nil
}

func runForeground(
	ctx context.Context,
	params execParams,
	timeout time.Duration,
	baseEnv map[string]string,
) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func shellCmd(ctx context.Context, command string) *exec.Cmd { _ = "STUB: not implemented"; return nil }

func mergedEnv(
	baseEnv map[string]string,
	extra map[string]string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func setEnv(env []string, k, v string) []string { _ = "STUB: not implemented"; return nil }

func exitCode(err error) int { _ = "STUB: not implemented"; return 0 }

func (m *Manager) startBackground(
	params execParams,
	timeout time.Duration,
	redact func(string) string,
) (*session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use cmd.Process.Wait() instead of cmd.Wait() because
// cmd.Wait() closes the pipe read ends returned by StdoutPipe
// and StderrPipe, which races with readFrom goroutines still
// reading from those pipes.  See the exec.StdoutPipe docs:
// "It is thus incorrect to call Wait before all reads from the
// pipe have completed."

func copyEnvMap(env map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func (m *Manager) commandRequest(
	ctx context.Context,
	params execParams,
) CommandRequest {
	_ = "STUB: not implemented"
	return *new(CommandRequest)
}

func (m *Manager) commandEnv(
	ctx context.Context,
	workdir string,
	extra map[string]string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func mergeEnvMaps(
	base map[string]string,
	extra map[string]string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) outputRedactor(
	req CommandRequest,
) func(string) string {
	_ = "STUB: not implemented"
	return nil
}

func copyCommandRequest(req CommandRequest) CommandRequest {
	_ = "STUB: not implemented"
	return *new(CommandRequest)
}

func applyOutputRedactor(
	redact func(string) string,
	output string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func currentProcessEnvMap() map[string]string { _ = "STUB: not implemented"; return nil }

func envListToMap(env []string) map[string]string { _ = "STUB: not implemented"; return nil }

func splitEnvPair(pair string) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (m *Manager) loginShellEnv(
	ctx context.Context,
	workdir string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func snapshotLoginShellEnv(
	ctx context.Context,
	workdir string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func waitDone(done <-chan struct{}, timeout time.Duration) { _ = "STUB: not implemented"; return }

func startPipes(
	cmd *exec.Cmd,
) (io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), *new(io.ReadCloser), *new(io.ReadCloser), nil
}

func (m *Manager) list() []processSession { _ = "STUB: not implemented"; return nil }

// ListSessions returns the current exec_command session snapshots.
func (m *Manager) ListSessions() []ProcessSession { _ = "STUB: not implemented"; return nil }

func (m *Manager) poll(id string, limit *int) (processPoll, error) {
	_ = "STUB: not implemented"
	return *new(processPoll), nil
}

func (m *Manager) log(
	id string,
	offset *int,
	limit *int,
) (processLog, error) {
	_ = "STUB: not implemented"
	return *new(processLog), nil
}

func (m *Manager) write(
	id string,
	data string,
	newline bool,
) (processWrite, error) {
	_ = "STUB: not implemented"
	return *new(processWrite), nil
}

func (m *Manager) kill(id string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) clearFinished(id string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) remove(id string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) get(id string) (*session, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Manager) cleanupExpired() { _ = "STUB: not implemented"; return }
