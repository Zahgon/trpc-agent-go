//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package golang provides internal Go AST parsing for code-aware knowledge ingestion.
package golang

import (
	"go/ast"
	"go/token"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast"
)

// Parser parses Go source content into code-aware AST nodes.
type Parser struct {
	extractor codeast.Extractor[*extractInput]
	analyzer  codeast.Analyzer[*analyzeInput]
}

type extractInput struct {
	pkg  *parsedPackage
	fset *token.FileSet
}

type analyzeInput struct {
	pkg *parsedPackage
}

type parserConfig struct {
	concurrency    int
	extractImports bool
}

// Option is a functional option for configuring the parser.
type Option func(*parserConfig)

// WithConcurrency sets the concurrency for parallel extraction.
func WithConcurrency(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtractImports enables or disables extracting file-level imports.
func WithExtractImports(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewParser creates a new Go AST parser.
func NewParser(opts ...Option) *Parser { _ = "STUB: not implemented"; return nil }

// ParseContent parses Go source content and returns semantic nodes plus reserved edge slots.
func (p *Parser) ParseContent(name, content string) (*codeast.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseDirectory parses a Go directory/module and returns semantic nodes across all files.
func (p *Parser) ParseDirectory(dirPath string) (*codeast.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func modulePathForDir(baseDir, dir string) string { _ = "STUB: not implemented"; return "" }

// ParseFileInfo extracts file-level metadata without requiring a full semantic extraction result.
func (p *Parser) ParseFileInfo(name, content string) (*codeast.FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type parsedPackage struct {
	ID     string
	Name   string
	Syntax []*ast.File
	Fset   *token.FileSet
}

func buildFileInfo(name string, fileNode *ast.File) *codeast.FileInfo {
	_ = "STUB: not implemented"
	return nil
}

// BuildNodeEmbeddingText builds the embedding payload for a parsed Go node.
func BuildNodeEmbeddingText(node *codeast.Node) string { _ = "STUB: not implemented"; return "" }

// BuildFileEmbeddingText builds the embedding payload for a whole Go file document.
func BuildFileEmbeddingText(content, name, packagePath string, imports []string) string {
	_ = "STUB: not implemented"
	return ""
}

func resolvePackagePath(fileName, packageName string) string { _ = "STUB: not implemented"; return "" }

func looksLikeLocalPath(name string) bool { _ = "STUB: not implemented"; return false }

func findNearestGoModule(startDir string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func parseGoModulePath(goModPath string) string { _ = "STUB: not implemented"; return "" }
