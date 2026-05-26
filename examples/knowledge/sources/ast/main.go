//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates AST-aware ingestion through repo source on a
// mixed-language repository.
package main

import (
	"context"
	"flag"
	"fmt"

	util "trpc.group/trpc-go/trpc-agent-go/examples/knowledge"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
)

var (
	vectorStore  = flag.String("vectorstore", "inmemory", "Vector store type: inmemory|pgvector|sqlitevec|tcvector|elasticsearch|milvus")
	dumpDir      = flag.String("dumpdir", "chunked", "Output directory for parsed documents; defaults to ./chunked and preserves source file structure under go-reader/proto-reader/repo-source")
	goRepoURL    = flag.String("gorepo", "https://github.com/trpc-group/trpc-go", "Remote Go repository URL used for Go reader and repo source demo")
	embedderMode = flag.String("embedder", "mock", "Embedder mode: auto|mock|openai. mock is useful for chunk preview and local validation without real embeddings")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	goRepo := *goRepoURL
	protoRepo := util.ExampleDataPath("ast/proto-lib")

	fmt.Println("🔮 AST Source Demo (Repo Source)")
	fmt.Println("================================")
	fmt.Printf("Go repository URL: %s\n", goRepo)
	fmt.Printf("Proto repository root: %s\n", protoRepo)

	fmt.Println("\n📦 Step 1: Repo source preview on a repository")
	fmt.Println("----------------------------------------------")
	demonstrateRepoSource(ctx, goRepo, protoRepo)
}

func demonstrateRepoSource(ctx context.Context, goRepoURL string, protoRepo string) {
	_ = "STUB: not implemented"
	return
}

func printDocumentPreview(docs []*document.Document, limit int) { _ = "STUB: not implemented"; return }

func printMetadataByPrefix(metadata map[string]any) { _ = "STUB: not implemented"; return }

func shorten(text string, limit int) string { _ = "STUB: not implemented"; return "" }

func dumpDocuments(section, root string, docs []*document.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func renderDocumentGroupDump(relPath string, docs []*document.Document) string {
	_ = "STUB: not implemented"
	return ""
}

func relativeSourcePath(root string, doc *document.Document) string {
	_ = "STUB: not implemented"
	return ""
}

func sanitizePathSegment(value string) string { _ = "STUB: not implemented"; return "" }

func toString(value any) string { _ = "STUB: not implemented"; return "" }

func toInt(value any) int { _ = "STUB: not implemented"; return 0 }

func chooseEmbedder(mode string, apiKey string) (embedder.Embedder, string, error) {
	_ = "STUB: not implemented"
	return *new(embedder.Embedder), "", nil
}

type mockEmbedder struct{}

func (mockEmbedder) GetEmbedding(ctx context.Context, text string) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m mockEmbedder) GetEmbeddingWithUsage(ctx context.Context, text string) ([]float64, map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (mockEmbedder) GetDimensions() int { _ = "STUB: not implemented"; return 0 }

func formatEmbeddingTextForDump(text string) string { _ = "STUB: not implemented"; return "" }
