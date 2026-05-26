//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package runner

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	agenttool "trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	errAGUIToolNameRequired        = "agui tool name is required"
	errAGUIToolNameRequiredAt      = "agui tool[%d]: %s"
	errConvertAGUIToolParameters   = "convert agui tool[%d] %q parameters: %w"
	errMarshalAGUIToolParameters   = "marshal agui tool parameters"
	errUnmarshalAGUIToolParameters = "unmarshal agui tool parameters"
	jsonSchemaTypeObject           = "object"
)

func appendExternalToolRunOption(
	opts []agent.RunOption,
	input *adapter.RunAgentInput,
) ([]agent.RunOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func externalToolsFromRunAgentInput(
	input *adapter.RunAgentInput,
) ([]agenttool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func aguiToolParametersToSchema(params any) (*agenttool.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type declarationOnlyTool struct {
	declaration *agenttool.Declaration
}

func (t *declarationOnlyTool) Declaration() *agenttool.Declaration {
	_ = "STUB: not implemented"
	return nil
}
