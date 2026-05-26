//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package admin

const langfuseTraceIDPlaceholder = "{{trace_id}}"

// LangfuseStatus describes the current Langfuse integration state that the
// admin surface exposes.
type LangfuseStatus struct {
	Enabled          bool   `json:"enabled"`
	Ready            bool   `json:"ready"`
	Error            string `json:"error,omitempty"`
	UIBaseURL        string `json:"ui_base_url,omitempty"`
	TraceURLTemplate string `json:"trace_url_template,omitempty"`
}

func normalizeLangfuseStatus(raw LangfuseStatus) LangfuseStatus {
	_ = "STUB: not implemented"
	return *new(LangfuseStatus)
}

func (s *Service) langfuseTraceURL(traceID string) string { _ = "STUB: not implemented"; return "" }
