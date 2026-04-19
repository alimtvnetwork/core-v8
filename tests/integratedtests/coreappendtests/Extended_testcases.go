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

var appendAnyItemsTestCases = []coretestcases.CaseV1{
	{
		Title:         "AppendAnyItemsToString comma -- not empty",
		ArrangeInput:  args.Map{
			"when": "comma joiner",
			"joiner": ", ",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var prependAnyItemsTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAnyItemsToString comma -- not empty",
		ArrangeInput:  args.Map{
			"when": "comma joiner",
			"joiner": ", ",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var prependAppendStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppendAnyItemsToString -- not empty",
		ArrangeInput:  args.Map{
			"when": "comma joiner",
			"joiner": ", ",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var prependAppendSkipNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppendSkipOnNil -- skips nil items (3 total: prefix + item1 + item3 + suffix)",
		ArrangeInput:  args.Map{"when": "nil in middle"},
		ExpectedInput: args.Map{"length": 4},
	},
}

var prependAppendNilPrependTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppend nil prepend -- skips prepend (item1 + suffix)",
		ArrangeInput:  args.Map{"when": "nil prepend"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var prependAppendNilAppendTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppend nil append -- skips append (prefix + item1)",
		ArrangeInput:  args.Map{"when": "nil append"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var mapAppendTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapStringStringAppend -- appends 2 items to existing 1",
		ArrangeInput:  args.Map{"when": "append 2 items"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var mapAppendSkipEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapStringStringAppend skip empty -- skips empty value",
		ArrangeInput:  args.Map{"when": "skip empty"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var mapAppendEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapStringStringAppend empty append -- unchanged",
		ArrangeInput:  args.Map{"when": "empty append map"},
		ExpectedInput: args.Map{"length": 1},
	},
}
