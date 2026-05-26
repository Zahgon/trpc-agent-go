//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates the todoenforcer extension.
//
// Run side-by-side comparisons:
//
//	# baseline: tool/todo without any enforcement (model can exit early)
//	go run ./examples/todoenforcer --enforce=false \
//	  --seed "Plan and execute a 4-step deployment: write config, run tests, deploy to staging, verify with smoke test."
//
//	# hardened: same prompt with todoenforcer wired in
//	go run ./examples/todoenforcer --enforce=true \
//	  --seed "Plan and execute a 4-step deployment: write config, run tests, deploy to staging, verify with smoke test."
//
// In the hardened run, watch the [enforce] lines: they show every
// blocked attempt the extension caught, the nudge it queued, and (if
// the model gives up via todo_declare_blocker) the formal blocker
// declaration. The key behavioural promise the demo verifies is
// that the model can no longer end the turn while open items
// remain — it must either work them, or formally declare a
// blocker.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent/extension/todoenforcer"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool/todo"
)

var (
	modelName  = flag.String("model", "deepseek-chat", "Model identifier (any OpenAI-compatible name)")
	streaming  = flag.Bool("streaming", true, "Stream the assistant response")
	variant    = flag.String("variant", "openai", "Provider variant passed to the OpenAI client")
	enforce    = flag.Bool("enforce", true, "Install the todoenforcer extension (set false to compare baseline behaviour)")
	maxRetries = flag.Int("max-retries", 3, "todoenforcer block-retry budget per invocation")
	maxTokens  = flag.Int("max-tokens", 4000, "Max completion tokens (lower if your provider rejects the default)")
	seed       = flag.String("seed", "", "Optional first user message; if set, the chat auto-starts with it")
	prefill    = flag.Bool("prefill-todos", false, "Pre-seed the session with one in_progress + one pending todo BEFORE the first turn. Useful when your provider's OpenAI-compatible tool-call adapter is unstable and you need to verify the enforcer fires without depending on the model successfully invoking todo_write first.")
)

const (
	appName   = "todoenforcer-demo"
	agentName = "todoenforcer-assistant"
)

func main() {
	flag.Parse()

	fmt.Println("todoenforcer extension demo")
	fmt.Printf("Model:       %s (variant=%s)\n", *modelName, *variant)
	fmt.Printf("Streaming:   %t\n", *streaming)
	fmt.Printf("Enforce:     %t (max-retries=%d)\n", *enforce, *maxRetries)
	fmt.Println("Commands:    /exit to quit, /list to print the current checklist")
	fmt.Println(strings.Repeat("=", 70))

	c := &chat{
		modelName:  *modelName,
		streaming:  *streaming,
		variant:    *variant,
		enforce:    *enforce,
		maxRetries: *maxRetries,
		maxTokens:  *maxTokens,
		prefill:    *prefill,
	}
	if err := c.run(); err != nil {
		log.Fatalf("chat failed: %v", err)
	}
}

// chat wraps the runner, session service and conversation loop.
type chat struct {
	modelName  string
	streaming  bool
	variant    string
	enforce    bool
	maxRetries int
	maxTokens  int
	prefill    bool

	runner    runner.Runner
	sessSvc   session.Service
	userID    string
	sessionID string
}

func (c *chat) run() error { _ = "STUB: not implemented"; return nil }

