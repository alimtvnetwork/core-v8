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

package coredatatests

import (
	"reflect"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var (
	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.ThreeFuncAny{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	funWrapCreationTestCases = []coretestcases.CaseV1{
		{
			Title: "someFunctionV1 => Calls dynamically with valid params, outputs as it should.",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				Third:    "f3",
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"someFunctionV1 => called with (f1, f2, f3) - some new stuff",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "someFunctionV1 => Calls dynamically with less param (null), outputs error args count mismatch.",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				Third:    nil,
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"error : ",
				"  someFunctionV1 [Func] =>",
				"    arguments count doesn't match for - count:",
				"      expected : 3",
				"      given    : 2",
				"    expected types listed :",
				"      - string",
				"      - string",
				"      - string",
				"    actual given types list :",
				"      - 0. string [value: f1]",
				"      - 1. string [value: f2]",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "someFunctionV1 => Calls dynamically with mismatch datatype for arg 2nd, it expects string but given int, outputs error",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   1,
				Third:    "f3",
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"error : ",
				"  someFunctionV1 =>",
				"      - Index {1}, 2nd args : Expected Type (string) != (int) Given Type",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving nil as a work func, doesn't panic but returns error.",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   1,
				Third:    "f3",
				WorkFunc: nil,
			},
			ExpectedInput: []string{
				"error : ",
				"  func-wrap is invalid:",
				"      given type: <nil>",
				"      name: ",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving (int) as a work func, doesn't panic but returns error.",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   1,
				Third:    "f3",
				WorkFunc: 1,
			},
			ExpectedInput: []string{
				"error : ",
				"  func-wrap is invalid:",
				"      given type: int",
				"      name: ",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "someFunctionV2 => Calls dynamically with valid params, outputs as it should.",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				WorkFunc: someFunctionV2,
			},
			ExpectedInput: []string{
				"someFunctionV2 => called with (f1, f2) - (string, error)",
				"some err v2",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "someFunctionV2 => Calls dynamically with valid params, outputs as it should.",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				WorkFunc: someFunctionV2,
			},
			ExpectedInput: []string{
				"someFunctionV2 => called with (f1, f2) - (string, error)",
				"some err v2",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "someFunctionV3 => Calls dynamically with valid params, outputs as it should.",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				WorkFunc: someFunctionV3,
			},
			ExpectedInput: []string{
				"5",
				"someFunctionV3 => called with (f1, f2) - (int, string, error)",
				"some err of v3",
			},
			VerifyTypeOf: commonType,
		},
	}
)
