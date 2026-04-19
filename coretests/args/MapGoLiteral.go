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
	"strings"
)

// GoLiteralLines returns the map entries formatted as Go literal lines,
// suitable for copy-pasting into test case definitions.
//
// Output format:
//
//	"key": value,
//
// Values are formatted using %#v for strings (quoted) and %v for others.
//
// Example:
//
//	m := args.Map{"isZero": false, "value": 5, "name": "hello"}
//	m.GoLiteralLines()
//	// Returns:
//	//   "isZero": false,
//	//   "name":   "hello",
//	//   "value":  5,
func (it Map) GoLiteralLines() []string {
	if len(it) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(it))

	for k := range it {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Find max key length for alignment
	maxKeyLen := 0
	for _, k := range keys {
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}

	lines := make([]string, len(keys))

	for i, key := range keys {
		padding := strings.Repeat(" ", maxKeyLen-len(key))
		val := it[key]

		var valStr string

		switch v := val.(type) {
		case string:
			valStr = fmt.Sprintf("%q", v)
		default:
			valStr = fmt.Sprintf("%v", v)
		}

		lines[i] = fmt.Sprintf(
			"%q: %s%s,",
			key,
			padding,
			valStr,
		)
	}

	return lines
}

// GoLiteralString returns the map formatted as a multi-line Go literal block.
//
// Example output:
//
//	"isZero": false,
//	"name":   "hello",
//	"value":  5,
func (it Map) GoLiteralString() string {
	lines := it.GoLiteralLines()

	return strings.Join(lines, "\n")
}