// prefillTodos simulates a mid-task resumption scenario. It
// reproduces, by hand, the exact pair of artefacts a real prior
// turn would have left behind:
//
//  1. session.State holds the current todo list (this is what
//     the enforcer's AfterModel reads to detect open items).
//  2. The session's event log holds the chat history that proves
//     the model itself already wrote the list — a user turn
//     asking for the work, an assistant turn calling todo_write,
//     and a tool turn carrying todo_write's response.
//
// Both are required for the demo to feel honest:
//
//   - Without (1) the enforcer never trips, since there are no
//     open items in state.
//   - Without (2) the model is blind to its own past — when the
//     user later says "please continue what you were doing", the
//     model has nothing in its visible chat history to continue
//     from. It will (correctly) say "I don't have any prior
//     context" and force the enforcer into a teaching role
//     ("here is a task you never knew about"), which is NOT the
//     real production semantics. The extension is meant to remind
//     a model of work it KNOWS it owes, not to dump unfamiliar
//     work onto it.
//
// We construct (2) using event.NewResponseEvent + AppendEvent,
// the same path examples/session/appendevent/ uses. The Branch
// field on each event matches the agent's branch (the agent
// name in a single-agent setup), so when the runner builds the
// next request these messages show up in the right context.
func (c *chat) prefillTodos(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Branch defaults to the agent name for a single-agent setup
// — same convention printTodos / GetTodos / the enforcer use.

// Build the synthetic prior turn. The arguments and result
// payload mirror what a real todo_write call would carry, so
// downstream consumers (event log replays, evaluation
// pipelines, ...) see a turn that is structurally
// indistinguishable from one the model actually produced.

// 1. The original user request that kicked the work off.

// 2. The assistant turn that called todo_write to plan.

// 3. The tool turn carrying todo_write's response.
// ToolName is set even though the spec doesn't strictly
// require it for assistant↔tool pairing — several
// Some OpenAI-compatible adapters reject tool messages
// whose name field is empty, even when the matching ID is
// present.

// newPrefillEvent is a thin wrapper that fills in the boilerplate
// every prefilled event needs: a non-streaming, non-partial
// response carrying exactly one Choice with the supplied message.
// IsValidContent() returns true for tool-call / tool-result
// messages even when Content is empty, which is why the assistant
// tool_call leg is well-formed despite having no text body.
func newPrefillEvent(invocationID, author string, msg model.Message) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (c *chat) setup() error { _ = "STUB: not implemented"; return nil }

// Two construction paths so the same demo can show both
// "raw tool/todo" baseline and "hardened by todoenforcer".
// Other than the WithExtensions / WithTools split, the agent
// configuration is identical, which keeps the comparison
// honest — any behavioural difference you see in [enforce]
// lines comes purely from the extension being installed.

// Install the extension. WithExtensions also routes the
// tools the enforcer contributes via extension.Registry.Tools
// — the enforcer ships both todo_write AND
// todo_declare_blocker, so we do NOT pass either via
// WithTools. Passing them here on top would trigger the
// name-collision dedup and silently drop the enforcer's
// copies; the dedup is correct but it's not what we want
// the demo to show.

// Baseline: hand-install todo so the model has a place to
// write a plan, but with no enforcement. The model is free
// to mark items pending and still emit a final answer —
// which is exactly the behaviour the extension exists to
// fix.

// observeEnforce is the OnEnforce callback. It runs on the model
// callback hot path, so we keep it cheap (just printing). We
// serialise prints with a mutex because the enforcer hook can
// race with the streaming output goroutine in processResponse —
// without it [enforce] lines occasionally interleave with
// assistant deltas mid-token.
var enforceMu sync.Mutex

func (c *chat) observeEnforce(evt todoenforcer.EnforceEvent) { _ = "STUB: not implemented"; return }

func (c *chat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *chat) processMessage(ctx context.Context, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

// processResponse renders the stream: tool calls, tool responses
// and assistant text are each formatted distinctly so the
// enforcement loop is easy to follow visually.
func (c *chat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// printTodos reads the canonical checklist from session state at
// the end of a turn. The enforcer reads from the same place, so
// what you see here is what the enforcer's AfterModel saw when it
// decided to block (or pass).
func (c *chat) printTodos() { _ = "STUB: not implemented"; return }

func formatTodos(todos []todo.Item) string { _ = "STUB: not implemented"; return "" }

// truncate keeps the demo output legible when a tool result is
// long (the JSON payload of a 10-item todo list is otherwise
// noisy enough to push the [enforce] lines off-screen).
func truncate(s string, n int) string { _ = "STUB: not implemented"; return "" }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }
