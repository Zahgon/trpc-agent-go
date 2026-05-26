//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package mysql

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/store"
	storage "trpc.group/trpc-go/trpc-agent-go/storage/mysql"
)

type mysqlStore struct {
	db        storage.Client
	tableName string
}

// New creates a MySQL-backed PromptIter store.
func New(opts ...Option) (store.Store, error) {
	_ = "STUB: not implemented"
	return *new(store.Store), nil
}

// Create persists one new PromptIter run.
func (s *mysqlStore) Create(ctx context.Context, appName string, run *engine.RunResult) error {
	_ = "STUB: not implemented"
	return nil
}

// Pass JSON as a UTF-8 string so the driver does not bind []byte as BINARY (MySQL JSON rejects binary charset).

// Get loads one persisted PromptIter run by app name and run ID.
func (s *mysqlStore) Get(ctx context.Context, appName, runID string) (*engine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update persists changes to one existing PromptIter run.
func (s *mysqlStore) Update(ctx context.Context, appName string, run *engine.RunResult) error {
	_ = "STUB: not implemented"
	return nil
}

// Close releases the underlying MySQL client.
func (s *mysqlStore) Close() error { _ = "STUB: not implemented"; return nil }

func validateRun(appName string, run *engine.RunResult) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRunKey(appName, runID string) error { _ = "STUB: not implemented"; return nil }
