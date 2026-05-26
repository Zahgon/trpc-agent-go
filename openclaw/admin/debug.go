//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package admin

import (
	"time"
)

type debugStatus struct {
	Enabled      bool               `json:"enabled"`
	BySessionDir string             `json:"by_session_dir,omitempty"`
	SessionCount int                `json:"session_count"`
	TraceCount   int                `json:"trace_count"`
	Error        string             `json:"error,omitempty"`
	Sessions     []debugSessionView `json:"sessions,omitempty"`
	RecentTraces []debugTraceView   `json:"recent_traces,omitempty"`
}

type debugSessionView struct {
	SessionID      string    `json:"session_id,omitempty"`
	TraceCount     int       `json:"trace_count"`
	LastTraceAt    time.Time `json:"last_trace_at,omitempty"`
	Channel        string    `json:"channel,omitempty"`
	RequestID      string    `json:"request_id,omitempty"`
	TraceID        string    `json:"trace_id,omitempty"`
	ProfileID      string    `json:"profile_id,omitempty"`
	ProfileVersion string    `json:"profile_version,omitempty"`
	ProfileAppName string    `json:"profile_app_name,omitempty"`
	TracePath      string    `json:"trace_path,omitempty"`
	LangfuseURL    string    `json:"langfuse_url,omitempty"`
	MetaURL        string    `json:"meta_url,omitempty"`
	EventsURL      string    `json:"events_url,omitempty"`
	ResultURL      string    `json:"result_url,omitempty"`
}

type debugTraceView struct {
	SessionID      string    `json:"session_id,omitempty"`
	StartedAt      time.Time `json:"started_at,omitempty"`
	Channel        string    `json:"channel,omitempty"`
	RequestID      string    `json:"request_id,omitempty"`
	MessageID      string    `json:"message_id,omitempty"`
	TraceID        string    `json:"trace_id,omitempty"`
	ProfileID      string    `json:"profile_id,omitempty"`
	ProfileVersion string    `json:"profile_version,omitempty"`
	ProfileAppName string    `json:"profile_app_name,omitempty"`
	TracePath      string    `json:"trace_path,omitempty"`
	LangfuseURL    string    `json:"langfuse_url,omitempty"`
	MetaURL        string    `json:"meta_url,omitempty"`
	EventsURL      string    `json:"events_url,omitempty"`
	ResultURL      string    `json:"result_url,omitempty"`
}

type debugTraceRef struct {
	TraceDir  string    `json:"trace_dir"`
	StartedAt time.Time `json:"started_at"`
	Channel   string    `json:"channel,omitempty"`
	RequestID string    `json:"request_id,omitempty"`
	MessageID string    `json:"message_id,omitempty"`
	TraceID   string    `json:"trace_id,omitempty"`
}

func (s *Service) debugStatus() debugStatus { _ = "STUB: not implemented"; return *new(debugStatus) }

func (s *Service) debugStatusForSession(sessionID string) debugStatus {
	_ = "STUB: not implemented"
	return *new(debugStatus)
}

func (s *Service) buildDebugStatus(sessionFilter string) debugStatus {
	_ = "STUB: not implemented"
	return *new(debugStatus)
}

func (s *Service) loadDebugTraces(
	root string,
	sessionFilter string,
) ([]debugTraceView, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const debugMetaTraceRefName = "trace.json"

func (s *Service) readDebugTrace(
	root string,
	bySessionRoot string,
	refPath string,
	sessionFilter string,
) (debugTraceView, bool, error) {
	_ = "STUB: not implemented"
	return *new(debugTraceView), false, nil
}

type debugRuntimeProfileRecord struct {
	Kind    string         `json:"kind"`
	Payload map[string]any `json:"payload,omitempty"`
}

func enrichDebugTraceProfile(traceAbs string, out *debugTraceView) {
	_ = "STUB: not implemented"
	return
}

func stringField(values map[string]any, key string) string { _ = "STUB: not implemented"; return "" }

func (s *Service) debugFileURL(tracePath string, name string) string {
	_ = "STUB: not implemented"
	return ""
}

func limitDebugTraces(
	items []debugTraceView,
	limit int,
) []debugTraceView {
	_ = "STUB: not implemented"
	return nil
}

func fileExists(path string) bool { _ = "STUB: not implemented"; return false }
