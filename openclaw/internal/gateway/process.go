//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package gateway

import (
	"context"
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwproto"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/debugrecorder"
)

type preparedMessageRun struct {
	userID              string
	sessionID           string
	requestID           string
	requestSystemPrompt string
	inbound             InboundMessage
	userMsg             model.Message
	streamOptions       *gwproto.MessageStreamOptions
	extensions          map[string]json.RawMessage
}

// ProcessMessage processes a gateway message request without an HTTP hop.
//
// It returns a JSON-serializable response payload and the HTTP-like status
// code that the /messages endpoint would use.
func (s *Server) ProcessMessage(
	ctx context.Context,
	req gwproto.MessageRequest,
) (rsp gwproto.MessageResponse, status int) {
	_ = "STUB: not implemented"
	return *new(gwproto.MessageResponse), 0
}

func (s *Server) prepareMessageRun(
	ctx context.Context,
	req gwproto.MessageRequest,
	opts *gwproto.MessageStreamOptions,
	trace *debugrecorder.Trace,
) (preparedMessageRun, *gwproto.MessageResponse, int) {
	_ = "STUB: not implemented"
	return *new(preparedMessageRun), nil, 0
}

func cloneStreamOptions(
	src *gwproto.MessageStreamOptions,
) *gwproto.MessageStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func cloneExtensions(
	src map[string]json.RawMessage,
) map[string]json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// CancelRequest cancels an in-flight run by request ID.
//
// It returns (canceled=false, status=200) when no matching run exists.
func (s *Server) CancelRequest(
	ctx context.Context,
	requestID string,
) (canceled bool, apiErr *gwproto.APIError, status int) {
	_ = "STUB: not implemented"
	return false, nil, 0
}

func (s *Server) ensureTrace(
	ctx context.Context,
	req gwproto.MessageRequest,
) (*debugrecorder.Trace, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type errString string

func (e errString) Error() string { _ = "STUB: not implemented"; return "" }
