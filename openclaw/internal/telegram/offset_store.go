//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telegram

import (
	"context"
)

const offsetStoreVersion = 1

// OffsetStore persists the getUpdates offset.
type OffsetStore interface {
	Read(ctx context.Context) (int, bool, error)
	Write(ctx context.Context, offset int) error
}

// FileOffsetStore stores offsets in a local JSON file.
type FileOffsetStore struct {
	path string
}

// NewFileOffsetStore creates a file-based offset store.
func NewFileOffsetStore(path string) (*FileOffsetStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type offsetStoreState struct {
	Version int `json:"version"`
	Offset  int `json:"offset"`
}

// Read returns the stored offset.
func (s *FileOffsetStore) Read(
	ctx context.Context,
) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// Write persists the next offset.
func (s *FileOffsetStore) Write(ctx context.Context, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

func randomHex(n int) string { _ = "STUB: not implemented"; return "" }
