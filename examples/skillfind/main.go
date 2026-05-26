//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates real public skill discovery, installation,
// and execution with Agent Skills.
package main

import (
	"context"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

const (
	skillFindAppName   = "skillfind-example"
	skillFindAgentName = "skillfind-agent"

	defaultModelNameValue = "gpt-5.2"
	defaultUserID         = "demo-user"

	exitCommand  = "exit"
	newCommand   = "/new"
	listCommand  = "/skills"
	resetCommand = "/reset-skills"
)

const agentInstructionBase = `
You are a skill-enabled assistant.

If the user asks to find, install, or try a public Agent Skill, load the
local skill named "skill-find" first.

After skill_install_github succeeds, use the returned skill_name with
skill_load.

Explain briefly which public skill you installed and what happened.`

const agentInstructionRunDisabled = `
Local execution is disabled for this demo. Do not call skill_run. Search,
install, and load skills only.`

const agentInstructionRunEnabled = `
Local execution is enabled for this demo. Do not call skill_run unless
the user explicitly asks to run the installed skill. Never execute
downloaded code automatically.`

type skillFindChat struct {
	modelName       string
	commonSkillsDir string
	userSkillsDir   string
	userID          string
	sessionID       string
	oneShotPrompt   string
	allowSkillRun   bool
	resetUserSkills bool

	repo   *skill.FSRepository
	runner runner.Runner
}

func main() {
	flags := parseFlags()

	chat := &skillFindChat{
		modelName:       flags.modelName,
		commonSkillsDir: flags.commonSkillsDir,
		userSkillsDir:   flags.userSkillsDir,
		userID:          flags.userID,
		sessionID:       newSessionID(),
		oneShotPrompt:   flags.oneShotPrompt,
		allowSkillRun:   flags.allowSkillRun,
		resetUserSkills: flags.resetUserSkills,
	}

	if err := chat.run(context.Background()); err != nil {
		log.Fatalf("skillfind example failed: %v", err)
	}
}

type cliFlags struct {
	modelName       string
	commonSkillsDir string
	userSkillsDir   string
	userID          string
	oneShotPrompt   string
	allowSkillRun   bool
	resetUserSkills bool
}

func parseFlags() cliFlags { _ = "STUB: not implemented"; return *new(cliFlags) }

func defaultModelName() string { _ = "STUB: not implemented"; return "" }

func defaultCommonSkillsDir() string { _ = "STUB: not implemented"; return "" }

func defaultUserSkillsDir(userID string) string { _ = "STUB: not implemented"; return "" }

func exampleDir() string { _ = "STUB: not implemented"; return "" }

func newSessionID() string { _ = "STUB: not implemented"; return "" }

func (c *skillFindChat) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *skillFindChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *skillFindChat) printBanner() { _ = "STUB: not implemented"; return }

func (c *skillFindChat) startInteractiveChat(
	ctx context.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *skillFindChat) processMessage(
	ctx context.Context,
	userMessage string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func printToolCalls(toolCalls []model.ToolCall) { _ = "STUB: not implemented"; return }

func compactText(value string) string { _ = "STUB: not implemented"; return "" }

func (c *skillFindChat) printInstalledSkills() { _ = "STUB: not implemented"; return }

func (c *skillFindChat) resetUserSkillRoot() error { _ = "STUB: not implemented"; return nil }

func intPtr(value int) *int { _ = "STUB: not implemented"; return nil }

func buildAgentInstruction(allowSkillRun bool) string { _ = "STUB: not implemented"; return "" }

func skillToolProfile(
	allowSkillRun bool,
) llmagent.SkillToolProfile {
	_ = "STUB: not implemented"
	return *new(llmagent.SkillToolProfile)
}

func promptLines(allowSkillRun bool) []string { _ = "STUB: not implemented"; return nil }
