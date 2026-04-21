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

func Test_PointerStringsDsc_Len(t *testing.T) {
	for caseIndex, tc := range pointerStringsDscLenTestCases {
		// Arrange
		var ps coredata.PointerStringsDsc
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			count, _ := input.GetAsInt("count")
			ptrs := make([]*string, count)
			for i := range ptrs {
				v := "item"
				ptrs[i] = &v
			}
			ps = coredata.PointerStringsDsc(ptrs)
		}

		// Act
		actual := args.Map{
			"length": ps.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PointerStringsDsc_Less(t *testing.T) {
	// Arrange
	// Case 0: both non-nil — descending "alpha" vs "beta"
	{
		a, b := "alpha", "beta"
		ps := coredata.PointerStringsDsc{&a, &b}

	// Act
		actual := args.Map{
			"lessAB": ps.Less(0, 1),
			"lessBA": ps.Less(1, 0),
		}

	// Assert
		pointerStringsDscLessTestCases[0].ShouldBeEqualMap(t, 0, actual)
	}

	// Case 1: nil-i returns false
	{
		b := "beta"
		ps := coredata.PointerStringsDsc{nil, &b}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsDscLessTestCases[1].ShouldBeEqualMap(t, 1, actual)
	}

	// Case 2: nil-j returns true
	{
		a := "alpha"
		ps := coredata.PointerStringsDsc{&a, nil}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsDscLessTestCases[2].ShouldBeEqualMap(t, 2, actual)
	}

	// Case 3: both nil returns false
	{
		ps := coredata.PointerStringsDsc{nil, nil}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsDscLessTestCases[3].ShouldBeEqualMap(t, 3, actual)
	}
}

func Test_PointerStringsDsc_Sort(t *testing.T) {
	tc := pointerStringsDscSortTestCases[0]

	// Arrange
	a, b, c := "alpha", "charlie", "beta"
	ps := coredata.PointerStringsDsc{&a, nil, &b, &c}

	// Act
	sort.Sort(ps)

	actual := args.Map{
		"first":     *ps[0],
		"lastIsNil": ps[len(ps)-1] == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
