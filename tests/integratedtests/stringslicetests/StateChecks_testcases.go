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

package stringslicetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var srcIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmpty returns true -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmpty returns true -- empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmpty returns false -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var srcHasAnyItemTestCases = []coretestcases.CaseV1{
	{
		Title: "HasAnyItem returns false -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "HasAnyItem returns true -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var srcEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty returns empty slice -- no input",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcIsEmptyPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyPtr returns true -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmptyPtr returns true -- empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmptyPtr returns false -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var srcHasAnyItemPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "HasAnyItemPtr returns false -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "HasAnyItemPtr returns true -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var srcEmptyPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "EmptyPtr returns empty slice -- no input",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}
