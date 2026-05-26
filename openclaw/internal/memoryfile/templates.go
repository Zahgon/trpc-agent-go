//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package memoryfile

import "strings"

var (
	defaultTemplateTrimmed       = strings.TrimSpace(DefaultTemplate())
	legacyDefaultTemplateTrimmed = strings.TrimSpace(
		legacyDefaultTemplate(),
	)
)

func DefaultTemplate() string { _ = "STUB: not implemented"; return "" }

// legacyDefaultTemplate returns the full text of the default template used
// before the wording was updated to scope-aware text. We still need to
// recognise untouched files created with the old template so they are not
// injected as real memory on the fallback path.
func legacyDefaultTemplate() string { _ = "STUB: not implemented"; return "" }

func IsDefaultTemplate(content string) bool { _ = "STUB: not implemented"; return false }
