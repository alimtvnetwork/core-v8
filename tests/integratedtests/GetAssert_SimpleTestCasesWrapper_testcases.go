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

package integratedtests

import (
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var (
	stringTestCases = []coretestcases.CaseV1{
		{
			Title: "dff text but length is okay",
			ArrangeInput: coretestcases.CaseV1{
				Title: "giving string - output split to lines by newlines",
				ArrangeInput: args.Map{
					"any": "some string contains\nnewline\nin between",
				},
				ActualInput: "some actual text v1! - length okay but diff text",
				ExpectedInput: []string{
					"diff text but not in length!",
				},
				VerifyTypeOf: commonType,
			},
			ExpectedInput: []string{
				"----------------------",
				"0 )  Title:\"giving string - output split to lines by newlines\"",
				"      Input:`args.Map{\"any\":\"some string contains\\nnewline\\nin between\"}` ,",
				"",
				"  Actual:",
				"  `\"some actual text v1! - length okay but diff text\"` ,",
				"",
				"Expected:",
				"  `[]string{\"diff text but not in length!\"}`",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf(coretestcases.CaseV1{}),
		},
		{
			Title: "empty slice both cases alright.",
			ArrangeInput: coretestcases.CaseV1{
				Title: "giving empty empty slice",
				ArrangeInput: args.Map{
					"any": []string{},
				},
				ActualInput:   "some actual text v2! - empty slice",
				ExpectedInput: []string{},
				VerifyTypeOf:  commonType,
			},
			ExpectedInput: []string{
				"----------------------",
				"1 )  Title:\"giving empty empty slice\"",
				"      Input:`args.Map{\"any\":[]string{}}` ,",
				"",
				"  Actual:",
				"  `\"some actual text v2! - empty slice\"` ,",
				"",
				"Expected:",
				"  `[]string{}`",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf(coretestcases.CaseV1{}),
		},
	}

	linesTestCases = []coretestcases.CaseV1{
		{
			Title: "lines method verify output.",
			ArrangeInput: coretestcases.CaseV1{
				Title: "giving slice with diff length to verify",
				ArrangeInput: args.Map{
					"1": "line 1",
					"2": "line 2",
					"3": "line 3",
				},
				ActualInput: []string{
					"arrange inputs 1",
					"arrange inputs 2",
					"arrange inputs 3",
					"arrange inputs 4",
				},
				ExpectedInput: []string{
					"diff text but not in length! L1",
					"diff text but not in length! L2",
				},
				VerifyTypeOf: commonType,
			},
			ExpectedInput: []string{
				"Title: giving slice with diff length to verify",
				"Arrange Lines:",
				"    arrange inputs 1",
				"    arrange inputs 2",
				"    arrange inputs 3",
				"    arrange inputs 4",
				"Expected Lines:",
				"    diff text but not in length! L1",
				"    diff text but not in length! L2",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf(coretestcases.CaseV1{}),
		},
		{
			Title: "empty slice.",
			ArrangeInput: coretestcases.CaseV1{
				Title: "giving empty empty slice",
				ArrangeInput: args.Map{
					"any": []string{},
				},
				ActualInput:   "some actual text v2! - empty slice",
				ExpectedInput: []string{},
				VerifyTypeOf:  commonType,
			},
			ExpectedInput: []string{
				"Title: giving empty empty slice",
				"Arrange Lines:",
				"    some actual text v2! - empty slice",
				"Expected Lines:",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf(coretestcases.CaseV1{}),
		},
	}
)
