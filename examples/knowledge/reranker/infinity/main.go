//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// This example demonstrates how to use Infinity Reranker and compare with embedding similarity.
//
// Integration Guide:
//
//	reranker, err := infinity.New(
//	    infinity.WithEndpoint("http://custom-host:7997/rerank"),
//	    infinity.WithTopN(5),
//	)
//	if err != nil {
//	    // handle error
//	}
//	k := knowledge.New(
//	    knowledge.WithReranker(reranker),
//	)
//
// Required environment variables:
//   - OPENAI_API_KEY: Your OpenAI API key for embeddings
//   - OPENAI_BASE_URL: (Optional) Custom OpenAI API endpoint
//   - INFINITY_URL: (Optional) Infinity reranker endpoint
package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/reranker"
)

var (
	endpoint       = flag.String("endpoint", getDefaultEndpoint(), "Infinity endpoint URL (default: INFINITY_URL env or http://localhost:7997/rerank)")
	modelName      = flag.String("model", "bge-reranker-v2-m3", "Reranker model name")
	embeddingModel = flag.String("embedding-model", "text-embedding-3-small", "Embedding model name")
)

func getDefaultEndpoint() string { _ = "STUB: not implemented"; return "" }

type testCase struct {
	name      string
	query     string
	documents []string
}

func main() {
	flag.Parse()
	ctx := context.Background()

	fmt.Printf("Using reranker endpoint: %s\n", *endpoint)
	fmt.Printf("Using embedding model: %s\n", *embeddingModel)
	fmt.Println()

	testCases := []testCase{
		{
			name:  "Lexical Overlap Trap",
			query: "How to kill a Python process?",
			documents: []string{
				"Use kill -9 PID or pkill python to terminate a Python process.",
				"Python is a non-venomous snake that kills prey by constriction.",
				"Python programming language was created by Guido van Rossum.",
				"The process of learning Python takes about 3 months.",
				"Kill is a Unix command to send signals to processes.",
			},
		},
		{
			name:  "Semantic Precision",
			query: "What year was Bitcoin created?",
			documents: []string{
				"Bitcoin was created in 2009 by Satoshi Nakamoto.",
				"Bitcoin is a decentralized digital currency.",
				"The Bitcoin whitepaper was published in 2008.",
				"Bitcoin mining requires significant computational power.",
				"Cryptocurrency has grown significantly since 2010.",
			},
		},
		{
			name:  "Implicit Answer",
			query: "Can I use React without Node.js?",
			documents: []string{
				"React can be included via CDN script tags without any build tools.",
				"Node.js is commonly used for React development.",
				"React is a JavaScript library for building user interfaces.",
				"npm is the package manager for Node.js.",
				"Create React App requires Node.js to be installed.",
			},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("\n%s\n", strings.Repeat("=", 70))
		fmt.Printf("Case: %s\n", tc.name)
		fmt.Printf("Query: %s\n", tc.query)
		fmt.Printf("%s\n", strings.Repeat("=", 70))

		runComparison(ctx, tc.query, tc.documents)
	}
}

func runComparison(ctx context.Context, queryText string, documents []string) {
	_ = "STUB: not implemented"
	return
}

// 1. Calculate embedding similarity scores

// 2. Rerank with Infinity

func printRerankerResults(results []*reranker.Result) { _ = "STUB: not implemented"; return }
