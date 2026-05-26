//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.

// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates comprehensive checkpoint functionality
// using the graph package. This example shows how to save, restore,
// and manage execution checkpoints in a graph-based workflow, enabling
// workflow resumption, time-travel debugging, and fault tolerance.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3" // SQLite driver (install with: go get github.com/mattn/go-sqlite3)

	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	agentlog "trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	// Default configuration values.
	defaultModelName = "deepseek-v4-flash"
	defaultUserID    = "demo-user"
	defaultAppName   = "checkpoint-workflow"
	defaultDBPath    = "checkpoints.db"

	// State keys for the workflow.
	stateKeyCounter    = "counter"
	stateKeyMessages   = "messages"
	stateKeyStepCount  = "step_count"
	stateKeyLastAction = "last_action"

	// Node names.
	nodeIncrement1 = "increment1"
	nodeIncrement2 = "increment2"
	nodeIncrement3 = "increment3"
	nodeFinal      = "final"

	// Messages.
	msgNodeExecuted     = "Node %s executed at %s"
	msgWorkflowComplete = "Workflow completed with counter: %d"

	// Commands.
	cmdRun     = "run"
	cmdList    = "list"
	cmdResume  = "resume"
	cmdBranch  = "branch"
	cmdTree    = "tree"
	cmdDelete  = "delete"
	cmdHistory = "history"
	cmdLatest  = "latest"
	cmdDemo    = "demo"
	cmdHelp    = "help"
	cmdExit    = "exit"
	cmdQuit    = "quit"
)

var (
	modelName = flag.String("model", defaultModelName,
		"Name of the model to use")
	storage = flag.String("storage", "memory",
		"Storage type: 'memory' or 'sqlite' or 'redis'")
	dbPath = flag.String("db", defaultDBPath,
		"Path to SQLite database file (only used with -storage=sqlite)")
	redisClientURL = flag.String("redis-url", "redis://localhost:6379",
		"Redis client URL (only used with -storage=redis)")
	verbose = flag.Bool("verbose", false,
		"Enable verbose output")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Printf("🚀 Advanced Checkpoint Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Storage: %s", *storage)
	if *storage == "sqlite" {
		fmt.Printf(" (DB: %s)", *dbPath)
	}
	fmt.Println()
	fmt.Printf("Verbose Mode: %v\n", *verbose)
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the workflow.
	workflow := &checkpointWorkflow{
		modelName:        *modelName,
		storageType:      *storage,
		dbPath:           *dbPath,
		verbose:          *verbose,
		redisClientURL:   *redisClientURL,
		currentNamespace: "checkpoint-demo",
	}
	if err := workflow.run(); err != nil {
		log.Fatalf("Workflow failed: %v", err)
	}
}

// checkpointWorkflow manages a workflow with comprehensive checkpoint support.
type checkpointWorkflow struct {
	modelName        string
	storageType      string
	dbPath           string
	redisClientURL   string
	verbose          bool
	logger           agentlog.Logger
	runner           runner.Runner
	saver            graph.CheckpointSaver
	manager          *graph.CheckpointManager
	graphAgent       *graphagent.GraphAgent
	currentLineageID string
	currentNamespace string
}

// run starts the checkpoint workflow.
func (w *checkpointWorkflow) run() error { _ = "STUB: not implemented"; return nil }

// Setup the workflow components.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive mode.

// setup creates the graph agent with checkpoint support and runner.
func (w *checkpointWorkflow) setup() error {
	_ = "STUB: not implemented"
	// Initialize logger.
	return nil
}

// Create checkpoint saver based on storage type.

// Create the workflow graph.

// Create GraphAgent with checkpoint support.

// Get the checkpoint manager from the executor.

// Create session service.

// Create runner with the graph agent.

// createWorkflowGraph creates a workflow graph with multiple nodes.
func (w *checkpointWorkflow) createWorkflowGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Create state schema with custom fields.
	return nil, nil
}

// Add custom state fields with proper type definitions.

// Explicitly define as int

// Explicitly define as []string

// Explicitly define as int

// Explicitly define as string

// Create the state graph.

// Add workflow nodes.

// Add workflow edges.

// Compile the graph.

// Node implementations.

