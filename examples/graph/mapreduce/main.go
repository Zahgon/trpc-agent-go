//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a Map-Reduce workflow over a document:
// 1) chunk the document, 2) retrieve top-K relevant chunks for a question,
// 3) summarize selected chunks in parallel, and 4) join partial summaries.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	defaultModelName = "deepseek-v4-flash"

	// Schema keys
	keyDocText          = "doc_text"
	keyChunks           = "chunks"
	keySelected         = "selected_chunks"
	keySelectedCount    = "selected_count"
	keyPartialSummaries = "partial_summaries"
	keyFinalAnswer      = "final_answer"
)

var (
	flagModel     = flag.String("model", defaultModelName, "Model name to use for LLM nodes")
	flagFile      = flag.String("file", "./sample.txt", "Path to the input document (text)")
	flagChunkSize = flag.Int("chunk-size", 800, "Chunk size (characters)")
	flagOverlap   = flag.Int("overlap", 100, "Chunk overlap (characters)")
	flagTopK      = flag.Int("top-k", 4, "Top K chunks to summarize in parallel")
	flagVerbose   = flag.Bool("verbose", false, "Verbose node/model events")
)

func main() {
	flag.Parse()
	fmt.Println("🗺️  Map-Reduce Document QA Example")
	fmt.Printf("Model: %s\n", *flagModel)
	fmt.Printf("File : %s\n", absOrSelf(*flagFile))
	fmt.Println(strings.Repeat("=", 56))
	demo := &mapReduceDemo{
		modelName:   *flagModel,
		chunkSize:   *flagChunkSize,
		overlap:     *flagOverlap,
		topK:        *flagTopK,
		verboseLogs: *flagVerbose,
	}
	if err := demo.setup(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "setup failed: %v\n", err)
		os.Exit(1)
	}

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer demo.runner.Close()

	if err := demo.runInteractive(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "run failed: %v\n", err)
		os.Exit(1)
	}
}

// mapReduceDemo holds runtime config and shared components.
type mapReduceDemo struct {
	modelName   string
	chunkSize   int
	overlap     int
	topK        int
	verboseLogs bool

	docText string

	runner    runner.Runner
	userID    string
	sessionID string
}

func (d *mapReduceDemo) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Load document text eagerly so node functions can reference it.
	return nil
}

// Build the graph.

// Create GraphAgent.

// Session service + runner

// buildGraph constructs the Map-Reduce graph.
func (d *mapReduceDemo) buildGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Base schema with message semantics + our fields.
	return nil, nil
}

// 1) Load + chunk

// 2) Retrieve top‑K for the query in user_input

// 3) Fan‑out create commands to map_summarize with per‑chunk input

// 3b) Map summarize each chunk with the question context

// 3c) Collect each partial summary into keyPartialSummaries

// 4) Prepare reduce prompt using collected partials and the question

// 4b) LLM final join

// 5) Finish node formats the final output

// Wiring

// Barrier: proceed to prepare_reduce only when all K partials are collected

// After prepare_reduce, call reduce then finish.

// nodeLoadAndChunk splits the preloaded document into overlapping chunks.
func (d *mapReduceDemo) nodeLoadAndChunk(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// nodeRetrieve selects top‑K chunks by simple lexical scoring against the question in user_input.
func (d *mapReduceDemo) nodeRetrieve(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// nodeCreateMapTasks fans out commands to map_summarize with per-chunk prompts.
func (d *mapReduceDemo) nodeCreateMapTasks(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Build per-task user input for the LLM summarizer node.

// nodeCollectPartial appends the last_response from map_summarize into partial_summaries.
func (d *mapReduceDemo) nodeCollectPartial(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// nodePrepareReduce creates a single user_input for the reduce LLM that contains
// the original question and each partial summary.
func (d *mapReduceDemo) nodePrepareReduce(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// nodeFinish exposes a consistent final message for the CLI.
func (d *mapReduceDemo) nodeFinish(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	// Prefer reduce_join output; fallback to last_response.
	return *new(any), nil
}

// Interactive shell: read question and run once per line.
func (d *mapReduceDemo) runInteractive(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *mapReduceDemo) runOnce(ctx context.Context, question string) error {
	_ = "STUB: not implemented"
	return nil
}

// Seed runtime state (doc_text is set by nodeLoadAndChunk from d.docText).

// Minimal event loop: print final answer when done.

// Try reduce_join output first.

// Helpers

// chunkText splits text into overlapping chunks of size chunkSize with given overlap.
func chunkText(text string, chunkSize, overlap int) []string { _ = "STUB: not implemented"; return nil }

// Advance by step = chunkSize - overlap

// scoreChunk computes a simple lexical score for how well chunk matches question.
func scoreChunk(question, chunk string) float64 {
	_ = "STUB: not implemented"
	// Lower‑case, split into alphanumerics, compute overlap count.
	return 0
}

// Normalize by sqrt length to dampen long chunks effect.

func tokenize(s string) []string { _ = "STUB: not implemented"; return nil }

// Replace non‑letters/digits with space

// Collapse spaces

func absOrSelf(p string) string { _ = "STUB: not implemented"; return "" }

// Small helpers to avoid import cycles with examples
func appendMapSliceReducer(existing, update any) any { _ = "STUB: not implemented"; return *new(any) }

// Minimal JSON unmarshal from []byte to T without adding a new dependency here.
func jsonUnmarshal(data []byte, out any) error { _ = "STUB: not implemented"; return nil }
