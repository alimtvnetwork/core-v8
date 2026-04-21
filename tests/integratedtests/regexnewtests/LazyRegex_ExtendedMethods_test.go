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
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/regexnew"
)

// =============================================================================
// LazyRegex.FullString
// =============================================================================

func Test_LazyRegex_FullString_FromLazyRegexExtendedMet(t *testing.T) {
	for caseIndex, tc := range lazyRegexFullStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString(params.pattern)
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		result := lazy.FullString()

		actual := args.Map{
			params.isNotEmpty: result != "",
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.CompileMust
// =============================================================================

func Test_LazyRegex_CompileMust_FromLazyRegexExtendedMet(t *testing.T) {
	for caseIndex, tc := range lazyRegexCompileMustTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString(params.pattern)
		lazy := regexnew.New.LazyLock(pattern)

		var panicked bool
		var result *regexp.Regexp

		// Act
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			result = lazy.CompileMust()
		}()

		actual := args.Map{
			params.regexNotNil: result != nil,
			params.panicked:    panicked,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.FirstMatchLine
// =============================================================================

func Test_LazyRegex_FirstMatchLine_FromLazyRegexExtendedMet(t *testing.T) {
	for caseIndex, tc := range lazyRegexFirstMatchLineTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString(params.pattern)
		content, _ := input.GetAsString(params.content)
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		firstMatch, isInvalidMatch := lazy.FirstMatchLine(content)

		actual := args.Map{
			params.firstMatch:     firstMatch,
			params.isInvalidMatch: isInvalidMatch,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.IsFailedMatchBytes
// =============================================================================

func Test_LazyRegex_IsFailedMatchBytes_FromLazyRegexExtendedMet(t *testing.T) {
	for caseIndex, tc := range lazyRegexIsFailedMatchBytesTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString(params.pattern)
		inputStr, _ := input.GetAsString(params.input)
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		isFailed := lazy.IsFailedMatchBytes([]byte(inputStr))

		actual := args.Map{
			params.isFailed: isFailed,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.MatchUsingFuncError
// =============================================================================

func Test_LazyRegex_MatchUsingFuncError(t *testing.T) {
	matchFunc := func(regex *regexp.Regexp, lookingTerm string) bool {
		return regex.MatchString(lookingTerm)
	}

	for caseIndex, tc := range lazyRegexMatchUsingFuncErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString(params.pattern)
		comparing, _ := input.GetAsString(params.comparing)
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		err := lazy.MatchUsingFuncError(comparing, matchFunc)

		actual := args.Map{
			params.hasError: err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.OnRequiredCompiledMust
// =============================================================================

func Test_LazyRegex_OnRequiredCompiledMust_FromLazyRegexExtendedMet(t *testing.T) {
	for caseIndex, tc := range lazyRegexOnRequiredCompiledMustTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString(params.pattern)
		lazy := regexnew.New.LazyLock(pattern)

		var panicked bool

		// Act
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			lazy.OnRequiredCompiledMust()
		}()

		actual := args.Map{
			params.panicked: panicked,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// LazyRegex.MustBeSafe
// =============================================================================

func Test_LazyRegex_MustBeSafe_FromLazyRegexExtendedMet(t *testing.T) {
	for caseIndex, tc := range lazyRegexMustBeSafeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString(params.pattern)
		lazy := regexnew.New.LazyLock(pattern)

		var panicked bool

		// Act
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			lazy.MustBeSafe()
		}()

		actual := args.Map{
			params.panicked: panicked,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
