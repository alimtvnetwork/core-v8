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

package isanytests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isAnyDefinedNullTestCases = []coretestcases.CaseV1{
	{
		Title: "nil is Null and not Defined",
		ArrangeInput: args.Map{
			"when":  "given nil value",
			"input": nil,
		},
		ExpectedInput: args.Map{
			"isDefined": "false",
			"isNull":    "true",
		},
	},
	{
		Title: "non-nil error is Defined and not Null",
		ArrangeInput: args.Map{
			"when":     "given a non-nil error",
			"input":    "error-marker",
			"useError": true,
		},
		ExpectedInput: args.Map{
			"isDefined": "true",
			"isNull":    "false",
		},
	},
	{
		Title: "empty string is Defined (not nil)",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: args.Map{
			"isDefined": "true",
			"isNull":    "false",
		},
	},
	{
		Title: "integer zero is Defined",
		ArrangeInput: args.Map{
			"when":  "given integer zero",
			"input": 0,
		},
		ExpectedInput: args.Map{
			"isDefined": "true",
			"isNull":    "false",
		},
	},
}

var isAnyBothTestCases = []coretestcases.CaseV1{
	{
		Title: "DefinedBoth(nil, non-nil) returns false",
		ArrangeInput: args.Map{
			"when":   "given nil and non-nil",
			"first":  nil,
			"second": "something",
		},
		ExpectedInput: args.Map{
			"definedBoth": "false",
			"nullBoth":    "false",
		},
	},
	{
		Title: "NullBoth(nil, nil) returns true",
		ArrangeInput: args.Map{
			"when":   "given both nil",
			"first":  nil,
			"second": nil,
		},
		ExpectedInput: args.Map{
			"definedBoth": "false",
			"nullBoth":    "true",
		},
	},
	{
		Title: "DefinedBoth(string, string) returns true",
		ArrangeInput: args.Map{
			"when":   "given both defined strings",
			"first":  "a",
			"second": "b",
		},
		ExpectedInput: args.Map{
			"definedBoth": "true",
			"nullBoth":    "false",
		},
	},
}
