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

package coreflecttests

import (
	"reflect"

	"github.com/alimtvnetwork/core-v8/coretests"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var (
	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.ThreeFuncAny{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	funWrapCreationTestCases = []coretestcases.CaseV1{
		{
			Title: "FuncWrap returns correct output -- someFunctionV1 with valid params",
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
			Title: "FuncWrap returns args count mismatch error -- someFunctionV1 with nil third param",
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
			Title: "FuncWrap returns type mismatch error -- someFunctionV1 with int instead of string at arg 2",
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
			Title: "FuncWrap returns invalid error -- nil work func given",
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
			Title: "FuncWrap returns invalid error -- int given as work func",
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
			Title: "FuncWrap returns string and error output -- someFunctionV2 with valid params",
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
			Title: "FuncWrap returns int, string, error output -- someFunctionV3 with valid params",
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
