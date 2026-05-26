//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package gwclient provides an in-process client for the gateway HTTP
// handler.
package gwclient

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwproto"
)

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"

	methodPost = "POST"
)

// Client invokes the gateway handler without a network hop.
type Client struct {
	handler      http.Handler
	messagesPath string
	streamPath   string
	cancelPath   string
}

// New creates a new client for the gateway messages endpoint.
func New(
	handler http.Handler,
	messagesPath string,
	cancelPath string,
) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWithStreamPath creates a client with an explicit stream path.
func NewWithStreamPath(
	handler http.Handler,
	messagesPath string,
	streamPath string,
	cancelPath string,
) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MessageRequest matches the gateway /messages JSON payload.
type MessageRequest = gwproto.MessageRequest

// MessageStreamOptions controls optional streaming behaviors.
type MessageStreamOptions = gwproto.MessageStreamOptions

// StreamEvent matches the gateway streaming event payload.
type StreamEvent = gwproto.StreamEvent

// APIError matches gateway error payloads.
type APIError = gwproto.APIError

// Usage matches gateway token usage payloads.
type Usage = gwproto.Usage

// MessageResponse matches the gateway /messages response JSON.
type MessageResponse struct {
	SessionID string    `json:"session_id,omitempty"`
	RequestID string    `json:"request_id,omitempty"`
	Reply     string    `json:"reply,omitempty"`
	Usage     *Usage    `json:"usage,omitempty"`
	Ignored   bool      `json:"ignored,omitempty"`
	Error     *APIError `json:"error,omitempty"`

	StatusCode int `json:"-"`
}

// ScheduledJobSummary is a transport-safe view of one scheduled job.
type ScheduledJobSummary struct {
	ID               string     `json:"id,omitempty"`
	Name             string     `json:"name,omitempty"`
	Enabled          bool       `json:"enabled"`
	Schedule         string     `json:"schedule,omitempty"`
	Message          string     `json:"message,omitempty"`
	MaxRuns          int        `json:"max_runs,omitempty"`
	RunCount         int        `json:"run_count,omitempty"`
	SuccessCount     int        `json:"success_count,omitempty"`
	FailureCount     int        `json:"failure_count,omitempty"`
	DeliveryFailures int        `json:"delivery_failures,omitempty"`
	EndsAt           *time.Time `json:"ends_at,omitempty"`
	OverlapPolicy    string     `json:"overlap_policy,omitempty"`
	NextRunAt        *time.Time `json:"next_run_at,omitempty"`
	LastStatus       string     `json:"last_status,omitempty"`
	LastError        string     `json:"last_error,omitempty"`
	LastOutput       string     `json:"last_output,omitempty"`
	DeliveryChannel  string     `json:"delivery_channel,omitempty"`
	DeliveryTarget   string     `json:"delivery_target,omitempty"`
}

// SendMessage sends one message to the gateway handler.
func (c *Client) SendMessage(
	ctx context.Context,
	req MessageRequest,
) (MessageResponse, error) {
	_ = "STUB: not implemented"
	return *new(MessageResponse), nil
}

type cancelRequest struct {
	RequestID string `json:"request_id,omitempty"`
}

type cancelResponse struct {
	Canceled bool `json:"canceled"`
}

type errorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

// Cancel attempts to cancel a request by request_id.
func (c *Client) Cancel(
	ctx context.Context,
	requestID string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type responseRecorder struct {
	header http.Header
	code   int
	body   bytes.Buffer
}

func newResponseRecorder() *responseRecorder { _ = "STUB: not implemented"; return nil }

func (r *responseRecorder) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (r *responseRecorder) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (r *responseRecorder) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *responseRecorder) Code() int { _ = "STUB: not implemented"; return 0 }

func (r *responseRecorder) BodyBytes() []byte { _ = "STUB: not implemented"; return nil }
