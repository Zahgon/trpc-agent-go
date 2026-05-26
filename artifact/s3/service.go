//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package s3 provides an S3-compatible artifact storage service implementation.
// It supports AWS S3, MinIO, DigitalOcean Spaces, Cloudflare R2, and other
// S3-compatible object storage services.
package s3

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/artifact"
	"trpc.group/trpc-go/trpc-agent-go/log"
	s3storage "trpc.group/trpc-go/trpc-agent-go/storage/s3"
)

// defaultContentType is the fallback MIME type for artifacts without one.
const defaultContentType = "application/octet-stream"

// Compile-time check that Service implements artifact.Service.
var _ artifact.Service = (*Service)(nil)

// Service is an S3-compatible implementation of the artifact service.
// It supports AWS S3, MinIO, DigitalOcean Spaces, Cloudflare R2, and other
// S3-compatible object storage services.
//
// The object name format used depends on whether the filename has a user namespace:
//   - For files with user namespace (starting with "user:"):
//     {app_name}/{user_id}/user/{filename}/{version}
//   - For regular session-scoped files:
//     {app_name}/{user_id}/{session_id}/{filename}/{version}
type Service struct {
	client     s3storage.Client
	ownsClient bool // true if we created the client, false if provided via WithClient
	logger     log.Logger
}

// NewService creates a new S3 artifact service.
// When using WithClient, the bucket parameter is ignored as the client already has one configured.
func NewService(ctx context.Context, bucket string, opts ...Option) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close releases any resources held by the service.
// If the client was provided externally via WithClient, it is not closed.
func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

// SaveArtifact saves an artifact to S3.
// It automatically determines the next version number by listing existing versions.
//
// Concurrency: This method is NOT safe for concurrent writes to the same filename.
// If multiple goroutines save the same artifact concurrently, they may compute
// the same version number, causing one write to overwrite the other.
// For concurrent access, use external synchronization or unique filenames.
func (s *Service) SaveArtifact(
	ctx context.Context,
	sessionInfo artifact.SessionInfo,
	filename string,
	art *artifact.Artifact,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// LoadArtifact loads an artifact from S3.
// If version is nil, the latest version is loaded.
func (s *Service) LoadArtifact(
	ctx context.Context,
	sessionInfo artifact.SessionInfo,
	filename string,
	version *int,
) (*artifact.Artifact, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Artifact not found

// ListArtifactKeys lists all artifact filenames within a session.
// It returns artifacts from both session scope and user scope.
func (s *Service) ListArtifactKeys(
	ctx context.Context,
	sessionInfo artifact.SessionInfo,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteArtifact deletes all versions of an artifact from S3.
func (s *Service) DeleteArtifact(
	ctx context.Context,
	sessionInfo artifact.SessionInfo,
	filename string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ListVersions lists all versions of an artifact.
func (s *Service) ListVersions(
	ctx context.Context,
	sessionInfo artifact.SessionInfo,
	filename string,
) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) listVersions(
	ctx context.Context,
	sessionInfo artifact.SessionInfo,
	filename string,
) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractFilename extracts the filename from an object key given a prefix.
// Object key format: {prefix}{filename}/{version}
// Returns the filename or empty string if the key doesn't match the expected format.
func extractFilename(objectKey, prefix string) string { _ = "STUB: not implemented"; return "" }

// validateSessionInfo checks that all required session info fields are present.
func validateSessionInfo(info artifact.SessionInfo) error { _ = "STUB: not implemented"; return nil }

// validateFilename checks that the filename is valid and safe.
// It rejects empty filenames, path traversal attempts, and other dangerous patterns.
func validateFilename(filename string) error { _ = "STUB: not implemented"; return nil }

// Check for path traversal and invalid characters
// Note: "user:" prefix is allowed for user-scoped artifacts
