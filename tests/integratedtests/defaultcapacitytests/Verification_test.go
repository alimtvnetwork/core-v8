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

package defaultcapacitytests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/defaultcapacity"
)

func Test_OfSplits_Verification(t *testing.T) {
	for caseIndex, testCase := range ofSplitsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		wholeLength := input.GetAsIntDefault("wholeLength", 0)
		limit := input.GetAsIntDefault("limit", 0)

		// Act
		result := defaultcapacity.OfSplits(wholeLength, limit)

		actual := args.Map{}
		expected := testCase.ExpectedInput.(args.Map)

		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		if _, has := expected["isPositive"]; has {
			actual["isPositive"] = result > 0
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PredictiveDefaultSmall_Verification(t *testing.T) {
	for caseIndex, testCase := range predictiveDefaultSmallTestCases {
		// Arrange
		input := testCase.ArrangeInput.(int)

		// Act
		result := defaultcapacity.PredictiveDefaultSmall(input)

		actual := args.Map{
			"isPositive": result > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PredictiveFiftyPercentIncrement_Verification(t *testing.T) {
	for caseIndex, testCase := range predictiveFiftyPercentTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		possibleLen := input.GetAsIntDefault("possibleLen", 0)
		additionalCap := input.GetAsIntDefault("additionalCap", 0)

		// Act
		result := defaultcapacity.PredictiveFiftyPercentIncrement(possibleLen, additionalCap)

		actual := args.Map{}
		expected := testCase.ExpectedInput.(args.Map)

		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		if _, has := expected["isPositive"]; has {
			actual["isPositive"] = result > 0
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
