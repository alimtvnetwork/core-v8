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
// LengthOfVerifierSegments
// ==========================================

func Test_RangeSegmentsValidator_LengthOfVerifierSegments_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorLengthTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)

		// Act
		actual := args.Map{
			"length": v.LengthOfVerifierSegments(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Validators
// ==========================================

func Test_RangeSegmentsValidator_Validators_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorValidatorsTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		v.SetActual(rangeSegActualLines)

		// Act
		validators := v.Validators()
		actual := args.Map{
			"hasValidators": len(validators) > 0,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyAll
// ==========================================

func Test_RangeSegmentsValidator_VerifyAll(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyAllTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyAll("header", rangeSegActualLines, params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifySimple
// ==========================================

func Test_RangeSegmentsValidator_VerifySimple_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifySimpleTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifySimple(rangeSegActualLines, params, false)
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

func Test_RangeSegmentsValidator_VerifyFirst_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyFirstTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyFirst("header", rangeSegActualLines, params, false)
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

func Test_RangeSegmentsValidator_VerifyUpto_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyUptoTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyUpto("header", rangeSegActualLines, params, 3, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyFirstDefault
// ==========================================

func Test_RangeSegmentsValidator_VerifyFirstDefault_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyFirstDefaultTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyFirstDefault(rangeSegActualLines, params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyUptoDefault
// ==========================================

func Test_RangeSegmentsValidator_VerifyUptoDefault_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyUptoDefaultTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyUptoDefault(rangeSegActualLines, params, 3, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// SetActual
// ==========================================

func Test_RangeSegmentsValidator_SetActual_FromRangeSegmentsValidat(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorSetActualTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)

		// Act
		result := v.SetActual(rangeSegActualLines)
		validators := v.Validators()
		actual := args.Map{
			"returnsSelf": result == v,
			"isMatch":     validators.IsMatch(true),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// SetActualOnAll (via Validators)
// ==========================================

func Test_RangeSegmentsValidator_SetActualOnAll(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorSetActualOnAllTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		v.SetActual(rangeSegActualLines)

		// Act
		validators := v.Validators()
		validators.SetActualOnAll(rangeSegActualLines...)
		actual := args.Map{
			"validatorCount": len(validators),
			"isMatch":        validators.IsMatch(true),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
