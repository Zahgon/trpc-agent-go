//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package octool

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/memoryfile"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
)

const (
	toolExecCommand = "exec_command"
	toolWriteStdin  = "write_stdin"
	toolKillSession = "kill_session"

	errExecToolNotConfigured  = "exec tool is not configured"
	errCommandRequired        = "command is required"
	errWriteToolNotConfigured = "write_stdin tool is not configured"
	errKillToolNotConfigured  = "kill_session tool is not configured"
	errSessionIDRequired      = "session id is required"

	envSessionUploadsDir = "OPENCLAW_SESSION_UPLOADS_DIR"
	envLastUploadPath    = "OPENCLAW_LAST_UPLOAD_PATH"
	envLastUploadHostRef = "OPENCLAW_LAST_UPLOAD_HOST_REF"
	envLastUploadName    = "OPENCLAW_LAST_UPLOAD_NAME"
	envLastUploadMIME    = "OPENCLAW_LAST_UPLOAD_MIME"
	envRecentUploadsJSON = "OPENCLAW_RECENT_UPLOADS_JSON"

	envLastImagePath    = "OPENCLAW_LAST_IMAGE_PATH"
	envLastImageHostRef = "OPENCLAW_LAST_IMAGE_HOST_REF"
	envLastImageName    = "OPENCLAW_LAST_IMAGE_NAME"
	envLastImageMIME    = "OPENCLAW_LAST_IMAGE_MIME"

	envLastAudioPath    = "OPENCLAW_LAST_AUDIO_PATH"
	envLastAudioHostRef = "OPENCLAW_LAST_AUDIO_HOST_REF"
	envLastAudioName    = "OPENCLAW_LAST_AUDIO_NAME"
	envLastAudioMIME    = "OPENCLAW_LAST_AUDIO_MIME"

	envLastVideoPath    = "OPENCLAW_LAST_VIDEO_PATH"
	envLastVideoHostRef = "OPENCLAW_LAST_VIDEO_HOST_REF"
	envLastVideoName    = "OPENCLAW_LAST_VIDEO_NAME"
	envLastVideoMIME    = "OPENCLAW_LAST_VIDEO_MIME"

	envLastPDFPath    = "OPENCLAW_LAST_PDF_PATH"
	envLastPDFHostRef = "OPENCLAW_LAST_PDF_HOST_REF"
	envLastPDFName    = "OPENCLAW_LAST_PDF_NAME"
	envLastPDFMIME    = "OPENCLAW_LAST_PDF_MIME"

	envMemoryFile = "OPENCLAW_MEMORY_FILE"

	recentUploadsLimit = 6

	execOutputMediaMarker    = "MEDIA:"
	execOutputMediaDirMarker = "MEDIA_DIR:"
	maxExecOutputMarkers     = 16
)

const (
	uploadKindImage = "image"
	uploadKindAudio = "audio"
	uploadKindVideo = "video"
	uploadKindPDF   = "pdf"
	uploadKindFile  = "file"
)

type execUploadMeta struct {
	Name     string `json:"name,omitempty"`
	Path     string `json:"path,omitempty"`
	HostRef  string `json:"host_ref,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Source   string `json:"source,omitempty"`
}

type execTool struct {
	mgr         *Manager
	uploads     *uploads.Store
	memoryStore *memoryfile.Store
}

// NewExecCommandTool creates the canonical host command tool.
func NewExecCommandTool(
	mgr *Manager,
	stores ...*uploads.Store,
) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

// NewExecCommandToolWithMemoryFileStore creates the canonical host command
// tool with file-based memory environment injection.
func NewExecCommandToolWithMemoryFileStore(
	mgr *Manager,
	uploadStore *uploads.Store,
	memoryStore *memoryfile.Store,
) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

func (t *execTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func execToolDescription(hasMemoryFile bool) string { _ = "STUB: not implemented"; return "" }

type execInput struct {
	Command       string            `json:"command"`
	Workdir       string            `json:"workdir,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	YieldTimeMS   *int              `json:"yield_time_ms,omitempty"`
	YieldMs       *int              `json:"yieldMs,omitempty"`
	Background    bool              `json:"background,omitempty"`
	TimeoutSec    *int              `json:"timeout_sec,omitempty"`
	TimeoutSecOld *int              `json:"timeoutSec,omitempty"`
	TTY           *bool             `json:"tty,omitempty"`
	PTY           *bool             `json:"pty,omitempty"`
}

