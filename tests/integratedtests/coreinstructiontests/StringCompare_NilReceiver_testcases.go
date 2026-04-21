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

package coreinstructiontests

import (
	"github.com/alimtvnetwork/core-v8/coreinstruction"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/coretests/results"
)

// =============================================================================
// StringCompare nil receiver test cases (migrated from CaseV1 string-dispatch)
// =============================================================================

var stringCompareNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Nil receiver - IsMatch returns true (vacuous truth)",
		Func:  (*coreinstruction.StringCompare).IsMatch,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - IsMatchFailed returns false",
		Func:  (*coreinstruction.StringCompare).IsMatchFailed,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - IsInvalid returns true",
		Func:  (*coreinstruction.StringCompare).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - IsDefined returns false",
		Func:  (*coreinstruction.StringCompare).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver - VerifyError returns nil",
		Func:  (*coreinstruction.StringCompare).VerifyError,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
}
