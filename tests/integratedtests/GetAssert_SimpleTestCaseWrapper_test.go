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

package integratedtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

func Test_SimpleTestCaseWrapper_String_Verification(t *testing.T) {
	for caseIndex, testCase := range stringTestCases {
		// Arrange
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(20)
		asserter := coretests.GetAssert.SimpleTestCaseWrapper
		actFunc := asserter.String
		caseV1 := testCase.ArrangeInput.(coretestcases.CaseV1)
		simplerWrapper := caseV1.AsSimpleTestCaseWrapper()

		// Act
		output := actFunc(
			caseIndex,
			simplerWrapper,
		)

		actualSlice.Adds(strings.Split(output, "\n")...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_SimpleTestCaseWrapper_Lines_Verification(t *testing.T) {
	for caseIndex, testCase := range linesTestCases {
		// Arrange
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(20)
		asserter := coretests.GetAssert.SimpleTestCaseWrapper
		caseV1 := testCase.ArrangeInput.(coretestcases.CaseV1)
		simplerWrapper := caseV1.AsSimpleTestCaseWrapper()
		prefixWithSpaceFunc := coretests.GetAssert.ToStringsWithSpace
		actFunc := asserter.Lines

		// Act
		arrange, expected := actFunc(
			simplerWrapper,
		)

		actualSlice.Add("Title: " + simplerWrapper.CaseTitle())
		actualSlice.Add("Arrange Lines:")
		actualSlice.Adds(prefixWithSpaceFunc(4, arrange)...)
		actualSlice.Add("Expected Lines:")
		actualSlice.Adds(prefixWithSpaceFunc(4, expected)...)

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
