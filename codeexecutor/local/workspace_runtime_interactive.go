//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package local

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

const (
	defaultInteractiveMaxLines = 20_000
	defaultInteractiveKillWait = 2 * time.Second
	envPathKey                 = "PATH"
	envPathExtKey              = "PATHEXT"
)

var errInteractiveTTYWindows = errors.New(
	"interactive tty is not supported on windows",
)

type interactiveSession struct {
	id      string
	command string

	cmd     *exec.Cmd
	stdin   io.WriteCloser
	closeIO func() error
	cancel  context.CancelFunc

	doneCh chan struct{}
	ioDone chan struct{}
	ioWG   sync.WaitGroup

	// pipeDrain closes parent-held stdout/stderr write ends (os.Pipe path only).
	// Called after cmd.Wait() so read goroutines observe EOF before ioWG.Wait().
	pipeDrain func()

	mu       sync.Mutex
	started  time.Time
	finished time.Time
	exitCode int
	timedOut bool
	duration time.Duration

	lineBase   int
	lines      []string
	partial    string
	pollCursor int
	maxLines   int
	closeOnce  sync.Once
	stdout     strings.Builder
	stderr     strings.Builder
}

func newInteractiveSession(
	id string,
	command string,
	maxLines int,
) *interactiveSession {
	_ = "STUB: not implemented"
	return nil
}

func (s *interactiveSession) ID() string { _ = "STUB: not implemented"; return "" }

func (s *interactiveSession) Poll(limit *int) codeexecutor.ProgramPoll {
	_ = "STUB: not implemented"
	return *new(codeexecutor.ProgramPoll)
}

func (s *interactiveSession) Log(
	offset *int,
	limit *int,
) codeexecutor.ProgramLog {
	_ = "STUB: not implemented"
	return *new(codeexecutor.ProgramLog)
}

func (s *interactiveSession) Write(
	data string,
	newline bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *interactiveSession) Kill(grace time.Duration) error { _ = "STUB: not implemented"; return nil }

func (s *interactiveSession) Close() error { _ = "STUB: not implemented"; return nil }

func (s *interactiveSession) State() codeexecutor.ProgramState {
	_ = "STUB: not implemented"
	return *new(codeexecutor.ProgramState)
}

func (s *interactiveSession) RunResult() codeexecutor.RunResult {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult)
}

func (s *interactiveSession) markDone(
	exitCode int,
	duration time.Duration,
	timedOut bool,
) {
	_ = "STUB: not implemented"
	return
}

func (s *interactiveSession) readFrom(
	r io.Reader,
	stream string,
) {
	_ = "STUB: not implemented"
	return
}

func (s *interactiveSession) appendOutput(
	chunk string,
	stream string,
) {
	_ = "STUB: not implemented"
	return
}

func (s *interactiveSession) trimLocked() { _ = "STUB: not implemented"; return }

// StartProgram starts an interactive program in the workspace.
func (r *Runtime) StartProgram(
	ctx context.Context,
	ws codeexecutor.Workspace,
	spec codeexecutor.InteractiveProgramSpec,
) (codeexecutor.ProgramSession, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.ProgramSession), nil
}

func (r *Runtime) buildProgramEnv(
	ws codeexecutor.Workspace,
	spec codeexecutor.RunProgramSpec,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLocalProgramCommand(
	ctx context.Context,
	cwd string,
	spec codeexecutor.RunProgramSpec,
	env []string,
) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func localProgramCommandPath(
	cwd string,
	name string,
	env []string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func envValue(env []string, key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func envKeyEqual(a string, b string) bool { _ = "STUB: not implemented"; return false }

func envKeyEqualForGOOS(a string, b string, goos string) bool {
	_ = "STUB: not implemented"
	return false
}

func localProgramCandidateNames(name string, pathExt string) []string {
	_ = "STUB: not implemented"
	return nil
}

func localProgramCandidateNamesForGOOS(
	name string,
	pathExt string,
	goos string,
	pathListSep string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func isBareLocalCommand(name string) bool { _ = "STUB: not implemented"; return false }

func isLocalExecutableFile(name string) bool { _ = "STUB: not implemented"; return false }

func isLocalExecutableFileForGOOS(name string, goos string) bool {
	_ = "STUB: not implemented"
	return false
}

func formatInteractiveCommand(cmd string, args []string) string {
	_ = "STUB: not implemented"
	return ""
}

func interactiveExitCode(err error) int { _ = "STUB: not implemented"; return 0 }

func startPipes(
	cmd *exec.Cmd,
) (io.WriteCloser, io.ReadCloser, io.ReadCloser, func(), error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), *new(io.ReadCloser), *new(io.ReadCloser), nil, nil
}
