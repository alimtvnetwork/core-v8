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
	"time"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var baseTime = time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
var laterTime = baseTime.Add(10 * time.Minute)

var timeCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "Time returns Equal -- identical times",
		ArrangeInput: args.Map{
			"when":  "given identical time values",
			"left":  baseTime,
			"right": baseTime,
		},
		ExpectedInput: "Equal", // compareResult
	},
	{
		Title: "Time returns LeftLess -- left before right",
		ArrangeInput: args.Map{
			"when":  "given left time before right time",
			"left":  baseTime,
			"right": laterTime,
		},
		ExpectedInput: "LeftLess", // compareResult
	},
	{
		Title: "Time returns LeftGreater -- left after right",
		ArrangeInput: args.Map{
			"when":  "given left time after right time",
			"left":  laterTime,
			"right": baseTime,
		},
		ExpectedInput: "LeftGreater", // compareResult
	},
	{
		Title: "Time returns LeftGreater -- small nanosecond difference forward",
		ArrangeInput: args.Map{
			"when":  "given left time slightly after right by nanoseconds",
			"left":  baseTime.Add(time.Duration(600000)),
			"right": baseTime,
		},
		ExpectedInput: "LeftGreater", // compareResult
	},
	{
		Title: "Time returns LeftLess -- small nanosecond difference reverse",
		ArrangeInput: args.Map{
			"when":  "given left time slightly before right by nanoseconds",
			"left":  baseTime,
			"right": baseTime.Add(time.Duration(600000)),
		},
		ExpectedInput: "LeftLess", // compareResult
	},
}
