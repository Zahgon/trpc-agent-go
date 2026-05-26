//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package gateway

import (
	"context"
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwproto"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/debugrecorder"
)

const (
	traceStatusOK       = "ok"
	traceStatusIgnored  = "ignored"
	traceStatusError    = "error"
	traceStatusCanceled = "canceled"

	streamToolExecCommand  = "exec_command"
	streamToolReadDocument = "read_document"
	streamToolReadSheet    = "read_spreadsheet"
	streamToolReadFile     = "fs_read_file"
	streamToolSaveFile     = "fs_save_file"
	streamToolListDir      = "fs_list_dir"
	streamToolSearch       = "fs_search"
	streamToolApplyPatch   = "apply_patch"

	streamToolArgCommand   = "command"
	streamToolArgPath      = "path"
	streamToolArgQuery     = "query"
	streamToolArgPattern   = "pattern"
	streamToolArgSkill     = "skill"
	streamToolArgDocs      = "docs"
	streamToolArgAction    = "action"
	streamToolArgOperation = "operation"
	streamToolArgURL       = "url"
	streamToolArgURLs      = "urls"
	streamToolArgFile      = "file"
	streamToolArgFilePath  = "file_path"
	streamToolArgFileName  = "file_name"
	streamToolArgFilename  = "filename"
	streamToolArgID        = "id"
	streamToolArgName      = "name"
	streamToolArgTitle     = "title"
	streamToolArgTarget    = "target"
	streamToolArgElement   = "element"
	streamToolArgSelector  = "selector"
	streamToolArgRef       = "ref"
	streamToolArgKey       = "key"
	streamToolArgJobID     = "job_id"
	streamToolArgSession   = "session_id"
	streamToolArgRow       = "row"
	streamToolArgStartRow  = "start_row"
	streamToolArgEndRow    = "end_row"
	streamToolArgSheet     = "sheet"

	streamToolDetailRowsPrefix  = "rows "
	streamToolDetailRowPrefix   = "row "
	streamToolDetailPagePrefix  = "page "
	streamToolDetailSheetPrefix = "sheet "
	streamToolDetailJobPrefix   = "job "
	streamToolDetailSessPrefix  = "session "
	streamToolDetailRefPrefix   = "ref "
	streamToolDetailKeyPrefix   = "key "
	streamToolDetailSeparator   = " "
	streamToolDetailMaxRunes    = 96
	streamToolDetailMaxParts    = 2
	streamToolPathEllipsis      = "..."
	streamToolPathSeparator     = "/"

	streamCommandCD        = "cd"
	streamCommandSet       = "set"
	streamCommandExport    = "export"
	streamCommandSource    = "source"
	streamCommandDot       = "."
	streamCommandFor       = "for"
	streamCommandShellLoop = "shell loop"
	streamCommandBash      = "bash"
	streamCommandSh        = "sh"
	streamCommandZsh       = "zsh"
	streamCommandGo        = "go"
	streamCommandGit       = "git"
	streamCommandRG        = "rg"
	streamCommandSed       = "sed"
	streamCommandCat       = "cat"
	streamCommandHead      = "head"
	streamCommandTail      = "tail"
	streamCommandLS        = "ls"
	streamCommandFind      = "find"
	streamCommandPytest    = "pytest"
	streamCommandPython    = "python"
	streamCommandPython3   = "python3"
	streamCommandNPM       = "npm"
	streamCommandPNPM      = "pnpm"
	streamCommandYarn      = "yarn"
	streamCommandBun       = "bun"

	streamCommandModuleFlag = "-m"
	streamCommandShellFlag  = "-c"
	streamCommandShellLogin = "-lc"
	streamCommandArgLimit   = 4

	streamSecretToken         = "token"
	streamSecretSecret        = "secret"
	streamSecretPassword      = "password"
	streamSecretAuthorization = "authorization"
	streamSecretBearer        = "bearer"
	streamSecretAPIKey        = "api_key"
	streamSecretAPIKeyFlat    = "apikey"
	streamSecretCookie        = "cookie"
	streamSecretKeySuffix     = "_key"
	streamSecretOpenAIKey     = "sk-"
	streamSecretGitHubToken   = "ghp_"
	streamSecretTencentToken  = "tgit_"

	progressSummaryPrepare   = "Preparing request"
	progressSummaryDoc       = "Reading document"
	progressSummarySheet     = "Reading spreadsheet"
	progressSummaryTool      = "Running local tool"
	progressSummaryAnswering = "Preparing final answer"
	progressSummaryGoTest    = "Running go test"
	progressSummaryPytest    = "Running pytest"
	progressSummaryNPMTest   = "Running npm test"
	progressSummaryGit       = "Running git command"
	progressSummaryInspect   = "Inspecting workspace"
)

