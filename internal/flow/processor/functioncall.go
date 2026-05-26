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
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// ErrorToolNotFound is the error message for tool not found.
	ErrorToolNotFound = "Error: tool not found"
	// ErrorCallableToolExecution is the error message for callable tool execution failed.
	ErrorCallableToolExecution = "Error: callable tool execution failed"
	// ErrorStreamableToolExecution is the error message for streamable tool execution failed.
	ErrorStreamableToolExecution = "Error: streamable tool execution failed"
	// ErrorMarshalResult is the error message for failed to marshal result.
	ErrorMarshalResult = "Error: failed to marshal result"
)

// funcRespCompletionTimeout is the default wait duration for ensuring a
// tool.response event has been processed by the session persistence layer.
const funcRespCompletionTimeout = 5 * time.Second

// summarizationSkipper is implemented by tools that can indicate whether
// the flow should skip a post-tool summarization step. This allows tools
// like AgentTool to mark their tool.response as final for the turn.
type summarizationSkipper interface {
	SkipSummarization() bool
}

// streamInnerPreference is implemented by tools that want to control whether
// the flow should treat them as streamable (forwarding inner deltas) or fall
// back to the callable path. When this returns false, the flow will not use
// the StreamableTool path even if the tool implements it.
type streamInnerPreference interface {
	StreamInner() bool
}

type innerTextModePreference interface {
	InnerTextMode() tool.InnerTextMode
}

type toolEventStateDelta struct {
	tool            tool.Tool
	invocation      *agent.Invocation
	sessionBaseline session.StateMap
	args            []byte
	choice          model.Choice
}

type resolvedToolContextKey struct{}
type stateDeltaSessionBaselineContextKey struct{}
type executingToolArgsContextKey struct{}

// toolResult holds the result of a single tool execution.
type toolResult struct {
	index      int
	event      *event.Event
	err        error
	stateDelta *toolEventStateDelta
	toolArgs   []byte
}

// Default message used when transferring to a sub-agent without an explicit message.
// Users can override or disable it via SetDefaultTransferMessage.
var defaultTransferMessage = "Task delegated from coordinator"

// SetDefaultTransferMessage configures the message to inject when a sub-agent is
// called without an explicit message (model directly calls the sub-agent name).
func SetDefaultTransferMessage(message string) { _ = "STUB: not implemented"; return }

// subAgentCall defines the input format for direct sub-agent tool calls.
// This handles cases where models call sub-agent names directly instead of using transfer_to_agent.
type subAgentCall struct {
	Message string `json:"message,omitempty"`
}

// FunctionCallResponseProcessor handles agent transfer operations after LLM responses.
type FunctionCallResponseProcessor struct {
	enableParallelTools bool
	toolCallbacks       *tool.Callbacks
	toolRetryPolicy     *tool.RetryPolicy
}

// FunctionCallResponseProcessorOption configures a function-call response processor.
type FunctionCallResponseProcessorOption func(*FunctionCallResponseProcessor)

// WithToolCallRetryPolicy sets the retry policy used for single callable tool invocations.
func WithToolCallRetryPolicy(policy *tool.RetryPolicy) FunctionCallResponseProcessorOption {
	_ = "STUB: not implemented"
	return *new(FunctionCallResponseProcessorOption)
}

