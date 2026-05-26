//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package a2ui

import (
	"context"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/translator"
)

const defaultA2UISource = "a2ui/v0.8"

// NewFactory wraps a base translator factory and applies A2UI options.
func NewFactory(innerFactory runner.TranslatorFactory, baseOpts []translator.Option, a2uiOpts ...Option) runner.TranslatorFactory {
	_ = "STUB: not implemented"
	return *new(runner.TranslatorFactory)
}

// a2uiTranslator adapts default translator output for A2UI streaming.
type a2uiTranslator struct {
	inner                translator.Translator
	parser               *parser
	receiving            bool
	source               string
	passThroughEventHook PassThroughEventHook
}

func newA2UITranslator(inner translator.Translator, opts ...Option) *a2uiTranslator {
	_ = "STUB: not implemented"
	return nil
}

// Translate runs the inner translator and converts text message chunks to RAW events.
func (t *a2uiTranslator) Translate(
	ctx context.Context,
	event *event.Event,
) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostRunFinalizationEvents finalizes pending A2UI text streams after a run ends.
func (t *a2uiTranslator) PostRunFinalizationEvents(ctx context.Context) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *a2uiTranslator) toRawEvents(lines []string) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

func parseA2UIJSONLine(line string) any { _ = "STUB: not implemented"; return *new(any) }
