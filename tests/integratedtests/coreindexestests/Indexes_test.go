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

package coreindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreindexes"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_HasIndex_Verification(t *testing.T) {
	for caseIndex, tc := range hasIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		indexesVal, _ := input.Get("indexes")
		indexes := indexesVal.([]int)
		current, _ := input.GetAsInt("current")

		// Act
		result := coreindexes.HasIndex(indexes, current)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LastIndex_Verification(t *testing.T) {
	for caseIndex, tc := range lastIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		length, _ := input.GetAsInt("length")

		// Act
		result := coreindexes.LastIndex(length)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsWithinIndexRange_Verification(t *testing.T) {
	for caseIndex, tc := range isWithinIndexRangeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		index, _ := input.GetAsInt("index")
		length, _ := input.GetAsInt("length")

		// Act
		result := coreindexes.IsWithinIndexRange(index, length)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
