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

package enumimpltests

import (
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var dynamicMapSimpleDiffCaseV1TestCases = []coretestcases.CaseV1{
	{
		Title: "Dynamic map simple diff [someKey2] mismatch verify",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey3": 3,
			},
			"right": map[string]any{
				"someKey":  1,
				"someKey2": 4,
				"someKey3": 3,
			},
			"checker": enumimpl.DefaultDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map simple diff [someKey2] mismatch verify",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Left Map - Has Diff from Right Map:",
			"",
			"  {",
			"  ",
			"    \"someKey2\":4",
			"  ",
			"  }",
			"",
			"- Right Map - Has Diff from Left Map:",
			"",
			"  {",
			"  ",
			"    \"someKey2\":2",
			"  ",
			"  }}",
		},
	},
	{
		Title: "Dynamic map simple diff [someKey2], [someKey4] mismatch verify",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey3": 3,
			},
			"right": map[string]any{
				"someKey":  1,
				"someKey4": 4,
				"someKey3": 3,
			},
			"checker": enumimpl.DefaultDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map simple diff [someKey2], [someKey4] mismatch verify",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Left Map - Has Diff from Right Map:",
			"",
			"  {",
			"  ",
			"    \"someKey4\":4",
			"  ",
			"  }",
			"",
			"- Right Map - Has Diff from Left Map:",
			"",
			"  {",
			"  ",
			"    \"someKey2\":2",
			"  ",
			"  }}",
		},
	},
	{
		Title: "Dynamic map simple diff all match - no diff",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey4": 4,
				"someKey3": 3,
			},
			"right": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey4": 4,
				"someKey3": 3,
			},
			"checker": enumimpl.DefaultDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"",
		},
	},
}
