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

package corevalidatortests

import (
	"reflect"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/tests/testwrappers/corevalidatortestwrappers"
)

var (
	arrangeArgsTwoTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.TwoAny{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	sliceValidatorTestCases = []corevalidatortestwrappers.SliceValidatorWrapper{
		{
			Case: coretestcases.CaseV1{
				Title: "Diff check against invalid comparisons, it will contain all the diff Index 0 - 2",
				ArrangeInput: []args.TwoAny{
					{
						First:  1,
						Second: byte(2),
					},
					{
						First:  1,
						Second: float64(5),
					},
					{
						First:  "1",
						Second: 1,
					},
				},
				ExpectedInput: []string{
					"Wrong expectation 1",
					"Wrong expectation 2",
					"Wrong expectation 3",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
			Validator: corevalidator.SliceValidator{
				Condition: corevalidator.DefaultTrimCoreCondition,
				ExpectedLines: []string{
					"",
					"=== Line-by-Line Diff (Case 0: Diff check against invalid comparisons, it will contain all the diff Index 0 - 2) ===",
					"    Actual lines: 3, Expected lines: 3",
					"  Line   0 [MISMATCH]:",
					"              actual : `0 : false (1, 2)`",
					"            expected : `Wrong expectation 1`",
					"  Line   1 [MISMATCH]:",
					"              actual : `1 : false (1, 5)`",
					"            expected : `Wrong expectation 2`",
					"  Line   2 [MISMATCH]:",
					"              actual : `2 : false (\"1\", 1)`",
					"            expected : `Wrong expectation 3`",
					"=== Total: 3 lines, 3 mismatches ===",
					"",
					"",
					"============================>",
					"0 ) Actual Received:",
					"    Diff check against invalid comparisons, it will contain all the diff Index 0 - 2",
					"============================>",
					"\"0 : false (1, 2)\",",
					"\"1 : false (1, 5)\",",
					"\"2 : false (\\\"1\\\", 1)\",",
					"============================>",
					"",
					"",
					"============================>",
					"0 )  Expected Input:",
					"     Diff check against invalid comparisons, it will contain all the diff Index 0 - 2",
					"============================>",
					"\"Wrong expectation 1\",",
					"\"Wrong expectation 2\",",
					"\"Wrong expectation 3\",",
					"============================>",
				},
			},
		},
	}

	sliceValidatorFirstErrorTestCases = []corevalidatortestwrappers.SliceValidatorWrapper{
		{
			Case: coretestcases.CaseV1{
				Title: "Diff check against invalid comparisons, it will only contain the first diff Index 0.",
				ArrangeInput: []args.TwoAny{
					{
						First:  1,
						Second: byte(2),
					},
					{
						First:  1,
						Second: float64(5),
					},
					{
						First:  "1",
						Second: 1,
					},
				},
				ExpectedInput: []string{
					"Wrong expectation 1",
					"Wrong expectation 2",
					"Wrong expectation 3",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
			Validator: corevalidator.SliceValidator{
				Condition: corevalidator.DefaultTrimCoreCondition,
				ExpectedLines: []string{
					"",
					"=== Line-by-Line Diff (Case 0: Diff check against invalid comparisons, it will only contain the first diff Index 0.) ===",
					"    Actual lines: 3, Expected lines: 3",
					"  Line   0 [MISMATCH]:",
					"              actual : `0 : false (1, 2)`",
					"            expected : `Wrong expectation 1`",
					"  Line   1 [MISMATCH]:",
					"              actual : `1 : false (1, 5)`",
					"            expected : `Wrong expectation 2`",
					"  Line   2 [MISMATCH]:",
					"              actual : `2 : false (\"1\", 1)`",
					"            expected : `Wrong expectation 3`",
					"=== Total: 3 lines, 3 mismatches ===",
					"",
					"",
					"============================>",
					"0 ) Actual Received:",
					"    Diff check against invalid comparisons, it will only contain the first diff Index 0.",
					"============================>",
					"\"0 : false (1, 2)\",",
					"\"1 : false (1, 5)\",",
					"\"2 : false (\\\"1\\\", 1)\",",
					"============================>",
					"",
					"",
					"============================>",
					"0 )  Expected Input:",
					"     Diff check against invalid comparisons, it will only contain the first diff Index 0.",
					"============================>",
					"\"Wrong expectation 1\",",
					"\"Wrong expectation 2\",",
					"\"Wrong expectation 3\",",
					"============================>",
				},
			},
		},
	}
)
