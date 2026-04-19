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

package conditionaltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_ValueOrZero_String_Verification(t *testing.T) {
	for caseIndex, testCase := range valueOrZeroStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.ValueOrZero[string](ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_ValueOrZero_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range valueOrZeroIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		result := conditional.ValueOrZero[int](ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_ValueOrZero_Bool_Verification(t *testing.T) {
	for caseIndex, testCase := range valueOrZeroBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *bool
		if !isNil {
			val, _ := input.Get("value")
			boolVal := val == true
			ptr = &boolVal
		}

		result := conditional.ValueOrZero[bool](ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_PtrOrZero_String_Verification(t *testing.T) {
	for caseIndex, testCase := range ptrOrZeroStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.PtrOrZero[string](ptr)
		actual := args.Map{
			"isNotNil": fmt.Sprintf("%v", result != nil),
			"value":    *result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PtrOrZero_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range ptrOrZeroIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		result := conditional.PtrOrZero[int](ptr)
		actual := args.Map{
			"isNotNil": fmt.Sprintf("%v", result != nil),
			"value":    fmt.Sprintf("%v", *result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NilVal_String_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.NilVal[string](ptr, onNil, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_NilVal_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsInt("onNil")
		onNonNil, _ := input.GetAsInt("onNonNil")

		// Act
		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		result := conditional.NilVal[int](ptr, onNil, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_NilValPtr_String_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValPtrStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.NilValPtr[string](ptr, onNil, onNonNil)
		actual := args.Map{
			"isNotNil": fmt.Sprintf("%v", result != nil),
			"value":    *result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
