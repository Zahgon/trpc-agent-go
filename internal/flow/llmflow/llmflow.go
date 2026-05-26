//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package llmflow provides an LLM-based flow implementation.
package llmflow

import (
	"context"
	"time"

	oteltrace "go.opentelemetry.io/otel/trace"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/internal/flow"
	"trpc.group/trpc-go/trpc-agent-go/internal/flow/processor"
	itelemetry "trpc.group/trpc-go/trpc-agent-go/internal/telemetry"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// Timeout for event completion signaling.
	eventCompletionTimeout    = 5 * time.Second
	generatedResponseIDPrefix = "llmflow-response-"
	queuedUserAuthor          = "user"

	errMsgNoModelResponse = "no response received from model"

	flowRunPanicLogFmt = "Flow execution panic (invocation: %s, " +
		"agent: %s): %v\n%s"

	flowRunPanicErrFmt = "flow panic: %v"

	// stateKeyToolsSnapshot is the invocation state key used to cache the
	// final tool list for a single Invocation. This ensures that the tool
	// set (including ToolSet-based tools and filters) stays stable for the
	// entire lifetime of an Invocation, even when underlying ToolSets are
	// dynamic.
	stateKeyToolsSnapshot = "llmflow:tools_snapshot"
	// stateKeyHasFilteredUserTools caches whether the final filtered tool
	// snapshot for this invocation still contains any user tool.
	stateKeyHasFilteredUserTools = "llmflow:has_filtered_user_tools"

	defaultContextCompactionThresholdRatio = 0.7
	contextCompactionFallbackWindow        = 8192
	contextCompactionMinTokens             = 2000
)

// InvocationHasFilteredUserTools reports whether the cached filtered tool
// snapshot for this invocation still contains any user tool.
func InvocationHasFilteredUserTools(invocation *agent.Invocation) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// Options contains configuration options for creating a Flow.
type Options struct {
	ChannelBufferSize               int // Buffer size for event channels (default: 256).
	ModelCallbacks                  *model.Callbacks
	BaseModelResolver               BaseModelResolver
	ModelSelector                   agent.ModelSelector
	SyncSummaryIntraRun             bool
	EnableContextCompaction         bool
	ContextCompactionThresholdRatio float64
}

// ModelBaseResolution describes the base model for one LLM call.
type ModelBaseResolution struct {
	Model              model.Model
	AllowAgentSelector bool
}

// BaseModelResolver resolves the base model before one LLM call.
type BaseModelResolver func(inv *agent.Invocation) ModelBaseResolution

// Flow provides the basic flow implementation.
type Flow struct {
	requestProcessors               []flow.RequestProcessor
	responseProcessors              []flow.ResponseProcessor
	channelBufferSize               int
	modelCallbacks                  *model.Callbacks
	baseModelResolver               BaseModelResolver
	modelSelector                   agent.ModelSelector
	syncSummaryIntraRun             bool
	enableContextCompaction         bool
	contextCompactionThresholdRatio float64
}

type contextCompactionTailProcessor interface {
	SupportsContextCompactionRebuild(
		invocation *agent.Invocation,
	) bool
	RebuildRequestForContextCompaction(
		ctx context.Context,
		invocation *agent.Invocation,
		req *model.Request,
	)
}

type contextCompactionRebuildPlan struct {
	beforeContent    *model.Request
	contentProcessor *processor.ContentRequestProcessor
	tailProcessors   []contextCompactionTailProcessor
}

type summarySnapshot struct {
	exists    bool
	summary   string
	updatedAt time.Time
}

