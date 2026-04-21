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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// IntegerOnce -- Core (Value + String + comparisons)
// =============================================================================

type integerOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue int
}

var integerOnceCoreTestCases = []integerOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce returns isZero true and isEmpty true -- input 0",
			ExpectedInput: args.Map{
				"value":          0,
				"string":         "0",
				"isZero":         true,
				"isEmpty":        true,
				"isAboveZero":    false,
				"isPositive":     false,
				"isLessThanZero": false,
				"isNegative":     false,
			},
		},
		InitValue: 0,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce returns isPositive true and isZero false -- input 42",
			ExpectedInput: args.Map{
				"value":          42,
				"string":         "42",
				"isZero":         false,
				"isEmpty":        false,
				"isAboveZero":    true,
				"isPositive":     true,
				"isLessThanZero": false,
				"isNegative":     false,
			},
		},
		InitValue: 42,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce returns isNegative true and isZero false -- input -3",
			ExpectedInput: args.Map{
				"value":          -3,
				"string":         "-3",
				"isZero":         false,
				"isEmpty":        false,
				"isAboveZero":    false,
				"isPositive":     false,
				"isLessThanZero": true,
				"isNegative":     true,
			},
		},
		InitValue: -3,
	},
}

// =============================================================================
// IntegerOnce -- Caching
// =============================================================================

var integerOnceCachingTestCases = []integerOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        42,
				"r2":        42,
				"callCount": 1,
			},
		},
		InitValue: 42,
	},
}

// =============================================================================
// IntegerOnce -- Comparisons (IsAbove, IsLessThan)
// =============================================================================

type integerOnceCompareTestCase struct {
	Case         coretestcases.CaseV1
	InitValue    int
	CompareValue int
}

var integerOnceCompareTestCases = []integerOnceCompareTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce returns isAbove true and isAboveEqual true -- input 10 compare 5",
			ExpectedInput: args.Map{
				"isAboveCompare":   true,
				"isAboveSelf":      false,
				"isAboveEqualSelf": true,
			},
		},
		InitValue:    10,
		CompareValue: 5,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce returns isLessThan true and isLessThanEqual true -- input 3 compare 5",
			ExpectedInput: args.Map{
				"isLessThanCompare":   true,
				"isLessThanSelf":      false,
				"isLessThanEqualSelf": true,
			},
		},
		InitValue:    3,
		CompareValue: 5,
	},
}

// =============================================================================
// IntegerOnce -- JSON
// =============================================================================

var integerOnceJsonTestCases = []integerOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "IntegerOnce.MarshalJSON returns '42' -- input 42",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "42",
			},
		},
		InitValue: 42,
	},
}
