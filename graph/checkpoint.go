//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.

// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"context"
	"time"
)

const (
	// CheckpointVersion is the current version of the checkpoint format.
	CheckpointVersion = 1

	// CheckpointSourceInput indicates the checkpoint was created from input.
	CheckpointSourceInput = "input"
	// CheckpointSourceLoop indicates the checkpoint was created from inside the loop.
	CheckpointSourceLoop = "loop"
	// CheckpointSourceUpdate indicates the checkpoint was created from manual update.
	CheckpointSourceUpdate = "update"
	// CheckpointSourceFork indicates the checkpoint was created as a copy.
	CheckpointSourceFork = "fork"
	// CheckpointSourceInterrupt indicates the checkpoint was created from an interrupt.
	CheckpointSourceInterrupt = "interrupt"

	// DefaultCheckpointNamespace is the default namespace for checkpoints.
	DefaultCheckpointNamespace = ""
	// DefaultChannelVersion is the default version for channels.
	DefaultChannelVersion = 1
	// DefaultMaxCheckpointsPerLineage is the default maximum number of checkpoints per lineage.
	DefaultMaxCheckpointsPerLineage = 100
)

// Special channel names for interrupt and resume functionality.
const (
	InterruptChannel = "__interrupt__"
	ResumeChannel    = "__resume__"
	ErrorChannel     = "__error__"
	ScheduledChannel = "__scheduled__"
)

// Checkpoint represents a snapshot of graph state at a specific point in time.
type Checkpoint struct {
	// Version is the version of the checkpoint format.
	Version int `json:"v"`
	// ID is the unique identifier for this checkpoint.
	ID string `json:"id"`
	// Timestamp is when the checkpoint was created.
	Timestamp time.Time `json:"ts"`
	// ChannelValues contains the values of channels at checkpoint time.
	ChannelValues map[string]any `json:"channel_values"`
	// ChannelVersions contains the versions of channels at checkpoint time.
	ChannelVersions map[string]int64 `json:"channel_versions"`
	// BarrierSets stores per-channel barrier seen sets.
	BarrierSets map[string][]string `json:"barrier_sets,omitempty"`
	// VersionsSeen tracks which versions each node has seen.
	VersionsSeen map[string]map[string]int64 `json:"versions_seen"`
	// ParentCheckpointID is the ID of the parent checkpoint (for branching).
	ParentCheckpointID string `json:"parent_checkpoint_id,omitempty"`
	// UpdatedChannels lists channels updated in the step that produced this
	// checkpoint.
	UpdatedChannels []string `json:"updated_channels,omitempty"`
	// PendingSends contains messages that haven't been sent yet.
	PendingSends []PendingSend `json:"pending_sends,omitempty"`
	// InterruptState contains information about the current interrupt state.
	InterruptState *InterruptState `json:"interrupt_state,omitempty"`
	// NextNodes contains the next nodes to execute (alternative to pendingWrites).
	NextNodes []string `json:"next_nodes,omitempty"`
	// NextChannels contains the next channels to trigger (alternative to pendingWrites).
	NextChannels []string `json:"next_channels,omitempty"`
}

// InterruptState represents the state of an interrupted execution.
type InterruptState struct {
	// NodeID is the ID of the node where execution was interrupted.
	NodeID string `json:"node_id"`
	// TaskID is the ID of the task that was interrupted.
	TaskID string `json:"task_id"`
	// InterruptValue is the value that was passed to interrupt().
	InterruptValue any `json:"interrupt_value"`
	// ResumeValues contains values to resume execution with.
	ResumeValues []any `json:"resume_values,omitempty"`
	// Step is the step number when the interrupt occurred.
	Step int `json:"step"`
	// Path is the execution path to the interrupted node.
	Path []string `json:"path,omitempty"`
}

