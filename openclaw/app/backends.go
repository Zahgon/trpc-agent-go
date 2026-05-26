//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package app

import (
	"trpc.group/trpc-go/trpc-agent-go/memory"
	memextractor "trpc.group/trpc-go/trpc-agent-go/memory/extractor"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/summary"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

func newSessionService(
	mdl model.Model,
	opts runOptions,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func newInMemorySessionBackend(
	deps registry.SessionDeps,
	_ registry.SessionBackendSpec,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func newRedisSessionBackend(
	deps registry.SessionDeps,
	spec registry.SessionBackendSpec,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func newMemoryService(
	mdl model.Model,
	opts runOptions,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

func resolveMemoryBackendType(raw string) string { _ = "STUB: not implemented"; return "" }

func newDisabledMemoryBackend(
	_ registry.MemoryDeps,
	_ registry.MemoryBackendSpec,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

func newInMemoryMemoryBackend(
	deps registry.MemoryDeps,
	spec registry.MemoryBackendSpec,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

func newRedisMemoryBackend(
	deps registry.MemoryDeps,
	spec registry.MemoryBackendSpec,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

func newSessionSummarizer(
	mdl model.Model,
	opts runOptions,
) (summary.SessionSummarizer, error) {
	_ = "STUB: not implemented"
	return *new(summary.SessionSummarizer), nil
}

// Override the token counter heuristic when configured.
// The framework default (4 runes/token) works well for English but
// underestimates Chinese text (~1–2 runes/token). Setting a lower
// value ensures summary triggers fire at the right time.

// Context-window aware: dynamically resolve the model's context
// window at evaluation time, trigger when delta tokens exceed a
// fraction of it (default 50%). Zero-configuration.

// Manual thresholds (original behavior).

func parseSummaryPolicy(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func newAutoMemoryExtractor(
	mdl model.Model,
	opts runOptions,
) (memextractor.MemoryExtractor, error) {
	_ = "STUB: not implemented"
	return *new(memextractor.MemoryExtractor), nil
}

// When no checkers are configured, ShouldExtract always returns
// true so extraction runs on every turn. We still validate
// MemoryAutoPolicy to catch misconfigurations early.

const summaryToolResultMaxRunes = 2000

// summaryToolResultFormatter truncates large tool results to avoid blowing
// the summarizer model's context window. web_fetch and similar tools can
// return 100 KB+ of HTML/Markdown; without truncation the conversation text
// sent to the summary LLM will exceed small-to-medium model limits.
func summaryToolResultFormatter(msg model.Message) string { _ = "STUB: not implemented"; return "" }
