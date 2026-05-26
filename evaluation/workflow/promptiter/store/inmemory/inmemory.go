//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package inmemory provides an in-memory PromptIter store implementation.
package inmemory

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/store"
)

type inMemoryStore struct {
	mu   sync.RWMutex
	runs map[string]map[string]*engine.RunResult
}

// New creates an in-memory PromptIter store.
func New() store.Store { _ = "STUB: not implemented"; return *new(store.Store) }

func (s *inMemoryStore) Create(_ context.Context, appName string, run *engine.RunResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *inMemoryStore) Get(_ context.Context, appName, runID string) (*engine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *inMemoryStore) Update(_ context.Context, appName string, run *engine.RunResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *inMemoryStore) Close() error { _ = "STUB: not implemented"; return nil }

func cloneRun(run *engine.RunResult) (*engine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateRun(appName string, run *engine.RunResult) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRunKey(appName, runID string) error { _ = "STUB: not implemented"; return nil }
