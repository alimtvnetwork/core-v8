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

package anycmptests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/anycmp"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_Cmp_SamePointer_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpSamePointerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_BothNil_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpBothNilTestCases {
		// Arrange & Act
		result := anycmp.Cmp(nil, nil)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_OneNil_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpOneNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_TypedNilBothNull_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpTypedNilBothNullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_TypedNilOneSide_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpTypedNilOneSideTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_BothNonNil_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpBothNonNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
