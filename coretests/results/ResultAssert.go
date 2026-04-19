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

package results

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ExpectAnyError is a sentinel error used in Expected results
// to indicate that the method is expected to return any non-nil error.
//
// Usage:
//
//	Expected: results.ResultAny{
//	    Error: results.ExpectAnyError,
//	},
var ExpectAnyError = fmt.Errorf("expect-any-error")

// ShouldMatchResult compares the Result against an expected ResultAny.
//
// The set of compared fields is auto-derived from the expected result:
//   - "panicked" — always compared
//   - "value"    — compared if expected.Value != nil
//   - "hasError" — compared if expected.Error != nil
//   - "returnCount" — compared if expected.ReturnCount != 0
//   - "isSafe"   — compared if explicitly requested via compareFields
//
// To override auto-detection, pass explicit field names in compareFields.
// When compareFields is non-empty, only those fields are compared.
func (it Result[T]) ShouldMatchResult(
	t *testing.T,
	caseIndex int,
	title string,
	expected ResultAny,
	compareFields ...string,
) {
	t.Helper()

	fields := compareFields

	if len(fields) == 0 {
		fields = deriveCompareFields(expected)
	}

	actualMap := it.ToMap()
	expectedMap := expected.ToMap()

	actFiltered := filterByFields(actualMap, fields)
	expFiltered := filterByFields(expectedMap, fields)

	actLines := actFiltered.CompileToStrings()
	expLines := expFiltered.CompileToStrings()

	errcore.AssertDiffOnMismatch(
		t,
		caseIndex,
		title,
		actLines,
		expLines,
	)
}

// deriveCompareFields auto-detects which fields to compare
// based on the expected ResultAny.
func deriveCompareFields(expected ResultAny) []string {
	fields := []string{"panicked"}

	if expected.Value != nil {
		fields = append(fields, "value")
	}

	if expected.Error != nil {
		fields = append(fields, "hasError")
	}

	if expected.ReturnCount != 0 {
		fields = append(fields, "returnCount")
	}

	return fields
}

// filterByFields returns a new args.Map containing only the specified keys.
func filterByFields(m args.Map, fields []string) args.Map {
	filtered := args.Map{}

	for _, key := range fields {
		if val, exists := m[key]; exists {
			filtered[key] = val
		} else {
			filtered[key] = fmt.Sprintf("<missing key: %s>", key)
		}
	}

	return filtered
}
