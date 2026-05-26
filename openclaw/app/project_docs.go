//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

const (
	projectDocFileName     = "AGENTS.md"
	projectDocOverrideName = "AGENTS.override.md"
	projectDocGitMarker    = ".git"
	projectDocMaxBytes     = 16 * 1024
)

var projectDocFileNames = []string{
	projectDocFileName,
	projectDocOverrideName,
}

func resolveProjectDocs(cwd string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func discoverProjectDocPaths(cwd string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func projectDocDirs(cwd string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func hasProjectDocRootMarker(dir string) bool { _ = "STUB: not implemented"; return false }

func readTrimmedTextFile(path string, limit int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func reverseStrings(values []string) { _ = "STUB: not implemented"; return }
