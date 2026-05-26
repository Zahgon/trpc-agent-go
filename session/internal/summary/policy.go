//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package summary

// SummaryDispatchPolicy controls which branch summaries are allowed to run and
// whether branch-triggered updates should also refresh the full-session summary.
type SummaryDispatchPolicy struct {
	FilterAllowlist    map[string]struct{}
	CascadeFullSession bool
}

// NewSummaryDispatchPolicy normalizes summary dispatch settings.
func NewSummaryDispatchPolicy(
	filterAllowlist []string,
	cascadeFullSession bool,
) SummaryDispatchPolicy {
	_ = "STUB: not implemented"
	return *new(SummaryDispatchPolicy)
}

// SummaryTargets returns the summary keys that should be refreshed for the
// given trigger filterKey.
func (p SummaryDispatchPolicy) SummaryTargets(filterKey string) []string {
	_ = "STUB: not implemented"
	return nil
}

// AllowsFilterKey reports whether the given filterKey may be summarized when a
// caller explicitly requests that key.
func (p SummaryDispatchPolicy) AllowsFilterKey(filterKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func (p SummaryDispatchPolicy) allowsBranch(filterKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func normalizeSummaryFilterAllowlist(
	filterAllowlist []string,
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func matchSummaryFilterKey(allowedKey, filterKey string) bool {
	_ = "STUB: not implemented"
	return false
}
