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

// ==========================================================================
// Test: IsMatch
// ==========================================================================

func Test_LazyRegex_IsMatch_FullDigit(t *testing.T) {
	tc := lazyRegexIsMatchFullDigitTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	actual := args.Map{
		params.isMatch: lazyRegex.IsMatch(compareInput),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LazyRegex_IsMatch_PartialMismatch(t *testing.T) {
	tc := lazyRegexIsMatchPartialMismatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	actual := args.Map{
		params.isMatch: lazyRegex.IsMatch(compareInput),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IsFailedMatch
// ==========================================================================

func Test_LazyRegex_IsFailedMatch_FromLazyRegexPatternMatc(t *testing.T) {
	tc := lazyRegexIsFailedMatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	actual := args.Map{
		params.isFailedMatch: lazyRegex.IsFailedMatch(compareInput),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: FirstMatchLine
// ==========================================================================

func Test_LazyRegex_FirstMatchLine_Found(t *testing.T) {
	tc := lazyRegexFirstMatchLineFoundTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	firstMatch, isInvalid := lazyRegex.FirstMatchLine(compareInput)

	actual := args.Map{
		params.firstMatch:     firstMatch,
		params.isInvalidMatch: isInvalid,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LazyRegex_FirstMatchLine_NotFound(t *testing.T) {
	tc := lazyRegexFirstMatchLineNotFoundTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	firstMatch, isInvalid := lazyRegex.FirstMatchLine(compareInput)

	actual := args.Map{
		params.firstMatch:     firstMatch,
		params.isInvalidMatch: isInvalid,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
