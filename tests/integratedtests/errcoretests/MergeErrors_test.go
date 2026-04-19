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

package errcoretests

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// buildErrorSlice creates []error from string messages interleaved with nils.
func buildErrorSlice(input args.Map) []error {
	errorMsgs, _ := input.GetAsStrings("errors")
	nilCount, _ := input.GetAsInt("nils")

	var errs []error
	// Interleave: nil first, then errors, then remaining nils
	halfNils := nilCount / 2
	for i := 0; i < halfNils; i++ {
		errs = append(errs, nil)
	}
	for _, msg := range errorMsgs {
		errs = append(errs, errors.New(msg))
	}
	for i := halfNils; i < nilCount; i++ {
		errs = append(errs, nil)
	}

	return errs
}

func Test_SliceToError_Verification(t *testing.T) {
	for caseIndex, testCase := range sliceToErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)

		var slice []string
		if !isNil {
			slice, _ = input.GetAsStrings("input")
		}

		// Act
		err := errcore.SliceToError(slice)

		// Assert — branch on whether ExpectedInput is args.Map or plain string
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			contain, _ := input.GetAsString("contain")
			actual := args.Map{
				"hasError":        fmt.Sprintf("%v", err != nil),
				"containsMessage": fmt.Sprintf("%v", err != nil && strings.Contains(err.Error(), contain)),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", err != nil))
		}
	}
}

func Test_SliceToErrorPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range sliceToErrorPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)

		var slice []string
		if !isNil {
			slice, _ = input.GetAsStrings("input")
		}

		// Act
		err := errcore.SliceToErrorPtr(slice)

		// Assert
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			contain, _ := input.GetAsString("contain")
			actual := args.Map{
				"hasError":        fmt.Sprintf("%v", err != nil),
				"containsMessage": fmt.Sprintf("%v", err != nil && strings.Contains(err.Error(), contain)),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", err != nil))
		}
	}
}

func Test_MergeErrors_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := buildErrorSlice(input)

		// Act
		merged := errcore.MergeErrors(errs...)

		// Assert
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			contain, _ := input.GetAsString("contain")
			actual := args.Map{
				"hasError":        fmt.Sprintf("%v", merged != nil),
				"containsMessage": fmt.Sprintf("%v", merged != nil && strings.Contains(merged.Error(), contain)),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", merged != nil))
		}
	}
}

func Test_SliceErrorsToStrings_Verification(t *testing.T) {
	for caseIndex, testCase := range sliceErrorsToStringsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := buildErrorSlice(input)

		// Act
		result := errcore.SliceErrorsToStrings(errs...)

		// Assert
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actual := args.Map{
				"count": fmt.Sprintf("%v", len(result)),
			}
			if len(result) > 0 {
				actual["first"] = result[0]
			}
			if len(result) > 1 {
				actual["second"] = result[1]
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", len(result)))
		}
	}
}

func Test_MergeErrorsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")
		errs := buildErrorSlice(input)

		// Act
		result := errcore.MergeErrorsToString(joiner, errs...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_MergeErrorsToStringDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsToStringDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := buildErrorSlice(input)

		// Act
		result := errcore.MergeErrorsToStringDefault(errs...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
