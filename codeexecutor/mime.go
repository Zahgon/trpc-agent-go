//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package codeexecutor

const (
	mimeTextPrefix = "text/"
	mimeAppJSON    = "application/json"
	mimeSuffixJSON = "+json"
)

// IsTextMIME reports whether mimeType describes a text format that is safe
// to inline as UTF-8 text.
func IsTextMIME(mimeType string) bool { _ = "STUB: not implemented"; return false }
