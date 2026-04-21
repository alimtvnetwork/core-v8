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

	"github.com/alimtvnetwork/core-v8/coredata/coregeneric"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_QW_Collection_Length_NilItems(t *testing.T) {
	// Arrange
	var c *coregeneric.Collection[string]

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for nil collection", actual)
}

func Test_QW_LinkedList_IndexAt_EndOfList(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()
	ll.Add("a")
	// Access index beyond list length — covers the out-of-range early return
	node := ll.IndexAt(5)

	// Act
	actual := args.Map{"result": node != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for out-of-range index", actual)
}

func Test_QW_MinOf_SecondSmaller(t *testing.T) {
	// Arrange
	// Cover the else branch (a >= b)
	result := coregeneric.MinOf(5, 3)

	// Act
	actual := args.Map{"result": result != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_QW_MaxOf_SecondLarger(t *testing.T) {
	// Arrange
	// Cover the else branch (a <= b)
	result := coregeneric.MaxOf(3, 5)

	// Act
	actual := args.Map{"result": result != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_QW_MinOfSlice_NonMinElements(t *testing.T) {
	// Arrange
	// Cover the case where v < result is false
	result := coregeneric.MinOfSlice([]int{3, 5, 1, 4})

	// Act
	actual := args.Map{"result": result != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}
