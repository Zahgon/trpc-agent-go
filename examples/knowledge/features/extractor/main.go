//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates the Docling extractor for converting linked PDF
// and HTML content to markdown, and shows how to integrate it with knowledge
// sources.
//
// Prerequisites:
//   - A running Docling Serve instance (default: http://localhost:5001)
//     docker run -p 5001:5001 ghcr.io/docling-project/docling-serve
//
// Example usage:
//
//	cd examples/knowledge/features/extractor
//	go run main.go
//	go run main.go -endpoint http://localhost:5001 -output ./output
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/extractor"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/extractor/docling"

	// Register readers.
	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/markdown"
	pdfreader "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/pdf"
)

var (
	endpoint  = flag.String("endpoint", "http://localhost:5001", "Docling Serve endpoint")
	outputDir = flag.String("output", "./output", "Directory to save extracted markdown files")
)

const demoPDFURL = "https://arxiv.org/pdf/1706.03762"
const demoHTMLURL = "https://www.rfc-editor.org/rfc/rfc9110.html"

var demoURLs = []string{
	demoPDFURL,
	demoHTMLURL,
}

func main() {
	flag.Parse()
	ctx := context.Background()

	fmt.Println("Docling Extractor Demo")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Endpoint:   %s\n", *endpoint)
	fmt.Printf("Output Dir: %s\n", *outputDir)
	fmt.Printf("PDF URL:    %s\n", demoPDFURL)
	fmt.Printf("HTML URL:   %s\n", demoHTMLURL)
	fmt.Println(strings.Repeat("=", 60))

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output dir: %v", err)
	}

	// Create the Docling extractor.
	// By default, OCR is enabled and images use placeholder mode.
	ext := docling.New(
		docling.WithEndpoint(*endpoint),
		docling.WithTimeout(10*time.Minute),
	)
	defer ext.Close()

	fmt.Printf("\nSupported formats: %v\n", ext.SupportedFormats())

	// --- Part 0: Baseline with built-in PDF reader (no extractor) ---
	// Show what the built-in PDF text reader produces for the linked arXiv PDF.
	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Println("Part 0: Built-in PDF Reader from URL (no extractor, text only)")
	fmt.Println(strings.Repeat("-", 60))

	pdfReader := pdfreader.New()
	fmt.Printf("\n[1] Reading with built-in PDF reader: %s\n", demoPDFURL)

	startTime := time.Now()
	docs, err := pdfReader.ReadFromURL(demoPDFURL)
	elapsed := time.Since(startTime)
	if err != nil {
		log.Printf("  Read failed: %v", err)
	} else {
		// Combine all chunks into one text for saving.
		var fullText strings.Builder
		for _, doc := range docs {
			fullText.WriteString(doc.Content)
			fullText.WriteString("\n")
		}

		txtPath := filepath.Join(*outputDir, "1706.03762_pdfreader.txt")
		if err := os.WriteFile(txtPath, []byte(fullText.String()), 0644); err != nil {
			log.Printf("  Failed to write %s: %v", txtPath, err)
		} else {
			fmt.Printf("  Chunks:   %d\n", len(docs))
			fmt.Printf("  Output:   %s (%d bytes)\n", txtPath, fullText.Len())
			fmt.Printf("  Time:     %v\n", elapsed)
			fmt.Printf("  Preview:\n")
			printPreview(fullText.String(), 500)
		}
	}

	// --- Part 1: Direct extraction ---
	// Extract the linked PDF and HTML page to markdown using the Docling extractor directly.
	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Println("Part 1: Direct Extraction from URLs (PDF/HTML -> Markdown)")
	fmt.Println(strings.Repeat("-", 60))

	for i, rawURL := range demoURLs {
		fmt.Printf("\n[%d] Extracting: %s\n", i+1, rawURL)

		data, err := downloadURL(ctx, rawURL)
		if err != nil {
			log.Printf("  Failed to fetch %s: %v", rawURL, err)
			continue
		}
		fmt.Printf("  Input size: %d bytes\n", len(data))

		startTime := time.Now()
		result, err := ext.Extract(ctx, data)
		elapsed := time.Since(startTime)
		if err != nil {
			log.Printf("  Extraction failed (%.1fs): %v", elapsed.Seconds(), err)
			continue
		}

		content, err := io.ReadAll(result.Reader)
		if err != nil {
			log.Printf("  Failed to read result: %v", err)
			continue
		}

		mdName := outputBaseNameFromURL(rawURL) + "_docling.md"
		mdPath := filepath.Join(*outputDir, mdName)
		if err := os.WriteFile(mdPath, content, 0644); err != nil {
			log.Printf("  Failed to write %s: %v", mdPath, err)
			continue
		}

		fmt.Printf("  Format:   %s\n", result.Format)
		fmt.Printf("  Output:   %s (%d bytes)\n", mdPath, len(content))
		fmt.Printf("  Time:     %v\n", elapsed)
		fmt.Printf("  Preview:\n")
		printPreview(string(content), 500)
	}

	// --- Part 2: URL Source with extractor ---
	// Show how url.WithExtractor integrates Docling for linked PDF and HTML content,
	// and write chunk results into per-source files.
	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Println("Part 2: URL Source with Docling Extractor (PDF/HTML -> Markdown -> Chunked Files)")
	fmt.Println(strings.Repeat("-", 60))

	demonstrateKnowledgeSource(ctx, ext, demoURLs, *outputDir)

	fmt.Println("\nDone!")
}

// demonstrateKnowledgeSource shows how to use url.WithExtractor to produce
// chunked documents directly and write chunks from the same source into one file.
func demonstrateKnowledgeSource(ctx context.Context, ext extractor.Extractor, urls []string, outputDir string) {
	_ = "STUB: not implemented"
	return
}

func buildGroupedChunkFileContent(rawURL string, docs []anyDocument) string {
	_ = "STUB: not implemented"
	return ""
}

type anyDocument struct {
	Name     string
	Content  string
	Metadata map[string]any
	URL      string
}

func groupDocumentsBySourceURL(docs []*document.Document) map[string][]anyDocument {
	_ = "STUB: not implemented"
	return nil
}

func downloadURL(ctx context.Context, rawURL string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func outputBaseNameFromURL(rawURL string) string { _ = "STUB: not implemented"; return "" }

// printPreview prints the first n characters of content with indentation.
func printPreview(content string, n int) { _ = "STUB: not implemented"; return }
