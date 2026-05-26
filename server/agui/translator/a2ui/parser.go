// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package a2ui

// parser splits streaming text into JSONL records.
//
// Each non-empty line is treated as one JSONL message.
// Lines may arrive in fragments across multiple Append calls.
// Trailing "\r" from CRLF line endings is removed.
// Blank lines are ignored.
type parser struct {
	pending string
}

// newParser creates a new JSONL parser.
func newParser() *parser {
	_ = "STUB: not implemented"

	// append appends streaming text and returns all completed JSONL lines.
	//
	// Incomplete trailing data is buffered until more text arrives or flush is called.
	return nil
}

func (p *parser) append(text string) []string { _ = "STUB: not implemented"; return nil }

// flush returns the final buffered line, if any, and resets the pending state.
//
// Blank final content is ignored.
func (p *parser) flush() []string { _ = "STUB: not implemented"; return nil }

// reset clears all buffered state.
func (p *parser) reset() { _ = "STUB: not implemented"; return }

func (p *parser) consume(data string, flush bool) []string { _ = "STUB: not implemented"; return nil }

func normalizeLine(s string) string { _ = "STUB: not implemented"; return "" }
