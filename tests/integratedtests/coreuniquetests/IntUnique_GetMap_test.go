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

package coreuniquetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreunique/intunique"
)

// ============================================================================
// GetMap with duplicates
// ============================================================================

func Test_IntUnique_GetMap_WithDuplicates_Ext(t *testing.T) {
	tc := extGetMapWithDuplicatesTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)
	clone := make([]int, len(slice))
	copy(clone, slice)

	// Act
	result := intunique.GetMap(&clone)

	// Assert
	actual := args.Map{
		"isNil":  result == nil,
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// GetMap nil
// ============================================================================

func Test_IntUnique_GetMap_Nil_Ext(t *testing.T) {
	tc := extGetMapNilTestCase

	// Act
	result := intunique.GetMap(nil)

	// Assert
	actual := args.Map{
		"isNil": result == nil,
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// GetMap empty
// ============================================================================

func Test_IntUnique_GetMap_Empty_Ext(t *testing.T) {
	tc := extGetMapEmptyTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)

	// Act
	result := intunique.GetMap(&slice)

	// Assert
	actual := args.Map{
		"isNil":  result == nil,
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// Get empty slice
// ============================================================================

func Test_IntUnique_Get_EmptySlice_Ext(t *testing.T) {
	tc := extGetEmptySliceTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)

	// Act
	result := intunique.Get(&slice)

	// Assert
	actual := args.Map{
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// Get single element
// ============================================================================

func Test_IntUnique_Get_SingleElement_Ext(t *testing.T) {
	tc := extGetSingleElementTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)

	// Act
	result := intunique.Get(&slice)

	// Assert
	actual := args.Map{
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}
