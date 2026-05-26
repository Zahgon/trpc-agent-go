//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates the SkillToolProfile option.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	appName        = "skilltoolprofile-example"
	agentName      = "skilltoolprofile-agent"
	userID         = "skilltoolprofile-user"
	defaultProfile = llmagent.SkillToolProfileFull
	demoSkillName  = "demo-profile"
)

func main() {
	var (
		flagProfile = flag.String(
			"profile",
			string(defaultProfile),
			"Skill tool profile: full|knowledge_only",
		)
		flagSkillsRoot = flag.String(
			"skills-root",
			defaultSkillsRoot(),
			"Skills root directory",
		)
	)
	flag.Parse()

	profile, err := parseProfile(*flagProfile)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := skill.NewFSRepository(*flagSkillsRoot)
	if err != nil {
		log.Fatalf("load skills repo: %v", err)
	}

	agt, executorLabel := newProfileAgent(profile, repo)
	r := runner.NewRunner(
		appName,
		agt,
		runner.WithSessionService(inmemory.NewSessionService()),
	)
	defer r.Close()

	sessionID := fmt.Sprintf(
		"skilltoolprofile-%d",
		time.Now().Unix(),
	)

	fmt.Printf("Profile: %s\n", profile)
	fmt.Printf("Skills root: %s\n", *flagSkillsRoot)
	fmt.Printf("Executor: %s\n", executorLabel)
	fmt.Println("Registered skill tools:")
	for _, name := range listSkillToolNames(agt.Tools()) {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()

	if profile == llmagent.SkillToolProfileKnowledgeOnly {
		fmt.Println(
			"Demo flow: skill_load -> skill_list_docs -> " +
				"skill_select_docs -> assistant",
		)
	} else {
		fmt.Println("Demo flow: skill_load -> skill_run -> assistant")
	}
	fmt.Printf("Session: %s\n\n", sessionID)

	ctx := context.Background()
	evCh, err := r.Run(ctx, userID, sessionID, model.NewUserMessage("demo"))
	if err != nil {
		log.Fatalf("run: %v", err)
	}
	for ev := range evCh {
		printEvent(ev)
	}
}

func defaultSkillsRoot() string { _ = "STUB: not implemented"; return "" }

func parseProfile(raw string) (llmagent.SkillToolProfile, error) {
	_ = "STUB: not implemented"
	return *new(llmagent.SkillToolProfile), nil
}

func newProfileAgent(
	profile llmagent.SkillToolProfile,
	repo skill.Repository,
) (*llmagent.LLMAgent, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// The CodeExecutor here exists so skill_run /
// workspace_exec can run; it is not a license to
// auto-execute fenced code from assistant replies. Keep
// that orthogonal switch explicitly off so the demo
// showcases only the skill-tool execution path.

func listSkillToolNames(ts []tool.Tool) []string { _ = "STUB: not implemented"; return nil }

func printEvent(evt *event.Event) { _ = "STUB: not implemented"; return }

func compactText(s string) string { _ = "STUB: not implemented"; return "" }

func agentDescription(profile llmagent.SkillToolProfile) string {
	_ = "STUB: not implemented"
	return ""
}

func agentInstruction(profile llmagent.SkillToolProfile) string {
	_ = "STUB: not implemented"
	return ""
}

type profileModel struct {
	profile llmagent.SkillToolProfile
	step    int
}

func newProfileModel(profile llmagent.SkillToolProfile) *profileModel {
	_ = "STUB: not implemented"
	return nil
}

func (m *profileModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *profileModel) GenerateContent(
	ctx context.Context,
	_ *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *profileModel) fullResponse() *model.Response { _ = "STUB: not implemented"; return nil }

func (m *profileModel) knowledgeOnlyResponse() *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func toolCallResponse(id string, toolName string, args string) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func assistantResponse(content string) *model.Response { _ = "STUB: not implemented"; return nil }
