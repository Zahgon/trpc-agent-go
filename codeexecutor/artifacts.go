//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package codeexecutor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/artifact"
)

// LoadArtifactHelper resolves artifact name@version via callback context.
// If version is nil, loads latest. Returns data, mime, actual version.
func LoadArtifactHelper(
	ctx context.Context, name string, version *int,
) ([]byte, string, int, error) {
	_ = "STUB: not implemented"
	return nil, "", 0, nil
}

func resolveArtifactVersion(
	ctx context.Context,
	svc artifact.Service,
	info artifact.SessionInfo,
	name string,
	version *int,
) int {
	_ = "STUB: not implemented"
	return 0
}

// ParseArtifactRef splits "name@version" into name and optional version.
func ParseArtifactRef(ref string) (string, *int, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func isDecimalVersion(s string) bool { _ = "STUB: not implemented"; return false }

func parseDecimalVersion(s string) int { _ = "STUB: not implemented"; return 0 }

// SaveArtifactHelper saves a file as artifact using callback context.
func SaveArtifactHelper(
	ctx context.Context, filename string, data []byte, mime string,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WithArtifactService attaches artifact.Service to context so lower
// layers (codeexecutor) can resolve artifacts without importing agent.
type artifactKey struct{}
type artifactSessionKey struct{}

// WithArtifactService stores an artifact.Service in the context.
// Callers retrieve it in lower layers to load/save artifacts
// without importing higher-level packages.
func WithArtifactService(
	ctx context.Context, svc artifact.Service,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ArtifactServiceFromContext fetches the artifact.Service previously
// stored by WithArtifactService. It returns the service and a boolean
// indicating presence.
func ArtifactServiceFromContext(
	ctx context.Context,
) (artifact.Service, bool) {
	_ = "STUB: not implemented"
	return *new(artifact.Service), false
}

// WithArtifactSession stores artifact session info in context.
func WithArtifactSession(
	ctx context.Context, info artifact.SessionInfo,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func artifactSessionFromContext(
	ctx context.Context,
) artifact.SessionInfo {
	_ = "STUB: not implemented"
	return *new(artifact.SessionInfo)
}
