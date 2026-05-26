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

type resolvedAgentPrompts struct {
	Instruction  string
	SystemPrompt string
}

func resolveAgentPrompts(opts runOptions) (resolvedAgentPrompts, error) {
	_ = "STUB: not implemented"
	return *new(resolvedAgentPrompts), nil
}

func resolveAgentPromptsWithGetwd(
	opts runOptions,
	getwd func() (string, error),
) (resolvedAgentPrompts, error) {
	_ = "STUB: not implemented"
	return *new(resolvedAgentPrompts), nil
}

func resolveAgentPromptsForDir(
	opts runOptions,
	cwd string,
) (resolvedAgentPrompts, error) {
	_ = "STUB: not implemented"
	return *new(resolvedAgentPrompts), nil
}

func joinPromptParts(parts ...string) string { _ = "STUB: not implemented"; return "" }

func buildAgentPrompt(
	inline string,
	files []string,
	dir string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readAgentPromptFile(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func readAgentPromptDir(dir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
