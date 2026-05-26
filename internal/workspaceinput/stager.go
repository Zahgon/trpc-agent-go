//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package workspaceinput stages conversation file inputs into executor
// workspaces so workspace-bound tools can operate on uploaded files without
// exposing staging steps to the model.
package workspaceinput

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	inputFromPrefix = "user_message://"
	inputModePut    = "put"
	inputNameFmt    = "upload_%d"

	keyFileIDPrefix = "file_id/"
	keySHA256Prefix = "sha256/"

	// DefaultName is used when an uploaded file does not carry a usable name.
	DefaultName = "upload"
	// HostPrefix is the trusted host-file prefix used by local runtimes.
	HostPrefix = "host://"

	warnPrefix = "user file input:"
	// WarnMissingRef is returned when a file part contains neither bytes nor a
	// stable reference.
	WarnMissingRef = warnPrefix + " missing bytes and file_id"
	// WarnNoDownloader is returned when the current model cannot dereference a
	// provider-side file id.
	WarnNoDownloader = warnPrefix + " model does not support file download"
	// WarnArtifactNoService is returned when an artifact ref is present but no
	// artifact service is available in the current invocation context.
	WarnArtifactNoService = warnPrefix +
		" artifact service is not configured"
)

// StagedInput describes one conversation file materialized into the workspace.
type StagedInput struct {
	Name         string `json:"name"`
	OriginalName string `json:"original_name,omitempty"`
	MIMEType     string `json:"mime_type,omitempty"`
	SizeBytes    int64  `json:"size_bytes,omitempty"`
}

// StageConversationFiles materializes conversation file inputs into
// work/inputs/ for the current invocation workspace.
func StageConversationFiles(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
) ([]StagedInput, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stageConversationFilesLocked(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	files []model.File,
) ([]StagedInput, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArtifactBaseName extracts a stable basename from an artifact:// ref.
func ArtifactBaseName(fileID string) string { _ = "STUB: not implemented"; return "" }

// SanitizeFileName converts a user-provided filename into a safe basename.
func SanitizeFileName(name string) string { _ = "STUB: not implemented"; return "" }

// UniqueFileName de-duplicates a sanitized upload name under work/inputs.
func UniqueFileName(
	used map[string]struct{},
	existingTo map[string]struct{},
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

// ResolveFileBytes resolves an uploaded file into bytes and a MIME type.
func ResolveFileBytes(
	ctx context.Context,
	mdl model.Model,
	f model.File,
) ([]byte, string, string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

func stageConversationFile(
	ctx context.Context,
	mdl model.Model,
	f model.File,
	idx int,
	usedNames map[string]struct{},
	existingTo map[string]struct{},
	existingByKey map[string]string,
	puts *[]codeexecutor.PutFile,
	md *codeexecutor.WorkspaceMetadata,
) (*StagedInput, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func fastKey(f model.File) (string, bool) { _ = "STUB: not implemented"; return "", false }

func reuseKey(f model.File, sanitizedName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func filesFromSession(sess *session.Session) []model.File { _ = "STUB: not implemented"; return nil }

func filesFromMessage(msg model.Message) []model.File { _ = "STUB: not implemented"; return nil }

func hostPathFromID(fileID string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func hostBytes(hostPath string, f model.File) ([]byte, string, string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

func artifactBytes(
	ctx context.Context,
	fileID string,
) ([]byte, string, string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

func withArtifactContext(
	ctx context.Context,
	inv *agent.Invocation,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
