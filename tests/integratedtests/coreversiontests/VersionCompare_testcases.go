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

package coreversiontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var versionCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "VersionCompare returns Equal -- identical versions v0.0.1",
		ArrangeInput: args.Map{
			"when":  "given equal versions",
			"left":  "v0.0.1",
			"right": "v0.0.1",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
	{
		Title: "VersionCompare returns LeftGreater -- left major v3.0 vs right v0.2.1",
		ArrangeInput: args.Map{
			"when":  "given left major version greater",
			"left":  "v3.0",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftGreater",
		},
	},
	{
		Title: "VersionCompare returns LeftLess -- left minor v0.0.2 vs right v0.2.1",
		ArrangeInput: args.Map{
			"when":  "given left minor version less",
			"left":  "v0.0.2",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftLess",
		},
	},
	{
		Title: "VersionCompare returns Equal -- zero-padded v4 vs v4.0",
		ArrangeInput: args.Map{
			"when":  "given v4 vs v4.0",
			"left":  "v4",
			"right": "v4.0",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
}
