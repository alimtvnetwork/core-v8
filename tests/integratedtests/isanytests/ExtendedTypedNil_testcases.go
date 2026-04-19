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

package isanytests

import (
	"errors"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var (
	nilError   error
	nilIntPtr  *int
	nilIntPtr2 *int
	liveError  = errors.New("")
	liveErrorX = errors.New("x")

	// -------------------------------------------------------------------------
	// nullTesting01 — Defined and Null on typed nils (error, *int)
	// -------------------------------------------------------------------------

	extendedDefinedTestCases = []coretestcases.CaseV1{
		{
			Title: "Defined on typed-nil error and *int",
			ArrangeInput: args.Map{
				"when":   "nil literal, live error, typed-nil error, typed-nil *int",
				"inputs": []any{nil, liveError, nilError, nilIntPtr},
			},
			ExpectedInput: args.Map{
				"result0": "false",
				"type0":   "<nil>",
				"result1": "true",
				"type1":   "*errors.errorString",
				"result2": "false",
				"type2":   "<nil>",
				"result3": "false",
				"type3":   "*int",
			},
		},
	}

	extendedNullTestCases = []coretestcases.CaseV1{
		{
			Title: "Null on typed-nil error and *int",
			ArrangeInput: args.Map{
				"when":   "nil literal, live error, typed-nil error, typed-nil *int",
				"inputs": []any{nil, liveError, nilError, nilIntPtr},
			},
			ExpectedInput: args.Map{
				"result0": "true",
				"type0":   "<nil>",
				"result1": "false",
				"type1":   "*errors.errorString",
				"result2": "true",
				"type2":   "<nil>",
				"result3": "true",
				"type3":   "*int",
			},
		},
	}

	// -------------------------------------------------------------------------
	// nullTesting02 — DefinedBoth and NullBoth with error and *int typed nils
	// -------------------------------------------------------------------------

	extendedDefinedBothTestCases = []coretestcases.CaseV1{
		{
			Title: "DefinedBoth with typed-nil error, *int, and live error",
			ArrangeInput: args.Map{
				"when": "migrated from nullTesting02",
				"pairs": []args.TwoAny{
					{First: nil, Second: liveErrorX},
					{First: nil, Second: nil},
					{First: nilIntPtr, Second: nilIntPtr2},
					{First: liveErrorX, Second: liveErrorX},
					{First: liveErrorX, Second: nilIntPtr},
				},
			},
			ExpectedInput: args.Map{
				"result0": "false",
				"types0":  "<nil>, *errors.errorString",
				"result1": "false",
				"types1":  "<nil>, <nil>",
				"result2": "false",
				"types2":  "*int, *int",
				"result3": "true",
				"types3":  "*errors.errorString, *errors.errorString",
				"result4": "false",
				"types4":  "*errors.errorString, *int",
			},
		},
	}

	extendedNullBothTestCases = []coretestcases.CaseV1{
		{
			Title: "NullBoth with typed-nil error, *int, and live error",
			ArrangeInput: args.Map{
				"when": "migrated from nullTesting02",
				"pairs": []args.TwoAny{
					{First: nil, Second: liveErrorX},
					{First: nil, Second: nil},
					{First: nilIntPtr, Second: nilIntPtr2},
					{First: liveErrorX, Second: liveErrorX},
					{First: liveErrorX, Second: nilIntPtr},
				},
			},
			ExpectedInput: args.Map{
				"result0": "false",
				"types0":  "<nil>, *errors.errorString",
				"result1": "true",
				"types1":  "<nil>, <nil>",
				"result2": "true",
				"types2":  "*int, *int",
				"result3": "false",
				"types3":  "*errors.errorString, *errors.errorString",
				"result4": "false",
				"types4":  "*errors.errorString, *int",
			},
		},
	}
)
