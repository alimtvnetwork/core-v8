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

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/internal/reflectinternal"
)

func Test_TwoParams_Method_Verification(t *testing.T) {
	for caseIndex, testCase := range versionTwoParamsVerificationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.FourAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		actualSlice.AppendFmt(
			"Testing for -> Version(%s)",
			someVersionV5,
		)

		// Act
		for index, input := range inputs {
			// "v5.8.1.5"
			f := input.First.(int)
			s := input.Second.(int)
			third := input.Third.(bool)
			theFunc := input.Fourth.(func(major, x int) bool)
			funcName := reflectinternal.GetFunc.NameOnly(theFunc)

			isMatch := theFunc(f, s)

			actualSlice.AppendFmt(
				comparisonMethodFmt,
				index,
				funcName,
				f,
				s,
				isMatch,
				third,
			)
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_ThreeParams_Method_Verification(t *testing.T) {
	for caseIndex, testCase := range versionThreeParamsVerificationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.FiveAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		actualSlice.AppendFmt(
			"Testing for -> Version(%s)",
			someVersionV5,
		)

		// Act
		for index, input := range inputs {
			// "v5.8.1.5"
			f := input.First.(int)
			s := input.Second.(int)
			third := input.Third.(int)
			fourth := input.Fourth.(bool)
			theFunc := input.Fifth.(func(major, x, y int) bool)
			funcName := reflectinternal.GetFunc.NameOnly(theFunc)

			isMatch := theFunc(f, s, third)

			actualSlice.AppendFmt(
				comparisonMethod3Fmt,
				index,
				funcName,
				f,
				s,
				third,
				isMatch,
				fourth,
			)
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