// PendingSend represents a message that hasn't been sent yet.
type PendingSend struct {
	// Channel is the channel to send to.
	Channel string `json:"channel"`
	// Value is the value to send.
	Value any `json:"value"`
	// TaskID is the ID of the task that created this send.
	TaskID string `json:"task_id,omitempty"`
}

// CheckpointMetadata contains metadata about a checkpoint.
type CheckpointMetadata struct {
	// Source indicates how the checkpoint was created.
	Source string `json:"source"`
	// Step is the step number (-1 for input, 0+ for loop steps).
	Step int `json:"step"`
	// Parents maps checkpoint namespaces to parent checkpoint IDs.
	Parents map[string]string `json:"parents"`
	// Additional metadata fields.
	Extra map[string]any `json:"extra,omitempty"`
	// IsResuming indicates if this checkpoint is being resumed from.
	IsResuming bool `json:"is_resuming,omitempty"`
}

// CheckpointTuple wraps a checkpoint with its configuration and metadata.
type CheckpointTuple struct {
	// Config contains the configuration used to create this checkpoint.
	Config map[string]any `json:"config"`
	// Checkpoint is the actual checkpoint data.
	Checkpoint *Checkpoint `json:"checkpoint"`
	// Metadata contains additional checkpoint information.
	Metadata *CheckpointMetadata `json:"metadata"`
	// ParentConfig is the configuration of the parent checkpoint.
	ParentConfig map[string]any `json:"parent_config,omitempty"`
	// PendingWrites contains writes that haven't been committed yet.
	PendingWrites []PendingWrite `json:"pending_writes,omitempty"`
}

// PendingWrite represents a write operation that hasn't been committed.
type PendingWrite struct {
	// TaskID is the ID of the task that created this write.
	TaskID string `json:"task_id"`
	// Channel is the channel being written to.
	Channel string `json:"channel"`
	// Value is the value being written.
	Value any `json:"value"`
	// Sequence is the global sequence number for deterministic replay.
	Sequence int64 `json:"sequence"`
}

// PutRequest contains all data needed to store a checkpoint.
type PutRequest struct {
	Config      map[string]any
	Checkpoint  *Checkpoint
	Metadata    *CheckpointMetadata
	NewVersions map[string]int64
}

// PutWritesRequest contains all data needed to store writes.
type PutWritesRequest struct {
	Config   map[string]any
	Writes   []PendingWrite
	TaskID   string
	TaskPath string
}

// PutFullRequest contains all data needed to atomically store a checkpoint with its writes.
type PutFullRequest struct {
	Config        map[string]any
	Checkpoint    *Checkpoint
	Metadata      *CheckpointMetadata
	NewVersions   map[string]int64
	PendingWrites []PendingWrite
}

// CheckpointSaver defines the interface for checkpoint storage implementations.
type CheckpointSaver interface {
	// Get retrieves a checkpoint by configuration.
	Get(ctx context.Context, config map[string]any) (*Checkpoint, error)
	// GetTuple retrieves a checkpoint tuple by configuration.
	GetTuple(ctx context.Context, config map[string]any) (*CheckpointTuple, error)
	// List retrieves checkpoints matching criteria.
	List(ctx context.Context, config map[string]any, filter *CheckpointFilter) ([]*CheckpointTuple, error)
	// Put stores a checkpoint.
	Put(ctx context.Context, req PutRequest) (map[string]any, error)
	// PutWrites stores intermediate writes linked to a checkpoint.
	PutWrites(ctx context.Context, req PutWritesRequest) error
	// PutFull atomically stores a checkpoint with its pending writes in a single transaction.
	PutFull(ctx context.Context, req PutFullRequest) (map[string]any, error)
	// DeleteLineage removes all checkpoints for a lineage.
	DeleteLineage(ctx context.Context, lineageID string) error
	// Close releases resources held by the saver.
	Close() error
}

// CheckpointTree represents the tree structure of checkpoints in a lineage.
type CheckpointTree struct {
	// Root is the root node of the tree.
	Root *CheckpointNode `json:"root"`
	// Branches maps checkpoint IDs to their nodes for quick access.
	Branches map[string]*CheckpointNode `json:"branches"`
}

