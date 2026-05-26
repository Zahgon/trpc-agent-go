//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package finalresponse

import (
	cjson "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/json"
	crouge "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/rouge"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/text"
	cxml "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/xml"
)

// options holds construction-time configuration for FinalResponseCriterion.
type options struct {
	// text configures text-based comparison.
	text *text.TextCriterion
	// json configures JSON-based comparison.
	json *cjson.JSONCriterion
	// rouge configures ROUGE scoring comparison.
	rouge *crouge.RougeCriterion
	// xml configures XML validation.
	xml *cxml.XMLCriterion
	// compareName selects a registered comparison implementation by name.
	compareName string
	// compare overrides built-in comparison when provided.
	compare CompareFunc
}

// newOptions applies functional options to build a criterion configuration.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option configures FinalResponseCriterion.
type Option func(*options)

// WithTextCriterion sets the text criterion.
func WithTextCriterion(criterion *text.TextCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithJSONCriterion sets the JSON criterion.
func WithJSONCriterion(criterion *cjson.JSONCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithXMLCriterion sets the XML criterion.
func WithXMLCriterion(criterion *cxml.XMLCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCompareName sets the name of the registered compare function.
func WithCompareName(compareName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCompare sets the custom compare function.
func WithCompare(compare CompareFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRougeCriterion sets the ROUGE criterion.
func WithRougeCriterion(criterion *crouge.RougeCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
