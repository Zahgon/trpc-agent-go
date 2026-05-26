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

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
)

const (
	recentUploadContextLimit = 6

	recentUploadContextHeader = "Recent chat files/media available " +
		"to tools in this session (newest first):"
	recentUploadKindHeader = "Latest matching file by kind " +
		"in this chat:"
	recentUploadContextFooter = "Use OPENCLAW_LAST_UPLOAD_* " +
		"for the newest file. For multiple recent files in " +
		"exec_command, use OPENCLAW_RECENT_UPLOADS_JSON. " +
		"For common file-reading tasks, prefer read_document " +
		"or read_spreadsheet before exec_command. If a single " +
		"matching upload is already present, call those tools " +
		"directly instead of printing OPENCLAW_* variables. " +
		"When calling message to send an existing upload back " +
		"to the user, prefer the stable host_ref from that " +
		"JSON or OPENCLAW_LAST_UPLOAD_HOST_REF instead of " +
		"guessing a local path. " +
		"This list may include both user uploads and bot-" +
		"generated files previously sent in this chat. " +
		"When the user says 'the PDF/audio/video I just " +
		"sent' or 'the file you just sent me', resolve " +
		"against this list first. If the " +
		"target file came from a recent bot reply, it is " +
		"still valid to reuse it. If the " +
		"user replies to an earlier media message, that " +
		"replied media is usually the intended target. If the " +
		"agent wants OpenClaw chat channels to auto-attach " +
		"generated outputs, it may include final reply lines like " +
		"`MEDIA: /path/to/file` or `MEDIA_DIR: /path/to/dir`; " +
		"those directive lines are hidden from users while the " +
		"matching files are sent. To send compatible audio as a " +
		"Telegram voice bubble, it may also include " +
		"`[[audio_as_voice]]` in the final reply. " +
		"Avoid repeating local machine " +
		"paths or placeholder Telegram filenames " +
		"in user-facing replies unless the user explicitly asks. " +
		"If the " +
		"requested media kind is not present here, say " +
		"which files are currently available in this chat."
)

func (s *Server) uploadContextMessages(
	ctx context.Context,
	userID string,
	sessionID string,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func buildUploadContextText(files []uploads.ListedFile) string {
	_ = "STUB: not implemented"
	return ""
}

func buildUploadKindSummary(files []uploads.ListedFile) string {
	_ = "STUB: not implemented"
	return ""
}

const uploadKindFileLabel = "file"

func formatUploadContextLine(file uploads.ListedFile) string { _ = "STUB: not implemented"; return "" }

func uploadContextName(file uploads.ListedFile) string { _ = "STUB: not implemented"; return "" }

func isUploadPlaceholderName(name string) bool { _ = "STUB: not implemented"; return false }

func describeUploadKind(name string, mimeType string) string { _ = "STUB: not implemented"; return "" }

func describeUploadSource(source string) string { _ = "STUB: not implemented"; return "" }

func channelFromSessionID(sessionID string) string { _ = "STUB: not implemented"; return "" }
