//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/runner"
	mcp "trpc.group/trpc-go/trpc-agent-go/tool/mcp"
)

type augmentCodeAgentRunner struct {
	runner  runner.Runner
	toolSet *mcp.ToolSet
}

func newAugmentCodeAgentRunner(modelName string) (*augmentCodeAgentRunner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *augmentCodeAgentRunner) RunCase(ctx context.Context, c comparisonCase) (*agentRunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *augmentCodeAgentRunner) Close() error { _ = "STUB: not implemented"; return nil }

const augmentCodeSearchAgentInstruction = `You are a repository code assistant.

When calling augment_code_search, always use:
- repo_owner: trpc-group
- repo_name:  trpc-agent-go
- branch:     main
`