func (w *checkpointWorkflow) incrementNode1(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (w *checkpointWorkflow) incrementNode2(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (w *checkpointWorkflow) incrementNode3(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (w *checkpointWorkflow) finalNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Create a safe copy of keys to avoid concurrent map iteration

// Log key state values before helper calls

// Helper functions for state access.

func (w *checkpointWorkflow) getCounter(state graph.State) int { _ = "STUB: not implemented"; return 0 }

func (w *checkpointWorkflow) getStepCount(state graph.State) int {
	_ = "STUB: not implemented"
	return 0
}

func (w *checkpointWorkflow) getMessages(state graph.State) []string {
	_ = "STUB: not implemented"
	return nil
}

// Try to handle []any case from JSON deserialization.

// startInteractiveMode starts the interactive command-line interface.
func (w *checkpointWorkflow) startInteractiveMode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse command and arguments.

// Support additional input when resuming (advanced feature)
// Join remaining parts to preserve spaces, trim simple quotes.

// runWorkflow executes the workflow with the given lineage ID.
func (w *checkpointWorkflow) runWorkflow(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create initial message.

// Create checkpoint config for this lineage.

// Run the workflow through the runner.
// Pass lineage_id and checkpoint namespace to enable checkpoint saving.

// Ensure lineage_id is set
// Set checkpoint namespace

// Create safe copy of keys to avoid concurrent map access

// Process streaming response.

// Check if checkpoints were created after execution

// resumeWorkflow resumes a workflow from a checkpoint.
func (w *checkpointWorkflow) resumeWorkflow(
	ctx context.Context, lineageID, checkpointID, additionalInput string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create checkpoint config.

// Ensure namespace matches storage

// Get the checkpoint to see its state.

// If no checkpoint found, just run normally.

// Removed unused source variable

// Update current lineage tracking.

// Create resume message based on input mode.

// Resume with additional input (advanced mode)

// Pure resume with no input (standard mode, following LangGraph pattern)

// Run with the checkpoint config.

// Ensure lineage_id is set
// Set checkpoint namespace

// Add checkpoint_id directly to runtime state for the executor to find it

// Use the actual checkpoint ID from the retrieved checkpoint (checkpoint is guaranteed to be non-nil here)

// The framework now handles type restoration automatically using schema information.
// No manual pre-population is needed.

// Process streaming response.

// branchCheckpoint creates a branch within the same lineage (fork).
func (w *checkpointWorkflow) branchCheckpoint(ctx context.Context, lineageID, checkpointID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the executor.

// Create config for the source checkpoint.

// Fork the checkpoint (keeps same lineage_id).

// Get the branched checkpoint ID.

// showTree displays the checkpoint tree for a lineage.
func (w *checkpointWorkflow) showTree(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the checkpoint manager.

// Get the checkpoint tree.

// Display the tree recursively.

// Display summary.

// Count branches (nodes with children).

// printTreeNode recursively prints a checkpoint tree node.
func (w *checkpointWorkflow) printTreeNode(node *graph.CheckpointNode, prefix string, isLast bool) {
	_ = "STUB: not implemented"
	return
}

// Determine the branch character.

// Get checkpoint info.

// Format the checkpoint display.

// Get counter value if available using extractRootState for proper type handling.

// Handle different number types from JSON deserialization.

// Print the node.

// Update prefix for children.

// Print children.

// listCheckpoints lists all checkpoints for a lineage.
func (w *checkpointWorkflow) listCheckpoints(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create config for the lineage.

// List checkpoints with a filter.

// Display checkpoints.

// Show state summary.

// showLatestCheckpoint displays the latest checkpoint for a lineage.
func (w *checkpointWorkflow) showLatestCheckpoint(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the latest checkpoint.

// Display checkpoint details.

// Show full state.

// Show messages if any.

// showHistory shows the execution history for a lineage.
func (w *checkpointWorkflow) showHistory(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// List all checkpoints.

// Display in chronological order (oldest first).

// Format header based on source.

// Show what happened.

// Show state values.

// Show messages if any.

// Show last 2 messages for context.

// Show checkpoint metadata.

// deleteLineage deletes all checkpoints for a lineage.
func (w *checkpointWorkflow) deleteLineage(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Confirm deletion.

// Delete the lineage.

// Clear current lineage if it was deleted.

// runDemo runs a demonstration sequence.
func (w *checkpointWorkflow) runDemo(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Step 1: Run workflow.

// Step 2: List checkpoints.

// Step 3: Show latest checkpoint.

// Step 4: Create a new workflow run.

// Step 5: Show execution history.

// Step 6: Create a branch within same lineage.

// Get the latest checkpoint ID for branching.

// Use the checkpoint's actual ID.

// Resume from the branched checkpoint.

// Get the latest checkpoint (which should be the branch).

// processStreamingResponse handles the streaming workflow response.
func (w *checkpointWorkflow) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Log all events for debugging.

// Try multiple approaches to detect node execution.

// Approach 1: Look for node execution events directly by checking the object field.

// This is a node completion event - the author is the node ID

// Approach 2: Parse node execution metadata regardless of author.

// Avoid double-counting; rely on explicit graph.node.complete events for counts.

// Track final state updates.

// Store current state for final display.

// Skip metadata keys
// Unmarshal the byte data to get the actual value

// Handle completion.

// extractRootState extracts the root state from a checkpoint.
func (w *checkpointWorkflow) extractRootState(checkpoint *graph.Checkpoint) map[string]any {
	_ = "STUB: not implemented"
	// State is stored directly in ChannelValues
	return nil
}

// Convert to map[string]any

// generateLineageID generates a new lineage ID.
func (w *checkpointWorkflow) generateLineageID() string { _ = "STUB: not implemented"; return "" }

// showHelp displays available commands.
func (w *checkpointWorkflow) showHelp() { _ = "STUB: not implemented"; return }
