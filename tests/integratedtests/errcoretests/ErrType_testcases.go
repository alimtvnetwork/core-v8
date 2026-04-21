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

package errcoretests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var errTypeCombineTestCases = []coretestcases.CaseV1{
	{
		Title: "Combine returns formatted string -- message 'some 2' and ref 'alim-1'",
		ArrangeInput: args.Map{
			"when":    "given message and reference",
			"message": "some 2",
			"ref":     "alim-1",
		},
		ExpectedInput: ".*some 2.*alim-1.*",
	},
	{
		Title: "Combine returns string with reference -- empty message, ref 'alim-2 no msg'",
		ArrangeInput: args.Map{
			"when":    "given empty message with reference",
			"message": "",
			"ref":     "alim-2 no msg",
		},
		ExpectedInput: ".*alim-2 no msg.*",
	},
	{
		Title: "Combine returns type name only -- empty message and empty ref",
		ArrangeInput: args.Map{
			"when":    "given both empty",
			"message": "",
			"ref":     "",
		},
		ExpectedInput: ".*Bytes data either nil or empty.*",
	},
}

var errMergeTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors returns nil -- both errors nil",
		ArrangeInput: args.Map{
			"when":     "given both nil errors",
			"hasError": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "MergeErrors returns error -- one non-nil error",
		ArrangeInput: args.Map{
			"when":     "given one real error",
			"hasError": true,
		},
		ExpectedInput: "false",
	},
}

var errTypeErrorNoRefsTestCases = []coretestcases.CaseV1{
	{
		Title: "ErrorNoRefs returns non-nil error -- message 'something broke'",
		ArrangeInput: args.Map{
			"when":    "given a message",
			"message": "something broke",
		},
		ExpectedInput: "true",
	},
	{
		Title: "ErrorNoRefs returns non-nil error -- empty message",
		ArrangeInput: args.Map{
			"when":    "given empty message",
			"message": "",
		},
		ExpectedInput: "true",
	},
}

var errTypeErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Error returns string with both -- message 'parsing failed' and ref 'line-42'",
		ArrangeInput: args.Map{
			"when":    "given message and ref",
			"message": "parsing failed",
			"ref":     "line-42",
		},
		ExpectedInput: ".*parsing failed.*line-42.*",
	},
	{
		Title: "Error returns string with message -- message 'some error', empty ref",
		ArrangeInput: args.Map{
			"when":    "given message only",
			"message": "some error",
			"ref":     "",
		},
		ExpectedInput: ".*some error.*",
	},
}
