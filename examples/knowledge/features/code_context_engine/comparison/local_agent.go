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

// localAgentConfig controls how the local code-search agent is wired up.
// The agent connects to our own MCP server (see comparison/../mcp/server.go)
// that exposes the AST-backed code_search tool over MCP.
type localAgentConfig struct {
	ModelName    string
	MCPServerURL string
	MCPToolName  string
}

type localCodeSearchAgentRunner struct {
	runner  runner.Runner
	toolSet *mcp.ToolSet
}

// newLocalCodeSearchAgentRunner builds an LLM agent whose only tool is the
// remote code_search exposed by our local MCP server. The MCP server owns the
// knowledge base (vector store + embedder + repo source + code_search) and is
// responsible for its own ingestion lifecycle.
func newLocalCodeSearchAgentRunner(cfg localAgentConfig) (*localCodeSearchAgentRunner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *localCodeSearchAgentRunner) RunCase(ctx context.Context, c comparisonCase) (*agentRunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *localCodeSearchAgentRunner) Close() error { _ = "STUB: not implemented"; return nil }

const localCodeSearchAgentInstruction = `You are a repository code assistant.`
