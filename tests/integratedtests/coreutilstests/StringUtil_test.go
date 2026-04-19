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

package coreutilstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

func Test_IsEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range isEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsEmpty(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsBlank_Verification(t *testing.T) {
	for caseIndex, testCase := range isBlankTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsBlank(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsEmptyOrWhitespace_Verification(t *testing.T) {
	for caseIndex, testCase := range isEmptyOrWhitespaceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsEmptyOrWhitespace(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_SafeSubstring_Verification(t *testing.T) {
	for caseIndex, testCase := range safeSubstringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		start, _ := input.GetAsInt("start")
		end, _ := input.GetAsInt("end")

		// Act
		result := stringutil.SafeSubstring(content, start, end)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_SplitLeftRight_Verification(t *testing.T) {
	for caseIndex, testCase := range splitLeftRightTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		separator, _ := input.GetAsString("separator")

		// Act
		left, right := stringutil.SplitLeftRight(inputStr, separator)
		actual := args.Map{
			"left":  left,
			"right": right,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsStartsWith_Verification(t *testing.T) {
	for caseIndex, testCase := range isStartsWithTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		startsWith, _ := input.GetAsString("startsWith")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase := isIgnoreCaseVal == true

		// Act
		result := stringutil.IsStartsWith(content, startsWith, isIgnoreCase)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsEndsWith_Verification(t *testing.T) {
	for caseIndex, testCase := range isEndsWithTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		endsWith, _ := input.GetAsString("endsWith")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase := isIgnoreCaseVal == true

		// Act
		result := stringutil.IsEndsWith(content, endsWith, isIgnoreCase)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_RemoveMany_Verification(t *testing.T) {
	for caseIndex, testCase := range removeManyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		removesRaw, _ := input.Get("removes")
		removes := removesRaw.([]string)

		// Act
		result := stringutil.RemoveMany(content, removes...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_ReplaceWhiteSpacesToSingle_Verification(t *testing.T) {
	for caseIndex, testCase := range replaceWhiteSpacesToSingleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.ReplaceWhiteSpacesToSingle(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
