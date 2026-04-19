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

package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList.Values() — multi-element iteration
// Covers LinkedListIter.go L36 (current = current.next)
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedList_Values_MultiElement(t *testing.T) {
	// Arrange
	tc := cov5ValuesIterMultiElementTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})

	// Act
	collected := make([]int, 0, 3)
	for v := range ll.Values() {
		collected = append(collected, v)
	}

	actual := args.Map{
		"count":  len(collected),
		"first":  collected[0],
		"second": collected[1],
		"third":  collected[2],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LinkedList_Values_BreakEarly_FromLinkedListValuesIter(t *testing.T) {
	// Arrange
	tc := cov5ValuesIterBreakEarlyTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})

	// Act
	collected := make([]int, 0, 1)
	for v := range ll.Values() {
		collected = append(collected, v)
		break
	}

	actual := args.Map{
		"count": len(collected),
		"first": collected[0],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
