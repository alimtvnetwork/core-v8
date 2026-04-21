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

package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/regexnew"
)

func Test_LazyRegex_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexCompileTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		regex, err := lazyRegex.Compile()

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
			params.isCompiled:  lazyRegex.IsCompiled(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_HasError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexHasErrorTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.hasError:  lazyRegex.HasError(),
			params.isInvalid: lazyRegex.IsInvalid(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_MatchBytes_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchBytesTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.isMatchBytes:       lazyRegex.IsMatchBytes([]byte(compareInput)),
			params.isFailedMatchBytes: lazyRegex.IsFailedMatchBytes([]byte(compareInput)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_MatchError_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexMatchErrorTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)
		matchErr := lazyRegex.MatchError(compareInput)

		actual := args.Map{
			params.isNoError: matchErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LazyRegex_String_Verification(t *testing.T) {
	for caseIndex, testCase := range lazyRegexStringTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		lazyRegex := regexnew.New.LazyLock(pattern)

		actual := args.Map{
			params.stringResult: lazyRegex.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