type streamOutcome struct {
	status string
	errMsg string
}

type progressState struct {
	startedAt  time.Time
	stage      gwproto.StreamProgressStage
	summary    string
	toolName   string
	toolDetail string
	toolCallID string
	toolStatus gwproto.StreamToolStatus
}

type streamToolArgKind int

const (
	streamToolArgKindText streamToolArgKind = iota
	streamToolArgKindPath
	streamToolArgKindURL
)

type streamToolDetailArg struct {
	key    string
	prefix string
	kind   streamToolArgKind
}

var streamGenericToolDetailArgs = []streamToolDetailArg{
	{key: streamToolArgSkill},
	{key: streamToolArgDocs, kind: streamToolArgKindPath},
	{key: streamToolArgAction},
	{key: streamToolArgOperation},
	{key: streamToolArgQuery},
	{key: streamToolArgPattern},
	{key: streamToolArgURL, kind: streamToolArgKindURL},
	{key: streamToolArgURLs, kind: streamToolArgKindURL},
	{key: streamToolArgPath, kind: streamToolArgKindPath},
	{key: streamToolArgFilePath, kind: streamToolArgKindPath},
	{key: streamToolArgFileName, kind: streamToolArgKindPath},
	{key: streamToolArgFile, kind: streamToolArgKindPath},
	{key: streamToolArgFilename, kind: streamToolArgKindPath},
	{key: streamToolArgID},
	{key: streamToolArgName},
	{key: streamToolArgTitle},
	{key: streamToolArgTarget},
	{key: streamToolArgElement},
	{key: streamToolArgSelector},
	{key: streamToolArgRef, prefix: streamToolDetailRefPrefix},
	{key: streamToolArgKey, prefix: streamToolDetailKeyPrefix},
	{key: streamToolArgJobID, prefix: streamToolDetailJobPrefix},
	{key: streamToolArgSession, prefix: streamToolDetailSessPrefix},
}

// StreamMessage processes a request and returns a stream of gateway
// events. Validation errors are returned as APIError/status pairs.
func (s *Server) StreamMessage(
	ctx context.Context,
	req gwproto.MessageRequest,
) (<-chan gwproto.StreamEvent, *gwproto.APIError, int) {
	_ = "STUB: not implemented"
	return nil, nil, 0
}

// StreamMessageWithOptions processes a request with opt-in stream behavior
// controls and returns a stream of gateway events.
func (s *Server) StreamMessageWithOptions(
	ctx context.Context,
	req gwproto.MessageRequest,
	opts *gwproto.MessageStreamOptions,
) (<-chan gwproto.StreamEvent, *gwproto.APIError, int) {
	_ = "STUB: not implemented"
	return nil, nil, 0
}

func (s *Server) streamMessage(
	ctx context.Context,
	req gwproto.MessageRequest,
	opts *gwproto.MessageStreamOptions,
) (<-chan gwproto.StreamEvent, *gwproto.APIError, int) {
	_ = "STUB: not implemented"
	return nil, nil, 0
}

