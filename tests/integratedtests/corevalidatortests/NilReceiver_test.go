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

import "testing"

// ==========================================
// SliceValidator — CaseNilSafe pattern
// ==========================================

func Test_SliceValidator_NilReceiver(t *testing.T) {
	for caseIndex, tc := range sliceValidatorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// SliceValidators — CaseNilSafe pattern
// ==========================================

func Test_SliceValidators_NilReceiver(t *testing.T) {
	for caseIndex, tc := range sliceValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// TextValidator — CaseNilSafe pattern
// ==========================================

func Test_TextValidator_NilReceiver(t *testing.T) {
	for caseIndex, tc := range textValidatorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// TextValidators — CaseNilSafe pattern
// ==========================================

func Test_TextValidators_NilReceiver(t *testing.T) {
	for caseIndex, tc := range textValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// BaseLinesValidators — CaseNilSafe pattern
// ==========================================

func Test_BaseLinesValidators_NilReceiver(t *testing.T) {
	for caseIndex, tc := range baseLinesValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// LinesValidators — CaseNilSafe pattern
// ==========================================

func Test_LinesValidators_NilReceiver_FromNilReceiver(t *testing.T) {
	for caseIndex, tc := range linesValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// LineValidator — CaseNilSafe pattern
// ==========================================

func Test_LineValidator_NilReceiver(t *testing.T) {
	for caseIndex, tc := range lineValidatorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
