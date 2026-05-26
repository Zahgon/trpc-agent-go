//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package conversation

import (
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/event"
)

// Annotation stores channel-provided conversation metadata.
type Annotation struct {
	HistoryMode   string            `json:"history_mode,omitempty"`
	StorageUserID string            `json:"storage_user_id,omitempty"`
	ActorID       string            `json:"actor_id,omitempty"`
	ActorLabel    string            `json:"actor_label,omitempty"`
	ActorLabels   map[string]string `json:"actor_labels,omitempty"`
	QuoteText     string            `json:"quote_text,omitempty"`
}

// MergeRequestExtension stores conversation metadata in request
// extensions.
func MergeRequestExtension(
	extensions map[string]json.RawMessage,
	annotation Annotation,
) (map[string]json.RawMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AnnotationFromRequestExtensions decodes request conversation metadata.
func AnnotationFromRequestExtensions(
	extensions map[string]json.RawMessage,
) (Annotation, bool, error) {
	_ = "STUB: not implemented"
	return *new(Annotation), false, nil
}

// AnnotationFromRuntimeState decodes runtime conversation metadata.
func AnnotationFromRuntimeState(
	state map[string]any,
) (Annotation, bool) {
	_ = "STUB: not implemented"
	return *new(Annotation), false
}

// RuntimeState returns runtime state for one run.
func RuntimeState(annotation Annotation) map[string]any { _ = "STUB: not implemented"; return nil }

// SetEventAnnotation persists conversation metadata on one event.
func SetEventAnnotation(
	evt *event.Event,
	annotation Annotation,
) error {
	_ = "STUB: not implemented"
	return nil
}

// AnnotationFromEvent decodes persisted event metadata.
func AnnotationFromEvent(
	evt event.Event,
) (Annotation, bool, error) {
	_ = "STUB: not implemented"
	return *new(Annotation), false, nil
}

type eventAnnotation struct {
	ActorID    string `json:"actor_id,omitempty"`
	ActorLabel string `json:"actor_label,omitempty"`
	QuoteText  string `json:"quote_text,omitempty"`
}

func decodeAnnotation(
	extensions map[string]json.RawMessage,
) (Annotation, bool, error) {
	_ = "STUB: not implemented"
	return *new(Annotation), false, nil
}

func isZeroRuntimeAnnotation(annotation Annotation) bool { _ = "STUB: not implemented"; return false }

func normalizeAnnotation(annotation Annotation) Annotation {
	_ = "STUB: not implemented"
	return *new(Annotation)
}

func normalizeActorLabels(
	labels map[string]string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func persistableEventAnnotation(
	annotation Annotation,
) (eventAnnotation, bool) {
	_ = "STUB: not implemented"
	return *new(eventAnnotation), false
}

func isZeroEventAnnotation(annotation eventAnnotation) bool {
	_ = "STUB: not implemented"
	return false
}
