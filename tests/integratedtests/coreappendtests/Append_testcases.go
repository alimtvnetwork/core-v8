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

package coreappendtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var prependAppendTestCases = []coretestcases.CaseV1{
	{
		Title: "PrependAppendAnyItemsToStringsSkipOnNil returns 3 items -- prepend and append provided",
		ArrangeInput: args.Map{
			"when":    "given prepend, append and middle items with nil",
			"prepend": "start",
			"append":  "end",
		},
		ExpectedInput: args.Map{
			"totalCount": "3",
			"firstItem":  "start",
			"lastItem":   "end",
		},
	},
	{
		Title: "PrependAppendAnyItemsToStringsSkipOnNil returns 2 items -- nil prepend",
		ArrangeInput: args.Map{
			"when":   "given nil prepend",
			"append": "end",
		},
		ExpectedInput: args.Map{
			"totalCount": "2",
			"firstItem":  "middle",
			"lastItem":   "end",
		},
	},
}
