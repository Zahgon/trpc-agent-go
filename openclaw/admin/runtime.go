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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/octool"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
)

const (
	defaultUploadsDir = "uploads"
	maxUploadRows     = 12
	maxUploadSessions = 8
	maxExecRows       = 12
)

type uploadFilters struct {
	Channel   string
	UserID    string
	SessionID string
	Kind      string
	MimeType  string
	Source    string
}

type execStatus struct {
	Enabled      bool              `json:"enabled"`
	SessionCount int               `json:"session_count"`
	RunningCount int               `json:"running_count"`
	Sessions     []execSessionView `json:"sessions,omitempty"`
}

type execSessionView struct {
	SessionID string `json:"session_id,omitempty"`
	Command   string `json:"command,omitempty"`
	Status    string `json:"status,omitempty"`
	StartedAt string `json:"started_at,omitempty"`
	DoneAt    string `json:"done_at,omitempty"`
	ExitCode  *int   `json:"exit_code,omitempty"`
}

type uploadsStatus struct {
	Enabled      bool                `json:"enabled"`
	Root         string              `json:"root,omitempty"`
	FileCount    int                 `json:"file_count"`
	TotalBytes   int64               `json:"total_bytes"`
	Error        string              `json:"error,omitempty"`
	KindCounts   []uploadKindCount   `json:"kind_counts,omitempty"`
	SourceCounts []uploadSourceCount `json:"source_counts,omitempty"`
	Files        []uploadView        `json:"files,omitempty"`
	Sessions     []uploadSessionView `json:"sessions,omitempty"`
}

type uploadKindCount struct {
	Kind  string `json:"kind,omitempty"`
	Count int    `json:"count"`
}

type uploadSourceCount struct {
	Source string `json:"source,omitempty"`
	Count  int    `json:"count"`
}

type uploadView struct {
	Channel      string    `json:"channel,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
	SessionID    string    `json:"session_id,omitempty"`
	Name         string    `json:"name,omitempty"`
	RelativePath string    `json:"relative_path,omitempty"`
	Kind         string    `json:"kind,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	Source       string    `json:"source,omitempty"`
	SizeBytes    int64     `json:"size_bytes"`
	ModifiedAt   time.Time `json:"modified_at,omitempty"`
	OpenURL      string    `json:"open_url,omitempty"`
	DownloadURL  string    `json:"download_url,omitempty"`
}

type uploadSessionView struct {
	Channel      string    `json:"channel,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
	SessionID    string    `json:"session_id,omitempty"`
	FileCount    int       `json:"file_count"`
	TotalBytes   int64     `json:"total_bytes"`
	LastModified time.Time `json:"last_modified,omitempty"`
}

func (s *Service) execStatus() execStatus { _ = "STUB: not implemented"; return *new(execStatus) }

func execSessionViewFromSession(
	session octool.ProcessSession,
) execSessionView {
	_ = "STUB: not implemented"
	return *new(execSessionView)
}

func (s *Service) uploadsStatus() uploadsStatus {
	_ = "STUB: not implemented"
	return *new(uploadsStatus)
}

func (s *Service) uploadsStatusFiltered(
	filters uploadFilters,
	fileLimit int,
	sessionLimit int,
) uploadsStatus {
	_ = "STUB: not implemented"
	return *new(uploadsStatus)
}

func resolveUploadsRoot(stateDir string) string { _ = "STUB: not implemented"; return "" }

func uploadViewsFromList(
	listed []uploads.ListedFile,
	limit int,
) ([]uploadView, int64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func uploadSessionsFromList(
	listed []uploads.ListedFile,
	limit int,
) []uploadSessionView {
	_ = "STUB: not implemented"
	return nil
}

func uploadKindCountsFromList(
	listed []uploads.ListedFile,
) []uploadKindCount {
	_ = "STUB: not implemented"
	return nil
}

func uploadSourceCountsFromList(
	listed []uploads.ListedFile,
) []uploadSourceCount {
	_ = "STUB: not implemented"
	return nil
}

func filterUploadList(
	listed []uploads.ListedFile,
	filters uploadFilters,
) []uploads.ListedFile {
	_ = "STUB: not implemented"
	return nil
}

func errorsIsNotExist(err error) bool { _ = "STUB: not implemented"; return false }

func uploadKindFromName(name string) string { _ = "STUB: not implemented"; return "" }

func uploadKindFromFile(file uploads.ListedFile) string { _ = "STUB: not implemented"; return "" }

func uploadFileURL(rel string, download bool) string { _ = "STUB: not implemented"; return "" }
