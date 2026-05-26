//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package skillprofile centralizes built-in skill tool profile
// normalization and tool-flag resolution.
package skillprofile

const (
	// Full enables the complete built-in skill tool suite, including
	// skill_run / skill_exec and the interactive exec-session helpers.
	// This profile is opt-in; callers that want skill-driven execution
	// must request it explicitly.
	Full = "full"
	// KnowledgeOnly enables progressive-disclosure skill tools used for
	// knowledge injection without exposing execution tools. This is the
	// default profile: agents that only configure a skill repository
	// get skill_load / skill_list_docs / skill_select_docs, and rely on
	// workspace_exec for any script execution.
	KnowledgeOnly = "knowledge_only"

	// ToolLoad is the built-in tool name for loading SKILL.md and optional docs.
	ToolLoad = "skill_load"
	// ToolListDocs is the built-in tool name for listing skill docs.
	ToolListDocs = "skill_list_docs"
	// ToolSelectDocs is the built-in tool name for selecting skill docs.
	ToolSelectDocs = "skill_select_docs"
	// ToolRun is the built-in tool name for non-interactive skill execution.
	ToolRun = "skill_run"
	// ToolExec is the built-in tool name for interactive skill execution.
	ToolExec = "skill_exec"
	// ToolWriteStdin is the built-in tool name for writing to skill_exec stdin.
	ToolWriteStdin = "skill_write_stdin"
	// ToolPollSession is the built-in tool name for polling skill_exec sessions.
	ToolPollSession = "skill_poll_session"
	// ToolKillSession is the built-in tool name for terminating skill_exec sessions.
	ToolKillSession = "skill_kill_session"
)

// Flags describes which built-in skill tools are enabled for a profile.
type Flags struct {
	Load        bool
	SelectDocs  bool
	ListDocs    bool
	Run         bool
	Exec        bool
	WriteStdin  bool
	PollSession bool
	KillSession bool
}

// Normalize canonicalizes a profile name. An empty or unknown value
// resolves to KnowledgeOnly, matching the framework default that keeps
// skill execution tools (skill_run / skill_exec / interactive helpers)
// off unless they are explicitly opted in via Full.
func Normalize(profile string) string { _ = "STUB: not implemented"; return "" }

// NormalizeTool canonicalizes a built-in skill tool name.
func NormalizeTool(name string) string { _ = "STUB: not implemented"; return "" }

// ResolveFlags returns the enabled built-in skill tool set.
//
// When allowedTools is nil, the resolved flags come from the profile preset.
// When allowedTools is non-nil, it is treated as an explicit whitelist that
// overrides the profile.
func ResolveFlags(profile string, allowedTools []string) (Flags, error) {
	_ = "STUB: not implemented"
	return *new(Flags), nil
}

// IsKnowledgeOnly reports whether the profile is knowledge-only after
// normalization.
func IsKnowledgeOnly(profile string) bool { _ = "STUB: not implemented"; return false }

// IsExplicitKnowledgeOnly reports whether the caller explicitly selected the
// knowledge-only profile. It distinguishes an explicit opt-in from the
// unconfigured default (empty string), which also normalizes to
// KnowledgeOnly but carries different semantics around convenience
// fallbacks (for example, llmagent auto-falling back to a local code
// executor when skills are enabled without one).
func IsExplicitKnowledgeOnly(profile string) bool { _ = "STUB: not implemented"; return false }

// Any reports whether any built-in skill tool is enabled.
func (f Flags) Any() bool { _ = "STUB: not implemented"; return false }

// HasKnowledgeTools reports whether any non-executing skill disclosure tools
// are enabled.
func (f Flags) HasKnowledgeTools() bool { _ = "STUB: not implemented"; return false }

// HasDocHelpers reports whether any doc inspection/selection helpers are
// enabled beyond skill_load itself.
func (f Flags) HasDocHelpers() bool { _ = "STUB: not implemented"; return false }

// RequiresExecutionTools reports whether the profile needs an executor.
func (f Flags) RequiresExecutionTools() bool { _ = "STUB: not implemented"; return false }

// RequiresExecSessionTools reports whether the profile exposes interactive
// exec session helpers.
func (f Flags) RequiresExecSessionTools() bool { _ = "STUB: not implemented"; return false }

// WithoutInteractiveExecution removes interactive execution capabilities while
// preserving non-interactive execution.
func (f Flags) WithoutInteractiveExecution() Flags { _ = "STUB: not implemented"; return *new(Flags) }

func presetFlags(profile string) Flags { _ = "STUB: not implemented"; return *new(Flags) }

func flagsFromAllowedTools(allowedTools []string) (Flags, error) {
	_ = "STUB: not implemented"
	return *new(Flags), nil
}

func validateFlags(flags Flags) error { _ = "STUB: not implemented"; return nil }
