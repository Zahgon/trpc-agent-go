//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package jsonrepair

// isHex reports whether the rune is a hexadecimal digit.
func isHex(char rune) bool { _ = "STUB: not implemented"; return false }

// isDigit reports whether the rune is an ASCII digit.
func isDigit(char rune) bool { _ = "STUB: not implemented"; return false }

// isValidStringCharacter reports whether the rune is allowed unescaped in a JSON string.
func isValidStringCharacter(char rune) bool { _ = "STUB: not implemented"; return false }

// isDelimiter reports whether the rune is treated as a delimiter by the parser.
func isDelimiter(char rune) bool { _ = "STUB: not implemented"; return false }

// isFunctionNameCharStart reports whether the rune can start an identifier.
func isFunctionNameCharStart(char rune) bool { _ = "STUB: not implemented"; return false }

// isFunctionNameChar reports whether the rune can be part of an identifier.
func isFunctionNameChar(char rune) bool { _ = "STUB: not implemented"; return false }

// isURLStart reports whether the text ends with a supported scheme prefix ending in "://".
func isURLStart(text string) bool { _ = "STUB: not implemented"; return false }

// isURLChar reports whether the rune is allowed inside a URL token.
func isURLChar(char rune) bool { _ = "STUB: not implemented"; return false }

// isUnquotedStringDelimiter reports whether the rune ends an unquoted string token.
func isUnquotedStringDelimiter(char rune) bool { _ = "STUB: not implemented"; return false }

// isStartOfValue reports whether the rune can start a JSON value.
func isStartOfValue(char rune) bool { _ = "STUB: not implemented"; return false }

// isControlCharacter reports whether the rune is a control character that must be escaped.
func isControlCharacter(char rune) bool { _ = "STUB: not implemented"; return false }

// isWhitespace reports whether the rune is treated as whitespace.
func isWhitespace(char rune) bool { _ = "STUB: not implemented"; return false }

// isWhitespaceExceptNewline reports whether the rune is whitespace excluding newline.
func isWhitespaceExceptNewline(char rune) bool { _ = "STUB: not implemented"; return false }

// isSpecialWhitespace reports whether the rune is a special whitespace character normalized to a space.
func isSpecialWhitespace(char rune) bool { _ = "STUB: not implemented"; return false }

// isQuote reports whether the rune is a supported quote character.
func isQuote(char rune) bool { _ = "STUB: not implemented"; return false }

// isDoubleQuoteLike reports whether the rune is a double-quote-like character.
func isDoubleQuoteLike(char rune) bool { _ = "STUB: not implemented"; return false }

// isDoubleQuote reports whether the rune is a double quote character.
func isDoubleQuote(char rune) bool {
	_ = "STUB: not implemented"

	// isSingleQuoteLike reports whether the rune is a single-quote-like character.
	return false
}

func isSingleQuoteLike(char rune) bool { _ = "STUB: not implemented"; return false }

// isSingleQuote reports whether the rune is a single quote character.
func isSingleQuote(char rune) bool { _ = "STUB: not implemented"; return false }

// stripLastOccurrence removes the last occurrence of char from text.
func stripLastOccurrence(text []rune, char rune, stripRemainingText bool) []rune {
	_ = "STUB: not implemented"
	return nil
}

// insertBeforeLastWhitespace inserts runes before any trailing whitespace in the text.
func insertBeforeLastWhitespace(text []rune, insert []rune) []rune {
	_ = "STUB: not implemented"
	return nil
}

// removeAtIndex removes count runes from text starting at start.
func removeAtIndex(text []rune, start int, count int) []rune { _ = "STUB: not implemented"; return nil }

// endsWithCommaOrNewline reports whether text ends with a comma or newline, ignoring trailing spaces.
func endsWithCommaOrNewline(text []rune) bool { _ = "STUB: not implemented"; return false }
