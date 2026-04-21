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
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/coretests/results"
	"github.com/alimtvnetwork/core-v8/corevalidator"
)

// =============================================================================
// SliceValidator nil receiver test cases
// (migrated from inline t.Error tests in SliceValidatorUnit_test.go
//  and SliceValidatorExtra_test.go)
// =============================================================================

var sliceValidatorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsValid on nil returns true",
		Func:  (*corevalidator.SliceValidator).IsValid,
		Args:  []any{true},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ActualLinesLength on nil returns 0",
		Func:  (*corevalidator.SliceValidator).ActualLinesLength,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "AllVerifyError on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyError,
		Args:  []any{&corevalidator.Parameter{CaseIndex: 0}},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "VerifyFirstError on nil returns nil",
		Func:  (*corevalidator.SliceValidator).VerifyFirstError,
		Args:  []any{&corevalidator.Parameter{CaseIndex: 0}},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "Dispose on nil does not panic",
		Func:  (*corevalidator.SliceValidator).Dispose,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "AllVerifyErrorExceptLast on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyErrorExceptLast,
		Args:  []any{&corevalidator.Parameter{CaseIndex: 0}},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "AllVerifyErrorQuick on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyErrorQuick,
		Args:  []any{0, "test", "a"},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "AllVerifyErrorTestCase on nil returns nil",
		Func:  (*corevalidator.SliceValidator).AllVerifyErrorTestCase,
		Args:  []any{0, "test", true},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
	{
		Title: "ActualLinesString on nil returns empty",
		Func:  (*corevalidator.SliceValidator).ActualLinesString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "ExpectingLinesString on nil returns empty",
		Func:  (*corevalidator.SliceValidator).ExpectingLinesString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "IsUsedAlready on nil returns false",
		Func:  (*corevalidator.SliceValidator).IsUsedAlready,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
