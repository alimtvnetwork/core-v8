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

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/corevalidator"
)

// ==========================================
// Length
// ==========================================

func Test_HeaderSliceValidators_Length(t *testing.T) {
	for caseIndex, tc := range headerSliceValidatorsLengthTestCases {
		// Arrange
		input := tc.ArrangeInput
		var v corevalidator.HeaderSliceValidators
		if input != nil {
			v = input.(corevalidator.HeaderSliceValidators)
		}

		// Act
		actual := args.Map{
			"length": v.Length(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// IsEmpty
// ==========================================

func Test_HeaderSliceValidators_IsEmpty(t *testing.T) {
	for caseIndex, tc := range headerSliceValidatorsIsEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput
		var v corevalidator.HeaderSliceValidators
		if input != nil {
			v = input.(corevalidator.HeaderSliceValidators)
		}

		// Act
		actual := args.Map{
			"isEmpty": v.IsEmpty(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// IsMatch
// ==========================================

func Test_HeaderSliceValidators_IsMatch(t *testing.T) {
	for caseIndex, tc := range headerSliceValidatorsIsMatchTestCases {
		// Arrange
		input := tc.ArrangeInput
		var v corevalidator.HeaderSliceValidators
		if input != nil {
			v = input.(corevalidator.HeaderSliceValidators)
		}

		// Act
		actual := args.Map{
			"isMatch": v.IsMatch(true),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// IsValid (delegates to IsMatch)
// ==========================================

func Test_HeaderSliceValidators_IsValid_FromHeaderSliceValidator(t *testing.T) {
	for caseIndex, tc := range headerSliceValidatorsIsMatchTestCases {
		// Arrange
		input := tc.ArrangeInput
		var v corevalidator.HeaderSliceValidators
		if input != nil {
			v = input.(corevalidator.HeaderSliceValidators)
		}

		// Act
		actual := args.Map{
			"isMatch": v.IsValid(true),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// SetActualOnAll
// ==========================================

func Test_HeaderSliceValidators_SetActualOnAll_Empty_FromHeaderSliceValidator(t *testing.T) {
	// should not panic
	var v corevalidator.HeaderSliceValidators
	v.SetActualOnAll("a", "b")
}

func Test_HeaderSliceValidators_SetActualOnAll_NonEmpty(t *testing.T) {
	v := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				ActualLines:   []string{"old"},
				ExpectedLines: []string{"new"},
			},
		},
	}

	// Act — should not panic
	v.SetActualOnAll("new")
}

// ==========================================
// VerifyAll
// ==========================================

func Test_HeaderSliceValidators_VerifyAll(t *testing.T) {
	for caseIndex, tc := range headerSliceValidatorsVerifyAllTestCases {
		// Arrange
		v := tc.ArrangeInput.(corevalidator.HeaderSliceValidators)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyAll("header", params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyFirst
// ==========================================

func Test_HeaderSliceValidators_VerifyFirst(t *testing.T) {
	for caseIndex, tc := range headerSliceValidatorsVerifyFirstTestCases {
		// Arrange
		v := tc.ArrangeInput.(corevalidator.HeaderSliceValidators)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyFirst(params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyUpto
// ==========================================

func Test_HeaderSliceValidators_VerifyUpto(t *testing.T) {
	for caseIndex, tc := range headerSliceValidatorsVerifyUptoTestCases {
		// Arrange
		v := tc.ArrangeInput.(corevalidator.HeaderSliceValidators)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyUpto(false, false, 1, params)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
