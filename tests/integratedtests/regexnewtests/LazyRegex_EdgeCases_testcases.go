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

package regexnewtests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// Empty pattern edge cases
// =============================================================================

var emptyPatternEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex returns undefined -- empty pattern",
		ExpectedInput: args.Map{
			params.isUndefined: true,
		},
	},
	{
		Title: "LazyRegex returns not applicable -- empty pattern",
		ExpectedInput: args.Map{
			params.isApplicable: false,
		},
	},
	{
		Title: "IsMatch returns false -- empty pattern",
		ExpectedInput: args.Map{
			params.isMatch: false,
		},
	},
	{
		Title: "IsFailedMatch returns true -- empty pattern",
		ExpectedInput: args.Map{
			params.isFailedMatch: true,
		},
	},
	{
		Title: "Compile returns error and nil regex -- empty pattern",
		ExpectedInput: args.Map{
			params.hasError:   true,
			params.regexIsNil: true,
		},
	},
}
