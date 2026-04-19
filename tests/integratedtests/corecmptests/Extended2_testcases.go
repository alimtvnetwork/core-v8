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
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isStringsEqualPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsStringsEqualPtr returns true -- both nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsStringsEqualPtr returns false -- left nil right non-nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": false,
			"right": []string{"a"},
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsStringsEqualPtr returns false -- right nil left non-nil",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": true,
			"left": []string{"a"},
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsStringsEqualPtr returns true -- equal slices",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": false,
			"left": []string{"a", "b"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsStringsEqualPtr returns false -- different length",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": false,
			"left": []string{"a"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var timePtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "TimePtr returns Equal -- both nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": true,
		},
		ExpectedInput: args.Map{"isEqual": true},
	},
	{
		Title:         "TimePtr returns NotEqual -- left nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": false,
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "TimePtr returns NotEqual -- right nil",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": true,
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "TimePtr returns Equal -- both same time",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": false,
			"sameTime": true,
		},
		ExpectedInput: args.Map{"isEqual": true},
	},
}
