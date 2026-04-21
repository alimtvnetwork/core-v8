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

package corevalidatortests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/corevalidator"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func Test_TestValidators_Verification(t *testing.T) {
	for caseIndex, testCase := range textValidatorsTestCases {
		// Arrange
		parameter := corevalidator.Parameter{
			CaseIndex:                  constants.Zero,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: testCase.IsSkipOnContentsEmpty,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            testCase.IsCaseSensitive,
		}

		err := testCase.Validators.AllVerifyErrorMany(
			&parameter,
			testCase.ComparingLines...,
		)

		errorLines := errcore.ErrorToSplitLines(err)

		sliceValidator := corevalidator.SliceValidator{
			Condition:     corevalidator.DefaultDisabledCoreCondition,
			CompareAs:     stringcompareas.Equal,
			ActualLines:   errorLines,
			ExpectedLines: testCase.ExpectationLines,
		}

		nextBaseParam := corevalidator.Parameter{
			CaseIndex:                  caseIndex,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            testCase.IsCaseSensitive,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(&nextBaseParam)
		isValid := validationFinalError == nil

		// Assert
		actLines := []string{fmt.Sprintf("%v", isValid)}
		expected := []string{"true"}

		testCase.ShouldBeEqual(t, caseIndex, actLines, expected)
	}
}
