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

package issettertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

func Test_Value_New_Verification(t *testing.T) {
	for caseIndex, testCase := range valueNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(string)

		// Act
		val, err := issetter.New(input)

		actual := args.Map{
			"hasError": err != nil,
			"name":     val.Name(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_GetBool_Verification(t *testing.T) {
	for caseIndex, testCase := range getBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(bool)

		// Act
		val := issetter.GetBool(input)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_NewBool_Verification(t *testing.T) {
	for caseIndex, testCase := range newBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(bool)

		// Act
		val := issetter.NewBool(input)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_BooleanLogic_Verification(t *testing.T) {
	for caseIndex, testCase := range booleanLogicTestCases {
		// Arrange
		input := testCase.ArrangeInput.(issetter.Value)

		// Act
		actual := args.Map{
			"isOn":       input.IsOn(),
			"isOff":      input.IsOff(),
			"isTrue":     input.IsTrue(),
			"isFalse":    input.IsFalse(),
			"isSet":      input.IsSet(),
			"isUnset":    input.IsUnset(),
			"isValid":    input.IsValid(),
			"isWildcard": input.IsWildcard(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_CombinedBooleans_Verification(t *testing.T) {
	for caseIndex, testCase := range combinedBooleansTestCases {
		// Arrange
		input := testCase.ArrangeInput.([]bool)

		// Act
		val := issetter.CombinedBooleans(input...)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_Conversions_Verification(t *testing.T) {
	for caseIndex, testCase := range conversionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(issetter.Value)

		// Act
		actual := args.Map{
			"toBooleanValue":  input.ToBooleanValue().Name(),
			"toSetUnsetValue": input.ToSetUnsetValue().Name(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_GetSet_Verification(t *testing.T) {
	for caseIndex, testCase := range getSetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(getSetInput)

		// Act
		val := issetter.GetSet(input.condition, input.trueVal, input.falseVal)
		result := val.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Value_IsOutOfRange_Verification(t *testing.T) {
	for caseIndex, testCase := range isOutOfRangeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(byte)

		// Act
		result := fmt.Sprintf("%v", issetter.IsOutOfRange(input))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
