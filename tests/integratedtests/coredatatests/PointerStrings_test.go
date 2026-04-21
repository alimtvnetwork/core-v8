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

func Test_PointerStrings_Len(t *testing.T) {
	for caseIndex, tc := range pointerStringsLenTestCases {
		// Arrange
		var ps coredata.PointerStrings
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			count, _ := input.GetAsInt("count")
			ptrs := make([]*string, count)
			for i := range ptrs {
				v := "item"
				ptrs[i] = &v
			}
			ps = coredata.PointerStrings(ptrs)
		}

		// Act
		actual := args.Map{
			"length": ps.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PointerStrings_Less(t *testing.T) {
	// Arrange
	// Case 0: both non-nil
	{
		a, b := "alpha", "beta"
		ps := coredata.PointerStrings{&a, &b}

	// Act
		actual := args.Map{
			"less01": ps.Less(0, 1),
			"less10": ps.Less(1, 0),
		}

	// Assert
		pointerStringsLessTestCases[0].ShouldBeEqualMap(t, 0, actual)
	}

	// Case 1: nil first
	{
		b := "beta"
		ps := coredata.PointerStrings{nil, &b}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsLessTestCases[1].ShouldBeEqualMap(t, 1, actual)
	}

	// Case 2: nil second
	{
		a := "alpha"
		ps := coredata.PointerStrings{&a, nil}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsLessTestCases[2].ShouldBeEqualMap(t, 2, actual)
	}

	// Case 3: both nil
	{
		ps := coredata.PointerStrings{nil, nil}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsLessTestCases[3].ShouldBeEqualMap(t, 3, actual)
	}
}

func Test_PointerStrings_Sort(t *testing.T) {
	tc := pointerStringsSortTestCases[0]

	// Arrange
	c, a, b := "charlie", "alpha", "beta"
	ps := coredata.PointerStrings{&c, nil, &a, &b}

	// Act
	sort.Sort(ps)

	actual := args.Map{
		"firstIsNil": ps[0] == nil,
		"second":     *ps[1],
		"third":      *ps[2],
		"fourth":     *ps[3],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
