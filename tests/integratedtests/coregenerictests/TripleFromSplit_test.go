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
)

// ==========================================
// Test: TripleFromSplit
// ==========================================

func Test_TripleFromSplit_Verification(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplit(inputStr, sep)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TripleFromSplitTrimmed
// ==========================================

func Test_TripleFromSplitTrimmed_Verification(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplitTrimmed(inputStr, sep)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TripleFromSplitN
// ==========================================

func Test_TripleFromSplitN_Verification(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitNTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplitN(inputStr, sep)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TripleFromSplitNTrimmed
// ==========================================

func Test_TripleFromSplitNTrimmed_Verification(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitNTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplitNTrimmed(inputStr, sep)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TripleFromSlice
// ==========================================

func Test_TripleFromSlice_Verification(t *testing.T) {
	for caseIndex, testCase := range tripleFromSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		parts, _ := input.GetAsStrings("parts")

		// Act
		triple := coregeneric.TripleFromSlice(parts)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
