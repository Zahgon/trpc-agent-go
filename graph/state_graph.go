//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"context"
	"time"

	oteltrace "go.opentelemetry.io/otel/trace"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/internal/responseusage"
	itelemetry "trpc.group/trpc-go/trpc-agent-go/internal/telemetry"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// StateGraph provides a fluent interface for building graphs.
// This is the primary public API for creating executable graphs.
//
// StateGraph provides:
//   - Type-safe state management with schemas and reducers
//   - Conditional routing and dynamic node execution
//   - Command support for combined state updates and routing
//
// Example usage:
//
//	schema := NewStateSchema().AddField("counter", StateField{...})
//	graph, err := NewStateGraph(schema).
//	  AddNode("increment", incrementFunc).
//	  SetEntryPoint("increment").
//	  SetFinishPoint("increment").
//	  Compile()
//
// The compiled Graph can then be executed with NewExecutor(graph).
type StateGraph struct {
	graph       *Graph
	buildErrors *stateGraphBuildErrors
}

type stateGraphBuildErrors struct {
	errs []error
}

// NewStateGraph creates a new graph builder with the given state schema.
func NewStateGraph(schema *StateSchema) *StateGraph { _ = "STUB: not implemented"; return nil }

func (sg *StateGraph) addBuildError(err error) { _ = "STUB: not implemented"; return }

func (sg *StateGraph) buildErr() error { _ = "STUB: not implemented"; return nil }

// Option is a function that configures a Node.
type Option func(*Node)

