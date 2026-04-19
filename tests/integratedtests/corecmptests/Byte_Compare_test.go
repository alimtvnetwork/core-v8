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

package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Byte_Compare_Verification(t *testing.T) {
	for caseIndex, testCase := range byteCompareTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsInt("left")
		right, _ := input.GetAsInt("right")

		// Act
		result := corecmp.Byte(byte(left), byte(right))

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsStringsEqualWithoutOrder_Verification(t *testing.T) {
	for caseIndex, testCase := range isStringsEqualWithoutOrderTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsStrings("left")
		right, _ := input.GetAsStrings("right")

		// Act
		result := corecmp.IsStringsEqualWithoutOrder(left, right)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_VersionSliceByte_Verification(t *testing.T) {
	for caseIndex, testCase := range versionSliceByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftRaw, _ := input.Get("left")
		rightRaw, _ := input.Get("right")

		var left, right []byte
		if leftRaw != nil {
			for _, v := range leftRaw.([]int) {
				left = append(left, byte(v))
			}
		}
		if rightRaw != nil {
			for _, v := range rightRaw.([]int) {
				right = append(right, byte(v))
			}
		}

		// Act
		result := corecmp.VersionSliceByte(left, right)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_VersionSliceInteger_Verification(t *testing.T) {
	for caseIndex, testCase := range versionSliceIntegerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftRaw, _ := input.Get("left")
		rightRaw, _ := input.Get("right")

		var left, right []int
		if leftRaw != nil {
			left = leftRaw.([]int)
		}
		if rightRaw != nil {
			right = rightRaw.([]int)
		}

		// Act
		result := corecmp.VersionSliceInteger(left, right)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsStringsEqual_BothNil_Verification(t *testing.T) {
	for caseIndex, testCase := range isStringsEqualBothNilTestCases {
		// Arrange & Act
		result := corecmp.IsStringsEqual(nil, nil)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsStringsEqual_OneNil_Verification(t *testing.T) {
	for caseIndex, testCase := range isStringsEqualOneNilTestCases {
		// Arrange & Act
		result := corecmp.IsStringsEqual([]string{"a"}, nil)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Compare_Extended_Methods(t *testing.T) {
	for caseIndex, testCase := range compareExtendedMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(value)

		// Act
		actual := args.Map{
			"isLess":                compare.IsLess(),
			"isLessEqual":          compare.IsLessEqual(),
			"isGreater":            compare.IsGreater(),
			"isGreaterEqual":       compare.IsGreaterEqual(),
			"isDefined":            compare.IsDefined(),
			"isInconclusive":       compare.IsInconclusive(),
			"isNotEqual":           compare.IsNotEqual(),
			"isNotEqualLogically":  compare.IsNotEqualLogically(),
			"isDefinedProperly":    compare.IsDefinedProperly(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
