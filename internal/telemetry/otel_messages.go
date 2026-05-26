//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telemetry

import (
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	otelPartTypeBlob             = "blob"
	otelPartTypeFile             = "file"
	otelPartTypeReasoning        = "reasoning"
	otelPartTypeText             = "text"
	otelPartTypeToolCall         = "tool_call"
	otelPartTypeToolCallResponse = "tool_call_response"
	otelPartTypeURI              = "uri"

	otelModalityAudio = "audio"
	otelModalityFile  = "file"
	otelModalityImage = "image"
	otelModalityVideo = "video"
)

// OTelMessagePart is the OpenTelemetry-aligned message part payload.
type OTelMessagePart struct {
	Type      string          `json:"type"`
	Content   string          `json:"content,omitempty"`
	Modality  string          `json:"modality,omitempty"`
	MIMEType  string          `json:"mime_type,omitempty"`
	URI       string          `json:"uri,omitempty"`
	FileID    string          `json:"file_id,omitempty"`
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
	Response  json.RawMessage `json:"response,omitempty"`
	Detail    string          `json:"detail,omitempty"`
	Filename  string          `json:"filename,omitempty"`
}

// OTelInputMessage is the OpenTelemetry-aligned payload for gen_ai.input.messages.otel.
type OTelInputMessage struct {
	Role  model.Role        `json:"role"`
	Parts []OTelMessagePart `json:"parts"`
	Name  string            `json:"name,omitempty"`
}

// OTelOutputMessage is the OpenTelemetry-aligned payload for gen_ai.output.messages.otel.
type OTelOutputMessage struct {
	Role         model.Role        `json:"role"`
	Parts        []OTelMessagePart `json:"parts"`
	Name         string            `json:"name,omitempty"`
	FinishReason string            `json:"finish_reason,omitempty"`
}

func marshalOTelTelemetryMessages(messages []model.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalOTelTelemetryChoices(choices []model.Choice) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func otelPartsFromModelMessage(msg model.Message) []OTelMessagePart {
	_ = "STUB: not implemented"
	return nil
}

func otelPartsFromContentParts(contentParts []model.ContentPart) []OTelMessagePart {
	_ = "STUB: not implemented"
	return nil
}

func otelPartFromContentPart(contentPart model.ContentPart) (OTelMessagePart, bool) {
	_ = "STUB: not implemented"
	return *new(OTelMessagePart), false
}

func otelPartFromImage(image *model.Image) (OTelMessagePart, bool) {
	_ = "STUB: not implemented"
	return *new(OTelMessagePart), false
}

func otelPartFromFile(file *model.File) (OTelMessagePart, bool) {
	_ = "STUB: not implemented"
	return *new(OTelMessagePart), false
}

func otelPartFromToolCall(toolCall model.ToolCall) OTelMessagePart {
	_ = "STUB: not implemented"
	return *new(OTelMessagePart)
}

func otelPartFromToolCallResponse(msg model.Message) (OTelMessagePart, bool) {
	_ = "STUB: not implemented"
	return *new(OTelMessagePart), false
}

func toolResponseRawMessage(msg model.Message) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func rawJSONOrJSONString(raw []byte) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func jsonValueOrString(raw []byte) any { _ = "STUB: not implemented"; return *new(any) }

func normalizedOTelMessageRole(role model.Role, fallback model.Role) model.Role {
	_ = "STUB: not implemented"
	return *new(model.Role)
}

func isZeroOTelTelemetryMessage(msg model.Message) bool { _ = "STUB: not implemented"; return false }

func fileMetadata(file *model.File) (string, string) { _ = "STUB: not implemented"; return "", "" }

func modalityFromMIMEType(mimeType string) string { _ = "STUB: not implemented"; return "" }

func imageMIMEType(image *model.Image) string { _ = "STUB: not implemented"; return "" }

func normalizeFormatAsMIME(format, category string) string { _ = "STUB: not implemented"; return "" }

func mimeTypeFromURL(rawURL string) string { _ = "STUB: not implemented"; return "" }

func mimeTypeFromName(name string) string { _ = "STUB: not implemented"; return "" }

func normalizeMIMEType(mimeType string) string { _ = "STUB: not implemented"; return "" }

func derefString(v *string) string { _ = "STUB: not implemented"; return "" }
