//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"reflect"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

var (
	timeType              = reflect.TypeOf(time.Time{})
	mapStringAnyType      = reflect.TypeOf(map[string]any(nil))
	mapStringBytesType    = reflect.TypeOf(map[string][]byte(nil))
	sliceAnyType          = reflect.TypeOf([]any(nil))
	sliceBytesType        = reflect.TypeOf([]byte(nil))
	sliceMessageOpsType   = reflect.TypeOf([]MessageOp(nil))
	modelMessagesType     = reflect.TypeOf([]model.Message(nil))
	modelContentPartsType = reflect.TypeOf([]model.ContentPart(nil))
	modelToolCallsType    = reflect.TypeOf([]model.ToolCall(nil))
)

// DeepCopier defines an interface for types that can perform deep copies of themselves.
type DeepCopier interface {
	// DeepCopy performs a deep copy of the object and returns a new copy.
	DeepCopy() any
}

type visitKind uint8

const (
	visitKindPointer visitKind = iota
	visitKindMap
	visitKindSlice
)

type visitKey struct {
	kind visitKind
	typ  reflect.Type
	ptr  uintptr
	len  int
}

type visitedMap map[visitKey]any

func newVisitedMap() visitedMap { _ = "STUB: not implemented"; return *new(visitedMap) }

func pointerVisitKey(ptr uintptr, typ reflect.Type) visitKey {
	_ = "STUB: not implemented"
	return *new(visitKey)
}

func mapVisitKey(ptr uintptr, typ reflect.Type) visitKey {
	_ = "STUB: not implemented"
	return *new(visitKey)
}

func sliceVisitKey(ptr uintptr, length int, typ reflect.Type) visitKey {
	_ = "STUB: not implemented"
	return *new(visitKey)
}

// deepCopyAny performs a deep copy of common JSON-serializable Go types to
// avoid sharing mutable references (maps/slices) across goroutines.
func deepCopyAny(value any) any { _ = "STUB: not implemented"; return *new(any) }

