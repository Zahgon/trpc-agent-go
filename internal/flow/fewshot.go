//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package flow

import "trpc.group/trpc-go/trpc-agent-go/model"

// InsertFewShotMessages inserts few-shot examples after the leading system message block.
func InsertFewShotMessages(
	messages []model.Message,
	examples [][]model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func flattenFewShotMessages(examples [][]model.Message) []model.Message {
	_ = "STUB: not implemented"
	return nil
}
