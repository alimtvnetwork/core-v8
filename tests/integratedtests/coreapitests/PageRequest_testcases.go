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

package coreapitests

import (
	"github.com/alimtvnetwork/core-v8/coredata/coreapi"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// IsPageSizeEmpty test cases
// =============================================================================

var pageRequestIsPageSizeEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsPageSizeEmpty - nil receiver returns true",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageSizeEmpty - zero returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 0, PageIndex: 1},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageSizeEmpty - negative returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: -1},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageSizeEmpty - positive returns false",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 10},
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// IsPageIndexEmpty test cases
// =============================================================================

var pageRequestIsPageIndexEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsPageIndexEmpty - nil receiver returns true",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageIndexEmpty - zero returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageIndex: 0, PageSize: 10},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageIndexEmpty - positive returns false",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageIndex: 2},
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// HasPageSize test cases
// =============================================================================

var pageRequestHasPageSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "HasPageSize - nil receiver returns false",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "false",
	},
	{
		Title: "HasPageSize - positive returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 25},
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// HasPageIndex test cases
// =============================================================================

var pageRequestHasPageIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "HasPageIndex - nil receiver returns false",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "false",
	},
	{
		Title: "HasPageIndex - positive returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageIndex: 3},
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// Clone test cases
// =============================================================================

var pageRequestCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone - nil receiver returns nil",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "true",
	},
}

var pageRequestCloneFieldsTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone - copies all fields",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 20, PageIndex: 5},
		},
		ExpectedInput: args.Map{
			"pageSize":  "20",
			"pageIndex": "5",
		},
	},
}

var pageRequestCloneIndependenceTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone - independence from original",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 20, PageIndex: 5},
		},
		ExpectedInput: args.Map{
			"pageSize":  "20",
			"pageIndex": "5",
		},
	},
}
