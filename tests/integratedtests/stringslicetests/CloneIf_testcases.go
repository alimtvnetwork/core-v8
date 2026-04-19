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

package stringslicetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// stringslice.CloneIf
// =============================================================================

var cloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneIf returns independent copy with length 2 -- isClone true with extra cap",
		ArrangeInput: args.Map{
			"when":          "given isClone true with extra cap",
			"input":         []string{"a", "b"},
			"isClone":       true,
			"additionalCap": 5,
		},
		ExpectedInput: args.Map{
			"resultLength":      "2",
			"item0":             "a",
			"item1":             "b",
			"isIndependentCopy": "true",
		},
	},
	{
		Title: "CloneIf returns original slice reference -- isClone false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []string{"x", "y"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Map{
			"resultLength":      "2",
			"item0":             "x",
			"item1":             "y",
			"isIndependentCopy": "false",
		},
	},
	{
		Title: "CloneIf returns empty non-independent -- nil input with isClone false",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone false",
			"isNil":         true,
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Map{
			"resultLength":      "0",
			"isIndependentCopy": "false",
		},
	},
	{
		Title: "CloneIf returns empty independent copy -- nil input with isClone true",
		ArrangeInput: args.Map{
			"when":          "given nil input with isClone true",
			"isNil":         true,
			"isClone":       true,
			"additionalCap": 3,
		},
		ExpectedInput: args.Map{
			"resultLength":      "0",
			"isIndependentCopy": "true",
		},
	},
}

// =============================================================================
// stringslice.AnyItemsCloneIf
// =============================================================================

var anyItemsCloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyItemsCloneIf returns independent copy with length 3 -- isClone true",
		ArrangeInput: args.Map{
			"when":          "given isClone true",
			"input":         []any{"a", 1, true},
			"isClone":       true,
			"additionalCap": 2,
		},
		ExpectedInput: args.Map{
			"resultLength":      "3",
			"item0":             "a",
			"item1":             "1",
			"item2":             "true",
			"isIndependentCopy": "true",
		},
	},
	{
		Title: "AnyItemsCloneIf returns original slice reference -- isClone false",
		ArrangeInput: args.Map{
			"when":          "given isClone false",
			"input":         []any{"x"},
			"isClone":       false,
			"additionalCap": 0,
		},
		ExpectedInput: args.Map{
			"resultLength":      "1",
			"item0":             "x",
			"isIndependentCopy": "false",
		},
	},
}
