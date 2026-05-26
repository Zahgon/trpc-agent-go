//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/admin"
)

const (
	adminPromptInstructionBundle = "agent_instruction"
	adminPromptSystemBundle      = "agent_system"

	adminPromptFilePerm = 0o600
	adminPromptDirPerm  = 0o700
	adminPromptFileExt  = ".md"
)

type adminPromptProvider struct {
	mu sync.RWMutex

	cwd string

	opts       runOptions
	controller *RuntimePromptController

	instructionOverride *string
	systemOverride      *string
}

func buildAdminPromptProvider(
	opts runOptions,
	controller *RuntimePromptController,
) admin.PromptsProvider {
	_ = "STUB: not implemented"
	return *new(admin.PromptsProvider)
}

func (p *adminPromptProvider) PromptsStatus() (
	admin.PromptsStatus,
	error,
) {
	_ = "STUB: not implemented"
	return *new(admin.PromptsStatus), nil
}

func (p *adminPromptProvider) SavePromptRuntime(
	bundleKey string,
	content string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *adminPromptProvider) SavePromptInline(
	bundleKey string,
	content string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *adminPromptProvider) SavePromptFile(
	bundleKey string,
	path string,
	content string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *adminPromptProvider) CreatePromptFile(
	bundleKey string,
	fileName string,
	content string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *adminPromptProvider) DeletePromptFile(
	bundleKey string,
	path string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *adminPromptProvider) bundleStateLocked(
	key string,
	title string,
	effective string,
	override *string,
) admin.PromptBundleState {
	_ = "STUB: not implemented"
	return *new(admin.PromptBundleState)
}

func (p *adminPromptProvider) bundleConfiguredPromptLocked(
	key string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *adminPromptProvider) bundleFilesLocked(
	key string,
) ([]admin.PromptFileState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *adminPromptProvider) lookupPromptFileLocked(
	bundleKey string,
	path string,
) (admin.PromptFileState, bool, error) {
	_ = "STUB: not implemented"
	return *new(admin.PromptFileState), false, nil
}

func (p *adminPromptProvider) bundleCreateDirLocked(
	key string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *adminPromptProvider) bundleCreateDirValueLocked(
	key string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *adminPromptProvider) bundleResolvedPathsLocked(
	key string,
) ([]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (p *adminPromptProvider) resolvePromptPaths(
	rawFiles []string,
	rawDir string,
) ([]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (p *adminPromptProvider) resolvePromptPath(raw string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *adminPromptProvider) applyLocked() error { _ = "STUB: not implemented"; return nil }

func buildAdminPromptPreview(
	instruction string,
	systemPrompt string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func adminPromptPreviewBlock(
	title string,
	content string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func adminPromptBundleSummary(key string) string { _ = "STUB: not implemented"; return "" }

func adminPromptConfiguredLabel(key string) string { _ = "STUB: not implemented"; return "" }

func adminPromptEffectiveLabel(key string) string { _ = "STUB: not implemented"; return "" }

func adminPromptSourceSummary(key string, fileCount int) string {
	_ = "STUB: not implemented"
	return ""
}

func displayAdminPromptFileLabel(path string) string { _ = "STUB: not implemented"; return "" }

func promptOverridePtr(content string) *string { _ = "STUB: not implemented"; return nil }

func writeAdminPromptFile(path string, content string) error { _ = "STUB: not implemented"; return nil }

func normalizeAdminPromptFileName(raw string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
