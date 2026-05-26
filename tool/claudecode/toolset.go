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

// NewToolSet constructs a Claude Code-compatible toolset.
func NewToolSet(opts ...Option) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

func newToolRuntime(baseAbs string, maxFileSize int64) *runtime {
	_ = "STUB: not implemented"
	return nil
}

func appendCoreTools(cc *compositeToolSet, rt *runtime, readOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

func appendWebTools(cc *compositeToolSet, options toolSetOptions) error {
	_ = "STUB: not implemented"
	return nil
}
