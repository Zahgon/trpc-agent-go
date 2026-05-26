//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package graph provides graph-based execution functionality.
package graph

import (
	"context"
	"sync"
	"sync/atomic"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph/internal/channel"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Special node identifiers for graph routing.
const (
	// Start represents the virtual start node for routing.
	Start = "__start__"
	// End represents the virtual end node for routing.
	End = "__end__"
)

// Error types for graph execution.
const (
	ErrorTypeGraphExecution  = "graph_execution_error"
	ErrorTypeInvalidNode     = "invalid_node_error"
	ErrorTypeInvalidState    = "invalid_state_error"
	ErrorTypeInvalidEdge     = "invalid_edge_error"
	ErrorTypeConditionalEdge = "conditional_edge_error"
	ErrorTypeStateValidation = "state_validation_error"
	ErrorTypeNodeExecution   = "node_execution_error"
	ErrorTypeCircularRef     = "circular_reference_error"
	ErrorTypeConcurrency     = "concurrency_error"
	ErrorTypeTimeout         = "timeout_error"
	ErrorTypeModelGeneration = "model_generation_error"
)

// NodeFunc is a function that can be executed by a node.
// Node function signature: (state) -> updated_state or Command.
type NodeFunc func(ctx context.Context, state State) (any, error)

// NodeResult represents the result of executing a node function.
// It can be either a State update or a Command for combined state update + routing.
type NodeResult any

// ConditionalFunc is a function that determines the next node(s) based on state.
// Conditional edge function signature.
type ConditionalFunc = func(ctx context.Context, state State) (string, error)

// MultiConditionalFunc returns multiple next nodes for parallel execution.
type MultiConditionalFunc = func(ctx context.Context, state State) ([]string, error)

// UniversalCondFunc is a function that determines the next node(s) based on state.
type UniversalCondFunc = func(ctx context.Context, state State) (ConditionResult, error)

// ConditionResult represents the result of executing a conditional edge function.
type ConditionResult struct {
	NextNodes []string
}

func wrapperCondFunc(condFunc any) UniversalCondFunc {
	_ = "STUB: not implemented"
	return *new(UniversalCondFunc)
}

// channelWriteEntry represents a write operation to a channel.
type channelWriteEntry struct {
	Channel  string
	Value    any
	SkipNone bool
	Mapper   func(any) any
}

// Node represents a node in the graph.
// Nodes are primarily functions with metadata.
type Node struct {
	ID          string
	Name        string
	Description string
	Function    NodeFunc
	Type        NodeType // Type of the node (function, llm, tool, etc.)

	// userInputKey is the state key used as one-shot input for LLM and
	// Agent nodes. When empty, StateKeyUserInput is used.
	userInputKey string

	// instruction stores the static instruction for LLM nodes.
	instruction string
	// llmModel stores the static model for LLM nodes.
	llmModel model.Model
	// baseTools stores the static tools configured directly on the node.
	baseTools map[string]tool.Tool

	toolSets             []tool.ToolSet
	refreshToolSetsOnRun bool
	// Per-node callbacks for fine-grained control
	callbacks *NodeCallbacks
	// Optional per-node cache policy. If nil, graph-level policy applies.
	cachePolicy *CachePolicy
	// Optional per-node cache key selector. When set, the executor applies this
	// selector to the sanitized node input before invoking the CachePolicy.KeyFunc.
	// The selector receives a sanitized map[string]any view and should return a
	// projection to be used for key derivation (e.g., subset of fields).
	cacheKeySelector func(map[string]any) any

	// Retry policies configured for this node. When empty, executor defaults
	// (if any) will be used. Policies are evaluated in order to determine
	// whether an error is retryable and which parameters to use.
	retryPolicies []RetryPolicy

	// Pregel-style extensions
	triggers []string            // Channels that trigger this node
	channels []string            // Channels this node reads from
	writers  []channelWriteEntry // Channels this node writes to
	mapper   func(any) any       // Input transformation function

	// Declared destinations for dynamic routing visualization and static checks.
	// Keys are target node IDs; values are optional labels.
	destinations map[string]string

	// ends holds per-node named ends mapping for stronger branch semantics.
	// Keys are symbolic branch names returned by node logic (e.g. "approved",
	// "rejected"). Values are destination node IDs (or the special End).
	// This allows nodes to route via Command.GoTo using symbolic names and for
	// conditional branches to resolve results with clearer, local semantics.
	ends map[string]string

	// It's effect just for LLM node
	modelCallbacks *model.Callbacks
	// just for tool node.
	toolCallbacks *tool.Callbacks
	// toolCallRetryPolicy configures single tool-call retry for tools nodes.
	toolCallRetryPolicy *tool.RetryPolicy

	// enableParallelTools toggles parallel execution for Tools nodes.
	// When true, multiple tool calls in a single assistant response are executed concurrently.
	// Default is false (serial execution) for compatibility and safety.
	enableParallelTools bool

	// llmGenerationConfig stores per-node generation configuration for LLM nodes.
	// If set, AddLLMNode forwards it to the underlying LLM runner.
	llmGenerationConfig *model.GenerationConfig

	// streamOutputName enables node-to-node streaming via agent.StreamHub.
	// For LLM and Agent nodes, it forwards streaming deltas to this stream.
	streamOutputName string

	// interruptBefore pauses execution before this node runs.
	interruptBefore bool
	// interruptAfter pauses execution after this node runs.
	interruptAfter bool

	// Subgraph (agent node) options
	agentInputMapper      SubgraphInputMapper
	agentOutputMapper     SubgraphOutputMapper
	agentIsolatedMessages bool
	agentEventScope       string
	// agentInputFromLastResponse indicates whether the agent node should
	// construct the child invocation's user input from the parent's
	// StateKeyLastResponse. When true, the framework will map
	// last_response -> user_input for this agent node before invoking the
	// sub-agent. This provides a concise way to implement "pass only the
	// result" pipelines between agent nodes without extra glue nodes.
	agentInputFromLastResponse bool

	// traceTransparent marks framework-owned agent nodes that may elide the
	// parent wrapper step from execution trace when all runtime guards pass.
	traceTransparent bool
}

// Edge represents an edge in the graph.
// Simplified edge pattern.
type Edge struct {
	From string
	To   string
}

// ConditionalEdge represents a conditional edge with routing logic.
type ConditionalEdge struct {
	From      string
	Condition UniversalCondFunc
	PathMap   map[string]string // Maps condition result to target node.
}

// Graph represents a directed graph of nodes and edges.
// This is the compiled runtime structure created by StateGraph.Compile().
// Users typically don't create Graph instances directly. Instead, use:
//   - StateGraph for building graphs with compatible patterns.
//
// The Graph type is the immutable runtime representation that gets executed
// by the Executor.
type Graph struct {
	mu               sync.RWMutex
	schema           *StateSchema
	nodes            map[string]*Node
	edges            map[string][]*Edge
	conditionalEdges map[string]*ConditionalEdge
	entryPoint       string
	// Pregel-style extensions
	channelManager *channel.Manager
	triggerToNodes map[string][]string // Maps channel names to nodes that are triggered

	// Caching
	cache       Cache
	cachePolicy *CachePolicy
	// Optional graph version used to scope cache namespaces, helping avoid
	// stale cache collisions across graph code changes or deployments.
	graphVersion string
}

// New creates a new empty graph with the given state schema.
func New(schema *StateSchema) *Graph { _ = "STUB: not implemented"; return nil }

// Node returns a node by ID.
func (g *Graph) Node(id string) (*Node, bool) { _ = "STUB: not implemented"; return nil, false }

// Nodes returns all nodes in the graph sorted by node ID.
func (g *Graph) Nodes() []*Node { _ = "STUB: not implemented"; return nil }

// Edges returns all outgoing edges from a node.
func (g *Graph) Edges(nodeID string) []*Edge { _ = "STUB: not implemented"; return nil }

// ConditionalEdge returns the conditional edge from a node.
func (g *Graph) ConditionalEdge(nodeID string) (*ConditionalEdge, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// EntryPoint returns the entry point node ID.
func (g *Graph) EntryPoint() string { _ = "STUB: not implemented"; return "" }

// Instruction returns the static instruction for an LLM node.
func (n *Node) Instruction() string { _ = "STUB: not implemented"; return "" }

// Model returns the static model for an LLM node.
func (n *Node) Model() model.Model {
	_ = "STUB: not implemented"

	// AgentEventScope returns the configured event scope for an agent node.
	return *new(model.Model)
}

func (n *Node) AgentEventScope() string { _ = "STUB: not implemented"; return "" }

// HasTools reports whether the node has statically configured tools.
func (n *Node) HasTools() bool { _ = "STUB: not implemented"; return false }

// EndTargets returns the concrete end targets declared on the node.
func (n *Node) EndTargets() []string { _ = "STUB: not implemented"; return nil }

// Tools returns the static visible tools of the node.
func (n *Node) Tools(ctx context.Context) []tool.Tool { _ = "STUB: not implemented"; return nil }

// Schema returns the state schema.
func (g *Graph) Schema() *StateSchema {
	_ = "STUB: not implemented"

	// Cache returns the graph-level cache (may be nil).
	return nil
}

func (g *Graph) Cache() Cache { _ = "STUB: not implemented"; return *new(Cache) }

// CachePolicy returns the graph-level cache policy (may be nil).
func (g *Graph) CachePolicy() *CachePolicy { _ = "STUB: not implemented"; return nil }

// setCache sets the graph-level cache.
func (g *Graph) setCache(c Cache) { _ = "STUB: not implemented"; return }

// setCachePolicy sets the graph-level cache policy.
func (g *Graph) setCachePolicy(p *CachePolicy) { _ = "STUB: not implemented"; return }

// setGraphVersion sets an optional version string used for cache namespacing.
func (g *Graph) setGraphVersion(v string) { _ = "STUB: not implemented"; return }

// cacheNamespace builds a per-node namespace including optional graph version.
func (g *Graph) cacheNamespace(nodeID string) string { _ = "STUB: not implemented"; return "" }

// clearCacheForNodes clears cache entries for the given node IDs.
func (g *Graph) clearCacheForNodes(nodes []string) { _ = "STUB: not implemented"; return }

// validate validates the graph structure.
func (g *Graph) validate() error { _ = "STUB: not implemented"; return nil }

// Validate declared destinations exist.

// fallthrough to ends check

// Validate per-node ends mapping targets exist.

// ExecutionContext contains context for graph execution.
type ExecutionContext struct {
	Graph        *Graph
	EventChan    chan<- *event.Event
	InvocationID string
	// Invocation is the per-run invocation context. Nodes may use it to
	// read invocation-scoped state (for example, {invocation:*} placeholders).
	Invocation *agent.Invocation
	// channels holds the per-execution Pregel channels. These are constructed
	// from the Graph's static channel definitions when the execution starts
	// and are never shared across concurrent runs.
	channels *channel.Manager
	// stateMutex protects State reads/writes.
	stateMutex sync.RWMutex
	State      State
	// completionIdentity* carry internal agent-node completion identity when
	// the public graph state does not expose StateKeyLastResponseID.
	completionIdentityText string
	completionIdentity     string
	// pendingMu protects pendingWrites operations.
	pendingMu     sync.Mutex
	pendingWrites []PendingWrite
	resumed       bool
	seq           atomic.Int64 // Atomic sequence counter for deterministic replay
	// tasksMutex protects pendingTasks queue operations.
	tasksMutex   sync.Mutex
	pendingTasks []*Task
	// versionsSeen tracks which channel versions each node has seen.
	// Map from nodeID -> channelName -> version number.
	versionsSeen   map[string]map[string]int64
	versionsSeenMu sync.RWMutex
	// lastCheckpoint holds the most recent checkpoint used for planning
	// when resuming. Keeping this per-execution avoids cross-run sharing
	// when a single Executor is reused concurrently.
	lastCheckpoint *Checkpoint
	// traceMu protects execution trace planning and task-to-step bookkeeping.
	traceMu sync.Mutex
	// traceChannelSources tracks the source step ids currently attached to a channel.
	traceChannelSources map[string][]string
	// traceChannelSourceSteps tracks the step number for last-value/ephemeral provenance.
	traceChannelSourceSteps map[string]int
	// traceBarrierChannelSources tracks the latest source step ids for each barrier sender.
	traceBarrierChannelSources map[string]map[string][]string
	// traceSourceStepIDsByTaskID tracks the trace source steps produced by each task.
	traceSourceStepIDsByTaskID map[string][]string
	// traceAgentNodeTasksByNodeID tracks transparent agent-node candidates.
	traceAgentNodeTasksByNodeID map[string]*traceTaskRegistryEntry
}

func (e *ExecutionContext) setCompletionIdentity(text, identity string) {
	_ = "STUB: not implemented"
	return
}

func (e *ExecutionContext) snapshotCompletionState(
	fields map[string]StateField,
) (State, string, string) {
	_ = "STUB: not implemented"
	return *new(State), "", ""
}

// Command represents a command that combines state updates with routing.
type Command struct {
	Update    State
	GoTo      string
	Resume    any
	ResumeMap map[string]any
}

// addNode adds a node to the graph.
func (g *Graph) addNode(node *Node) error { _ = "STUB: not implemented"; return nil }

// addEdge adds an edge to the graph.
func (g *Graph) addEdge(edge *Edge) error { _ = "STUB: not implemented"; return nil }

// Allow Start and End as special nodes

// addConditionalEdge adds a conditional edge to the graph.
func (g *Graph) addConditionalEdge(condEdge *ConditionalEdge) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate condition presence and exclusivity.

// Validate all target nodes in path map

// setEntryPoint sets the entry point of the graph.
func (g *Graph) setEntryPoint(nodeID string) error { _ = "STUB: not implemented"; return nil }

// Pregel-style methods

// addChannel adds a channel to the graph.
func (g *Graph) addChannel(name string, channelType channel.Behavior) {
	_ = "STUB: not implemented"
	return
}

// getChannel retrieves a channel definition by name. This is primarily used
// during graph construction and in tests. Runtime execution should operate on
// the per-execution channels stored in ExecutionContext.
func (g *Graph) getChannel(name string) (*channel.Channel, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// getAllChannels returns all channel definitions in the graph. Callers must
// treat the returned channels as immutable templates (name + behavior).
// Per-execution channel state (values, versions, availability) is stored in
// ExecutionContext.channels.
func (g *Graph) getAllChannels() map[string]*channel.Channel { _ = "STUB: not implemented"; return nil }

// getTriggerToNodes returns the mapping of channels to triggered nodes.
func (g *Graph) getTriggerToNodes() map[string][]string { _ = "STUB: not implemented"; return nil }

// addNodeTrigger adds a trigger relationship between a channel and a node.
func (g *Graph) addNodeTrigger(channelName string, nodeID string) {
	_ = "STUB: not implemented"
	return
}

// Deduplicate

// addNodeWriter adds a writer to a node.
func (g *Graph) addNodeWriter(nodeID string, writer channelWriteEntry) {
	_ = "STUB: not implemented"
	return
}

// addNodeTrigger adds a trigger to a node.
func (g *Graph) addNodeTriggerChannel(nodeID string, channelName string) {
	_ = "STUB: not implemented"
	return
}

// addNodeChannel adds a channel that a node reads from.
func (g *Graph) addNodeChannel(nodeID string, channelName string) {
	_ = "STUB: not implemented"
	return
}
