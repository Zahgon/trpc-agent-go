//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package searchfilter

// Equal creates a condition for equality comparison.
func Equal(field string, value any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// NotEqual creates a condition for inequality comparison.
func NotEqual(field string, value any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// GreaterThan creates a condition for greater than comparison.
func GreaterThan(field string, value any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// GreaterThanOrEqual creates a condition for greater than or equal comparison.
func GreaterThanOrEqual(field string, value any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// LessThan creates a condition for less than comparison.
func LessThan(field string, value any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// LessThanOrEqual creates a condition for less than or equal comparison.
func LessThanOrEqual(field string, value any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// In creates a condition for IN operator.
func In(field string, values ...any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// NotIn creates a condition for NOT IN operator.
func NotIn(field string, values ...any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// Like creates a condition for LIKE operator (pattern matching).
func Like(field string, pattern string) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// NotLike creates a condition for NOT LIKE operator.
func NotLike(field string, pattern string) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// Between creates a condition for BETWEEN operator.
func Between(field string, min, max any) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// And creates a condition that combines multiple conditions with AND logic.
func And(conditions ...*UniversalFilterCondition) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// Or creates a condition that combines multiple conditions with OR logic.
func Or(conditions ...*UniversalFilterCondition) *UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}