// New creates a new basic flow instance with the provided processors.
// Processors are immutable after creation.
func New(
	requestProcessors []flow.RequestProcessor,
	responseProcessors []flow.ResponseProcessor,
	opts Options,
) *Flow {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the flow in a loop until completion.
func (f *Flow) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Configurable buffered channel for events.

// Mark the invocation so the runner skips redundant async
// summary enqueue when sync intra-run summary handles it.

// Optionally resume from pending tool calls before starting a new
// LLM cycle. This covers scenarios where the previous run stopped
// after an assistant tool_call response but before tools executed.

// emit start event and wait for completion notice.

// Run sync intra-run summary only between iterations.

// Run one step (one LLM call cycle).

// Treat context cancellation as graceful termination (common in streaming
// pipelines where the client closes the stream after final event).

// Send error event through channel instead of just logging.

// Exit conditions.
// If no events were produced in this step, treat as terminal to avoid busy loop.
// Also break when EndInvocation is set or a final response is observed.

func recoverFlowRunPanic(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func flowInvocationID(invocation *agent.Invocation) string { _ = "STUB: not implemented"; return "" }

func flowAgentName(invocation *agent.Invocation) string { _ = "STUB: not implemented"; return "" }

func traceSnapshotFromMessages(messages []model.Message) *atrace.Snapshot {
	_ = "STUB: not implemented"
	return nil
}

func traceSnapshotFromEvent(evt *event.Event) *atrace.Snapshot {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flow) maybeConsumeQueuedUserMessages(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

func flowEventWaitTimeout(ctx context.Context) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// maybeResumePendingToolCalls inspects the latest session events and, when
// RunOptions.Resume is enabled, executes any pending tool calls before the
// next LLM request. A pending tool call is defined as the latest persisted
// event being an assistant response that contains tool calls but no tool
// results after it.
func (f *Flow) maybeResumePendingToolCalls(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (f *Flow) maybeSyncSummaryIntraRun(
	ctx context.Context,
	invocation *agent.Invocation,
) {
	_ = "STUB: not implemented"
	return
}

func (f *Flow) emitStartEventAndWait(ctx context.Context, invocation *agent.Invocation,
	eventChan chan<- *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for completion notice.
// Ensure that the events of the previous agent or the previous step have been synchronized to the session.

func (f *Flow) selectModelForStep(
	ctx context.Context,
	invocation *agent.Invocation,
) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

func runModelSelector(
	ctx context.Context,
	selector agent.ModelSelector,
	invocation *agent.Invocation,
) (selected model.Model, err error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// runOneStep executes one step of the flow (one LLM call cycle).
// Returns the last event generated, or nil if no events.
func (f *Flow) runOneStep(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil,

		// Initialize empty LLM request.
		nil
}

// Initialize tools map

// 1. Preprocess (prepare request).

// 2. Call LLM (get response sequence).

// 3. Process streaming responses.

// processStreamingResponses handles the streaming response processing logic.
func (f *Flow) processStreamingResponses(
	ctx context.Context,
	invocation *agent.Invocation,
	observabilityInvocation *agent.Invocation,
	llmRequest *model.Request,
	responseSeq model.Seq[*model.Response],
	eventChan chan<- *event.Event,
	span oteltrace.Span,
	startedSpan bool,
) (lastEvent *event.Event, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle after model callbacks.

// Repair tool call arguments in place when needed.

// 4. Create and send LLM response using the clean constructor.

// 5. Check context cancellation.

// 6. Postprocess response.

// handleAfterModelCallbacks processes after model callbacks.
func (f *Flow) handleAfterModelCallbacks(
	ctx context.Context,
	eventInvocation *agent.Invocation,
	invocation *agent.Invocation,
	llmRequest *model.Request,
	response *model.Response,
	eventChan chan<- *event.Event,
) (context.Context, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// createLLMResponseEvent creates a new LLM response event.
func (f *Flow) createLLMResponseEvent(
	eventInvocation *agent.Invocation,
	optionsInvocation *agent.Invocation,
	response *model.Response,
	llmRequest *model.Request,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func invocationFromContextOrDefault(
	ctx context.Context,
	invocation *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func invocationViewForModel(
	invocation *agent.Invocation,
	callModel model.Model,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func metricsInvocationForCurrent(
	current *agent.Invocation,
	base *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func observabilityInvocationForCurrent(
	current *agent.Invocation,
	base *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func trackModelResponseTelemetry(
	response *model.Response,
	tracker *itelemetry.ChatMetricsTracker,
) {
	_ = "STUB: not implemented"
	return
}

func responseUsageTimingInfo(invocation *agent.Invocation) *model.TimingInfo {
	_ = "STUB: not implemented"
	return nil
}

func applyPartialEventMetadataOverrides(
	ev *event.Event,
	response *model.Response,
	invocation *agent.Invocation,
) {
	_ = "STUB: not implemented"
	return
}

func collectLongRunningToolIDs(ToolCalls []model.ToolCall, tools map[string]tool.Tool) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flow) runAfterModelCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	response *model.Response,
) (context.Context, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func runAfterModelCallbackSet(
	ctx context.Context,
	callbacks *model.Callbacks,
	req *model.Request,
	response *model.Response,
) (context.Context, *model.Response, bool, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, false, nil
}

// preprocess handles pre-LLM call preparation using request processors.
func (f *Flow) preprocess(
	ctx context.Context,
	invocation *agent.Invocation,
	llmRequest *model.Request,
	eventChan chan<- *event.Event,
) *contextCompactionRebuildPlan {
	_ = "STUB: not implemented"
	return nil
}

// Run request processors - they send events directly to the channel.

// Add tools to the request with optional filtering.

// Sanitize invalid tool calls in history to avoid poisoning future requests.

func normalizeContextCompactionThresholdRatio(ratio float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (f *Flow) maybeCompactContextBeforeLLM(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	rebuildPlan *contextCompactionRebuildPlan,
) *model.Request {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flow) rebuildRequestForContextCompaction(
	ctx context.Context,
	invocation *agent.Invocation,
	rebuildPlan *contextCompactionRebuildPlan,
) *model.Request {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flow) supportsSyncSummaryRetry() bool { _ = "STUB: not implemented"; return false }

func cloneRequestForContextCompaction(req *model.Request) *model.Request {
	_ = "STUB: not implemented"
	return nil
}

func cloneMessagesForContextCompaction(msgs []model.Message) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func cloneMessageForContextCompaction(msg model.Message) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func cloneContentPartsForContextCompaction(
	parts []model.ContentPart,
) []model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func cloneContentPartForContextCompaction(
	part model.ContentPart,
) model.ContentPart {
	_ = "STUB: not implemented"
	return *new(model.ContentPart)
}

func cloneToolCallsForContextCompaction(
	toolCalls []model.ToolCall,
) []model.ToolCall {
	_ = "STUB: not implemented"
	return nil
}

func cloneGenerationConfigForContextCompaction(
	cfg model.GenerationConfig,
) model.GenerationConfig {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfig)
}

func cloneStructuredOutputForContextCompaction(
	out *model.StructuredOutput,
) *model.StructuredOutput {
	_ = "STUB: not implemented"
	return nil
}

func cloneJSONMapForContextCompaction(
	src map[string]any,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func snapshotSummary(sess *session.Session, filterKey string) summarySnapshot {
	_ = "STUB: not implemented"
	return *new(summarySnapshot)
}

func (s summarySnapshot) advanced(next summarySnapshot) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldSyncCompactContext(
	ctx context.Context,
	inv *agent.Invocation,
	req *model.Request,
	ratio float64,
	counter model.TokenCounter,
) bool {
	_ = "STUB: not implemented"
	return false
}

func contextCompactionThreshold(inv *agent.Invocation, ratio float64) int {
	_ = "STUB: not implemented"
	return 0
}

// UserToolsProvider is an optional interface that agents can implement to expose
// which tools were explicitly registered by the user (WithTools, WithToolSets)
// vs framework-added tools (Knowledge, SubAgents).
//
// User tools are subject to filtering via WithToolFilter.
// Framework tools are never filtered and always available to the agent.
type UserToolsProvider interface {
	UserTools() []tool.Tool
}

// ToolFilterProvider is an optional interface that agents can implement to provide
type ToolFilterProvider interface {
	FilterTools(ctx context.Context) []tool.Tool
}

// InvocationToolSurfaceProvider is an optional interface that exposes
// invocation-scoped tools and user-tool classification.
type InvocationToolSurfaceProvider interface {
	InvocationToolSurface(
		ctx context.Context,
		invocation *agent.Invocation,
	) ([]tool.Tool, map[string]bool)
}

// getFilteredTools returns the list of tools for this invocation after applying the filter.
//
// User tools (can be filtered):
//   - Tools registered via WithTools
//   - Tools registered via WithToolSets
//
// Framework tools (never filtered):
//   - transfer_to_agent (auto-added when SubAgents are configured)
//   - knowledge_search / agentic_knowledge_search (auto-added when Knowledge is configured)
//
// This method is called during the preprocess stage, before sending the request to the model.
func (f *Flow) getFilteredTools(ctx context.Context, invocation *agent.Invocation) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// Get user tools (if the agent supports it).
// User tools are those explicitly registered via WithTools and
// WithToolSets. Framework tools (Knowledge, SubAgents) are never filtered.

// If no filter is specified, return all tools for this invocation.

// Apply the filter function to each tool.
// Framework tools are never filtered.

// Determine if this is a user tool or framework tool.

// Framework tools are always included (never filtered).

// User tool: apply the filter function.

// Sort tools by name to ensure stable order for better prompt cache hit rate.
// Map iteration order is random in Go, so sorting ensures consistent tool ordering
// across requests, which improves cache efficiency.

func appendRunOptionTools(
	allTools []tool.Tool,
	userToolNames map[string]bool,
	hasUserToolTracking bool,
	opts agent.RunOptions,
) ([]tool.Tool, map[string]bool, bool, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil, false, nil
}

func appendRunOptionToolList(
	allTools []tool.Tool,
	userToolNames map[string]bool,
	hasUserToolTracking bool,
	seen map[string]bool,
	tools []tool.Tool,
) ([]tool.Tool, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectToolNames(tools []tool.Tool) map[string]bool { _ = "STUB: not implemented"; return nil }

func copyToolNames(src map[string]bool) map[string]bool { _ = "STUB: not implemented"; return nil }

func setVisibleExternalToolNames(
	invocation *agent.Invocation,
	tools []tool.Tool,
	externalNames map[string]bool,
) {
	_ = "STUB: not implemented"
	return
}

func toolName(tl tool.Tool) string { _ = "STUB: not implemented"; return "" }

func hasTrackedUserTool(
	tools []tool.Tool,
	hasUserToolTracking bool,
	userToolNames map[string]bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// callLLM performs the actual LLM call using core/model.
func (f *Flow) callLLM(
	ctx context.Context,
	invocation *agent.Invocation,
	llmRequest *model.Request,
	callModel model.Model,
) (context.Context, model.Seq[*model.Response], error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// Enforce optional per-invocation LLM call limit. When the limit is not
// configured (<= 0), this is a no-op and preserves existing behavior.

// Run before model callbacks if they exist.

func (f *Flow) runBeforeModelCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	llmRequest *model.Request,
) (context.Context, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func withInvocationContextIfMissing(ctx context.Context, invocation *agent.Invocation) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func invocationFromContextOrFallback(ctx context.Context, fallback *agent.Invocation) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func runBeforeModelCallbacksWith(
	ctx context.Context,
	invocation *agent.Invocation,
	llmRequest *model.Request,
	callbacks *model.Callbacks,
) (context.Context, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func wrapBeforeModelCallbacksWithInvocation(
	callbacks *model.Callbacks,
	invocation *agent.Invocation,
) *model.Callbacks {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flow) generateContentSeq(
	ctx context.Context,
	invocation *agent.Invocation,
	llmRequest *model.Request,
	callModel model.Model,
) (model.Seq[*model.Response], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeResponseIDs(seq model.Seq[*model.Response]) model.Seq[*model.Response] {
	_ = "STUB: not implemented"
	return nil
}

func normalizeResponseID(resp *model.Response, currentID *string) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Preserve one stable ID for the entire active response stream.

// postprocess handles post-LLM call processing using response processors.
func (f *Flow) postprocess(
	ctx context.Context,
	invocation *agent.Invocation,
	llmRequest *model.Request,
	llmResponse *model.Response,
	eventChan chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Run response processors - they send events directly to the channel.

// WaitEventTimeout returns the remaining time until the context deadline.
// If the context has no deadline, it returns the default event completion timeout.
func WaitEventTimeout(ctx context.Context) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
