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

package defaultcapacitytests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var ofSplitsTestCases = []coretestcases.CaseV1{
	{
		Title: "No limit (-1) returns OfSearch result",
		ArrangeInput: args.Map{
			"wholeLength": 100,
			"limit":       -1,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title: "Limit >= wholeLength returns OfSearch result",
		ArrangeInput: args.Map{
			"wholeLength": 50,
			"limit":       100,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title: "Limit == wholeLength returns OfSearch result",
		ArrangeInput: args.Map{
			"wholeLength": 50,
			"limit":       50,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title: "Limit < wholeLength returns limit",
		ArrangeInput: args.Map{
			"wholeLength": 100,
			"limit":       25,
		},
		ExpectedInput: args.Map{"result": 25},
	},
}

var predictiveDefaultSmallTestCases = []coretestcases.CaseV1{
	{
		Title:         "Positive length yields positive result",
		ArrangeInput:  100,
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title:         "Zero length yields positive result from additionalCap",
		ArrangeInput:  0,
		ExpectedInput: args.Map{"isPositive": true},
	},
}

var predictiveFiftyPercentTestCases = []coretestcases.CaseV1{
	{
		Title:         "Positive length with additional cap",
		ArrangeInput:  args.Map{
			"possibleLen": 100,
			"additionalCap": 10,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title:         "Zero length returns additional cap",
		ArrangeInput:  args.Map{
			"possibleLen": 0,
			"additionalCap": 5,
		},
		ExpectedInput: args.Map{"result": 5},
	},
}
