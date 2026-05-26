//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package templateresolver

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/invocationsaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/samplesaggregator"
)

const (
	// SampleAggregatorMajorityVoteName identifies the default samples aggregator.
	SampleAggregatorMajorityVoteName = "majority_vote"
	// InvocationAggregatorAverageName identifies the default invocations aggregator.
	InvocationAggregatorAverageName = "average"
)

// ResolveSamplesAggregator returns the samples aggregator identified by name.
func ResolveSamplesAggregator(name string) (samplesaggregator.SamplesAggregator, error) {
	_ = "STUB: not implemented"
	return *new(samplesaggregator.SamplesAggregator), nil
}

// ResolveInvocationsAggregator returns the invocations aggregator identified by name.
func ResolveInvocationsAggregator(name string) (invocationsaggregator.InvocationsAggregator, error) {
	_ = "STUB: not implemented"
	return *new(invocationsaggregator.InvocationsAggregator), nil
}
