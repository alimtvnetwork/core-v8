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
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/issetter"
)

func Test_Value_IsOnLogically(t *testing.T) {
	for caseIndex, tc := range isOnLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actual := args.Map{"result": value.IsOnLogically()}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsOffLogically(t *testing.T) {
	for caseIndex, tc := range isOffLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actual := args.Map{"result": value.IsOffLogically()}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_WildcardApply(t *testing.T) {
	for caseIndex, tc := range wildcardApplyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		boolInput := input["input"].(bool)

		// Act
		actual := args.Map{"result": value.WildcardApply(boolInput)}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsWildcardOrBool(t *testing.T) {
	for caseIndex, tc := range isWildcardOrBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		boolInput := input["input"].(bool)

		// Act
		actual := args.Map{"result": value.IsWildcardOrBool(boolInput)}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_ToByteCondition_FromValueLogic(t *testing.T) {
	for caseIndex, tc := range toByteConditionTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		trueVal := input["trueVal"].(byte)
		falseVal := input["falseVal"].(byte)
		invalidVal := input["invalidVal"].(byte)

		// Act
		actual := args.Map{"result": int(value.ToByteCondition(trueVal, falseVal, invalidVal))}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_ToByteConditionWithWildcard(t *testing.T) {
	for caseIndex, tc := range toByteConditionWithWildcardTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		wildcardVal := input["wildcardVal"].(byte)
		trueVal := input["trueVal"].(byte)
		falseVal := input["falseVal"].(byte)
		invalidVal := input["invalidVal"].(byte)

		// Act
		actual := args.Map{"result": int(value.ToByteConditionWithWildcard(wildcardVal, trueVal, falseVal, invalidVal))}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsDefinedLogically_FromValueLogic(t *testing.T) {
	for caseIndex, tc := range isDefinedLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actual := args.Map{"result": value.IsDefinedLogically()}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsUndefinedLogically_FromValueLogic(t *testing.T) {
	for caseIndex, tc := range isUndefinedLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actual := args.Map{"result": value.IsUndefinedLogically()}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsPositive(t *testing.T) {
	for caseIndex, tc := range isPositiveTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actual := args.Map{"result": value.IsPositive()}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsNegative(t *testing.T) {
	for caseIndex, tc := range isNegativeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actual := args.Map{"result": value.IsNegative()}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_GetSetBoolOnInvalid(t *testing.T) {
	for caseIndex, tc := range getSetBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		setter := input["setter"].(bool)

		result := v.GetSetBoolOnInvalid(setter)

		// Act
		actual := args.Map{
			"result":        result,
			"isTrueOrFalse": v.IsTrue() || v.IsFalse(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_LazyEvaluateBool(t *testing.T) {
	for caseIndex, tc := range lazyEvaluateBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		called := false

		result := v.LazyEvaluateBool(func() { called = true })

		// Act
		actual := args.Map{
			"called":       called,
			"returnedTrue": result,
			"isTrue":       v.IsTrue(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_LazyEvaluateSet(t *testing.T) {
	for caseIndex, tc := range lazyEvaluateSetTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		called := false

		result := v.LazyEvaluateSet(func() { called = true })

		// Act
		actual := args.Map{
			"called":       called,
			"returnedTrue": result,
			"isSet":        v.IsSet(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
