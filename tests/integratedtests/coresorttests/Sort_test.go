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

package coresorttests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coresort/intsort"
	"github.com/alimtvnetwork/core-v8/coresort/strsort"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// =============================================================================
// intsort.Quick — ascending
// =============================================================================

func Test_IntSort_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		clone := make([]int, len(original))
		copy(clone, original)

		// Act
		result := intsort.Quick(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// intsort.QuickDsc — descending
// =============================================================================

func Test_IntSort_QuickDsc_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickDscTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		clone := make([]int, len(original))
		copy(clone, original)

		// Act
		result := intsort.QuickDsc(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// intsort.QuickPtr — ascending pointer sort
// =============================================================================

func Test_IntSort_QuickPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		ptrs := toIntPtrs(original)

		// Act
		result := intsort.QuickPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatIntPtrs(*result))
	}
}

// =============================================================================
// intsort.QuickDscPtr — descending pointer sort
// =============================================================================

func Test_IntSort_QuickDscPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickDscPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		ptrs := toIntPtrs(original)

		// Act
		result := intsort.QuickDscPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatIntPtrs(*result))
	}
}

// =============================================================================
// strsort.Quick — ascending
// =============================================================================

func Test_StrSort_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		clone := make([]string, len(inputVal))
		copy(clone, inputVal)

		// Act
		result := strsort.Quick(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// strsort.QuickDsc — descending
// =============================================================================

func Test_StrSort_QuickDsc_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickDscTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		clone := make([]string, len(inputVal))
		copy(clone, inputVal)

		// Act
		result := strsort.QuickDsc(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// strsort.QuickPtr — ascending pointer sort
// =============================================================================

func Test_StrSort_QuickPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		ptrs := toStrPtrs(inputVal)

		// Act
		result := strsort.QuickPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatStrPtrs(*result))
	}
}

// =============================================================================
// strsort.QuickDscPtr — descending pointer sort
// =============================================================================

func Test_StrSort_QuickDscPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickDscPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		ptrs := toStrPtrs(inputVal)

		// Act
		result := strsort.QuickDscPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatStrPtrs(*result))
	}
}
