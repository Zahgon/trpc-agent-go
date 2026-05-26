//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package finalresponse defines criteria for comparing agent final responses.
package finalresponse

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	cjson "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/json"
	crouge "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/rouge"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/text"
	cxml "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/xml"
)

// CompareFunc defines custom final response comparison logic.
type CompareFunc func(actual, expected *evalset.Invocation) (bool, error)

// FinalResponseCriterion provides comparison rules for final response messages.
type FinalResponseCriterion struct {
	// Text compares the final response content as plain text.
	Text *text.TextCriterion `json:"text,omitempty"`
	// JSON compares the final response content as JSON.
	JSON *cjson.JSONCriterion `json:"json,omitempty"`
	// Rouge scores the final response content with ROUGE.
	Rouge *crouge.RougeCriterion `json:"rouge,omitempty"`
	// XML validates the final response content as XML.
	XML *cxml.XMLCriterion `json:"xml,omitempty"`
	// CompareName selects a registered comparison implementation by name.
	CompareName string `json:"compareName,omitempty"`
	// Compare allows overriding the built-in matching logic.
	Compare CompareFunc `json:"-"`
}

// New creates a FinalResponseCriterion with the provided options.
func New(opt ...Option) *FinalResponseCriterion { _ = "STUB: not implemented"; return nil }

// Match compares the final responses of actual and expected invocations using the provided context.
func (c *FinalResponseCriterion) Match(ctx context.Context, actual, expected *evalset.Invocation) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *FinalResponseCriterion) hasConfiguredCriterion() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *FinalResponseCriterion) matchFinalResponseContent(ctx context.Context, actual, expected *evalset.Invocation) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// matchContentAsText compares two strings using a TextCriterion.
func matchContentAsText(actual, expected string, criterion *text.TextCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

// matchContentAsJSON validates or compares JSON content using a JSONCriterion.
func matchContentAsJSON(actual, expected string, criterion *cjson.JSONCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

// matchContentAsXML validates a string using an XMLCriterion.
func matchContentAsXML(actual, expected string, criterion *cxml.XMLCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

// matchContentAsRouge scores and validates two strings using a RougeCriterion.
func matchContentAsRouge(ctx context.Context, actual, expected string, criterion *crouge.RougeCriterion) error {
	_ = "STUB: not implemented"
	return nil
}
