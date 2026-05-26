//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"context"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph/internal/channel"
	itelemetry "trpc.group/trpc-go/trpc-agent-go/internal/telemetry"
)

const (
	// AuthorGraphExecutor is the author of the graph executor.
	AuthorGraphExecutor = "graph-executor"
)

var (
	defaultChannelBufferSize     = 256
	defaultMaxSteps              = 100
	defaultStepTimeout           = time.Duration(0) // No timeout by default, users can set if needed.
	defaultCheckpointSaveTimeout = 10 * time.Second // Default timeout for checkpoint save operations.
	defaultBarrierWaitTimeout    = 5 * time.Second  // Default timeout for barrier completion waits.
)

func defaultMaxConcurrency() int { _ = "STUB: not implemented"; return 0 }

// Executor executes a graph with the given initial state using the configured
// execution engine (default: Pregel-style BSP).
//
// Runtime isolation principle:
//   - Executor is designed to be reusable across concurrent runs.
//   - It must not hold per-run mutable state; such data lives in ExecutionContext.
//   - Checkpoint-derived artifacts (e.g., lastCheckpoint, pendingWrites) are
//     carried inside ExecutionContext and never stored on the Executor.
//
// This makes it safe to share a single Executor instance between many
// concurrent invocations without cross-run interference.
type Executor struct {
	graph                 *Graph
	channelBufferSize     int
	maxSteps              int
	maxConcurrency        int
	executionEngine       ExecutionEngine
	stepTimeout           time.Duration
	nodeTimeout           time.Duration
	checkpointSaveTimeout time.Duration
	checkpointSaver       CheckpointSaver
	checkpointManager     *CheckpointManager
	// defaultRetry holds executor-level retry policies used when a node
	// does not have explicit retryPolicies configured.
	defaultRetry []RetryPolicy
}

// ExecutorOption is a function that configures an Executor.
type ExecutorOption func(*ExecutorOptions)

// ExecutorOptions contains configuration options for creating an Executor.
type ExecutorOptions struct {
	// ChannelBufferSize is the buffer size for event channels (default: 256).
	ChannelBufferSize int
	// MaxSteps is the maximum number of steps for graph execution.
	MaxSteps int
	// MaxConcurrency is the maximum number of tasks executed in parallel.
	//
	// When <= 0, it defaults to runtime.GOMAXPROCS(0).
	MaxConcurrency int
	// StepTimeout is the timeout for each step (default: 0 = no timeout).
	StepTimeout time.Duration
	// NodeTimeout is the timeout for individual node execution
	// (default: derived from StepTimeout/2 when StepTimeout>0; otherwise no timeout).
	NodeTimeout time.Duration
	// CheckpointSaveTimeout is the timeout for saving checkpoints (default: 10s).
	CheckpointSaveTimeout time.Duration
	// CheckpointSaver is the checkpoint saver for persisting graph state.
	CheckpointSaver CheckpointSaver
	// ExecutionEngine controls how the graph is scheduled and executed.
	//
	// The default is ExecutionEngineBSP.
	ExecutionEngine ExecutionEngine
	// DefaultRetryPolicies are applied to nodes without explicit policies.
	DefaultRetryPolicies []RetryPolicy
}

// WithChannelBufferSize sets the buffer size for event channels.
func WithChannelBufferSize(size int) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithMaxSteps sets the maximum number of steps for graph execution.
func WithMaxSteps(maxSteps int) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithMaxConcurrency sets the maximum number of tasks executed in parallel.
//
// When max <= 0, it uses the default value (runtime.GOMAXPROCS(0)).
func WithMaxConcurrency(max int) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithStepTimeout sets the timeout for each step.
func WithStepTimeout(timeout time.Duration) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithNodeTimeout sets the timeout for individual node execution.
func WithNodeTimeout(timeout time.Duration) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithCheckpointSaver sets the checkpoint saver for the executor.
func WithCheckpointSaver(saver CheckpointSaver) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithCheckpointSaveTimeout sets the timeout for checkpoint save operations.
func WithCheckpointSaveTimeout(timeout time.Duration) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithExecutionEngine sets the execution engine for scheduling.
//
// The default is ExecutionEngineBSP.
func WithExecutionEngine(engine ExecutionEngine) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// WithDefaultRetryPolicy sets executor-level retry policies used by nodes
// that do not define their own. Policies are evaluated in order.
func WithDefaultRetryPolicy(policies ...RetryPolicy) ExecutorOption {
	_ = "STUB: not implemented"
	return *new(ExecutorOption)
}

