//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package usersimulation

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
)

type options struct {
	stopSignal            *string
	maxAllowedInvocations *int
	userIDSupplier        func(context.Context) string
	sessionIDSupplier     func(context.Context) string
	systemPromptBuilder   SystemPromptBuilder
}

// Option configures the default simulator implementation.
type Option func(*options)

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithStopSignal overrides the scenario stop signal for the default simulator.
func WithStopSignal(signal string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxAllowedInvocations overrides the scenario turn limit for the default simulator.
func WithMaxAllowedInvocations(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserIDSupplier overrides the internal simulator user ID supplier.
func WithUserIDSupplier(supplier func(context.Context) string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSessionIDSupplier overrides the internal simulator session ID supplier.
func WithSessionIDSupplier(supplier func(context.Context) string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// SystemPromptBuilder builds the simulator's initial system prompt text.
type SystemPromptBuilder func(ctx context.Context, scenario *evalset.ConversationScenario) string

// WithSystemPromptBuilder overrides the initial system prompt builder for the default simulator.
func WithSystemPromptBuilder(builder SystemPromptBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
