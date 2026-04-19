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
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/errcore"
)

type rwxCompileValueTestCase struct {
	Case                      coretestcases.CaseV1
	Existing, Input, Expected chmodins.RwxOwnerGroupOther
}

// ShouldBeEqual asserts actLines match expectedLines using
// the embedded Case.Title, with optional context lines for diagnostics.
func (it rwxCompileValueTestCase) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	actLines []string,
	expectedLines []string,
	contextLines ...string,
) {
	t.Helper()

	errcore.AssertDiffOnMismatch(
		t,
		caseIndex,
		it.Case.Title,
		actLines,
		expectedLines,
		contextLines...,
	)
}

var rwxCompileValueTestCases = []rwxCompileValueTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Existing [rwx,r-x,r--] Applied by [*-x,**x,-w-] should result [r-x,r-x,-w-]",
		},
		Existing: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Input: chmodins.RwxOwnerGroupOther{
			Owner: "*-x",
			Group: "**x",
			Other: "-w-",
		},
		Expected: chmodins.RwxOwnerGroupOther{
			Owner: "r-x",
			Group: "r-x",
			Other: "-w-",
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Existing [rwx,r--,--x] Applied by [***,**x,-w*] should result [rwx,r-x,-wx]",
		},
		Existing: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r--",
			Other: "--x",
		},
		Input: chmodins.RwxOwnerGroupOther{
			Owner: "***",
			Group: "**x",
			Other: "-w*",
		},
		Expected: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "-wx",
		},
	},
}
