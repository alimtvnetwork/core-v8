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

func Test_If_String_Verification(t *testing.T) {
	for caseIndex, testCase := range ifStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValue, _ := input.GetAsString("trueValue")
		falseValue, _ := input.GetAsString("falseValue")

		// Act
		result := conditional.If[string](isTrue, trueValue, falseValue)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_If_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range ifIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValue, _ := input.GetAsInt("trueValue")
		falseValue, _ := input.GetAsInt("falseValue")

		// Act
		result := conditional.If[int](isTrue, trueValue, falseValue)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_NilDef_String_Verification(t *testing.T) {
	for caseIndex, testCase := range nilDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsString("defVal")

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.NilDef[string](ptr, defVal)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_IfFunc_String_Verification(t *testing.T) {
	for caseIndex, testCase := range ifFuncStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValue, _ := input.GetAsString("trueValue")
		falseValue, _ := input.GetAsString("falseValue")

		// Act
		result := conditional.IfFunc[string](
			isTrue,
			func() string { return trueValue },
			func() string { return falseValue },
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_IfTrueFunc_String_Verification(t *testing.T) {
	for caseIndex, testCase := range ifTrueFuncStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValue, _ := input.GetAsString("trueValue")

		// Act
		result := conditional.IfTrueFunc[string](
			isTrue,
			func() string { return trueValue },
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_IfSlice_String_Verification(t *testing.T) {
	for caseIndex, testCase := range ifSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValueRaw, _ := input.Get("trueValue")
		falseValueRaw, _ := input.Get("falseValue")
		trueValue := trueValueRaw.([]string)
		falseValue := falseValueRaw.([]string)

		// Act
		result := conditional.IfSlice[string](isTrue, trueValue, falseValue)
		actual := args.Map{
			"length": fmt.Sprintf("%v", len(result)),
			"first":  result[0],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NilCheck_Verification(t *testing.T) {
	for caseIndex, testCase := range nilCheckTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		// Act
		var canBeEmpty any
		if !isNil {
			canBeEmpty = "something"
		}
		result := conditional.NilCheck(canBeEmpty, onNil, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_DefOnNil_Verification(t *testing.T) {
	for caseIndex, testCase := range defOnNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNonNil, _ := input.GetAsString("onNonNil")

		// Act
		var canBeEmpty any
		if !isNil {
			val, _ := input.GetAsString("value")
			canBeEmpty = val
		}
		result := conditional.DefOnNil(canBeEmpty, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
