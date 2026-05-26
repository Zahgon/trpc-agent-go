//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package skills

import (
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/deps"
)

const (
	skillFileName = "SKILL.md"

	openClawMetadataKey = "openclaw"

	openClawBaseDirPlaceholder = "{baseDir}"
)

var errNoFrontMatter = errors.New("no yaml front matter")

type parsedFrontMatter struct {
	Name        string
	Description string
	Metadata    map[string]any
}

type openClawMetadata struct {
	Always     bool                   `yaml:"always"`
	SkillKey   string                 `yaml:"skillKey"`
	PrimaryEnv string                 `yaml:"primaryEnv"`
	Emoji      string                 `yaml:"emoji"`
	Homepage   string                 `yaml:"homepage"`
	OS         []string               `yaml:"os"`
	Requires   openClawRequires       `yaml:"requires"`
	Install    []openClawInstallEntry `yaml:"install"`
}

type openClawRequires = deps.Requirement

type openClawInstallEntry = deps.InstallAction

func parseFrontMatterFile(path string) (parsedFrontMatter, error) {
	_ = "STUB: not implemented"
	return *new(parsedFrontMatter), nil
}

func parseFrontMatter(content string) (parsedFrontMatter, error) {
	_ = "STUB: not implemented"
	return *new(parsedFrontMatter), nil
}

func parseOpenClawMetadata(
	fm parsedFrontMatter,
) (openClawMetadata, bool, error) {
	_ = "STUB: not implemented"
	return *new(openClawMetadata), false, nil
}

func asString(v any) string { _ = "STUB: not implemented"; return "" }

func normalizeStringAnyMap(v any) map[string]any { _ = "STUB: not implemented"; return nil }

func normalizeMetadata(v any) map[string]any { _ = "STUB: not implemented"; return nil }
