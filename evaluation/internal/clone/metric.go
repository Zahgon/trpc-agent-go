//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package clone

import (
	criterionjson "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/json"
	criterionlength "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/length"
	criterionllm "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/llm"
	criterionrouge "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/rouge"
	criteriontext "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/text"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/tooltrajectory"
	criterionxml "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/xml"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/finalresponse"
)

// CloneEvalMetric clones a metric configuration.
func CloneEvalMetric(src *metric.EvalMetric) (*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneCriterion(src *criterion.Criterion) (*criterion.Criterion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneToolTrajectoryCriterion(src *tooltrajectory.ToolTrajectoryCriterion) (*tooltrajectory.ToolTrajectoryCriterion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneToolTrajectoryStrategy(src *tooltrajectory.ToolTrajectoryStrategy) (*tooltrajectory.ToolTrajectoryStrategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneFinalResponseCriterion(src *finalresponse.FinalResponseCriterion) (*finalresponse.FinalResponseCriterion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneTextCriterion(src *criteriontext.TextCriterion) *criteriontext.TextCriterion {
	_ = "STUB: not implemented"
	return nil
}

func cloneJSONCriterion(src *criterionjson.JSONCriterion) (*criterionjson.JSONCriterion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneLengthCriterion(src *criterionlength.LengthCriterion) *criterionlength.LengthCriterion {
	_ = "STUB: not implemented"
	return nil
}

func cloneRougeCriterion(src *criterionrouge.RougeCriterion) *criterionrouge.RougeCriterion {
	_ = "STUB: not implemented"
	return nil
}

func cloneXMLCriterion(src *criterionxml.XMLCriterion) *criterionxml.XMLCriterion {
	_ = "STUB: not implemented"
	return nil
}

func cloneLLMCriterion(src *criterionllm.LLMCriterion) (*criterionllm.LLMCriterion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneJudgeTemplateOptions(src *criterionllm.JudgeTemplateOptions) *criterionllm.JudgeTemplateOptions {
	_ = "STUB: not implemented"
	return nil
}

func cloneTemplateVariableBindings(src []*criterionllm.TemplateVariableBinding) []*criterionllm.TemplateVariableBinding {
	_ = "STUB: not implemented"
	return nil
}

func cloneTemplateVariableBinding(src *criterionllm.TemplateVariableBinding) *criterionllm.TemplateVariableBinding {
	_ = "STUB: not implemented"
	return nil
}

func cloneRubrics(src []*criterionllm.Rubric) []*criterionllm.Rubric {
	_ = "STUB: not implemented"
	return nil
}

func cloneRubric(src *criterionllm.Rubric) *criterionllm.Rubric {
	_ = "STUB: not implemented"
	return nil
}

func cloneJudgeModelOptions(src *criterionllm.JudgeModelOptions) (*criterionllm.JudgeModelOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
