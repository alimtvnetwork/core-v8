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

package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ==========================================
// Test: PairFromSplit
// ==========================================

func Test_PairFromSplit(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplit(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSplitTrimmed
// ==========================================

func Test_PairFromSplitTrimmed(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplitTrimmed(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSplitFull
// ==========================================

func Test_PairFromSplitFull(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitFullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplitFull(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSplitFullTrimmed
// ==========================================

func Test_PairFromSplitFullTrimmed(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitFullTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplitFullTrimmed(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSlice
// ==========================================

func Test_PairFromSlice(t *testing.T) {
	for caseIndex, testCase := range pairFromSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		parts, ok := input.GetAsStrings("parts")
		if !ok {
			errcore.HandleErrMessage("parts is required")
		}

		// Act
		pair := coregeneric.PairFromSlice(parts)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
