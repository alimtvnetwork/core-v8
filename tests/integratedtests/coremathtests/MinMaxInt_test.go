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

package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_MaxInt_Verification(t *testing.T) {
	for caseIndex, tc := range maxIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MaxInt(a, b)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinInt_Verification(t *testing.T) {
	for caseIndex, tc := range minIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MinInt(a, b)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MaxByte_Verification(t *testing.T) {
	for caseIndex, tc := range maxByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MaxByte(byte(a), byte(b))

		// Assert
		actual := args.Map{"result": int(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinByte_Verification(t *testing.T) {
	for caseIndex, tc := range minByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MinByte(byte(a), byte(b))

		// Assert
		actual := args.Map{"result": int(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToByte_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToByte(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToInt8_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToInt8TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt8(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToByte_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToByte(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
