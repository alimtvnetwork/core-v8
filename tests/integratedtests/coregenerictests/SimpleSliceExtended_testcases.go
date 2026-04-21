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

package coregenerictests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var simpleSliceAddIfTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddIf true adds item",
		ArrangeInput:  args.Map{
			"isAdd": true,
			"item": "x",
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "AddIf false skips item",
		ArrangeInput:  args.Map{
			"isAdd": false,
			"item": "x",
		},
		ExpectedInput: args.Map{"length": 0},
	},
}

var simpleSliceAddsIfTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddsIf true adds items",
		ArrangeInput:  args.Map{"isAdd": true},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "AddsIf false skips items",
		ArrangeInput:  args.Map{"isAdd": false},
		ExpectedInput: args.Map{"length": 0},
	},
}

var simpleSliceMethodsTestCases = []coretestcases.CaseV1{
	{
		Title:        "FirstOrDefault, LastOrDefault, Count, HasAnyItem, HasItems, HasIndex, Items, ForEach, CountFunc, String on populated slice",
		ArrangeInput: args.Map{"items": []string{"a", "b", "c"}},
		ExpectedInput: args.Map{
			"firstOrDefault": "a",
			"lastOrDefault":  "c",
			"count":          3,
			"hasAnyItem":     true,
			"hasItems":       true,
			"hasIndex1":      true,
			"hasIndex5":      false,
			"itemsLen":       3,
			"forEachCount":   3,
			"countFuncGt0":   true,
			"stringNotEmpty": true,
		},
	},
}

var simpleSliceEmptyMethodsTestCases = []coretestcases.CaseV1{
	{
		Title:        "FirstOrDefault, LastOrDefault on empty slice",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"firstOrDefault": "",
			"lastOrDefault":  "",
		},
	},
}

var simpleSliceAddSliceTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddSlice with items",
		ArrangeInput:  args.Map{"items": []string{"a", "b"}},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "AddSlice with empty",
		ArrangeInput:  args.Map{"items": []string{}},
		ExpectedInput: args.Map{"length": 0},
	},
}

var simpleSliceAddFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddFunc adds result of function",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{
			"length": 1,
			"first": "gen",
		},
	},
}

var simpleSliceInsertAtTestCases = []coretestcases.CaseV1{
	{
		Title:         "InsertAt valid index",
		ArrangeInput:  args.Map{
			"index": 1,
			"item": "inserted",
		},
		ExpectedInput: args.Map{
			"length": 4,
			"atIndex": "inserted",
		},
	},
	{
		Title:         "InsertAt negative index -- no change",
		ArrangeInput:  args.Map{
			"index": -1,
			"item": "x",
		},
		ExpectedInput: args.Map{"length": 3},
	},
}

var linkedListNodeExtendedTestCases = []coretestcases.CaseV1{
	{
		Title:        "EndOfChain, Clone, ListPtr, String on node chain",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"endLength":    3,
			"cloneHasNext": false,
			"listPtrLen":   3,
			"stringOk":     true,
		},
	},
}
