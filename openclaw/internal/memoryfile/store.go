//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package memoryfile

import (
	"context"
	"sync"
)

const (
	rootDirName    = "memory"
	memoryFileName = "MEMORY.md"

	dirPerm           = 0o700
	filePerm          = 0o600
	tempPatternSuffix = ".tmp-*"
)

type Store struct {
	root string

	mu sync.Mutex
}

func DefaultRoot(stateDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func NewStore(root string) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Store) Root() string { _ = "STUB: not implemented"; return "" }

func (s *Store) MemoryDir(appName string, userID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Store) MemoryPath(appName string, userID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Store) EnsureMemory(
	ctx context.Context,
	appName string,
	userID string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Store) UpdateMemory(
	ctx context.Context,
	appName string,
	userID string,
	update func(current string) (string, error),
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SaveResolvedMemoryFile writes a validated file path while holding the
// same store lock used by logical memory updates.
func (s *Store) SaveResolvedMemoryFile(
	ctx context.Context,
	path string,
	content string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) ReadFile(path string, maxBytes int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Store) resolveFilePath(path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Store) DeleteUser(
	ctx context.Context,
	appName string,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) removeScopedDir(ctx context.Context, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

func sanitizePathPart(raw string) string { _ = "STUB: not implemented"; return "" }

func writeFileAtomic(path string, data []byte) error { _ = "STUB: not implemented"; return nil }

func fileExists(path string) bool { _ = "STUB: not implemented"; return false }

func contextErr(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
