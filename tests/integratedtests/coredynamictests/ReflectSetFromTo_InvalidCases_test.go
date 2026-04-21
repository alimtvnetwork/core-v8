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

	"github.com/alimtvnetwork/core-v8/converters"
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/corevalidator"
	"github.com/alimtvnetwork/core-v8/internal/trydo"
	"github.com/alimtvnetwork/core-v8/tests/testwrappers/coredynamictestwrappers"
)

// Test_ReflectSetFromTo_ValidCases
//
// Valid Inputs:
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
func Test_ReflectSetFromTo_InvalidCases_Verification(t *testing.T) {
	for caseIndex, testCase := range coredynamictestwrappers.ReflectSetFromToInvalidTestCases {
		// Act
		wrappedResult := trydo.ErrorFuncWrapPanic(
			func() error {
				return coredynamic.ReflectSetFromTo(
					testCase.From,
					testCase.To,
				)
			},
		)

		err := wrappedResult.Error
		testCase.SetActual(wrappedResult)

		// Assert - error expectation
		hasErr := fmt.Sprintf("%v", err != nil)
		expectedHasErr := fmt.Sprintf("%v", testCase.IsExpectingError)

		actLines := []string{hasErr}
		expected := []string{expectedHasErr}

		// Assert - validator verification
		parameter := &corevalidator.Parameter{
			CaseIndex:                  caseIndex,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            true,
		}

		finalErr := getFinalVerificationError(
			testCase,
			testCase.Validator,
			parameter,
			wrappedResult,
		)

		actLines = append(actLines, fmt.Sprintf("%v", finalErr == nil))
		expected = append(expected, "true")

		testCase.ShouldBeEqual(t, caseIndex, actLines, expected)
	}
}

func getFinalVerificationError(
	testCase coredynamictestwrappers.FromToTestWrapper,
	validator corevalidator.TextValidator,
	parameter *corevalidator.Parameter,
	wrappedResult trydo.WrappedErr,
) error {
	if testCase.HasPanic {
		return validator.VerifyDetailError(
			parameter,
			wrappedResult.ExceptionString(),
		)
	}

	if testCase.IsExpectingError {
		return validator.VerifyDetailError(
			parameter,
			wrappedResult.ErrorString(),
		)
	}

	return validator.VerifyDetailError(
		parameter,
		converters.AnyTo.ValueString(testCase.ExpectedValue),
	)
}
