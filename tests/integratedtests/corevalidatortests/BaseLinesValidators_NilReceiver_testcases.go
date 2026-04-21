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
// BaseLinesValidators nil receiver test cases
// (migrated from inline t.Error tests in BaseLinesValidators_test.go)
// =============================================================================

var baseLinesValidatorsNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "LinesValidatorsLength on nil returns 0",
		Func:  (*corevalidator.BaseLinesValidators).LinesValidatorsLength,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmptyLinesValidators on nil returns true",
		Func:  (*corevalidator.BaseLinesValidators).IsEmptyLinesValidators,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasLinesValidators on nil returns false",
		Func:  (*corevalidator.BaseLinesValidators).HasLinesValidators,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}

// =============================================================================
// LinesValidators nil receiver test cases
// (migrated from inline t.Error tests in BaseLinesValidators_test.go)
// =============================================================================

var linesValidatorsNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Length on nil returns 0",
		Func:  (*corevalidator.LinesValidators).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*corevalidator.LinesValidators).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
