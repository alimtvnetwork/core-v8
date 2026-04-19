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

func newStringSearchFromMap(input args.Map) *coreinstruction.StringSearch {
	method, _ := input.GetAsString("method")
	search, _ := input.GetAsString("search")

	var compareMethod stringcompareas.Variant
	switch method {
	case "equal":
		compareMethod = stringcompareas.Equal
	case "contains":
		compareMethod = stringcompareas.Contains
	case "startsWith":
		compareMethod = stringcompareas.StartsWith
	case "endsWith":
		compareMethod = stringcompareas.EndsWith
	case "regex":
		compareMethod = stringcompareas.Regex
	default:
		compareMethod = stringcompareas.Equal
	}

	return &coreinstruction.StringSearch{
		CompareMethod: compareMethod,
		Search:        search,
	}
}

func Test_StringSearch_IsMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchIsMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ss := newStringSearchFromMap(input)
		content, _ := input.GetAsString("content")

		// Act
		actual := args.Map{
			"isMatch":       ss.IsMatch(content),
			"isMatchFailed": ss.IsMatchFailed(content),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

func Test_StringSearch_IsAllMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchIsAllMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ss := newStringSearchFromMap(input)
		contents, _ := input.GetAsStrings("contents")

		// Act
		actual := args.Map{
			"isAllMatch":       ss.IsAllMatch(contents...),
			"isAnyMatchFailed": ss.IsAnyMatchFailed(contents...),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

func Test_StringSearch_State_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil, _ := isNilVal.(bool)

		var ss *coreinstruction.StringSearch
		if !isNil {
			ss = newStringSearchFromMap(input)
		}

		// Act
		actual := args.Map{
			"isEmpty": ss.IsEmpty(),
			"isExist": ss.IsExist(),
			"has":     ss.Has(),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

func Test_StringSearch_VerifyError_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchVerifyErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil, _ := isNilVal.(bool)
		content, _ := input.GetAsString("content")

		var ss *coreinstruction.StringSearch
		if !isNil {
			ss = newStringSearchFromMap(input)
		}

		// Act
		err := ss.VerifyError(content)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", err != nil),
		)
	}
}
