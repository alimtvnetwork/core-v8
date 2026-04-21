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
// Map — int to string
// ==========================================

var mapIntToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Map transforms int collection to string collection",
		ArrangeInput: args.Map{
			"when":  "given ints [1, 2, 3]",
			"items": []int{1, 2, 3},
		},
		ExpectedInput: []string{"3", "#1", "#2", "#3"},
	},
}

var mapEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Map on empty collection returns empty",
		ArrangeInput: args.Map{
			"when": "given empty int collection",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

var mapNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Map on nil collection returns empty",
		ArrangeInput: args.Map{
			"when": "given nil collection",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================
// Map — string to int (length)
// ==========================================

var mapStringToIntTestCases = []coretestcases.CaseV1{
	{
		Title: "Map transforms strings to their lengths",
		ArrangeInput: args.Map{
			"when":  "given strings [hi, hello, x]",
			"items": []string{"hi", "hello", "x"},
		},
		ExpectedInput: []string{"3", "2", "5", "1"},
	},
}

// ==========================================
// FlatMap — string to chars
// ==========================================

var flatMapTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMap flattens nested slices into single collection",
		ArrangeInput: args.Map{
			"when":  "given strings split to chars",
			"items": []string{"ab", "cd"},
		},
		ExpectedInput: []string{"4", "a", "b", "c", "d"},
	},
}

var flatMapEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMap on empty collection returns empty",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================
// Reduce — sum
// ==========================================

var reduceSumTestCases = []coretestcases.CaseV1{
	{
		Title: "Reduce sums int collection",
		ArrangeInput: args.Map{
			"when":  "given ints [10, 20, 30]",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: "60",
	},
}

var reduceEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Reduce on empty returns initial value",
		ArrangeInput: args.Map{
			"when":    "given empty int collection",
			"initial": 99,
		},
		ExpectedInput: "99",
	},
}

// ==========================================
// Reduce — string concat
// ==========================================

var reduceConcatTestCases = []coretestcases.CaseV1{
	{
		Title: "Reduce concatenates strings",
		ArrangeInput: args.Map{
			"when":  "given strings [a, b, c]",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: "a-b-c",
	},
}

// ==========================================
// Map chained — Map then Filter
// ==========================================

var mapThenFilterTestCases = []coretestcases.CaseV1{
	{
		Title: "Map then Filter produces correct subset",
		ArrangeInput: args.Map{
			"when":  "given ints mapped to doubled then filtered > 5",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{"3", "6", "8", "10"},
	},
}
