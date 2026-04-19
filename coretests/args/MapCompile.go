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

package args

import (
	"fmt"
	"sort"
)

// CompileToStrings converts all map values to strings using %v format
// and returns sorted "key : value" lines.
//
// This enables test cases to store raw typed values (int, bool, etc.)
// while producing deterministic, human-readable string lines for comparison.
//
// Example:
//
//	m := args.Map{"isZero": false, "value": 5}
//	m.CompileToStrings()
//	// Returns: []string{"isZero : false", "value : 5"}
func (it Map) CompileToStrings() []string {
	if len(it) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(it))

	for k := range it {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	lines := make([]string, len(keys))

	for i, key := range keys {
		lines[i] = fmt.Sprintf(
			"%s : %v",
			key,
			it[key],
		)
	}

	return lines
}

// CompileToString converts all map values to a single sorted string
// with one "key : value" per line.
//
// Useful for quick debugging or single-string comparison.
func (it Map) CompileToString() string {
	lines := it.CompileToStrings()

	result := ""

	for i, line := range lines {
		if i > 0 {
			result += "\n"
		}

		result += line
	}

	return result
}
