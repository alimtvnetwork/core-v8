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
// AddLock
// ==========================================

var collectionAddLockTestCases = []coretestcases.CaseV1{
	{
		Title: "AddLock appends item thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent AddLock calls",
			"count": 100,
		},
		ExpectedInput: "100",
	},
}

// ==========================================
// AddsLock
// ==========================================

var collectionAddsLockTestCases = []coretestcases.CaseV1{
	{
		Title: "AddsLock appends multiple items thread-safely",
		ArrangeInput: args.Map{
			"when":  "given concurrent AddsLock calls",
			"count": 50,
			"batch": 2,
		},
		ExpectedInput: "100",
	},
}

// ==========================================
// LengthLock
// ==========================================

var collectionLengthLockTestCases = []coretestcases.CaseV1{
	{
		Title: "LengthLock returns correct count under concurrency",
		ArrangeInput: args.Map{
			"when":  "given items added then LengthLock called",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: "3",
	},
}

// ==========================================
// IsEmptyLock
// ==========================================

var collectionIsEmptyLockTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyLock returns true for empty collection",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: "true",
	},
}

var collectionIsEmptyLockNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyLock returns false for non-empty collection",
		ArrangeInput: args.Map{
			"when": "given non-empty collection",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// ItemsLock
// ==========================================

var collectionItemsLockTestCases = []coretestcases.CaseV1{
	{
		Title: "ItemsLock returns independent copy",
		ArrangeInput: args.Map{
			"when":  "given collection with items",
			"items": []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"length":        2,
			"first":         "x",
			"last":          "y",
			"isIndependent": true,
		},
	},
}

// ==========================================
// ClearLock
// ==========================================

var collectionClearLockTestCases = []coretestcases.CaseV1{
	{
		Title: "ClearLock removes all items thread-safely",
		ArrangeInput: args.Map{
			"when":  "given collection then ClearLock",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================
// AddCollectionLock
// ==========================================

var collectionAddCollectionLockTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCollectionLock merges thread-safely",
		ArrangeInput: args.Map{
			"when":   "given two collections merged with lock",
			"first":  []string{"a"},
			"second": []string{"b", "c"},
		},
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "a",
			"last":   "c",
		},
	},
}
