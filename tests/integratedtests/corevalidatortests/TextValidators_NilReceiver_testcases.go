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
// TextValidators nil receiver test cases
// (migrated from inline t.Error tests in TextValidators_test.go)
// =============================================================================

var textValidatorsNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Dispose on nil does not panic",
		Func:  (*corevalidator.TextValidators).Dispose,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*corevalidator.TextValidators).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
	},
	{
		Title: "VerifyErrorMany on nil returns nil",
		Func:  (*corevalidator.TextValidators).VerifyErrorMany,
		Args:  []any{true, &corevalidator.Parameter{CaseIndex: 0}, "a"},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
}
