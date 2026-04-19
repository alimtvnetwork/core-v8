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
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// GroupBy — group by first character
// ==========================================

var collectionGroupByTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupBy returns grouped map -- 5 strings by first character",
		ArrangeInput: args.Map{
			"items": []string{"apple", "avocado", "banana", "blueberry", "cherry"},
		},
		ExpectedInput: args.Map{
			"groupA": "a:2",
			"groupB": "b:2",
			"groupC": "c:1",
		},
	},
	{
		Title: "GroupBy returns empty map -- empty input",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: args.Map{},
	},
	{
		Title: "GroupBy returns single group -- all items start with 'a'",
		ArrangeInput: args.Map{
			"items": []string{"ant", "ape", "ace"},
		},
		ExpectedInput: "a:3",
	},
}

// ==========================================
// GroupByCount
// ==========================================

var collectionGroupByCountTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupByCount returns occurrence counts -- [red,blue,red,green,blue,red]",
		ArrangeInput: args.Map{
			"items": []string{"red", "blue", "red", "green", "blue", "red"},
		},
		ExpectedInput: args.Map{
			"blueCount":  "blue:2",
			"greenCount": "green:1",
			"redCount":   "red:3",
		},
	},
	{
		Title: "GroupByCount returns empty -- empty input",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: "0",
	},
}
