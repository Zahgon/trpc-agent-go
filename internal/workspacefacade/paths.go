//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package workspacefacade hosts implementation primitives shared by the
// codeexecutor/workspaceio facade and the tool/workspaceexec LLM tools.
// Nothing in this package is part of the public API.
package workspacefacade

// HasGlobMeta reports whether s contains any glob metacharacter.
func HasGlobMeta(s string) bool { _ = "STUB: not implemented"; return false }

// IsWorkspaceEnvPath reports whether s starts with a recognized
// workspace env-prefixed path such as $WORK_DIR/... or
// ${SKILLS_DIR}/....
func IsWorkspaceEnvPath(s string) bool { _ = "STUB: not implemented"; return false }

// NormalizeArtifactPath validates a single-file path used by artifact
// publishing entry points (workspace_save_artifact LLM tool and
// Workspace.SaveArtifact). Globs and parent traversal are rejected.
// The returned path is always workspace-relative, clean, and confirmed
// to live under one of the supported publish roots (work/, out/,
// runs/).
func NormalizeArtifactPath(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// IsAllowedPublishArtifactPath reports whether rel resolves to a
// workspace path under work/, out/, or runs/.
func IsAllowedPublishArtifactPath(rel string) bool { _ = "STUB: not implemented"; return false }

// IsAllowedWorkspaceRoot reports whether rel resolves to one of the
// workspace roots reachable for read/write/exec operations
// (skills/, work/, out/, runs/). Looser than
// IsAllowedPublishArtifactPath, which excludes skills/ because artifact
// publishing must not target skill assets.
func IsAllowedWorkspaceRoot(rel string) bool { _ = "STUB: not implemented"; return false }

// NormalizeWorkspaceCWD canonicalises a working-directory string for
// program execution against the current invocation's workspace. It is
// the single source of truth for "is this Cwd safe?" — both the
// workspace_exec LLM tool and Workspace.RunProgram forward to it so a
// single rule set governs Cwd containment.
//
// Behavior:
//   - empty / whitespace-only returns ".", meaning workspace root.
//   - glob metacharacters are rejected.
//   - $WORK / $OUT / $RUNS / $SKILLS env-prefixed paths are expanded
//     to their workspace-relative form via codeexecutor.NormalizeGlobs.
//   - absolute paths (leading "/") are stripped to workspace-relative
//     and then checked against IsAllowedWorkspaceRoot.
//   - relative paths that traverse out of the workspace ("..", "../*")
//     are rejected outright.
//   - any other relative path must resolve under
//     IsAllowedWorkspaceRoot (skills/, work/, out/, runs/).
func NormalizeWorkspaceCWD(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

const (
	envVarPrefix = "$"
	envVarLBrace = "${"
	envVarRBrace = "}"
)

// hasEnvPrefix reports whether s starts with an env reference such as
// $name or ${name} followed by either nothing or a path separator.
func hasEnvPrefix(s, name string) bool { _ = "STUB: not implemented"; return false }
