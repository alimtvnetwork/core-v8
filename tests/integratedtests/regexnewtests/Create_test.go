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

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range createTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex, err := regexnew.New.DefaultLock(pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Create_IsMatchLock_Verification(t *testing.T) {
	for caseIndex, testCase := range createIsMatchLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		actual := args.Map{
			params.isMatch: regexnew.IsMatchLock(pattern, compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Create_IsMatchFailed_Verification(t *testing.T) {
	for caseIndex, testCase := range createIsMatchFailedTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		actual := args.Map{
			params.isFailed: regexnew.IsMatchFailed(pattern, compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Test: MatchError
// ==========================================================================

func Test_MatchError_Match_FromCreate(t *testing.T) {
	tc := matchErrorMatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchError(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MatchError_Mismatch(t *testing.T) {
	tc := matchErrorMismatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchError(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: MatchErrorLock
// ==========================================================================

func Test_MatchErrorLock_Match_FromCreate(t *testing.T) {
	tc := matchErrorLockMatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchErrorLock(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MatchErrorLock_Mismatch(t *testing.T) {
	tc := matchErrorLockMismatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchErrorLock(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
