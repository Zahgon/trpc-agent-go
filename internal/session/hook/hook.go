//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package hook provides internal hook execution utilities for session services.
package hook

import (
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// RunAppendEventHooks executes AppendEvent hooks chain.
// The final hook performs the actual storage operation.
func RunAppendEventHooks(
	hooks []session.AppendEventHook,
	ctx *session.AppendEventContext,
	final session.AppendEventHook,
) error {
	_ = "STUB: not implemented"
	// Wrap final as a hook that ignores next (it's the terminal)
	return nil
}

// RunGetSessionHooks executes GetSession hooks chain.
// The final hook performs the actual storage retrieval.
func RunGetSessionHooks(
	hooks []session.GetSessionHook,
	ctx *session.GetSessionContext,
	final session.GetSessionHook,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	// Wrap final as a hook that ignores next (it's the terminal)
	return nil, nil
}