func (s *Server) handleMessagesStream(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

type streamMessageRequest struct {
	gwproto.MessageRequest
	StreamOptions *gwproto.MessageStreamOptions `json:"stream_options,omitempty"`
}

func writeSSEEvent(
	w http.ResponseWriter,
	flusher http.Flusher,
	evt gwproto.StreamEvent,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Server) streamLocked(
	ctx context.Context,
	run preparedMessageRun,
	trace *debugrecorder.Trace,
	out chan<- gwproto.StreamEvent,
) streamOutcome {
	_ = "STUB: not implemented"
	return *new(streamOutcome)
}

func shouldSendProgressForEvent(
	opts *gwproto.MessageStreamOptions,
	sentText bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func sendStreamEvent(
	ctx context.Context,
	out chan<- gwproto.StreamEvent,
	evt gwproto.StreamEvent,
) bool {
	_ = "STUB: not implemented"
	return false
}

type progressUpdate struct {
	stage      gwproto.StreamProgressStage
	summary    string
	toolName   string
	toolDetail string
	toolCallID string
	toolStatus gwproto.StreamToolStatus
}

func progressUpdateFromRunnerEvent(
	evt *event.Event,
) (progressUpdate, bool) {
	_ = "STUB: not implemented"
	return *new(progressUpdate), false
}

func firstToolCall(rsp *model.Response) (model.ToolCall, bool) {
	_ = "STUB: not implemented"
	return *new(model.ToolCall), false
}

func firstToolResult(rsp *model.Response) (model.Message, bool) {
	_ = "STUB: not implemented"
	return *new(model.Message), false
}

func progressFromToolCall(
	toolCall model.ToolCall,
) (progressUpdate, bool) {
	_ = "STUB: not implemented"
	return *new(progressUpdate), false
}

func progressFromToolResult(rsp *model.Response) progressUpdate {
	_ = "STUB: not implemented"
	return *new(progressUpdate)
}

func toolDetailFromToolCall(toolCall model.ToolCall) string { _ = "STUB: not implemented"; return "" }

func execCommandToolDetail(toolCall model.ToolCall) string { _ = "STUB: not implemented"; return "" }

func readDocumentToolDetail(toolCall model.ToolCall) string { _ = "STUB: not implemented"; return "" }

func readSpreadsheetToolDetail(toolCall model.ToolCall) string {
	_ = "STUB: not implemented"
	return ""
}

func toolPathDetail(toolCall model.ToolCall) string { _ = "STUB: not implemented"; return "" }

func searchToolDetail(toolCall model.ToolCall) string { _ = "STUB: not implemented"; return "" }

func genericToolDetail(toolCall model.ToolCall) string { _ = "STUB: not implemented"; return "" }

func toolDetailArgValue(kind streamToolArgKind, value string) string {
	_ = "STUB: not implemented"
	return ""
}

func containsToolDetailPart(parts []string, part string) bool {
	_ = "STUB: not implemented"
	return false
}

func stringArgFromToolCall(
	toolCall model.ToolCall,
	key string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func toolCallArgs(toolCall model.ToolCall) (map[string]any, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func stringArgFromMap(args map[string]any, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func stringValuesFromMap(args map[string]any, key string) []string {
	_ = "STUB: not implemented"
	return nil
}

func stringValuesFromList(values []any) []string { _ = "STUB: not implemented"; return nil }

func execCommandProgressSummary(toolCall model.ToolCall) string {
	_ = "STUB: not implemented"
	return ""
}

func normalizeExecCommand(command string) string { _ = "STUB: not implemented"; return "" }

func looksLikeWorkspaceInspection(command string) bool { _ = "STUB: not implemented"; return false }

func shellCommandDetail(command string) string { _ = "STUB: not implemented"; return "" }

func shellCommandSegments(command string) []string { _ = "STUB: not implemented"; return nil }

func shellCommandSegmentDetail(segment string) string { _ = "STUB: not implemented"; return "" }

func splitShellSegments(command string) []string { _ = "STUB: not implemented"; return nil }

func appendShellSegment(segments []string, segment string) []string {
	_ = "STUB: not implemented"
	return nil
}

func shellFields(segment string) []string { _ = "STUB: not implemented"; return nil }

func appendShellField(fields []string, field string) []string {
	_ = "STUB: not implemented"
	return nil
}

func trimShellCommandPrefix(fields []string) []string { _ = "STUB: not implemented"; return nil }

func isShellEnvAssignment(field string) bool { _ = "STUB: not implemented"; return false }

func cleanCommandFields(segment string) []string { _ = "STUB: not implemented"; return nil }

func safeCommandName(name string) string { _ = "STUB: not implemented"; return "" }

func skipShellSetupCommand(command string) bool { _ = "STUB: not implemented"; return false }

func commandWithSafeArgs(
	command string,
	args []string,
	limit int,
) string {
	_ = "STUB: not implemented"
	return ""
}

func isSensitiveCommandFlag(arg string) bool { _ = "STUB: not implemented"; return false }

func isSensitiveCommandKeyToken(arg string) bool { _ = "STUB: not implemented"; return false }

func commandWithPathArg(command string, args []string) string { _ = "STUB: not implemented"; return "" }

func shellWrapperCommandDetail(args []string) string { _ = "STUB: not implemented"; return "" }

func pythonCommandDetail(command string, args []string) string {
	_ = "STUB: not implemented"
	return ""
}

func packageCommandDetail(command string, args []string) string {
	_ = "STUB: not implemented"
	return ""
}

func safeCommandArgDetail(arg string) string { _ = "STUB: not implemented"; return "" }

func lastSafePathArg(args []string) string { _ = "STUB: not implemented"; return "" }

func safePathDetail(path string) string { _ = "STUB: not implemented"; return "" }

func safeURLDetail(rawURL string) string { _ = "STUB: not implemented"; return "" }

func safeDetailToken(value string) string { _ = "STUB: not implemented"; return "" }

func sanitizeStreamToolDetail(detail string) string { _ = "STUB: not implemented"; return "" }

func isSafeToolDetailRune(char rune) bool { _ = "STUB: not implemented"; return false }

func looksSensitiveCommandArg(arg string) bool { _ = "STUB: not implemented"; return false }

func isSensitiveToolArgKey(key string) bool { _ = "STUB: not implemented"; return false }

func normalizeSensitiveKey(key string) string { _ = "STUB: not implemented"; return "" }

func looksSensitiveValue(value string) bool { _ = "STUB: not implemented"; return false }

func looksLikeJWT(value string) bool { _ = "STUB: not implemented"; return false }

func isBase64URLLike(value string) bool { _ = "STUB: not implemented"; return false }

func readDocumentProgressSummary(toolCall model.ToolCall) string {
	_ = "STUB: not implemented"
	return ""
}

func readSpreadsheetProgressSummary(toolCall model.ToolCall) string {
	_ = "STUB: not implemented"
	return ""
}

func sendProgressUpdate(
	ctx context.Context,
	out chan<- gwproto.StreamEvent,
	run preparedMessageRun,
	state *progressState,
	update progressUpdate,
) bool {
	_ = "STUB: not implemented"
	return false
}

func singleStreamEvents(
	events ...gwproto.StreamEvent,
) <-chan gwproto.StreamEvent {
	_ = "STUB: not implemented"
	return nil
}

func contextErrMessage(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func resolvedStreamRequestID(
	requestID string,
	fallback string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func apiErrorFromEvent(evt *event.Event) *gwproto.APIError { _ = "STUB: not implemented"; return nil }

func streamDeltaText(
	evt *event.Event,
	sentText bool,
) string {
	_ = "STUB: not implemented"
	return ""
}

func streamPublicDelta(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func streamPublicCompleted(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func shouldSendPublicCompleted(
	evt *event.Event,
	publicReply string,
	lastPublicCompleted string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func streamThoughtDelta(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func streamThoughtCompleted(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func shouldSendThoughtCompleted(
	evt *event.Event,
	thoughtReply string,
	pendingThought bool,
	lastThoughtCompleted string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func fullTextFromResponse(rsp *model.Response) string { _ = "STUB: not implemented"; return "" }

func fullPublicFromResponse(rsp *model.Response) string { _ = "STUB: not implemented"; return "" }

func fullThoughtFromResponse(rsp *model.Response) string { _ = "STUB: not implemented"; return "" }

func deltaTextFromResponse(rsp *model.Response) string { _ = "STUB: not implemented"; return "" }

func deltaPublicFromResponse(rsp *model.Response) string { _ = "STUB: not implemented"; return "" }

func deltaThoughtFromResponse(rsp *model.Response) string { _ = "STUB: not implemented"; return "" }

func responseHasPublicContent(rsp *model.Response) bool { _ = "STUB: not implemented"; return false }
