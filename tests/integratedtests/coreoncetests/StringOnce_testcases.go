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
// StringOnce -- Core (Value + String queries)
// =============================================================================

type stringOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
}

var stringOnceCoreTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce returns value 'hello' and isEmpty false -- 'hello' input",
			ExpectedInput: args.Map{
				"value":               "hello",
				"string":              "hello",
				"isEmpty":             false,
				"isEmptyOrWhitespace": false,
			},
		},
		InitValue: "hello",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce returns isEmpty true and isEmptyOrWhitespace true -- empty input",
			ExpectedInput: args.Map{
				"value":               "",
				"string":              "",
				"isEmpty":             true,
				"isEmptyOrWhitespace": true,
			},
		},
		InitValue: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce returns isEmpty false and isEmptyOrWhitespace true -- whitespace input",
			ExpectedInput: args.Map{
				"value":               "   ",
				"string":              "   ",
				"isEmpty":             false,
				"isEmptyOrWhitespace": true,
			},
		},
		InitValue: "   ",
	},
}

// =============================================================================
// StringOnce -- Caching
// =============================================================================

var stringOnceCachingTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        "hello",
				"r2":        "hello",
				"r3":        "hello",
				"callCount": 1,
			},
		},
		InitValue: "hello",
	},
}

// =============================================================================
// StringOnce -- String matching
// =============================================================================

type stringOnceMatchTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
	MatchArg  string
}

var stringOnceMatchTestCases = []stringOnceMatchTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.IsEqual returns true for 'abc' and false for 'xyz' -- 'abc' input",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "abc",
		MatchArg:  "abc",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.IsContains returns true for 'world' and false for 'xyz' -- 'hello world' input",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "hello world",
		MatchArg:  "world",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.HasPrefix returns true for 'prefix' and false for 'data' -- 'prefix-data' input",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "prefix-data",
		MatchArg:  "prefix",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.HasSuffix returns true for 'suffix' and false for 'data' -- 'data-suffix' input",
			ExpectedInput: args.Map{
				"matchResult":   true,
				"noMatchResult": false,
			},
		},
		InitValue: "data-suffix",
		MatchArg:  "suffix",
	},
}

// =============================================================================
// StringOnce -- Split
// =============================================================================

type stringOnceSplitTestCase struct {
	Case      coretestcases.CaseV1
	InitValue string
	Splitter  string
	Method    string // "splitBy", "splitLeftRight", "splitLeftRightTrim"
}

var stringOnceSplitTestCases = []stringOnceSplitTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.SplitBy returns 3 parts -- 'a,b,c' split by ','",
			ExpectedInput: args.Map{
				"partsLength": 3,
				"firstPart":   "a",
				"lastPart":    "c",
			},
		},
		InitValue: "a,b,c",
		Splitter:  ",",
		Method:    "splitBy",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.SplitLeftRight returns 'key' and 'value' -- 'key=value' split by '='",
			ExpectedInput: args.Map{
				"left":  "key",
				"right": "value",
			},
		},
		InitValue: "key=value",
		Splitter:  "=",
		Method:    "splitLeftRight",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.SplitLeftRight returns full left and empty right -- 'nosplit' no separator found",
			ExpectedInput: args.Map{
				"left":  "nosplit",
				"right": "",
			},
		},
		InitValue: "nosplit",
		Splitter:  "=",
		Method:    "splitLeftRight",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.SplitLeftRightTrim returns trimmed 'key' and 'value' -- ' key = value ' split by '='",
			ExpectedInput: args.Map{
				"left":  "key",
				"right": "value",
			},
		},
		InitValue: " key = value ",
		Splitter:  "=",
		Method:    "splitLeftRightTrim",
	},
}

// =============================================================================
// StringOnce -- JSON
// =============================================================================

var stringOnceJsonTestCases = []stringOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "StringOnce.MarshalJSON returns '\"json\"' -- 'json' input",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "\"json\"",
			},
		},
		InitValue: "json",
	},
}
