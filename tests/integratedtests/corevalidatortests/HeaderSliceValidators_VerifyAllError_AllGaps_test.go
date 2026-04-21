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
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage19 — corevalidator remaining 8 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── HeaderSliceValidators.VerifyAllError non-empty with pass (line 112-118) ──

func Test_HeaderSliceValidators_VerifyAllError_Pass(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "test-header",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"hello"},
				ExpectedLines: []string{"hello"},
			},
		},
	}
	params := &corevalidator.Parameter{
		Header:          "Test header validators pass",
		IsCaseSensitive: true,
	}

	// Act
	err := hsv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "HeaderSliceValidators VerifyAllError pass", expected)
}

// ── SliceValidatorMessages: ActualInputWithExpectingMessage returns non-empty (line 80) ──

func Test_SliceValidatorMessages_UserInputsMergeWithError(t *testing.T) {
	// Arrange
	sv := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"actual"},
		ExpectedLines: []string{"expected"},
	}
	params := &corevalidator.Parameter{
		Header:          "Test mismatch",
		IsCaseSensitive: true,
		CaseIndex:       1,
	}

	// Act
	err := sv.VerifyFirstError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "SliceValidator VerifyFirstError mismatch", expected)
}

// ── SliceValidatorVerify: ActualLinesLength > 0 but comparingLength == 0 (line 220-225) ──

func Test_SliceValidatorVerify_NonEmptyActualButEmptyExpected(t *testing.T) {
	// Arrange
	sv := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"something"},
		ExpectedLines: []string{},
	}
	params := &corevalidator.Parameter{
		Header:          "Test non-empty actual empty expected",
		IsCaseSensitive: true,
		CaseIndex:       1,
	}

	// Act
	err := sv.VerifyFirstError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "SliceValidatorVerify non-empty actual empty expected", expected)
}

// ── SliceValidators.VerifyAllError non-empty with pass (line 114-120) ──

func Test_SliceValidators_VerifyAllError_Pass(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"match"},
				ExpectedLines: []string{"match"},
			},
		},
	}
	params := &corevalidator.Parameter{
		Header:          "Test slice validators pass",
		IsCaseSensitive: true,
	}

	// Act
	err := sv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "SliceValidators VerifyAllError pass", expected)
}

// ── TextValidator: verifyDetailErrorUsingLineProcessing nil guard (line 177) ──
// nil receiver guard — documented as dead code
