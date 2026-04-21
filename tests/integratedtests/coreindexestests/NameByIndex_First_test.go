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

	"github.com/alimtvnetwork/core-v8/coreindexes"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ============================================================================
// NameByIndex
// ============================================================================

func Test_NameByIndex_First(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.NameByIndex(0)}

	// Assert
	expected := args.Map{"result": "First"}
	expected.ShouldBeEqual(t, 0, "NameByIndex returns First -- index 0", actual)
}

func Test_NameByIndex_Tenth(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.NameByIndex(9)}

	// Assert
	expected := args.Map{"result": "Tenth"}
	expected.ShouldBeEqual(t, 0, "NameByIndex returns Tenth -- index 9", actual)
}

func Test_NameByIndex_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.NameByIndex(99)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NameByIndex returns empty -- out of range", actual)
}

// ============================================================================
// Of
// ============================================================================

func Test_Of_Found(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.Of([]int{10, 20, 30}, 20)}

	// Assert
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "Of returns index -- found", actual)
}

func Test_Of_NotFound(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.Of([]int{10, 20, 30}, 99)}

	// Assert
	expected := args.Map{"result": -1}
	expected.ShouldBeEqual(t, 0, "Of returns -1 -- not found", actual)
}

func Test_Of_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.Of([]int{}, 1)}

	// Assert
	expected := args.Map{"result": -1}
	expected.ShouldBeEqual(t, 0, "Of returns -1 -- empty slice", actual)
}

// ============================================================================
// SafeEndingIndex
// ============================================================================

func Test_SafeEndingIndex_WithinBounds(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.SafeEndingIndex(10, 5)}

	// Assert
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "SafeEndingIndex returns requested -- within bounds", actual)
}

func Test_SafeEndingIndex_BeyondBounds(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.SafeEndingIndex(5, 10)}

	// Assert
	expected := args.Map{"result": 4}
	expected.ShouldBeEqual(t, 0, "SafeEndingIndex returns lastIndex -- beyond bounds", actual)
}

// ============================================================================
// IsInvalidIndex
// ============================================================================

func Test_IsInvalidIndex_Negative(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsInvalidIndex(-1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsInvalidIndex returns true -- -1", actual)
}

func Test_IsInvalidIndex_Zero(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsInvalidIndex(0)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsInvalidIndex returns false -- 0", actual)
}

func Test_IsInvalidIndex_Positive(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsInvalidIndex(5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsInvalidIndex returns false -- 5", actual)
}

// ============================================================================
// HasIndexPlusRemoveIndex
// ============================================================================

func Test_HasIndexPlusRemoveIndex_Found(t *testing.T) {
	// Arrange
	indexes := []int{1, 2, 3}
	found := coreindexes.HasIndexPlusRemoveIndex(&indexes, 2)

	// Act
	actual := args.Map{
		"found": found,
		"len": len(indexes),
	}

	// Assert
	expected := args.Map{
		"found": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "HasIndexPlusRemoveIndex removes and returns true -- found", actual)
}

func Test_HasIndexPlusRemoveIndex_NotFound(t *testing.T) {
	// Arrange
	indexes := []int{1, 2, 3}
	found := coreindexes.HasIndexPlusRemoveIndex(&indexes, 99)

	// Act
	actual := args.Map{
		"found": found,
		"len": len(indexes),
	}

	// Assert
	expected := args.Map{
		"found": false,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "HasIndexPlusRemoveIndex returns false -- not found", actual)
}
