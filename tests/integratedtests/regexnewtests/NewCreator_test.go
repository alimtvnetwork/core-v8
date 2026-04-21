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

// =============================================================================
// New.Default
// =============================================================================

func Test_NewCreator_Default_Verification(t *testing.T) {
	for caseIndex, testCase := range newDefaultTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex, err := regexnew.New.Default(pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.DefaultLockIf
// =============================================================================

func Test_NewCreator_DefaultLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range newDefaultLockIfTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		isLock, _ := input.GetAsBool(params.isLock)

		// Act
		regex, err := regexnew.New.DefaultLockIf(isLock, pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.DefaultApplicableLock
// =============================================================================

func Test_NewCreator_DefaultApplicableLock_Verification(t *testing.T) {
	for caseIndex, testCase := range newDefaultApplicableLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex, err, isApplicable := regexnew.New.DefaultApplicableLock(pattern)

		actual := args.Map{
			params.regexNotNil:  regex != nil,
			params.hasError:     err != nil,
			params.isApplicable: isApplicable,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.LazyRegex.NewLockIf
// =============================================================================

func Test_NewLazyRegexCreator_NewLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range newLazyRegexNewLockIfTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)
		isLock, _ := input.GetAsBool(params.isLock)

		// Act
		lazy := regexnew.New.LazyRegex.NewLockIf(isLock, pattern)

		actual := args.Map{
			params.isDefined:    lazy.IsDefined(),
			params.isApplicable: lazy.IsApplicable(),
			params.isMatch:      lazy.IsMatch(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.LazyRegex.AllPatternsMap
// =============================================================================

func Test_NewLazyRegexCreator_AllPatternsMap_Verification(t *testing.T) {
	for caseIndex, testCase := range allPatternsMapTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// ensure at least one pattern exists in the map
		regexnew.New.LazyLock(pattern)

		// Act
		allPatterns := regexnew.New.LazyRegex.AllPatternsMap()

		actual := args.Map{
			params.isNotEmpty: len(allPatterns) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
