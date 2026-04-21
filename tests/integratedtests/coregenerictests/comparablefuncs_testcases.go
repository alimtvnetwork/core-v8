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

// ==========================================
// ContainsAll
// ==========================================

var containsAllTrueTestCase = coretestcases.CaseV1{
	Title: "ContainsAll returns true -- all search items [1,3,5] present in [1,2,3,4,5]",
	ArrangeInput: args.Map{
		"when":        "given collection containing all search items",
		"items":       []int{1, 2, 3, 4, 5},
		"searchItems": []int{1, 3, 5},
	},
	ExpectedInput: "true",
}

var containsAllFalseTestCase = coretestcases.CaseV1{
	Title: "ContainsAll returns false -- search item 99 missing from [1,2,3]",
	ArrangeInput: args.Map{
		"when":        "given collection missing one search item",
		"items":       []int{1, 2, 3},
		"searchItems": []int{1, 2, 99},
	},
	ExpectedInput: "false",
}

// ==========================================
// ContainsAny
// ==========================================

var containsAnyTrueTestCase = coretestcases.CaseV1{
	Title: "ContainsAny returns true -- search item 3 found in [1,2,3]",
	ArrangeInput: args.Map{
		"when":        "given collection with one matching item",
		"items":       []int{1, 2, 3},
		"searchItems": []int{99, 3, 100},
	},
	ExpectedInput: "true",
}

var containsAnyFalseTestCase = coretestcases.CaseV1{
	Title: "ContainsAny returns false -- no search items [88,99,100] in [1,2,3]",
	ArrangeInput: args.Map{
		"when":        "given collection with no matching items",
		"items":       []int{1, 2, 3},
		"searchItems": []int{88, 99, 100},
	},
	ExpectedInput: "false",
}

// ==========================================
// RemoveItem
// ==========================================

var removeItemFoundTestCase = coretestcases.CaseV1{
	Title: "RemoveItem returns true -- removing 2 from [1,2,3,2,4]",
	ArrangeInput: args.Map{
		"when":       "given collection with duplicates, remove first 2",
		"items":      []int{1, 2, 3, 2, 4},
		"removeItem": 2,
	},
	ExpectedInput: args.Map{
		"removed":   true,
		"newLength": 4,
	},
}

var removeItemMissingTestCase = coretestcases.CaseV1{
	Title: "RemoveItem returns false -- removing 99 from [1,3,5]",
	ArrangeInput: args.Map{
		"when":       "given collection without target item",
		"items":      []int{1, 3, 5},
		"removeItem": 99,
	},
	ExpectedInput: args.Map{
		"removed":   false,
		"newLength": 3,
	},
}

// ==========================================
// RemoveAllItems
// ==========================================

var removeAllItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "RemoveAllItems returns 3 removed -- removing all 2s from [1,2,3,2,4,2]",
		ArrangeInput: args.Map{
			"when":  "given collection with multiple 2s",
			"items": []int{1, 2, 3, 2, 4, 2},
		},
		ExpectedInput: args.Map{
			"removedCount": 3,
			"newLength":    3,
		},
	},
}

// ==========================================
// ToHashset
// ==========================================

var toHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "ToHashset returns 3 unique items -- from [1,2,3,2,1]",
		ArrangeInput: args.Map{
			"when":  "given collection with duplicates",
			"items": []int{1, 2, 3, 2, 1},
		},
		ExpectedInput: args.Map{
			"uniqueCount": 3,
			"has1":        true,
			"has2":        true,
			"has3":        true,
			"has99":       false,
		},
	},
}

// ==========================================
// DistinctSimpleSlice
// ==========================================

var distinctSimpleSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "DistinctSimpleSlice returns 4 unique preserving order -- from [3,1,2,1,3,4]",
		ArrangeInput: args.Map{
			"when":  "given simple slice with duplicates",
			"items": []int{3, 1, 2, 1, 3, 4},
		},
		ExpectedInput: args.Map{
			"uniqueCount":  4,
			"firstElement": 3,
			"lastElement":  4,
		},
	},
}

// ==========================================
// ContainsSimpleSliceItem
// ==========================================

var containsSimpleSliceItemTestCases = []coretestcases.CaseV1{
	{
		Title: "ContainsSimpleSliceItem returns true/false -- existing and missing items in [10,20,30]",
		ArrangeInput: args.Map{
			"when":  "given simple slice containing target",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: args.Map{
			"containsExisting": true,
			"containsMissing":  false,
		},
	},
}
