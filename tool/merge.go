//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tool

// Mergeable interface for custom types that want to define their own merging logic
type Mergeable interface {
	Merge(other any) any
}

// Merge merges a slice of items of the same type
// Supports: strings, numbers, byte slices, slices, arrays, maps, structs, and custom Mergeable types
func Merge[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// Handle the first element to determine the type and operation

// Check if type implements Mergeable interface

// For unsupported types, return the first element

// mergeStrings concatenates string values
func mergeStrings[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// mergeInts sums integer values
func mergeInts[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// Convert back to the original type

// mergeUints sums unsigned integer values
func mergeUints[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// Convert back to the original type

// mergeFloats sums floating point values
func mergeFloats[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// Convert back to the original type

// mergeSlices concatenates slice values
func mergeSlices[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// Handle special case for byte slices

// Generic slice concatenation

// mergeArrays concatenates array values into a new array
func mergeArrays[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// Note: Arrays are fixed size, so we assume all arrays in ts are of the same type and size.
// Plsease use slices if you need dynamic size.

// mergeMaps merges map values
func mergeMaps[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// mergeStructs merges struct values by combining fields using field-by-field merging
func mergeStructs[T any](ts []T) T { _ = "STUB: not implemented"; return *new(T) }

// Create a new struct instance

// Process each field

// Skip unexported fields

// Copy the value from the first struct for unexported fields

// Collect all field values from all structs

// Create a slice of the appropriate type and use Merge

// Convert fieldSlice to []any and use Merge directly

// Fallback: use the last value
