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
	"os"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// =============================================================================
// errors.Is verification for MergeErrors (errors.Join)
// =============================================================================

func Test_MergeErrors_ErrorsIs_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsIsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sentinel := input["sentinel"].(error)
		errs := input["errors"].([]error)

		// Act
		merged := errcore.MergeErrors(errs...)

		// Assert
		actual := args.Map{
			"hasError":   fmt.Sprintf("%v", merged != nil),
			"errorsIsOk": fmt.Sprintf("%v", errors.Is(merged, sentinel)),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// errors.Is verification for ConcatMessageWithErr (fmt.Errorf %w)
// =============================================================================

func Test_ConcatMessageWithErr_ErrorsIs_Verification(t *testing.T) {
	for caseIndex, testCase := range concatMessageErrorsIsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sentinel := input["sentinel"].(error)
		msg, _ := input.GetAsString("message")

		// Act
		wrapped := errcore.ConcatMessageWithErr(msg, sentinel)

		// Assert
		actual := args.Map{
			"hasError":        fmt.Sprintf("%v", wrapped != nil),
			"errorsIsOk":      fmt.Sprintf("%v", errors.Is(wrapped, sentinel)),
			"containsMessage": fmt.Sprintf("%v", wrapped != nil && len(wrapped.Error()) > len(sentinel.Error())),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// errors.As verification for MergeErrors
// =============================================================================

func Test_MergeErrors_ErrorsAs_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsAsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := input["errors"].([]error)

		// Act
		merged := errcore.MergeErrors(errs...)

		// Assert
		var pathErr *os.PathError
		actual := args.Map{
			"hasError":   fmt.Sprintf("%v", merged != nil),
			"errorsAsOk": fmt.Sprintf("%v", errors.As(merged, &pathErr)),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// errors.As verification for ConcatMessageWithErr
// =============================================================================

func Test_ConcatMessageWithErr_ErrorsAs_Verification(t *testing.T) {
	for caseIndex, testCase := range concatMessageErrorsAsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		err := input["error"].(error)
		msg, _ := input.GetAsString("message")

		// Act
		wrapped := errcore.ConcatMessageWithErr(msg, err)

		// Assert
		var pathErr *os.PathError
		actual := args.Map{
			"hasError":   fmt.Sprintf("%v", wrapped != nil),
			"errorsAsOk": fmt.Sprintf("%v", errors.As(wrapped, &pathErr)),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ConcatMessageWithErr nil passthrough
// =============================================================================

func Test_ConcatMessageWithErr_NilPassthrough_Verification(t *testing.T) {
	for caseIndex, testCase := range concatMessageNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		msg, _ := input.GetAsString("message")

		// Act
		result := errcore.ConcatMessageWithErr(msg, nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result == nil))
	}
}

// =============================================================================
// MergeErrors preserves multiple sentinels
// =============================================================================

func Test_MergeErrors_MultipleSentinels_Verification(t *testing.T) {
	for caseIndex, testCase := range mergeErrorsMultiSentinelTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errs := input["errors"].([]error)
		sentinels := input["sentinels"].([]error)

		// Act
		merged := errcore.MergeErrors(errs...)

		// Assert
		allMatch := true
		for _, s := range sentinels {
			if !errors.Is(merged, s) {
				allMatch = false
				break
			}
		}

		actual := args.Map{
			"hasError":       fmt.Sprintf("%v", merged != nil),
			"allSentinelsOk": fmt.Sprintf("%v", allMatch),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
