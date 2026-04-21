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
	"errors"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// ErrorOnce -- Core (Value + state queries)
// =============================================================================

type errorOnceTestCase struct {
	Case      coretestcases.CaseV1
	InitError string // empty means nil error
}

var errorOnceCoreTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns hasError true and message 'fail' -- error 'fail'",
			ExpectedInput: args.Map{
				"hasError":   true,
				"isValid":    false,
				"isSuccess":  false,
				"isEmpty":    false,
				"isInvalid":  true,
				"isFailed":   true,
				"hasAnyItem": true,
				"isDefined":  true,
				"message":    "fail",
			},
		},
		InitError: "fail",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isValid true and message empty -- nil error",
			ExpectedInput: args.Map{
				"hasError":   false,
				"isValid":    true,
				"isSuccess":  true,
				"isEmpty":    true,
				"isInvalid":  false,
				"isFailed":   false,
				"hasAnyItem": false,
				"isDefined":  false,
				"message":    "",
			},
		},
		InitError: "",
	},
}

// =============================================================================
// ErrorOnce -- Caching
// =============================================================================

var errorOnceCachingTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":        "fail",
				"r2":        "fail",
				"r3":        "fail",
				"callCount": 1,
			},
		},
		InitError: "fail",
	},
}

// =============================================================================
// ErrorOnce -- IsNullOrEmpty
// =============================================================================

var errorOnceNullOrEmptyTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isNullOrEmpty true -- nil error",
			ExpectedInput: args.Map{
				"isNullOrEmpty": true,
			},
		},
		InitError: "",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isNullOrEmpty true -- empty string error",
			ExpectedInput: args.Map{
				"isNullOrEmpty": true,
			},
		},
		InitError: "empty-marker", // special: will create errors.New("")
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce returns isNullOrEmpty false -- error 'msg'",
			ExpectedInput: args.Map{
				"isNullOrEmpty": false,
			},
		},
		InitError: "msg",
	},
}

// =============================================================================
// ErrorOnce -- IsMessageEqual
// =============================================================================

type errorOnceMessageEqualTestCase struct {
	Case      coretestcases.CaseV1
	InitError string
	MatchMsg  string
}

var errorOnceMessageEqualTestCases = []errorOnceMessageEqualTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.IsMessageEqual returns true for 'match' and false for 'other' -- error 'match'",
			ExpectedInput: args.Map{
				"isMessageEqualMatch": true,
				"isMessageEqualOther": false,
			},
		},
		InitError: "match",
		MatchMsg:  "match",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.IsMessageEqual returns false for all -- nil error",
			ExpectedInput: args.Map{
				"isMessageEqualMatch": false,
				"isMessageEqualOther": false,
			},
		},
		InitError: "",
		MatchMsg:  "anything",
	},
}

// =============================================================================
// ErrorOnce -- ConcatNew
// =============================================================================

type errorOnceConcatTestCase struct {
	Case      coretestcases.CaseV1
	InitError string
	ExtraMsg  string
}

var errorOnceConcatTestCases = []errorOnceConcatTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.ConcatNewString returns string containing 'base' and 'extra' -- error 'base'",
			ExpectedInput: args.Map{
				"containsBase":  true,
				"containsExtra": true,
			},
		},
		InitError: "base",
		ExtraMsg:  "extra",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.ConcatNewString returns only additional message -- nil error",
			ExpectedInput: args.Map{
				"result": "\"only\"",
			},
		},
		InitError: "",
		ExtraMsg:  "only",
	},
}

// =============================================================================
// ErrorOnce -- JSON
// =============================================================================

var errorOnceJsonTestCases = []errorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.MarshalJSON returns '\"marshal\"' -- error 'marshal'",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "\"marshal\"",
			},
		},
		InitError: "marshal",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "ErrorOnce.MarshalJSON returns '\"\"' -- nil error",
			ExpectedInput: args.Map{
				"noError":        true,
				"marshaledValue": "\"\"",
			},
		},
		InitError: "",
	},
}

// unused import guard
var _ = errors.New
