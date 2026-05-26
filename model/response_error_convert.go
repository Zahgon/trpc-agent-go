//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package model

type errorTyper interface {
	ErrorType() string
}

type errorCoder interface {
	ErrorCode() string
}

type codeStringer interface {
	Code() string
}

type codeInter interface {
	Code() int
}

type codeInt32er interface {
	Code() int32
}

type codeInt64er interface {
	Code() int64
}

// ResponseErrorFromError converts an error into a ResponseError.
//
// This helper is primarily used when an internal workflow error needs to be
// serialized into an event stream. It attempts to preserve structured fields
// (type/code) when the error carries them.
func ResponseErrorFromError(err error, fallbackType string) *ResponseError {
	_ = "STUB: not implemented"
	return nil
}

func errorTypeFromError(err error) string { _ = "STUB: not implemented"; return "" }

func errorCodeFromError(err error) string { _ = "STUB: not implemented"; return "" }

func stringPtr(s string) *string { _ = "STUB: not implemented"; return nil }
