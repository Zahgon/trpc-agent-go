//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package proto provides internal proto AST parsing logic for knowledge documents.
package proto

import (
	"google.golang.org/protobuf/types/descriptorpb"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast"
)

// Parser parses proto files into internal codeast results.
type Parser struct {
	extractor codeast.Extractor[*extractInput]
	analyzer  codeast.Analyzer[*analyzeInput]
}

// NewParser creates a new proto parser.
func NewParser() *Parser { _ = "STUB: not implemented"; return nil }

// ParseContent parses a proto file content and extracts all AST entities.
func (p *Parser) ParseContent(name, content string) (*codeast.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildRPCSignature builds a human-readable RPC signature.
func BuildRPCSignature(name, inputType, outputType string, clientStreaming, serverStreaming bool) string {
	_ = "STUB: not implemented"
	return ""
}

// QualifiedName returns the fully qualified name with proto package prefix.
func QualifiedName(pkg, name string) string { _ = "STUB: not implemented"; return "" }

// QualifiedNameWithParent returns the fully qualified name with parent scope prefix.
func QualifiedNameWithParent(name, parent string) string { _ = "STUB: not implemented"; return "" }

// ExtractCode extracts code from source lines by start/end line numbers.
func ExtractCode(lines []string, startLine, endLine int) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildNodeEmbeddingText builds embedding payload for a proto node.
func BuildNodeEmbeddingText(node *codeast.Node) string { _ = "STUB: not implemented"; return "" }

// BuildFileEmbeddingText builds embedding payload for file-level documents.
func BuildFileEmbeddingText(content, fileName string, fileMetadata map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// ExtractFileMetadata parses proto content and returns normalized file-level metadata.
func ExtractFileMetadata(content string) map[string]any { _ = "STUB: not implemented"; return nil }

// ExtractOptionString extracts option string values like go_package/java_package.
func ExtractOptionString(content, optionName string) string { _ = "STUB: not implemented"; return "" }

// ExtractFileOptions extracts go_package/java_package from descriptor options.
func ExtractFileOptions(fd *descriptorpb.FileDescriptorProto) (goPackage, javaPackage string) {
	_ = "STUB: not implemented"
	return "", ""
}

// ShortTypeName returns short name from a qualified type name.
func ShortTypeName(typeName string) string { _ = "STUB: not implemented"; return "" }

func extractFileMetadataFromDescriptor(fd *descriptorpb.FileDescriptorProto) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
