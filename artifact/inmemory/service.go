//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package inmemory provides an in-memory implementation of the artifact service.
package inmemory

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/artifact"
)

// Service is an in-memory implementation of the artifact service.
// It is suitable for testing and development environments.
type Service struct {
	// mutex protects concurrent access to the artifacts map
	mutex sync.RWMutex
	// artifacts stores artifacts by path, with each path containing a list of versions
	artifacts map[string][]*artifact.Artifact
}

// NewService creates a new in-memory artifact service.
func NewService() *Service { _ = "STUB: not implemented"; return nil }

// SaveArtifact saves an artifact to the in-memory storage.
func (s *Service) SaveArtifact(ctx context.Context, sessionInfo artifact.SessionInfo, filename string, art *artifact.Artifact) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// LoadArtifact gets an artifact from the in-memory storage.
func (s *Service) LoadArtifact(ctx context.Context, sessionInfo artifact.SessionInfo, filename string, version *int) (*artifact.Artifact, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the latest version (last element)

// ListArtifactKeys lists all the artifact filenames within a session.
func (s *Service) ListArtifactKeys(ctx context.Context, sessionInfo artifact.SessionInfo) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteArtifact deletes an artifact.
func (s *Service) DeleteArtifact(ctx context.Context, sessionInfo artifact.SessionInfo, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// Artifact doesn't exist, but this is not an error in the Python implementation

// ListVersions lists all versions of an artifact.
func (s *Service) ListVersions(ctx context.Context, sessionInfo artifact.SessionInfo, filename string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
