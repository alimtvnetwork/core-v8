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
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/isany"
)

func Test_SliceValidator(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorTestCases {
		// Arrange
		inputs := testCase.
			Case.
			ArrangeInput.([]args.TwoAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				"%d : %t (%s, %s)",
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
			)
		}

		actLines := actualSlice.Strings()
		actualError := testCase.Case.VerifyAllEqual(
			caseIndex,
			actLines...,
		)
		validator := testCase.Validator
		errLines := errcore.ErrorToSplitLines(actualError)

		// Assert
		validator.AssertAllQuick(
			t,
			caseIndex,
			testCase.Case.Title,
			errLines...,
		)
	}
}

func Test_SliceValidator_FirstError(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorFirstErrorTestCases {
		// Arrange
		inputs := testCase.
			Case.
			ArrangeInput.([]args.TwoAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				"%d : %t (%s, %s)",
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
			)
		}

		actLines := actualSlice.Strings()
		actualError := testCase.Case.VerifyFirst(
			caseIndex,
			stringcompareas.Equal,
			actLines,
		)
		validator := testCase.Validator
		errLines := errcore.ErrorToSplitLines(actualError)

		// Assert
		validator.AssertAllQuick(
			t,
			caseIndex,
			testCase.Case.Title,
			errLines...,
		)
	}
}
