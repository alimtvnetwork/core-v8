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

package coreindexestests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var hasIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "HasIndex returns true -- index 3 in [1,3,5]",
		ArrangeInput: args.Map{
			"when":    "given matching index",
			"indexes": []int{1, 3, 5},
			"current": 3,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "HasIndex returns false -- index 4 not in [1,3,5]",
		ArrangeInput: args.Map{
			"when":    "given non-matching index",
			"indexes": []int{1, 3, 5},
			"current": 4,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var lastIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "LastIndex returns 4 -- length 5",
		ArrangeInput: args.Map{
			"when":   "given length 5",
			"length": 5,
		},
		ExpectedInput: args.Map{"result": 4},
	},
	{
		Title: "LastIndex returns 0 -- length 1",
		ArrangeInput: args.Map{
			"when":   "given length 1",
			"length": 1,
		},
		ExpectedInput: args.Map{"result": 0},
	},
}

var isWithinIndexRangeTestCases = []coretestcases.CaseV1{
	{
		Title: "IsWithinIndexRange returns true -- index 2 in length 5",
		ArrangeInput: args.Map{
			"when":   "given index within range",
			"index":  2,
			"length": 5,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IsWithinIndexRange returns false -- index 5 in length 5",
		ArrangeInput: args.Map{
			"when":   "given index beyond range",
			"index":  5,
			"length": 5,
		},
		ExpectedInput: args.Map{"result": false},
	},
}
