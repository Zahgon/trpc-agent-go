//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates comprehensive interrupt and resume functionality
// using the graph package. It shows how to create a graph-based agent that
// can be interrupted at specific points and resumed with user input, using
// Runner for orchestration and GraphAgent for execution.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver (install with: go get github.com/mattn/go-sqlite3)

	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	agentlog "trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	// Default configuration values.
	defaultModelName     = "deepseek-v4-flash"
	defaultUserID        = "interrupt-user"
	defaultAppName       = "interrupt-workflow"
	defaultDBPath        = "interrupt-checkpoints.db"
	defaultLineagePrefix = "interrupt-demo"

	// State keys for the workflow.
	stateKeyCounter   = "counter"
	stateKeyMessages  = "messages"
	stateKeyUserInput = "user_input"
	stateKeyApproved  = "approved"
	stateKeyStepCount = "step_count"
	stateKeyLastNode  = "last_node"

	// Node names.
	nodeIncrement       = "increment"
	nodeRequestApproval = "request_approval"
	nodeSecondApproval  = "second_approval"
	nodeProcessApproval = "process_approval"
	nodeFinalize        = "finalize"

	// Messages.
	msgApprovalRequest   = "Please approve the current state (yes/no):"
	msgSecondApproval    = "This requires a second approval (yes/no):"
	msgUserApproved      = "user approved: %t"
	msgUserRejected      = "user rejected - stopping execution"
	msgUserApprovedCont  = "user approved - continuing execution"
	msgExecutionComplete = "execution completed successfully"
	msgNodeExecuted      = "Node %s executed at %s"
	msgWorkflowComplete  = "Workflow completed with counter: %d"

	// Commands.
	cmdRun       = "run"
	cmdInterrupt = "interrupt"
	cmdResume    = "resume"
	cmdList      = "list"
	cmdTree      = "tree"
	cmdHistory   = "history"
	cmdLatest    = "latest"
	cmdDelete    = "delete"
	cmdStatus    = "status"
	cmdDemo      = "demo"
	cmdHelp      = "help"
	cmdExit      = "exit"
	cmdQuit      = "quit"
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
	interactiveMode = flag.Bool("interactive", true,
		"Enable interactive command-line mode")
	lineageFlag = flag.String("lineage", "",
		"Lineage ID for checkpointing (default: auto-generated)")
)

// interruptWorkflow manages a workflow with comprehensive interrupt and resume support.
type interruptWorkflow struct {
	modelName        string
	storageType      string
	dbPath           string
	redisClientURL   string
	verbose          bool
	logger           agentlog.Logger
	runner           runner.Runner
	graphAgent       *graphagent.GraphAgent
	saver            graph.CheckpointSaver
	manager          *graph.CheckpointManager
	currentLineageID string
	currentNamespace string
	userID           string
	sessionID        string
}

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Printf("🔄 Advanced Interrupt & Resume Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Storage: %s", *storage)
	if *storage == "sqlite" {
		fmt.Printf(" (DB: %s)", *dbPath)
	}
	fmt.Println()
	fmt.Printf("Verbose Mode: %v\n", *verbose)
	fmt.Printf("Interactive Mode: %v\n", *interactiveMode)
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the workflow.
	workflow := &interruptWorkflow{
		modelName:   *modelName,
		storageType: *storage,
		dbPath:      *dbPath,
		verbose:     *verbose,
		userID:      defaultUserID,
	}

	// Setup lineage ID.
	if *lineageFlag != "" {
		workflow.currentLineageID = *lineageFlag
	} else {
		workflow.currentLineageID = fmt.Sprintf("%s-%d", defaultLineagePrefix,
			time.Now().Unix())
	}
	workflow.sessionID = fmt.Sprintf("session-%d", time.Now().Unix())
	workflow.currentNamespace = "interrupt-demo"

	if err := workflow.run(); err != nil {
		fmt.Printf("❌ Workflow failed: %v\n", err)
		os.Exit(1)
	}
}

// run starts the interrupt workflow.
func (w *interruptWorkflow) run() error { _ = "STUB: not implemented"; return nil }

// Setup the workflow components.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive mode.

// Single command execution.

// setup creates the graph agent with interrupt support and runner.
func (w *interruptWorkflow) setup() error {
	_ = "STUB: not implemented"
	// Initialize logger.
	return nil
}

// Create checkpoint saver based on storage type.

// Create the workflow graph.

// Create GraphAgent with checkpoint support.

// Get the checkpoint manager from the executor.

// Create session service.

// Create Runner.

// createInterruptGraph creates the interrupt-capable workflow graph.
func (w *interruptWorkflow) createInterruptGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Define schema with comprehensive state tracking.
	return nil, nil
}

// Build graph.

// Node 1: Increment counter and track execution.

// Increment by 10 to make it more visible

// Node 2: Request user approval (first interrupt point).

// Check if we should skip interrupts (for run command).

// Create interrupt payload with rich context.

// Interrupt will check for resume values automatically.
// If there's a resume value, it returns it without interrupting.
// If not, it creates an interrupt.
// Use the node ID as the interrupt key to match what executor sets as TaskID

// Process resume value.

// When skipping interrupts, auto-approve.

// Node 3: Second approval for complex workflow (optional interrupt).

// Skip second approval if first was rejected.

// Check if we should skip interrupts (for run command).

// Create second interrupt for approved workflows.

// Second interrupt point.
// Interrupt will check for resume values automatically.

// Use the node ID as the interrupt key to match what executor sets as TaskID

// Process second approval.

// When skipping interrupts, auto-approve second approval.

// Node 4: Process final approval decision.

// Node 5: Finalize workflow.

// Setup workflow edges.

