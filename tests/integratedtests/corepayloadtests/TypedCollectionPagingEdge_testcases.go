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

package corepayloadtests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// Edge cases for paging — single item, page size 1, large page size
// =============================================================================

var typedCollectionPagingEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollection returns 1 page -- 1 item page size 1",
		ArrangeInput: args.Map{
			"when":     "1 item with page size 1",
			"count":    1,
			"pageSize": 1,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 1,
		},
	},
	{
		Title: "GetPagedCollection returns 3 pages -- 3 items page size 1",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 1",
			"count":    3,
			"pageSize": 1,
		},
		ExpectedInput: args.Map{
			"pageCount":  3,
			"page1Items": 1,
			"page2Items": 1,
			"page3Items": 1,
		},
	},
	{
		Title: "GetPagedCollection returns 1 page -- 5 items page size equals count",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 5",
			"count":    5,
			"pageSize": 5,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 5,
		},
	},
	{
		Title: "GetPagedCollection returns 1 page -- 3 items page size 100 larger than count",
		ArrangeInput: args.Map{
			"when":     "3 items with page size 100",
			"count":    3,
			"pageSize": 100,
		},
		ExpectedInput: args.Map{
			"pageCount":  1,
			"page1Items": 3,
		},
	},
}

// =============================================================================
// GetSinglePageCollection edge cases
// =============================================================================

var typedCollectionSinglePageEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetSinglePageCollection returns 3 items -- middle page 2 of 9 items page size 3",
		ArrangeInput: args.Map{
			"when":      "page 2 of 3 items per page from 9 items",
			"count":     9,
			"pageSize":  3,
			"pageIndex": 2,
		},
		ExpectedInput: args.Map{
			"pageItemCount": 3,
			"item0":         "user-3",
			"item1":         "user-4",
			"item2":         "user-5",
		},
	},
	{
		Title: "GetSinglePageCollection returns 1 item -- page 2 page size 1 from 5 items",
		ArrangeInput: args.Map{
			"when":      "page 2 of page size 1 from 5 items",
			"count":     5,
			"pageSize":  1,
			"pageIndex": 2,
		},
		ExpectedInput: args.Map{
			"pageItemCount": 1,
			"item0":         "user-1",
		},
	},
}

// =============================================================================
// GetPagedCollectionWithInfo edge cases
// =============================================================================

var typedCollectionPagedWithInfoEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollectionWithInfo returns correct metadata -- 1 item page size 5",
		ArrangeInput: args.Map{
			"when":     "1 item with page size 5",
			"count":    1,
			"pageSize": 5,
		},
		ExpectedInput: args.Map{
			"pageCount":          1,
			"p1CurrentPageIndex": 1,
			"p1TotalPages":       1,
			"p1PerPageItems":     5,
			"p1TotalItems":       1,
		},
	},
	{
		Title: "GetPagedCollectionWithInfo returns correct metadata -- 4 items page size 2 exact division",
		ArrangeInput: args.Map{
			"when":     "4 items with page size 2",
			"count":    4,
			"pageSize": 2,
		},
		ExpectedInput: args.Map{
			"pageCount":          2,
			"p1CurrentPageIndex": 1,
			"p1TotalPages":       2,
			"p1PerPageItems":     2,
			"p1TotalItems":       4,
			"p2CurrentPageIndex": 2,
			"p2TotalPages":       2,
			"p2PerPageItems":     2,
			"p2TotalItems":       4,
		},
	},
}

// =============================================================================
// GetPagesSize edge cases
// =============================================================================

var typedCollectionPagesSizeEdgeCases = []coretestcases.CaseV1{
	{
		Title: "GetPagesSize returns 5 -- 5 items page size 1",
		ArrangeInput: args.Map{
			"when":     "5 items with page size 1",
			"count":    5,
			"pageSize": 1,
		},
		ExpectedInput: args.Map{
			"pagesSize": 5,
		},
	},
	{
		Title: "GetPagesSize returns 1 -- 1 item page size 10",
		ArrangeInput: args.Map{
			"when":     "1 item with page size 10",
			"count":    1,
			"pageSize": 10,
		},
		ExpectedInput: args.Map{
			"pagesSize": 1,
		},
	},
}

// ==========================================================================
// Paging empty with GetPagedCollectionWithInfo
// ==========================================================================

var typedCollectionPagingWithInfoEmptyTestCase = coretestcases.CaseV1{
	Title: "GetPagedCollectionWithInfo returns 1 page 0 items -- empty collection",
	ExpectedInput: args.Map{
		"pageCount":    1,
		"firstPageLen": 0,
	},
}
