//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package agent

import (
	"trpc.group/trpc-go/trpc-agent-go/internal/surfacepatch"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// SurfacePatch represents one node's runtime surface overrides.
type SurfacePatch struct {
	patch surfacepatch.Patch
}

// SetInstruction sets the instruction surface override.
func (p *SurfacePatch) SetInstruction(text string) { _ = "STUB: not implemented"; return }

// SetGlobalInstruction sets the global instruction surface override.
func (p *SurfacePatch) SetGlobalInstruction(text string) { _ = "STUB: not implemented"; return }

// SetFewShot sets the few-shot surface override.
func (p *SurfacePatch) SetFewShot(examples [][]model.Message) { _ = "STUB: not implemented"; return }

// SetModel sets the model surface override.
func (p *SurfacePatch) SetModel(m model.Model) { _ = "STUB: not implemented"; return }

// SetTools sets the tool surface override and clears appended tools.
func (p *SurfacePatch) SetTools(tools []tool.Tool) { _ = "STUB: not implemented"; return }

// AppendTools appends tools to the node's runtime tool surface.
func (p *SurfacePatch) AppendTools(tools []tool.Tool) { _ = "STUB: not implemented"; return }

// SetSkillRepository sets the skill repository surface override.
func (p *SurfacePatch) SetSkillRepository(repo skill.Repository) { _ = "STUB: not implemented"; return }

// WithSurfacePatchForNode applies one node's runtime surface overrides to this run.
func WithSurfacePatchForNode(nodeID string, patch SurfacePatch) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}
