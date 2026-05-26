//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package debugrecorder provides an opt-in, file-based debug recorder
// for OpenClaw runtime and channels.
package debugrecorder

import (
	"compress/gzip"
	"context"
	"os"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwproto"
)

const (
	modeFull Mode = "full"
	modeSafe Mode = "safe"

	defaultDateDirLayout = "20060102"
	defaultTimeLayout    = "150405"

	defaultTraceDirPerm = 0o700
	defaultFilePerm     = 0o600

	defaultAttachmentsDir = "attachments"
	defaultBySessionDir   = "by-session"

	eventsFileName     = "events.jsonl"
	eventsGzipFileName = eventsFileName + ".gz"
	metaFileName       = "meta.json"
	resultFileName     = "result.json"
	traceRefName       = "trace.json"

	KindTraceStart     = "trace.start"
	KindTraceEnd       = "trace.end"
	KindText           = "text"
	KindError          = "error"
	KindGatewayReq     = "gateway.request"
	KindGatewayRsp     = "gateway.response"
	KindGatewayRun     = "gateway.run.start"
	KindCronRun        = "cron.run.start"
	KindCronDelivery   = "cron.delivery"
	KindRuntimeProfile = "runtime.profile"
	KindModelReq       = "model.chat.request"
	KindRunnerEvent    = "runner.event"

	ProviderOpenAIChatCompletions = "openai.chat.completions"

	KindTelegramMessage    = "telegram.message"
	KindTelegramAttachment = "telegram.attachment"

	errEmptyDir  = "debug recorder: empty dir"
	errEmptyKind = "debug trace: empty kind"

	modelRequestDataURLPrefix   = "data:"
	modelRequestBase64Delimiter = ";base64,"
	modelRequestDefaultMIMEType = "application/octet-stream"
	modelRequestInlineBlobName  = "inline"
	modelRequestFieldBlob       = "blob"
	modelRequestFieldData       = "data"
	modelRequestFieldFileData   = "file_data"
	modelRequestFieldMIMEType   = "mime_type"
	modelRequestFieldURL        = "url"

	maxTraceBaseLen     = 96
	maxSafeComponentLen = 64
	traceSuffixBytes    = 4

	gzipCompressionLevel = gzip.DefaultCompression
)

type Mode string

func ParseMode(raw string) (Mode, error) { _ = "STUB: not implemented"; return *new(Mode), nil }

type Recorder struct {
	dir  string
	mode Mode
	now  func() time.Time
}

func New(dir string, mode Mode) (*Recorder, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Recorder) Dir() string { _ = "STUB: not implemented"; return "" }

func (r *Recorder) Mode() Mode { _ = "STUB: not implemented"; return *new(Mode) }

type TraceStart struct {
	AppName   string `json:"app_name,omitempty"`
	Channel   string `json:"channel,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	Thread    string `json:"thread,omitempty"`
	MessageID string `json:"message_id,omitempty"`
	RequestID string `json:"request_id,omitempty"`
	TraceID   string `json:"trace_id,omitempty"`
	Source    string `json:"source,omitempty"`
}

type TraceEnd struct {
	Status   string        `json:"status,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
	Error    string        `json:"error,omitempty"`
}

type Trace struct {
	root string
	mode Mode

	startedAt time.Time
	metaPath  string
	traceRef  string
	traceID   string

	mu     sync.Mutex
	events *os.File
	closed bool
}

