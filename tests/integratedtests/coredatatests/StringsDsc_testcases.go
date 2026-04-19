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
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var stringsDscLenTestCases = []coretestcases.CaseV1{
	{
		Title: "StringsDsc nil slice Len returns 0",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "StringsDsc empty slice Len returns 0",
		ArrangeInput: args.Map{
			"values": []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "StringsDsc with elements Len returns count",
		ArrangeInput: args.Map{
			"values": []string{"charlie", "alpha", "beta"},
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var stringsDscSortTestCases = []coretestcases.CaseV1{
	{
		Title: "StringsDsc sorts descending",
		ArrangeInput: args.Map{
			"values": []string{"alpha", "charlie", "beta"},
		},
		ExpectedInput: args.Map{
			"first": "charlie",
			"last":  "alpha",
		},
	},
	{
		Title: "StringsDsc single element unchanged",
		ArrangeInput: args.Map{
			"values": []string{"only"},
		},
		ExpectedInput: args.Map{
			"first": "only",
			"last":  "only",
		},
	},
}
