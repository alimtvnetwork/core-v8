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

func Test_Integers_Len(t *testing.T) {
	for caseIndex, tc := range integersLenTestCases {
		// Arrange
		var integers coredata.Integers
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			if vals, ok := input["values"]; ok {
				integers = coredata.Integers(vals.([]int))
			}
		}

		// Act
		actual := args.Map{
			"length": integers.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Integers_Less(t *testing.T) {
	tc := integersLessTestCases[0]

	// Arrange
	integers := coredata.Integers{5, 3, 8}

	// Act
	actual := args.Map{
		"less10": integers.Less(1, 0),
		"less01": integers.Less(0, 1),
		"less00": integers.Less(0, 0),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Integers_Sort(t *testing.T) {
	for caseIndex, tc := range integersSortTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		src := input["values"].([]int)
		integers := make(coredata.Integers, len(src))
		copy(integers, src)

		// Act
		sort.Sort(integers)

		actual := args.Map{
			"first": integers[0],
			"last":  integers[len(integers)-1],
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
