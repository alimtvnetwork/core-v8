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

package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ==========================================
// Test: IsEmpty
// ==========================================

func Test_LeftRight_IsEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsEmptyTestCases {
		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.IsEmpty()))
	}
}

// ==========================================
// Test: HasLeft
// ==========================================

func Test_LeftRight_HasLeft_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasLeftTestCases {
		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.HasLeft()))
	}
}

// ==========================================
// Test: HasRight
// ==========================================

func Test_LeftRight_HasRight_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasRightTestCases {
		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.HasRight()))
	}
}

// ==========================================
// Test: IsLeftEmpty
// ==========================================

func Test_LeftRight_IsLeftEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsLeftEmptyTestCases {
		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.IsLeftEmpty()))
	}
}

// ==========================================
// Test: IsRightEmpty
// ==========================================

func Test_LeftRight_IsRightEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsRightEmptyTestCases {
		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.IsRightEmpty()))
	}
}

// ==========================================
// Test: DeserializeLeft
// ==========================================

func Test_LeftRight_DeserializeLeft_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightDeserializeLeftTestCases {
		// Arrange
		result := tc.LR.DeserializeLeft()

		if result == nil {

		// Assert
			tc.Case.ShouldBeEqual(t, caseIndex, "true")
		} else {

		// Act
			actual := args.Map{
				"isNil":    result == nil,
				"hasError": result.HasError(),
			}

			tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

// ==========================================
// Test: DeserializeRight
// ==========================================

func Test_LeftRight_DeserializeRight_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightDeserializeRightTestCases {
		result := tc.LR.DeserializeRight()

		if result == nil {

		// Assert
			tc.Case.ShouldBeEqual(t, caseIndex, "true")
		} else {
			tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result == nil))
		}
	}
}

// ==========================================
// Test: TypeStatus
// ==========================================

func Test_LeftRight_TypeStatus_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightTypeStatusTestCases {
		// Arrange
		status := tc.LR.TypeStatus()

		// Act
		actual := args.Map{
			"isSame":             fmt.Sprintf("%v", status.IsSame),
			"isLeftUnknownNull":  fmt.Sprintf("%v", status.IsLeftUnknownNull),
			"isRightUnknownNull": fmt.Sprintf("%v", status.IsRightUnknownNull),
		}

		if expectedMap, ok := tc.Case.ExpectedInput.(args.Map); ok {
			if _, has := expectedMap["isLeftPointer"]; has {
				actual["isLeftPointer"] = fmt.Sprintf("%v", status.IsLeftPointer)
			}
			if _, has := expectedMap["isRightPointer"]; has {
				actual["isRightPointer"] = fmt.Sprintf("%v", status.IsRightPointer)
			}
			if _, has := expectedMap["isSameRegardlessPointer"]; has {
				actual["isSameRegardlessPointer"] = fmt.Sprintf("%v", status.IsSameRegardlessPointer())
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
