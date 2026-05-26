//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package surfacepatch stores runtime node surface patches in invocation configs.
package surfacepatch

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	configsKey           = "__trpc_agent_internal_node_surface_patches__"
	rootNodeIDConfigsKey = "__trpc_agent_internal_surface_root_node_id__"
)

type textSlot struct {
	set   bool
	value string
}

type fewShotSlot struct {
	set   bool
	value [][]model.Message
}

type modelSlot struct {
	set   bool
	value model.Model
}

type toolsSlot struct {
	set    bool
	value  []tool.Tool
	append []tool.Tool
}

type skillRepoSlot struct {
	set   bool
	value skill.Repository
}

// Patch represents one node's runtime surface overrides.
type Patch struct {
	instruction       textSlot
	globalInstruction textSlot
	fewShot           fewShotSlot
	model             modelSlot
	tools             toolsSlot
	skillRepo         skillRepoSlot
}

// SetInstruction sets the instruction surface override.
func (p *Patch) SetInstruction(text string) { _ = "STUB: not implemented"; return }

// SetGlobalInstruction sets the global instruction surface override.
func (p *Patch) SetGlobalInstruction(text string) { _ = "STUB: not implemented"; return }

// SetFewShot sets the few-shot surface override.
func (p *Patch) SetFewShot(examples [][]model.Message) { _ = "STUB: not implemented"; return }

// SetModel sets the model surface override.
func (p *Patch) SetModel(m model.Model) { _ = "STUB: not implemented"; return }

// SetTools sets the tool surface override and clears appended tools.
func (p *Patch) SetTools(tools []tool.Tool) { _ = "STUB: not implemented"; return }

// AppendTools appends tools to the effective tool surface.
func (p *Patch) AppendTools(tools []tool.Tool) { _ = "STUB: not implemented"; return }

// SetSkillRepository sets the skill repository surface override.
func (p *Patch) SetSkillRepository(repo skill.Repository) { _ = "STUB: not implemented"; return }

// Instruction returns the instruction surface override.
func (p Patch) Instruction() (string, bool) { _ = "STUB: not implemented"; return "", false }

// GlobalInstruction returns the global instruction surface override.
func (p Patch) GlobalInstruction() (string, bool) { _ = "STUB: not implemented"; return "", false }

// FewShot returns the few-shot surface override.
func (p Patch) FewShot() ([][]model.Message, bool) { _ = "STUB: not implemented"; return nil, false }

// Model returns the model surface override.
func (p Patch) Model() (model.Model, bool) {
	_ = "STUB: not implemented"
	return *new(model.Model), false
}

// Tools returns the tool surface override.
func (p Patch) Tools() ([]tool.Tool, bool) { _ = "STUB: not implemented"; return nil, false }

// ApplyTools returns the effective tool surface after applying this patch.
func (p Patch) ApplyTools(base []tool.Tool) ([]tool.Tool, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SkillRepository returns the skill repository surface override.
func (p Patch) SkillRepository() (skill.Repository, bool) {
	_ = "STUB: not implemented"
	return *new(skill.Repository), false
}

// IsEmpty reports whether the patch carries any surface override.
func (p Patch) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Merge returns a copy where values from other override the same surface type.
func (p Patch) Merge(other Patch) Patch { _ = "STUB: not implemented"; return *new(Patch) }

// Clone returns a defensive copy of the patch.
func (p Patch) Clone() Patch { _ = "STUB: not implemented"; return *new(Patch) }

// WithPatch stores a node patch in custom configs.
func WithPatch(cfgs map[string]any, nodeID string, patch Patch) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// PatchForNode returns the merged patch for one node id.
func PatchForNode(cfgs map[string]any, nodeID string) (Patch, bool) {
	_ = "STUB: not implemented"
	return *new(Patch), false
}

// WithRootNodeID stores one invocation's surface lookup root node id.
func WithRootNodeID(cfgs map[string]any, nodeID string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// RootNodeID returns the surface lookup root node id from configs when present.
func RootNodeID(cfgs map[string]any, fallback string) string { _ = "STUB: not implemented"; return "" }

type nodePatches map[string]Patch

func nodePatchesFromConfigs(cfgs map[string]any) nodePatches {
	_ = "STUB: not implemented"
	return *new(nodePatches)
}

func cloneNodePatches(in nodePatches) nodePatches {
	_ = "STUB: not implemented"
	return *new(nodePatches)
}

func copyConfigs(in map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func cloneFewShot(in [][]model.Message) [][]model.Message { _ = "STUB: not implemented"; return nil }

func cloneTools(in []tool.Tool) []tool.Tool { _ = "STUB: not implemented"; return nil }

func appendTools(base []tool.Tool, appended []tool.Tool) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}