// startInteractiveMode starts the interactive command-line interface.
func (w *interruptWorkflow) startInteractiveMode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse command and arguments.

// Could be checkpoint-id or user input

// runWorkflow executes the workflow with the given lineage ID.
func (w *interruptWorkflow) runWorkflow(ctx context.Context, lineageID string, waitForInterrupt bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Create initial message.

// Create checkpoint config for this lineage.

// Run the workflow through the runner.
// Pass lineage_id and checkpoint namespace to enable checkpoint saving.
// Also pass skip_interrupts flag to control interrupt behavior.

// Process events and track execution.

// Show node execution progress for interrupt workflows.
// Node events have Author=<node-name> and Object="graph.node.start"/"graph.node.complete".

// Track the last node for completion messages.

// Check if an interrupt checkpoint was created by examining actual checkpoints.
// This is more reliable than trying to parse event streams.

// Look for interrupt checkpoint.

// resumeWorkflow resumes execution from a checkpoint.
func (w *interruptWorkflow) resumeWorkflow(ctx context.Context, lineageID, checkpointID, userInput string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the lineage exists before attempting resume.

// Prompt for user input if not provided.

// Create resume command. Only set the resume value for the specific interrupt
// that is currently active. We need to check which interrupt is active.

// Check which interrupt is currently active by looking at the latest checkpoint.
// This leverages the ResumeMap design - we only set the resume value for the
// specific interrupt that's currently active, not future ones.
// First, list all checkpoints to find the interrupted one

// Find the latest interrupted checkpoint

// Fallback to using Latest() if no interrupted checkpoint found in list

// Use the TaskID from the interrupt state as the key.
// This automatically handles any interrupt without needing to know specific names.

// Add checkpoint ID if specified.

// Process events.

// Track node execution for verbose output.

// Track execution.

// Check if the workflow completed or was interrupted again.

// Check the latest checkpoint

// If the latest checkpoint is a regular checkpoint (not interrupt),
// and it's from a step after the resume, the workflow likely completed or progressed

// Check if we reached the final node

// There's a new interrupt

// listCheckpoints lists available checkpoints with interrupt-specific details.
func (w *interruptWorkflow) listCheckpoints(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle nil namespace properly

// Use current time as placeholder since CreatedAt may not be available

// Display state summary.

// Highlight interrupts with detailed information.

// showLatestCheckpoint displays detailed information about the latest checkpoint.
func (w *interruptWorkflow) showLatestCheckpoint(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Look for interrupt checkpoints first, then fallback to latest.

// Prefer interrupt checkpoint if available, otherwise use the first one (most recent).

// Use first (most recent) checkpoint

// Handle nil namespace properly

// Use current time as placeholder

// Display comprehensive state.

// Show last 3 messages

// Display interrupt details if present.

// showTree displays checkpoint tree structure with interrupt indicators.
func (w *interruptWorkflow) showTree(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Build parent-child relationships.

// Display tree structure.

// printCheckpointTree recursively prints the checkpoint tree structure.
func (w *interruptWorkflow) printCheckpointTreeNode(checkpoint *graph.CheckpointTuple, prefix string, connector string, checkpointMap map[string]*graph.CheckpointTuple, childMap map[string][]*graph.CheckpointTuple) {
	_ = "STUB: not implemented"
	// Extract state information.
	return
}

// Create display info.

// Use current time as placeholder

// Choose icon based on interrupt status.

// Print the connector and node on the same line

// Print children recursively

// Determine connector and continuation prefix

// Recursively print child

// printCheckpointTree is the entry point for tree printing.
func (w *interruptWorkflow) printCheckpointTree(checkpoint *graph.CheckpointTuple, prefix string, checkpointMap map[string]*graph.CheckpointTuple, childMap map[string][]*graph.CheckpointTuple, isRoot bool) {
	_ = "STUB: not implemented"
	return
}

// showHistory displays execution history with interrupt markers.
func (w *interruptWorkflow) showHistory(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Sort by step.

// Use current time as placeholder

// Extract state.

// Show recent messages.

// Highlight interrupts.

// showInterruptStatus shows current interrupt status (interrupt-specific command).
func (w *interruptWorkflow) showInterruptStatus(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Look for interrupt checkpoints specifically.

// Find the interrupt checkpoint.

// Use current time as placeholder

// No interrupt checkpoint found, get latest checkpoint for completed status.

// Use current time as placeholder

// Show final state.

// deleteLineage deletes all checkpoints for a lineage.
func (w *interruptWorkflow) deleteLineage(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// For non-interactive usage (like tests), skip confirmation.
// In interactive usage, we could add confirmation, but for simplicity
// and to make tests work, we'll proceed directly.

// Auto-confirm for testing purposes

// runDemo runs a comprehensive demonstration of interrupt features.
func (w *interruptWorkflow) runDemo(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Step 1: Run until first interrupt.

// Step 2: Show interrupt status.

// Step 3: List checkpoints with details.

// Step 4: Resume with approval to hit second interrupt.

// Step 5: Show updated status.

// Step 6: Resume and complete.

// Step 7: Show final tree structure.

// Step 8: Show execution history.

// showHelp displays available commands with comprehensive interrupt documentation.
func (w *interruptWorkflow) showHelp() { _ = "STUB: not implemented"; return }

// Helper functions.

func getInt(s graph.State, key string) int { _ = "STUB: not implemented"; return 0 }

func getBool(s graph.State, key string) bool { _ = "STUB: not implemented"; return false }

func getStrs(s graph.State, key string) []string { _ = "STUB: not implemented"; return nil }

func generateLineageID() string { _ = "STUB: not implemented"; return "" }

// extractRootState extracts the root state from a checkpoint.
func (w *interruptWorkflow) extractRootState(checkpoint *graph.Checkpoint) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
