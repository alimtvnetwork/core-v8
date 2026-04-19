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
// MapStringStringOnce -- Core
// =============================================================================

type mapSSOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue map[string]string
}

var mapSSOnceCoreTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce returns length 2 and hasAnyItem true -- {a:1,b:2} input",
			ExpectedInput: args.Map{
				"length":     2,
				"isEmpty":    false,
				"hasAnyItem": true,
			},
		},
		InitValue: map[string]string{"a": "1", "b": "2"},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce returns length 0 and isEmpty true -- empty map input",
			ExpectedInput: args.Map{
				"length":     0,
				"isEmpty":    true,
				"hasAnyItem": false,
			},
		},
		InitValue: map[string]string{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce returns length 0 and isEmpty true -- nil input",
			ExpectedInput: args.Map{
				"length":     0,
				"isEmpty":    true,
				"hasAnyItem": false,
			},
		},
		InitValue: nil,
	},
}

// =============================================================================
// MapStringStringOnce -- Lookup (Has, IsContains, IsMissing, GetValue)
// =============================================================================

var mapSSOnceContainsTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce.Has returns true and GetValue returns 'v1' -- {k1:v1,k2:v2} input",
			ExpectedInput: args.Map{
				"hasK1":      true,
				"containsK2": true,
				"isMissingX": true,
				"hasAllK1K2": true,
				"getK1":      "v1",
			},
		},
		InitValue: map[string]string{"k1": "v1", "k2": "v2"},
	},
}

// =============================================================================
// MapStringStringOnce -- Keys / Values / Sorted
// =============================================================================

var mapSSOnceKeysValuesTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce.AllKeysSorted returns [a,b] and AllValuesSorted returns [1,2] -- {b:2,a:1} input",
			ExpectedInput: args.Map{
				"keysLen":          2,
				"valuesLen":        2,
				"sortedFirstKey":   "a",
				"sortedLastKey":    "b",
				"sortedFirstValue": "1",
				"sortedLastValue":  "2",
			},
		},
		InitValue: map[string]string{"b": "2", "a": "1"},
	},
}

// =============================================================================
// MapStringStringOnce -- IsEqual
// =============================================================================

var mapSSOnceIsEqualTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce.IsEqual returns true for same and false for different -- {a:1} input",
			ExpectedInput: args.Map{
				"isEqualSame":    true,
				"isEqualDiffVal": false,
				"isEqualDiffKey": false,
				"isEqualDiffLen": false,
			},
		},
		InitValue: map[string]string{"a": "1"},
	},
}

// =============================================================================
// MapStringStringOnce -- Caching
// =============================================================================

var mapSSOnceCachingTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce.Value returns cached result -- initializer runs once",
			ExpectedInput: args.Map{
				"callCount": 1,
				"length":    2,
			},
		},
		InitValue: map[string]string{"x": "1", "y": "2"},
	},
}

// =============================================================================
// MapStringStringOnce -- JSON
// =============================================================================

var mapSSOnceJsonTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "MapStringStringOnce.MarshalJSON returns bytes without error -- {a:1} input",
			ExpectedInput: args.Map{
				"noError":             true,
				"dataLengthAboveZero": true,
			},
		},
		InitValue: map[string]string{"a": "1"},
	},
}

// =============================================================================
// MapStringStringOnce -- Constructor
// =============================================================================

var mapSSOnceConstructorTestCases = []mapSSOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewMapStringStringOnce returns length 1 -- single key-value input",
			ExpectedInput: args.Map{
				"length": 1,
			},
		},
		InitValue: map[string]string{"k": "v"},
	},
}
