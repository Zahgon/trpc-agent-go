//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates how SkillLoadMode affects skill load lifetime.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

const (
	appName       = "skillloadmode-example"
	agentName     = "skillloadmode-agent"
	userID        = "skillloadmode-user"
	skillName     = "demo-skill"
	defaultMode   = llmagent.SkillLoadModeTurn
	defaultRoot   = "./skills"
	stateValueOne = "1"
)

func main() {
	var (
		flagMode = flag.String(
			"mode",
			defaultMode,
			"SkillLoadMode: once | turn | session",
		)
		flagSkillsRoot = flag.String(
			"skills-root",
			defaultRoot,
			"Skills root directory",
		)
		flagToolResults = flag.Bool(
			"tool-results",
			false,
			"Materialize loaded skill content into tool results",
		)
	)
	flag.Parse()

	mode := strings.ToLower(strings.TrimSpace(*flagMode))
	if mode != llmagent.SkillLoadModeOnce &&
		mode != llmagent.SkillLoadModeTurn &&
		mode != llmagent.SkillLoadModeSession {
		log.Fatalf("invalid -mode: %q", mode)
	}

	repo, err := skill.NewFSRepository(*flagSkillsRoot)
	if err != nil {
		log.Fatalf("load skills repo: %v", err)
	}

	mock := newStepModel(skillName)
	agt := llmagent.New(
		agentName,
		llmagent.WithModel(mock),
		llmagent.WithSkills(repo),
		llmagent.WithSkillLoadMode(mode),
		llmagent.WithSkillsLoadedContentInToolResults(*flagToolResults),
	)

	svc := inmemory.NewSessionService()
	r := runner.NewRunner(
		appName,
		agt,
		runner.WithSessionService(svc),
	)
	defer r.Close()

	sessionID := fmt.Sprintf(
		"skillloadmode-%d",
		time.Now().Unix(),
	)

	fmt.Printf("SkillLoadMode: %s\n", mode)
	fmt.Printf("Tool result materialization: %t\n", *flagToolResults)
	fmt.Printf("Skills root: %s\n", *flagSkillsRoot)
	fmt.Printf("Session: %s\n\n", sessionID)

	ctx := context.Background()

	fmt.Println("Turn 1: model calls skill_load")
	runOnce(ctx, r, sessionID)
	printSkillState(ctx, svc, sessionID)

	fmt.Println("\nTurn 2: no tool calls (observe auto-clearing)")
	runOnce(ctx, r, sessionID)
	printSkillState(ctx, svc, sessionID)
}

func runOnce(
	ctx context.Context,
	r runner.Runner,
	sessionID string,
) {
	_ = "STUB: not implemented"
	return
}

func printSelectedEvent(evt *event.Event) { _ = "STUB: not implemented"; return }

func printSkillState(
	ctx context.Context,
	svc session.Service,
	sessionID string,
) {
	_ = "STUB: not implemented"
	return
}

func listSkillStateKeys(state session.StateMap) []string { _ = "STUB: not implemented"; return nil }

type stepModel struct {
	skill string
	step  int
}

func newStepModel(skillName string) *stepModel { _ = "STUB: not implemented"; return nil }

func (m *stepModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *stepModel) GenerateContent(
	ctx context.Context,
	_ *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toolCallResponse(
	id string,
	toolName string,
	args []byte,
) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func assistantResponse(content string) *model.Response { _ = "STUB: not implemented"; return nil }
