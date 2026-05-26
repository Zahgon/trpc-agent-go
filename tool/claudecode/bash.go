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
	"context"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func newBashTool(runtime *runtime) (tool.Tool, error) {
	_ = "STUB: not implemented"
	return *new(tool.Tool), nil
}

func runForegroundCommand(ctx context.Context, runtime *runtime, in bashInput) (bashOutput, error) {
	_ = "STUB: not implemented"
	return *new(bashOutput), nil
}

func bashTimeout(timeout *int) int { _ = "STUB: not implemented"; return 0 }

func runBackgroundCommand(runtime *runtime, command string) (bashOutput, error) {
	_ = "STUB: not implemented"
	return *new(bashOutput), nil
}

func backgroundTaskStatus(waitErr error, state *os.ProcessState) string {
	_ = "STUB: not implemented"
	return ""
}

func backgroundTaskExitCode(waitErr error, state *os.ProcessState) int {
	_ = "STUB: not implemented"
	return 0
}

func errorsIsDeadlineExceeded(err error) bool { _ = "STUB: not implemented"; return false }

func bashDescription() string { _ = "STUB: not implemented"; return "" }
