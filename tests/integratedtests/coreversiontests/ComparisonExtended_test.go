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

package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/corecmp"
	"github.com/alimtvnetwork/core-v8/corecomparator"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coreversion"
	"github.com/alimtvnetwork/core-v8/enums/versionindexes"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func Test_ComparisonValueIndexes_Verification(t *testing.T) {
	for caseIndex, testCase := range comparisonValueIndexesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, ok := input.GetAsString("left")
		if !ok {
			errcore.HandleErrMessage("left is required")
		}
		rightStr, ok := input.GetAsString("right")
		if !ok {
			errcore.HandleErrMessage("right is required")
		}

		// Act
		leftV := coreversion.New.Create(leftStr)
		rightV := coreversion.New.Create(rightStr)
		result := leftV.ComparisonValueIndexes(
			&rightV,
			versionindexes.AllVersionIndexes...,
		)

		// Assert
		actual := args.Map{
			"result": result.Name(),
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_VersionSliceInteger_Verification(t *testing.T) {
	for caseIndex, testCase := range versionSliceIntegerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, ok := input.GetAsString("left")
		if !ok {
			errcore.HandleErrMessage("left is required")
		}
		rightStr, ok := input.GetAsString("right")
		if !ok {
			errcore.HandleErrMessage("right is required")
		}

		// Act
		leftV := coreversion.New.Create(leftStr)
		rightV := coreversion.New.Create(rightStr)
		leftValues := leftV.AllVersionValues()
		rightValues := rightV.AllVersionValues()
		result := corecmp.VersionSliceInteger(leftValues, rightValues)

		// Assert
		actual := args.Map{
			"result": result.Name(),
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsAtLeast_Verification(t *testing.T) {
	for caseIndex, testCase := range isAtLeastTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, ok := input.GetAsString("left")
		if !ok {
			errcore.HandleErrMessage("left is required")
		}
		rightStr, ok := input.GetAsString("right")
		if !ok {
			errcore.HandleErrMessage("right is required")
		}

		// Act
		result := coreversion.IsAtLeast(leftStr, rightStr)

		// Assert
		actual := args.Map{
			"isAtLeast": result,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsLower_Verification(t *testing.T) {
	for caseIndex, testCase := range isLowerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, ok := input.GetAsString("left")
		if !ok {
			errcore.HandleErrMessage("left is required")
		}
		rightStr, ok := input.GetAsString("right")
		if !ok {
			errcore.HandleErrMessage("right is required")
		}

		// Act
		result := coreversion.IsLower(leftStr, rightStr)

		// Assert
		actual := args.Map{
			"isLower": result,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsExpectedVersion_Verification(t *testing.T) {
	for caseIndex, testCase := range isExpectedVersionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, ok := input.GetAsString("left")
		if !ok {
			errcore.HandleErrMessage("left is required")
		}
		rightStr, ok := input.GetAsString("right")
		if !ok {
			errcore.HandleErrMessage("right is required")
		}
		expected := input["expected"].(corecomparator.Compare)

		// Act
		result := coreversion.IsExpectedVersion(expected, leftStr, rightStr)

		// Assert
		actual := args.Map{
			"isExpected": result,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
