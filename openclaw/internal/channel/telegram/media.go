//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telegram

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwclient"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwproto"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/debugrecorder"
	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

const (
	downloadFailedMessage = "Failed to download attachment."
	attachmentTooLargeMsg = "Attachment is too large to process."
	unsupportedMediaMsg   = "Unsupported attachment format."

	defaultAttachmentName = "attachment"
	defaultDocumentName   = "document"
	defaultPhotoName      = "photo"
	defaultVoiceName      = "voice"
	defaultAudioName      = "audio"
	defaultVideoName      = "video"
	defaultAnimationName  = "animation"
	defaultVideoNoteName  = "video-note"

	audioFormatWAV = "wav"
	audioFormatMP3 = "mp3"

	mimeImageJPEG = "image/jpeg"
	mimeImagePNG  = "image/png"
	mimeImageWEBP = "image/webp"

	mimeAudioMPEG = "audio/mpeg"
	mimeAudioWAV  = "audio/wav"
	mimeAudioOGG  = "audio/ogg"
	mimeVideoMP4  = "video/mp4"
	mimeImageGIF  = "image/gif"

	bytesPerKiB int64 = 1 << 10
	bytesPerMiB int64 = 1 << 20
)

const (
	attachmentKindPhoto     = "photo"
	attachmentKindDocument  = "document"
	attachmentKindVideo     = "video"
	attachmentKindVoice     = "voice"
	attachmentKindAudio     = "audio"
	attachmentKindAnimation = "animation"
	attachmentKindVideoNote = "video_note"
)

type userError struct {
	userMessage string
	err         error
}

func (e *userError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *userError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (c *Channel) buildGatewayRequest(
	ctx context.Context,
	fromID string,
	thread string,
	sessionID string,
	requestID string,
	msg tgapi.Message,
) (gwclient.MessageRequest, error) {
	_ = "STUB: not implemented"
	return *new(gwclient.MessageRequest), nil
}

func (c *Channel) appendMessageParts(
	ctx context.Context,
	parts []gwproto.ContentPart,
	msg *tgapi.Message,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasTelegramAttachments(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func replyHasPhoto(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func replyHasDocument(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func replyHasAudio(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func replyHasVoice(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func replyHasVideo(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func replyHasAnimation(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func replyHasVideoNote(msg *tgapi.Message) bool { _ = "STUB: not implemented"; return false }

func incomingReplyTo(msg tgapi.Message) int { _ = "STUB: not implemented"; return 0 }

func joinMessageText(a, b string) string { _ = "STUB: not implemented"; return "" }

func (c *Channel) appendPhotoPart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	photos []tgapi.PhotoSize,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) appendDocumentPart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	doc *tgapi.Document,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) appendVideoPart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	video *tgapi.Video,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) appendAnimationPart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	animation *tgapi.Animation,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) appendVideoNotePart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	videoNote *tgapi.VideoNote,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) appendNamedVideoLikePart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	kind string,
	fileID string,
	fileName string,
	mimeType string,
	fileSize int64,
	fallbackName string,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) appendVoicePart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	voice *tgapi.Voice,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) appendAudioPart(
	ctx context.Context,
	parts []gwproto.ContentPart,
	audio *tgapi.Audio,
	maxBytes int64,
) ([]gwproto.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapDownloadError(err error, maxBytes int64) error { _ = "STUB: not implemented"; return nil }

func attachmentTooLargeMessage(maxBytes int64) string { _ = "STUB: not implemented"; return "" }

func formatByteLimit(maxBytes int64) string { _ = "STUB: not implemented"; return "" }

func appendStoredFilePart(
	parts []gwproto.ContentPart,
	name string,
	mimeType string,
	data []byte,
) []gwproto.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func mimeTypeForImageFormat(format string) string { _ = "STUB: not implemented"; return "" }

func mimeTypeForAudioFormat(format string) string { _ = "STUB: not implemented"; return "" }

func isSupportedAudioFormat(format string) bool { _ = "STUB: not implemented"; return false }

type telegramAttachmentSummary struct {
	Kind         string                `json:"kind,omitempty"`
	FileID       string                `json:"file_id,omitempty"`
	Name         string                `json:"name,omitempty"`
	Format       string                `json:"format,omitempty"`
	MimeType     string                `json:"mime_type,omitempty"`
	ReportedSize int64                 `json:"reported_size,omitempty"`
	UserMessage  string                `json:"user_message,omitempty"`
	Error        string                `json:"error,omitempty"`
	Blob         debugrecorder.BlobRef `json:"blob,omitempty"`
}

func recordAttachment(
	trace *debugrecorder.Trace,
	summary telegramAttachmentSummary,
) {
	_ = "STUB: not implemented"
	return
}

func storeBlob(
	trace *debugrecorder.Trace,
	name string,
	data []byte,
) debugrecorder.BlobRef {
	_ = "STUB: not implemented"
	return *new(debugrecorder.BlobRef)
}

func userMessageFromErr(err error) string { _ = "STUB: not implemented"; return "" }

func fallbackFilename(primary, filePath, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

func fallbackMediaFilename(
	primary string,
	filePath string,
	fallback string,
	mimeType string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func fallbackDocumentFilename(
	primary string,
	filePath string,
	mimeType string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func documentFallbackBase(filePath string, mimeType string) string {
	_ = "STUB: not implemented"
	return ""
}

func looksGeneratedTelegramFileName(name string) bool { _ = "STUB: not implemented"; return false }

func mediaExtFromPathOrMIME(filePath, mimeType string) string { _ = "STUB: not implemented"; return "" }

func normalizeMediaMIME(
	fileName string,
	filePath string,
	mimeType string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func isVideoMedia(
	fileName string,
	filePath string,
	mimeType string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func inferImageFormat(filePath string, data []byte) string { _ = "STUB: not implemented"; return "" }

func imageFormatFromExt(ext string) string { _ = "STUB: not implemented"; return "" }

func imageFormatFromContentType(contentType string) string { _ = "STUB: not implemented"; return "" }

func inferAudioFormat(filename, filePath, mimeType string) string {
	_ = "STUB: not implemented"
	return ""
}

func audioFormatFromExt(ext string) string { _ = "STUB: not implemented"; return "" }

func audioFormatFromMimeType(mimeType string) string { _ = "STUB: not implemented"; return "" }
