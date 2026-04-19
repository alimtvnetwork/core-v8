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

package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_SimpleSliceValidator_SetActual_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorSetActualTestCase

	// Arrange
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a", "b"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}

	// Act
	result := v.SetActual([]string{"a", "b"})

	actual := args.Map{
		"sameInstance": result == v,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleSliceValidator_SliceValidator_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorSliceValidatorTestCase

	// Arrange
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual([]string{"a"})

	// Act
	sv := v.SliceValidator()

	actual := args.Map{
		"isNotNil": sv != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleSliceValidator_VerifyAll_FromSimpleSliceValidator(t *testing.T) {
	for caseIndex, tc := range simpleSliceValidatorVerifyAllTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		expectedLines := input["expected"].([]string)
		actualLines := input["actual"].([]string)

		expected := corestr.New.SimpleSlice.Direct(false, expectedLines)
		v := &corevalidator.SimpleSliceValidator{
			Expected:  expected,
			Condition: corevalidator.DefaultDisabledCoreCondition,
			CompareAs: stringcompareas.Equal,
		}
		v.SetActual(actualLines)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyAll(actualLines, params)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSliceValidator_VerifyFirst_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorVerifyFirstTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	expectedLines := input["expected"].([]string)
	actualLines := input["actual"].([]string)

	expected := corestr.New.SimpleSlice.Direct(false, expectedLines)
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual(actualLines)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := v.VerifyFirst(actualLines, params)

	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleSliceValidator_VerifyUpto_FromSimpleSliceValidator(t *testing.T) {
	tc := simpleSliceValidatorVerifyUptoTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	expectedLines := input["expected"].([]string)
	actualLines := input["actual"].([]string)
	length, _ := input.GetAsInt("length")

	expected := corestr.New.SimpleSlice.Direct(false, expectedLines)
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual(actualLines)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := v.VerifyUpto(actualLines, params, length)

	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
