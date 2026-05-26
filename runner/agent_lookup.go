//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package runner

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

func (r *runner) loadRegisteredAgent(
	ctx context.Context,
	agentName string,
	ro agent.RunOptions,
) (agent.Agent, bool, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), false, nil
}

func (r *runner) resolveAwaitUserReplyRoute(
	ctx context.Context,
	route agent.AwaitUserReplyRoute,
	ro agent.RunOptions,
) (agent.Agent, string, bool, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), "", false, nil
}

func splitAgentPath(path string) []string { _ = "STUB: not implemented"; return nil }
