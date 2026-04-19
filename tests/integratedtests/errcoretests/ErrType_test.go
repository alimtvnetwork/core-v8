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
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_ErrType_Combine_Verification(t *testing.T) {
	for caseIndex, testCase := range errTypeCombineTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")
		ref, _ := input.GetAsString("ref")
		errType := errcore.BytesAreNilOrEmptyType

		// Act
		result := errType.Combine(message, ref)

		// Assert
		testCase.ShouldBeRegex(
			t,
			caseIndex,
			result,
		)
	}
}

func Test_ErrCore_MergeErrors_Verification(t *testing.T) {
	for caseIndex, testCase := range errMergeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		hasError, _ := input.GetAsBool(params.hasError)

		var primaryErr, secondaryErr error
		if hasError {
			primaryErr = errors.New("test error")
		}

		// Act
		merged := errcore.MergeErrors(primaryErr, secondaryErr)
		isNil := fmt.Sprintf("%v", merged == nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			isNil,
		)
	}
}

func Test_ErrType_ErrorNoRefs_Verification(t *testing.T) {
	for caseIndex, testCase := range errTypeErrorNoRefsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")
		errType := errcore.BytesAreNilOrEmptyType

		// Act
		result := errType.ErrorNoRefs(message)
		hasError := fmt.Sprintf("%v", result != nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, hasError)
	}
}

func Test_ErrType_Error_Verification(t *testing.T) {
	for caseIndex, testCase := range errTypeErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")
		ref, _ := input.GetAsString("ref")
		errType := errcore.BytesAreNilOrEmptyType

		// Act
		result := errType.Error(message, ref)

		// Assert
		testCase.ShouldBeRegex(t, caseIndex, result.Error())
	}
}

func Test_MustBeEmpty_NilError_DoesNotPanic(t *testing.T) {
	// Should not panic
	errcore.MustBeEmpty(nil)
}

func Test_MustBeEmpty_WithError_Panics(t *testing.T) {
	// Arrange
	defer func() {
		recovered := recover()

	// Act
		actual := args.Map{"result": recovered == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected panic but did not get one", actual)
	}()

	errcore.MustBeEmpty(errors.New("must fail"))
}

func Test_HandleErr_NilError_DoesNotPanic(t *testing.T) {
	// Should not panic
	errcore.HandleErr(nil)
}
