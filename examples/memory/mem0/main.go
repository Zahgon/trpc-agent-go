//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	memorymem0 "trpc.group/trpc-go/trpc-agent-go/memory/mem0"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	sessioninmemory "trpc.group/trpc-go/trpc-agent-go/session/inmemory"
)

const defaultModelName = "deepseek-v4-flash"

var (
	modelName  = flag.String("model", defaultModelName, "Chat model name")
	appName    = flag.String("app", "mem0-integration-demo", "Application name used for mem0 ownership")
	userID     = flag.String("user", "demo-user", "User ID used for mem0 ownership")
	sessionID  = flag.String("session", "", "Session ID (default: generated from timestamp)")
	waitFor    = flag.Duration("wait-timeout", 90*time.Second, "How long to wait for the memory to become readable")
	withCustom = flag.Bool("with-options", false, "If set, calls IngestSession directly with custom per-request IngestOption values to demonstrate the option pattern (bypasses the runner's default ingestion)")
	tagValue   = flag.String("tag", "support", "Value attached to mem0 metadata when -with-options is set")
)

func main() {
	flag.Parse()
	if os.Getenv("MEM0_API_KEY") == "" {
		log.Fatal("MEM0_API_KEY is required")
	}
	if os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("OPENAI_API_KEY is required")
	}

	sid := *sessionID
	if sid == "" {
		sid = fmt.Sprintf("mem0-%d", time.Now().Unix())
	}
	token := fmt.Sprintf("Mem0IntegrationDemo-%d", time.Now().UnixNano())
	userMessage := fmt.Sprintf("For future reference, my dog is named %s. Please reply briefly.", token)

	mem0Svc, err := newMem0Service(*waitFor)
	if err != nil {
		log.Fatalf("create mem0 service: %v", err)
	}
	defer mem0Svc.Close()

	chatAgent := llmagent.New(
		"mem0-demo-agent",
		llmagent.WithModel(openai.New(*modelName)),
		llmagent.WithDescription("A concise assistant with mem0-backed long-term memory integration."),
		llmagent.WithTools(mem0Svc.Tools()),
	)

	sessSvc := sessioninmemory.NewSessionService()
	runnerOpts := []runner.Option{
		runner.WithSessionService(sessSvc),
	}
	// When -with-options is set the example bypasses the runner-driven
	// ingestion so it can call IngestSession directly with custom per-request
	// options. Otherwise the runner attaches its default options
	// (WithIngestRunID(sess.ID) and WithIngestAgentID(invocation.AgentName)).
	if !*withCustom {
		runnerOpts = append(runnerOpts, runner.WithSessionIngestor(mem0Svc))
	}
	r := runner.NewRunner(*appName, chatAgent, runnerOpts...)
	defer r.Close()

	ctx := context.Background()
	fmt.Printf("Model: %s\nApp: %s\nUser: %s\nSession: %s\nToken: %s\n", *modelName, *appName, *userID, sid, token)
	fmt.Printf("Message: %s\n", userMessage)
	fmt.Println(strings.Repeat("=", 60))

	result, err := runOnce(ctx, r, *userID, sid, model.NewUserMessage(userMessage))
	if err != nil {
		log.Fatalf("runner failed: %v", err)
	}
	if len(result.toolCalls) > 0 {
		fmt.Printf("Tool calls: %s\n", strings.Join(result.toolCalls, ", "))
	} else {
		fmt.Println("Tool calls: <none>")
	}
	if reply := strings.TrimSpace(result.reply); reply != "" {
		fmt.Printf("Assistant: %s\n", reply)
	}
	fmt.Println()

	// When custom options are requested, exercise the per-request IngestOption
	// API directly so callers can verify metadata/agent_id/run_id round-trip
	// to mem0 records.
	if *withCustom {
		sess, err := lookupSession(ctx, sessSvc, *userID, sid)
		if err != nil {
			log.Fatalf("lookup session: %v", err)
		}
		if err := mem0Svc.IngestSession(ctx, sess,
			session.WithIngestMetadata(map[string]any{
				"trpc_demo_tag":   *tagValue,
				"trpc_demo_token": token,
			}),
			session.WithIngestAgentID("billing-bot"),
			session.WithIngestRunID(fmt.Sprintf("ticket-%d", time.Now().UnixNano())),
		); err != nil {
			log.Fatalf("custom IngestSession: %v", err)
		}
		fmt.Println("Submitted IngestSession with custom IngestOption set; verifying metadata...")
	}

	entries, err := waitForToken(ctx, mem0Svc, memory.UserKey{AppName: *appName, UserID: *userID}, token, *waitFor)
	if err != nil {
		// Mem0 sometimes takes longer than the demo timeout to make natively
		// ingested memories visible via search (the underlying ingest is async
		// even with async_mode=false). Warn rather than fail so the per-request
		// option demo (above) is still observable to the caller.
		log.Printf("warning: memory not yet searchable via SearchMemories: %v", err)
		return
	}
	fmt.Printf("Stored memories (%d):\n", len(entries))
	for i, entry := range entries {
		fmt.Printf("  %d. %s\n", i+1, entry.Memory.Memory)
		if extras := summariseExtras(entry); extras != "" {
			fmt.Printf("     %s\n", extras)
		}
	}
}

// summariseExtras renders the option-driven fields (tags, agent_id, run_id)
// that mem0 echoes back inside Entry.Memory.Topics / Memory.Metadata so the
// example clearly shows the per-request options round-tripping.
func summariseExtras(entry *memory.Entry) string { _ = "STUB: not implemented"; return "" }

func newMem0Service(timeout time.Duration) (*memorymem0.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mem0Host() string { _ = "STUB: not implemented"; return "" }

type runResult struct {
	toolCalls []string
	reply     string
}

func runOnce(ctx context.Context, r runner.Runner, userID, sessionID string, msg model.Message) (*runResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectResponse(out *runResult, seen map[string]struct{}, evt *event.Event) {
	_ = "STUB: not implemented"
	return
}

// lookupSession fetches the persisted session the runner just produced from
// the supplied session.Service. It is used by the -with-options demo to call
// IngestSession directly with custom per-request IngestOption values.
func lookupSession(ctx context.Context, svc session.Service, userID, sessionID string) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func waitForToken(ctx context.Context, svc *memorymem0.Service, userKey memory.UserKey, token string, timeout time.Duration) ([]*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
