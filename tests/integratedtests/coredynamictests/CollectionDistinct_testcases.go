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
// Distinct
// ==========================================

var collectionDistinctTestCases = []coretestcases.CaseV1{
	{
		Title: "Distinct returns unique items preserving order -- [a,b,a,c,b,a]",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a", "c", "b", "a"},
		},
		ExpectedInput: args.Map{
			"distinctCount": 3,
			"item0":         "a",
			"item1":         "b",
			"item2":         "c",
		},
	},
	{
		Title: "Distinct returns same items -- already unique [x,y,z]",
		ArrangeInput: args.Map{
			"items": []string{"x", "y", "z"},
		},
		ExpectedInput: args.Map{
			"distinctCount": 3,
			"item0":         "x",
			"item1":         "y",
			"item2":         "z",
		},
	},
	{
		Title: "Distinct returns empty -- empty input",
		ArrangeInput: args.Map{
			"items": []string{},
		},
		ExpectedInput: "0",
	},
}

// ==========================================
// DistinctCount
// ==========================================

var collectionDistinctCountTestCases = []coretestcases.CaseV1{
	{
		Title: "DistinctCount returns 3 -- [a,b,a,c,b]",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a", "c", "b"},
		},
		ExpectedInput: "3",
	},
}

// ==========================================
// IsDistinct
// ==========================================

var collectionIsDistinctTestCases = []coretestcases.CaseV1{
	{
		Title: "IsDistinct returns true -- unique [a,b,c]",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsDistinct returns false -- duplicates [a,b,a]",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "a"},
		},
		ExpectedInput: "false",
	},
}
