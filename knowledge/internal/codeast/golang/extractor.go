//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package golang

import (
	"go/ast"
	"go/token"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast"
)

type defaultExtractor struct {
	concurrency    int
	extractImports bool
}

func newDefaultExtractor(concurrency int, extractImports bool) *defaultExtractor {
	_ = "STUB: not implemented"
	return nil
}

func (e *defaultExtractor) Extract(input *extractInput) ([]*codeast.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *defaultExtractor) extractFile(pkg *parsedPackage, fset *token.FileSet, file *ast.File) []*codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *defaultExtractor) extractFunction(pkg *parsedPackage, fset *token.FileSet, decl *ast.FuncDecl, chunkIndex int) *codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *defaultExtractor) buildFunctionSignature(fset *token.FileSet, decl *ast.FuncDecl, receiverType string) string {
	_ = "STUB: not implemented"
	return ""
}

func writeFunctionPrefix(sig *strings.Builder, decl *ast.FuncDecl, receiverType string) {
	_ = "STUB: not implemented"
	return
}

func writeFunctionTypeParams(sig *strings.Builder, fset *token.FileSet, decl *ast.FuncDecl) {
	_ = "STUB: not implemented"
	return
}

func writeFunctionParams(sig *strings.Builder, fset *token.FileSet, decl *ast.FuncDecl) {
	_ = "STUB: not implemented"
	return
}

func writeReceiverSignature(sig *strings.Builder, decl *ast.FuncDecl, receiverType string) {
	_ = "STUB: not implemented"
	return
}

func writeResultSignature(sig *strings.Builder, fset *token.FileSet, results *ast.FieldList) {
	_ = "STUB: not implemented"
	return
}

func requiresResultParens(results *ast.FieldList) bool { _ = "STUB: not implemented"; return false }

func formatFieldList(fset *token.FileSet, fields *ast.FieldList) string {
	_ = "STUB: not implemented"
	return ""
}

func formatField(fset *token.FileSet, field *ast.Field) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *defaultExtractor) extractType(pkg *parsedPackage, fset *token.FileSet, spec *ast.TypeSpec, genDecl *ast.GenDecl, doc *ast.CommentGroup, chunkIndex int) []*codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *defaultExtractor) extractVariable(pkg *parsedPackage, fset *token.FileSet, spec *ast.ValueSpec, genDecl *ast.GenDecl, doc *ast.CommentGroup, chunkIndex int) []*codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func newNode(entityType codeast.EntityType, shortName, id, fullName, code, signature, comment, filePath string,
	lineStart, lineEnd, chunkIndex int,
) *codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func extractImportsFromASTFile(file *ast.File) []string { _ = "STUB: not implemented"; return nil }

func getCodeWithComment(fset *token.FileSet, node ast.Node, doc *ast.CommentGroup) string {
	_ = "STUB: not implemented"
	return ""
}

func getCodeWithGenDecl(fset *token.FileSet, spec ast.Node, genDecl *ast.GenDecl, doc *ast.CommentGroup) string {
	_ = "STUB: not implemented"
	return ""
}

func typeToString(fset *token.FileSet, expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

func fieldListToString(fset *token.FileSet, fields *ast.FieldList) string {
	_ = "STUB: not implemented"
	return ""
}

func receiverBaseTypeName(fset *token.FileSet, expr ast.Expr) string {
	_ = "STUB: not implemented"
	return ""
}
