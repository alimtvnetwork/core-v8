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

package coreinstructiontests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func newStringCompareFromMap(input args.Map) *coreinstruction.StringCompare {
	method, _ := input.GetAsString("method")
	search, _ := input.GetAsString("search")
	content, _ := input.GetAsString("content")

	isIgnoreCaseRaw, _ := input.Get("isIgnoreCase")
	isIgnoreCase, _ := isIgnoreCaseRaw.(bool)

	switch method {
	case "equal":
		return coreinstruction.NewStringCompareEqual(search, content)
	case "contains":
		return coreinstruction.NewStringCompareContains(isIgnoreCase, search, content)
	case "startsWith":
		return coreinstruction.NewStringCompareStartsWith(isIgnoreCase, search, content)
	case "endsWith":
		return coreinstruction.NewStringCompareEndsWith(isIgnoreCase, search, content)
	case "regex":
		return coreinstruction.NewStringCompareRegex(search, content)
	default:

		return coreinstruction.NewStringCompare(
			stringcompareas.Equal,
			isIgnoreCase,
			search,
			content,
		)
	}
}

func Test_StringCompare_IsMatch(t *testing.T) {
	for caseIndex, testCase := range stringCompareIsMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sc := newStringCompareFromMap(input)

		// Act
		actual := args.Map{
			"isMatch":       sc.IsMatch(),
			"isMatchFailed": sc.IsMatchFailed(),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

func Test_StringCompare_VerifyError(t *testing.T) {
	for caseIndex, testCase := range stringCompareVerifyErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sc := newStringCompareFromMap(input)

		// Act
		hasErr := fmt.Sprintf("%v", sc.VerifyError() != nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, hasErr)
	}
}

// ==========================================
// Nil Receiver — CaseNilSafe pattern
// ==========================================

func Test_StringCompare_NilReceiver(t *testing.T) {
	for caseIndex, tc := range stringCompareNilSafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
