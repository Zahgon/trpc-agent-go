//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package aggregator

import (
	"context"

	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// MessageBuilder encodes one aggregation request into one runner message.
type MessageBuilder func(ctx context.Context, request *Request) (*model.Message, error)

const defaultMessageTemplateText = `Aggregate PromptIter gradients for a single surface.

You will receive one aggregation request with all sample-level gradients that belong to the same surface.
Return exactly one JSON object with a Gradients field.

Requirements:
- Keep the response as raw JSON only.
- Do not wrap the response in markdown code fences.
- Return only merged gradient items. The caller will attach the target surface identity and type.
- Merge duplicated or overlapping gradients when appropriate.
- Drop clearly empty or redundant gradient items.

Request JSON:
{{ toPrettyJSON . }}
`

func defaultMessageBuilder() MessageBuilder { _ = "STUB: not implemented"; return *new(MessageBuilder) }

type promptData struct {
	Surface   promptSurface
	Gradients []promptGradient
}

type promptSurface struct {
	Type astructure.SurfaceType
}

type promptGradient struct {
	Severity promptiter.LossSeverity
	Gradient string
}

func newPromptData(request *Request) promptData { _ = "STUB: not implemented"; return *new(promptData) }

// toPrettyJSON renders one value as indented JSON for prompts.
func toPrettyJSON(value any) (string, error) { _ = "STUB: not implemented"; return "", nil }
