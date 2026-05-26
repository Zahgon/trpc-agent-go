//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package clone

import "trpc.group/trpc-go/trpc-agent-go/model"

func cloneMessages(src []*model.Message) ([]*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneMessage(src *model.Message) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneContentParts(src []model.ContentPart) []model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func cloneImage(src *model.Image) *model.Image { _ = "STUB: not implemented"; return nil }

func cloneAudio(src *model.Audio) *model.Audio { _ = "STUB: not implemented"; return nil }

func cloneFile(src *model.File) *model.File { _ = "STUB: not implemented"; return nil }

func cloneToolCalls(src []model.ToolCall) ([]model.ToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneGenerationConfig(src *model.GenerationConfig) *model.GenerationConfig {
	_ = "STUB: not implemented"
	return nil
}
