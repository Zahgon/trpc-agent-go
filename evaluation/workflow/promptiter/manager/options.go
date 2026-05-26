//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package manager

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/store"
)

// Option configures the PromptIter run manager.
type Option func(*options)

type options struct {
	store                store.Store
	storedResultSlimming engine.RunResultSlimming
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithStore sets the store used to persist PromptIter runs.
func WithStore(store store.Store) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStoredResultSlimming omits selected fields before runs are persisted.
func WithStoredResultSlimming(slimming engine.RunResultSlimming) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
