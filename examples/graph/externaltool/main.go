//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates model‑orchestrated external tools that run
// outside the graph process. The Large Language Model (LLM) first returns
// a tool call (for example: extract document content), the client executes
// that tool out‑of‑process and feeds the result back, then the model may
// request another tool (for example: summarize the extracted content) and
// continue. We implement this by intercepting tool calls in a callback,
// emitting a graph interrupt with the tool call details, and resuming with
// the user‑supplied tool result.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultModelName = "deepseek-v4-flash"
	userID           = "demo-user"
	sessionID        = "demo-session"

	nodePrepare   = "prepare_input"
	nodeAssistant = "assistant_plan"
	nodeTools     = "external_tools"
	nodeFinish    = "finalize"

	// External + internal tool names for the demo.
	toolNameFetch     = "external_fetch"
	toolNameSummarize = "summarize_text"
	toolNameFormat    = "format_bullets"

	// Interrupt key for passing external tool results back to callback.
	interruptKeyTool = "external_tool_result"

	stateKeyQuestion = "user_question"
)

var modelName = flag.String("model", defaultModelName,
	"Name of the model to use for the assistant")

func main() {
	flag.Parse()

	ctx := context.Background()
	workflow := &externalToolWorkflow{
		modelName: *modelName,
	}
	if err := workflow.setup(); err != nil {
		fmt.Printf("failed to set up workflow: %v\n", err)
		os.Exit(1)
	}

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer workflow.runner.Close()

	if err := workflow.interactive(ctx); err != nil {
		fmt.Printf("workflow ended with error: %v\n", err)
		os.Exit(1)
	}
}

// externalToolWorkflow wires together the runner, graph, and CLI helpers.
type externalToolWorkflow struct {
	modelName string
	runner    runner.Runner
	saver     graph.CheckpointSaver
	manager   *graph.CheckpointManager

	// coordinator intercepts tool calls and raises interrupts so the
	// client can execute tools externally and feed results back.
	coordinator *externalCoordinator

	currentLineage string
	pending        *pendingResume
}

// setup prepares the graph, agent, runner, and checkpoint services.
func (w *externalToolWorkflow) setup() error { _ = "STUB: not implemented"; return nil }

// buildGraph constructs the minimal graph with LLM, tool, and finish nodes.
func (w *externalToolWorkflow) buildGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only external_fetch is external; others run as normal tools.

// Lower temperature to reduce chatty clarifications and make
// tool usage more consistent for this demo.

// Lower temperature to reduce chatty clarifications and keep tools
// deterministic for this demo.

// assistantPrompt defines the system prompt passed to the LLM node.
func assistantPrompt() string { _ = "STUB: not implemented"; return "" }

