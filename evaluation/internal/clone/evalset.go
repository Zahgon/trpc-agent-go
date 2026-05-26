//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package clone

import "trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"

// CloneEvalSet clones an eval set and all nested eval cases.
func CloneEvalSet(src *evalset.EvalSet) (*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CloneEvalCase clones an eval case and all nested invocations.
func CloneEvalCase(src *evalset.EvalCase) (*evalset.EvalCase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneEvalCaseRubrics(src []*evalset.EvalCaseRubric) []*evalset.EvalCaseRubric {
	_ = "STUB: not implemented"
	return nil
}

func cloneEvalCaseRubric(src *evalset.EvalCaseRubric) *evalset.EvalCaseRubric {
	_ = "STUB: not implemented"
	return nil
}

func cloneInvocations(src []*evalset.Invocation) ([]*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneInvocation(src *evalset.Invocation) (*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneTools(src []*evalset.Tool) ([]*evalset.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneTool(src *evalset.Tool) (*evalset.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneSessionInput(src *evalset.SessionInput) (*evalset.SessionInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
