//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package registry provides runtime registration and resolution for metric extensions.
package registry

import (
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/finalresponse"
	criterionjson "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/json"
	criterionrouge "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/rouge"
	criteriontext "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/text"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/tooltrajectory"
)

// Registry resolves runtime metric extensions from registered names.
type Registry interface {
	// RegisterTextCompare registers a named text compare function.
	RegisterTextCompare(name string, fn criteriontext.CompareFunc) error
	// RegisterJSONCompare registers a named JSON compare function.
	RegisterJSONCompare(name string, fn criterionjson.CompareFunc) error
	// RegisterToolTrajectoryCompare registers a named tool trajectory compare function.
	RegisterToolTrajectoryCompare(name string, fn tooltrajectory.CompareFunc) error
	// RegisterFinalResponseCompare registers a named final response compare function.
	RegisterFinalResponseCompare(name string, fn finalresponse.CompareFunc) error
	// RegisterRougeTokenizer registers a named ROUGE tokenizer.
	RegisterRougeTokenizer(name string, tok criterionrouge.Tokenizer) error
	// Resolve resolves registered names into runtime implementations on the metric.
	Resolve(evalMetric *metric.EvalMetric) error
}

type registry struct {
	mu                     sync.RWMutex
	textCompares           map[string]criteriontext.CompareFunc
	jsonCompares           map[string]criterionjson.CompareFunc
	toolTrajectoryCompares map[string]tooltrajectory.CompareFunc
	finalResponseCompares  map[string]finalresponse.CompareFunc
	rougeTokenizers        map[string]criterionrouge.Tokenizer
}

// New creates a metric extension registry.
func New() Registry { _ = "STUB: not implemented"; return *new(Registry) }

// RegisterTextCompare registers a named text compare function.
func (r *registry) RegisterTextCompare(name string, fn criteriontext.CompareFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterJSONCompare registers a named JSON compare function.
func (r *registry) RegisterJSONCompare(name string, fn criterionjson.CompareFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterToolTrajectoryCompare registers a named tool trajectory compare function.
func (r *registry) RegisterToolTrajectoryCompare(name string, fn tooltrajectory.CompareFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterFinalResponseCompare registers a named final response compare function.
func (r *registry) RegisterFinalResponseCompare(name string, fn finalresponse.CompareFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRougeTokenizer registers a named ROUGE tokenizer.
func (r *registry) RegisterRougeTokenizer(name string, tok criterionrouge.Tokenizer) error {
	_ = "STUB: not implemented"
	return nil
}

// Resolve resolves registered names into runtime implementations on the metric.
func (r *registry) Resolve(evalMetric *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) resolveToolTrajectoryCriterion(criterion *tooltrajectory.ToolTrajectoryCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) resolveToolTrajectoryStrategy(strategy *tooltrajectory.ToolTrajectoryStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) resolveFinalResponseCriterion(criterion *finalresponse.FinalResponseCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) resolveTextCriterion(criterion *criteriontext.TextCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) resolveJSONCriterion(criterion *criterionjson.JSONCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) resolveRougeCriterion(criterion *criterionrouge.RougeCriterion) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) lookupTextCompare(name string) (criteriontext.CompareFunc, error) {
	_ = "STUB: not implemented"
	return *new(criteriontext.CompareFunc), nil
}

func (r *registry) lookupJSONCompare(name string) (criterionjson.CompareFunc, error) {
	_ = "STUB: not implemented"
	return *new(criterionjson.CompareFunc), nil
}

func (r *registry) lookupToolTrajectoryCompare(name string) (tooltrajectory.CompareFunc, error) {
	_ = "STUB: not implemented"
	return *new(tooltrajectory.CompareFunc), nil
}

func (r *registry) lookupFinalResponseCompare(name string) (finalresponse.CompareFunc, error) {
	_ = "STUB: not implemented"
	return *new(finalresponse.CompareFunc), nil
}

func (r *registry) lookupRougeTokenizer(name string) (criterionrouge.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(criterionrouge.Tokenizer), nil
}
