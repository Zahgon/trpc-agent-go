//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package file

import (
	"context"
)

type workspaceIndex struct {
	files []string
	dirs  []string
}

func buildWorkspaceIndex(ctx context.Context) workspaceIndex {
	_ = "STUB: not implemented"
	return *new(workspaceIndex)
}

func matchWorkspacePaths(
	ctx context.Context,
	dir string,
	pattern string,
	caseSensitive bool,
) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func matchWorkspacePattern(
	pattern string,
	name string,
	caseSensitive bool,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
