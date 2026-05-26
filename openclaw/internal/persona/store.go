//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package persona

import (
	"context"
	"errors"
	"sync"
)

const (
	storeVersion = 1

	storeDirName  = "persona"
	storeFileName = "presets.json"
	storeTempExt  = ".tmp"

	storeDirPerm  = 0o700
	storeFilePerm = 0o600

	scopeKindDM     = "dm"
	scopeKindThread = "thread"
	scopeSep        = ":"

	PresetDefault    = "default"
	PresetGirlfriend = "girlfriend"
	PresetConcise    = "concise"
	PresetCoach      = "coach"
	PresetCreative   = "creative"

	personaTaskCompletionPrompt = " Keep this persona subordinate to " +
		"task completion. Do not answer only with what you will " +
		"do next when the requested content or action can be " +
		"completed now."
)

var ErrUnknownPreset = errors.New("persona: unknown preset")

type Preset struct {
	ID          string
	Name        string
	Description string
	Prompt      string
}

var presetList = []Preset{
	{
		ID:          PresetDefault,
		Name:        "Default",
		Description: "Use the normal assistant behavior.",
	},
	{
		ID:          PresetGirlfriend,
		Name:        "Girlfriend",
		Description: "Warm, playful, and affectionate companion tone.",
		Prompt: "Adopt a warm, playful, and affectionate companion " +
			"tone. Prefer natural, low-pressure language. Light " +
			"flirting is fine when the user's tone welcomes it. " +
			"Stay respectful, emotionally grounded, and honest. " +
			"Do not claim real-world exclusivity, dependency, or " +
			"obligations." + personaTaskCompletionPrompt,
	},
	{
		ID:          PresetConcise,
		Name:        "Concise",
		Description: "Direct, brief, and action-first replies.",
		Prompt: "Be direct, brief, and low-friction. Lead with the " +
			"answer or concrete action. Keep wording tight unless " +
			"the user explicitly asks for depth." +
			personaTaskCompletionPrompt,
	},
	{
		ID:          PresetCoach,
		Name:        "Coach",
		Description: "Structured, pragmatic, and goal-oriented.",
		Prompt: "Act like a pragmatic coach. Give clear structure, " +
			"challenge vague thinking, and convert goals into " +
			"concrete next steps. Stay supportive but do not " +
			"sugarcoat tradeoffs." + personaTaskCompletionPrompt,
	},
	{
		ID:          PresetCreative,
		Name:        "Creative",
		Description: "More imaginative, vivid, and idea-rich.",
		Prompt: "Lean imaginative, vivid, and idea-rich. Offer " +
			"varied angles, names, examples, and alternatives " +
			"when it helps the task." + personaTaskCompletionPrompt,
	},
}

var presetAliases = map[string]string{
	"":      PresetDefault,
	"none":  PresetDefault,
	"off":   PresetDefault,
	"reset": PresetDefault,
	"gf":    PresetGirlfriend,
}

type Store struct {
	path string

	mu    sync.Mutex
	state storeState
}

type storeState struct {
	Version int               `json:"version"`
	Scopes  map[string]string `json:"scopes,omitempty"`
}

func DefaultStorePath(stateDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func NewStore(path string) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

func List() []Preset { _ = "STUB: not implemented"; return nil }

func Lookup(id string) (Preset, bool) { _ = "STUB: not implemented"; return *new(Preset), false }

func DefaultPreset() Preset { _ = "STUB: not implemented"; return *new(Preset) }

func DMScopeKey(channel string, userID string) string { _ = "STUB: not implemented"; return "" }

func ThreadScopeKey(channel string, thread string) string { _ = "STUB: not implemented"; return "" }

func ScopeKeyFromSession(
	channel string,
	userID string,
	sessionID string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *Store) Get(scopeKey string) (Preset, error) {
	_ = "STUB: not implemented"
	return *new(Preset), nil
}

func (s *Store) Set(
	ctx context.Context,
	scopeKey string,
	presetID string,
) (Preset, error) {
	_ = "STUB: not implemented"
	return *new(Preset), nil
}

func (s *Store) ForgetUser(
	ctx context.Context,
	channel string,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) load() error { _ = "STUB: not implemented"; return nil }

func (s *Store) persistLocked(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func buildScopeKey(
	channel string,
	kind string,
	id string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func normalizeChannel(channel string) string { _ = "STUB: not implemented"; return "" }

func normalizeScopeKey(scopeKey string) string { _ = "STUB: not implemented"; return "" }

func normalizePresetID(id string) string { _ = "STUB: not implemented"; return "" }

func contextErr(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
