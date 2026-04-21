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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var dynamicMapDiffMessageCaseV1TestCases = []coretestcases.CaseV1{
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
		},
		ExpectedInput: []string{
			"Dynamic map diff string compiled must be same",
			"",
			"Difference Between Map:",
			"",
			"{{",
			"",
			"  \"not-exist-in-left\":2,",
			"  \"not-exist-in-right\":3,",
			"  \"exist-in-left-right-diff-val\":5",
			"",
			"}}",
		},
	},
}
