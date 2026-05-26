//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package decode provides typed output decoders for PromptIter internals.
package decode

import (
	irunner "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/internal/runner"
)

// DecodeOutputJSON decodes one runner output into one typed JSON payload.
//
// It first accepts exact structured outputs of type T or *T. When the runner
// returns a generic JSON-compatible value, it falls back to marshaling that
// payload back into JSON and decoding it into T. When structured output is
// absent, it falls back to decoding the final assistant content as JSON.
func DecodeOutputJSON[T any](output *irunner.Output) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeStructuredOutput decodes exact typed payloads and generic JSON objects.
func decodeStructuredOutputExact[T any](payload any) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeFinalContentJSON decodes the final assistant content as JSON.
func decodeFinalContentJSON[T any](content string) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
