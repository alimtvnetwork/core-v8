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

package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// IntegersOnce -- Core
// =============================================================================

type integersOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue []int
}

var integersOnceCoreTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce returns length 3 and isEmpty false -- [3,1,2] input",
			ExpectedInput: args.Map{
				"length":  3,
				"isEmpty": false,
				"isZero":  false,
			},
		},
		InitValue: []int{3, 1, 2},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce returns length 0 and isEmpty true -- empty input",
			ExpectedInput: args.Map{
				"length":  0,
				"isEmpty": true,
				"isZero":  true,
			},
		},
		InitValue: []int{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce returns length 0 and isEmpty true -- nil input",
			ExpectedInput: args.Map{
				"length":  0,
				"isEmpty": true,
				"isZero":  true,
			},
		},
		InitValue: nil,
	},
}

// =============================================================================
// IntegersOnce -- Sorted
// =============================================================================

var integersOnceSortedTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce.Sorted returns [1,2,3] -- [3,1,2] input",
			ExpectedInput: args.Map{
				"first": 1,
				"last":  3,
			},
		},
		InitValue: []int{3, 1, 2},
	},
}

// =============================================================================
// IntegersOnce -- RangesMap / RangesBoolMap
// =============================================================================

var integersOnceRangesTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce returns rangesMapLength 3 -- [10,20,30] input",
			ExpectedInput: args.Map{
				"rangesMapLength":    3,
				"rangesBoolMapLen":   3,
				"uniqueMapLen":       3,
				"rangesMapHas10":     true,
				"rangesBoolMapHas20": true,
			},
		},
		InitValue: []int{10, 20, 30},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce returns rangesMapLength 0 -- empty input",
			ExpectedInput: args.Map{
				"rangesMapLength":    0,
				"rangesBoolMapLen":   0,
				"uniqueMapLen":       0,
				"rangesMapHas10":     false,
				"rangesBoolMapHas20": false,
			},
		},
		InitValue: []int{},
	},
}

// =============================================================================
// IntegersOnce -- IsEqual
// =============================================================================

var integersOnceIsEqualTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce.IsEqual returns true for same and false for different -- [1,2,3] input",
			ExpectedInput: args.Map{
				"isEqualSame":    true,
				"isEqualDiff":    false,
				"isEqualDiffLen": false,
			},
		},
		InitValue: []int{1, 2, 3},
	},
}

// =============================================================================
// IntegersOnce -- Caching
// =============================================================================

var integersOnceCachingTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce.Value caches -- initializer runs once",
			ExpectedInput: args.Map{
				"callCount": 1,
				"length":    2,
			},
		},
		InitValue: []int{5, 10},
	},
}

// =============================================================================
// IntegersOnce -- JSON
// =============================================================================

var integersOnceJsonTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegersOnce.MarshalJSON returns '[1,2]' -- [1,2] input",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "[1,2]",
			},
		},
		InitValue: []int{1, 2},
	},
}

// =============================================================================
// IntegersOnce -- Constructor
// =============================================================================

var integersOnceConstructorTestCases = []integersOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewIntegersOnce returns correct length -- [1,2,3] input",
			ExpectedInput: args.Map{
				"length": 3,
			},
		},
		InitValue: []int{1, 2, 3},
	},
}
