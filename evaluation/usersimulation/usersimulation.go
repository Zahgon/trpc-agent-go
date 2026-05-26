//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package usersimulation

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Simulator starts a simulated user conversation for one eval case.
type Simulator interface {
	// Start creates a conversation handle for a single eval case.
	Start(ctx context.Context, req *StartRequest) (Conversation, error)
}

// Conversation advances a simulated user one turn at a time.
type Conversation interface {
	// Next returns the next simulated user action for the target agent.
	Next(ctx context.Context, req *TurnRequest) (*Decision, error)
	// Close releases resources owned by the conversation.
	Close() error
}

// StartRequest carries the stable inputs for one simulated conversation.
type StartRequest struct {
	// EvalCaseID identifies the eval case that owns this conversation.
	EvalCaseID string
	// Scenario defines the simulated conversation plan.
	Scenario *evalset.ConversationScenario
	// InitialSession contains the target agent session seed.
	InitialSession *evalset.SessionInput
	// SessionID identifies the target agent session.
	SessionID string
}

// TurnRequest carries the target agent output from the previous turn.
type TurnRequest struct {
	// LastTargetResponse is the final response from the previous target turn.
	LastTargetResponse *model.Message
}

// Decision describes the next action produced by the simulator.
type Decision struct {
	// Message is the next user message for the target agent.
	Message *model.Message
	// Stop signals that the conversation should end before another target turn.
	Stop bool
}

var _ Simulator = (*userSimulator)(nil)
var _ Conversation = (*conversation)(nil)

type userSimulator struct {
	simRunner runner.Runner
	options   *options
}

type conversation struct {
	simRunner              runner.Runner
	initialSession         *evalset.SessionInput
	startingPrompt         string
	stopSignal             string
	maxAllowedInvocations  int
	simUserID              string
	simSessionID           string
	injectedContextMessage []model.Message
	started                bool
	generatedInputs        int
	closed                 bool
}

// New builds the default simulator implementation backed by a runner.
func New(simRunner runner.Runner, opt ...Option) (Simulator, error) {
	_ = "STUB: not implemented"
	return *new(Simulator), nil
}

// Start creates a new default simulated conversation.
func (s *userSimulator) Start(ctx context.Context, req *StartRequest) (Conversation, error) {
	_ = "STUB: not implemented"
	return *new(Conversation), nil
}

// Next returns the next user message or a stop decision.
func (c *conversation) Next(ctx context.Context, req *TurnRequest) (*Decision, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close marks the conversation as closed.
func (c *conversation) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conversation) generateWithRunner(ctx context.Context, lastTargetResponse *model.Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func collectFinalResponseContent(events <-chan *event.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func eventFinalResponseContent(evt *event.Event) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func buildScenarioContextMessages(
	ctx context.Context,
	scenario *evalset.ConversationScenario,
	builder SystemPromptBuilder,
) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildDefaultSystemPrompt(ctx context.Context, scenario *evalset.ConversationScenario) string {
	_ = "STUB: not implemented"
	return ""
}

func buildEffectiveScenario(
	scenario *evalset.ConversationScenario,
	stopSignal string,
	maxAllowedInvocations int,
) *evalset.ConversationScenario {
	_ = "STUB: not implemented"
	return nil
}

func containsStopSignal(content string, stopSignal string) bool {
	_ = "STUB: not implemented"
	return false
}
