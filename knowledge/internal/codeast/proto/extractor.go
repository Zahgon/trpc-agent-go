//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package proto

import (
	"github.com/bufbuild/protocompile/ast"
	parserpkg "github.com/bufbuild/protocompile/parser"
	"google.golang.org/protobuf/types/descriptorpb"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast"
)

type extractInput struct {
	fileNode     *ast.FileNode
	result       parserpkg.Result
	fileName     string
	protoPackage string
	syntax       string
	goPackage    string
	javaPackage  string
	imports      []string
	lines        []string
}

type defaultExtractor struct{}

func newDefaultExtractor() *defaultExtractor { _ = "STUB: not implemented"; return nil }

func (e *defaultExtractor) Extract(input *extractInput) ([]*codeast.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type entityExtractor struct {
	fileNode        *ast.FileNode
	result          parserpkg.Result
	fileName        string
	protoPackage    string
	syntax          string
	goPackage       string
	javaPackage     string
	imports         []string
	lines           []string
	allServiceNames []string
}

func (e *entityExtractor) extract() []*codeast.Node { _ = "STUB: not implemented"; return nil }

func (e *entityExtractor) collectAllServiceNames(fd *descriptorpb.FileDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (e *entityExtractor) addFileMetadata(node *codeast.Node, includeServices bool) {
	_ = "STUB: not implemented"
	return
}

func (e *entityExtractor) extractService(svc *descriptorpb.ServiceDescriptorProto, chunkIndex *int) []*codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *entityExtractor) collectRPCMethods(svc *descriptorpb.ServiceDescriptorProto) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *entityExtractor) extractRPC(method *descriptorpb.MethodDescriptorProto, svcFullName string, chunkIndex *int) *codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *entityExtractor) extractMessage(msg *descriptorpb.DescriptorProto, parentPkg string, chunkIndex *int) []*codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *entityExtractor) extractEnum(enum *descriptorpb.EnumDescriptorProto, parentPkg string, chunkIndex *int) *codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *entityExtractor) createEntityNode(code, name string, nodeType codeast.EntityType, fullName, comment, signature string, startLine, endLine int, chunkIndex *int) *codeast.Node {
	_ = "STUB: not implemented"
	return nil
}

func (e *entityExtractor) nodeLineRange(node ast.Node) (startLine, endLine int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (e *entityExtractor) extractCode(startLine, endLine int) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *entityExtractor) extractComment(node ast.Node) string {
	_ = "STUB: not implemented"
	return ""
}
