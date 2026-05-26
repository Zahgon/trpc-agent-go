//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package multimodal converts AG-UI multimodal content into internal model messages.
package multimodal

import (
	"github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/types"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	// CustomEventNameUserMessage is the custom event name used to persist user input messages in the track stream.
	CustomEventNameUserMessage = "trpc-agent-go.user_message"
)

// UserMessageFromInputContents converts AG-UI multimodal input contents into a model user message.
func UserMessageFromInputContents(contents []types.InputContent) (model.Message, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), nil
}

func contentPartFromInputContent(part types.InputContent) (*model.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func contentPartFromBinaryInput(part types.InputContent) (*model.ContentPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fileFromBinaryURL(part types.InputContent, mimeType, fileURL string) *model.File {
	_ = "STUB: not implemented"
	return nil
}

func fileFromBinaryID(part types.InputContent) *model.File { _ = "STUB: not implemented"; return nil }

func fileNameFromArtifactRef(fileID string) string { _ = "STUB: not implemented"; return "" }

func decodeBase64Payload(payload string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
