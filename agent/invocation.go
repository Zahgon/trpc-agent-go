//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package agent

import (
	"context"
	"reflect"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"trpc.group/trpc-go/trpc-agent-go/artifact"
	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/internal/tracecapture"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/searchfilter"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// WaitNoticeWithoutTimeout is the timeout duration for waiting without timeout
	WaitNoticeWithoutTimeout = 0 * time.Second

	// AppendEventNoticeKeyPrefix is the prefix for append event notice keys
	AppendEventNoticeKeyPrefix = "append_event:"

	// BranchDelimiter is the delimiter for branch
	BranchDelimiter = "/"

	// EventFilterKeyDelimiter is the delimiter for event filter key
	EventFilterKeyDelimiter = "/"

	// flusherStateKey is the invocation state key used by flush.Attach.
	flusherStateKey = "__flush_session__"
	// barrierStateKey is the invocation state key used by internal barrier flag.
	barrierStateKey = "__graph_barrier__"
	// appenderStateKey is the invocation state key used by internal appender
	// attachment (see internal/state/appender).
	appenderStateKey = "__append_event__"

	// streamHubStateKey is the invocation state key used by the graph to
	// share ephemeral streams across node invocations within the same run.
	streamHubStateKey = "__graph_stream_hub__"
	// surfaceRootNodeIDStateKey stores one invocation's mounted surface root node id.
	surfaceRootNodeIDStateKey = "__trpc_agent_internal_surface_root_node_id_state__"
	// teamMemberTraceRootStateKey stores one invocation's mounted team member trace root.
	teamMemberTraceRootStateKey = "__trpc_agent_internal_team_member_trace_root_state__"

	// SyncSummaryIntraRunStateKey is set on the invocation by the
	// flow when sync intra-run summary is active.
	// Runner checks this key to skip redundant async summary
	// enqueue during the same run.
	SyncSummaryIntraRunStateKey = "__sync_summary_intra_run__"
)

// TransferInfo contains information about a pending agent transfer.
type TransferInfo struct {
	// TargetAgentName is the name of the agent to transfer control to.
	TargetAgentName string
	// Message is the message to send to the target agent.
	Message string
}

// Invocation represents the context for a flow execution.
type Invocation struct {
	// Agent is the agent that is being invoked.
	Agent Agent
	// AgentName is the name of the agent that is being invoked.
	AgentName string
	// InvocationID is the ID of the invocation.
	InvocationID string
	// Branch records agent execution chain information.
	// In multi-agent mode, this is useful for tracing agent execution trajectories.
	Branch string
	// EndInvocation is a flag that indicates if the invocation is complete.
	EndInvocation bool
	// Session is the session that is being used for the invocation.
	Session *session.Session
	// SessionService is the session service used by this invocation.
	SessionService session.Service
	// Model is the model that is being used for the invocation.
	Model model.Model
	// Message is the message that is being sent to the agent.
	Message model.Message
	// RunOptions is the options for the Run method.
	RunOptions RunOptions
	// TransferInfo contains information about a pending agent transfer.
	TransferInfo *TransferInfo

	// Plugins provides runner-scoped hooks applied to this invocation.
	Plugins PluginManager

	// StructuredOutput defines how the model should produce structured output for this invocation.
	StructuredOutput *model.StructuredOutput
	// StructuredOutputType is the Go type to unmarshal the final JSON into.
	StructuredOutputType reflect.Type

	// MemoryService is the service for managing memory.
	MemoryService memory.Service
	// ArtifactService is the service for managing artifacts.
	ArtifactService artifact.Service

	// noticeChannels is used to signal when events are written to the session.
	noticeChannels map[string]chan any
	noticeMu       *sync.Mutex

	// eventFilterKey is used to filter events for flow or agent
	eventFilterKey string

	// parent is the parent invocation, if any
	parent *Invocation
	// traceCapture stores the shared execution trace capture for one root run.
	traceCapture *tracecapture.Capture
	traceMu      sync.Mutex
	// entryPredecessorStepIDs stores the predecessor step ids passed to this invocation entry.
	entryPredecessorStepIDs []string
	// traceNodeID stores the mounted static root node id for this invocation.
	traceNodeID string

	// state stores invocation-scoped state data (lazy initialized).
	// Can be used by callbacks, middleware, or any invocation-scoped logic.
	state   map[string]any
	stateMu sync.RWMutex

	// MaxLLMCalls is an optional upper bound on the number of LLM calls
	// allowed for this invocation. When the value is:
	//   - > 0: the limit is enforced for this invocation.
	//   - <= 0: no limit is applied (default, preserves existing behavior).
	//
	// Typical usage:
	//   - LLMAgent copies its per-agent limits into these fields in setupInvocation.
	//   - Other agent implementations may set them explicitly when constructing
	//     invocations. If left at zero, IncLLMCallCount/IncToolIteration are no-ops.
	MaxLLMCalls int

	// MaxToolIterations is an optional upper bound on how many tool-call
	// iterations are allowed for this invocation. A "tool iteration" is defined
	// as an assistant response that contains tool calls and triggers the
	// FunctionCallResponseProcessor. When the value is:
	//   - > 0: the limit is enforced for this invocation.
	//   - <= 0: no limit is applied (default, preserves existing behavior).
	MaxToolIterations int

	// timingInfo stores timing information for the first LLM call in this invocation.
	timingInfo *model.TimingInfo

	// llmCallCount tracks how many LLM calls have been made in this invocation.
	// This is used together with MaxLLMCalls to enforce a per-invocation limit.
	// Note: counters are invocation-scoped. When child invocations are created
	// via Clone (for example, in transfer_to_agent or AgentTool), the counters
	// start from zero for each invocation.
	llmCallCount int

	// toolIterationCount tracks how many tool call iterations have been processed
	// in this invocation. This is used together with MaxToolIterations
	// to guard against unbounded tool_call -> LLM -> tool_call loops.
	toolIterationCount int
}