// NewExecutor creates a new graph executor.
func NewExecutor(graph *Graph, opts ...ExecutorOption) (*Executor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply function options.

// Calculate node timeout: use provided value or derive from step timeout if step timeout is set.

// Only derive from step timeout if step timeout is explicitly set.

// Create checkpoint manager if saver is provided.

// Task represents a task to be executed in a step.
type Task struct {
	NodeID             string              // NodeID is the ID of the node to execute.
	Input              any                 // Input is the input of the task.
	Writes             []channelWriteEntry // Writes is the writes of the task.
	Triggers           []string            // Triggers is the triggers of the task.
	TaskID             string              // TaskID is the ID of the task.
	TaskPath           []string            // TaskPath is the path of the task.
	Overlay            State               // Overlay is the overlay state of the task.
	PredecessorStepIDs []string            // PredecessorStepIDs is the direct predecessor step ids of the task.
}

// Step represents a single step in execution.
type Step struct {
	StepNumber      int             // StepNumber is the number of the step.
	Tasks           []*Task         // Tasks is the tasks of the step.
	State           State           // State is the state of the step.
	UpdatedChannels map[string]bool // UpdatedChannels is the updated channels of the step.
}

// Execute executes the graph with the given initial state.
func (e *Executor) Execute(
	ctx context.Context,
	initialState State,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the internal event channel used by graph execution.

// Start execution in a goroutine.

// Check if this is an interrupt error.

// For interrupt errors, we don't emit an error event.
// The interrupt will be handled by the caller.

// Emit error event for other errors.

func (e *Executor) forwardExecutionEvents(
	ctx context.Context,
	invocation *agent.Invocation,
	src <-chan *event.Event,
	dst chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func shouldHideExecutorBarrierEvents(invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

func isGraphNodeBarrierEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func notifySuppressedBarrierCompletion(
	ctx context.Context,
	invocation *agent.Invocation,
	evt *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

// executeGraph executes the graph using Pregel-style BSP execution.
func (e *Executor) executeGraph(
	ctx context.Context,
	initialState State,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
	startTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Build execution context (including per-execution channels) from the prepared state.

// Initialize per-execution input channels from the prepared state.

// prepareCheckpointAndState initializes or restores state and checkpointing.
func (e *Executor) prepareCheckpointAndState(
	ctx context.Context,
	initialState State,
	invocation *agent.Invocation,
) (State, map[string]any, bool, int, *Checkpoint, []PendingWrite, error) {
	_ = "STUB: not implemented"
	return *new(State), nil, false, 0, nil, nil, nil
}

// resumeOrInitWithSaver handles state preparation when checkpoint saver is set.
func (e *Executor) resumeOrInitWithSaver(
	ctx context.Context,
	initialState State,
	invocation *agent.Invocation,
) (State, map[string]any, bool, int, *Checkpoint, []PendingWrite, error) {
	_ = "STUB: not implemented"
	return *new(State), nil, false, 0, nil, nil, nil
}

// restoreStateFromCheckpoint converts checkpoint channel values back into state.
// It mirrors the original inline logic: convert values to schema field types,
// then add any missing schema defaults or zero values so downstream nodes see
// consistent shapes, exactly as prior to refactor.
func (e *Executor) restoreStateFromCheckpoint(tuple *CheckpointTuple) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// mergeInitialStateNonInternal merges caller-provided initial values that are
// not internal (do not start with "_"). By default checkpoint-restored values
// win, but runtime-state override keys can explicitly override the restored
// state during resume.
func (e *Executor) mergeInitialStateNonInternal(
	restored,
	initial State,
	resumeStateOverrideKeys map[string]struct{},
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// applyExecutableNextNodes sets StateKeyNextNodes when suitable.
// This is important for initial checkpoints (step -1) that have the entry
// point set, so a forked resume can continue from the beginning.
func (e *Executor) applyExecutableNextNodes(
	restored State,
	tuple *CheckpointTuple,
) {
	_ = "STUB: not implemented"
	return
}

func (e *Executor) applyGraphInterruptInputs(
	restored State,
	tuple *CheckpointTuple,
) {
	_ = "STUB: not implemented"
	return
}

func copyGraphInterruptInputsState(
	inputs map[string]State,
) map[string]State {
	_ = "STUB: not implemented"
	return nil
}

func copyGraphInterruptInputsAny(inputs map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func copyGraphInterruptInputsAnySlice(
	inputs map[string][]any,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func copyGraphInterruptInputsStateSlice(
	inputs map[string][]State,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// processResumeCommand applies resume-related fields from the initial state.
func (e *Executor) processResumeCommand(execState, initialState State) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// Apply resume values if present.

// restoreVersionsSeen restores per-node versionsSeen from the last
// checkpoint.
func (e *Executor) restoreVersionsSeen(
	resumed bool,
	lastCheckpoint *Checkpoint,
) map[string]map[string]int64 {
	_ = "STUB: not implemented"
	return nil
}

// buildChannelManager creates per-execution channels from the graph's static
// channel definitions.
func (e *Executor) buildChannelManager() *channel.Manager { _ = "STUB: not implemented"; return nil }

// restoreChannelVersions seeds channel versions from the last checkpoint for
// resumed executions.
func (e *Executor) restoreChannelVersions(
	execCtx *ExecutionContext,
	resumed bool,
	lastCheckpoint *Checkpoint,
) {
	_ = "STUB: not implemented"
	return
}

// restoreBarrierSets restores barrier sets from the last checkpoint for resumed
// executions.
func (e *Executor) restoreBarrierSets(
	execCtx *ExecutionContext,
	resumed bool,
	lastCheckpoint *Checkpoint,
) {
	_ = "STUB: not implemented"
	return
}

// buildExecutionContext constructs the execution context including versionsSeen.
func (e *Executor) buildExecutionContext(
	eventChan chan<- *event.Event,
	invocationID string,
	state State,
	resumed bool,
	lastCheckpoint *Checkpoint,
) *ExecutionContext {
	_ = "STUB: not implemented"
	return nil
}

// For resumed executions, seed channel versions from the last checkpoint so
// version-based triggering semantics can continue to function correctly.

// runBspLoop runs the BSP execution loop from the given start step.
func (e *Executor) runBspLoop(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	checkpointConfig *map[string]any,
	startStep int,
	extInterrupt *externalInterruptWatcher,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *Executor) runBspStep(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	checkpointConfig *map[string]any,
	startStep int,
	step int,
	extInterrupt *externalInterruptWatcher,
) (stop bool, executed bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func (e *Executor) stepContext(
	ctx context.Context,
) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (e *Executor) planTasksForBspStep(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	startStep int,
	step int,
) ([]*Task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nextNodesFromTasks(tasks []*Task) []string { _ = "STUB: not implemented"; return nil }

func (e *Executor) metaExtraForPlannedExternalInterrupt(
	tasks []*Task,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) stateFields() map[string]StateField { _ = "STUB: not implemented"; return nil }

func stateFromAny(v any) (State, bool) { _ = "STUB: not implemented"; return *new(State), false }

func (e *Executor) maybeHandleExternalInterruptBeforeStep(
	ctx context.Context,
	stepCtx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
	checkpointConfig *map[string]any,
	extInterrupt *externalInterruptWatcher,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Executor) maybeHandleStaticInterruptBeforeStep(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
	checkpointConfig *map[string]any,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Executor) newStepExecutionReportIfNeeded(
	extInterrupt *externalInterruptWatcher,
) *stepExecutionReport {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) executeStepWithInterruptHandling(
	stepCtx context.Context,
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
	checkpointConfig *map[string]any,
	report *stepExecutionReport,
	extInterrupt *externalInterruptWatcher,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) handleExecuteStepError(
	stepCtx context.Context,
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
	checkpointConfig map[string]any,
	report *stepExecutionReport,
	extInterrupt *externalInterruptWatcher,
	stepErr error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) shouldForceExternalInterrupt(
	extInterrupt *externalInterruptWatcher,
	ctx context.Context,
	stepErr error,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Executor) handleForcedExternalInterrupt(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
	checkpointConfig map[string]any,
	report *stepExecutionReport,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) tasksToRerun(
	tasks []*Task,
	report *stepExecutionReport,
) []*Task {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) nextNodesForForcedInterrupt(
	execCtx *ExecutionContext,
	rerun []*Task,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) metaExtraForForcedInterrupt(
	report *stepExecutionReport,
	rerun []*Task,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) updateChannelsForStep(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	step int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) maybeHandleStaticInterruptAfterStep(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
	checkpointConfig *map[string]any,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Executor) maybeCreateLoopCheckpoint(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	checkpointConfig *map[string]any,
	step int,
) {
	_ = "STUB: not implemented"
	return
}

// buildCompletionEvent prepares the completion event with a state snapshot.
func (e *Executor) buildCompletionEvent(
	execCtx *ExecutionContext,
	startTime time.Time,
	stepsExecuted int,
) *event.Event {
	_ = "STUB: not implemented"
	// Take a deep snapshot of the final state under read lock.
	// IMPORTANT: Skip volatile/non-serializable keys (e.g., Session, callbacks, exec context)
	// to avoid racing on their internal maps/slices managed by other goroutines.
	return nil
}

func resolveCompletionResponseID(
	finalState State,
	identityText string,
	identity string,
) string {
	_ = "STUB: not implemented"
	return ""
}

// createCheckpointAndSave creates a checkpoint and persists any pending writes
// associated with the current step atomically, updating the provided config with the
// returned value from saver.PutFull (which may include the new checkpoint_id).
func (e *Executor) createCheckpointAndSave(
	ctx context.Context,
	invocation *agent.Invocation,
	config *map[string]any,
	source string,
	step int,
	execCtx *ExecutionContext,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Checkpoint saver is nil.

// Create checkpoint object.

// Set parent checkpoint ID from config if available.

// Set parent checkpoint ID.

// Created checkpoint object.

// Create metadata.

// Get pending writes atomically.

// Clear after copying.

// Track new versions for channels that were updated on this execution.

// Persist all per-run channel versions for correct resume semantics.
// Version-based triggering relies on monotonic channel versions even when a
// channel is not currently "available" (it may have been acknowledged).

// Set next nodes and channels for recovery.

// For initial checkpoints, set the entry point as the next node.
// This ensures that if someone forks and resumes from this checkpoint,
// the workflow will start from the beginning.

// Use PutFull for atomic storage.

// Successfully saved checkpoint.
// Clear step marks after checkpoint creation.

// Update external config with the new checkpoint_id.

// Updated config with new checkpoint ID.

func shouldEmitCheckpointLifecycleEvents(
	invocation *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldEmitPregelStepEvents(invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

func emitTerminalGraphErrorEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

// applyPendingWrites replays pending writes into channels to rebuild frontier.
func (e *Executor) applyPendingWrites(ctx context.Context, invocation *agent.Invocation,
	execCtx *ExecutionContext, writes []PendingWrite) {
	_ = "STUB: not implemented"
	return
}

// Sort writes by sequence number for deterministic replay.

// Emit channel update event to mirror live execution behavior.

// getConfigKeys helper to extract keys from config map for logging
func getConfigKeys(config map[string]any) []string { _ = "STUB: not implemented"; return nil }

// initializeState initializes the execution state with schema defaults.
func (e *Executor) initializeState(initialState State) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// Add schema defaults for missing fields.

// Use default function if available, otherwise provide zero value.

// initializeChannels initializes channels with input state.
// If updateChannels is false, only registers channels without triggering updates.
// Channels are created on the per-execution channel manager stored in the
// ExecutionContext to avoid sharing mutable channel state across runs.
func (e *Executor) initializeChannels(execCtx *ExecutionContext, state State, updateChannels bool) {
	_ = "STUB: not implemented"
	return
}

// planStep determines which nodes to execute in the current step.
func (e *Executor) planStep(ctx context.Context, invocation *agent.Invocation,
	execCtx *ExecutionContext, step int) ([]*Task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if we have nodes to execute from a resumed checkpoint stored in state
// This needs to be checked regardless of step number when resuming

// Create tasks for the nodes stored in the state

// Remove the special key from state after using it

// If there are pending tasks produced by prior fan-out, schedule them first.

// Check if this is the first step (entry point).

// Use the normal entry point

// Plan based on channel triggers.

// planBasedOnChannelTriggers creates tasks for nodes triggered by channel updates.
func (e *Executor) planBasedOnChannelTriggers(execCtx *ExecutionContext, step int) []*Task {
	_ = "STUB: not implemented"
	return nil
}

// If this is a resumed execution, use version-based triggering

// Use traditional availability-based triggering

// planBasedOnVersionTriggers creates tasks based on per-node version tracking.
func (e *Executor) planBasedOnVersionTriggers(execCtx *ExecutionContext, step int) []*Task {
	_ = "STUB: not implemented"
	return nil
}

// planBasedOnAvailabilityTriggers creates tasks based on channel availability.
func (e *Executor) planBasedOnAvailabilityTriggers(
	execCtx *ExecutionContext,
	step int,
	triggerToNodes map[string][]string,
) []*Task {
	_ = "STUB: not implemented"
	return nil
}

// Don't log error for virtual end node - it's expected.

// createTask creates a task for a node.
func (e *Executor) createTask(nodeID string, state State, step int) *Task {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) createTaskWithPredecessors(
	nodeID string,
	state State,
	step int,
	predecessors []string,
) *Task {
	_ = "STUB: not implemented"
	// Handle virtual end node - it doesn't need to be executed.
	return nil
}

func (e *Executor) createTriggeredTasks(
	execCtx *ExecutionContext,
	step int,
	nodeTriggers map[string][]string,
) []*Task {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) tracePredecessorsForChannels(
	execCtx *ExecutionContext,
	channelNames []string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) clearTraceChannelSources(execCtx *ExecutionContext, channelNames []string) {
	_ = "STUB: not implemented"
	return
}

func (e *Executor) recordTraceChannelSource(
	execCtx *ExecutionContext,
	channelName string,
	senderKey string,
	stepID string,
	step int,
) {
	_ = "STUB: not implemented"
	return
}

func (e *Executor) recordTraceChannelSources(
	execCtx *ExecutionContext,
	channelName string,
	senderKey string,
	stepIDs []string,
	step int,
) {
	_ = "STUB: not implemented"
	return
}

func traceChannelBehavior(execCtx *ExecutionContext, channelName string) channel.Behavior {
	_ = "STUB: not implemented"
	return *new(channel.Behavior)
}

func (e *Executor) recordTraceSourceStepIDs(execCtx *ExecutionContext, taskID string, stepIDs []string) {
	_ = "STUB: not implemented"
	return
}

func (e *Executor) traceSourceStepIDsForTask(execCtx *ExecutionContext, taskID string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) clearTraceSourceStepIDs(execCtx *ExecutionContext, taskID string) {
	_ = "STUB: not implemented"
	return
}

func normalizeTraceStepIDs(stepIDs []string) []string { _ = "STUB: not implemented"; return nil }

func consumeGraphInterruptInput(state State, nodeID string) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

func consumeGraphInterruptInputState(
	state State,
	inputs map[string]State,
	nodeID string,
) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

func consumeGraphInterruptInputStates(
	state State,
	inputs map[string][]State,
	nodeID string,
) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

func consumeGraphInterruptInputAny(
	state State,
	inputs map[string]any,
	nodeID string,
) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

func consumeGraphInterruptInputAnySlice(
	state State,
	inputs map[string][]any,
	nodeID string,
) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

func consumeStateFromStateSlice(values []State) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

func consumeStateFromAnySlice(values []any) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

func cleanupGraphInterruptInputs(state State, inputs map[string]any) {
	_ = "STUB: not implemented"
	return
}

// executeStep executes all tasks concurrently.
func (e *Executor) executeStep(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
	report *stepExecutionReport,
) error {
	_ = "STUB: not implemented"
	// Emit execution step event.
	return nil
}

func sameNodeDuplicateTaskSet(tasks []*Task) map[*Task]bool { _ = "STUB: not implemented"; return nil }

func (e *Executor) workerCount(taskCount int) int { _ = "STUB: not implemented"; return 0 }

func (e *Executor) dispatchTasks(tasksCh chan<- *Task, tasks []*Task) {
	_ = "STUB: not implemented"
	return
}

func (e *Executor) executeStepTask(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	step int,
	report *stepExecutionReport,
	sameNodeDuplicate bool,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) taskInvocationContext(
	ctx context.Context,
	invocation *agent.Invocation,
	t *Task,
) (*agent.Invocation, context.Context) {
	_ = "STUB: not implemented"
	return nil, *new(context.Context)
}

// emitExecutionStepEvent emits the execution step event.
func (e *Executor) emitExecutionStepEvent(ctx context.Context, invocation *agent.Invocation,
	execCtx *ExecutionContext, tasks []*Task, step int) {
	_ = "STUB: not implemented"
	return
}

func (e *Executor) emitNodeBarrierAndWait(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	nodeID string,
	step int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// executeSingleTask executes a single task and handles all its events.
func (e *Executor) executeSingleTask(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	step int,
	report *stepExecutionReport,
	sameNodeDuplicate bool,
) error {
	_ = "STUB: not implemented"
	// Initialize node execution context with retry policies and metadata.
	return nil
}

// Run before node callbacks.

// Ensure pre-callback state mutations are visible to the node function.
// We pass the callback-mutated state copy as the task input so that
// executeNodeFunction uses it (instead of rebuilding from the global state).
// This preserves overlay application done in buildTaskStateCopy and respects
// any in-place state changes made by before-node callbacks.

// Attempt cache lookup; if hit, handle cached result and return.

// Execute with retry logic (emits completion event downstream on success).

// initializeNodeContext initializes the node execution context with all
// necessary metadata, policies, and callbacks.
func (e *Executor) initializeNodeContext(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	step int,
	sameNodeDuplicate bool,
) *nodeExecutionContext {
	_ = "STUB: not implemented"
	// Get node type and determine retry policies for metadata.
	return nil
}

// Best-effort max attempts hint for start event (first policy wins).

// Emit node start event with attempt metadata.

// Create callback context.

func (e *Executor) canCreateTransparentAgentNodeCandidate(
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	nodeType NodeType,
	nodePolicies []RetryPolicy,
	callbacks *NodeCallbacks,
	sameNodeDuplicate bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func hasEffectiveNodeCallbacks(callbacks *NodeCallbacks) bool {
	_ = "STUB: not implemented"
	return false
}

func hasEffectiveAgentCallbacks(callbacks *agent.Callbacks) bool {
	_ = "STUB: not implemented"
	return false
}

func pluginAgentCallbacksFromInvocation(invocation *agent.Invocation) *agent.Callbacks {
	_ = "STUB: not implemented"
	return nil
}

// getNodeRetryPolicies retrieves retry policies for the given node.
// Returns node-specific policies if available, otherwise returns default policies.
func (e *Executor) getNodeRetryPolicies(nodeID string) []RetryPolicy {
	_ = "STUB: not implemented"
	return nil
}

// getMaxAttemptsHint extracts the max attempts hint from retry policies.
// Returns the MaxAttempts value from the first policy, or 0 if not available.
func (e *Executor) getMaxAttemptsHint(policies []RetryPolicy) int {
	_ = "STUB: not implemented"
	return 0
}

// attemptCacheLookup attempts to retrieve a cached result for the task.
// Returns true and the cached result if found, false otherwise.
func (e *Executor) attemptCacheLookup(t *Task) (bool, any) {
	_ = "STUB: not implemented"
	return false, *new(any)
}

// Apply optional cache key selector (node-level) to focus on relevant inputs.

func applyCacheKeySelector(
	selector func(map[string]any) any,
	input any,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// handleCachedResult processes a cache hit by running callbacks and handling
// the result without executing the node function.
func (e *Executor) handleCachedResult(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	result any,
	step int,
	nodeCtx *nodeExecutionContext,
) error {
	_ = "STUB: not implemented"
	// Run after node callbacks on cache hit.
	return nil
}

// Handle result and process channel writes.

// Update versions seen for this node after successful execution.

// Process conditional edges after node execution.

// Emit node completion event with cache-hit metadata.

// nodeExecutionContext holds node execution related state.
type nodeExecutionContext struct {
	nodeType        NodeType
	nodeStart       time.Time
	nodePolicies    []RetryPolicy
	callbackCtx     *NodeCallbackContext
	stateCopy       State
	mergedCallbacks *NodeCallbacks
	traceStepID     string
	metricRecorder  *workflowMetricRecorder
	traceTask       *traceTaskMetadata
}

type workflowMetricRecorder struct {
	once       sync.Once
	start      time.Time
	attributes itelemetry.WorkflowAttributes
}

func (r *workflowMetricRecorder) recordSuccess(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (r *workflowMetricRecorder) recordError(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

func (r *workflowMetricRecorder) record(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// executeTaskWithRetry executes the task with retry logic.
func (e *Executor) executeTaskWithRetry(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	step int,
	nodeCtx *nodeExecutionContext,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Track total elapsed for optional policy.MaxElapsedTime evaluation.

// Execute single attempt.

// Handle successful execution.

// Check if should retry.

// reserved for future metrics

// executeSingleAttempt executes a single attempt of the node function.
func (e *Executor) executeSingleAttempt(
	ctx context.Context,
	execCtx *ExecutionContext,
	t *Task,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// finalizeSuccessfulExecution handles all post-execution steps after successful node execution.
func (e *Executor) finalizeSuccessfulExecution(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	result any,
	step int,
	nodeCtx *nodeExecutionContext,
) error {
	_ = "STUB: not implemented"
	// Run after node callbacks on success.
	return nil
}

// Handle result and process channel writes.

// After successful writes, persist cache entry if cache policy exists.

// Use the same sanitized input used for lookup (post-callback state copy).

// Update versions seen for this node after successful execution.

// Process conditional edges after node execution.

// Emit node completion event for the overall node run (no cache hit).

func (e *Executor) finalizeFailedExecution(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	result any,
	retryErr error,
	nodeErr error,
	step int,
	nodeCtx *nodeExecutionContext,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) finalizeRecoveredExecution(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	result any,
	step int,
	nodeCtx *nodeExecutionContext,
) error {
	_ = "STUB: not implemented"
	return nil
}

// retryContext holds retry evaluation state.
type retryContext struct {
	attempt    int
	totalStart time.Time
	err        error
}

// evaluateRetryDecision determines if a retry should occur and handles error reporting.
func (e *Executor) evaluateRetryDecision(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	step int,
	nodeCtx *nodeExecutionContext,
	retryCtx *retryContext,
) (bool, error) {
	_ = "STUB: not implemented"
	// Interrupt errors should not be retried.
	return false, nil
}

// Run on-node-error callbacks for observability (both intermediate and final).

// Evaluate retry policy.

// No retry policy matched -> emit error and exit.

// Check if retry budget is exhausted.

// Emit error event with retrying metadata and wait.

// checkRetryBudget checks if retry attempts or time budget is exhausted.
func (e *Executor) checkRetryBudget(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	nodeID string,
	step int,
	nodeCtx *nodeExecutionContext,
	retryCtx *retryContext,
	pol RetryPolicy,
	maxAttempts int,
) (bool, error) {
	_ = "STUB: not implemented"
	// Check attempt budget.
	return false, nil
}

// Check elapsed time budget.

// waitBeforeRetry handles the delay before retry and deadline checking.
func (e *Executor) waitBeforeRetry(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	nodeID string,
	step int,
	nodeCtx *nodeExecutionContext,
	retryCtx *retryContext,
	pol RetryPolicy,
	maxAttempts int,
) (bool, error) {
	_ = "STUB: not implemented"
	// Compute delay and clamp to parent context deadline if present.
	return false, nil
}

// Emit error event with retrying metadata.

// Sleep or abort if context canceled.

// getNodeType retrieves the node type for a given node ID.
func (e *Executor) getNodeType(nodeID string) NodeType {
	_ = "STUB: not implemented"
	return *new(NodeType)
}

// Default fallback.

// getNodeName retrieves the node name for a given node ID.
func (e *Executor) getNodeName(nodeID string) string { _ = "STUB: not implemented"; return "" }

// Default to node ID if node not found.

func (e *Executor) newWorkflowMetricRecorder(
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	nodeID string,
	nodeType NodeType,
	start time.Time,
) *workflowMetricRecorder {
	_ = "STUB: not implemented"
	return nil
}

// getSessionID retrieves the session ID from the execution context.
func (e *Executor) getSessionID(execCtx *ExecutionContext) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *Executor) getSessionIdentity(execCtx *ExecutionContext) (appName string, userID string) {
	_ = "STUB: not implemented"
	return "", ""
}

// newNodeContext creates a context for a single node execution with timeout.
func (e *Executor) newNodeContext(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// newNodeCallbackContext builds callback context for node lifecycle events.
func (e *Executor) newNodeCallbackContext(
	execCtx *ExecutionContext,
	nodeID string,
	nodeType NodeType,
	step int,
	start time.Time,
) *NodeCallbackContext {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) traceNodeIDForTask(invocation *agent.Invocation, t *Task) string {
	_ = "STUB: not implemented"
	return ""
}

func traceSnapshotFromValue(value any) *atrace.Snapshot { _ = "STUB: not implemented"; return nil }

func marshalTraceSnapshot(value any) *atrace.Snapshot { _ = "STUB: not implemented"; return nil }

// buildTaskStateCopy returns the per-task input state, including overlay.
func (e *Executor) buildTaskStateCopy(execCtx *ExecutionContext, t *Task) State {
	_ = "STUB: not implemented"
	// Always construct an isolated state copy so node code can freely mutate
	// without racing with other goroutines. Skip or shallow-copy unsafe keys
	// whose internals may be mutated concurrently by other subsystems.
	return *new(State)
}

// Apply overlay if present to form the isolated input view.

// Inject execution context helpers used by nodes.

// getMergedCallbacks merges global and per-node callbacks for a node.
func (e *Executor) getMergedCallbacks(stateCopy State, nodeID string) *NodeCallbacks {
	_ = "STUB: not implemented"
	return nil
}

// runBeforeCallbacks executes before-node callbacks and handles early result.
func (e *Executor) runBeforeCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	callbacks *NodeCallbacks,
	cbCtx *NodeCallbackContext,
	stateCopy State,
	execCtx *ExecutionContext,
	t *Task,
	nodeType NodeType,
	nodeStart time.Time,
	metricRecorder *workflowMetricRecorder,
	step int,
	traceStepID string,
	traceTask *traceTaskMetadata,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// We need to skip intermediate nodes after routed.

// runAfterCallbacks executes after-node callbacks and returns an override.
func (e *Executor) runAfterCallbacks(
	ctx context.Context,
	invocation *agent.Invocation,
	callbacks *NodeCallbacks,
	cbCtx *NodeCallbackContext,
	stateCopy State,
	result any,
	nodeErr error,
	execCtx *ExecutionContext,
	nodeID string,
	nodeType NodeType,
	step int,
	metricRecorder *workflowMetricRecorder,
) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func runAfterNodeCallbacks(
	ctx context.Context,
	callbacks *NodeCallbacks,
	callbackCtx *NodeCallbackContext,
	state State,
	result any,
	nodeErr error,
) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

// mergeNodeCallbacks merges global and per-node callbacks.
// Global callbacks are executed first, followed by per-node callbacks.
// This allows per-node callbacks to override or extend global behavior.
func (e *Executor) mergeNodeCallbacks(global, perNode *NodeCallbacks) *NodeCallbacks {
	_ = "STUB: not implemented"
	return nil
}

// Create a new merged callbacks instance.

// Add global callbacks first for Before and OnNodeError.

// For per-node callbacks, Before callbacks execute after global.

// For After callbacks, execute per-node first, then global, so per-node can
// shape/override the result before global observers run.

// emitNodeStartEvent emits the node start event.
func (e *Executor) emitNodeStartEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	nodeID string,
	nodeType NodeType,
	step int,
	startTime time.Time,
	extra ...NodeEventOption,
) {
	_ = "STUB: not implemented"
	return
}

// Extract model input for LLM nodes.

// Build event with optional extra metadata (e.g., retries)

// executeNodeFunction executes the actual node function.
func (e *Executor) executeNodeFunction(
	ctx context.Context,
	execCtx *ExecutionContext,
	t *Task,
) (res any, err error) {
	_ = "STUB: not implemented"
	// Recover from panics in user-provided node functions to prevent
	// the whole service from crashing. Convert to error so the normal
	// error handling path (callbacks, events, checkpointing) can run.
	return *new(any), nil
}

// Prefer the prebuilt task input which is already a deep copy created by
// buildTaskStateCopy. If missing (e.g., legacy paths), deep-copy the
// current global state here as a fallback.

// Apply overlay if present to form the isolated input view.

// Inject execution context helpers used by nodes.

// Only inject node-level callbacks if configured to avoid overwriting
// state-level callbacks with nil.

// emitNodeErrorEvent emits the node error event.
func (e *Executor) emitNodeErrorEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	nodeID string,
	nodeType NodeType,
	step int,
	err error,
	extra ...NodeEventOption,
) {
	_ = "STUB: not implemented"
	return
}

// handleNodeResult handles the result from node execution.
func (e *Executor) handleNodeResult(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	result any,
	step int,
) (bool, error) {
	_ = "STUB: not implemented"
	// Even if result is nil, static edge writes should still occur so that
	// downstream nodes can be triggered. Only skip static writes when we
	// have explicit routing (Command with GoTo) or fan-out ([]*Command).
	return false, nil
}

// Handle node result by concrete type.

// State update.

// Single command.

// Resolve GoTo via per-node ends if provided.

// If the command explicitly routes via GoTo, avoid also writing to
// channels from static edges for this task to prevent double-triggering
// the downstream node (once via GoTo, once via edge writes).

// Fan-out commands.
// Fan-out: enqueue tasks with overlays.

// Resolve per-node ends for each command before enqueue.

// Process channel writes when not explicitly routed. This ensures that
// nodes with nil results (e.g., pure routing/start nodes) still trigger
// their outgoing static edges.

func (e *Executor) ensureTraceSourceForTask(
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	result any,
	stepErr error,
	traceTask *traceTaskMetadata,
) {
	_ = "STUB: not implemented"
	return
}

// resolveTargetByEnds resolves a symbolic target name using the node's per-node
// ends mapping. If no mapping is found, returns an empty string.
func (e *Executor) resolveTargetByEnds(fromNodeID, target string) string {
	_ = "STUB: not implemented"
	return ""
}

// enqueueCommands enqueues a set of commands as pending tasks for subsequent steps.
func (e *Executor) enqueueCommands(
	execCtx *ExecutionContext,
	t *Task,
	cmds []*Command,
	step int,
	predecessors []string,
) {
	_ = "STUB: not implemented"
	return
}

// Command fan-out tasks are scheduled for the next planning cycle.
// Preserve planning-cycle uniqueness with the next BSP step number and a per-command suffix.

// Get a copy of the current global state to merge with each command

// Merge global state with command-specific overlay

// Resolve writers/triggers from the target node rather than the source task.

// Create task with merged state and target node channel config.

// updateStateFromResult updates the execution context state from a State result.
func (e *Executor) updateStateFromResult(execCtx *ExecutionContext, stateResult State) {
	_ = "STUB: not implemented"
	return
}

// Sanitize: drop internal/ephemeral keys from user node updates.
// These keys (e.g., exec_context) are maintained by the executor and
// may contain concurrently-mutated maps. Accepting them causes
// reflective deep copies to iterate maps while other goroutines write.
// That leads to "concurrent map iteration and map write" panics.

// Use schema-based reducers when available for proper merging.

// Fallback to direct assignment if no schema available.

// syncResumeState copies resume-related keys from a node-local state view into the shared executor state.
func (e *Executor) syncResumeState(execCtx *ExecutionContext, source State) {
	_ = "STUB: not implemented"
	return
}

// syncResumeKey applies a specific resume key mutation from the node state.
func syncResumeKey(target, source State, key string) { _ = "STUB: not implemented"; return }

// handleCommandResult handles a Command result from node execution.
func (e *Executor) handleCommandResult(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	cmdResult *Command,
	step int,
	taskID string,
) error {
	_ = "STUB: not implemented"
	// Update state with command updates.
	return nil
}

// Handle GoTo routing.

// handleCommandRouting handles the routing specified by a Command.
func (e *Executor) handleCommandRouting(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	taskID string,
	sourceStepIDs []string,
	targetNode string,
	step int,
) {
	_ = "STUB: not implemented"
	// Create trigger channel for the target node (including self).
	return
}

// Ensure the per-execution channel exists and write to it.

// Emit channel update event.

// processChannelWrites processes the channel writes for a task.
func (e *Executor) processChannelWrites(ctx context.Context, invocation *agent.Invocation,
	execCtx *ExecutionContext, taskID string, writes []channelWriteEntry, step int) {
	_ = "STUB: not implemented"
	return
}

// Emit channel update event.

// Accumulate into pendingWrites to be saved with the next checkpoint.

// Use atomic increment for deterministic replay

// restoreCheckpointValueWithSchema restores a checkpoint value to its proper type using schema information.
func (e *Executor) restoreCheckpointValueWithSchema(value any, field StateField) any {
	_ = "STUB: not implemented"
	// Skip if already the correct type.
	return *new(any)
}

// Approach 1: Use Default as template if available.

// Use a pointer to the template for unmarshaling.

// Approach 2: Use reflection to create correct type.

// Fallback: return value as-is.

// emitChannelUpdateEvent emits a channel update event.
func (e *Executor) emitChannelUpdateEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	channelName string,
	channelType channel.Behavior,
	triggeredNodes []string,
) {
	_ = "STUB: not implemented"
	return
}

// emitNodeCompleteEvent emits the node completion event.
func (e *Executor) emitNodeCompleteEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	nodeID string,
	nodeType NodeType,
	step int,
	startTime time.Time,
	cacheHit bool,
) {
	_ = "STUB: not implemented"
	return
}

// Attach cache-hit metadata if supported by the event schema.
// We piggyback on node metadata extension by encoding a boolean marker inside
// the existing Node metadata via StateDelta in NewNodeCompleteEvent.
// Here we cannot mutate the event payload directly, so we re-emit a separate
// event that carries cache info is not strictly necessary. As a lightweight
// approach, when cacheHit=true we append a synthetic key in OutputKeys to
// aid debugging without breaking compatibility.

// Best-effort hint: add a virtual output key for observability.

// getEffectiveCachePolicy returns the node-level cache policy if set, otherwise the graph-level policy.
func (e *Executor) getEffectiveCachePolicy(nodeID string) *CachePolicy {
	_ = "STUB: not implemented"
	return nil
}

// updateChannels processes channel updates and emits events.
func (e *Executor) updateChannels(ctx context.Context, invocation *agent.Invocation,
	execCtx *ExecutionContext, step int) error {
	_ = "STUB: not implemented"
	return nil
}

// emitUpdateStepEvent emits the update step event.
func (e *Executor) emitUpdateStepEvent(ctx context.Context, invocation *agent.Invocation, execCtx *ExecutionContext, step int) {
	_ = "STUB: not implemented"
	return
}

// emitStateUpdateEvent emits the state update event.
func (e *Executor) emitStateUpdateEvent(ctx context.Context, invocation *agent.Invocation, execCtx *ExecutionContext) {
	_ = "STUB: not implemented"
	return
}

// getUpdatedChannels returns a list of updated channel names for this execution.
func (e *Executor) getUpdatedChannels(execCtx *ExecutionContext) []string {
	_ = "STUB: not implemented"
	return nil
}

// selectRetryPolicy selects the first matching retry policy for the given error.
// Returns whether a match was found, the chosen policy, and its MaxAttempts
// (falling back to 1 when unspecified or invalid).
func (e *Executor) selectRetryPolicy(err error, policies []RetryPolicy) (bool, RetryPolicy, int) {
	_ = "STUB: not implemented"
	return false, *new(RetryPolicy), 0
}

// getUpdatedChannelsInStep returns a list of channels updated in the current step.
func (e *Executor) getUpdatedChannelsInStep(execCtx *ExecutionContext, step int) []string {
	_ = "STUB: not implemented"
	return nil
}

// getTriggeredNodes returns the list of nodes triggered by a channel.
func (e *Executor) getTriggeredNodes(channelName string) []string {
	_ = "STUB: not implemented"
	return nil
}

// processConditionalEdges evaluates conditional edges for a node and creates dynamic channels.
func (e *Executor) processConditionalEdges(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	t *Task,
	step int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Evaluate the conditional function.

// Deduplicate results to avoid double triggers.

// Skip empty branch keys; they are treated as no-op.

// processConditionalResult processes the result of a conditional edge evaluation.
func (e *Executor) processConditionalResult(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	condEdge *ConditionalEdge,
	result string,
	step int,
	sourceTaskID string,
) error {
	_ = "STUB: not implemented"
	// Determine target by precedence:
	// 1) explicit PathMap mapping
	// 2) node-level ends mapping (symbolic -> concrete)
	// 3) treat result as a concrete node id
	// First, check explicit PathMap mapping.
	return nil
}

// Then resolve by node-level ends mapping (symbolic -> concrete).

// Finally, fallback to treating the result as a concrete node id.

// Create and trigger the target channel.

// Trigger the target by writing to the channel.

// handleInterrupt handles an interrupt during graph execution.
func (e *Executor) handleInterrupt(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	interrupt *InterruptError,
	step int,
	checkpointConfig map[string]any,
	metaExtra map[string]any,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create an interrupt checkpoint with the current state.

// Set interrupt state in the checkpoint.

// IMPORTANT: Set parent checkpoint ID from current config to maintain
// proper tree structure.

// Setting parent checkpoint ID for interrupt

// Create metadata for the interrupt checkpoint.

// Set next nodes for recovery
// IMPORTANT:
// - For internal interrupts (from graph.Interrupt within a node), the
//   interrupted node needs to be re-executed to complete its work, so we
//   include it in NextNodes.
// - For static interrupts before a step executes, channel-based frontier
//   discovery is unavailable; callers may provide NextNodes explicitly.

// Store the interrupt checkpoint using PutFull for consistency
// Use a new context to ensure checkpoint saves even if main context is canceled.
// Use configured timeout, fallback to default if not set.

// Update the config with new checkpoint ID for proper parent tracking

// Replace ctx with a fresh eventCtx derived from background to avoid cancel warning.

// Emit interrupt event.

// Return the interrupt error to propagate it to the caller.

func (e *Executor) maybeEmitCheckpointInterruptEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	checkpointID string,
	step int,
	duration time.Duration,
	ok bool,
) {
	_ = "STUB: not implemented"
	return
}

// createCheckpointFromState creates a checkpoint from the current execution state.
func (e *Executor) createCheckpointFromState(state State, step int, execCtx *ExecutionContext) *Checkpoint {
	_ = "STUB: not implemented"
	// Convert state to channel values, ensuring we capture the latest state
	// including any updates from nodes that haven't been written to channels yet.
	// No deep copy is required here
	return nil
}

// Create versions seen from execution context.

// Create checkpoint.

// Use step-specific channels if step is provided, otherwise fallback to all available

func (e *Executor) collectChannelVersions(
	execCtx *ExecutionContext,
) map[string]int64 {
	_ = "STUB: not implemented"
	return nil
}

// getNextNodes determines which nodes should be executed next based on the current state.
func (e *Executor) getNextNodes(execCtx *ExecutionContext) []string {
	_ = "STUB: not implemented"
	return nil
}

// Check for nodes that are ready to execute based on channel triggers

// Remove duplicates

// getNextChannels determines which channels should be triggered next.
func (e *Executor) getNextChannels(execCtx *ExecutionContext) []string {
	_ = "STUB: not implemented"
	return nil
}

// getNextChannelsInStep determines which channels were updated in the current step.
func (e *Executor) getNextChannelsInStep(execCtx *ExecutionContext, step int) []string {
	_ = "STUB: not implemented"
	return nil
}

// clearChannelStepMarks clears the step marks for all channels after checkpoint creation.
func (e *Executor) clearChannelStepMarks(execCtx *ExecutionContext) {
	_ = "STUB: not implemented"
	return
}

// CheckpointManager returns the executor's checkpoint manager.
// Returns nil if no checkpoint saver was configured.
func (e *Executor) CheckpointManager() *CheckpointManager { _ = "STUB: not implemented"; return nil }

// updateVersionsSeen updates the versions seen by a node after task execution.
func (e *Executor) updateVersionsSeen(execCtx *ExecutionContext, nodeID string, triggers []string) {
	_ = "STUB: not implemented"
	return
}

// Initialize map for node if needed.

// Record current version of all trigger channels this node has seen.

// shouldTriggerNode checks if a node should be triggered based on version tracking.
func (e *Executor) shouldTriggerNode(
	nodeID string,
	channelName string,
	currentVersion int64,
	lastCheckpoint *Checkpoint,
) bool {
	_ = "STUB: not implemented"
	return false
}

// No checkpoint or no version tracking - should trigger.

// Get what this node has seen before.

// Node has never run - should trigger.

// Check if node has seen this channel version.

// Node hasn't seen this channel before - should trigger.

// Only trigger if channel has newer version than what node has seen.

// Fork creates a new branch from an existing checkpoint within the same lineage.
// This allows exploring alternative execution paths from any checkpoint.
func (e *Executor) Fork(ctx context.Context, config map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the source checkpoint.

// Fork the checkpoint (creates new ID and sets parent).

// Create metadata for the fork.

// Save the forked checkpoint with same lineage_id.

// Copy pending writes from the source to ensure resumed execution can continue.
// If the source has pending writes, we need to preserve them in the fork.

// Use PutFull to save both checkpoint and pending writes atomically.
