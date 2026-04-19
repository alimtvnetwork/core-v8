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

package keymktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var keyLegendGroupIntRangeTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupIntRange returns correct key range -- range 5 to 10",
		ArrangeInput: args.Map{
			"when":    "given range 5 to 10",
			"root":    "cimux",
			"package": "main",
			"group":   "myg",
			"state":   "stateName",
			"startId": 5,
			"endId":   10,
		},
		ExpectedInput: args.Map{
			"count":    "6",
			"firstKey": "cimux-main-5-stateName",
			"lastKey":  "cimux-main-10-stateName",
		},
	},
}

var keyLegendUserStringWithoutStateTestCases = []coretestcases.CaseV1{
	{
		Title: "UserStringWithoutState returns root-package-group-user -- user 'mynewuser1'",
		ArrangeInput: args.Map{
			"when":    "given user mynewuser1",
			"root":    "cimux",
			"package": "main",
			"group":   "myg",
			"state":   "stateName",
			"user":    "mynewuser1",
		},
		ExpectedInput: "cimux-main-myg-mynewuser1",
	},
}

var keyLegendUpToStateTestCases = []coretestcases.CaseV1{
	{
		Title: "UpToState returns root-package-group-state-user -- user 'my-user'",
		ArrangeInput: args.Map{
			"when":    "given user my-user",
			"root":    "cimux",
			"package": "main",
			"group":   "myg",
			"state":   "stateName",
			"user":    "my-user",
		},
		ExpectedInput: "cimux-main-myg-stateName-my-user",
	},
}
