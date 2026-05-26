//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package a2a

const (
	stateDeltaEnvelopeEncodingKey   = "encoding"
	stateDeltaEnvelopePayloadKey    = "payload"
	stateDeltaEnvelopeEncodingBytes = "bytes"
	stateDeltaEnvelopeEncodingNil   = "nil"
)

// EncodeStateDeltaMetadata converts Event.StateDelta into A2A metadata using a
// single lossless envelope format for every entry.
func EncodeStateDeltaMetadata(stateDelta map[string][]byte) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// DecodeStateDeltaMetadata restores Event.StateDelta from A2A metadata encoded
// by EncodeStateDeltaMetadata.
func DecodeStateDeltaMetadata(raw any) map[string][]byte { _ = "STUB: not implemented"; return nil }

func decodeStateDeltaEnvelope(entry map[string]any) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
