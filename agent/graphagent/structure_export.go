//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package graphagent

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent/structure"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Export exports the static structure of the graph agent.
func (ga *GraphAgent) Export(
	ctx context.Context,
	exportChild structure.ChildExporter,
) (*structure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendAgentNodeChildSnapshot(
	snapshot *structure.Snapshot,
	rebased *structure.Snapshot,
) {
	_ = "STUB: not implemented"
	return
}

func nodeKindFromGraphNodeType(nodeType graph.NodeType) structure.NodeKind {
	_ = "STUB: not implemented"
	return *new(structure.NodeKind)
}

func exportGraphNodeSurfaces(
	ctx context.Context,
	node *graph.Node,
	nodeID string,
) []structure.Surface {
	_ = "STUB: not implemented"
	return nil
}

func toolRefsFromTools(tools []tool.Tool) []structure.ToolRef {
	_ = "STUB: not implemented"
	return nil
}

func stringPtr(value string) *string { _ = "STUB: not implemented"; return nil }

func collectConditionalTargets(g *graph.Graph, node *graph.Node) []string {
	_ = "STUB: not implemented"
	return nil
}

func uniqueTargets(targets []string) []string { _ = "STUB: not implemented"; return nil }
