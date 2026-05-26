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

func newReadTool(runtime *runtime) (tool.Tool, error) {
	_ = "STUB: not implemented"
	return *new(tool.Tool), nil
}

func readText(runtime *runtime, snapshot localFileSnapshot, in readInput) (readOutput, error) {
	_ = "STUB: not implemented"
	return *new(readOutput), nil
}

func readNotebook(runtime *runtime, snapshot localFileSnapshot, in readInput) (readOutput, error) {
	_ = "STUB: not implemented"
	return *new(readOutput), nil
}

func readPDF(runtime *runtime, snapshot localFileSnapshot, in readInput) (readOutput, error) {
	_ = "STUB: not implemented"
	return *new(readOutput), nil
}

func readImage(runtime *runtime, snapshot localFileSnapshot, in readInput) (readOutput, error) {
	_ = "STUB: not implemented"
	return *new(readOutput), nil
}

func readDescription() string { _ = "STUB: not implemented"; return "" }
