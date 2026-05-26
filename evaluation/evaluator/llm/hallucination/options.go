//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package hallucination

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/invocationsaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/messagesconstructor"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/samplesaggregator"
)

type options struct {
	messagesConstructor   messagesconstructor.MessagesConstructor
	responsescorer        responsescorer.ResponseScorer
	samplesAggregator     samplesaggregator.SamplesAggregator
	invocationsAggregator invocationsaggregator.InvocationsAggregator
}

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option customizes Hallucination evaluator dependencies.
type Option func(*options)

// WithMessagesConstructor sets the prompt builder for hallucination evaluation.
func WithMessagesConstructor(messagesConstructor messagesconstructor.MessagesConstructor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithResponsescorer sets the response scorer implementation.
func WithResponsescorer(responsescorer responsescorer.ResponseScorer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSamplesAggregator sets how multiple judge samples are reduced.
func WithSamplesAggregator(samplesAggregator samplesaggregator.SamplesAggregator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithInvocationsAggregator sets how per-invocation scores are aggregated.
func WithInvocationsAggregator(invocationsAggregator invocationsaggregator.InvocationsAggregator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
