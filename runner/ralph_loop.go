//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package runner

import (
	"context"
	"errors"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultRalphMaxIterations = 10

	defaultPromiseTagOpen  = "<promise>"
	defaultPromiseTagClose = "</promise>"

	defaultRalphEventBuffer = 256

	ralphVerifyShell = "bash"
	ralphVerifyFlag  = "-lc"
)

var errRalphLoopMissingStopCondition = errors.New(
	"ralph loop: missing completion promise, verify command, and verifiers",
)

// VerifyResult describes one verifier outcome.
type VerifyResult struct {
	// Passed reports whether the completion condition is satisfied.
	Passed bool
	// Feedback is a human-readable failure message for the next iteration.
	// It is ignored when Passed is true.
	Feedback string
}

// Verifier checks whether the task is complete.
//
// Verifiers are only called after an agent iteration completes.
// All configured verifiers must pass before Ralph Loop allows the run to
// stop.
type Verifier interface {
	Verify(
		ctx context.Context,
		invocation *agent.Invocation,
		lastEvent *event.Event,
	) (VerifyResult, error)
}

// RalphLoopConfig controls a runner-level Ralph Loop mode.
//
// Ralph Loop is an "outer loop": instead of trusting the Large Language Model
// (LLM) to decide when it is done, the runner keeps iterating until a
// verifiable completion condition is met (or max iterations is reached).
type RalphLoopConfig struct {
	// MaxIterations is the maximum number of agent iterations to run.
	// When <= 0, a safe default is used.
	MaxIterations int

	// CompletionPromise is an optional stop signal.
	//
	// When set, the loop stops only if the agent output contains:
	//   <promise>CompletionPromise</promise>
	//
	// The tag strings are configurable via PromiseTagOpen/PromiseTagClose.
	CompletionPromise string
	PromiseTagOpen    string
	PromiseTagClose   string

	// VerifyCommand is an optional command to run after each iteration.
	// The loop stops only if the command exits with code 0.
	VerifyCommand string
	VerifyWorkDir string
	VerifyTimeout time.Duration
	VerifyEnv     map[string]string
	VerifyRunner  RalphLoopCommandRunner

	// Verifiers is an optional list of additional completion checks.
	// All verifiers must pass.
	Verifiers []Verifier
}

// RalphLoopCommandRunner executes VerifyCommand.
type RalphLoopCommandRunner interface {
	Run(
		ctx context.Context,
		spec RalphLoopCommandSpec,
	) (RalphLoopCommandResult, error)
}

// RalphLoopCommandSpec describes a command execution.
type RalphLoopCommandSpec struct {
	Command string
	WorkDir string
	Env     map[string]string
	Timeout time.Duration
}

// RalphLoopCommandResult captures command output and exit status.
type RalphLoopCommandResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
	Duration time.Duration
	TimedOut bool
}

// WithRalphLoop enables Ralph Loop mode on a Runner.
//
// It wraps the selected agent so the run will keep iterating until:
//   - CompletionPromise is detected (when set), AND
//   - VerifyCommand exits with code 0 (when set), AND
//   - All Verifiers pass (when set), OR
//   - MaxIterations is reached.
func WithRalphLoop(cfg RalphLoopConfig) Option { _ = "STUB: not implemented"; return *new(Option) }

func wrapAgentsWithRalphLoop(
	agents map[string]agent.Agent,
	cfg RalphLoopConfig,
) {
	_ = "STUB: not implemented"
	return
}

func wrapAgentWithRalphLoop(
	ag agent.Agent,
	cfg RalphLoopConfig,
) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

type ralphLoopAgent struct {
	inner agent.Agent
	cfg   RalphLoopConfig
}

func (a *ralphLoopAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

func (a *ralphLoopAgent) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

func (a *ralphLoopAgent) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

func (a *ralphLoopAgent) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func (a *ralphLoopAgent) Run(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ralphLoopAgent) runLoop(
	ctx context.Context,
	base *agent.Invocation,
	out chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (a *ralphLoopAgent) newInnerInvocation(
	base *agent.Invocation,
	entryPredecessors []string,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func (a *ralphLoopAgent) forwardEvents(
	ctx context.Context,
	events <-chan *event.Event,
	out chan<- *event.Event,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (a *ralphLoopAgent) verifyIteration(
	ctx context.Context,
	base *agent.Invocation,
	lastFull *event.Event,
) (bool, string, error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

func (a *ralphLoopAgent) verifiersSatisfied(
	ctx context.Context,
	base *agent.Invocation,
	lastFull *event.Event,
) (bool, string, error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

func hasNonNilVerifier(verifiers []Verifier) bool { _ = "STUB: not implemented"; return false }

func (a *ralphLoopAgent) promiseSatisfied(lastFull *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *ralphLoopAgent) missingPromiseReport() string { _ = "STUB: not implemented"; return "" }

func (a *ralphLoopAgent) commandSatisfied(
	ctx context.Context,
) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (a *ralphLoopAgent) appendFeedback(
	ctx context.Context,
	base *agent.Invocation,
	iter int,
	feedback string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *ralphLoopAgent) emitStopError(
	ctx context.Context,
	base *agent.Invocation,
	out chan<- *event.Event,
	message string,
) {
	_ = "STUB: not implemented"
	return
}

func validateRalphLoopConfig(cfg RalphLoopConfig) error { _ = "STUB: not implemented"; return nil }

func normalizeRalphLoopConfig(cfg RalphLoopConfig) RalphLoopConfig {
	_ = "STUB: not implemented"
	return *new(RalphLoopConfig)
}

func firstTagText(
	evt *event.Event,
	open string,
	closeTag string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func firstTagTextInString(
	text string,
	open string,
	closeTag string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func textFromContentParts(parts []model.ContentPart) string { _ = "STUB: not implemented"; return "" }

func normalizePromiseText(s string) string { _ = "STUB: not implemented"; return "" }

func formatCommandFailure(
	cmd string,
	res RalphLoopCommandResult,
) string {
	_ = "STUB: not implemented"
	return ""
}

type hostRalphLoopRunner struct{}

func (hostRalphLoopRunner) Run(
	ctx context.Context,
	spec RalphLoopCommandSpec,
) (RalphLoopCommandResult, error) {
	_ = "STUB: not implemented"
	return *new(RalphLoopCommandResult), nil
}

//nolint:gosec // VerifyCommand is explicitly provided by the caller.

func mergeEnv(overrides map[string]string) []string { _ = "STUB: not implemented"; return nil }
