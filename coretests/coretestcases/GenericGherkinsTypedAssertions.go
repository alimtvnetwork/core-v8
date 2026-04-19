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

package coretestcases

import (
	"fmt"
	"testing"
)

// ShouldMatchExpected asserts that the given result matches
// the Expected field value. Uses the typed Expected field directly
// instead of converting to string lines.
//
// This is preferred over string-based comparison for boolean tests.
func (it *GenericGherkins[TInput, TExpect]) ShouldMatchExpected(
	t *testing.T,
	caseIndex int,
	result any,
) {
	t.Helper()

	title := it.CaseTitle()
	expected := fmt.Sprintf("%v", it.Expected)
	actual := fmt.Sprintf("%v", result)

	if actual == expected {
		return
	}

	t.Errorf(
		"Case %d (%s): got %s, want %s",
		caseIndex,
		title,
		actual,
		expected,
	)
}

// ShouldMatchExpectedFirst asserts that the given result matches
// the Expected field value using caseIndex=0. Use for named single
// test cases (non-loop).
func (it *GenericGherkins[TInput, TExpect]) ShouldMatchExpectedFirst(
	t *testing.T,
	result any,
) {
	t.Helper()

	it.ShouldMatchExpected(
		t,
		0,
		result,
	)
}
