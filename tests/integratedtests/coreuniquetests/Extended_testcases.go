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

package coreuniquetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// GetMap
// ==========================================================================

var extGetMapWithDuplicatesTestCase = coretestcases.CaseV1{
	Title: "GetMap returns unique map -- slice with duplicates [1,2,2,3]",
	ArrangeInput: args.Map{
		"when":  "given slice with duplicates",
		"input": []int{1, 2, 2, 3},
	},
	ExpectedInput: args.Map{
		"isNil":  false,
		"length": 3,
	},
}

var extGetMapNilTestCase = coretestcases.CaseV1{
	Title: "GetMap returns nil -- nil slice input",
	ArrangeInput: args.Map{
		"when": "given nil slice",
	},
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var extGetMapEmptyTestCase = coretestcases.CaseV1{
	Title: "GetMap returns empty map -- empty slice input",
	ArrangeInput: args.Map{
		"when":  "given empty slice",
		"input": []int{},
	},
	ExpectedInput: args.Map{
		"isNil":  false,
		"length": 0,
	},
}

// ==========================================================================
// Get with empty slice
// ==========================================================================

var extGetEmptySliceTestCase = coretestcases.CaseV1{
	Title: "Get returns same slice -- empty slice input",
	ArrangeInput: args.Map{
		"when":  "given empty slice",
		"input": []int{},
	},
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// ==========================================================================
// Get with single element
// ==========================================================================

var extGetSingleElementTestCase = coretestcases.CaseV1{
	Title: "Get returns same slice -- single element slice",
	ArrangeInput: args.Map{
		"when":  "given single element slice",
		"input": []int{42},
	},
	ExpectedInput: args.Map{
		"length": 1,
	},
}
