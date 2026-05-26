//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates that a sub-agent's skill_load does not leak
// loaded skill bodies/docs into the coordinator's prompt.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
	agenttool "trpc.group/trpc-go/trpc-agent-go/tool/agent"
)

const (
	appName = "skillisolation-example"
	userID  = "skillisolation-user"

	coordinatorAgentName = "skillisolation-coordinator"
	childAgentName       = "skillisolation-child"

	demoSkillName      = "demo-skill"
	defaultSkillsRoot  = "./skills"
	defaultModelName   = "gpt-5"
	loadedMarkerPrefix = "[Loaded] "
	loadedMarker       = loadedMarkerPrefix + demoSkillName
)

func main() {
	var (
		flagModel = flag.String(
			"model",
			defaultModelName,
			"OpenAI-compatible model name",
		)
		flagSkillsRoot = flag.String(
			"skills-root",
			defaultSkillsRoot,
			"Skills root directory",
		)
	)
	flag.Parse()

	repo, err := skill.NewFSRepository(*flagSkillsRoot)
	if err != nil {
		log.Fatalf("load skills repo: %v", err)
	}

	mdl := openai.New(*flagModel)

	child := llmagent.New(
		childAgentName,
		llmagent.WithModel(mdl),
		llmagent.WithSkills(repo),
		llmagent.WithEnableCodeExecutionResponseProcessor(false),
		llmagent.WithInstruction(childInstruction()),
		llmagent.WithInputSchema(agentToolInputSchema()),
		llmagent.WithGenerationConfig(model.GenerationConfig{
			Stream: false,
		}),
	)
	childTool := agenttool.NewTool(child)

	coordinator := llmagent.New(
		coordinatorAgentName,
		llmagent.WithModel(mdl),
		llmagent.WithSkills(repo),
		llmagent.WithEnableCodeExecutionResponseProcessor(false),
		llmagent.WithTools([]tool.Tool{childTool}),
		llmagent.WithInstruction(coordinatorInstruction()),
		llmagent.WithInputSchema(agentToolInputSchema()),
		llmagent.WithModelCallbacks(coordinatorCallbacks()),
		llmagent.WithGenerationConfig(model.GenerationConfig{
			Stream: false,
		}),
	)

	svc := inmemory.NewSessionService()
	r := runner.NewRunner(
		appName,
		coordinator,
		runner.WithSessionService(svc),
	)
	defer r.Close()

	sessionID := fmt.Sprintf("skillisolation-%d", time.Now().Unix())
	fmt.Printf("Session: %s\n", sessionID)

	ctx := context.Background()
	msg := model.NewUserMessage(
		"Call the sub-agent tool. " +
			"The sub-agent must load demo-skill via skill_load.",
	)
	events, err := r.Run(ctx, userID, sessionID, msg)
	if err != nil {
		log.Fatalf("run: %v", err)
	}
	printTranscript(events)

	sess, err := svc.GetSession(ctx, session.Key{
		AppName:   appName,
		UserID:    userID,
		SessionID: sessionID,
	})
	if err != nil {
		log.Fatalf("get session: %v", err)
	}

	printStateSummary(sess)
}

func coordinatorCallbacks() *model.Callbacks { _ = "STUB: not implemented"; return nil }

func systemMessage(msgs []model.Message) string { _ = "STUB: not implemented"; return "" }

func loadedSkillNames(
	inv *agent.Invocation,
	agentName string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func printStateSummary(sess *session.Session) { _ = "STUB: not implemented"; return }

func hasStateKey(state session.StateMap, key string) bool { _ = "STUB: not implemented"; return false }

func listSkillStateKeys(state session.StateMap) []string { _ = "STUB: not implemented"; return nil }

func drain(events <-chan *event.Event) { _ = "STUB: not implemented"; return }

func printTranscript(events <-chan *event.Event) { _ = "STUB: not implemented"; return }

func printEvent(evt *event.Event) { _ = "STUB: not implemented"; return }

func childInstruction() string { _ = "STUB: not implemented"; return "" }

func coordinatorInstruction() string { _ = "STUB: not implemented"; return "" }

func agentToolInputSchema() map[string]any { _ = "STUB: not implemented"; return nil }
