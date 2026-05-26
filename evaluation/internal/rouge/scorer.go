//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package rouge

import (
	"context"
)

// Compute returns ROUGE scores for a single target and prediction pair.
// Compute returns an empty map when no ROUGE types are configured.
func Compute(ctx context.Context, target, prediction string, opt ...Option) (map[string]Score, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// computeMulti computes ROUGE scores for multiple targets and selects the maximum F-measure per type.
func computeMulti(ctx context.Context, targets []string, prediction string, opt ...Option) (map[string]Score, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateRougeType validates a ROUGE type identifier such as rouge1, rougeL, or rougeLsum.
func validateRougeType(rougeType string) error { _ = "STUB: not implemented"; return nil }

// parseRougeN parses a ROUGE-N type string and returns the N value.
func parseRougeN(rougeType string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// scoreNGrams computes ROUGE-N precision, recall, and F-measure for tokenized inputs.
func scoreNGrams(targetTokens, predTokens []string, n int) Score {
	_ = "STUB: not implemented"
	return *new(Score)
}

// createNGrams builds a multiset of n-grams keyed by a delimiter-joined token sequence.
func createNGrams(tokens []string, n int) map[string]int { _ = "STUB: not implemented"; return nil }

// scoreLCS computes ROUGE-L precision, recall, and F-measure using the LCS length.
func scoreLCS(targetTokens, predTokens []string) Score {
	_ = "STUB: not implemented"
	return *new(Score)
}

// lcsLength computes the length of the longest common subsequence.
func lcsLength(ref, can []string) int { _ = "STUB: not implemented"; return 0 }

// maxInt returns the larger of a and b.
func maxInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

// scoreSummaryLCS computes rougeLsum using summary-level LCS aggregation.
func scoreSummaryLCS(target, prediction string, tok Tokenizer, splitSummaries bool) (Score, error) {
	_ = "STUB: not implemented"
	return *new(Score), nil
}

// getSentences returns sentence strings using either newline splitting or a sentence tokenizer.
func getSentences(text string, splitSummaries bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// summaryLevelLCS computes rougeLsum and prevents double-counting matched tokens.
func summaryLevelLCS(refSent, canSent [][]string) Score {
	_ = "STUB: not implemented"
	return *new(Score)
}

// unionLCS returns the union of token indices from LCS matches across candidate sentences.
func unionLCS(ref []string, cans [][]string) []string { _ = "STUB: not implemented"; return nil }

// findUnion merges and sorts indices from multiple LCS paths.
func findUnion(lcsList [][]int) []int { _ = "STUB: not implemented"; return nil }

// lcsInd returns indices of one LCS between ref and can.
func lcsInd(ref, can []string) []int { _ = "STUB: not implemented"; return nil }

// lcsTable builds the dynamic programming table for LCS reconstruction.
func lcsTable(ref, can []string) [][]int { _ = "STUB: not implemented"; return nil }

// backtrackNoRec reconstructs a single LCS index sequence without recursion.
func backtrackNoRec(table [][]int, ref, can []string) []int { _ = "STUB: not implemented"; return nil }
