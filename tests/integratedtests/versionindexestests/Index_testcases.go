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

package versionindexestests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var jsonRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "Patch JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Patch index",
			"index": "Patch",
		},
		ExpectedInput: args.Map{
			"indexName":  "Patch",
			"indexValue": "2",
		},
	},
	{
		Title: "Major JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Major index",
			"index": "Major",
		},
		ExpectedInput: args.Map{
			"indexName":  "Major",
			"indexValue": "0",
		},
	},
	{
		Title: "Build JSON roundtrip produces valid JSON string",
		ArrangeInput: args.Map{
			"when":  "given Build index",
			"index": "Build",
		},
		ExpectedInput: args.Map{
			"indexName":  "Build",
			"indexValue": "3",
		},
	},
}

var nameAndNameValueTestCases = []coretestcases.CaseV1{
	{
		Title: "Minor Name returns Minor",
		ArrangeInput: args.Map{
			"when":  "given Minor index",
			"index": "Minor",
		},
		ExpectedInput: args.Map{
			"name":      "Minor",
			"nameValue": "Minor(1)",
		},
	},
	{
		Title: "Patch Name returns Patch",
		ArrangeInput: args.Map{
			"when":  "given Patch index",
			"index": "Patch",
		},
		ExpectedInput: args.Map{
			"name":      "Patch",
			"nameValue": "Patch(2)",
		},
	},
}

var jsonParseSelfInjectTestCases = []coretestcases.CaseV1{
	{
		Title: "JsonParseSelfInject overwrites Minor with Patch JSON",
		ArrangeInput: args.Map{
			"when":   "given Patch JSON injected into Minor",
			"source": "Patch",
			"target": "Minor",
		},
		ExpectedInput: args.Map{
			"resultName":      "Patch",
			"resultNameValue": "Patch(2)",
		},
	},
	{
		Title: "JsonParseSelfInject overwrites Build with Major JSON",
		ArrangeInput: args.Map{
			"when":   "given Major JSON injected into Build",
			"source": "Major",
			"target": "Build",
		},
		ExpectedInput: args.Map{
			"resultName":      "Major",
			"resultNameValue": "Major(0)",
		},
	},
}
