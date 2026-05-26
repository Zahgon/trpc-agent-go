//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package fileref parses and reads file references like workspace://... and
// artifact://....
//
// Tools use these references to share a unified file view.
package fileref

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/internal/toolcache"
)

const (
	schemeSep = "://"

	// SchemeArtifact is the artifact:// file ref scheme.
	SchemeArtifact = "artifact"
	// SchemeWorkspace is the workspace:// file ref scheme.
	SchemeWorkspace = "workspace"

	// ArtifactPrefix is the "artifact://" prefix.
	ArtifactPrefix = SchemeArtifact + schemeSep
	// WorkspacePrefix is the "workspace://" prefix.
	WorkspacePrefix = SchemeWorkspace + schemeSep

	// SchemeHost is the host:// file ref scheme used by local runtimes.
	SchemeHost = "host"
	// HostPrefix is the "host://" prefix.
	HostPrefix = SchemeHost + schemeSep

	filePrefix = "file://"
)

const errArtifactNameEmpty = "artifact name is empty"

// Ref is a parsed file reference.
//
// When Scheme is empty, Path is a caller-defined local path
// (for example, relative to a file tool base directory).
type Ref struct {
	Scheme          string
	Path            string
	ArtifactName    string
	ArtifactVersion *int
	Raw             string
}

// WorkspaceRef builds a workspace:// reference for the given relative path.
func WorkspaceRef(rel string) string { _ = "STUB: not implemented"; return "" }

// Parse parses raw into a Ref.
//
// When the returned Ref has an empty Scheme, the caller should treat Path as
// a local path (for example, relative to a tool base directory).
func Parse(raw string) (Ref, error) { _ = "STUB: not implemented"; return *new(Ref), nil }

// IsInternalFileRef reports whether raw is an internal or local-only file ref
// that should not be forwarded to model providers as a provider file_id.
func IsInternalFileRef(raw string) bool { _ = "STUB: not implemented"; return false }

// DisplayName returns a safe basename for a supported internal file ref.
func DisplayName(raw string) string { _ = "STUB: not implemented"; return "" }

func refBaseName(raw string) string { _ = "STUB: not implemented"; return "" }

func cleanRelPath(p string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// TryRead reads raw if it is a supported file reference.
//
// When handled is false, raw is not a reference and the caller should treat
// it as a local path.
func TryRead(
	ctx context.Context,
	raw string,
) (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

// WorkspaceFiles returns files exported from skill_run output_files in ctx.
func WorkspaceFiles(
	ctx context.Context,
) []toolcache.SkillRunOutputFile {
	_ = "STUB: not implemented"
	return nil
}

func loadArtifactFromContext(
	ctx context.Context,
	name string,
	version *int,
) ([]byte, string, int, error) {
	_ = "STUB: not implemented"
	return nil, "", 0, nil
}

func withArtifactContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
