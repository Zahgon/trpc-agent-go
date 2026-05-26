//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package responsejson decodes structured JSON judge responses.
package responsejson

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// UnmarshalContent decodes the first response choice content as JSON into dst.
func UnmarshalContent(resp *model.Response, dst any) error { _ = "STUB: not implemented"; return nil }

func trimCodeFence(content string) string { _ = "STUB: not implemented"; return "" }
