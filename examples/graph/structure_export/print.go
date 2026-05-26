//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/agent/structure"
)

func printSnapshot(snapshot *structure.Snapshot) { _ = "STUB: not implemented"; return }

func formatSurfaceValue(surface structure.Surface) string { _ = "STUB: not implemented"; return "" }

func printHighlights(rootNodeID string, edges []structure.Edge) { _ = "STUB: not implemented"; return }

func printBranchPoints(outgoing map[string][]string) bool { _ = "STUB: not implemented"; return false }

func printFanInPoints(rootNodeID string, incoming map[string][]string) bool {
	_ = "STUB: not implemented"
	return false
}

func printLoopRegions(edges []structure.Edge, outgoing map[string][]string) bool {
	_ = "STUB: not implemented"
	return false
}

func canReach(outgoing map[string][]string, start string, target string) bool {
	_ = "STUB: not implemented"
	return false
}

func filterSources(sources []string, ignored string) []string {
	_ = "STUB: not implemented"
	return nil
}

func findStronglyConnectedComponents(
	edges []structure.Edge,
	outgoing map[string][]string,
) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func dfsFinishOrder(
	nodeID string,
	outgoing map[string][]string,
	visited map[string]bool,
	order *[]string,
) {
	_ = "STUB: not implemented"
	return
}

func dfsCollectComponent(
	nodeID string,
	reverse map[string][]string,
	visited map[string]bool,
	component *[]string,
) {
	_ = "STUB: not implemented"
	return
}

func isLoopComponent(component []string, outgoing map[string][]string) bool {
	_ = "STUB: not implemented"
	return false
}
