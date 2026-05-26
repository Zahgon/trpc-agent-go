//go:build windows

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
	"time"
)

func preparePipeCommand(_ *exec.Cmd) { _ = "STUB: not implemented"; return }

func preparePTYCommand(_ *exec.Cmd) { _ = "STUB: not implemented"; return }

func commandProcessGroupID(_ *exec.Cmd) int { _ = "STUB: not implemented"; return 0 }

func terminateProcessTree(
	_ context.Context,
	process *os.Process,
	_ int,
	_ time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}
