//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package structure

import (
	"context"
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var errNilAgent = errors.New("agent is nil")

type exportState struct {
	stack []agent.Agent
}

// Export exports a normalized static structure snapshot for the given agent.
func Export(ctx context.Context, a agent.Agent) (*Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exportWithState(
	ctx context.Context,
	a agent.Agent,
	state *exportState,
) (*Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *exportState) exportChild(
	ctx context.Context,
	a agent.Agent,
) (*Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeSnapshot(raw *Snapshot) (*Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type edgeKey struct {
	from string
	to   string
}

func cloneSnapshot(raw *Snapshot) *Snapshot { _ = "STUB: not implemented"; return nil }

func cloneSurfaceValue(value SurfaceValue) SurfaceValue {
	_ = "STUB: not implemented"
	return *new(SurfaceValue)
}

func cloneToolRefs(refs []ToolRef) []ToolRef { _ = "STUB: not implemented"; return nil }

func cloneToolSchema(schema *tool.Schema) *tool.Schema { _ = "STUB: not implemented"; return nil }

func cloneSchemaMap(in map[string]*tool.Schema) map[string]*tool.Schema {
	_ = "STUB: not implemented"
	return nil
}

func cloneSchemaValues(in []any) []any { _ = "STUB: not implemented"; return nil }

func cloneSchemaValue(value any) any { _ = "STUB: not implemented"; return *new(any) }

func cloneFewShot(value []FewShotExample) []FewShotExample { _ = "STUB: not implemented"; return nil }

func normalizeSurfaceValue(value SurfaceValue) SurfaceValue {
	_ = "STUB: not implemented"
	return *new(SurfaceValue)
}

func uniqueToolRefs(refs []ToolRef) []ToolRef { _ = "STUB: not implemented"; return nil }

func uniqueSkillRefs(refs []SkillRef) []SkillRef { _ = "STUB: not implemented"; return nil }

func opaqueLeafSnapshot(a agent.Agent) *Snapshot { _ = "STUB: not implemented"; return nil }

func escapeNodeIDSegment(name string) string { _ = "STUB: not implemented"; return "" }

func validateSurfaceValue(surfaceType SurfaceType, value SurfaceValue) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTextSurfaceValue(value SurfaceValue) error { _ = "STUB: not implemented"; return nil }

func validateFewShotSurfaceValue(value SurfaceValue) error { _ = "STUB: not implemented"; return nil }

func validateModelSurfaceValue(value SurfaceValue) error { _ = "STUB: not implemented"; return nil }

func validateToolSurfaceValue(value SurfaceValue) error { _ = "STUB: not implemented"; return nil }

func validateSkillSurfaceValue(value SurfaceValue) error { _ = "STUB: not implemented"; return nil }

func (s *exportState) push(a agent.Agent) { _ = "STUB: not implemented"; return }

func (s *exportState) pop() { _ = "STUB: not implemented"; return }

func (s *exportState) containsRecursiveAgentInstance(a agent.Agent) bool {
	_ = "STUB: not implemented"
	return false
}

func samePointerAgentInstance(left agent.Agent, right agent.Agent) bool {
	_ = "STUB: not implemented"
	return false
}

func isNilAgent(a agent.Agent) bool { _ = "STUB: not implemented"; return false }
