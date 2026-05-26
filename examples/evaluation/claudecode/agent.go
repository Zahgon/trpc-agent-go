//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/claudecode"
)

func newClaudeCodeEvalAgent(bin, outputFormat, workDir, logDir string) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func newLogHook(outDir string) claudecode.RawOutputHook {
	_ = "STUB: not implemented"
	return *new(claudecode.RawOutputHook)
}
