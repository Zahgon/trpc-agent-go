//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package rouge

import (
	"sync"

	"github.com/neurosnap/sentences"
)

var (
	// englishSentenceTokenizerOnce ensures the Punkt model is loaded once.
	englishSentenceTokenizerOnce sync.Once
	// englishSentenceTokenizer holds the initialized sentence tokenizer instance.
	englishSentenceTokenizer *sentences.DefaultSentenceTokenizer
	// englishSentenceTokenizerErr caches any initialization error.
	englishSentenceTokenizerErr error
)

// nltkSentTokenizeEnglish splits English text into sentences using Punkt training data.
// This function aims to match NLTK's sent_tokenize behavior for rougeLsum sentence splitting.
func nltkSentTokenizeEnglish(text string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// splitLeadingStandalonePeriodsLikeNLTK splits leading standalone periods into separate sentences.
// NLTK's PunktSentenceTokenizer treats ". ." patterns as standalone sentences in several edge cases.
func splitLeadingStandalonePeriodsLikeNLTK(s string) []string {
	_ = "STUB: not implemented"
	return nil
}

// isWhitespaceASCII reports whether the byte is an ASCII whitespace character.
func isWhitespaceASCII(b byte) bool { _ = "STUB: not implemented"; return false }