func (r *Recorder) Start(start TraceStart) (*Trace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Recorder) newTraceDir(
	now time.Time,
	start TraceStart,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type traceRef struct {
	TraceDir  string    `json:"trace_dir"`
	StartedAt time.Time `json:"started_at"`
	Channel   string    `json:"channel,omitempty"`
	SessionID string    `json:"session_id,omitempty"`
	RequestID string    `json:"request_id,omitempty"`
	MessageID string    `json:"message_id,omitempty"`
	TraceID   string    `json:"trace_id,omitempty"`
}

func (r *Recorder) writeSessionIndex(
	root string,
	now time.Time,
	start TraceStart,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *Recorder) newSessionIndexDir(
	now time.Time,
	start TraceStart,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func sessionIndexComponent(start TraceStart) string { _ = "STUB: not implemented"; return "" }

func sessionIndexBase(now time.Time, start TraceStart) string { _ = "STUB: not implemented"; return "" }

func (t *Trace) Dir() string { _ = "STUB: not implemented"; return "" }

func (t *Trace) Mode() Mode { _ = "STUB: not implemented"; return *new(Mode) }

type record struct {
	Time    time.Time `json:"time"`
	Kind    string    `json:"kind"`
	Payload any       `json:"payload,omitempty"`
}

func (t *Trace) Record(kind string, payload any) error { _ = "STUB: not implemented"; return nil }

func (t *Trace) RecordText(text string) error { _ = "STUB: not implemented"; return nil }

func (t *Trace) RecordError(err error) error { _ = "STUB: not implemented"; return nil }

type BlobRef struct {
	Ref    string `json:"ref,omitempty"`
	SHA256 string `json:"sha256,omitempty"`
	Size   int    `json:"size"`
	Name   string `json:"name,omitempty"`
}

type RequestSummary struct {
	Channel   string `json:"channel,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	Thread    string `json:"thread,omitempty"`
	MessageID string `json:"message_id,omitempty"`
	Text      string `json:"text,omitempty"`

	RequestSystemPrompt string `json:"request_system_prompt,omitempty"`

	UserID    string `json:"user_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	RequestID string `json:"request_id,omitempty"`

	ContentParts []ContentPartSummary `json:"content_parts,omitempty"`
}

type ContentPartSummary struct {
	Type string `json:"type,omitempty"`

	Text     string                `json:"text,omitempty"`
	Image    *ImagePartSummary     `json:"image,omitempty"`
	Audio    *AudioPartSummary     `json:"audio,omitempty"`
	File     *FilePartSummary      `json:"file,omitempty"`
	Location *gwproto.LocationPart `json:"location,omitempty"`
	Link     *gwproto.LinkPart     `json:"link,omitempty"`
}

type ImagePartSummary struct {
	URL    string  `json:"url,omitempty"`
	Data   BlobRef `json:"data,omitempty"`
	Detail string  `json:"detail,omitempty"`
	Format string  `json:"format,omitempty"`
}

type AudioPartSummary struct {
	URL    string  `json:"url,omitempty"`
	Data   BlobRef `json:"data,omitempty"`
	Format string  `json:"format,omitempty"`
}

type FilePartSummary struct {
	Filename string  `json:"filename,omitempty"`
	Data     BlobRef `json:"data,omitempty"`
	FileID   string  `json:"file_id,omitempty"`
	Format   string  `json:"format,omitempty"`
	URL      string  `json:"url,omitempty"`
}

func SummarizeRequest(
	t *Trace,
	req gwproto.MessageRequest,
) (RequestSummary, error) {
	_ = "STUB: not implemented"
	return *new(RequestSummary), nil
}

func (t *Trace) StoreBlob(name string, data []byte) (BlobRef, error) {
	_ = "STUB: not implemented"
	return *new(BlobRef), nil
}

func (t *Trace) Close(end TraceEnd) error { _ = "STUB: not implemented"; return nil }

func (t *Trace) SetTraceID(traceID string) error { _ = "STUB: not implemented"; return nil }

type traceKey struct{}
type recorderKey struct{}

func WithTrace(ctx context.Context, t *Trace) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func TraceFromContext(ctx context.Context) *Trace { _ = "STUB: not implemented"; return nil }

func WithRecorder(ctx context.Context, r *Recorder) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func RecorderFromContext(ctx context.Context) *Recorder { _ = "STUB: not implemented"; return nil }

type modelRequestRecord struct {
	Provider string `json:"provider,omitempty"`
	Request  any    `json:"request,omitempty"`
}

type inlineDataSummary struct {
	MIMEType string  `json:"mime_type,omitempty"`
	Blob     BlobRef `json:"blob,omitempty"`
}

func RecordModelRequest(
	ctx context.Context,
	provider string,
	payload any,
) error {
	_ = "STUB: not implemented"
	return nil
}

type modelRequestPayloadSanitizer struct {
	trace           *Trace
	inlineBlobCount int
}

func sanitizeModelRequestPayload(
	trace *Trace,
	provider string,
	payload any,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *modelRequestPayloadSanitizer) walk(
	value any,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *modelRequestPayloadSanitizer) replaceInlineData(
	key string,
	value any,
) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func (s *modelRequestPayloadSanitizer) nextInlineBlobName(
	mimeType string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func isInlineDataField(key string) bool { _ = "STUB: not implemented"; return false }

func parseDataURL(raw string) (string, []byte, bool) {
	_ = "STUB: not implemented"
	return "", nil, false
}

func fileExtForMIMEType(mimeType string) string { _ = "STUB: not implemented"; return "" }

func writeJSONFile(path string, v any) error { _ = "STUB: not implemented"; return nil }

func ResolveEventsFilePath(traceDir string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func ReadEventsFile(traceDir string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func compressEventsFile(path string) error { _ = "STUB: not implemented"; return nil }

func verifyGzipFile(path string, wantHash []byte, wantSize int64) error {
	_ = "STUB: not implemented"
	return nil
}

func readGzipFile(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func writeTraceIDJSON(path string, traceID string) error { _ = "STUB: not implemented"; return nil }

func safeComponent(raw string) string { _ = "STUB: not implemented"; return "" }

func randomHex(nBytes int) (string, error) { _ = "STUB: not implemented"; return "", nil }
