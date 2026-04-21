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

package coredatatests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var pointerStringsDscLenTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStringsDsc nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "PointerStringsDsc with elements Len returns count",
		ArrangeInput: args.Map{
			"count": 3,
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var pointerStringsDscLessTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStringsDsc Less both non-nil descending",
		ExpectedInput: args.Map{
			"lessAB": false,
			"lessBA": true,
		},
	},
	{
		Title: "PointerStringsDsc Less nil-i returns false",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "PointerStringsDsc Less nil-j returns true",
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "PointerStringsDsc Less both nil returns false",
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var pointerStringsDscSortTestCases = []coretestcases.CaseV1{
	{
		Title: "PointerStringsDsc sort.Sort places values descending, nil last",
		ExpectedInput: args.Map{
			"first":     "charlie",
			"lastIsNil": true,
		},
	},
}