// CheckpointNode represents a node in the checkpoint tree.
type CheckpointNode struct {
	// Checkpoint is the checkpoint tuple at this node.
	Checkpoint *CheckpointTuple `json:"checkpoint"`
	// Children are the child nodes (forks from this checkpoint).
	Children []*CheckpointNode `json:"children"`
	// Parent is the parent node (null for root).
	Parent *CheckpointNode `json:"-"` // Avoid circular JSON.
}

// CheckpointFilter defines filtering criteria for listing checkpoints.
type CheckpointFilter struct {
	// Before limits results to checkpoints created before this config.
	Before map[string]any `json:"before,omitempty"`
	// Limit is the maximum number of checkpoints to return.
	Limit int `json:"limit,omitempty"`
	// Metadata filters checkpoints by metadata fields.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CheckpointConfig provides a structured way to handle checkpoint configuration.
type CheckpointConfig struct {
	// LineageID is the unique identifier for the conversation lineage.
	LineageID string
	// CheckpointID is the specific checkpoint to retrieve.
	CheckpointID string
	// Namespace is the checkpoint namespace.
	Namespace string
	// ResumeMap maps task namespaces to resume values.
	ResumeMap map[string]any
	// Extra contains additional configuration fields.
	Extra map[string]any
}

// NewCheckpoint creates a new checkpoint with the given data.
func NewCheckpoint(
	channelValues map[string]any,
	channelVersions map[string]int64,
	versionsSeen map[string]map[string]int64,
) *Checkpoint {
	_ = "STUB: not implemented"
	return nil
}

// NewCheckpointMetadata creates new checkpoint metadata.
func NewCheckpointMetadata(source string, step int) *CheckpointMetadata {
	_ = "STUB: not implemented"
	return nil
}

// NewCheckpointConfig creates a new checkpoint configuration.
func NewCheckpointConfig(lineageID string) *CheckpointConfig { _ = "STUB: not implemented"; return nil }

// Use default empty namespace to align with LangGraph's design.

// WithCheckpointID sets the checkpoint ID.
func (c *CheckpointConfig) WithCheckpointID(checkpointID string) *CheckpointConfig {
	_ = "STUB: not implemented"
	return nil
}

// WithNamespace sets the namespace.
func (c *CheckpointConfig) WithNamespace(namespace string) *CheckpointConfig {
	_ = "STUB: not implemented"
	return nil
}

// WithResumeMap sets the resume map.
func (c *CheckpointConfig) WithResumeMap(resumeMap map[string]any) *CheckpointConfig {
	_ = "STUB: not implemented"
	return nil
}

// WithExtra sets additional configuration.
func (c *CheckpointConfig) WithExtra(key string, value any) *CheckpointConfig {
	_ = "STUB: not implemented"
	return nil
}

// ToMap converts the config to a map for backward compatibility.
func (c *CheckpointConfig) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

// Always include namespace to ensure consistency (even if empty).

// Add extra fields.

// NewCheckpointFilter creates a new checkpoint filter.
func NewCheckpointFilter() *CheckpointFilter { _ = "STUB: not implemented"; return nil }

// WithBefore sets the before filter.
func (f *CheckpointFilter) WithBefore(before map[string]any) *CheckpointFilter {
	_ = "STUB: not implemented"
	return nil
}

// WithLimit sets the limit.
func (f *CheckpointFilter) WithLimit(limit int) *CheckpointFilter {
	_ = "STUB: not implemented"
	return nil

	// WithMetadata sets metadata filter.
}

func (f *CheckpointFilter) WithMetadata(key string, value any) *CheckpointFilter {
	_ = "STUB: not implemented"
	return nil
}

// Copy creates a deep copy of the checkpoint.
func (c *Checkpoint) Copy() *Checkpoint { _ = "STUB: not implemented"; return nil }

// Deep copy channel values.

// Deep copy channel versions.

// Deep copy barrier sets.

// Deep copy versions seen.

// Deep copy updated channels.

// Deep copy pending sends.

// Deep copy interrupt state.

// Deep copy next nodes and channels.

// Preserve original ID for true copy.

// Fork creates a copy of the checkpoint with a new ID and sets parent relationship.
// This is used for branching and creating new checkpoints based on existing ones.
func (c *Checkpoint) Fork() *Checkpoint { _ = "STUB: not implemented"; return nil }

// Create a true copy first.

// Set the parent to the current checkpoint's ID.

// Generate a new ID for the forked checkpoint.

// Update timestamp to current time.

// GetCheckpointID extracts checkpoint ID from configuration.
func GetCheckpointID(config map[string]any) string { _ = "STUB: not implemented"; return "" }

// GetLineageID extracts lineage ID from configuration.
func GetLineageID(config map[string]any) string { _ = "STUB: not implemented"; return "" }

// GetNamespace extracts namespace from configuration.
func GetNamespace(config map[string]any) string { _ = "STUB: not implemented"; return "" }

// GetResumeMap extracts resume map from configuration.
func GetResumeMap(config map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// CreateCheckpointConfig creates a checkpoint configuration (legacy function).
func CreateCheckpointConfig(lineageID string, checkpointID string, namespace string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Use default empty namespace to align with LangGraph's design.

// CheckpointManager provides high-level checkpoint management functionality.
type CheckpointManager struct {
	saver CheckpointSaver
}

// NewCheckpointManager creates a new checkpoint manager.
func NewCheckpointManager(saver CheckpointSaver) *CheckpointManager {
	_ = "STUB: not implemented"
	return nil
}

// CreateCheckpoint creates a new checkpoint from the current state.
func (cm *CheckpointManager) CreateCheckpoint(
	ctx context.Context, config map[string]any, state State, source string, step int,
) (*Checkpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert state to channel values with deep copy to prevent races when
// the saver serializes the checkpoint concurrently with node execution.

// Create channel versions (simple incrementing integers for now).

// Create versions seen (simplified for now).

// Create checkpoint.

// Create metadata.

// Store checkpoint.

// ResumeFromCheckpoint resumes execution from a specific checkpoint.
func (cm *CheckpointManager) ResumeFromCheckpoint(
	ctx context.Context, config map[string]any,
) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// Convert channel values back to state.

// ListCheckpoints lists checkpoints for a lineage.
func (cm *CheckpointManager) ListCheckpoints(
	ctx context.Context, config map[string]any, filter *CheckpointFilter,
) ([]*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteLineage removes all checkpoints for a lineage.
func (cm *CheckpointManager) DeleteLineage(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Latest returns the most recent checkpoint for a lineage and namespace.
func (cm *CheckpointManager) Latest(
	ctx context.Context, lineageID, namespace string,
) (*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get retrieves a checkpoint by configuration.
func (cm *CheckpointManager) Get(
	ctx context.Context, config map[string]any,
) (*Checkpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTuple retrieves a checkpoint tuple by configuration.
func (cm *CheckpointManager) GetTuple(
	ctx context.Context, config map[string]any,
) (*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Goto jumps to a specific checkpoint by ID.
func (cm *CheckpointManager) Goto(
	ctx context.Context, lineageID, namespace, checkpointID string,
) (*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put stores a checkpoint.
func (cm *CheckpointManager) Put(
	ctx context.Context, req PutRequest,
) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BranchFrom creates a new checkpoint branch from an existing one within the same lineage.
func (cm *CheckpointManager) BranchFrom(
	ctx context.Context,
	lineageID, namespace, checkpointID, newNamespace string,
) (*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the source checkpoint.

// Create a new checkpoint in the new namespace.

// Fork creates a new ID.

// Determine step from source if available.

// Store the new checkpoint.

// Return the new checkpoint tuple.

// BranchToNewLineage creates a new checkpoint in a different lineage from an existing checkpoint.
func (cm *CheckpointManager) BranchToNewLineage(
	ctx context.Context,
	sourceLineageID, sourceNamespace, sourceCheckpointID string,
	newLineageID, newNamespace string,
) (*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the source checkpoint.
// If namespace is empty and we're getting latest, we need to search across all namespaces.

// When getting latest without a specific checkpoint ID, search all namespaces if namespace is empty.

// Fetch full tuple to get step as well.

// Create a new checkpoint in the new lineage.

// Fork creates a new ID

// Create metadata with source information.

// Store the new checkpoint.

// Return the new checkpoint tuple.

// ResumeFromLatest resumes execution from the latest checkpoint with a resume command.
func (cm *CheckpointManager) ResumeFromLatest(ctx context.Context, lineageID, namespace string, cmd *Command) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

// Get the latest checkpoint.

// Convert channel values back to state.

// Add the resume command.

// IsInterrupted checks if a checkpoint represents an interrupted execution.
func (c *Checkpoint) IsInterrupted() bool { _ = "STUB: not implemented"; return false }

// GetInterruptValue returns the interrupt value if the checkpoint is interrupted.
func (c *Checkpoint) GetInterruptValue() any { _ = "STUB: not implemented"; return *new(any) }

// GetResumeValues returns the resume values for the interrupted execution.
func (c *Checkpoint) GetResumeValues() []any { _ = "STUB: not implemented"; return nil }

// AddResumeValue adds a resume value to the checkpoint.
func (c *Checkpoint) AddResumeValue(value any) { _ = "STUB: not implemented"; return }

// SetInterruptState sets the interrupt state for the checkpoint.
func (c *Checkpoint) SetInterruptState(nodeID, taskID string, interruptValue any, step int, path []string) {
	_ = "STUB: not implemented"
	return
}

// ClearInterruptState clears the interrupt state.
func (c *Checkpoint) ClearInterruptState() { _ = "STUB: not implemented"; return }

// GetCheckpointTree builds the tree structure of checkpoints in a lineage.
func (cm *CheckpointManager) GetCheckpointTree(
	ctx context.Context, lineageID string,
) (*CheckpointTree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get all checkpoints for the lineage.

// Build the tree structure.

// First pass: create all nodes.

// Second pass: establish parent-child relationships.

// Parent not found, treat as root.

// No parent, this is a root node.

// Sort children by timestamp for consistent ordering.

// Find the primary root (oldest without parent).

// ListChildren returns the direct children of a checkpoint.
func (cm *CheckpointManager) ListChildren(
	ctx context.Context, config map[string]any,
) ([]*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the parent checkpoint to find its ID.

// Get all checkpoints in the lineage.

// Filter to find children.

// Sort by timestamp.

// GetParent returns the parent checkpoint of the given checkpoint.
func (cm *CheckpointManager) GetParent(
	ctx context.Context, config map[string]any,
) (*CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the current checkpoint.

// Check if it has a parent.

// No parent.

// Prefer using explicit ParentConfig if available (handles cross-namespace parents).

// Try using provided parent config first.

// If not found (possibly due to namespace mismatch), try cross-namespace search.

// Fallback: use same-namespace lookup.

// deepCopy performs a deep copy using JSON marshaling/unmarshaling for safety.
func deepCopy(src any) any { _ = "STUB: not implemented"; return *new(any) }

// Marshal to JSON.

// If marshaling fails, return the original value.

// Unmarshal to a generic map with number preservation.

// Preserve number types as json.Number.

// If unmarshaling fails, return the original value.

// deepCopyMap performs a deep copy of a map[string]any.
func deepCopyMap(src map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// Fallback: create a new map and copy values.

// deepCopyStringSlice performs a deep copy of a []string.
func deepCopyStringSlice(src []string) []string { _ = "STUB: not implemented"; return nil }
