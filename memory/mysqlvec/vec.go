//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mysqlvec

// serializeVector converts a float64 embedding into the binary format
// accepted by MySQL 9.0+ VECTOR type (little-endian float32 sequence).
func serializeVector(embedding []float64) []byte { _ = "STUB: not implemented"; return nil }

// vectorToString converts a float64 embedding to MySQL STRING_TO_VECTOR
// format: "[1.0, 2.0, 3.0, ...]". Used for MySQL 9.0+ vector queries.
func vectorToString(embedding []float64) string { _ = "STUB: not implemented"; return "" }

// cosineSimilarity computes the cosine similarity between two float64 vectors.
// Returns 0 if either vector has zero magnitude.
func cosineSimilarity(a, b []float64) float64 { _ = "STUB: not implemented"; return 0 }

// deserializeVector converts a little-endian float32 binary blob back to
// float64 slice. Used for brute-force search on MySQL 8.x.
func deserializeVector(data []byte) ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }
