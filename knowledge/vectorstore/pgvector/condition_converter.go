//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package pgvector

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/searchfilter"
)

var comparisonOperators = map[string]string{
	searchfilter.OperatorEqual:              "=",
	searchfilter.OperatorNotEqual:           "!=",
	searchfilter.OperatorGreaterThan:        ">",
	searchfilter.OperatorGreaterThanOrEqual: ">=",
	searchfilter.OperatorLessThan:           "<",
	searchfilter.OperatorLessThanOrEqual:    "<=",
}

type condConvertResult struct {
	cond string
	args []any
}

// pgVectorConverter converts a filter condition to a postgres vector query.
type pgVectorConverter struct {
	metadataFieldName string
}

// Convert converts a filter condition to a postgres vector query filter.
func (c *pgVectorConverter) Convert(cond *searchfilter.UniversalFilterCondition) (*condConvertResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertFieldName converts metadata.xxx fields to JSONB syntax.
func (c *pgVectorConverter) convertFieldName(field string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *pgVectorConverter) convertCondition(cond *searchfilter.UniversalFilterCondition) (*condConvertResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pgVectorConverter) buildInCondition(cond *searchfilter.UniversalFilterCondition) (*condConvertResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pgVectorConverter) buildLogicalCondition(cond *searchfilter.UniversalFilterCondition) (*condConvertResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pgVectorConverter) buildComparisonCondition(cond *searchfilter.UniversalFilterCondition) (*condConvertResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pgVectorConverter) buildLikeCondition(cond *searchfilter.UniversalFilterCondition) (*condConvertResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pgVectorConverter) buildBetweenCondition(cond *searchfilter.UniversalFilterCondition) (*condConvertResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
