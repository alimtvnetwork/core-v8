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
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

type rwxWrapperManyApplyTestCase struct {
	Case      coretestcases.CaseV1
	SingleRwx chmodhelper.SingleRwx
}

var rwxWrapperManyApplyTestCases = []rwxWrapperManyApplyTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Apply r-x on Other class",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "r-x",
			ClassType: chmodclasstype.Other,
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Apply --- on Other class",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "---",
			ClassType: chmodclasstype.Other,
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Apply --x on Other class",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "--x",
			ClassType: chmodclasstype.Other,
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Apply r-x on Other class (duplicate verify)",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "r-x",
			ClassType: chmodclasstype.Other,
		},
	},
}