func deepCopyAnyWithVisited(value any, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// deepCopyFastPath handles common JSON-friendly types without reflection.
func deepCopyFastPath(value any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func deepCopyFastPathWithVisited(value any, visited visitedMap) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func deepCopyPrimitiveFastPath(value any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func deepCopyNumericFastPath(value any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func deepCopyMapStringAny(in map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func deepCopyMapStringAnyWithVisited(
	in map[string]any,
	visited visitedMap,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyMapStringBytes(in map[string][]byte) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyMapStringBytesWithVisited(
	in map[string][]byte,
	visited visitedMap,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func deepCopySliceAny(in []any) []any { _ = "STUB: not implemented"; return nil }

func deepCopySliceAnyWithVisited(in []any, visited visitedMap) []any {
	_ = "STUB: not implemented"
	return nil
}

func cloneSlice[T any](in []T) []T { _ = "STUB: not implemented"; return nil }

func cloneFastPathSlice[T any](in []T) []T { _ = "STUB: not implemented"; return nil }

func deepCopyMessageOps(in []MessageOp) ([]MessageOp, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func canDeepCopyMessageOpFastPath(op MessageOp) bool { _ = "STUB: not implemented"; return false }

func canDeepCopyMessageOpsFastPath(in []MessageOp) bool { _ = "STUB: not implemented"; return false }

func deepCopyMessageOpsWithVisited(
	in []MessageOp,
	visited visitedMap,
) ([]MessageOp, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func deepCopyMessageOp(op MessageOp) (MessageOp, bool) {
	_ = "STUB: not implemented"
	return *new(MessageOp), false
}

func deepCopyMessageOpWithVisited(
	op MessageOp,
	visited visitedMap,
) (MessageOp, bool) {
	_ = "STUB: not implemented"
	return *new(MessageOp), false
}

func deepCopyModelMessages(in []model.Message) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelMessagesWithVisited(
	in []model.Message,
	visited visitedMap,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelContentParts(in []model.ContentPart) []model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelContentPartsWithVisited(
	in []model.ContentPart,
	visited visitedMap,
) []model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelImage(in *model.Image) *model.Image { _ = "STUB: not implemented"; return nil }

func deepCopyModelImageWithVisited(
	in *model.Image,
	visited visitedMap,
) *model.Image {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelAudio(in *model.Audio) *model.Audio { _ = "STUB: not implemented"; return nil }

func deepCopyModelAudioWithVisited(
	in *model.Audio,
	visited visitedMap,
) *model.Audio {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelFile(in *model.File) *model.File { _ = "STUB: not implemented"; return nil }

func deepCopyModelFileWithVisited(
	in *model.File,
	visited visitedMap,
) *model.File {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelToolCalls(in []model.ToolCall) []model.ToolCall {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyModelToolCallsWithVisited(
	in []model.ToolCall,
	visited visitedMap,
) []model.ToolCall {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyStringPointerWithVisited(
	in *string,
	visited visitedMap,
) *string {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyIntPointerWithVisited(
	in *int,
	visited visitedMap,
) *int {
	_ = "STUB: not implemented"
	return nil
}

func deepCopyBytesWithVisited(
	in []byte,
	visited visitedMap,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

// deepCopyReflect performs a deep copy using reflection with cycle detection.
func deepCopyReflect(rv reflect.Value, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func copyInterface(rv reflect.Value, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func copyPointer(rv reflect.Value, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func copyMap(rv reflect.Value, visited visitedMap) any { _ = "STUB: not implemented"; return *new(any) }

func copySlice(rv reflect.Value, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func copyArray(rv reflect.Value, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func copyStruct(rv reflect.Value, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func isTimeType(rt reflect.Type) bool { _ = "STUB: not implemented"; return false }

func copyTime(value reflect.Value) any { _ = "STUB: not implemented"; return *new(any) }

// isJSONUnsafeKind reports whether a reflect.Kind cannot be handled
// by encoding/json (chan, func, unsafe pointer).
func isJSONUnsafeKind(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

func deepCopyByInterface(value any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func deepCopyByReflectValue(value reflect.Value) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func valueIsJSONUnsafe(value any) bool { _ = "STUB: not implemented"; return false }

// hasJSONUnsafeField returns true when a struct type contains exported
// fields that encoding/json cannot serialize.
func hasJSONUnsafeField(rt reflect.Type) bool { _ = "STUB: not implemented"; return false }

func hasJSONUnsafeType(rt reflect.Type, visiting map[reflect.Type]bool) bool {
	_ = "STUB: not implemented"
	return false
}

// jsonSafeCopy produces a deep copy of value that is safe for
// encoding/json.Marshal. Structs containing chan/func fields are
// converted to map[string]any with those fields omitted.
func jsonSafeCopy(value any) any { _ = "STUB: not implemented"; return *new(any) }

func jsonSafeCopyWithVisited(value any, visited visitedMap) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// jsonSafeFastPath handles common JSON-friendly types without
// reflection, delegating nested values to jsonSafeCopyWithVisited.
// For maps, unsafe values are dropped to match jsonSafeCopyMap behavior.
func jsonSafeFastPath(value any, visited visitedMap) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Skip non-serializable values.

// jsonSafeReflect is like deepCopyReflect but converts structs that
// contain non-serializable fields into map[string]any representations
// so that the result is always safe for json.Marshal.
func jsonSafeReflect(
	rv reflect.Value,
	visited visitedMap,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Drop non-serializable values entirely.

func jsonSafeCopyPointer(
	rv reflect.Value,
	visited visitedMap,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Cache a placeholder before descending to break pointer cycles.

func jsonSafeCopyMap(
	rv reflect.Value,
	visited visitedMap,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Skip non-serializable map values.

func jsonSafeCopySlice(
	rv reflect.Value,
	visited visitedMap,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func jsonSafeCopyArray(
	rv reflect.Value,
	visited visitedMap,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// jsonSafeCopyStruct converts a struct to map[string]any when it
// contains non-serializable fields; otherwise deep-copies normally.
func jsonSafeCopyStruct(
	rv reflect.Value,
	visited visitedMap,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// No unsafe fields; deep-copy preserving original type.

// structToJSONSafeMap converts a struct value into a map[string]any,
// skipping fields whose types are not JSON-serializable.
func structToJSONSafeMap(
	rv reflect.Value,
	visited visitedMap,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Skip chan/func/unsafe-pointer fields.

// shouldSkipJSONField reports whether a struct field should be ignored
// when checking or generating JSON-safe outputs.
func shouldSkipJSONField(ft reflect.StructField) bool { _ = "STUB: not implemented"; return false }

// mapValueIsJSONUnsafe reports whether a map value is a non-serializable
// value that should be removed from JSON-safe output.
func mapValueIsJSONUnsafe(value reflect.Value) bool { _ = "STUB: not implemented"; return false }

// splitJSONTag splits a json struct tag value on commas.
func splitJSONTag(tag string) []string { _ = "STUB: not implemented"; return nil }
