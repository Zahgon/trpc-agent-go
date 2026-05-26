//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Copyright The OpenTelemetry Authors
// Copyright (C) 2025 Tencent. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
//

package tracetransform

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
)

// KeyValues transforms a slice of attribute KeyValues into OTLP key-values.
func KeyValues(attrs []attribute.KeyValue) []*commonpb.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// Iterator transforms an attribute iterator into OTLP key-values.
func Iterator(iter attribute.Iterator) []*commonpb.KeyValue { _ = "STUB: not implemented"; return nil }

// ResourceAttributes transforms a Resource OTLP key-values.
func ResourceAttributes(res *resource.Resource) []*commonpb.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// KeyValue transforms an attribute KeyValue into an OTLP key-value.
func KeyValue(kv attribute.KeyValue) *commonpb.KeyValue { _ = "STUB: not implemented"; return nil }

// Value transforms an attribute Value into an OTLP AnyValue.
func Value(v attribute.Value) *commonpb.AnyValue { _ = "STUB: not implemented"; return nil }

func boolSliceValues(vals []bool) []*commonpb.AnyValue { _ = "STUB: not implemented"; return nil }

func int64SliceValues(vals []int64) []*commonpb.AnyValue { _ = "STUB: not implemented"; return nil }

func float64SliceValues(vals []float64) []*commonpb.AnyValue { _ = "STUB: not implemented"; return nil }

func stringSliceValues(vals []string) []*commonpb.AnyValue { _ = "STUB: not implemented"; return nil }
