//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package codeexecutor holds workspace metadata helpers and constants.
package codeexecutor

import (
	"context"
	"io/fs"
	"sync"
	"time"
)

// Well-known subdirectories in a workspace.
const (
	// DirSkills contains session-scoped skill working copies. Skills
	// staged here are writable by default so third-party scripts can
	// emit cache files, temporary outputs, and Python bytecode next to
	// their source. Callers that need a stable, canonical skill tree
	// should treat the upstream skill repository as the source of
	// truth; the copy under this directory is a working copy tied to
	// the current session.
	DirSkills = "skills"
	// DirWork contains writable shared intermediates.
	DirWork = "work"
	// DirRuns contains per-run working directories.
	DirRuns = "runs"
	// DirOut contains collected outputs for artifacting.
	DirOut = "out"
	// MetaFileName is the metadata file name at workspace root.
	MetaFileName = "metadata.json"
)

// Additional environment variable keys injected at runtime.
const (
	EnvSkillsDir = "SKILLS_DIR"
	EnvWorkDir   = "WORK_DIR"
	EnvOutputDir = "OUTPUT_DIR"
	EnvRunDir    = "RUN_DIR"
	EnvSkillName = "SKILL_NAME"
)

const (
	metadataFileMode       fs.FileMode = 0o600
	legacyMetadataTmpName              = ".metadata.tmp"
	metadataTmpPrefix                  = ".metadata."
	metadataTmpSuffix                  = ".tmp"
	metadataTmpPartCount               = 4
	metadataRandomHexLen               = 16
	metadataNoRandomSuffix             = "norand"
	emptyMetadataLockKey               = "__empty_workspace__"
)

var (
	metadataTmpCounter uint64
	metadataLocks      = newWorkspaceMetadataLocker()
)

type workspaceMetadataLocker struct {
	mu    sync.Mutex
	locks map[string]*workspaceMetadataLock
}

type workspaceMetadataLock struct {
	ch   chan struct{}
	refs int
}

func newWorkspaceMetadataLocker() *workspaceMetadataLocker { _ = "STUB: not implemented"; return nil }

// MetadataTempFileName returns a unique workspace-relative temporary file
// name suitable for atomically replacing metadata.json.
func MetadataTempFileName() string { _ = "STUB: not implemented"; return "" }

// NewWorkspaceMetadata returns a metadata value initialized with defaults.
func NewWorkspaceMetadata() WorkspaceMetadata {
	_ = "STUB: not implemented"
	return *new(WorkspaceMetadata)
}

// IsMetadataCorruptError reports whether err came from decoding workspace
// metadata JSON.
func IsMetadataCorruptError(err error) bool { _ = "STUB: not implemented"; return false }

// IsMetadataTempFileName reports whether name is a metadata temp file created
// by MetadataTempFileName.
func IsMetadataTempFileName(name string) bool { _ = "STUB: not implemented"; return false }

// IsRootMetadataTempPath reports whether rel identifies a root-level
// workspace metadata temp file.
func IsRootMetadataTempPath(rel string) bool { _ = "STUB: not implemented"; return false }

func metadataRandomSuffix() string { _ = "STUB: not implemented"; return "" }

func isMetadataRandomSuffix(s string) bool { _ = "STUB: not implemented"; return false }

// WithWorkspaceMetadataLock serializes metadata read-modify-write operations
// for the same workspace within this process. The callback should keep the
// critical section limited to metadata load, mutation, and save.
func WithWorkspaceMetadataLock(
	ctx context.Context,
	root string,
	fn func(context.Context) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func workspaceMetadataLockKey(root string) string { _ = "STUB: not implemented"; return "" }

func (k *workspaceMetadataLocker) lock(
	ctx context.Context,
	key string,
) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *workspaceMetadataLocker) releaseRef(
	key string,
	kl *workspaceMetadataLock,
) {
	_ = "STUB: not implemented"
	return
}

// WorkspaceMetadata describes staged skills and recent activity.
type WorkspaceMetadata struct {
	Version    int                  `json:"version"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	LastAccess time.Time            `json:"last_access"`
	Skills     map[string]SkillMeta `json:"skills"`
	Inputs     []InputRecord        `json:"inputs,omitempty"`
	Outputs    []OutputRecord       `json:"outputs,omitempty"`
	// Prepared records the last-known converged state for each
	// workspace requirement keyed by Requirement.Key(). It is used by
	// the workspaceprep reconciler to skip work whose fingerprint is
	// unchanged and whose sentinel (for example the target file) is
	// still present. The map is a local per-workspace cache; session
	// state remains the authoritative source of "what should exist".
	Prepared map[string]PreparedRecord `json:"prepared,omitempty"`
}

// PreparedRecord captures a single successfully-applied workspace
// requirement. It is written by the reconciler after a successful
// apply and read on subsequent reconciles to decide whether to skip.
type PreparedRecord struct {
	Key         string    `json:"key"`
	Kind        string    `json:"kind"`
	Fingerprint string    `json:"fingerprint"`
	Target      string    `json:"target,omitempty"`
	PreparedAt  time.Time `json:"prepared_at"`
}

// SkillMeta records a staged skill snapshot.
type SkillMeta struct {
	Name     string    `json:"name"`
	RelPath  string    `json:"rel_path"`
	Digest   string    `json:"digest"`
	Mounted  bool      `json:"mounted"`
	StagedAt time.Time `json:"staged_at"`
}

// InputRecord tracks a staged input resolution.
type InputRecord struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Resolved  string    `json:"resolved,omitempty"`
	Version   *int      `json:"version,omitempty"`
	Mode      string    `json:"mode,omitempty"`
	Timestamp time.Time `json:"ts"`
}

// OutputRecord tracks an output collection run.
type OutputRecord struct {
	Globs     []string  `json:"globs"`
	SavedAs   []string  `json:"saved_as,omitempty"`
	Versions  []int     `json:"versions,omitempty"`
	LimitsHit bool      `json:"limits_hit"`
	Timestamp time.Time `json:"ts"`
}

// EnsureLayout creates standard workspace subdirectories and a
// metadata file when absent. It returns full paths for convenience.
func EnsureLayout(root string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize metadata if missing.

// LoadMetadata loads metadata.json from workspace root. When missing,
// an empty metadata with defaults is returned without error.
func LoadMetadata(root string) (WorkspaceMetadata, error) {
	_ = "STUB: not implemented"
	return *new(WorkspaceMetadata), nil
}

// SaveMetadata writes metadata.json to the workspace root.
func SaveMetadata(root string, md WorkspaceMetadata) error { _ = "STUB: not implemented"; return nil }

// DirDigest computes a stable digest of a directory tree. It walks
// the tree, sorts entries, and hashes relative path and contents.
func DirDigest(root string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Normalize to slash for stability.