// WithName sets the name of the node.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDescription sets the description of the node.
func WithDescription(description string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNodeType sets the type of the node.
func WithNodeType(nodeType NodeType) Option { _ = "STUB: not implemented"; return *new(Option) }

func withTraceTransparent() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserInputKey sets the state key used as one-shot user input for LLM and
// Agent nodes. When empty, StateKeyUserInput is used.
func WithUserInputKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithToolSets sets the ToolSets for the node. This is a declarative
// per-node configuration used by AddLLMNode to build the LLM runner.
func WithToolSets(toolSets []tool.ToolSet) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRefreshToolSetsOnRun controls whether tools from ToolSets are
// refreshed from the underlying ToolSet on each node run.
// When false (default), tools from ToolSets are resolved once when the
// node is created. When true, the graph will call ToolSet.Tools again
// when building the tools map for each execution.
func WithRefreshToolSetsOnRun(refresh bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableParallelTools enables parallel tool execution for a Tools node.
// When enabled, if the last assistant message contains multiple tool calls,
// they will be executed concurrently and their responses will be merged in
// the original order. By default, tools run serially for compatibility.
func WithEnableParallelTools(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCacheKeyFields sets a cache key selector that derives the cache key
// input from a subset of fields in the sanitized input map. This helps avoid
// including unrelated or volatile keys in the cache key.
func WithCacheKeyFields(fields ...string) Option {
	_ = "STUB: not implemented"
	// copy fields to avoid external mutation
	return *new(Option)
}

// WithCacheKeySelector sets a custom selector for deriving the cache key input
// from the sanitized input map. The returned value will be passed to the
// CachePolicy.KeyFunc.
func WithCacheKeySelector(selector func(map[string]any) any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNodeCachePolicy sets a cache policy for this node.
// When set, the executor will attempt to cache the node's final result using this policy.
func WithNodeCachePolicy(policy *CachePolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRetryPolicy sets retry policies for the node. Policies are evaluated
// in order when an error occurs to determine whether to retry and what
// backoff to apply. Passing multiple policies allows matching by different
// conditions (e.g., network vs. HTTP status).
func WithRetryPolicy(policies ...RetryPolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithInterruptBefore pauses execution before this node runs.
func WithInterruptBefore() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInterruptAfter pauses execution after this node runs.
func WithInterruptAfter() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGenerationConfig sets the generation config for an LLM node.
// Effective only for nodes added via AddLLMNode.
func WithGenerationConfig(cfg model.GenerationConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamOutput enables node-to-node streaming for this node.
//
// For LLM and Agent nodes, streaming deltas are forwarded to the named stream.
// Function nodes can write to streams directly via OpenStreamWriter.
func WithStreamOutput(streamName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDestinations declares potential dynamic routing targets for a node.
// This is used for static validation (existence) and visualization only.
// It does not influence runtime execution.
func WithDestinations(dests map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEndsMap declares per-node named ends and their concrete destinations.
// The map keys are local symbolic names (e.g., "approved"), and values are
// concrete node IDs (or the special End) this node may route to.
// These ends are used at runtime to resolve Command.GoTo and conditional
// branch results, and at compile time for stronger validation.
func WithEndsMap(ends map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnds declares per-node named ends where the symbolic names are also the
// destination node IDs. Equivalent to WithEndsMap({name: name}).
func WithEnds(names ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPreNodeCallback sets a callback that will be executed before this specific node.
// This callback is specific to this node and will be executed in addition to any global callbacks.
func WithPreNodeCallback(callback BeforeNodeCallback) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPostNodeCallback sets a callback that will be executed after this specific node.
// This callback is specific to this node and will be executed in addition to any global callbacks.
func WithPostNodeCallback(callback AfterNodeCallback) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNodeErrorCallback sets a callback that will be executed when this specific node fails.
// This callback is specific to this node and will be executed in addition to any global callbacks.
func WithNodeErrorCallback(callback OnNodeErrorCallback) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNodeCallbacks sets multiple callbacks for this specific node.
// This allows setting multiple callbacks at once for convenience.
func WithNodeCallbacks(callbacks *NodeCallbacks) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Merge the provided callbacks with existing ones

// WithToolCallbacks sets tool callbacks for this specific node.
// This allows configuring tool callbacks at the node level.
// When both node-level and state-level callbacks are present, node-level
// callbacks take precedence.
// This option is only effective for tool nodes.
func WithToolCallbacks(callbacks *tool.Callbacks) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithToolCallRetryPolicy sets callable tool-call retry policy for tool nodes.
func WithToolCallRetryPolicy(policy *tool.RetryPolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAgentNodeEventCallback sets a callback that will be executed when an agent event is emitted.
// This callback is specific to this node and will be executed in addition to any global callbacks.
func WithAgentNodeEventCallback(callback AgentEventCallback) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Subgraph I/O mapping and scope utilities

// SubgraphResult captures a subgraph's outputs exposed to the parent mapper.
// RawStateDelta provides the original serialized final-state snapshot map from
// the subgraph's terminal graph.execution event. Callers can decode values
// with custom types if needed. Note that FinalState is reconstructed by JSON
// decoding, which may coerce numbers to float64 and complex structures to
// map[string]any.
//
// FallbackStateDelta and FallbackState are only populated when the child ends
// with a fatal error before emitting graph.execution. They carry the best-
// effort business state accumulated from the fatal path and are intentionally
// kept separate from FinalState/RawStateDelta so callers can distinguish a
// normal terminal snapshot from fatal fallback state.
type SubgraphResult struct {
	LastResponse       string
	FinalState         State
	RawStateDelta      map[string][]byte
	FallbackState      State
	FallbackStateDelta map[string][]byte
	StructuredOutput   any // Structured output from sub-agent.
}

// EffectiveState returns the normal final state when it exists, otherwise the
// fatal fallback state.
func (r SubgraphResult) EffectiveState() State { _ = "STUB: not implemented"; return *new(State) }

// EffectiveStateDelta returns the normal final-state delta when it exists,
// otherwise the fatal fallback delta.
func (r SubgraphResult) EffectiveStateDelta() map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// SubgraphInputMapper projects parent state into child runtime state.
// The returned state replaces the runtime state passed to the child.
type SubgraphInputMapper func(parent State) State

// SubgraphOutputMapper converts subgraph results into parent state updates.
// Returning nil or an empty State means "no updates" will be applied.
// Note: Prefer returning nil when there are no updates to write back;
// this reads clearer and is equivalent to applying an empty update.
type SubgraphOutputMapper func(parent State, result SubgraphResult) State

// WithSubgraphInputMapper sets a mapper used to build the child runtime state.
func WithSubgraphInputMapper(f SubgraphInputMapper) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSubgraphOutputMapper sets a mapper that writes subgraph outputs back to parent state.
func WithSubgraphOutputMapper(f SubgraphOutputMapper) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSubgraphIsolatedMessages toggles seeding of session messages to the child.
// When true, the child GraphAgent runs with include_contents=none.
// Docs note: This effectively sets CfgKeyIncludeContents="none" in the child
// runtime state so the child does not inject session history and only sees the
// projected input from the parent.
func WithSubgraphIsolatedMessages(isolate bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSubgraphInputFromLastResponse maps the parent's last_response to the
// child sub-agent's user_input for this agent node.
//
// Use this option when you want the downstream agent to consume only the
// upstream agent's result as its current-round input, without injecting the
// session history. This keeps agent nodes as black boxes while enabling
// explicit result passing.
//
// Note: For even stricter isolation from session history, combine with
// WithSubgraphIsolatedMessages(true).
func WithSubgraphInputFromLastResponse() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSubgraphEventScope customizes the child's event filter scope.
// Docs note: Scope may be hierarchical (can include '/').
// If empty, it defaults to the child agent name.
// The final filterKey becomes parent/scope (no UUID).
// This keeps the filterKey stable across turns.
func WithSubgraphEventScope(scope string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithModelCallbacks sets the model callbacks for LLM node.
func WithModelCallbacks(callbacks *model.Callbacks) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// tracingDisabled reports whether tracing is disabled for the invocation.
func tracingDisabled(invocation *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

// tracingDisabledInContext reports whether tracing is disabled for the invocation in context.
func tracingDisabledInContext(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// startNodeSpan returns a no-op span when tracing is disabled and otherwise starts a new span.
func startNodeSpan(ctx context.Context, spanName string) (context.Context, oteltrace.Span, bool) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(oteltrace.Span), false
}

// startNodeSpanForInvocation returns a no-op span when tracing is disabled for the invocation or context.
func startNodeSpanForInvocation(ctx context.Context, invocation *agent.Invocation, spanName string) (context.Context, oteltrace.Span, bool) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(oteltrace.Span), false
}

func workflowTypeFromNodeType(nodeType NodeType) itelemetry.WorkflowType {
	_ = "STUB: not implemented"
	return *new(itelemetry.WorkflowType)
}

func executeNodeWithWorkflowTrace(
	ctx context.Context,
	id string,
	nodeType NodeType,
	state State,
	function NodeFunc,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// AddNode adds a node with the given ID and function.
// The name and description of the node can be set with the options.
// This automatically sets up Pregel-style channel configuration.
func (sg *StateGraph) AddNode(id string, function NodeFunc, opts ...Option) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// Default to function type

// Automatically set up Pregel-style configuration.
// Create a trigger channel for this node.

// AddLLMNode adds a node that uses the model package directly.
func (sg *StateGraph) AddLLMNode(
	id string,
	llmModel model.Model,
	instruction string,
	tools map[string]tool.Tool,
	opts ...Option,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// Automatically set up Pregel-style configuration.
// Create a trigger channel for this node.

// AddToolsNode adds a node that uses the tools package directly.
func (sg *StateGraph) AddToolsNode(
	id string,
	tools map[string]tool.Tool,
	opts ...Option,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// AddAgentNode adds a node that uses a sub-agent by name.
// The agent name should correspond to a sub-agent in the GraphAgent's sub-agent list.
func (sg *StateGraph) AddAgentNode(
	id string,
	opts ...Option,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// Add agent node type option.

// AddSubgraphNode is a sugar alias of AddAgentNode to emphasize subgraph semantics.
func (sg *StateGraph) AddSubgraphNode(id string, opts ...Option) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// WithInterruptBeforeNodes enables static interrupts before the given nodes.
func (sg *StateGraph) WithInterruptBeforeNodes(
	nodeIDs ...string,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// WithInterruptAfterNodes enables static interrupts after the given nodes.
func (sg *StateGraph) WithInterruptAfterNodes(
	nodeIDs ...string,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

func (sg *StateGraph) setStaticInterruptNodes(
	nodeIDs []string,
	before bool,
) {
	_ = "STUB: not implemented"
	return
}

// channelUpdateMarker value for marking channel updates.
const channelUpdateMarker = "update"

const (
	joinChannelFromSeparator = ":from:"
	joinKeyLenBytes          = 8
)

// AddEdge adds a normal edge between two nodes.
// This automatically sets up Pregel-style channel configuration.
func (sg *StateGraph) AddEdge(from, to string) *StateGraph { _ = "STUB: not implemented"; return nil }

// Automatically set up Pregel-style channel for the edge.

// Set up trigger relationship (node subscribes) and trigger mapping.

// Add writer to source node.

// Non-nil sentinel to mark update.

// AddJoinEdge adds a join edge that waits for all start nodes to complete
// before triggering the end node.
func (sg *StateGraph) AddJoinEdge(fromNodes []string, to string) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

func joinChannelName(to string, starts []string) string { _ = "STUB: not implemented"; return "" }

func joinKeyForStarts(starts []string) string { _ = "STUB: not implemented"; return "" }

func normalizeJoinStarts(fromNodes []string) []string { _ = "STUB: not implemented"; return nil }

// AddConditionalEdges adds conditional routing from a node.
func (sg *StateGraph) AddConditionalEdges(
	from string,
	condFunc any,
	pathMap map[string]string,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// AddMultiConditionalEdges adds multi-conditional routing from a node.
// The condition returns multiple branch keys for parallel routing.
func (sg *StateGraph) AddMultiConditionalEdges(
	from string,
	condFunc MultiConditionalFunc,
	pathMap map[string]string,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// AddToolsConditionalEdges adds conditional routing from a LLM node to a tools node.
// If the last message has tool calls, route to the tools node.
// Otherwise, route to the fallback node.
func (sg *StateGraph) AddToolsConditionalEdges(
	fromLLMNode string,
	toToolsNode string,
	fallbackNode string,
) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// SetEntryPoint sets the entry point of the graph.
// This is equivalent to addEdge(Start, nodeId).
func (sg *StateGraph) SetEntryPoint(nodeID string) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// Also add an edge from Start to make it explicit

// SetFinishPoint adds an edge from the node to End.
// This is equivalent to addEdge(nodeId, End).
func (sg *StateGraph) SetFinishPoint(nodeID string) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// Compile compiles the graph and returns it for execution.
func (sg *StateGraph) Compile() (*Graph, error) { _ = "STUB: not implemented"; return nil, nil }

// WithNodeCallbacks adds node callbacks to the graph state schema.
// This allows users to register callbacks that will be executed during node execution.
func (sg *StateGraph) WithNodeCallbacks(callbacks *NodeCallbacks) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// WithCache sets the graph-level cache implementation.
func (sg *StateGraph) WithCache(cache Cache) *StateGraph { _ = "STUB: not implemented"; return nil }

// WithCachePolicy sets the default cache policy for all nodes (can be overridden per-node).
func (sg *StateGraph) WithCachePolicy(policy *CachePolicy) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// WithGraphVersion sets an optional version string used for cache namespacing.
// This helps avoid stale cache collisions across graph code changes or deployments.
func (sg *StateGraph) WithGraphVersion(version string) *StateGraph {
	_ = "STUB: not implemented"
	return nil
}

// ClearCache clears caches for the specified nodes. If nodes is empty, it clears all nodes currently in the graph.
func (sg *StateGraph) ClearCache(nodes ...string) *StateGraph {
	_ = "STUB: not implemented"
	return nil

	// collect all nodes
}

// MustCompile compiles the graph or panics if invalid.
func (sg *StateGraph) MustCompile() *Graph { _ = "STUB: not implemented"; return nil }

// LLMNodeFuncOption is a function that configures the LLM node function.
type LLMNodeFuncOption func(*llmRunner)

// WithLLMNodeID sets the node ID for the LLM node function.
func WithLLMNodeID(nodeID string) LLMNodeFuncOption {
	_ = "STUB: not implemented"
	return *new(LLMNodeFuncOption)
}

// WithLLMUserInputKey sets the one-shot input state key used by the LLM node.
// When empty, StateKeyUserInput is used.
func WithLLMUserInputKey(key string) LLMNodeFuncOption {
	_ = "STUB: not implemented"
	return *new(LLMNodeFuncOption)
}

// WithLLMRefreshToolSetsOnRun controls whether tools from ToolSets are
// refreshed from the underlying ToolSet on each LLM node run.
func WithLLMRefreshToolSetsOnRun(refresh bool) LLMNodeFuncOption {
	_ = "STUB: not implemented"
	return *new(LLMNodeFuncOption)
}

func mergeToolsWithToolSets(
	ctx context.Context,
	base map[string]tool.Tool,
	toolSets []tool.ToolSet,
) map[string]tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func cloneToolsMap(tools map[string]tool.Tool) map[string]tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// WithLLMToolSets sets the tool sets for the LLM node function.
func WithLLMToolSets(toolSets []tool.ToolSet) LLMNodeFuncOption {
	_ = "STUB: not implemented"
	return *new(LLMNodeFuncOption)
}

// WithLLMGenerationConfig sets the generation configuration for the LLM runner.
func WithLLMGenerationConfig(cfg model.GenerationConfig) LLMNodeFuncOption {
	_ = "STUB: not implemented"
	return *new(LLMNodeFuncOption)
}

// WithLLMStreamOutput sets the stream name used for node-to-node streaming.
func WithLLMStreamOutput(streamName string) LLMNodeFuncOption {
	_ = "STUB: not implemented"
	return *new(LLMNodeFuncOption)
}

// NewLLMNodeFunc creates a NodeFunc that uses the model package directly.
// This implements LLM node functionality using the model package interface.
func NewLLMNodeFunc(
	llmModel model.Model,
	instruction string,
	tools map[string]tool.Tool,
	opts ...LLMNodeFuncOption,
) NodeFunc {
	_ = "STUB: not implemented"
	return *new(NodeFunc)
}

// llmRunner encapsulates LLM execution dependencies to avoid long parameter
// lists.
type llmRunner struct {
	llmModel             model.Model
	instruction          string
	tools                map[string]tool.Tool
	toolSets             []tool.ToolSet
	refreshToolSetsOnRun bool
	nodeID               string
	generationConfig     model.GenerationConfig
	userInputKey         string
	streamOutputName     string
}

// execute implements the three-stage rule for LLM execution.
func (r *llmRunner) execute(ctx context.Context, state State, span oteltrace.Span) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (r *llmRunner) executeOneShotStage(
	ctx context.Context,
	state State,
	oneShotMsgs []model.Message,
	span oteltrace.Span,
	clearUpdate State,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Preallocate the common fast-path operations slice to avoid re-slicing
// growth on hot paths.

func (r *llmRunner) executeUserInputStage(
	ctx context.Context,
	state State,
	userInputKey string,
	userInput string,
	span oteltrace.Span,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Clear user input after execution.

func (r *llmRunner) executeHistoryStage(ctx context.Context, state State, span oteltrace.Span) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func applyInvocationRequestOverrides(
	request *model.Request,
	invocation *agent.Invocation,
	nodeID string,
) {
	_ = "STUB: not implemented"
	return
}

func extractModelResponseSummary(result any) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func selectGraphNodeModel(
	ctx context.Context,
	invocation *agent.Invocation,
	nodeID string,
	baseModel model.Model,
) (model.Model, *agent.Invocation, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil, nil
}

func graphModelInvocationView(
	invocation *agent.Invocation,
	nodeID string,
	callModel model.Model,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func runGraphModelSelector(
	ctx context.Context,
	selector agent.ModelSelector,
	invocation *agent.Invocation,
) (selected model.Model, err error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

type graphModelCall struct {
	ctx            context.Context
	tools          map[string]tool.Tool
	nodeID         string
	callInvocation *agent.Invocation
	callModel      model.Model
	span           oteltrace.Span
	startedSpan    bool
}

func (r *llmRunner) prepareModelCall(
	ctx context.Context,
	state State,
	span oteltrace.Span,
) (graphModelCall, error) {
	_ = "STUB: not implemented"
	return *new(graphModelCall), nil
}

func rootInvocationForGraphModelCall(ctx context.Context, state State) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func (r *llmRunner) executeModel(
	ctx context.Context,
	state State,
	messages []model.Message,
	span oteltrace.Span,
	instructionUsed string,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Sanitize invalid tool calls in history to avoid poisoning future requests.

// Build model input metadata from the original state and instruction
// so events accurately reflect both instruction and user input.

// processInstruction resolves placeholder variables in the instruction.
// It supports the same syntax as LLMAgent, including {invocation:*} values
// stored on the current invocation.
func (r *llmRunner) processInstruction(state State) string { _ = "STUB: not implemented"; return "" }

// extractAssistantMessage extracts the assistant message from model result.
func extractAssistantMessage(result any) *model.Message { _ = "STUB: not implemented"; return nil }

// extractResponseID extracts response ID from model result.
func extractResponseID(result any) string { _ = "STUB: not implemented"; return "" }

// ensureSystemHead ensures system prompt is at the head if provided.
func ensureSystemHead(in []model.Message, sys string) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// extractExecutionContext extracts execution context from state.
func extractExecutionContext(state State) (invocationID, sessionID, appName, userID string, eventChan chan<- *event.Event) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

func executionContextFromState(state State) *ExecutionContext {
	_ = "STUB: not implemented"
	return nil
}

// modelResponseConfig contains configuration for processing model responses.
type modelResponseConfig struct {
	Response         *model.Response
	Invocation       *agent.Invocation
	StableInvocation *agent.Invocation
	Tracker          *itelemetry.ChatMetricsTracker
	PartialUsage     *responseusage.PartialState
	ModelCallbacks   *model.Callbacks
	EventChan        chan<- *event.Event
	InvocationID     string
	SessionID        string
	LLMModel         model.Model
	Request          *model.Request
	Span             oteltrace.Span
	// NodeID, when provided, is used as the event author.
	NodeID string
}

func responseModelError(rsp *model.Response) error { _ = "STUB: not implemented"; return nil }

func invocationFromContextOrDefault(
	ctx context.Context,
	invocation *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func shouldDisableModelExecutionEvents(invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

func applyBeforeModelPluginCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	request *model.Request,
	span oteltrace.Span,
) (context.Context, bool, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), false, nil, nil
}

func applyBeforeModelCallbacks(
	ctx context.Context,
	callbacks *model.Callbacks,
	request *model.Request,
	span oteltrace.Span,
) (context.Context, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func singleResponseStream(response *model.Response) modelResponseStream {
	_ = "STUB: not implemented"
	return *new(modelResponseStream)
}

func applyAfterModelPluginCallbacks(
	ctx context.Context,
	args *model.AfterModelArgs,
	span oteltrace.Span,
) (context.Context, bool, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), false, nil, nil
}

func applyAfterModelCallbacks(
	ctx context.Context,
	callbacks *model.Callbacks,
	args *model.AfterModelArgs,
	span oteltrace.Span,
) (context.Context, *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func modelResponseAuthor(config modelResponseConfig) string { _ = "STUB: not implemented"; return "" }

func applyPartialEventMetadataOverrides(
	ev *event.Event,
	resp *model.Response,
	invocation *agent.Invocation,
) {
	_ = "STUB: not implemented"
	return
}

func emitModelResponseEvent(
	ctx context.Context,
	config modelResponseConfig,
	eventInvocation *agent.Invocation,
	optionsInvocation *agent.Invocation,
	ev *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldEmitModelResponseEvent(
	rsp *model.Response,
	invocation *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Emit the final (Done) response when it carries ReasoningContent so
// the runner can persist it into the session. Without this, thinking
// content from graph LLM nodes is lost across conversation turns.

// processModelResponse processes a single model response.
func processModelResponse(ctx context.Context, config modelResponseConfig) (context.Context, *event.Event, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// responseHasReasoningContent reports whether any choice in the response
// carries reasoning/thinking content.
func responseHasReasoningContent(rsp *model.Response) bool { _ = "STUB: not implemented"; return false }

func shouldEmitModelResponse(rsp *model.Response) bool { _ = "STUB: not implemented"; return false }

type modelResponseStream struct {
	// Ch is the response channel returned by model.Model.GenerateContent.
	Ch <-chan *model.Response
	// Seq is the iterator returned by model.IterModel.GenerateContentIter.
	Seq model.Seq[*model.Response]
}

func generateModelStream(
	ctx context.Context,
	llmModel model.Model,
	request *model.Request,
	span oteltrace.Span,
) (modelResponseStream, error) {
	_ = "STUB: not implemented"
	// Generate content.
	return *new(modelResponseStream), nil
}

func runModelStream(
	ctx context.Context,
	invocation *agent.Invocation,
	modelCallbacks *model.Callbacks,
	llmModel model.Model,
	request *model.Request,
	beforeGenerate func(context.Context),
) (context.Context, modelResponseStream, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(modelResponseStream), nil
}

// Set span attributes for model execution.

// runModel preserves the pre-refactor test-facing helper signature by
// adapting iterator-based model streams back to the legacy channel form.
func runModel(
	ctx context.Context,
	modelCallbacks *model.Callbacks,
	llmModel model.Model,
	request *model.Request,
) (context.Context, <-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// NewToolsNodeFunc creates a NodeFunc that uses the tools package directly.
// This implements tools node functionality using the tools package interface.
func NewToolsNodeFunc(tools map[string]tool.Tool, opts ...Option) NodeFunc {
	_ = "STUB: not implemented"
	return *new(NodeFunc)
}

func newToolsNodeRuntime(
	tools map[string]tool.Tool,
	opts ...Option,
) (NodeFunc, map[string]tool.Tool) {
	_ = "STUB: not implemented"
	return *new(NodeFunc), nil
}

// Capture whether to execute tools in parallel.

// Capture tool callbacks configured on the node.

// Extract and validate messages from state.

// Extract execution context for event emission.

// Determine which callbacks to use: node-configured takes precedence over state.

// Process all tool calls and collect results.

func resolveToolsNodeRuntimeTools(
	ctx context.Context,
	state State,
	node *Node,
	baseTools map[string]tool.Tool,
	staticTools map[string]tool.Tool,
) map[string]tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// copyRuntimeStateFiltered creates a shallow copy of the parent state excluding
// internal/ephemeral keys that should not leak into a child sub-agent's
// Invocation.RunOptions.RuntimeState (e.g., exec context, callbacks, session).
//
// Important: This is a shallow copy (only key bindings are copied); complex
// values (map/slice) remain shared references. Avoid concurrent mutation of the
// same complex object from parent/child. If isolation is required, deep copy in
// SubgraphInputMapper.
func copyRuntimeStateFiltered(parent State) State { _ = "STUB: not implemented"; return *new(State) }

type executorProvider interface {
	Executor() *Executor
}

type subgraphInterruptInfo struct {
	parentNodeID      string
	childAgentName    string
	childCheckpointID string
	childCheckpointNS string
	childLineageID    string
	childTaskID       string
}

type extractedPregelInterrupt struct {
	interrupt    *InterruptError
	lineageID    string
	checkpointID string
	checkpointNS string
}

func subgraphInterruptInfoFromState(
	state State,
) (subgraphInterruptInfo, bool) {
	_ = "STUB: not implemented"
	return *new(subgraphInterruptInfo), false
}

func extractPregelInterruptInfo(e *event.Event) (*extractedPregelInterrupt, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func extractPregelInterrupt(e *event.Event) (*InterruptError, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func latestInterruptedCheckpointID(
	ctx context.Context,
	targetAgent agent.Agent,
	lineageID string,
	namespace string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resumeCommandForSubgraph(
	state State,
	childTaskID string,
) *Command {
	_ = "STUB: not implemented"
	return nil
}

const includeContentsNone = "none"

type agentNodeConfig struct {
	callbacks           *NodeCallbacks
	inputMapper         SubgraphInputMapper
	outputMapper        SubgraphOutputMapper
	isolated            bool
	scope               string
	inputFromLast       bool
	llmGenerationConfig *model.GenerationConfig
	userInputKey        string
	streamOutputName    string
}

func agentNodeConfigFromOptions(opts ...Option) agentNodeConfig {
	_ = "STUB: not implemented"
	return *new(agentNodeConfig)
}

func targetAgentFromState(state State, agentName string) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func initialChildStateForAgentNode(parent State, inputMapper SubgraphInputMapper) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func applyDefaultCheckpointNamespace(child State, targetAgent agent.Agent) {
	_ = "STUB: not implemented"
	return
}

func applyIsolatedMessages(child State, isolated bool) { _ = "STUB: not implemented"; return }

func applyCheckpointResumeFields(child State, info subgraphInterruptInfo) {
	_ = "STUB: not implemented"
	return
}

func clearResumeChannelsIfNeeded(parent State, child State, cmd *Command) {
	_ = "STUB: not implemented"
	return
}

func consumeResumeMapEntryIfNeeded(parent State, childTaskID string, cmd *Command) {
	_ = "STUB: not implemented"
	return
}

func applyResumeCommandForAgentNode(
	parent State,
	child State,
	info subgraphInterruptInfo,
) {
	_ = "STUB: not implemented"
	return
}

func applySubgraphResumeForAgentNode(parent State, child State, nodeID string) {
	_ = "STUB: not implemented"
	return
}

func buildChildStateForAgentNode(
	parent State,
	nodeID string,
	targetAgent agent.Agent,
	cfg agentNodeConfig,
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func mapParentInputFromLastResponse(
	state State,
	enabled bool,
	userInputKey string,
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func setSubgraphInterruptState(
	ctx context.Context,
	state State,
	nodeID string,
	agentName string,
	targetAgent agent.Agent,
	childState State,
	invocation *agent.Invocation,
	childTaskID string,
	interruptInfo *extractedPregelInterrupt,
) {
	_ = "STUB: not implemented"
	return
}

// Unify local/remote handling: rely on interrupt event metadata rather than
// local executor/checkpoint-manager lookups.

// Fallback: when metadata does not carry checkpoint_id, best-effort lookup
// from local checkpoint manager for local GraphAgent subgraphs.

func finalizeAgentNodeOutput(
	state State,
	nodeID string,
	streamRes agentEventStreamResult,
	outputMapper SubgraphOutputMapper,
	userInputKey string,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// NewAgentNodeFunc creates a NodeFunc that looks up and uses a sub-agent by name.
// The agent name should correspond to a sub-agent in the parent GraphAgent's sub-agent list.
func NewAgentNodeFunc(agentName string, opts ...Option) NodeFunc {
	_ = "STUB: not implemented"
	return *new(NodeFunc)
}

type agentNodeRuntime struct {
	agentName string
	config    agentNodeConfig
}

func newAgentNodeRuntime(agentName string, cfg agentNodeConfig) *agentNodeRuntime {
	_ = "STUB: not implemented"
	return nil
}

func (r *agentNodeRuntime) NodeFunc() NodeFunc { _ = "STUB: not implemented"; return *new(NodeFunc) }

func (r *agentNodeRuntime) Run(
	ctx context.Context,
	state State,
	traceTask *traceTaskMetadata,
) (any, error) {
	_ = "STUB: not implemented"
	// Extract execution context for event emission.
	return *new(any), nil
}

// Extract current node ID from state.

// Optionally map parent's last_response to user_input for this agent node.

// Build invocation for the target agent with custom runtime state and scope.

// Emit agent execution start event.

// Execute the target agent through RunWithPlugins so that Runner-scoped
// PluginManager AgentCallbacks (BeforeAgent/AfterAgent) consistently
// apply to sub-agents invoked via agent-nodes, matching chain/parallel/
// cycle/transfer behavior. See issue #1432.
//
// Important: wrap the context with the sub-invocation so downstream
// callbacks (model/tool) can access it via agent.InvocationFromContext(ctx).

// Emit agent execution error event.

// Process agent event stream and capture completion state.

// Emit agent execution complete event.

func recordAgentNodeTraceTerminals(
	metadata *traceTaskMetadata,
	invocation *agent.Invocation,
) {
	_ = "STUB: not implemented"
	return
}

func agentNodeChildTerminalStepIDs(
	invocation *agent.Invocation,
	entryPredecessors []string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func stateStringOr(state State, key string, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

type agentEventStreamResult struct {
	lastResponse     string
	lastResponseID   string
	finalState       State
	rawDelta         map[string][]byte
	fallbackState    State
	fallbackRawDelta map[string][]byte
	structuredOutput any
	interrupt        *InterruptError
	interruptInfo    *extractedPregelInterrupt
	terminalErr      error
	terminalErrMeta  agentTerminalErrorMeta
	finalError       *model.ResponseError
}

type agentTerminalErrorMeta struct {
	InvocationID string
	FilterKey    string
}

type agentDeltaStreamTap struct {
	writer       *agent.StreamWriter
	lastResponse *string

	sawDelta bool
	broken   bool
}

func newAgentDeltaStreamTap(
	ctx context.Context,
	streamName string,
	lastResponse *string,
) (*agentDeltaStreamTap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *agentDeltaStreamTap) WriteDelta(ev *event.Event) { _ = "STUB: not implemented"; return }

func (t *agentDeltaStreamTap) Close(errp *error) { _ = "STUB: not implemented"; return }

func agentDeltaFromEvent(ev *event.Event) string { _ = "STUB: not implemented"; return "" }

func runAgentEventCallbacks(
	ctx context.Context,
	nodeCallbacks *NodeCallbacks,
	nodeID string,
	agentName string,
	state State,
	ev *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func updateAgentStreamResultFromEvent(
	ctx context.Context,
	res *agentEventStreamResult,
	ev *event.Event,
	invalidateSuccessResult bool,
	trackTerminalErrors bool,
) {
	_ = "STUB: not implemented"
	return
}

func updateAgentLastResponse(res *agentEventStreamResult, ev *event.Event) {
	_ = "STUB: not implemented"
	return
}

func updateAgentLastResponseValue(lastResponse *string, lastResponseID *string, ev *event.Event) {
	_ = "STUB: not implemented"
	return
}

func updateAgentStructuredOutput(res *agentEventStreamResult, ev *event.Event) {
	_ = "STUB: not implemented"
	return
}

func updateAgentInterrupt(res *agentEventStreamResult, ev *event.Event) {
	_ = "STUB: not implemented"
	return
}

// Keep the latest interrupt metadata observed in the stream.
// In nested subgraphs, a deeper child interrupt may appear first, followed by
// the immediate child interrupt propagated upward. Using the latest metadata
// avoids pinning parent resume state to a deeper descendant checkpoint.

func clearAgentSuccessResultOnError(res *agentEventStreamResult, ev *event.Event) {
	_ = "STUB: not implemented"
	return
}

func shouldInvalidateAgentSuccessResult(ev *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func clearAgentTerminalErrorOnContinuedOutput(
	res *agentEventStreamResult,
	ev *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func isAgentRecoveryEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

func matchesAgentTerminalErrorSource(
	meta agentTerminalErrorMeta,
	ev *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

func internalAgentEventWithInvocationFields(
	invocation *agent.Invocation,
	ev *event.Event,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func updateAgentTerminalError(res *agentEventStreamResult, ev *event.Event) {
	_ = "STUB: not implemented"
	return
}

func isTerminalAgentErrorEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

func isTerminalAgentSuccessEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

func updateAgentFinalState(
	ctx context.Context,
	res *agentEventStreamResult,
	ev *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func extractSubgraphFinalState(
	ctx context.Context,
	ev *event.Event,
) (State, map[string][]byte, bool) {
	_ = "STUB: not implemented"
	return *new(State), nil, false
}

func decodeSubgraphStateDelta(
	ctx context.Context,
	stateDelta map[string][]byte,
) (State, map[string][]byte, bool) {
	_ = "STUB: not implemented"
	return *new(State), nil, false
}

// Some transports normalize JSON string scalars into plain strings,
// which means the raw bytes are no longer valid JSON here. Preserve
// the raw string instead of silently dropping the key.

// processAgentEventStream processes the event stream from the target agent.
// This function handles forwarding events and capturing completion state.
func processAgentEventStream(
	ctx context.Context,
	invocation *agent.Invocation,
	agentEventChan <-chan *event.Event,
	nodeCallbacks *NodeCallbacks,
	nodeID string,
	state State,
	eventChan chan<- *event.Event,
	agentName string,
	streamName string,
) (res agentEventStreamResult, err error) {
	_ = "STUB: not implemented"
	return *new(agentEventStreamResult), nil
}

// Forward the event to the parent event channel.

func mergeAgentEventCallbacks(
	state State,
	perNode *NodeCallbacks,
) *NodeCallbacks {
	_ = "STUB: not implemented"
	return nil
}

func captureAgentFallbackState(
	res *agentEventStreamResult,
	ev *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func mergeAgentFallbackStateDelta(
	dst map[string][]byte,
	src map[string][]byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func shouldPropagateAgentFallbackStateKey(
	key string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldPropagateAgentFallbackState(
	err *model.ResponseError,
) bool {
	_ = "STUB: not implemented"
	return false
}

// buildAgentInvocationWithStateScopeAndInputKey builds an invocation for the
// target agent
// using a custom runtime state and an optional event filter scope segment.
//
// FilterKey Strategy:
// The FilterKey is built using a stable, deterministic pattern based on the agent
// hierarchy rather than random UUIDs. This ensures that:
//  1. Multi-turn conversations work correctly with BranchFilterModePrefix
//  2. Child agent responses from previous requests are visible in subsequent requests
//  3. The prefix matching in event.Filter() works as expected across requests
//
// Format: parentKey/agentName (without UUID)
// Example: "input/knowledge" instead of "input/knowledge/random-uuid"
func buildAgentInvocationWithStateScopeAndInputKey(
	ctx context.Context,
	parentState State,
	runtime State,
	targetAgent agent.Agent,
	nodeID string,
	scope string,
	userInputKey string,
	traceTask *traceTaskMetadata,
) *agent.Invocation {
	_ = "STUB: not implemented"
	// Extract user input from parent state.
	return nil
}

// Extract session from parent state.

// Clone from parent invocation if available to preserve linkage and filtering.

// Preserve the parent's visibility preference.
// The agent node captures completion snapshots from either raw
// graph.execution events or visible rewritten completion snapshots.

// Build a stable FilterKey without UUID to ensure multi-turn conversations

// Create standalone invocation.

// Use stable FilterKey based on agent name only (no UUID).

func currentTraceTaskPredecessors(traceTask *traceTaskMetadata) []string {
	_ = "STUB: not implemented"
	return nil
}

func currentTraceStepPredecessors(state State) []string { _ = "STUB: not implemented"; return nil }

func buildAgentInvocationWithStateAndScope(
	ctx context.Context,
	parentState State,
	runtime State,
	targetAgent agent.Agent,
	nodeID string,
	scope string,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func buildAgentNodeTraceNodeID(
	parentInvocation *agent.Invocation,
	nodeID string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func buildAgentNodeSurfaceRoot(
	parentInvocation *agent.Invocation,
	nodeID string,
) string {
	_ = "STUB: not implemented"
	return ""
}

const (
	errCallbackBeforeTool = "callback before tool error: %w"
	errCallbackAfterTool  = "callback after tool error: %w"
)

func runBeforeToolPluginCallbacks(
	ctx context.Context,
	toolCall model.ToolCall,
	decl *tool.Declaration,
	state State,
) (context.Context, model.ToolCall, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(model.ToolCall), *new(any), nil
}

func runBeforeToolCallbacks(
	ctx context.Context,
	toolCall model.ToolCall,
	decl *tool.Declaration,
	toolCallbacks *tool.Callbacks,
	state State,
) (context.Context, model.ToolCall, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(model.ToolCall), *new(any), nil
}

func ensureCallableTool(
	t tool.Tool,
	toolName string,
) (tool.CallableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool), nil
}

func runAfterToolPluginCallbacks(
	ctx context.Context,
	toolCall model.ToolCall,
	decl *tool.Declaration,
	result any,
	runErr error,
) (context.Context, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any), nil
}

func runAfterToolCallbacks(
	ctx context.Context,
	toolCall model.ToolCall,
	decl *tool.Declaration,
	result any,
	runErr error,
	toolCallbacks *tool.Callbacks,
) (context.Context, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any), nil
}

// extractMetaFromResult extracts metadata from tool result when available.
func extractMetaFromResult(result any) map[string]any { _ = "STUB: not implemented"; return nil }

// runTool executes a tool with before/after callbacks and returns the result.
// Parameters:
//   - ctx: context for cancellation and tracing
//   - toolCall: the tool call to execute, including function name and arguments
//   - toolCallbacks: callbacks to execute before and after tool execution
//   - t: the tool implementation to execute
//
// Returns:
//   - context.Context: the updated context from callbacks (if any)
//   - any: the result from tool execution or custom callback result
//   - []byte: the modified arguments after before-tool callbacks (for telemetry)
//   - error: any error that occurred during execution
func runTool(
	ctx context.Context,
	toolCall model.ToolCall,
	toolCallbacks *tool.Callbacks,
	t tool.Tool,
	state State,
) (context.Context, any, []byte, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(any), nil, nil
}

func runToolWithEventContexts(
	ctx context.Context,
	toolCall model.ToolCall,
	toolCallbacks *tool.Callbacks,
	t tool.Tool,
	state State,
	retryPolicy *tool.RetryPolicy,
) (context.Context, *agent.Invocation, context.Context, *agent.Invocation, any, []byte, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, *new(context.Context), nil, *new(any), nil, nil
}

func extractResultError(result any) bool { _ = "STUB: not implemented"; return false }

// extractModelInput extracts the model input from state and instruction.
func extractModelInput(state State, instruction, userInputKey string) string {
	_ = "STUB: not implemented"

	// Get user input if available.
	return ""
}

// Add instruction if provided.

// getModelName extracts the model name from the model instance.
func getModelName(llmModel model.Model) string { _ = "STUB: not implemented"; return "" }

// emitModelStartEvent emits a model execution start event.
func emitModelStartEvent(
	ctx context.Context,
	baseInvocation *agent.Invocation,
	currentInvocation *agent.Invocation,
	eventChan chan<- *event.Event,
	invocationID, modelName, nodeID, modelInput string,
	startTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// emitModelCompleteEvent emits a model execution complete event.
func emitModelCompleteEvent(
	ctx context.Context,
	baseInvocation *agent.Invocation,
	currentInvocation *agent.Invocation,
	eventChan chan<- *event.Event,
	invocationID, modelName, nodeID, modelInput, modelOutput, responseID string,
	startTime, endTime time.Time,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func emitInvocationScopedEvent(
	ctx context.Context,
	baseInvocation *agent.Invocation,
	currentInvocation *agent.Invocation,
	eventChan chan<- *event.Event,
	invocationID string,
	ev *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func modelExecutionEventRequestID(baseInvocation, currentInvocation *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func modelExecutionEventParentInvocationID(baseInvocation, currentInvocation *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func modelExecutionEventBranch(baseInvocation, currentInvocation *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func modelExecutionEventFilterKey(baseInvocation, currentInvocation *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

// modelExecutionConfig contains configuration for model execution with events.
type modelExecutionConfig struct {
	Invocation     *agent.Invocation
	ModelCallbacks *model.Callbacks
	LLMModel       model.Model
	Request        *model.Request
	EventChan      chan<- *event.Event
	InvocationID   string
	SessionID      string
	AppName        string
	UserID         string
	NodeID         string // Add NodeID for parallel execution support
	NodeResultKey  string // Add NodeResultKey for configurable result key pattern
	DeltaStream    *agent.StreamWriter
	BeforeGenerate func(context.Context)
	Span           oteltrace.Span
}

const (
	errMsgNoModelResponse = "no response received from model"
	errMsgNoModelChoices  = "model returned no choices"
)

type modelDeltaStreamTap struct {
	writer *agent.StreamWriter

	sawDelta bool
	broken   bool
}

func newModelDeltaStreamTap(w *agent.StreamWriter) *modelDeltaStreamTap {
	_ = "STUB: not implemented"
	return nil
}

func (t *modelDeltaStreamTap) WriteDelta(resp *model.Response) { _ = "STUB: not implemented"; return }

func (t *modelDeltaStreamTap) WriteFinalIfNoDelta(final *model.Response) {
	_ = "STUB: not implemented"
	return
}

func modelDeltaFromResponse(resp *model.Response) string { _ = "STUB: not implemented"; return "" }

func modelMessageFromResponse(resp *model.Response) string { _ = "STUB: not implemented"; return "" }

func validateFinalModelResponse(
	span oteltrace.Span,
	resp *model.Response,
) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectToolCallsFromResponse(
	toolCalls []model.ToolCall,
	resp *model.Response,
) []model.ToolCall {
	_ = "STUB: not implemented"
	return nil
}

func mergeToolCallsIntoFinalResponse(
	resp *model.Response,
	toolCalls []model.ToolCall,
) {
	_ = "STUB: not implemented"
	return
}

func hasAfterModelCallbacks(
	invocation *agent.Invocation,
	modelCallbacks *model.Callbacks,
) bool {
	_ = "STUB: not implemented"
	return false
}

func hasRegisteredAfterModelCallbacks(callbacks *model.Callbacks) bool {
	_ = "STUB: not implemented"
	return false
}

func nextReusableModelEvent(
	reusableEvents []event.Event,
	reusableEventIdx *int,
) *event.Event {
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

func emitFastModelResponseEvent(
	ctx context.Context,
	eventInvocation *agent.Invocation,
	config modelExecutionConfig,
	response *model.Response,
	author string,
	partialEventIDsDisabled bool,
	partialEventTimestampsDisabled bool,
	reusableEvent *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func traceProcessedModelResponse(
	span oteltrace.Span,
	tracker *itelemetry.ChatMetricsTracker,
	invocation *agent.Invocation,
	request *model.Request,
	response *model.Response,
	lastEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

type modelResponseProcessor struct {
	ctx                            context.Context
	config                         modelExecutionConfig
	stableInvocation               *agent.Invocation
	observabilityInvocation        *agent.Invocation
	invocation                     *agent.Invocation
	tracker                        *itelemetry.ChatMetricsTracker
	timingInfo                     *model.TimingInfo
	partialUsageState              responseusage.PartialState
	tap                            *modelDeltaStreamTap
	reusableEvents                 []event.Event
	reusableEventIdx               int
	author                         string
	fastResponsePath               bool
	partialEventIDsDisabled        bool
	partialEventTimestampsDisabled bool
	lastEvent                      *event.Event
	finalResponse                  *model.Response
	toolCalls                      []model.ToolCall
}

func newModelResponseProcessor(
	ctx context.Context,
	config modelExecutionConfig,
	invocation *agent.Invocation,
	runErr *error,
) *modelResponseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func observabilityInvocationView(
	invocation *agent.Invocation,
	config modelExecutionConfig,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func hasObservabilityModelValue(
	invocation *agent.Invocation,
	config modelExecutionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

func hasObservabilityInvocationIDValue(
	invocation *agent.Invocation,
	config modelExecutionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

func hasObservabilitySessionIDValue(
	invocation *agent.Invocation,
	config modelExecutionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

func hasObservabilityUserIDValue(
	invocation *agent.Invocation,
	config modelExecutionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

func hasObservabilityAppNameValue(
	invocation *agent.Invocation,
	config modelExecutionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldRefreshObservabilitySessionView(
	currentSession *session.Session,
	invocation *agent.Invocation,
	config modelExecutionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldRefreshObservabilityInvocationView(
	currentView *agent.Invocation,
	invocation *agent.Invocation,
	config modelExecutionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

func refreshObservabilityInvocationView(
	currentView *agent.Invocation,
	invocation *agent.Invocation,
	config modelExecutionConfig,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func (p *modelResponseProcessor) close() { _ = "STUB: not implemented"; return }

func (p *modelResponseProcessor) consume(stream modelResponseStream) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *modelResponseProcessor) handleResponse(response *model.Response) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *modelResponseProcessor) finalize() (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// executeModelAndProcessResponses runs the model call and handles response-side
// behavior such as streaming, callbacks, response events, telemetry, and final
// response assembly. DisableModelExecutionEvents only controls the model
// lifecycle events emitted around this call site; it does not bypass model
// execution or response processing.
func executeModelAndProcessResponsesWithContext(
	ctx context.Context,
	config modelExecutionConfig,
) (context.Context, *agent.Invocation, any, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, *new(any), nil
}

func executeModelAndProcessResponses(
	ctx context.Context,
	config modelExecutionConfig,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// executeModelWithEvents preserves the previous helper name for existing
// tests while delegating to the refactored response-processing pipeline.
func executeModelWithEvents(ctx context.Context, config modelExecutionConfig) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// extractToolCallsFromState extracts and validates tool calls from the state.
// It scans backwards from the end to find the most recent assistant message with tool calls,
// stopping when it encounters a user message.
func extractToolCallsFromState(state State, span oteltrace.Span) ([]model.ToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Scan backwards to find the most recent assistant message with tool calls.
// Stop when encountering a user message to ensure proper tool call pairing.

// Stop scanning when we encounter a user message.
// This ensures we don't process tool calls from previous conversation turns.

// Skip system, tool, and other message types.

// toolCallsConfig contains configuration for processing tool calls.
type toolCallsConfig struct {
	ToolCalls    []model.ToolCall
	Tools        map[string]tool.Tool
	InvocationID string
	EventChan    chan<- *event.Event
	Span         oteltrace.Span
	State        State
	// EnableParallel controls whether multiple tool calls are executed concurrently.
	// When false or when there is only one tool call, execution is serial.
	EnableParallel bool
	// ToolCallbacks specifies tool callbacks to use.
	// If nil, callbacks will be extracted from State.
	ToolCallbacks *tool.Callbacks
	// RetryPolicy specifies callable tool-call retry policy for this tools node.
	RetryPolicy *tool.RetryPolicy
}

// processToolCalls executes all tool calls and returns the resulting messages.
func processToolCalls(ctx context.Context, config toolCallsConfig) ([]model.Message, error) {
	_ = "STUB: not implemented"
	// Use callbacks from config if provided; otherwise extract from state.
	return nil, nil
}

// Serial path or single tool call.

// Parallel path: execute each tool call in its own goroutine while
// preserving the original order in the resulting messages slice.

// On error, cancel siblings but still report result so collector can exit cleanly.

// Aggregate while preserving order.

// Only set when message exists; zero value is fine otherwise.

// singleToolCallConfig contains configuration for executing a single tool call.
type singleToolCallConfig struct {
	ToolCall      model.ToolCall
	Tools         map[string]tool.Tool
	InvocationID  string
	EventChan     chan<- *event.Event
	Span          oteltrace.Span
	ToolCallbacks *tool.Callbacks
	State         State
	RetryPolicy   *tool.RetryPolicy
}

// executeSingleToolCall executes a single tool call with event emission.
func executeSingleToolCall(ctx context.Context, config singleToolCallConfig) (model.Message, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), nil
}

// Extract current node ID from state for event authoring.

// Keep the original invocation as a fallback when callbacks return a bare context.

// Emit tool execution start event with modified arguments.

// Do not emit error payload for interrupt so clients treat it as pause.

// Set result to interrupt value when no result is provided.

// Emit tool execution complete event.

// Marshal result to JSON.

func invocationFromContextOrFallback(ctx context.Context, fallback *agent.Invocation) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func invocationOrFallback(invocation *agent.Invocation, fallback *agent.Invocation) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func invocationIDOrFallback(invocation *agent.Invocation, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

// emitToolStartEvent emits a tool execution start event.
func emitToolStartEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
	invocationID, toolName, toolID, nodeID string,
	startTime time.Time,
	arguments []byte,
	responseID string,
) {
	_ = "STUB: not implemented"
	return
}

// toolCompleteEventConfig contains configuration for tool complete events.
type toolCompleteEventConfig struct {
	EventChan    chan<- *event.Event
	InvocationID string
	ToolName     string
	ToolID       string
	NodeID       string
	ResponseID   string
	StartTime    time.Time
	Result       any
	Error        error
	Arguments    []byte
}

// emitToolCompleteEvent emits a tool execution complete event.
func emitToolCompleteEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	config toolCompleteEventConfig,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// extractToolCallbacks extracts tool callbacks from the state.
func extractToolCallbacks(state State) (*tool.Callbacks, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// MessagesStateSchema creates a state schema optimized for message-based workflows.
func MessagesStateSchema() *StateSchema { _ = "STUB: not implemented"; return nil }

// buildAgentInvocation builds an invocation for the target agent.
func buildAgentInvocation(ctx context.Context, state State, targetAgent agent.Agent) *agent.Invocation {
	_ = "STUB: not implemented"
	// Delegate to the unified builder with default runtime state and empty scope.
	return nil
}

// emitAgentStartEvent emits an agent execution start event.
func emitAgentStartEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocationID, nodeID string,
	startTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// emitAgentCompleteEvent emits an agent execution complete event.
func emitAgentCompleteEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocationID, nodeID string,
	startTime, endTime time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// emitAgentErrorEvent emits an agent execution error event.
func emitAgentErrorEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocationID, nodeID string,
	startTime, endTime time.Time,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

// findSubAgentByName looks up a sub-agent by name from the parent agent.
func findSubAgentByName(parentAgent any, agentName string) agent.Agent {
	_ = "STUB: not implemented"
	// Try to cast to an interface that has SubAgents method.
	return *new(agent.Agent)
}
