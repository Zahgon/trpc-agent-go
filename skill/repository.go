//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights
// reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package skill provides a model-agnostic Agent Skills repository.
// A skill is a folder containing a SKILL.md file with an optional YAML front
// matter block and a Markdown body, plus optional doc files.
//
// If the front matter is missing or does not specify `name`, the skill name
// falls back to the folder name.
package skill

import (
	"bufio"
	"sync"
)

// skillFile is the canonical skill definition filename.
const skillFile = "SKILL.md"

// SkillFile is the canonical skill definition filename exposed for
// prompt rendering and other callers that need the on-disk path.
const SkillFile = skillFile

// EnvSkillsRoot is the environment variable name that points to the
// skills repository root directory used by examples and runtimes.
// Defining it here avoids repeated string literals across the codebase.
const EnvSkillsRoot = "SKILLS_ROOT"

// Summary contains the minimal information for a skill.
type Summary struct {
	Name        string
	Description string
}

// Doc represents an auxiliary document of a skill.
type Doc struct {
	Path    string
	Content string
}

// Skill contains full content of a skill.
type Skill struct {
	Summary Summary
	Body    string
	Docs    []Doc
}

// Repository is a source of skills.
type Repository interface {
	// Summaries returns all available skill summaries.
	Summaries() []Summary
	// Get returns a full skill by name.
	Get(name string) (*Skill, error)
	// Path returns the directory path that contains the given skill.
	// It allows staging the whole skill folder for execution.
	Path(name string) (string, error)
}

// RootedRepository optionally exposes the configured skill roots.
// Runtimes can use this to render compact directory locators in prompts.
type RootedRepository interface {
	Repository
	Roots() []string
}

// RefreshableRepository can rescan its backing skill sources.
type RefreshableRepository interface {
	Repository
	Refresh() error
}

// FSRepository implements Repository backed by filesystem roots.
type FSRepository struct {
	roots []string
	mu    sync.RWMutex
	// name -> directory path that contains SKILL.md
	index map[string]string
}

// NewFSRepository creates a FSRepository scanning the given roots.
func NewFSRepository(roots ...string) (*FSRepository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Refresh rescans the configured roots and replaces the skill index.
func (r *FSRepository) Refresh() error { _ = "STUB: not implemented"; return nil }

// Path returns the directory path that contains the given skill.
// It allows staging the whole skill folder for execution.
func (r *FSRepository) Path(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Roots returns the configured filesystem roots.
func (r *FSRepository) Roots() []string { _ = "STUB: not implemented"; return nil }

func scanRoots(roots []string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Record first occurrence; later ones ignored.

// Summaries implements Repository.
func (r *FSRepository) Summaries() []Summary { _ = "STUB: not implemented"; return nil }

// Get implements Repository.
func (r *FSRepository) Get(name string) (*Skill, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *FSRepository) readDocs(dir string) []Doc { _ = "STUB: not implemented"; return nil }

// parseSummary returns front matter name/description only.
func parseSummary(path string) (Summary, error) {
	_ = "STUB: not implemented"
	return *new(Summary), nil
}

// parseFull returns front matter and the Markdown body.
func parseFull(path string) (Summary, string, error) {
	_ = "STUB: not implemented"
	return *new(Summary), "", nil
}

// readFrontMatter reads YAML front matter block into a simple map.
// It uses gopkg.in/yaml.v3 to correctly handle multi-line block scalars
// (e.g. "description: |-\n  text") that the previous hand-rolled parser missed.
func readFrontMatter(r *bufio.Reader) (map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// splitFrontMatter splits text into a front-matter map and the Markdown body.
// It uses gopkg.in/yaml.v3 to correctly handle multi-line block scalars
// (e.g. "description: |-\n  text") that the previous hand-rolled parser missed.
func splitFrontMatter(text string) (map[string]string, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// No closing delimiter; treat whole as body.

// isBlockScalarIndicator reports whether val (the text after "key: ") is a
// YAML block scalar indicator, meaning the value spans multiple indented lines.
func isBlockScalarIndicator(val string) bool { _ = "STUB: not implemented"; return false }

// parseFrontMatterYAML parses a YAML front-matter block into a flat
// map[string]string using a hybrid strategy:
//
//   - Plain single-line values ("key: some value #with hash") are extracted
//     with a simple line-split so that unquoted '#' characters are not
//     misinterpreted as YAML comments.
//   - Block scalar values ("key: |-\n  line1\n  line2") are collected into a
//     minimal YAML snippet and parsed with gopkg.in/yaml.v3, which correctly
//     handles all block-scalar chomping indicators (|, |-,|+, >, >-, >+).
func parseFrontMatterYAML(src string) map[string]string { _ = "STUB: not implemented"; return nil }

// Skip blank lines at the top level.

// Expect "key: value" or "key:".

// Collect the indicator line plus all indented continuation lines
// into a mini YAML document and let yaml.v3 parse it properly.

// Continuation lines are indented (start with space/tab) or blank.

// Plain single-line value: use the raw text as-is so that '#' and
// other special characters are preserved without YAML interpretation.

func ioReadAll(r *bufio.Reader) (string, error) { _ = "STUB: not implemented"; return "", nil }

func isDocFile(name string) bool { _ = "STUB: not implemented"; return false }

// State keys used for skills.
const (
	StateKeyLoadedPrefix = "temp:skill:loaded:"
	StateKeyDocsPrefix   = "temp:skill:docs:"
	// StateKeyLoadedOrderPrefix stores the per-agent skill touch order
	// used by max-loaded-skills eviction.
	StateKeyLoadedOrderPrefix = "temp:skill:loaded_order:"
	// StateKeyLoadedByAgentPrefix scopes skill-loaded markers by agent name.
	// This prevents a sub-agent's skill_load from leaking into a parent agent's
	// prompt in multi-agent runs that share a Session.
	StateKeyLoadedByAgentPrefix = "temp:skill:loaded_by_agent:"
	// StateKeyDocsByAgentPrefix scopes doc selections by agent name.
	StateKeyDocsByAgentPrefix = "temp:skill:docs_by_agent:"
	// StateKeyLoadedOrderByAgentPrefix scopes the skill touch order by agent
	// name so each agent keeps its own recent-skill window.
	StateKeyLoadedOrderByAgentPrefix = "temp:skill:loaded_order_by_agent:"
	// StateKeyArtifacts stores per-tool-call artifact refs for replay. The value
	// is a JSON object like:
	// {"tool_call_id":"...","artifacts":[{"name":"...","version":3,
	// "ref":"artifact://...@3"}]}
	StateKeyArtifacts = "temp:skill:artifacts"
)
