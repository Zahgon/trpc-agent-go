//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package llmagent

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent/structure"
	"trpc.group/trpc-go/trpc-agent-go/prompt"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Export exports the static structure of the LLM agent.
func (a *LLMAgent) Export(
	ctx context.Context,
	exportChild structure.ChildExporter,
) (*structure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exportToolRefs(tools []tool.Tool) []structure.ToolRef { _ = "STUB: not implemented"; return nil }

func exportSkillRefs(summaries []skill.Summary) []structure.SkillRef {
	_ = "STUB: not implemented"
	return nil
}

func stringPtr(value string) *string { _ = "STUB: not implemented"; return nil }

func exportTextSurfaceValue(text prompt.Text) structure.SurfaceValue {
	_ = "STUB: not implemented"
	return *new(structure.SurfaceValue)
}

func promptSyntaxPtr(value structure.PromptSyntax) *structure.PromptSyntax {
	_ = "STUB: not implemented"
	return nil
}
