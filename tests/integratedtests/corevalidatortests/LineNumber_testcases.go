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
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var lineNumberHasLineNumberTestCases = []coretestcases.CaseV1{
	{
		Title: "LineNumber 5 HasLineNumber returns true",
		ArrangeInput: args.Map{
			"lineNumber": 5,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber 0 HasLineNumber returns true",
		ArrangeInput: args.Map{
			"lineNumber": 0,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber -1 HasLineNumber returns false",
		ArrangeInput: args.Map{
			"lineNumber": -1,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "LineNumber -2 HasLineNumber returns false",
		ArrangeInput: args.Map{
			"lineNumber": -2,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var lineNumberIsMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "LineNumber 3 IsMatch 3 returns true",
		ArrangeInput: args.Map{
			"lineNumber": 3,
			"input":      3,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber 3 IsMatch 5 returns false",
		ArrangeInput: args.Map{
			"lineNumber": 3,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "LineNumber 3 IsMatch -1 returns true (skip check)",
		ArrangeInput: args.Map{
			"lineNumber": 3,
			"input":      -1,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber -1 IsMatch 5 returns true (skip check)",
		ArrangeInput: args.Map{
			"lineNumber": -1,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber -1 IsMatch -1 returns true",
		ArrangeInput: args.Map{
			"lineNumber": -1,
			"input":      -1,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "LineNumber 0 IsMatch 0 returns true",
		ArrangeInput: args.Map{
			"lineNumber": 0,
			"input":      0,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var lineNumberVerifyErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "LineNumber 2 VerifyError 2 returns nil",
		ArrangeInput: args.Map{
			"lineNumber": 2,
			"input":      2,
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "LineNumber 2 VerifyError 5 returns error",
		ArrangeInput: args.Map{
			"lineNumber": 2,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "LineNumber -1 VerifyError 5 returns nil (skip)",
		ArrangeInput: args.Map{
			"lineNumber": -1,
			"input":      5,
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}
