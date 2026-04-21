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

package chmodhelpertests

import (
	"github.com/alimtvnetwork/core-v8/coretests"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var (
	dirFilesWithContentCreateReadTestCases = []coretestcases.CaseV1{
		{
			Title: "DirFilesWithContent - testing any file reading - writing",
			ArrangeInput: []args.OneAny{
				{
					First: pathInstructionsV3,
				},
			},
			ExpectedInput: []string{
				"0 : file-1.txt",
				"         0. some lines",
				"         1. alim",
				"0 : file-2.txt",
				"         0. some lines file - 2",
				"         1. alim",
				"0 : file-3.txt",
				"         0. some lines file - 3",
				"         1. alim",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.OneAny{}),
		},
	}
)
