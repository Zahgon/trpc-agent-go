//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package runner captures common runner outputs for PromptIter workflow adapters.
package runner

import (
	"trpc.group/trpc-go/trpc-agent-go/event"
)

// Output stores the structured and textual outputs collected from one runner invocation.
type Output struct {
	// StructuredOutput keeps the latest in-memory structured payload emitted by the runner.
	StructuredOutput any
	// FinalContent keeps the latest textual assistant content emitted by the runner.
	FinalContent string
}

// CaptureOutput consumes one runner event stream and extracts common output artifacts.
func CaptureOutput(events <-chan *event.Event) (*Output, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
