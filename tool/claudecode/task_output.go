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
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func newTaskOutputTool(runtime *runtime) (tool.Tool, error) {
	_ = "STUB: not implemented"
	return *new(tool.Tool), nil
}

func readTaskSnapshot(runtime *runtime, taskID string) (*taskOutputTask, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func snapshotBackgroundTask(runtime *runtime, taskID string) (*backgroundTask, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func taskOutputDescription() string { _ = "STUB: not implemented"; return "" }
