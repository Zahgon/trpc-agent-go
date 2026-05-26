//go:build !windows

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
	"os"
	"os/exec"
	"syscall"
	"time"
)

const (
	processLivenessSignal = syscall.Signal(0)
	processKillPoll       = 10 * time.Millisecond
)

func preparePipeCommand(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }

func preparePTYCommand(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }

func ensureSysProcAttr(cmd *exec.Cmd) *syscall.SysProcAttr { _ = "STUB: not implemented"; return nil }

func commandProcessGroupID(cmd *exec.Cmd) int { _ = "STUB: not implemented"; return 0 }

// Both spawn paths make the child the leader of its owned group.
// Pipe mode sets Setpgid=true, and PTY mode starts a new session.
// That keeps PGID == PID here, while signalProcessTree still falls
// back to direct process signals if group signaling fails.

func terminateProcessTree(
	ctx context.Context,
	process *os.Process,
	processGroupID int,
	grace time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func signalProcessTree(
	process *os.Process,
	processGroupID int,
	signal syscall.Signal,
) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForProcessTreeExit(
	ctx context.Context,
	process *os.Process,
	processGroupID int,
	grace time.Duration,
) bool {
	_ = "STUB: not implemented"
	return false
}

func processTreeAlive(
	process *os.Process,
	processGroupID int,
) bool {
	_ = "STUB: not implemented"
	return false
}
