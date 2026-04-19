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

package corerangestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_StartEndRanges_ValidCases(t *testing.T) {
	for _, testCase := range validStartEndRangesTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.StartEndInt)
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				convey.So(
					actualRanges,
					should.Equal,
					testCase.ExpectedInput,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}

func Test_StartEndString_Functions_Result_Verification(t *testing.T) {
	for caseIndex, testCase := range startEndRangesStringFunctionsVerificationTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.StartEndInt)
		sliceValidator := testCase.Validator
		slice := corestr.New.SimpleSlice.Cap(100)

		for i, input := range arrangeInputs {
			slice.AppendFmt("StartEnd : %s", input.String())

			slice.AppendFmt(
				"    [%d] func : %s        | result : %s",
				i,
				"String",
				input.String(),
			)

			slice.AppendFmt(
				"    [%d] func : %s   | result : %s",
				i,
				"StringColon",
				input.StringColon(),
			)

			slice.AppendFmt(
				"    [%d] func : %s  | result : %s",
				i,
				"StringHyphen",
				input.StringHyphen(),
			)

			slice.AppendFmt(
				"    [%d] func : %s   | result : %s",
				i,
				"StringSpace",
				input.StringSpace(),
			)
		}

		// Act
		actual := slice.Strings()
		testCase.SetActual(actual)
		sliceValidator.SetActual(actual)

		nextBaseParam := corevalidator.Parameter{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam,
		)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				errcore.PrintErrorWithTestIndex(
					caseIndex,
					testCase.Title,
					validationFinalError,
				)

				convey.So(
					validationFinalError,
					should.BeNil,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}
