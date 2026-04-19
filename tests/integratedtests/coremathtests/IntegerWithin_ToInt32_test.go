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

func Test_IntegerWithin_ToInt32_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToInt32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt32(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToUnsignedInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToUint16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToUnsignedInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToUnsignedInt32_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToUint32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToUnsignedInt32(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToUnsignedInt64_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToUint64TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToUnsignedInt64(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToInt8_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToInt8TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToInt8(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToInt32_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToInt32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToInt32(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToUnsignedInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToUint16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToUnsignedInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToUnsignedInt64_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToUint64TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToUnsignedInt64(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MaxFloat32_Verification(t *testing.T) {
	for caseIndex, tc := range maxFloat32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		aRaw, _ := input.Get("a")
		bRaw, _ := input.Get("b")
		a := float32(aRaw.(float64))
		b := float32(bRaw.(float64))

		// Act
		result := coremath.MaxFloat32(a, b)

		// Assert
		actual := args.Map{"result": float64(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinFloat32_Verification(t *testing.T) {
	for caseIndex, tc := range minFloat32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		aRaw, _ := input.Get("a")
		bRaw, _ := input.Get("b")
		a := float32(aRaw.(float64))
		b := float32(bRaw.(float64))

		// Act
		result := coremath.MinFloat32(a, b)

		// Assert
		actual := args.Map{"result": float64(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
