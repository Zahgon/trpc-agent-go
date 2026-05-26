//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package repo provides repository-based knowledge source implementation.
package repo

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

const defaultRepoSourceName = "Repository Source"

// Source represents a knowledge source for a code repository.
type Source struct {
	repository     Repository
	hasRepository  bool
	multiRepoError bool
	name           string
	metadata       map[string]any
	readers        map[string]reader.Reader
	fileExtensions []string
	recursive      bool
	transformers   []transform.Transformer
	skipDirs       []string
	skipSuffixes   []string
}

// Repository describes one repository input and its version/scope configuration.
type Repository struct {
	URL         string
	Dir         string
	Branch      string
	Tag         string
	Commit      string
	Subdir      string
	RepoName    string
	Description string
	RepoURL     string
}

// New creates a new repository knowledge source.
func New(opts ...Option) *Source { _ = "STUB: not implemented"; return nil }

func (s *Source) initializeReaders() { _ = "STUB: not implemented"; return }

// ReadDocuments reads all repository inputs and returns documents.
func (s *Source) ReadDocuments(ctx context.Context) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Step 1: directory-level language parsers (e.g. Go).

// Step 2: code language file parsers (e.g. Proto) – before plain-text readers.

// Step 3: plain-text / doc file readers (e.g. md, txt).

// fileClassification groups file paths by processing priority, matching trpc-ast-rag order:
// directory-level parsers first, code-file parsers second, plain-text readers last.
type fileClassification struct {
	dirTypes      []string                       // file types with directoryReader (e.g. "go"), sorted
	codeFiles     []string                       // code language files (e.g. .proto), sorted
	textFiles     []string                       // plain-text/doc files (e.g. .md, .txt), sorted
	allowedByType map[string]map[string]struct{} // allowed repo-relative paths per fileType
}

func resolveScanRoot(repoRoot, subdir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Source) classifyFiles(repoRoot string, filePaths []string) (*fileClassification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Source) resolvedRepository() (Repository, bool) {
	_ = "STUB: not implemented"
	return *new(Repository), false
}

// Name returns the source name.
func (s *Source) Name() string {
	_ = "STUB: not implemented"

	// Type returns the source type.
	return ""
}

func (s *Source) Type() string { _ = "STUB: not implemented"; return "" }

// GetMetadata returns source metadata.
func (s *Source) GetMetadata() map[string]any { _ = "STUB: not implemented"; return nil }

// RepositoryDescriptor exposes the configured repository name/description for
// tool-layer prompt construction. It is intentionally source-level metadata and
// is not copied into per-chunk document metadata.
func (s *Source) RepositoryDescriptor() (name, description string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

type repoInfo struct {
	name   string
	url    string
	branch string
}

type checkoutTargetKind string

const (
	checkoutTargetDefault checkoutTargetKind = "default"
	checkoutTargetBranch  checkoutTargetKind = "branch"
	checkoutTargetTag     checkoutTargetKind = "tag"
	checkoutTargetCommit  checkoutTargetKind = "commit"
)

type directoryReader interface {
	ReadFromDirectory(dirPath string) ([]*document.Document, error)
}

// isCodeFileType returns true for file types handled by language/code parsers
// (but not via directoryReader). These run before plain-text file readers to
// match trpc-ast-rag's processing order: language parsers first, text readers last.
func isCodeFileType(fileType string) bool { _ = "STUB: not implemented"; return false }

func (s *Source) resolveRepository(ctx context.Context, repository Repository) (string, *repoInfo, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil, nil
}

func resolveCheckoutTarget(repository Repository) (checkoutTargetKind, string) {
	_ = "STUB: not implemented"
	return *new(checkoutTargetKind), ""
}

func cloneRemoteRepository(ctx context.Context, repository Repository, tmpDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Source) getFilePaths(dirPath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip git submodule directories that have not been checked out.
// Populated submodules (with actual source files) are scanned normally.

// isUnpopulatedGitLink reports whether dirPath is a git submodule directory
// that has not been checked out. Such directories contain a .git pointer file
// (not the usual .git directory) but no actual source files.
// Fully cloned submodules have a .git file plus real source files and are
// allowed through.
func isUnpopulatedGitLink(dirPath string) bool { _ = "STUB: not implemented"; return false }

// No .git entry, or .git is a directory (a regular repo root) → not a submodule link.

// .git is a file → this directory is a git submodule.
// If it contains nothing else, the submodule has not been checked out.

// has other content → submodule is populated

// buildBaseMetadata returns common metadata shared by all documents produced from this source.
func (s *Source) buildBaseMetadata(repoRoot string, info *repoInfo) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (s *Source) processFile(filePath, repoRoot string, info *repoInfo) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Source) processDirectory(dirPath, fileType, repoRoot string, info *repoInfo, allowedPaths map[string]struct{}) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toRelativeRepoPath(repoRoot string, raw any) string { _ = "STUB: not implemented"; return "" }

func (s *Source) shouldSkipDir(name string) bool { _ = "STUB: not implemented"; return false }

func (s *Source) shouldSkipFile(name string) bool { _ = "STUB: not implemented"; return false }

func looksLikeGitURL(input string) bool { _ = "STUB: not implemented"; return false }

func runGit(ctx context.Context, dir string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func chooseRepoName(explicit, rawInput, fallbackPath string) string {
	_ = "STUB: not implemented"
	return ""
}

func chooseRepoURL(explicit, input string) string { _ = "STUB: not implemented"; return "" }

// firstNonEmpty returns the first non-empty string from vals.
func firstNonEmpty(vals ...string) string { _ = "STUB: not implemented"; return "" }