// prepareInput stores the cleaned user question in the shared state.
func (w *externalToolWorkflow) prepareInput(
	ctx context.Context,
	state graph.State,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// finishConversation ensures the assistant produced a response.
func (w *externalToolWorkflow) finishConversation(
	ctx context.Context,
	state graph.State,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// wrapToolsNode injects the current state before delegating to the tool node.
func (w *externalToolWorkflow) wrapToolsNode(
	base graph.NodeFunc,
) graph.NodeFunc {
	_ = "STUB: not implemented"
	return *new(graph.NodeFunc)
}

// interactive provides the command-line interface.
func (w *externalToolWorkflow) interactive(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Submit a simple placeholder content for external extract.

// printHelp displays available commands.
func (w *externalToolWorkflow) printHelp() { _ = "STUB: not implemented"; return }

// ask launches a new run unless an interrupt is waiting for resume.
func (w *externalToolWorkflow) ask(ctx context.Context, input string) error {
	_ = "STUB: not implemented"
	return nil
}

// resume continues a paused run with the provided tool result.
func (w *externalToolWorkflow) resume(
	ctx context.Context,
	content string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Provide a simple placeholder content to the external tool.

// runAndStream executes the graph and streams events to the terminal.
func (w *externalToolWorkflow) runAndStream(
	ctx context.Context,
	msg model.Message,
	runtimeState map[string]any,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// capturePending stores checkpoint information needed for resume.
func (w *externalToolWorkflow) capturePending(
	ctx context.Context,
	raw any,
) error {
	_ = "STUB: not implemented"
	return nil
}

// printPending informs the user about the paused state.
func (w *externalToolWorkflow) printPending() { _ = "STUB: not implemented"; return }

// printToolCalls displays tool call requests issued by the LLM.
func (w *externalToolWorkflow) printToolCalls(evt *event.Event) { _ = "STUB: not implemented"; return }

// printToolResult shows tool response events after resume.
func (w *externalToolWorkflow) printToolResult(evt *event.Event) { _ = "STUB: not implemented"; return }

// printStreaming streams assistant output as it arrives.
func (w *externalToolWorkflow) printStreaming(
	evt *event.Event,
	started bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// extractInterrupt decodes interrupt metadata from an event, if any.
func extractInterrupt(evt *event.Event) *graph.PregelStepMetadata {
	_ = "STUB: not implemented"
	return nil
}

// newLineage generates a unique lineage id for checkpointing.
func (w *externalToolWorkflow) newLineage() string { _ = "STUB: not implemented"; return "" }

// pendingResume holds data required to resume an interrupted run.
type pendingResume struct {
	checkpointID string
	prompt       string
	raw          any
}

// formatInterruptPrompt converts the interrupt payload into a
// human-friendly text block.
func formatInterruptPrompt(value any) string { _ = "STUB: not implemented"; return "" }

// ----- External tool interception layer -----

// externalCoordinator implements a ToolCallbacks.BeforeTool that pauses the
// graph and resumes with client‑provided tool results.
type externalCoordinator struct {
	mu    sync.Mutex
	state graph.State
}

func (c *externalCoordinator) setState(state graph.State) { _ = "STUB: not implemented"; return }

func (c *externalCoordinator) clearState() { _ = "STUB: not implemented"; return }

func (c *externalCoordinator) callbacks() *tool.Callbacks { _ = "STUB: not implemented"; return nil }

// before intercepts tool execution, raises an interrupt and waits for
// the client to submit the tool result.
func (c *externalCoordinator) before(
	ctx context.Context,
	args *tool.BeforeToolArgs,
) (*tool.BeforeToolResult, error) {
	_ = "STUB: not implemented"
	// Only intercept the external extract tool; others run normally.
	return nil, nil
}

// Convert a simple string into {"content": string} object.

func (c *externalCoordinator) getState() graph.State {
	_ = "STUB: not implemented"
	return *new(graph.State)
}

// parseExternalResult converts the resume payload into a Go value.
// Accepts either a raw JSON string or a map with field "result".
func parseExternalResult(v any) any {
	_ = "STUB: not implemented"
	// Preferred: plain string becomes content field.
	return *new(any)
}

// If client provided a map with content, pass it.

// Fallback: stringify.

// decodeJSONAny decodes a JSON string to an arbitrary Go value.
// (decodeJSONAny removed; not needed in simplified flow)

// declaredTool wraps a declaration as a non‑callable tool so that the
// BeforeTool callback can fully control execution.
type declaredToolWrapper struct{ d *tool.Declaration }

func (w declaredToolWrapper) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func declaredTool(d *tool.Declaration) tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

// Declarations for demo tools.
func fetchDecl() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func summarizeDecl() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func formatDecl() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// ----- Internal logic tools (callable) -----

// summarizerTool implements a simple in-process summarization.
type summarizerTool struct{}

func (summarizerTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (summarizerTool) Call(
	ctx context.Context, jsonArgs []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Create a tiny, readable summary for demo purposes.

// formatterTool formats text into simple bullets.
type formatterTool struct{}

func (formatterTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (formatterTool) Call(
	ctx context.Context, jsonArgs []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// clip returns a string trimmed to n runes.
func clip(s string, n int) string { _ = "STUB: not implemented"; return "" }

// toBullets turns text into a bullet list.
func toBullets(s string) string {
	_ = "STUB: not implemented"
	// Split by sentences; keep it simple for demo.
	return ""
}