func (t *execTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type writeTool struct {
	mgr *Manager
}

// NewWriteStdinTool creates the stdin continuation tool.
func NewWriteStdinTool(mgr *Manager) tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func (t *writeTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type writeInput struct {
	SessionID     string `json:"session_id,omitempty"`
	SessionIDOld  string `json:"sessionId,omitempty"`
	Chars         string `json:"chars,omitempty"`
	YieldTimeMS   *int   `json:"yield_time_ms,omitempty"`
	YieldMs       *int   `json:"yieldMs,omitempty"`
	AppendNewline *bool  `json:"append_newline,omitempty"`
	Submit        *bool  `json:"submit,omitempty"`
}

func (t *writeTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type killTool struct {
	mgr *Manager
}

// NewKillSessionTool creates the session termination tool.
func NewKillSessionTool(mgr *Manager) tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func (t *killTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type killInput struct {
	SessionID    string `json:"session_id,omitempty"`
	SessionIDOld string `json:"sessionId,omitempty"`
}

func (t *killTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

const (
	defaultWriteYield = 200
)

func mapPollResult(
	sessionID string,
	poll processPoll,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func annotateExecResult(out *execResult) { _ = "STUB: not implemented"; return }

func addExecOutputMarkers(
	out map[string]any,
	output string,
) {
	_ = "STUB: not implemented"
	return
}

func parseExecOutputMarkers(output string) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitExecOutputMarker(line string) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func appendExecOutputMarker(
	out []string,
	seen map[string]struct{},
	path string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func firstInt(values ...*int) *int { _ = "STUB: not implemented"; return nil }

func firstBool(values ...*bool) bool { _ = "STUB: not implemented"; return false }

func resolveWorkdir(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func mergeExecEnv(
	base map[string]string,
	extra map[string]string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func uploadEnvFromContext(
	ctx context.Context,
	store *uploads.Store,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func memoryFileEnvFromContext(
	ctx context.Context,
	store *memoryfile.Store,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func addLatestKindUploadEnv(
	env map[string]string,
	recent []execUploadMeta,
	kind string,
	pathKey string,
	hostRefKey string,
	nameKey string,
	mimeKey string,
) {
	_ = "STUB: not implemented"
	return
}

func latestUploadOfKind(
	recent []execUploadMeta,
	kind string,
) (execUploadMeta, bool) {
	_ = "STUB: not implemented"
	return *new(execUploadMeta), false
}

func recentUploadsFromInvocation(
	inv *agent.Invocation,
	store *uploads.Store,
	limit int,
) []execUploadMeta {
	_ = "STUB: not implemented"
	return nil
}

func appendRecentUploadsFromMessage(
	out []execUploadMeta,
	seen map[string]struct{},
	msg model.Message,
	limit int,
) []execUploadMeta {
	_ = "STUB: not implemented"
	return nil
}

func appendRecentUploadsFromStore(
	out []execUploadMeta,
	seen map[string]struct{},
	store *uploads.Store,
	inv *agent.Invocation,
	limit int,
) []execUploadMeta {
	_ = "STUB: not implemented"
	return nil
}

func uploadScopeFromInvocation(
	inv *agent.Invocation,
) (uploads.Scope, bool) {
	_ = "STUB: not implemented"
	return *new(uploads.Scope), false
}

func uploadChannelFromSessionID(sessionID string) string { _ = "STUB: not implemented"; return "" }

func uploadKindFromMeta(name string, mimeType string) string { _ = "STUB: not implemented"; return "" }

var _ tool.CallableTool = (*execTool)(nil)
var _ tool.CallableTool = (*writeTool)(nil)
var _ tool.CallableTool = (*killTool)(nil)
