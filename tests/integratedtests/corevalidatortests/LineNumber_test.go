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

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
)

func Test_LineNumber_HasLineNumber(t *testing.T) {
	for caseIndex, tc := range lineNumberHasLineNumberTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		lineNum, _ := input.GetAsInt("lineNumber")
		ln := corevalidator.LineNumber{LineNumber: lineNum}

		// Act
		actual := args.Map{
			"result": ln.HasLineNumber(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LineNumber_IsMatch(t *testing.T) {
	for caseIndex, tc := range lineNumberIsMatchTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		lineNum, _ := input.GetAsInt("lineNumber")
		inputNum, _ := input.GetAsInt("input")
		ln := corevalidator.LineNumber{LineNumber: lineNum}

		// Act
		actual := args.Map{
			"result": ln.IsMatch(inputNum),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LineNumber_VerifyError(t *testing.T) {
	for caseIndex, tc := range lineNumberVerifyErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		lineNum, _ := input.GetAsInt("lineNumber")
		inputNum, _ := input.GetAsInt("input")
		ln := corevalidator.LineNumber{LineNumber: lineNum}

		// Act
		err := ln.VerifyError(inputNum)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
