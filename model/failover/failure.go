//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package failover

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type failureError struct {
	failures []failureRecord
}

type failureRecord struct {
	candidate string
	message   string
	errType   string
}

func newFailureError(failures []failureRecord) error { _ = "STUB: not implemented"; return nil }

func (e *failureError) Error() string { _ = "STUB: not implemented"; return "" }

func buildFailureMessage(failures []failureRecord) string { _ = "STUB: not implemented"; return "" }

func failuresFromError(err error, fallback []failureRecord) []failureRecord {
	_ = "STUB: not implemented"
	return nil
}

func appendFailure(
	failures []failureRecord,
	candidateName string,
	message string,
	errType string,
) []failureRecord {
	_ = "STUB: not implemented"
	return nil
}

func buildFailureResponse(failures []failureRecord) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func failureResponseType(failures []failureRecord) string { _ = "STUB: not implemented"; return "" }
