//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package gateway

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwproto"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
)

const (
	defaultMaxContentPartBytes int64 = 8 << 20

	defaultContentPartTimeout = 15 * time.Second
	defaultMaxRedirects       = 5

	headerUserAgent = "User-Agent"

	contentPartUserAgent = "trpc-agent-go/openclaw-gateway"

	imageDetailAuto = "auto"

	mimeOctetStream = "application/octet-stream"

	audioFormatWAV = "wav"
	audioFormatMP3 = "mp3"

	errContentPartTooLarge = "content part too large"

	telegramChannelName = "telegram"

	audioTranscriptSeparator = "\n\n"
	audioTranscriptLabelFmt  = "Audio transcript %d:"

	kindGatewayAudioTranscript = "gateway.audio.transcription"
)

type partFetcher interface {
	Fetch(ctx context.Context, rawURL string, maxBytes int64) (fetched, error)
}

type fetched struct {
	Data        []byte
	ContentType string
	Filename    string
}

type partURLPolicy struct {
	allowPrivate    bool
	allowedPatterns []string
	resolver        hostResolver
}

type hostResolver interface {
	LookupIPAddr(ctx context.Context, host string) ([]net.IPAddr, error)
}

func (p partURLPolicy) Validate(
	ctx context.Context,
	u *url.URL,
) error {
	_ = "STUB: not implemented"
	return nil
}

func matchesAnyPattern(u *url.URL, patterns []string) bool { _ = "STUB: not implemented"; return false }

func matchPattern(u *url.URL, pattern string) bool { _ = "STUB: not implemented"; return false }

func matchHost(hostname, target string) bool { _ = "STUB: not implemented"; return false }

func validatePublicHost(
	ctx context.Context,
	host string,
	resolver hostResolver,
) error {
	_ = "STUB: not implemented"
	return nil
}

func isPrivateOrLocalIP(addr netip.Addr) bool { _ = "STUB: not implemented"; return false }

type validatingFetcher struct {
	next   partFetcher
	policy partURLPolicy
}

func (f validatingFetcher) Fetch(
	ctx context.Context,
	rawURL string,
	maxBytes int64,
) (fetched, error) {
	_ = "STUB: not implemented"
	return *new(fetched), nil
}

type urlPartFetcher struct {
	client       *http.Client
	maxRedirects int
	policy       partURLPolicy
}

func newURLPartFetcher(policy partURLPolicy) *urlPartFetcher { _ = "STUB: not implemented"; return nil }

func (f *urlPartFetcher) Fetch(
	ctx context.Context,
	rawURL string,
	maxBytes int64,
) (fetched, error) {
	_ = "STUB: not implemented"
	return *new(fetched), nil
}

func normalizeContentType(raw string) string { _ = "STUB: not implemented"; return "" }

func filenameFromHeaders(resp *http.Response, u *url.URL) string {
	_ = "STUB: not implemented"
	return ""
}

func readLimited(r io.Reader, maxBytes int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inboundFromRequest(
	req gwproto.MessageRequest,
	text string,
) InboundMessage {
	_ = "STUB: not implemented"
	return *new(InboundMessage)
}

func (s *Server) normalizeUserMessage(
	ctx context.Context,
	req gwproto.MessageRequest,
) (model.Message, string, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), "", nil
}

func joinText(a, b string) string { _ = "STUB: not implemented"; return "" }

func (s *Server) normalizeContentParts(
	ctx context.Context,
	parts []gwproto.ContentPart,
	scopes ...uploads.Scope,
) ([]model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *Server) rewriteAudioMessage(
	ctx context.Context,
	channel string,
	text string,
	parts []model.ContentPart,
) (string, []model.ContentPart, string, string) {
	_ = "STUB: not implemented"
	return "", nil, "", ""
}

func (s *Server) transcribeAudioParts(
	ctx context.Context,
	parts []model.ContentPart,
) ([]string, []model.ContentPart) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) transcribeAudioPart(
	ctx context.Context,
	index int,
	audio *model.Audio,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func formatAudioTranscriptText(transcripts []string) string { _ = "STUB: not implemented"; return "" }

func traceAudioTranscription(
	ctx context.Context,
	index int,
	audio *model.Audio,
	transcript string,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) normalizeContentPart(
	ctx context.Context,
	part gwproto.ContentPart,
	scopes ...uploads.Scope,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func normalizeTextPart(
	part gwproto.ContentPart,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func normalizeImagePart(
	part gwproto.ContentPart,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *Server) normalizeAudioPart(
	ctx context.Context,
	part gwproto.ContentPart,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *Server) normalizeAudioURL(
	ctx context.Context,
	audio *gwproto.AudioPart,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func inferAudioFormat(f fetched) string { _ = "STUB: not implemented"; return "" }

func isSupportedAudioFormat(format string) bool { _ = "STUB: not implemented"; return false }

func (s *Server) normalizeFilePart(
	ctx context.Context,
	part gwproto.ContentPart,
	scope uploads.Scope,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func normalizeFileID(file *gwproto.FilePart) *model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) normalizeFileURL(
	ctx context.Context,
	file *gwproto.FilePart,
	scope uploads.Scope,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *Server) normalizeFileData(
	ctx context.Context,
	file *gwproto.FilePart,
	scope uploads.Scope,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *Server) persistNormalizedFile(
	ctx context.Context,
	scope uploads.Scope,
	name string,
	mimeType string,
	data []byte,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func uploadScopeFromRequest(req gwproto.MessageRequest) uploads.Scope {
	_ = "STUB: not implemented"
	return *new(uploads.Scope)
}

func firstUploadScope(scopes ...uploads.Scope) uploads.Scope {
	_ = "STUB: not implemented"
	return *new(uploads.Scope)
}

func inferMimeTypeFromName(name string) string { _ = "STUB: not implemented"; return "" }

func normalizeLinkPart(
	part gwproto.ContentPart,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func normalizeLocationPart(
	part gwproto.ContentPart,
) (*model.ContentPart, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *Server) fetchContentPart(
	ctx context.Context,
	rawURL string,
) (fetched, error) {
	_ = "STUB: not implemented"
	return *new(fetched), nil
}
