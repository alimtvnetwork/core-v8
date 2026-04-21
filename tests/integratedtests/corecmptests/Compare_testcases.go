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

package corecmptests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var integerCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "Integer returns Equal -- same values 5 and 5",
		ArrangeInput: args.Map{
			"when":  "given equal integers",
			"left":  5,
			"right": 5,
		},
		ExpectedInput: "Equal", // compareResult
	},
	{
		Title: "Integer returns LeftLess -- left 3 right 7",
		ArrangeInput: args.Map{
			"when":  "given left less than right",
			"left":  3,
			"right": 7,
		},
		ExpectedInput: "LeftLess", // compareResult
	},
	{
		Title: "Integer returns LeftGreater -- left 10 right 2",
		ArrangeInput: args.Map{
			"when":  "given left greater than right",
			"left":  10,
			"right": 2,
		},
		ExpectedInput: "LeftGreater", // compareResult
	},
}

var isStringsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStringsEqual returns true -- identical slices",
		ArrangeInput: args.Map{
			"when":  "given identical string slices",
			"left":  []string{"a", "b", "c"},
			"right": []string{"a", "b", "c"},
		},
		ExpectedInput: "true", // isEqual
	},
	{
		Title: "IsStringsEqual returns false -- different values",
		ArrangeInput: args.Map{
			"when":  "given different string slices",
			"left":  []string{"a", "b"},
			"right": []string{"a", "c"},
		},
		ExpectedInput: "false", // isEqual
	},
	{
		Title: "IsStringsEqual returns false -- different lengths",
		ArrangeInput: args.Map{
			"when":  "given slices of different length",
			"left":  []string{"a"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: "false", // isEqual
	},
}
