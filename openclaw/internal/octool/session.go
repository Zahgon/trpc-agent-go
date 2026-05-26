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

type session struct {
	id      string
	command string
	redact  func(string) string

	cmd     *exec.Cmd
	stdin   io.WriteCloser
	closeIO func() error
	cancel  context.CancelFunc

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
}

func newSession(id, command string, maxLines int) *session { _ = "STUB: not implemented"; return nil }

func newSessionID() string { _ = "STUB: not implemented"; return "" }

func (s *session) running() bool { _ = "STUB: not implemented"; return false }

func (s *session) doneAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *session) markDone(exitCode int) { _ = "STUB: not implemented"; return }

func (s *session) readFrom(r io.Reader) { _ = "STUB: not implemented"; return }

func (s *session) appendOutput(chunk string) { _ = "STUB: not implemented"; return }

func (s *session) trimLocked() { _ = "STUB: not implemented"; return }

func (s *session) tail(lines int) string { _ = "STUB: not implemented"; return "" }

func (s *session) allOutput() (string, int) { _ = "STUB: not implemented"; return "", 0 }

type processSession struct {
	SessionID string `json:"sessionId"`
	Command   string `json:"command"`
	Status    string `json:"status"`
	StartedAt string `json:"startedAt"`
	DoneAt    string `json:"doneAt,omitempty"`
	ExitCode  *int   `json:"exitCode,omitempty"`
}

// ProcessSession is the exported view of one exec_command session.
type ProcessSession = processSession

func (s *session) snapshot() processSession { _ = "STUB: not implemented"; return *new(processSession) }

type processPoll struct {
	Status     string `json:"status"`
	Output     string `json:"output,omitempty"`
	Offset     int    `json:"offset"`
	NextOffset int    `json:"nextOffset"`
	ExitCode   *int   `json:"exitCode,omitempty"`
}

func (s *session) poll(limit *int) processPoll { _ = "STUB: not implemented"; return *new(processPoll) }

type processLog struct {
	Output     string `json:"output,omitempty"`
	Offset     int    `json:"offset"`
	NextOffset int    `json:"nextOffset"`
}

func (s *session) log(offset *int, limit *int) processLog {
	_ = "STUB: not implemented"
	return *new(processLog)
}

type processWrite struct {
	OK bool `json:"ok"`
}

func (s *session) write(data string, newline bool) (processWrite, error) {
	_ = "STUB: not implemented"
	return *new(processWrite), nil
}

func (s *session) kill(grace time.Duration) error { _ = "STUB: not implemented"; return nil }

func sortSessions(sessions []processSession) { _ = "STUB: not implemented"; return }
