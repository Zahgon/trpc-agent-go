//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package skills

import (
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	watchRefreshReasonManual = "manual"
	watchRefreshReasonWatch  = "watch"
)

var defaultWatchIgnoredNames = map[string]struct{}{
	".cache":        {},
	".git":          {},
	".mypy_cache":   {},
	".pytest_cache": {},
	".venv":         {},
	"__pycache__":   {},
	"build":         {},
	"dist":          {},
	"node_modules":  {},
	"venv":          {},
}

type WatchConfig struct {
	Enabled      bool
	Debounce     time.Duration
	WatchBundled bool
	BundledRoot  string
}

type WatchStatus struct {
	Enabled           bool       `json:"enabled"`
	WatchBundled      bool       `json:"watch_bundled"`
	DebounceMS        int        `json:"debounce_ms,omitempty"`
	Roots             []string   `json:"roots,omitempty"`
	Generation        int64      `json:"generation,omitempty"`
	LastRefreshAt     *time.Time `json:"last_refresh_at,omitempty"`
	LastRefreshReason string     `json:"last_refresh_reason,omitempty"`
	LastChangedPath   string     `json:"last_changed_path,omitempty"`
	LastError         string     `json:"last_error,omitempty"`
}

type WatchService struct {
	repo *Repository
	cfg  WatchConfig

	stateMu sync.RWMutex
	watchMu sync.RWMutex

	watcher *fsnotify.Watcher
	roots   []string
	watched map[string]struct{}

	done chan struct{}
	wg   sync.WaitGroup

	refreshMu sync.Mutex

	generation        int64
	lastRefreshAt     time.Time
	lastRefreshReason string
	lastChangedPath   string
	lastError         string
}

func NewWatchService(
	repo *Repository,
	roots []string,
	cfg WatchConfig,
) *WatchService {
	_ = "STUB: not implemented"
	return nil
}

func (s *WatchService) Close() error { _ = "STUB: not implemented"; return nil }

func (s *WatchService) Refresh() error { _ = "STUB: not implemented"; return nil }

func (s *WatchService) Status() *WatchStatus { _ = "STUB: not implemented"; return nil }

func (s *WatchService) loop() { _ = "STUB: not implemented"; return }

func (s *WatchService) relevantEvent(
	event fsnotify.Event,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (s *WatchService) refresh(
	reason string,
	changedPath string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *WatchService) syncWatches() error { _ = "STUB: not implemented"; return nil }

func (s *WatchService) desiredWatchDirs() map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func collectWatchDirs(root string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func (s *WatchService) recordError(err error) { _ = "STUB: not implemented"; return }

func normalizeWatchRoots(
	roots []string,
	cfg WatchConfig,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func normalizeWatchPath(raw string) string { _ = "STUB: not implemented"; return "" }

func nearestExistingWatchParent(path string) string { _ = "STUB: not implemented"; return "" }

func isIgnoredWatchPath(path string) bool { _ = "STUB: not implemented"; return false }

func isIgnoredWatchName(name string) bool { _ = "STUB: not implemented"; return false }

func isIgnorableWatchError(err error) bool { _ = "STUB: not implemented"; return false }

func watchDirExists(path string) bool { _ = "STUB: not implemented"; return false }
