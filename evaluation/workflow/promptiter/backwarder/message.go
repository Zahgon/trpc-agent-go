//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package backwarder

import (
	"context"

	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// MessageBuilder encodes one backward request into one runner message.
type MessageBuilder func(ctx context.Context, request *Request) (*model.Message, error)

const defaultMessageTemplateText = `Compute PromptIter backward attribution for one step.

You will receive one backward request for a single executed step.
Return exactly one JSON object with Gradients and Upstream fields.

Requirements:
- Keep the response as raw JSON only.
- Do not wrap the response in markdown code fences.
- Attribute gradients only to listed gradient surfaces.
- Route upstream gradients only to listed predecessor steps.
- Do not broadcast the same gradient packet to every predecessor.

Request JSON:
{{ toPrettyJSON . }}
`

func defaultMessageBuilder() MessageBuilder { _ = "STUB: not implemented"; return *new(MessageBuilder) }

type promptData struct {
	Node             promptNode
	Input            *atrace.Snapshot         `json:",omitempty"`
	Output           *atrace.Snapshot         `json:",omitempty"`
	Error            string                   `json:",omitempty"`
	GradientSurfaces []promptSurface          `json:",omitempty"`
	OtherSurfaces    []promptContextSurface   `json:",omitempty"`
	Predecessors     []promptPredecessor      `json:",omitempty"`
	Incoming         []promptIncomingGradient `json:",omitempty"`
}

type promptNode struct {
	Kind astructure.NodeKind `json:",omitempty"`
	Name string              `json:",omitempty"`
}

type promptSurface struct {
	SurfaceID string
	Type      astructure.SurfaceType
	Value     astructure.SurfaceValue
}

type promptContextSurface struct {
	Type  astructure.SurfaceType
	Value astructure.SurfaceValue
}

type promptPredecessor struct {
	StepID string
	Output *atrace.Snapshot `json:",omitempty"`
	Error  string           `json:",omitempty"`
}

type promptIncomingGradient struct {
	Severity promptiter.LossSeverity
	Gradient string
}

func newPromptData(request *Request) promptData { _ = "STUB: not implemented"; return *new(promptData) }

// toPrettyJSON renders one value as indented JSON for prompts.
func toPrettyJSON(value any) (string, error) { _ = "STUB: not implemented"; return "", nil }