// DefaultWaitNoticeTimeoutErr is the default error returned when a wait notice times out.
var DefaultWaitNoticeTimeoutErr = NewWaitNoticeTimeoutError("wait notice timeout.")

// WaitNoticeTimeoutError represents an error that signals the wait notice timeout.
type WaitNoticeTimeoutError struct {
	// Message contains the stop reason
	Message string
}

// Error implements the error interface.
func (e *WaitNoticeTimeoutError) Error() string {
	_ = "STUB: not implemented"

	// AsWaitNoticeTimeoutError checks if an error is a AsWaitNoticeTimeoutError using errors.As.
	return ""
}

func AsWaitNoticeTimeoutError(err error) (*WaitNoticeTimeoutError, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// NewWaitNoticeTimeoutError creates a new AsWaitNoticeTimeoutError with the given message.
func NewWaitNoticeTimeoutError(message string) *WaitNoticeTimeoutError {
	_ = "STUB: not implemented"
	return nil
}

// RunOption is a function that configures a RunOptions.
type RunOption func(*RunOptions)

// ModelSelector selects the model for one framework-managed LLM call.
// The invocation's Model is the base model for this call when the selector is
// invoked. Returning nil with nil error keeps that base model. Returning an
// error fails the current call before the request is built. A selector may be
// called concurrently by different runs and must protect any shared state it
// owns.
type ModelSelector func(ctx context.Context, inv *Invocation) (model.Model, error)

type runControlConfig struct {
	DisableGraphCompletionEvent bool
	DisableGraphExecutorEvents  bool
	EventChannelBufferSize      int
	PropagateChildAgentErrors   bool
}

// NewRunOptions builds a RunOptions value from RunOption functions.
func NewRunOptions(opts ...RunOption) RunOptions {
	_ = "STUB: not implemented"
	return *new(RunOptions)
}

// TraceStartedCallback receives the root span context for a run.
type TraceStartedCallback func(oteltrace.SpanContext)

// WithAppName overrides the runner's default app name for this specific run.
//
// This enables a single runner to serve multiple projects or tenants
// by isolating session and memory data under different app names.
// When not set, the runner uses its constructor-provided default app name.
func WithAppName(name string) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithRuntimeState sets the runtime state for the RunOptions.
func WithRuntimeState(state map[string]any) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithModelRequestExtraFields merges provider-specific top-level fields into
// each model request created during this run. Request-level fields take
// precedence over model-level extra fields in adapters that support merging.
func WithModelRequestExtraFields(fields map[string]any) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// MergeRuntimeState merges runtime state into existing RunOptions state.
//
// When a key already exists, the new value replaces the old one.
func MergeRuntimeState(state map[string]any) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithAgent sets the agent instance for this run only.
func WithAgent(a Agent) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithAgentByName sets the agent name that should be resolved for this run.
func WithAgentByName(name string) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// GetRuntimeStateValue retrieves a typed value from the runtime state.
//
// Returns the typed value and true if the key exists and the type matches,
// or the zero value and false otherwise.
//
// Example:
//
//	if userID, ok := GetRuntimeStateValue[string](&inv.RunOptions, "user_id"); ok {
//	    log.Printf("User ID: %s", userID)
//	}
//	if roomID, ok := GetRuntimeStateValue[int](&inv.RunOptions, "room_id"); ok {
//	    log.Printf("Room ID: %d", roomID)
//	}
func GetRuntimeStateValue[T any](opts *RunOptions, key string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// WithKnowledgeFilter sets the metadata filter for the RunOptions.
func WithKnowledgeFilter(filter map[string]any) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithKnowledgeConditionedFilter sets the complex condition filter for the RunOptions.
func WithKnowledgeConditionedFilter(filter *searchfilter.UniversalFilterCondition) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithMessages sets the caller-supplied conversation history for this run.
// Runner uses this history to auto-seed an empty Session (once) and to
// populate `invocation.Message` via RunWithMessages for compatibility. The
// content processor itself does not read this field; it derives messages from
// Session events and may fall back to a single `invocation.Message` when the
// Session is empty.
func WithMessages(messages []model.Message) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithInjectedContextMessages appends per-run messages that are injected into the
// model request context but are not persisted into the session transcript.
func WithInjectedContextMessages(messages []model.Message) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// UserMessageRewriteArgs contains stable metadata for one user message rewrite.
type UserMessageRewriteArgs struct {
	AppName         string
	UserID          string
	SessionID       string
	RequestID       string
	OriginalMessage model.Message
}

// UserMessageRewriter rewrites one current-turn user input into an ordered
// message sequence. The returned order is the persistence order for the turn,
// and the last message becomes invocation.Message.
type UserMessageRewriter func(
	ctx context.Context,
	args *UserMessageRewriteArgs,
) ([]model.Message, error)

// WithUserMessageRewriter rewrites the current-turn input into an ordered
// message sequence before runner persists it into the session transcript.
func WithUserMessageRewriter(rewriter UserMessageRewriter) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithResume enables or disables resume mode for this run.
// When enabled, flows like llmflow may inspect the existing Session history
// and resume unfinished work (for example, executing pending tool calls)
// before issuing a new model call.
func WithResume(enabled bool) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithPersistInterruptedAssistant controls whether a cancelled streaming run
// persists already-emitted assistant text as a final assistant message.
//
// By default this is not set, so the Runner uses its own default. The built-in
// Runner default is false to preserve the "cancel discards partial text"
// session semantics.
func WithPersistInterruptedAssistant(enabled bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithGraphEmitFinalModelResponses controls whether graph-based agents emit
// final (Done=true) model responses as events.
//
// When disabled (default), graph Large Language Model (LLM) nodes only emit
// streaming chunks (Done=false), which matches the pre-#901 behavior.
func WithGraphEmitFinalModelResponses(enabled bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithGraphTerminalMessagesOnly limits caller-visible graph message events to
// terminal nodes only.
//
// When disabled (default), all graph Large Language Model (LLM) nodes and
// sub-agent nodes may emit caller-visible message events.
func WithGraphTerminalMessagesOnly(enabled bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithStreamMode sets StreamMode selection for this run.
//
// When StreamModeMessages is present, graph-based Large Language Model (LLM)
// nodes will also emit their final (Done=true) model responses by default.
// If you need to override that behavior, call WithGraphEmitFinalModelResponses
// after WithStreamMode.
func WithStreamMode(modes ...StreamMode) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithDisableGraphCompletionEvent disables emitting the final graph completion event.
func WithDisableGraphCompletionEvent(disable bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithDisableGraphExecutorEvents disables emitting graph executor lifecycle events.
func WithDisableGraphExecutorEvents(disable bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithEventChannelBufferSize overrides the event channel buffer size for this run
// on supported flow and agent implementations.
//
// When size <= 0, supported implementations use their configured default.
func WithEventChannelBufferSize(size int) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithPropagateChildAgentErrors enables strict propagation for terminal child
// agent errors observed through agent-node event streams.
//
// When disabled (default), agent nodes preserve the legacy compatibility
// behavior: child error events remain observable in the stream but do not
// automatically fail the parent graph.
func WithPropagateChildAgentErrors(enabled bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithDisableTracing requests supported agent and flow execution paths to skip
// creating OpenTelemetry spans for this run.
func WithDisableTracing(disable bool) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithDisableResponseUsageTracking disables attaching usage and timing info to streaming responses.
func WithDisableResponseUsageTracking(disable bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithDisableModelExecutionEvents disables emitting model execution events for this run.
func WithDisableModelExecutionEvents(disable bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithDisablePartialEventIDs disables generating IDs for partial response events.
func WithDisablePartialEventIDs(disable bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithDisablePartialEventTimestamps disables generating timestamps for partial response events.
func WithDisablePartialEventTimestamps(disable bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithRequestID sets the request id for the RunOptions.
func WithRequestID(requestID string) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithEventFilterKey sets the invocation event filter key for this run.
//
// This controls the FilterKey injected into emitted events and the default
// filter prefix used by ContentRequestProcessor when building LLM context.
func WithEventFilterKey(filterKey string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithDetachedCancel enables running a job that ignores parent context
// cancellation.
//
// When enabled, Runner will remove the cancellation signal from the
// execution context while still preserving context values and enforcing
// timeouts and deadlines.
func WithDetachedCancel(enabled bool) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithMaxRunDuration sets the maximum duration for a single run.
//
// Runner will enforce the smaller of:
//   - the parent context deadline (if any)
//   - MaxRunDuration (if > 0)
func WithMaxRunDuration(d time.Duration) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithSpanAttributes sets custom span attributes for the RunOptions.
func WithSpanAttributes(attrs ...attribute.KeyValue) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithTraceStartedCallback registers a callback for the run root span.
func WithTraceStartedCallback(
	callback TraceStartedCallback,
) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithModel sets the model for this specific run.
// This allows temporarily switching the model for a single request without
// affecting other requests or the agent's default model configuration.
//
// Example:
//
//	runner.Run(ctx, userID, sessionID, message,
//	    agent.WithModel(customModel),
//	)
func WithModel(m model.Model) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithModelName sets the model name for this specific run.
// The agent will look up the model by name from its registered models.
// This is useful when the agent has multiple models registered via WithModels.
//
// Example:
//
//	runner.Run(ctx, userID, sessionID, message,
//	    agent.WithModelName("gpt-4"),
//	)
func WithModelName(name string) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithModelContextWindow sets the model context window for this specific run.
// This is useful for user-defined or private models whose names should not be
// registered in the process-wide model registry.
func WithModelContextWindow(tokens int) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// ModelContextWindowFromRunOptions returns the context window configured by
// WithModelContextWindow.
func ModelContextWindowFromRunOptions(opts *RunOptions) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// WithModelSelector sets the model selector for this specific run.
// The selector is called before each framework-managed LLM call and takes
// precedence over any agent-level selector.
func WithModelSelector(selector ModelSelector) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithCodeExecutor sets the code executor for this specific run.
// If set, it temporarily overrides the agent's default code executor for this
// request only.
func WithCodeExecutor(exec codeexecutor.CodeExecutor) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithStream enables or disables streaming for this specific run.
//
// When set, it overrides the agent's default Stream setting for this Run.
func WithStream(stream bool) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

// WithInstruction sets the instruction for this specific run.
// If set, it temporarily overrides the agent's instruction for this request
// only. This does not modify the agent instance.
func WithInstruction(instruction string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithGlobalInstruction sets the global instruction (system prompt) for this
// specific run.
// If set, it temporarily overrides the agent's global instruction for this
// request only. This does not modify the agent instance.
func WithGlobalInstruction(instruction string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithStructuredOutputJSONSchema sets a JSON schema structured output for this run.
func WithStructuredOutputJSONSchema(name string, schema map[string]any, strict bool, description string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithStructuredOutputJSON sets a JSON schema structured output for this run.
// The schema is constructed automatically from the provided example type.
func WithStructuredOutputJSON(examplePtr any, strict bool, description string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func newStructuredOutput(name string, schema map[string]any, strict bool, description string) *model.StructuredOutput {
	_ = "STUB: not implemented"
	return nil
}

// WithToolFilter sets a custom tool filter function for this specific run.
// The filter function receives a context and a tool, and returns true if the tool should be included.
//
// This is useful for:
//   - Permission control: restrict tool access based on user roles or runtime conditions
//   - Cost optimization: reduce token usage by limiting tool descriptions
//   - Feature isolation: limit capabilities for specific use cases
//   - Dynamic filtering: filter tools based on runtime state, session data, etc.
//
// Example - Simple name-based filtering:
//
//	runner.Run(ctx, userID, sessionID, message,
//	    agent.WithToolFilter(tool.NewIncludeToolNamesFilter("calculator", "time_tool")),
//	)
//
// Example - Custom logic with runtime state:
//
//	runner.Run(ctx, userID, sessionID, message,
//	    agent.WithToolFilter(func(ctx context.Context, t tool.Tool) bool {
//	        // Access invocation from context if needed
//	        inv, _ := agent.InvocationFromContext(ctx)
//	        userLevel, _ := inv.Session.Get("user_level").(string)
//
//	        // Premium users get all tools
//	        if userLevel == "premium" {
//	            return true
//	        }
//
//	        // Free users only get basic tools
//	        toolName := t.Declaration().Name
//	        return toolName == "calculator" || toolName == "time_tool"
//	    }),
//	)
//
// Note: Framework tools (knowledge_search, transfer_to_agent) are never filtered
// and will always be available regardless of the filter function.
//
// Note: This is a "soft" constraint. Tools should still implement their own
// authorization logic for security.
func WithToolFilter(filter tool.FilterFunc) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithAdditionalTools appends tools that are visible only for this run.
//
// Additional tools are treated as user tools, so WithToolFilter can still
// hide them. If an additional tool has the same name as an already available
// tool, the already available tool wins for that run.
func WithAdditionalTools(tools []tool.Tool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithExternalTools appends caller-executed tools for this run.
//
// External tools are visible to the model like additional tools, but the
// framework will not execute them. When the model calls one, the run stops
// after the assistant tool_call response. The caller should execute the tool
// externally and continue with model.NewToolMessage.
func WithExternalTools(tools []tool.Tool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithToolExecutionFilter sets which tools the framework will execute.
//
// This is different from WithToolFilter:
//   - WithToolFilter controls which tools are visible to the model.
//   - WithToolExecutionFilter controls which tool calls are auto-executed
//     after the model requests them.
//
// When the filter returns false for a tool, the tool call is not executed
// and the run ends after emitting the assistant tool_call response. The
// caller can then execute the tool externally and provide a RoleTool
// message with the tool result to continue.
func WithToolExecutionFilter(filter tool.FilterFunc) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func appendRunTools(opts *RunOptions, tools []tool.Tool) { _ = "STUB: not implemented"; return }

func declarationName(tl tool.Tool) string { _ = "STUB: not implemented"; return "" }

// WithToolCallArgumentsJSONRepairEnabled enables best-effort JSON repair for tool call arguments.
func WithToolCallArgumentsJSONRepairEnabled(enabled bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithA2ARequestOptions sets the A2A request options for the RunOptions.
// These options will be passed to A2A agent's SendMessage and StreamMessage calls.
// This allows passing dynamic HTTP headers or other request-specific options for each run.
func WithA2ARequestOptions(opts ...any) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithCustomAgentConfigs sets custom agent configurations.
// This allows passing agent-specific configurations at runtime without modifying the agent implementation.
//
// Parameters:
//   - configs: A map where the key is the agent type identifier and the value is the agent-specific config.
//     It's recommended to use the agent's defined RunOptionKey constant as the key and a typed options struct as the value.
//
// Usage:
//
//	// Example: Configure a custom LLM agent using its defined key and options struct
//	import customllm "your.module/agents/customllm"
//
//	runner.Run(ctx, userID, sessionID, message,
//	    agent.WithCustomAgentConfigs(map[string]any{
//	        customllm.RunOptionKey: customllm.RunOptions{
//	            "custom-context": "context",
//	        },
//	    }),
//	)
//
//
//	// In your custom agent implementation, retrieve the config:
//	func (a *CustomLLMAgent) Run(ctx context.Context, inv *agent.Invocation) (<-chan *event.Event, error) {
//	    config := inv.GetCustomAgentConfig(RunOptionKey)
//	    if opts, ok := config.(RunOptions); ok {
//	        client := NewLLMClient(opts.APIKey, opts.Model, opts.Temperature)
//	        // Use the configuration...
//	    }
//	    // ...
//	}
//
// Note:
//   - This function creates a shallow copy of the configs map to prevent external modifications.
//   - The stored configuration should be treated as read-only. Do not modify it after retrieval.
func WithCustomAgentConfigs(configs map[string]any) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// Create a shallow copy to prevent external modifications

func getRunControlConfig(opts *RunOptions) runControlConfig {
	_ = "STUB: not implemented"
	return *new(runControlConfig)
}

func setRunControlConfig(opts *RunOptions, cfg runControlConfig) { _ = "STUB: not implemented"; return }

// RunOptions is the options for the Run method.
type RunOptions struct {
	// AppName overrides the runner's default app name for this specific run.
	//
	// When set, the runner uses this value instead of its constructor-provided
	// app name for session keys, memory operations, and event filter keys.
	// This enables a single runner instance to serve multiple projects or
	// tenants, isolating their session and memory data by app name.
	//
	// If empty, the runner falls back to its default app name.
	AppName string

	// RuntimeState contains key-value pairs that will be merged into the initial state
	// for this specific run. This allows callers to pass dynamic parameters
	// (e.g., room ID, user context) without modifying the agent's base initial state.
	RuntimeState map[string]any

	// EventFilterKey overrides the invocation's event filter key used for
	// scoping session events (event.FilterKey) included in LLM context.
	//
	// Runner applies this value via WithInvocationEventFilterKey when it
	// constructs the invocation. When using Runner, the value should
	// typically start with the runner app name (e.g., "<appName>/...") so
	// sessions hooks and summaries continue to work as expected.
	EventFilterKey string

	// KnowledgeFilter contains metadata key-value pairs for the knowledge filter
	KnowledgeFilter map[string]any

	// KnowledgeConditionedFilter contains complex condition filter for the knowledge search
	KnowledgeConditionedFilter *searchfilter.UniversalFilterCondition

	// Messages allows callers to provide a full conversation history to Runner.
	// Runner will seed an empty Session with this history automatically and
	// then rely on Session events for subsequent turns. The content processor
	// ignores this field and reads only from Session events (or falls back to
	// `invocation.Message` when no events exist).
	Messages []model.Message

	// InjectedContextMessages allows callers to inject additional context messages
	// into the model request for this run. These messages are not persisted into
	// session events and therefore must be provided on every run if needed.
	InjectedContextMessages []model.Message

	// UserMessageRewriter rewrites the current-turn input into an ordered
	// message sequence before runner persists it into the session transcript.
	UserMessageRewriter UserMessageRewriter

	// Resume indicates whether this run should attempt to resume from existing
	// session context before making a new model call. When true, flows may
	// inspect the latest session events (for example, assistant messages with
	// pending tool calls) and complete unfinished work prior to issuing a new
	// LLM request.
	Resume bool

	// PersistInterruptedAssistant controls whether Runner persists already
	// emitted assistant text as a final assistant message when a streaming run
	// is cancelled before a normal final assistant response is produced.
	//
	// nil means the Runner default applies. The built-in Runner default is
	// false to preserve the legacy cancellation semantics.
	PersistInterruptedAssistant *bool

	// GraphEmitFinalModelResponses controls event emission for graph-based
	// Large Language Model (LLM) nodes.
	//
	// When false (default), graph LLM nodes only emit streaming chunks
	// (Done=false).
	//
	// When true, graph LLM nodes also emit the final model response
	// (Done=true). In that mode, callers should be prepared to receive
	// assistant messages from intermediate nodes.
	//
	// When enabled, Runner may omit echoing the final assistant message
	// in its runner-completion event to avoid duplicates.
	GraphEmitFinalModelResponses bool

	// GraphTerminalMessagesOnly limits caller-visible graph message events to
	// terminal nodes only.
	//
	// When false (default), graph message-capable nodes may all emit
	// caller-visible events.
	//
	// When true, GraphAgent keeps internal state propagation unchanged, but
	// caller-visible message events are forwarded only for terminal LLM nodes
	// and terminal sub-agent nodes.
	GraphTerminalMessagesOnly bool

	// StreamModeEnabled indicates whether the caller explicitly configured
	// StreamModes for this run.
	StreamModeEnabled bool

	// StreamModes selects which categories of events are forwarded to callers.
	//
	// When StreamModeEnabled is false, runners should not apply any stream
	// filtering and preserve the existing behavior.
	StreamModes []StreamMode

	// DisableTracing requests supported agent and flow execution paths to skip
	// creating OpenTelemetry spans for this run.
	DisableTracing bool

	// DisableResponseUsageTracking disables attaching usage and timing info to streaming responses.
	DisableResponseUsageTracking bool

	// DisableModelExecutionEvents disables emitting model execution start/complete events.
	DisableModelExecutionEvents bool

	// DisablePartialEventIDs disables generating IDs for partial response events.
	DisablePartialEventIDs bool

	// DisablePartialEventTimestamps disables generating timestamps for partial response events.
	DisablePartialEventTimestamps bool

	// ExecutionTraceEnabled enables in-process execution trace recording for this run.
	ExecutionTraceEnabled bool

	// RequestID is the request id of the request.
	RequestID string

	// DetachedCancel controls whether Runner ignores parent context
	// cancellation for this run.
	DetachedCancel bool

	// MaxRunDuration bounds the total execution time for this run.
	// When set, Runner enforces the smaller of:
	//   - the parent context deadline (if any)
	//   - MaxRunDuration
	MaxRunDuration time.Duration

	// SpanAttributes carries custom span attributes for this run.
	SpanAttributes []attribute.KeyValue

	// TraceStartedCallbacks run when the root span starts for this run.
	TraceStartedCallbacks []TraceStartedCallback

	// A2ARequestOptions contains A2A client request options that will be passed to
	// A2A agent's SendMessage and StreamMessage calls. This allows callers to pass
	// dynamic HTTP headers or other request-specific options for each run.
	//
	// Note: This field uses any type to avoid direct dependency on trpc-a2a-go/client package.
	// Users should pass client.RequestOption values (e.g., client.WithRequestHeader).
	// The a2aagent package will validate the option types at runtime.
	A2ARequestOptions []any

	// CustomAgentConfigs stores configurations for custom agents.
	// Key: agent type, Value: agent-specific config.
	CustomAgentConfigs map[string]any

	// Agent overrides the runner's default agent for this run.
	Agent Agent

	// AgentByName instructs the runner to resolve an agent by name for this run.
	AgentByName string

	// Model is the model to use for this specific run.
	// If set, it temporarily overrides the agent's default model for this request only.
	// This allows per-request model switching without affecting other concurrent requests.
	Model model.Model

	// ModelName is the name of the model to use for this specific run.
	// The agent will look up the model by name from its registered models.
	// If both Model and ModelName are set, Model takes precedence.
	ModelName string
	// ModelSelector selects the model before each framework-managed LLM call.
	ModelSelector ModelSelector

	// ModelContextWindow is the model context window for this specific run.
	// If set, it takes precedence over model instance configuration and the
	// process-wide model registry.
	ModelContextWindow int

	// ModelRequestExtraFields contains provider-specific top-level request body
	// fields for model calls made during this run.
	//
	// Adapters that support extra fields merge these with model-level extra
	// fields, with these request-level values taking precedence.
	ModelRequestExtraFields map[string]any

	// CodeExecutor is the code executor to use for this specific run.
	// If set, it temporarily overrides the agent's default code executor for
	// this request only.
	CodeExecutor codeexecutor.CodeExecutor

	// Stream overrides GenerationConfig.Stream for this run when non-nil.
	//
	// This is useful when you want to switch between streaming and
	// non-streaming responses per request without rebuilding the agent.
	Stream *bool

	// Instruction overrides the agent's instruction for this run.
	// If set, it temporarily overrides the agent's instruction for this request
	// only.
	Instruction string

	// GlobalInstruction overrides the agent's global instruction (system prompt)
	// for this run.
	// If set, it temporarily overrides the agent's global instruction for
	// this request only.
	GlobalInstruction string

	// StructuredOutput defines how the model should produce structured output for this run.
	StructuredOutput *model.StructuredOutput

	// StructuredOutputType is the Go type to unmarshal the final JSON into for this run.
	StructuredOutputType reflect.Type

	// ToolFilter is a custom function to filter tools for this run.
	// If set, only tools for which the filter returns true will be available to the model.
	// If nil, all registered tools will be available (default behavior).
	//
	// The filter function receives:
	//   - ctx: The context with invocation information (use agent.InvocationFromContext)
	//   - tool: The tool being filtered
	//
	// This filtering happens at the request preparation stage, before sending to the model.
	// The model will only see the tool descriptions for tools that pass the filter.
	//
	// Note: Framework tools (knowledge_search, transfer_to_agent) are never filtered
	// and will always be included regardless of the filter function's return value.
	//
	// Example:
	//   agent.WithToolFilter(tool.NewIncludeToolNamesFilter("calculator", "time_tool"))
	//   agent.WithToolFilter(func(ctx context.Context, t tool.Tool) bool {
	//       return t.Declaration().Name == "calculator"
	//   })
	ToolFilter tool.FilterFunc

	// AdditionalTools contains tools that are visible only for this run.
	//
	// These tools are treated as user tools and are therefore affected by
	// ToolFilter. They are appended to the effective tool surface without
	// mutating the agent's registered tools.
	AdditionalTools []tool.Tool

	// ExternalTools contains caller-executed tools that are visible only for
	// this run. The framework exposes them to the model, but does not execute
	// them after the model returns a tool call.
	ExternalTools []tool.Tool

	// ExternalToolNames contains the accepted caller-executed tool names for
	// this run. LLM flows set it after the invocation tool surface rejects
	// collisions with existing tools.
	ExternalToolNames map[string]bool

	// ToolExecutionFilter controls which tools are executed by the
	// framework when the model returns tool calls.
	//
	// This is different from ToolFilter:
	//   - ToolFilter controls which tools are sent to (and callable by) the
	//     model.
	//   - ToolExecutionFilter controls which tool calls are auto-executed by
	//     the framework after the model requests them.
	//
	// When this filter is set and returns false for a tool, the tool call
	// is not executed by the agent. The run stops after emitting the
	// assistant tool_call response so the caller can execute the tool
	// externally and later provide tool results (RoleTool messages).
	ToolExecutionFilter tool.FilterFunc
	// ToolCallArgumentsJSONRepairEnabled enables best-effort JSON repair for tool call arguments.
	// When nil, JSON repair is disabled by default.
	ToolCallArgumentsJSONRepairEnabled *bool

	// runControlConfig stores internal event and buffering controls.
	runControlConfig runControlConfig
}

// ShouldExecuteTool reports whether the framework should execute a tool call.
//
// External tools are always caller-executed and therefore return false. The
// ToolExecutionFilter is evaluated only for non-external tools.
func (opts RunOptions) ShouldExecuteTool(
	ctx context.Context,
	tl tool.Tool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (opts RunOptions) isExternalTool(tl tool.Tool) bool { _ = "STUB: not implemented"; return false }

func sameRunTool(a tool.Tool, b tool.Tool) bool { _ = "STUB: not implemented"; return false }

// IsGraphCompletionEventDisabled reports whether this invocation hides terminal graph completion events.
func IsGraphCompletionEventDisabled(inv *Invocation) bool { _ = "STUB: not implemented"; return false }

// IsGraphExecutorEventsDisabled reports whether this invocation hides graph executor lifecycle events.
func IsGraphExecutorEventsDisabled(inv *Invocation) bool { _ = "STUB: not implemented"; return false }

// GetEventChannelBufferSize returns the invocation-specific event channel buffer size override.
func GetEventChannelBufferSize(inv *Invocation) int { _ = "STUB: not implemented"; return 0 }

// ShouldPropagateChildAgentErrors reports whether terminal child agent errors
// should fail the parent graph by default.
func ShouldPropagateChildAgentErrors(inv *Invocation) bool { _ = "STUB: not implemented"; return false }

// NewInvocation create a new invocation
func NewInvocation(invocationOpts ...InvocationOptions) *Invocation {
	_ = "STUB: not implemented"
	return nil
}

// Clone clone a new invocation
func (inv *Invocation) Clone(invocationOpts ...InvocationOptions) *Invocation {
	_ = "STUB: not implemented"
	return nil
}

// seted by WithInvocationBranch

// View returns an isolated invocation view that preserves identity.
func (inv *Invocation) View(invocationOpts ...InvocationOptions) *Invocation {
	_ = "STUB: not implemented"
	return nil
}

// SyncView copies execution-visible state from a view while preserving RunOptions.
func (inv *Invocation) SyncView(view *Invocation) { _ = "STUB: not implemented"; return }

func (inv *Invocation) cloneState() map[string]any { _ = "STUB: not implemented"; return nil }

func (inv *Invocation) cloneViewState() map[string]any { _ = "STUB: not implemented"; return nil }

func (inv *Invocation) cloneStateByFilter(
	include func(string) bool,
	cloneValue func(string, any) any,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func includeAllStateKeys(string) bool { _ = "STUB: not implemented"; return false }

func keepStateValue(_ string, value any) any { _ = "STUB: not implemented"; return *new(any) }

func cloneViewStateValue(key string, value any) any { _ = "STUB: not implemented"; return *new(any) }

func isCloneStateKey(key string) bool { _ = "STUB: not implemented"; return false }

// cloneStateValue isolates common mutable custom state for invocation views.
// Known mutable types such as bytes.Buffer, strings.Builder, and big.Int are
// copied explicitly. Maps, slices, pointers, arrays, and fully exported
// structs are cloned recursively. Opaque structs with unexported fields are
// kept by reference to avoid unsafe copies of no-copy state such as locks.
func cloneStateValue(value any) any { _ = "STUB: not implemented"; return *new(any) }

type reflectVisit struct {
	typ      reflect.Type
	ptr      uintptr
	length   int
	capacity int
}

func cloneStateReflectValue(
	value reflect.Value,
	visited map[reflectVisit]reflect.Value,
) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func cloneKnownStateValue(value any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func cloneBytes(value []byte) []byte { _ = "STUB: not implemented"; return nil }

func cloneStatePointerValue(
	value reflect.Value,
	visited map[reflectVisit]reflect.Value,
) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func cloneStateAsType(
	value reflect.Value,
	typ reflect.Type,
) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func cloneStateMapValue(
	value reflect.Value,
	visited map[reflectVisit]reflect.Value,
) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// Keep keys unchanged; cloning pointer keys changes lookup identity.

func cloneStateElement(
	value reflect.Value,
	elemType reflect.Type,
	visited map[reflectVisit]reflect.Value,
) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func cloneStateSliceValue(
	value reflect.Value,
	visited map[reflectVisit]reflect.Value,
) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func cloneStateArrayValue(
	value reflect.Value,
	visited map[reflectVisit]reflect.Value,
) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func cloneStateStructValue(
	value reflect.Value,
	visited map[reflectVisit]reflect.Value,
) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// Opaque structs may carry no-copy state such as locks.

// GetEventFilterKey get event filter key.
func (inv *Invocation) GetEventFilterKey() string { _ = "STUB: not implemented"; return "" }

// GetParentInvocation get parent invocation.
func (inv *Invocation) GetParentInvocation() *Invocation { _ = "STUB: not implemented"; return nil }

// InjectIntoEvent inject invocation information into event.
func InjectIntoEvent(inv *Invocation, e *event.Event) { _ = "STUB: not implemented"; return }

// EmitEvent inject invocation information into event and emit it to channel.
func EmitEvent(ctx context.Context, inv *Invocation, ch chan<- *event.Event,
	e *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAppendEventNoticeKey get append event notice key.
func GetAppendEventNoticeKey(eventID string) string { _ = "STUB: not implemented"; return "" }

// SetState sets a value in the invocation state.
//
// This is a general-purpose key-value store scoped to the invocation lifecycle.
// It can be used by callbacks, middleware, or any invocation-scoped logic.
//
// Recommended key naming conventions:
//   - Agent callbacks: "agent:xxx" (e.g., "agent:start_time")
//   - Model callbacks: "model:xxx" (e.g., "model:start_time")
//   - Tool callbacks: "tool:<toolName>:<toolCallID>:xxx" (e.g., "tool:calculator:call_abc123:start_time")
//   - Middleware: "middleware:xxx" (e.g., "middleware:request_id")
//   - Custom logic: "custom:xxx" (e.g., "custom:user_context")
//
// Note: Tool callbacks should include tool call ID to support concurrent calls.
//
// Example:
//
//	inv.SetState("agent:start_time", time.Now())
//	inv.SetState("model:start_time", time.Now())
//	inv.SetState("tool:calculator:call_abc123:start_time", time.Now())
//	inv.SetState("middleware:request_id", "req-123")
//	inv.SetState("custom:user_context", userCtx)
func (inv *Invocation) SetState(key string, value any) { _ = "STUB: not implemented"; return }

// GetState retrieves a value from the invocation state.
//
// Returns the value and true if the key exists, or nil and false otherwise.
//
// Example:
//
//	if startTime, ok := inv.GetState("agent:start_time"); ok {
//	    duration := time.Since(startTime.(time.Time))
//	}
//	if startTime, ok := inv.GetState("tool:calculator:call_abc123:start_time"); ok {
//	    duration := time.Since(startTime.(time.Time))
//	}
func (inv *Invocation) GetState(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// GetStateValue retrieves a typed value from the invocation state.
//
// Returns the typed value and true if the key exists and the type matches,
// or the zero value and false otherwise.
//
// Example:
//
//	if startTime, ok := GetStateValue[time.Time](inv, "agent:start_time"); ok {
//	    duration := time.Since(startTime)
//	}
//	if requestID, ok := GetStateValue[string](inv, "middleware:request_id"); ok {
//	    log.Printf("Request ID: %s", requestID)
//	}
func GetStateValue[T any](inv *Invocation, key string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// GetOrCreateTimingInfo gets or creates timing info for this invocation.
// Only the first LLM call will create and populate timing info; subsequent calls reuse it.
// This ensures timing metrics only reflect the first LLM call in scenarios with multiple calls (e.g., tool calls).
func (inv *Invocation) GetOrCreateTimingInfo() *model.TimingInfo {
	_ = "STUB: not implemented"
	return nil
}

// IncLLMCallCount increments the LLM call counter for this invocation and
// enforces the optional MaxLLMCalls limit. When the limit is not set or
// non-positive, no restriction is applied. When the limit is exceeded, a
// StopError is returned so callers can terminate the flow early.
func (inv *Invocation) IncLLMCallCount() error { _ = "STUB: not implemented"; return nil }

// No limit configured, preserve existing behavior.

// IncToolIteration increments the tool iteration counter and reports whether
// the MaxToolIterations limit has been exceeded. A "tool iteration" is
// defined as an assistant response that contains tool calls and triggers the
// FunctionCallResponseProcessor. When the limit is not set or non-positive,
// this method always returns false, preserving existing behavior.
func (inv *Invocation) IncToolIteration() bool { _ = "STUB: not implemented"; return false }

// No limit configured, preserve existing behavior.

// DeleteState removes a value from the invocation state.
//
// Example:
//
//	inv.DeleteState("agent:start_time")
//	inv.DeleteState("tool:calculator:call_abc123:start_time")
func (inv *Invocation) DeleteState(key string) { _ = "STUB: not implemented"; return }

// AddNoticeChannelAndWait add notice channel and wait it complete
func (inv *Invocation) AddNoticeChannelAndWait(ctx context.Context, key string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// no timeout, maybe wait for ever

// AddNoticeChannel add a new notice channel
func (inv *Invocation) AddNoticeChannel(ctx context.Context, key string) chan any {
	_ = "STUB: not implemented"
	return nil
}

// NotifyCompletion notify completion signal to waiting task
func (inv *Invocation) NotifyCompletion(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// channel not found, create a new one and close it.
// May involve notification followed by waiting.

// channel found, close it if it's not closed

// CleanupNotice cleanup all notice channel
// The 'Invocation' instance created via the NewInvocation method ​​should be disposed​​
// upon completion to prevent resource leaks.
func (inv *Invocation) CleanupNotice(ctx context.Context) { _ = "STUB: not implemented"; return }

// GetCustomAgentConfig retrieves configuration for a specific custom agent type.
//
// Parameters:
//   - agentKey: The agent type identifier (typically the agent's RunOptionKey constant)
//
// Returns:
//   - The configuration value if found, nil otherwise
//
// Usage:
//
//	func (a *CustomLLMAgent) Run(ctx context.Context, inv *agent.Invocation) (<-chan *event.Event, error) {
//	    config := inv.GetCustomAgentConfig(RunOptionKey)
//	    if opts, ok := config.(RunOptions); ok {
//	        client := NewLLMClient(opts.APIKey, opts.Model)
//	        // ...
//	    }
//	}
//
// Note: The returned config should be treated as read-only. Do not modify it.
func (inv *Invocation) GetCustomAgentConfig(agentKey string) any {
	_ = "STUB: not implemented"
	return *new(any)
}
