//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package hallucination builds judge prompts for hallucination evaluation.
package hallucination

import (
	"context"
	"regexp"
	"text/template"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/messagesconstructor"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

var (
	noValidationContext = "No validation context was captured."
	segmentationPrompt  = `
# Mission

Segment the final answer into sentence-level or bullet-level claims.
You must preserve the original wording and punctuation exactly.

# Rules

1. Output one claim per line using the exact format <sentence>...</sentence>.
2. Cover the whole final answer in order.
3. Split bullet items and numbered items into separate claims when they carry separate meaning.
4. Keep short stylistic or process-only statements as their own claims if they appear in the final answer.
5. Do not add explanations, numbering, or any text outside <sentence> tags.
6. If the final answer is empty, output exactly:
<sentence>[EMPTY_RESPONSE]</sentence>

# Input

<final_answer>
{{.FinalResponse}}
</final_answer>
`
	segmentationPromptTemplate = template.Must(template.New("segmentationPrompt").Parse(segmentationPrompt))
	validatorPrompt            = `
# Mission

Your mission is to detect hallucinations in the agent's segmented final answer.
You will be given:
- a textual context captured during execution, and
- segmented sentences from the final answer.

Evaluate every input sentence against the textual context.
You must cover the whole segmented answer.

# Labels

supported: The sentence is fully supported by the textual context.
unsupported: The sentence is not supported by the textual context.
contradictory: The textual context directly contradicts the sentence.
disputed: The textual context contains both supporting and contradicting evidence.
not_applicable: The sentence does not require factual grounding, such as greetings, stylistic fillers, or process-only statements.

# Key Evaluation Principles

1. Only use the provided textual context as trusted evidence.
   Do not use external knowledge, common-sense guessing, or the final answer itself as evidence.
2. Be strict.
   If the textual context does not fully support a sentence, label it as unsupported.
3. Use the sentence IDs from the segmented input exactly.
   Do not renumber, merge, or skip them.
4. Keep the reason concise.
   Mention the decisive grounding signal briefly when useful.
5. Map labels to verdicts exactly as follows.
   supported => yes
   not_applicable => yes
   unsupported => no
   contradictory => no
   disputed => no
6. If the final answer is empty, output exactly one block with:
   ID: 1
   Reason: No final answer was produced.
   Label: unsupported
   Verdict: no

# Output Format

Repeat the following block for every sentence, starting with a new line:

ID: [1..N]
Reason: [A concise explanation of the label.]
Label: [supported|unsupported|contradictory|disputed|not_applicable]
Verdict: [yes|no]

Output only these blocks.

# Input

<context>
{{.ValidationContext}}
</context>

<segmented_sentences>
{{.SegmentedSentences}}
</segmented_sentences>
`
	validatorPromptTemplate = template.Must(template.New("validatorPrompt").Parse(validatorPrompt))
	sentenceRegex           = regexp.MustCompile(`(?s)<sentence>(.*?)</sentence>`)
)

type hallucinationMessagesConstructor struct {
}

// New returns a messages constructor for hallucination evaluation.
func New() messagesconstructor.MessagesConstructor {
	_ = "STUB: not implemented"
	return *new(messagesconstructor.MessagesConstructor)
}

// ConstructMessages builds judge prompts for hallucination evaluation.
func (e *hallucinationMessagesConstructor) ConstructMessages(ctx context.Context, actuals, _ []*evalset.Invocation,
	evalMetric *metric.EvalMetric) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildValidationContext(actuals []*evalset.Invocation) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func buildSegmentedSentences(ctx context.Context, finalResponse string, evalMetric *metric.EvalMetric) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type hallucinationPromptData struct {
	ValidationContext  string
	FinalResponse      string
	SegmentedSentences string
}
