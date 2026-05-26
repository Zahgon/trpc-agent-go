//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

var nonExecutableCodeLanguages = map[string]struct{}{
	"":          {},
	"css":       {},
	"csv":       {},
	"html":      {},
	"json":      {},
	"markdown":  {},
	"md":        {},
	"mermaid":   {},
	"plain":     {},
	"plaintext": {},
	"text":      {},
	"txt":       {},
	"xml":       {},
	"yaml":      {},
	"yml":       {},
}

// CodeExecutionResponseProcessor processes code execution responses from the model.
type CodeExecutionResponseProcessor struct {
}

// NewCodeExecutionResponseProcessor creates a new instance of CodeExecutionResponseProcessor.
// This processor is responsible for handling code execution responses from the model.
func NewCodeExecutionResponseProcessor() *CodeExecutionResponseProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessResponse processes the model response, extracts code blocks, executes them,
// and emits events for the code execution result.
func (p *CodeExecutionResponseProcessor) ProcessResponse(
	ctx context.Context, invocation *agent.Invocation, req *model.Request, rsp *model.Response, ch chan<- *event.Event) {
	_ = "STUB: not implemented"
	return
}

// todo: truncate the content

//  [Step 2] Executes the code and emit 2 Events for code and execution result.

// Add tag for error result

//  [Step 3] Skip processing the original model response to continue code generation loop.

func codeExecutorForInvocation(
	invocation *agent.Invocation,
) codeexecutor.CodeExecutor {
	_ = "STUB: not implemented"
	return *new(codeexecutor.CodeExecutor)
}

func autoExecutableCodeBlocks(
	content string,
	delimiter codeexecutor.CodeBlockDelimiter,
) []codeexecutor.CodeBlock {
	_ = "STUB: not implemented"
	return nil
}

func isAutoExecutableCodeLanguage(language string) bool { _ = "STUB: not implemented"; return false }
