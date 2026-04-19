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

var dynamicMapDiffMessageV2CaseV1TestCases = []coretestcases.CaseV1{
	{
		Title: "Dynamic map diff string compiled must be same",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"exist":                        1,
				"not-exist-in-right":           3,
				"exist-in-left-right-diff-val": 5,
			},
			"right": map[string]any{
				"exist":                        1,
				"not-exist-in-left":            2,
				"exist-in-left-right-diff-val": 6,
			},
			"checker": enumimpl.LeftRightDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map diff string compiled must be same",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Left Map - Has Diff from Right Map:",
			"",
			"  {",
			"  ",
			"    \"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",",
			"    \"not-exist-in-left\":\"{\"Left\":null,\"Right\":2}\"",
			"  ",
			"  }",
			"",
			"- Right Map - Has Diff from Left Map:",
			"",
			"  {",
			"  ",
			"    \"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",",
			"    \"not-exist-in-right\":\"3 (type:int) - left - key is missing!\"",
			"  ",
			"  }}",
		},
	},
	{
		Title: "Dynamic map diff - no changes",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"a":  1,
				"b":  3,
				"cl": 5,
			},
			"right": map[string]any{
				"a":  1,
				"b":  3,
				"cl": 5,
			},
			"checker": enumimpl.LeftRightDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"",
		},
	},
	{
		Title: "Dynamic map diff - right hand key missing - cl int 5",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"a":  1,
				"b":  3,
				"cl": 5,
			},
			"right": map[string]any{
				"a": 1,
				"b": 3,
			},
			"checker": enumimpl.LeftRightDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map diff - right hand key missing - cl int 5",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Right Map - Has Diff from Left Map:",
			"",
			"  {",
			"  ",
			"    \"cl\":\"5 (type:int) - left - key is missing!\"",
			"  ",
			"  }}",
		},
	},
	{
		Title: "Dynamic map diff - left hand key missing - cl {left, right}",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"a": 1,
				"b": 3,
			},
			"right": map[string]any{
				"a":  1,
				"b":  3,
				"cl": 5,
			},
			"checker": enumimpl.LeftRightDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map diff - left hand key missing - cl {left, right}",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Left Map - Has Diff from Right Map:",
			"",
			"  {",
			"  ",
			"    \"cl\":\"{\"Left\":null,\"Right\":5}\"",
			"  ",
			"  }}",
		},
	},
	{
		Title: "Dynamic map diff - left cl - key missing",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"a": 1,
				"b": 3,
			},
			"right": map[string]any{
				"a":  1,
				"b":  3,
				"cl": 5,
			},
			"checker": enumimpl.LeftRightDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map diff - left cl - key missing",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Left Map - Has Diff from Right Map:",
			"",
			"  {",
			"  ",
			"    \"cl\":\"{\"Left\":null,\"Right\":5}\"",
			"  ",
			"  }}",
		},
	},
}