// NewFunctionCallResponseProcessor creates a new transfer response processor.
func NewFunctionCallResponseProcessor(
	enableParallelTools bool,
	toolCallbacks *tool.Callbacks,
	opts ...FunctionCallResponseProcessorOption,
) *FunctionCallResponseProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessResponse implements the flow.ResponseProcessor interface.
// It checks for transfer requests and handles agent handoffs by actually calling
// the target agent's Run method.
func (p *FunctionCallResponseProcessor) ProcessResponse(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	rsp *model.Response,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Enforce optional per-invocation tool iteration limit. A "tool iteration"
// is defined as an assistant response that contains tool calls and reaches
// this processor. When the limit is not configured (<= 0), this check is a
// no-op and preserves existing behavior.

// Mark the invocation as ended so the flow will not issue another LLM call.

// Emit an error response event describing the limit breach instead of
// executing any tools. This makes the termination visible to callers
// while avoiding additional model or tool invocations.

// Option one: set invocation.EndInvocation is true, and stop next step.
// Option two: emit error event, maybe the LLM can correct this error and also need to wait for notice completion.
// maybe the Option two is better.
// Allow users to intervene in error handling through callbacks.

// If the tool indicates skipping outer summarization, mark the invocation to end
// after this tool response so the flow does not perform an extra LLM call.

func (p *FunctionCallResponseProcessor) toolExecutionDecision(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	rsp *model.Response,
) (deferred bool, executable bool, unknown bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (p *FunctionCallResponseProcessor) handleFunctionCallsAndSendEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	llmResponse *model.Response,
	tools map[string]tool.Tool,
	eventChan chan<- *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func funcRespWaitTimeout(ctx context.Context) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// handleFunctionCalls executes tool calls and returns a merged response event.
func (p *FunctionCallResponseProcessor) handleFunctionCalls(
	ctx context.Context,
	invocation *agent.Invocation,
	llmResponse *model.Response,
	tools map[string]tool.Tool,
	eventChan chan<- *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If parallel tools are enabled AND multiple tool calls, execute concurrently

// executeSingleToolCallSequential runs one tool call and returns its event.
func (p *FunctionCallResponseProcessor) executeSingleToolCallSequential(
	ctx context.Context,
	invocation *agent.Invocation,
	llmResponse *model.Response,
	tools map[string]tool.Tool,
	eventChan chan<- *event.Event,
	index int,
	toolCall model.ToolCall,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *FunctionCallResponseProcessor) executeSingleToolCallSequentialResult(
	ctx context.Context,
	invocation *agent.Invocation,
	llmResponse *model.Response,
	tools map[string]tool.Tool,
	eventChan chan<- *event.Event,
	index int,
	toolCall model.ToolCall,
) (toolResult, error) {
	_ = "STUB: not implemented"
	return *new(toolResult), nil
}

// Create error choice for ignorable errors

// Return critical errors (e.g., stop errors) immediately

// executeToolCallsInParallel runs multiple tool calls concurrently and merges
// their results into a single event.
func (p *FunctionCallResponseProcessor) executeToolCallsInParallel(
	ctx context.Context,
	invocation *agent.Invocation,
	llmResponse *model.Response,
	toolCalls []model.ToolCall,
	tools map[string]tool.Tool,
	eventChan chan<- *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runParallelToolCall executes one tool call and reports the result.
func (p *FunctionCallResponseProcessor) runParallelToolCall(
	ctx context.Context,
	wg *sync.WaitGroup,
	invocation *agent.Invocation,
	llmResponse *model.Response,
	tools map[string]tool.Tool,
	eventChan chan<- *event.Event,
	resultChan chan<- toolResult,
	index int,
	tc model.ToolCall,
) {
	_ = "STUB: not implemented"
	return
}

// Recover from panics to avoid breaking sibling goroutines.

// Trace the tool execution for observability.

// Execute the tool (streamable or callable) with callbacks.

// Handle errors based on whether they are ignorable or critical.

// Only propagate the error if it's not ignorable (e.g., stop errors)

// No error and at least one choice means we have tool result messages.

// Include declaration for telemetry even when tool is missing.

// Send result back to aggregator.

func (p *FunctionCallResponseProcessor) buildToolCallResponseEvent(
	invocation *agent.Invocation,
	llmResponse *model.Response,
	choices []model.Choice,
	tools map[string]tool.Tool,
	toolCall model.ToolCall,
	index int,
	toolArgs []byte,
	skipSummarization bool,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func annotateToolChoicesWithName(choices []model.Choice, toolName string) {
	_ = "STUB: not implemented"
	return
}

func annotateToolCallArgs(
	ev *event.Event,
	toolCall model.ToolCall,
	toolArgs []byte,
) {
	_ = "STUB: not implemented"
	return
}

func setToolCallArgs(
	ev *event.Event,
	toolCallID string,
	toolArgs []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *FunctionCallResponseProcessor) decorateToolCallResponseEvent(
	ev *event.Event,
	tools map[string]tool.Tool,
	toolCall model.ToolCall,
	skipSummarization bool,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// attachStateDelta copies tool-provided state delta to the event.
func (p *FunctionCallResponseProcessor) attachStateDelta(
	inv *agent.Invocation,
	tl tool.Tool,
	args []byte,
	choice *model.Choice,
	ev *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// annotateSkipSummarization marks an event to skip outer summarization.
func (p *FunctionCallResponseProcessor) annotateSkipSummarization(
	ev *event.Event,
	tl tool.Tool,
	dynamic bool,
) {
	_ = "STUB: not implemented"
	return
}

func (p *FunctionCallResponseProcessor) buildToolEventStateDelta(
	ctx context.Context,
	invocation *agent.Invocation,
	args []byte,
	choices []model.Choice,
) *toolEventStateDelta {
	_ = "STUB: not implemented"
	return nil
}

func newStateDeltaSnapshot(
	ctx context.Context,
	invocation *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func newParallelInvocationView(
	invocation *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func invocationView(invocation *agent.Invocation) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func cloneStateDeltaSession(invocation *agent.Invocation) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func withStateDeltaSessionBaseline(
	ctx context.Context,
	invocation *agent.Invocation,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func stateDeltaBaselineSession(
	ctx context.Context,
	invocation *agent.Invocation,
) *session.Session {
	_ = "STUB: not implemented"
	return nil
}

func stateDeltaSessionBaseline(ctx context.Context) session.StateMap {
	_ = "STUB: not implemented"
	return *new(session.StateMap)
}

func preserveStateDeltaInvocationDefaults(
	view *agent.Invocation,
	base *agent.Invocation,
) {
	_ = "STUB: not implemented"
	return
}

func newStateDeltaInvocationView(
	invocation *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func withResolvedToolContext(
	ctx context.Context,
	tl tool.Tool,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func resolvedToolFromContext(ctx context.Context) (tool.Tool, bool) {
	_ = "STUB: not implemented"
	return *new(tool.Tool), false
}

func (p *FunctionCallResponseProcessor) attachStateDeltaToToolResults(
	invocation *agent.Invocation,
	results []toolResult,
) []*event.Event {
	_ = "STUB: not implemented"
	return nil
}

func applyPriorStateDeltas(
	sess *session.Session,
	baseline session.StateMap,
	events []*event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func changedSessionKeys(
	baseline session.StateMap,
	current session.StateMap,
) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func applyReplayableStateDelta(
	sess *session.Session,
	changed map[string]bool,
	e *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func stateDeltaInvocationView(
	invocation *agent.Invocation,
	stateDelta *toolEventStateDelta,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

// lookupDeclaration returns a declaration or a safe placeholder.
func (p *FunctionCallResponseProcessor) lookupDeclaration(
	tools map[string]tool.Tool, name string,
) *tool.Declaration {
	_ = "STUB: not implemented"
	return nil
}

// sendToolResult sends without blocking when the context is cancelled.
func (p *FunctionCallResponseProcessor) sendToolResult(
	ctx context.Context, ch chan<- toolResult, res toolResult,
) {
	_ = "STUB: not implemented"
	return
}

// buildMergedParallelEvent merges child tool events or builds minimal choices.
func (p *FunctionCallResponseProcessor) buildMergedParallelEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	llmResponse *model.Response,
	tools map[string]tool.Tool,
	toolCalls []model.ToolCall,
	toolResults []toolResult,
	toolCallEvents []*event.Event,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// executeToolCall executes a single tool call and returns the choice.
// Parameters:
//   - ctx: context for cancellation and tracing
//   - invocation: agent invocation context containing agent name, model info, etc.
//   - toolCall: the tool call to execute, including function name and arguments
//   - tools: map of available tools by name
//   - index: index of this tool call in the batch (for error reporting)
//   - eventChan: channel for emitting events during execution
//
// Returns:
//   - context.Context: updated context from callbacks (if any)
//   - []model.Choice: tool response choices (nil if no response is emitted)
//   - []byte: the modified arguments after before-tool callbacks (for telemetry)
//   - bool: shouldIgnoreError - true if the error is ignorable (e.g., tool not found, marshal error), false for critical errors (e.g., stop errors)
//   - bool: skipSummarization - true if callbacks requested ending the turn
//     after the tool response
//   - error: any error that occurred during execution (no longer swallowed)
func (p *FunctionCallResponseProcessor) executeToolCall(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	tools map[string]tool.Tool,
	index int,
	eventChan chan<- *event.Event,
) (context.Context, []model.Choice, []byte, bool, bool, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil, false, false, nil
}

// Execute the tool with callbacks.

// Only return error when it's a stop error

//  allow to return nil not provide function response.

// Marshal failures (for example, NaN in floats) do not
// affect the overall flow. Downgrade to warning to avoid
// noisy alerts while still surfacing the issue.

// resolveToolCallTarget resolves the callable tool, applies compatibility remapping,
// and evaluates the optional tool execution filter.
func (p *FunctionCallResponseProcessor) resolveToolCallTarget(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	tools map[string]tool.Tool,
) (model.ToolCall, tool.Tool, bool, error) {
	_ = "STUB: not implemented"
	return *new(model.ToolCall), *new(tool.Tool), false, nil
}

// Compatibility: map sub-agent name calls to transfer_to_agent if present.

// applyToolResultMessagesCallback invokes the optional ToolResultMessages callback and
// converts its return value into choices. It returns:
//   - customChoices: the choices derived from callback output
//   - overridden: whether the default tool message should be replaced
//   - err: non-nil when the callback itself fails
func (p *FunctionCallResponseProcessor) applyToolResultMessagesCallback(
	ctx context.Context,
	toolCall model.ToolCall,
	tl tool.Tool,
	result any,
	modifiedArgs []byte,
	index int,
	defaultMsg model.Message,
) ([]model.Choice, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// No override.

// When a callback is provided and returns non-empty messages,
// the framework defers entirely to the callback for correctness.

func ensureToolResultMessageName(
	msg model.Message,
	toolCall model.ToolCall,
) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

// createErrorChoice creates an error choice for tool execution failures.
func (p *FunctionCallResponseProcessor) createErrorChoice(index int, toolID string,
	errorMsg string) *model.Choice {
	_ = "STUB: not implemented"
	return nil
}

// collectParallelToolResults drains resultChan and preserves order by index.
// It returns only non-nil events.
func (p *FunctionCallResponseProcessor) collectParallelToolResults(
	ctx context.Context,
	resultChan <-chan toolResult,
	toolCallsCount int,
) ([]toolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Channel closed, all results received.

// Context cancelled, return what we have.

func (p *FunctionCallResponseProcessor) runBeforeToolPluginCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	toolDeclaration *tool.Declaration,
) (context.Context, model.ToolCall, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(model.ToolCall), *new(any), nil
}

func (p *FunctionCallResponseProcessor) runBeforeToolCallbacks(
	ctx context.Context,
	toolCall model.ToolCall,
	toolDeclaration *tool.Declaration,
) (context.Context, model.ToolCall, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(model.ToolCall), *new(any), nil
}

func (p *FunctionCallResponseProcessor) runAfterToolPluginCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	toolDeclaration *tool.Declaration,
	toolResult any,
	toolErr error,
) (context.Context, any, bool, bool, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any), false, false, nil
}

func (p *FunctionCallResponseProcessor) runAfterToolCallbacks(
	ctx context.Context,
	toolCall model.ToolCall,
	toolDeclaration *tool.Declaration,
	toolResult any,
	toolErr error,
) (context.Context, any, bool, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any), false, nil
}

// extractMetaFromResult extracts metadata from tool result.
// For MCP mcpToolResult, returns the Meta field.
func extractMetaFromResult(result any) map[string]any { _ = "STUB: not implemented"; return nil }

// Check for our wrapped mcpToolResult type (from tool/mcp package)

// executeToolWithCallbacks executes a tool with before/after callbacks.
// Returns (context, result, modifiedArguments, suppressDefaultToolMessage,
// skipSummarization, error).
func (p *FunctionCallResponseProcessor) executeToolWithCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	tl tool.Tool,
	eventChan chan<- *event.Event,
) (context.Context, any, []byte, bool, bool, error) {
	_ = "STUB: not implemented"
	// Inject tool call ID into context for callbacks to use.
	return *new(context.Context), *new(any), nil, false, false, nil
}

// Repair tool call arguments in place when needed.

// Execute the actual tool.

// When the after-tool callback replaced the result with a CustomResult,
// the original tool execution error should be cleared so that the
// caller uses the replacement result as the tool response message
// instead of discarding it due to a non-nil error.

func rememberExecutingToolArgs(ctx context.Context, args []byte) { _ = "STUB: not implemented"; return }

// afterCallbackReplacedResult returns true when the after-tool callback has
// replaced the original (failed) tool result with a non-nil custom result.
// In that case the original toolErr should be cleared so that the framework
// sends the replacement result as the tool response message to the model.
// StopError is excluded because it carries a control-flow signal that must
// not be silently swallowed by a callback result replacement.
func afterCallbackReplacedResult(toolErr error, toolResult any) bool {
	_ = "STUB: not implemented"
	return false
}

// isStreamable returns true if the tool supports streaming and its stream
// preference is enabled.
func isStreamable(t tool.Tool) bool {
	_ = "STUB: not implemented"
	// Check if the tool has a stream preference and if it is enabled.
	return false
}

// executeTool executes the tool based on its capabilities.
func (f *FunctionCallResponseProcessor) executeTool(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	tl tool.Tool,
	eventChan chan<- *event.Event,
) (context.Context, any, bool, error) {
	_ = "STUB: not implemented"
	// originalTool refers to the actual underlying tool used to determine
	// whether streaming is supported. If tl is a NamedTool, use its
	// inner original tool instead of the wrapper itself.
	return *new(context.Context), *new(any), false, nil
}

// Prefer streaming execution if the tool supports it.

// Safe to cast since isStreamable checks for StreamableTool.

// Fallback to callable tool execution if supported.

// executeCallableTool executes a callable tool.
func (p *FunctionCallResponseProcessor) executeCallableTool(
	ctx context.Context,
	toolCall model.ToolCall,
	tl tool.CallableTool,
) (context.Context, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any), nil
}

func extractResultError(result any) bool { _ = "STUB: not implemented"; return false }

func buildDefaultToolMessage(
	toolCallID string,
	result any,
) (model.Message, error) {
	_ = "STUB: not implemented"
	// Preserve legacy tool message serialization for default fallback content.
	// Use marshalJSONNoHTMLEscape so that <, >, & in tool output (e.g. Go source
	// code containing "<-done") are preserved verbatim instead of being escaped
	// to \u003c, \u003e, \u0026 which confuses LLMs reading the content.
	return *new(model.Message), nil
}

// marshalJSONNoHTMLEscape serializes v to JSON without escaping <, >, & characters.
// Standard json.Marshal escapes these for HTML safety, but tool results are never
// embedded in HTML and the escaped sequences (\u003c, \u003e, \u0026) confuse LLMs
// that read the output as source code (e.g. Go channel operations "<-done").
func marshalJSONNoHTMLEscape(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// json.Encoder.Encode appends a trailing newline; trim it for Marshal parity.

type structuredStreamErrorOptIn interface {
	TRPCAgentGoStructuredStreamErrorsOptIn() bool
}

func streamableToolCallContext(
	ctx context.Context,
	tl tool.StreamableTool,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func shouldRequestStructuredStreamErrors(tl tool.StreamableTool) bool {
	_ = "STUB: not implemented"
	return false
}

func innerTextModeForTool(tl tool.StreamableTool) tool.InnerTextMode {
	_ = "STUB: not implemented"
	return *new(tool.InnerTextMode)
}

// executeStreamableTool executes a streamable tool.
func (f *FunctionCallResponseProcessor) executeStreamableTool(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	tl tool.StreamableTool,
	eventChan chan<- *event.Event,
) (context.Context, any, bool, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any), false, nil
}

// Process stream chunks, handling:
// Case 1: Raw sub-agent event passthrough.
// Case 2: Plain text-like chunk. Emit partial tool.response event.

// If we forwarded inner events, still return the merged content as the tool
// result so it can be recorded in the tool response message for the next LLM
// turn (to satisfy providers that require tool messages). The UI example
// suppresses printing these aggregated strings to avoid duplication; they are
// primarily for model consumption.

type streamFinalResult struct {
	seen  bool
	value any
}

type streamInnerEventState struct {
	pendingGraphToolErrorEvent *event.Event
}

type normalizedFinalResultChunk struct {
	result     any
	stateDelta map[string][]byte
}

type syntheticStateOnlyToolChoiceKey struct{}

func markSyntheticStateOnlyToolChoice(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func hasSyntheticStateOnlyToolChoice(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// consumeStream reads all chunks from the reader and processes them.
func (f *FunctionCallResponseProcessor) consumeStream(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	reader *tool.StreamReader,
	eventChan chan<- *event.Event,
	innerTextMode tool.InnerTextMode,
	structuredErrors bool,
) ([]any, streamFinalResult, error) {
	_ = "STUB: not implemented"
	return nil, *new(streamFinalResult), nil
}

// appendInnerEventContent extracts textual content from an inner event and appends it.
func (f *FunctionCallResponseProcessor) appendInnerEventContent(
	ev *event.Event,
	contents *[]any,
) {
	_ = "STUB: not implemented"
	return
}

// buildPartialToolResponseEvent constructs a partial tool.response event.
func (f *FunctionCallResponseProcessor) buildPartialToolResponseEvent(
	inv *agent.Invocation,
	toolCall model.ToolCall,
	text string,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (f *FunctionCallResponseProcessor) buildStateDeltaToolResponseEvent(
	inv *agent.Invocation,
	toolCall model.ToolCall,
	stateDelta map[string][]byte,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func cloneEventStateDelta(stateDelta map[string][]byte) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func streamToolEventError(ev *event.Event) error { _ = "STUB: not implemented"; return nil }

func isGraphToolExecutionErrorEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

func isRetryingGraphNodeErrorEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

// marshalChunkToText converts a chunk content into a string representation.
func marshalChunkToText(content any) string { _ = "STUB: not implemented"; return "" }

// compactToolResults keeps received tool results while preserving order.
// Results without an event may still carry executed args for merged metadata.
func (p *FunctionCallResponseProcessor) compactToolResults(
	results []toolResult,
) []toolResult {
	_ = "STUB: not implemented"
	return nil
}

func newToolCallResponseEvent(
	invocation *agent.Invocation,
	functionCallResponse *model.Response,
	functionResponses []model.Choice) *event.Event {
	_ = "STUB: not implemented"
	// Create function response event.
	return nil
}

func newMinimalToolCallResponseEvent(
	invocation *agent.Invocation,
	functionCallResponse *model.Response,
	toolCall model.ToolCall,
	index int,
	toolArgs []byte,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func newMinimalToolChoice(
	toolCall model.ToolCall,
	index int,
) model.Choice {
	_ = "STUB: not implemented"
	return *new(model.Choice)
}

func mergeParallelToolCallResponseEvents(es []*event.Event) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// collectMergedChoices collects the choices from all events.
func collectMergedChoices(es []*event.Event) []model.Choice { _ = "STUB: not implemented"; return nil }

// Add nil checks to prevent panic

// collectStateDelta collects the state delta from all events.
func collectStateDelta(es []*event.Event) map[string][]byte { _ = "STUB: not implemented"; return nil }

func collectToolCallArgs(es []*event.Event) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// findBaseEvent finds a valid base event for metadata.
func findBaseEvent(es []*event.Event) *event.Event { _ = "STUB: not implemented"; return nil }

// buildMergedToolResponse builds the merged tool response.
func buildMergedToolResponse(baseEvent *event.Event, mergedChoices []model.Choice) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// buildMergedEvent builds the merged event.
func buildMergedEvent(baseEvent *event.Event, resp *model.Response) *event.Event {
	_ = "STUB: not implemented"
	// If we have a base event, carry over invocation, author and branch.
	return nil
}

// Fallback: construct without base metadata.

// shouldSkipSummarization checks if any event prefers skipping summarization.
func shouldSkipSummarization(es []*event.Event) bool { _ = "STUB: not implemented"; return false }

func markSkipSummarization(ev *event.Event) { _ = "STUB: not implemented"; return }

func toolPrefersSkipSummarization(tl tool.Tool) bool { _ = "STUB: not implemented"; return false }

// findCompatibleTool attempts to map a requested (missing) tool name to a compatible tool.
// For models that directly call sub-agent names, map to transfer_to_agent when available.
func findCompatibleTool(requested string, tools map[string]tool.Tool, invocation *agent.Invocation) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

// convertToolArguments converts original args to the mapped tool args when needed.
// When mapping sub-agent name -> transfer_to_agent, wrap message and set agent_name.
func convertToolArguments(originalName string, originalArgs []byte, targetName string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// processStreamChunk handles a single streamed chunk and updates contents and events.
func (f *FunctionCallResponseProcessor) processStreamChunk(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	chunk tool.StreamChunk,
	eventChan chan<- *event.Event,
	contents *[]any,
	finalResult *streamFinalResult,
	innerEventState *streamInnerEventState,
	innerTextMode tool.InnerTextMode,
	structuredErrors bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func normalizeFinalResultChunk(content any) (*normalizedFinalResultChunk, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f *FunctionCallResponseProcessor) handleFinalResultChunk(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	eventChan chan<- *event.Event,
	finalResult *streamFinalResult,
	innerEventState *streamInnerEventState,
	finalChunk *normalizedFinalResultChunk,
	structuredErrors bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FunctionCallResponseProcessor) handleStreamInnerEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	contents *[]any,
	innerEventState *streamInnerEventState,
	ev *event.Event,
	innerTextMode tool.InnerTextMode,
	structuredErrors bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func filterForwardedInnerTextEvent(
	ev *event.Event,
	mode tool.InnerTextMode,
) (*event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func shouldEmitFilteredInnerEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

func clearForwardedMessageText(msg *model.Message) bool { _ = "STUB: not implemented"; return false }

func removeForwardedTextContentParts(
	parts []model.ContentPart,
) ([]model.ContentPart, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func responseHasForwardablePayload(rsp *model.Response) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *FunctionCallResponseProcessor) handlePlainStreamChunk(
	ctx context.Context,
	invocation *agent.Invocation,
	toolCall model.ToolCall,
	eventChan chan<- *event.Event,
	contents *[]any,
	innerEventState *streamInnerEventState,
	content any,
	structuredErrors bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func flushPendingGraphToolError(
	state *streamInnerEventState,
	nextEvent *event.Event,
	structuredErrors bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
