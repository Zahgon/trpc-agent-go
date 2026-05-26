//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package structure provides internal helpers for static structure export.
package structure

import (
	"trpc.group/trpc-go/trpc-agent-go/agent/structure"
)

// PathAllocator allocates stable child paths under one parent node.
type PathAllocator struct {
	parentNodeID string
	used         map[string]int
}

// NewPathAllocator creates a new path allocator for one parent node.
func NewPathAllocator(parentNodeID string) *PathAllocator { _ = "STUB: not implemented"; return nil }

// Next returns the next stable child path for the given local name.
func (a *PathAllocator) Next(localName string) string { _ = "STUB: not implemented"; return "" }

// EscapeLocalName escapes one path segment into a stable node-id segment.
func EscapeLocalName(name string) string { _ = "STUB: not implemented"; return "" }

// JoinNodeID joins a parent node id and a local name into a child node id.
func JoinNodeID(parentNodeID string, localName string) string { _ = "STUB: not implemented"; return "" }

// RebaseSnapshot rewrites one snapshot to a new mounted root node id.
func RebaseSnapshot(
	snapshot *structure.Snapshot,
	newRootNodeID string,
) (*structure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TerminalNodeIDs returns the static terminal node ids of a snapshot.
func TerminalNodeIDs(snapshot *structure.Snapshot) []string { _ = "STUB: not implemented"; return nil }

// RootOnly returns a snapshot that keeps only the root node and its surfaces.
func RootOnly(snapshot *structure.Snapshot) *structure.Snapshot {
	_ = "STUB: not implemented"
	return nil
}

func joinEscapedNodeID(parentNodeID string, escaped string) string {
	_ = "STUB: not implemented"
	return ""
}

func rebaseNodeID(nodeID string, oldRoot string, newRoot string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func reachableNodeIDs(
	entryNodeID string,
	nodeIDs map[string]struct{},
	adjacency map[string][]string,
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
