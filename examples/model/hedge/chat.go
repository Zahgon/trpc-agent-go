//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/runner"
)

type hedgeChat struct {
	config    appConfig
	runner    runner.Runner
	userID    string
	sessionID string
}

func (c *hedgeChat) run() error { _ = "STUB: not implemented"; return nil }

func (c *hedgeChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *hedgeChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *hedgeChat) startNewSession() { _ = "STUB: not implemented"; return }

func (c *hedgeChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func newSessionID() string { _ = "STUB: not implemented"; return "" }
