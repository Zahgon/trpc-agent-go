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
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

func (r *runner) applyAwaitUserReplyRoute(
	ctx context.Context,
	key session.Key,
	sess *session.Session,
	message model.Message,
	ro agent.RunOptions,
) (agent.RunOptions, string, error) {
	_ = "STUB: not implemented"
	return *new(agent.RunOptions), "", nil
}

func (r *runner) clearOverriddenAwaitUserReplyRoute(
	ctx context.Context,
	key session.Key,
	sess *session.Session,
	ro agent.RunOptions,
) (agent.RunOptions, string, error) {
	_ = "STUB: not implemented"
	return *new(agent.RunOptions), "", nil
}

func (r *runner) clearAwaitUserReplyRoute(
	ctx context.Context,
	key session.Key,
	sess *session.Session,
) error {
	_ = "STUB: not implemented"
	return nil
}
