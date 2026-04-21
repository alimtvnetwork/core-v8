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
// ByteOnce -- Core
// =============================================================================

type byteOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue byte
}

var byteOnceCoreTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce returns value 42 and isPositive true -- input 42",
			ExpectedInput: args.Map{
				"value":      42,
				"int":        42,
				"string":     "42",
				"isEmpty":    false,
				"isZero":     false,
				"isNegative": false,
				"isPositive": true,
			},
		},
		InitValue: 42,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce returns isZero true and isEmpty true -- input 0",
			ExpectedInput: args.Map{
				"value":      0,
				"int":        0,
				"string":     "0",
				"isEmpty":    true,
				"isZero":     true,
				"isNegative": false,
				"isPositive": false,
			},
		},
		InitValue: 0,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce returns isPositive true -- input 255 max byte value",
			ExpectedInput: args.Map{
				"value":      255,
				"int":        255,
				"string":     "255",
				"isEmpty":    false,
				"isZero":     false,
				"isNegative": false,
				"isPositive": true,
			},
		},
		InitValue: 255,
	},
}

// =============================================================================
// ByteOnce -- Caching
// =============================================================================

var byteOnceCachingTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        10,
				"r2":        10,
				"callCount": 1,
			},
		},
		InitValue: 10,
	},
}

// =============================================================================
// ByteOnce -- JSON
// =============================================================================

var byteOnceJsonTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce.MarshalJSON returns '99' -- input 99",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "99",
			},
		},
		InitValue: 99,
	},
}

// =============================================================================
// ByteOnce -- Serialize
// =============================================================================

var byteOnceSerializeTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ByteOnce.Serialize returns '77' -- input 77",
			ExpectedInput: args.Map{
				"noError":         true,
				"serializedValue": "77",
			},
		},
		InitValue: 77,
	},
}

// =============================================================================
// ByteOnce -- Constructor
// =============================================================================

var byteOnceConstructorTestCases = []byteOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewByteOnce returns correct value -- input 5",
			ExpectedInput: args.Map{
				"constructedValue": 5,
			},
		},
		InitValue: 5,
	},
}
