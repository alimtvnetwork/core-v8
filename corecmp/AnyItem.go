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

package corecmp

import "github.com/alimtvnetwork/core/corecomparator"

// AnyItem compares two any values for equality using Go's built-in == operator.
//
// Comparison logic (evaluated in order):
//  1. Both nil           → Equal
//  2. One nil, other not → NotEqual
//  3. left == right      → Equal (works for comparable built-in types)
//  4. Otherwise          → Inconclusive (types may not be comparable via ==)
//
// Returns:
//   - corecomparator.Equal        — both values are nil or structurally equal
//   - corecomparator.NotEqual     — exactly one value is nil
//   - corecomparator.Inconclusive — values differ or are not comparable via ==
//
// Note: This function does NOT perform deep comparison. Slices, maps, and
// structs without a comparable implementation will always return Inconclusive
// when they differ. Use reflect.DeepEqual or typed comparators for those cases.
func AnyItem(left, right any) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	if left == right {
		return corecomparator.Equal
	}

	return corecomparator.Inconclusive
}
