//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package evalset

// defaultEvalSetFileSuffix is the default suffix for eval set files.
const defaultEvalSetFileSuffix = ".evalset.json"

// Locator provides Build and List methods for locating eval set files.
type Locator interface {
	// Build builds the path of an eval set file for the given appName and evalSetID.
	Build(baseDir, appName, evalSetID string) string
	// List lists all eval set IDs for the given appName.
	List(baseDir, appName string) ([]string, error)
}

// locator is the default Locator implementation.
type locator struct {
}

// Build builds the path of an eval set file.
func (l *locator) Build(baseDir, appName, evalSetID string) string {
	_ = "STUB: not implemented"
	return ""
}

// List lists all eval set IDs for the given appName.
func (l *locator) List(baseDir, appName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
