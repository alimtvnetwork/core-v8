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

// ==========================================
// Contains
// ==========================================

var collectionContainsTestCases = []coretestcases.CaseV1{
	{
		Title: "Contains returns true for existing item",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "c"},
			"search": "b",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Contains returns false for missing item",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "c"},
			"search": "z",
		},
		ExpectedInput: "false",
	},
	{
		Title: "Contains returns false for empty collection",
		ArrangeInput: args.Map{
			"items":  []string{},
			"search": "a",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// IndexOf
// ==========================================

var collectionIndexOfTestCases = []coretestcases.CaseV1{
	{
		Title: "IndexOf returns correct index",
		ArrangeInput: args.Map{
			"items":  []string{"x", "y", "z"},
			"search": "y",
		},
		ExpectedInput: "1",
	},
	{
		Title: "IndexOf returns -1 for missing item",
		ArrangeInput: args.Map{
			"items":  []string{"x", "y", "z"},
			"search": "w",
		},
		ExpectedInput: "-1",
	},
}

// ==========================================
// HasAll
// ==========================================

var collectionHasAllTestCases = []coretestcases.CaseV1{
	{
		Title: "HasAll returns true when all present",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "c", "d"},
			"search": []string{"b", "d"},
		},
		ExpectedInput: "true",
	},
	{
		Title: "HasAll returns false when one missing",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b"},
			"search": []string{"a", "z"},
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// LastIndexOf
// ==========================================

var collectionLastIndexOfTestCases = []coretestcases.CaseV1{
	{
		Title: "LastIndexOf returns last occurrence",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "a", "c"},
			"search": "a",
		},
		ExpectedInput: "2",
	},
}

// ==========================================
// Count
// ==========================================

var collectionCountTestCases = []coretestcases.CaseV1{
	{
		Title: "Count returns correct occurrence count",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "a", "a", "c"},
			"search": "a",
		},
		ExpectedInput: "3",
	},
	{
		Title: "Count returns 0 for missing item",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b"},
			"search": "z",
		},
		ExpectedInput: "0",
	},
}
