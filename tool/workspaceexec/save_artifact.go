//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package workspaceexec

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// SaveArtifactTool persists an existing workspace file as an artifact.
type SaveArtifactTool struct {
	exec *ExecTool
}

type saveArtifactInput struct {
	Path string `json:"path"`
}

type saveArtifactOutput struct {
	Path      string `json:"path"`
	SavedAs   string `json:"saved_as"`
	Version   int    `json:"version"`
	Ref       string `json:"ref"`
	MIMEType  string `json:"mime_type,omitempty"`
	SizeBytes int64  `json:"size_bytes"`
}

type artifactStateRef struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
	Ref     string `json:"ref"`
}

type saveArtifactStateDelta struct {
	ToolCallID string             `json:"tool_call_id"`
	Artifacts  []artifactStateRef `json:"artifacts"`
}

// NewSaveArtifactTool creates a tool for persisting final workspace files.
func NewSaveArtifactTool(exec *ExecTool) *SaveArtifactTool { _ = "STUB: not implemented"; return nil }

// Declaration returns the schema for workspace_save_artifact.
func (t *SaveArtifactTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Call persists an existing workspace file through the artifact service.
func (t *SaveArtifactTool) Call(
	ctx context.Context,
	input []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// StateDelta returns a replayable artifact ref list when workspace_save_artifact
// successfully persists a workspace file via Artifact service.
func (t *SaveArtifactTool) StateDelta(
	toolCallID string,
	_ []byte,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// SupportsArtifactSave reports whether the current invocation can
// persist artifacts for workspace tools. It forwards to
// workspacefacade.ArtifactSaveSkipReasonInv so the predicate stays in
// sync with workspacefacade.ArtifactSaveSkipReason (used by Call).
func SupportsArtifactSave(inv *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

var _ tool.Tool = (*SaveArtifactTool)(nil)
var _ tool.CallableTool = (*SaveArtifactTool)(nil)
