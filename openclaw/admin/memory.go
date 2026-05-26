//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package admin

import (
	"context"
	"time"
)

const (
	memoryFileName            = "MEMORY.md"
	maxMemoryFilePreviewBytes = 4 * 1024
	maxMemoryFilePreviewRunes = 220
	maxMemoryPreviewLines     = 3

	memoryCardIDPrefix      = "memory-file-"
	memoryFilePerm          = 0o600
	memoryTempPatternSuffix = ".tmp-*"
)

type MemoryFileStore interface {
	Root() string
	ReadFile(path string, maxBytes int) (string, error)
}

type MemoryUserLabelResolver interface {
	ResolveMemoryUserLabel(appName string, userID string) string
}

type MemoryUserLabelResolverFunc func(
	appName string,
	userID string,
) string

func (f MemoryUserLabelResolverFunc) ResolveMemoryUserLabel(
	appName string,
	userID string,
) string {
	_ = "STUB: not implemented"
	return ""
}

type memoryFileSaver interface {
	SaveResolvedMemoryFile(
		ctx context.Context,
		path string,
		content string,
	) error
}

type memoryStatus struct {
	Enabled      bool             `json:"enabled"`
	FileEnabled  bool             `json:"file_enabled"`
	Backend      string           `json:"backend,omitempty"`
	Root         string           `json:"root,omitempty"`
	FileCount    int              `json:"file_count"`
	TotalBytes   int64            `json:"total_bytes"`
	LastModified *time.Time       `json:"last_modified,omitempty"`
	Error        string           `json:"error,omitempty"`
	Files        []memoryFileView `json:"files,omitempty"`
}

type memoryFileView struct {
	AppName      string    `json:"app_name,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
	UserLabel    string    `json:"user_label,omitempty"`
	RelativePath string    `json:"relative_path,omitempty"`
	Path         string    `json:"path,omitempty"`
	OpenURL      string    `json:"open_url,omitempty"`
	LoadURL      string    `json:"load_url,omitempty"`
	CardID       string    `json:"card_id,omitempty"`
	SearchValue  string    `json:"search_value,omitempty"`
	Preview      string    `json:"preview,omitempty"`
	SizeBytes    int64     `json:"size_bytes"`
	ModifiedAt   time.Time `json:"modified_at"`
}

type memoryFileDetail struct {
	AppName      string    `json:"app_name,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
	UserLabel    string    `json:"user_label,omitempty"`
	RelativePath string    `json:"relative_path,omitempty"`
	OpenURL      string    `json:"open_url,omitempty"`
	LoadURL      string    `json:"load_url,omitempty"`
	Content      string    `json:"content,omitempty"`
	SizeBytes    int64     `json:"size_bytes"`
	ModifiedAt   time.Time `json:"modified_at"`
}

func (s *Service) memoryStatus() memoryStatus { _ = "STUB: not implemented"; return *new(memoryStatus) }

func (s *Service) memoryStatusSummary() memoryStatus {
	_ = "STUB: not implemented"
	return *new(memoryStatus)
}

func (s *Service) memoryStatusWithFiles(includeFiles bool) memoryStatus {
	_ = "STUB: not implemented"
	return *new(memoryStatus)
}

func configuredMemoryRoot(
	store MemoryFileStore,
) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func memoryFileViews(
	store MemoryFileStore,
	includePreview bool,
) ([]memoryFileView, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func memoryFileViewsWithResolver(
	store MemoryFileStore,
	resolver MemoryUserLabelResolver,
	includePreview bool,
) ([]memoryFileView, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeMemoryPathPart(part string) string { _ = "STUB: not implemented"; return "" }

func summarizeMemoryPreview(
	text string,
	maxLines int,
	maxRunes int,
) string {
	_ = "STUB: not implemented"
	return ""
}

func resolveMemoryFile(root string, relPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func normalizeMemoryRelativePath(relPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readMemoryFileDetail(
	root string,
	relPath string,
) (memoryFileDetail, error) {
	_ = "STUB: not implemented"
	return *new(memoryFileDetail), nil
}

func readMemoryFileDetailWithResolver(
	root string,
	relPath string,
	resolver MemoryUserLabelResolver,
) (memoryFileDetail, error) {
	_ = "STUB: not implemented"
	return *new(memoryFileDetail), nil
}

func saveMemoryFile(
	ctx context.Context,
	store MemoryFileStore,
	relPath string,
	content string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func memoryScopeFromRelativePath(relPath string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func buildMemorySearchValue(
	appName string,
	userID string,
	userLabel string,
	relPath string,
	preview string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func resolveMemoryUserLabel(
	resolver MemoryUserLabelResolver,
	appName string,
	userID string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func memoryCardID(relPath string) string { _ = "STUB: not implemented"; return "" }

func writeMemoryFileAtomic(path string, data []byte) error { _ = "STUB: not implemented"; return nil }
