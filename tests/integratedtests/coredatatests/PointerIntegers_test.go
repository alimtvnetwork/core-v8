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

	"github.com/alimtvnetwork/core/coredata"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_PointerIntegers_Len(t *testing.T) {
	for caseIndex, tc := range pointerIntegersLenTestCases {
		// Arrange
		var pi coredata.PointerIntegers
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			count, _ := input.GetAsInt("count")
			ptrs := make([]*int, count)
			for i := range ptrs {
				v := i + 1
				ptrs[i] = &v
			}
			pi = coredata.PointerIntegers(ptrs)
		}

		// Act
		actual := args.Map{
			"length": pi.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PointerIntegers_Less(t *testing.T) {
	// Arrange
	// Case 0: both non-nil
	{
		a, b := 3, 8
		pi := coredata.PointerIntegers{&a, &b}

	// Act
		actual := args.Map{
			"lessIJ": pi.Less(0, 1),
			"lessJI": pi.Less(1, 0),
		}

	// Assert
		pointerIntegersLessTestCases[0].ShouldBeEqualMap(t, 0, actual)
	}

	// Case 1: nil-i
	{
		b := 5
		pi := coredata.PointerIntegers{nil, &b}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersLessTestCases[1].ShouldBeEqualMap(t, 1, actual)
	}

	// Case 2: nil-j
	{
		a := 5
		pi := coredata.PointerIntegers{&a, nil}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersLessTestCases[2].ShouldBeEqualMap(t, 2, actual)
	}

	// Case 3: both nil
	{
		pi := coredata.PointerIntegers{nil, nil}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersLessTestCases[3].ShouldBeEqualMap(t, 3, actual)
	}
}

func Test_PointerIntegers_Sort(t *testing.T) {
	tc := pointerIntegersSortTestCases[0]

	// Arrange
	a, b, c := 5, 1, 3
	pi := coredata.PointerIntegers{&a, nil, &b, &c}

	// Act
	sort.Sort(pi)

	actual := args.Map{
		"firstIsNil": pi[0] == nil,
		"second":     *pi[1],
		"last":       *pi[len(pi)-1],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
