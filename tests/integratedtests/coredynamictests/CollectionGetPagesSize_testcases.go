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

package coredynamictests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var collectionGetPagesSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagesSize returns 0 -- 5 items, eachPageSize=0",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=0",
			"items":        []int{1, 2, 3, 4, 5},
			"eachPageSize": 0,
		},
		ExpectedInput: "0",
	},
	{
		Title: "GetPagesSize returns 0 -- 3 items, eachPageSize=-3",
		ArrangeInput: args.Map{
			"when":         "given eachPageSize=-3",
			"items":        []int{1, 2, 3},
			"eachPageSize": -3,
		},
		ExpectedInput: "0",
	},
	{
		Title: "GetPagesSize returns 0 -- empty collection, eachPageSize=5",
		ArrangeInput: args.Map{
			"when":         "given empty collection with valid page size",
			"items":        []int{},
			"eachPageSize": 5,
		},
		ExpectedInput: "0",
	},
	{
		Title: "GetPagesSize returns 1 -- 3 items, eachPageSize=5",
		ArrangeInput: args.Map{
			"when":         "given 3 items with eachPageSize=5",
			"items":        []int{1, 2, 3},
			"eachPageSize": 5,
		},
		ExpectedInput: "1",
	},
	{
		Title: "GetPagesSize returns 1 -- 5 items exactly fill eachPageSize=5",
		ArrangeInput: args.Map{
			"when":         "given 5 items with eachPageSize=5",
			"items":        []int{1, 2, 3, 4, 5},
			"eachPageSize": 5,
		},
		ExpectedInput: "1",
	},
	{
		Title: "GetPagesSize returns 2 -- 6 items, eachPageSize=5",
		ArrangeInput: args.Map{
			"when":         "given 6 items with eachPageSize=5",
			"items":        []int{1, 2, 3, 4, 5, 6},
			"eachPageSize": 5,
		},
		ExpectedInput: "2",
	},
	{
		Title: "GetPagesSize returns 3 -- 10 items, eachPageSize=4",
		ArrangeInput: args.Map{
			"when":         "given 10 items with eachPageSize=4",
			"items":        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			"eachPageSize": 4,
		},
		ExpectedInput: "3",
	},
	{
		Title: "GetPagesSize returns 4 -- 4 items, eachPageSize=1",
		ArrangeInput: args.Map{
			"when":         "given 4 items with eachPageSize=1",
			"items":        []int{1, 2, 3, 4},
			"eachPageSize": 1,
		},
		ExpectedInput: "4",
	},
	{
		Title: "GetPagesSize returns 1 -- single item, eachPageSize=1",
		ArrangeInput: args.Map{
			"when":         "given 1 item with eachPageSize=1",
			"items":        []int{42},
			"eachPageSize": 1,
		},
		ExpectedInput: "1",
	},
}
