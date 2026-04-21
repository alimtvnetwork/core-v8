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

package codestacktests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var traceCollectionBasicTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection empty -- Length 0 IsEmpty true",
		ArrangeInput: args.Map{
			"when":  "empty collection",
			"count": 0,
		},
		ExpectedInput: args.Map{
			"length":     0,
			"isEmpty":    true,
			"hasAnyItem": false,
		},
	},
	{
		Title: "TraceCollection with items -- Length > 0 IsEmpty false",
		ArrangeInput: args.Map{
			"when":  "collection with 3 items",
			"count": 3,
		},
		ExpectedInput: args.Map{
			"length":     3,
			"isEmpty":    false,
			"hasAnyItem": true,
		},
	},
}

var traceCollectionFirstLastTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection First/Last -- returns correct items",
		ArrangeInput: args.Map{
			"when": "collection with 3 items",
		},
		ExpectedInput: args.Map{
			"firstPkg": "pkg1",
			"lastPkg":  "pkg3",
			"lastIdx":  2,
		},
	},
}

var traceCollectionSkipTakeTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Skip 1 -- returns remaining items",
		ArrangeInput: args.Map{
			"when": "skip first item",
			"skip": 1,
		},
		ExpectedInput: args.Map{
			"length":   2,
			"firstPkg": "pkg2",
		},
	},
	{
		Title: "TraceCollection Take 2 -- returns first 2 items",
		ArrangeInput: args.Map{
			"when": "take first 2",
			"take": 2,
		},
		ExpectedInput: args.Map{
			"length":  2,
			"lastPkg": "pkg2",
		},
	},
}

var traceCollectionReverseTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Reverse 3 items -- reverses order",
		ArrangeInput: args.Map{
			"when": "reverse 3 item collection",
		},
		ExpectedInput: args.Map{
			"firstPkg": "pkg3",
			"lastPkg":  "pkg1",
		},
	},
}

var traceCollectionFilterTestCases = []coretestcases.CaseV1{
	{
		Title: "FilterPackageNameTraceCollection -- filters by package",
		ArrangeInput: args.Map{
			"when":    "filter by pkg2",
			"package": "pkg2",
		},
		ExpectedInput: args.Map{
			"length":   1,
			"firstPkg": "pkg2",
		},
	},
}
