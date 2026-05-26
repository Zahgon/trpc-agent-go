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
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// FormatResponseErrorLabel converts a response error into a telemetry label.
func FormatResponseErrorLabel(respErr *model.ResponseError, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

// ToErrorType converts an error to an error type.
func ToErrorType(err error, errorType string) string { _ = "STUB: not implemented"; return "" }
