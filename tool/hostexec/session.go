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
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

const (
	programStatusRunning = "running"
	programStatusExited  = "exited"
)

type session struct {
	id      string
	command string

	cmd     *exec.Cmd
	stdin   io.WriteCloser
	closeIO func() error
	cancel  context.CancelFunc

	processGroupID int

	doneCh chan struct{}
	ioDone chan struct{}
	ioWG   sync.WaitGroup

	mu       sync.Mutex
	started  time.Time
	finished time.Time
	exitCode int

	lineBase   int
	lines      []string
	partial    string
	pollCursor int
	maxLines   int
	closeOnce  sync.Once
}

func newSession(id string, command string, maxLines int) *session {
	_ = "STUB: not implemented"
	return nil
}

func newSessionID() string { _ = "STUB: not implemented"; return "" }

func (s *session) running() bool { _ = "STUB: not implemented"; return false }

func (s *session) doneAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *session) markDone(exitCode int) { _ = "STUB: not implemented"; return }

func (s *session) readFrom(reader io.Reader) { _ = "STUB: not implemented"; return }

func (s *session) appendOutput(chunk string) { _ = "STUB: not implemented"; return }

func (s *session) trimLocked() { _ = "STUB: not implemented"; return }

func (s *session) tail(lines int) string { _ = "STUB: not implemented"; return "" }

func trimOutputTail(output string, lines int) string { _ = "STUB: not implemented"; return "" }

func (s *session) pollTail(lines int) string { _ = "STUB: not implemented"; return "" }

func (s *session) allOutput() (string, int) { _ = "STUB: not implemented"; return "", 0 }

type processPoll struct {
	Status     string
	Output     string
	Offset     int
	NextOffset int
	ExitCode   *int
}

func (s *session) poll(limit *int) processPoll { _ = "STUB: not implemented"; return *new(processPoll) }

func (s *session) write(data string, newline bool) error { _ = "STUB: not implemented"; return nil }

func killProcess(process *os.Process) error { _ = "STUB: not implemented"; return nil }

func (s *session) kill(
	ctx context.Context,
	grace time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *session) close() error { _ = "STUB: not implemented"; return nil }
