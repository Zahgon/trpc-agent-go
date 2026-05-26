//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package plugin

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	defaultGlobalInstructionPluginName = "global_instruction"
	doubleNewline                      = "\n\n"
)

// GlobalInstruction prepends a system message to every model request.
//
// This is useful for enforcing organization-wide policies or shared behaviors
// without repeating configuration on each Agent.
type GlobalInstruction struct {
	name        string
	instruction string
}

// NewGlobalInstruction creates a GlobalInstruction plugin with a default
// name.
func NewGlobalInstruction(instruction string) *GlobalInstruction {
	_ = "STUB: not implemented"
	return nil
}

// NewNamedGlobalInstruction creates a GlobalInstruction plugin with a custom
// name. Names must be unique per Runner.
func NewNamedGlobalInstruction(
	name string,
	instruction string,
) *GlobalInstruction {
	_ = "STUB: not implemented"
	return nil
}

// Name implements Plugin.
func (p *GlobalInstruction) Name() string {
	_ = "STUB: not implemented"

	// Register implements Plugin.
	return ""
}

func (p *GlobalInstruction) Register(r *Registry) { _ = "STUB: not implemented"; return }

func (p *GlobalInstruction) beforeModel(
	_ context.Context,
	args *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyGlobalInstruction(req *model.Request, instr string) { _ = "STUB: not implemented"; return }
