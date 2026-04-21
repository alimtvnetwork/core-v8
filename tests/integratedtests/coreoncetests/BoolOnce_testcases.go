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
// BoolOnce -- Core (Value + String)
// =============================================================================

type boolOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue bool
}

var boolOnceCoreTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns true and String returns 'true' -- init true",
			ExpectedInput: args.Map{
				"value":  true,
				"string": "true",
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns false and String returns 'false' -- init false",
			ExpectedInput: args.Map{
				"value":  false,
				"string": "false",
			},
		},
		InitValue: false,
	},
}

// =============================================================================
// BoolOnce -- Caching (call count verification)
// =============================================================================

var boolOnceCachingTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns cached true -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        true,
				"r2":        true,
				"r3":        true,
				"callCount": 1,
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.Value returns cached false -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        false,
				"r2":        false,
				"r3":        false,
				"callCount": 1,
			},
		},
		InitValue: false,
	},
}

// =============================================================================
// BoolOnce -- JSON
// =============================================================================

var boolOnceJsonTestCases = []boolOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.MarshalJSON returns 'true' -- init true",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "true",
			},
		},
		InitValue: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BoolOnce.MarshalJSON returns 'false' -- init false",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "false",
			},
		},
		InitValue: false,
	},
}
