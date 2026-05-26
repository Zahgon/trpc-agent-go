//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package hostexec

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"sync"
	"time"
)

const (
	defaultYieldMS    = 10_000
	defaultTimeoutS   = 1_800
	defaultLogTail    = 40
	defaultMaxLines   = 20_000
	defaultJobTTL     = 30 * time.Minute
	defaultKillGrace  = 2 * time.Second
	timeoutKillGrace  = time.Duration(0)
	defaultIODrain    = 1 * time.Second
	maxTimeoutSeconds = int64((1<<63)-1) /
		int64(time.Second)
)

var errUnknownSession = errors.New("unknown session id")

type manager struct {
	mu       sync.Mutex
	sessions map[string]*session

	maxLines int
	jobTTL   time.Duration
	baseEnv  map[string]string

	clock func() time.Time
}

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
	Status    string
	Output    string
	ExitCode  *int
	SessionID string
}

func newManager() *manager { _ = "STUB: not implemented"; return nil }

func (m *manager) exec(
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

func timeoutDuration(timeoutS int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func shellCmd(
	_ context.Context,
	command string,
) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nosemgrep: go.lang.security.audit.dangerous-exec-command
// hostexec intentionally executes trusted host commands.

//nolint:gosec

func shellSpec() (string, []string, error) { _ = "STUB: not implemented"; return "", nil, nil }

func mergedEnv(
	baseEnv map[string]string,
	extra map[string]string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func setEnv(env []string, key string, value string) []string { _ = "STUB: not implemented"; return nil }

func exitCode(err error) int { _ = "STUB: not implemented"; return 0 }

func (m *manager) startBackground(
	params execParams,
	timeout time.Duration,
) (*session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func startSession(
	id string,
	params execParams,
	timeout time.Duration,
	baseEnv map[string]string,
	maxLines int,
) (*session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use cmd.Process.Wait() instead of cmd.Wait() because
// cmd.Wait() closes StdoutPipe/StderrPipe readers before
// the readFrom goroutines are done consuming those pipes.

func waitDone(
	done <-chan struct{},
	timeout time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

func startPipes(
	cmd *exec.Cmd,
) (io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), *new(io.ReadCloser), *new(io.ReadCloser), nil
}

func (m *manager) poll(
	id string,
	limit *int,
) (processPoll, error) {
	_ = "STUB: not implemented"
	return *new(processPoll), nil
}

func (m *manager) write(
	id string,
	data string,
	newline bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) kill(id string) error { _ = "STUB: not implemented"; return nil }

func (m *manager) killContext(
	ctx context.Context,
	id string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) clearFinished(id string) error { _ = "STUB: not implemented"; return nil }

func (m *manager) get(id string) (*session, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *manager) cleanupExpired() { _ = "STUB: not implemented"; return }

func (m *manager) close() error { _ = "STUB: not implemented"; return nil }

func intPtr(value int) *int { _ = "STUB: not implemented"; return nil }
