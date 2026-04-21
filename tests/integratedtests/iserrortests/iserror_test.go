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

package iserrortests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/iserror"
)

// ==========================================
// Test: Empty / Defined / NotEmpty
// ==========================================

func Test_Empty_Defined_NotEmpty_Verification(t *testing.T) {
	errorsToTest := []error{nil, errSample1}

	for caseIndex, testCase := range emptyTestCases {
		// Arrange
		err := errorsToTest[caseIndex]

		// Act
		actual := args.Map{
			"isEmpty":    fmt.Sprintf("%v", iserror.Empty(err)),
			"isDefined":  fmt.Sprintf("%v", iserror.Defined(err)),
			"isNotEmpty": fmt.Sprintf("%v", iserror.NotEmpty(err)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Equal / NotEqual
// ==========================================

func Test_Equal_NotEqual_Verification(t *testing.T) {
	type errorPair struct {
		left, right error
	}

	pairs := []errorPair{
		{errSample1, errSample1},
		{nil, nil},
		{nil, errSample1},
		{errSame, errSameDup},
		{errSample1, errSample2},
	}

	for caseIndex, testCase := range equalTestCases {
		// Arrange
		pair := pairs[caseIndex]

		// Act
		actual := args.Map{
			"isEqual":    fmt.Sprintf("%v", iserror.Equal(pair.left, pair.right)),
			"isNotEqual": fmt.Sprintf("%v", iserror.NotEqual(pair.left, pair.right)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: AllDefined
// ==========================================

func Test_AllDefined_Verification(t *testing.T) {
	inputs := [][]error{
		{errSample1, errSample2, errSame},
		{errSample1, nil, errSame},
		{},
	}

	for caseIndex, testCase := range allDefinedTestCases {
		// Arrange
		errs := inputs[caseIndex]

		// Act
		actLines := []string{
			fmt.Sprintf("%v", iserror.AllDefined(errs...)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AnyDefined
// ==========================================

func Test_AnyDefined_Verification(t *testing.T) {
	inputs := [][]error{
		{nil, nil, errSample1},
		{nil, nil, nil},
		{},
	}

	for caseIndex, testCase := range anyDefinedTestCases {
		// Arrange
		errs := inputs[caseIndex]

		// Act
		actLines := []string{
			fmt.Sprintf("%v", iserror.AnyDefined(errs...)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AllEmpty
// ==========================================

func Test_AllEmpty_Verification(t *testing.T) {
	inputs := [][]error{
		{nil, nil, nil},
		{nil, errSample1, nil},
		{},
	}

	for caseIndex, testCase := range allEmptyTestCases {
		// Arrange
		errs := inputs[caseIndex]

		// Act
		actLines := []string{
			fmt.Sprintf("%v", iserror.AllEmpty(errs...)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AnyEmpty
// ==========================================

func Test_AnyEmpty_Verification(t *testing.T) {
	inputs := [][]error{
		{errSample1, nil, errSample2},
		{errSample1, errSample2},
		{},
	}

	for caseIndex, testCase := range anyEmptyTestCases {
		// Arrange
		errs := inputs[caseIndex]

		// Act
		actLines := []string{
			fmt.Sprintf("%v", iserror.AnyEmpty(errs...)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: EqualString / NotEqualString
// ==========================================

func Test_EqualString_NotEqualString_Verification(t *testing.T) {
	for caseIndex, testCase := range equalStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		actual := args.Map{
			"isEqual":    fmt.Sprintf("%v", iserror.EqualString(left, right)),
			"isNotEqual": fmt.Sprintf("%v", iserror.NotEqualString(left, right)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
