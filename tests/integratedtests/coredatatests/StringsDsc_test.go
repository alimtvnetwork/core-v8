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
	"sort"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_StringsDsc_Len(t *testing.T) {
	for caseIndex, tc := range stringsDscLenTestCases {
		// Arrange
		var strings coredata.StringsDsc
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			if vals, ok := input["values"]; ok {
				strings = coredata.StringsDsc(vals.([]string))
			}
		}

		// Act
		actual := args.Map{
			"length": strings.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsDsc_Sort(t *testing.T) {
	for caseIndex, tc := range stringsDscSortTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		src := input["values"].([]string)
		strs := make(coredata.StringsDsc, len(src))
		copy(strs, src)

		// Act
		sort.Sort(strs)

		actual := args.Map{
			"first": strs[0],
			"last":  strs[len(strs)-1],
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
