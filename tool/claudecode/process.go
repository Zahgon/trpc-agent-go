//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package claudecode

import (
	"bytes"
	"context"
	"os"
	"sync"
)

type capturedProcessResult struct {
	Stdout   []byte
	Stderr   []byte
	ExitCode int
}

type waitedProcessState struct {
	State *os.ProcessState
	Err   error
}

type processCapture struct {
	stdout bytes.Buffer
	stderr bytes.Buffer
	wg     sync.WaitGroup
}

func runCapturedProcess(
	ctx context.Context,
	dir string,
	env []string,
	bin string,
	args ...string,
) (capturedProcessResult, error) {
	_ = "STUB: not implemented"
	return *new(capturedProcessResult), nil
}

func startProcess(
	dir string,
	env []string,
	stdoutFile *os.File,
	stderrFile *os.File,
	bin string,
	args ...string,
) (*os.Process, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processPipes() (
	*os.File,
	*os.File,
	*os.File,
	*os.File,
	*os.File,
	func() error,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil
}

func processEnv(extra []string) []string { _ = "STUB: not implemented"; return nil }

func startProcessCapture(stdoutReader *os.File, stderrReader *os.File) *processCapture {
	_ = "STUB: not implemented"
	return nil
}

func (c *processCapture) wait() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func waitForProcess(ctx context.Context, proc *os.Process) waitedProcessState {
	_ = "STUB: not implemented"
	return *new(waitedProcessState)
}
