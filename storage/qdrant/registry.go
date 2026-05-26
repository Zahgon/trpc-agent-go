//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package qdrant

import "sync"

var (
	registryMu     sync.RWMutex
	qdrantRegistry = make(map[string][]ClientBuilderOpt)
)

// RegisterQdrantInstance registers a named Qdrant instance with its configuration options.
// If an instance with the same name already exists, it will be overwritten.
func RegisterQdrantInstance(name string, opts ...ClientBuilderOpt) {
	_ = "STUB: not implemented"
	return
}

// GetQdrantInstance retrieves the configuration options for a named Qdrant instance.
// Returns a copy of the options and true if found, or nil and false if not found.
func GetQdrantInstance(name string) ([]ClientBuilderOpt, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// copy to prevent external modifications

// UnregisterQdrantInstance removes a named Qdrant instance from the registry.
func UnregisterQdrantInstance(name string) { _ = "STUB: not implemented"; return }

// ListQdrantInstances returns a list of all registered instance names.
func ListQdrantInstances() []string { _ = "STUB: not implemented"; return nil }
