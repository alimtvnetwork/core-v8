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

package corecsvtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var anyItemsToStringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "Single item returns its string",
		ArrangeInput:  args.Map{"items": []any{"hello"}},
		ExpectedInput: args.Map{"notEmpty": true},
	},
	{
		Title:         "Multiple items returns csv",
		ArrangeInput:  args.Map{"items": []any{"a", "b", "c"}},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var compileStringersToCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title:         "Empty funcs returns empty",
		ArrangeInput:  args.Map{"count": 0},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title:         "Single quote mode",
		ArrangeInput:  args.Map{
			"count": 1,
			"includeQuote": true,
			"singleQuote": true,
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "Double quote mode",
		ArrangeInput:  args.Map{
			"count": 1,
			"includeQuote": true,
			"singleQuote": false,
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "No quote mode",
		ArrangeInput:  args.Map{
			"count": 1,
			"includeQuote": false,
			"singleQuote": false,
		},
		ExpectedInput: args.Map{"length": 1},
	},
}
