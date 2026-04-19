// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package coregeneric

// =============================================================================
// Pair[L, R] — type aliases for common combinations
// =============================================================================

type StringStringPair = Pair[string, string]
type StringIntPair = Pair[string, int]
type StringInt64Pair = Pair[string, int64]
type StringFloat64Pair = Pair[string, float64]
type StringBoolPair = Pair[string, bool]
type StringAnyPair = Pair[string, any]
type IntIntPair = Pair[int, int]
type IntStringPair = Pair[int, string]
type IntAnyPair = Pair[int, any]
type AnyAnyPair = Pair[any, any]

// =============================================================================
// Triple[A, B, C] — type aliases for common combinations
// =============================================================================

type StringStringStringTriple = Triple[string, string, string]
type StringIntStringTriple = Triple[string, int, string]
type StringAnyAnyTriple = Triple[string, any, any]
type AnyAnyAnyTriple = Triple[any, any, any]
