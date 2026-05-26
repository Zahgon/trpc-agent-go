//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package elasticsearch provides Elasticsearch-based vector storage implementation.
package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/searchfilter"
)

// esConverter converts a filter condition to an Elasticsearch query.
type esConverter struct {
	metadataFieldName string
}

// convertFieldName converts metadata.xxx fields to ES field path.
func (c *esConverter) convertFieldName(field string) string { _ = "STUB: not implemented"; return "" }

// Convert converts a filter condition to an Elasticsearch query filter.
func (c *esConverter) Convert(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) convertCondition(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) buildLogicalCondition(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

// OperatorOr

func (c *esConverter) buildComparisonCondition(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) convertEqual(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) convertNotEqual(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) convertRange(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) buildBetweenCondition(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) buildInCondition(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

func (c *esConverter) buildLikeCondition(cond *searchfilter.UniversalFilterCondition) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}
