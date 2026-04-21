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

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coreunique/intunique"
)

// ═══════════════════════════════════════════
// GetMap
// ═══════════════════════════════════════════

func Test_GetMap_NilSlice(t *testing.T) {
	// Arrange
	result := intunique.GetMap(nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "GetMap returns nil -- nil input", actual)
}

func Test_GetMap_EmptySlice(t *testing.T) {
	// Arrange
	input := []int{}
	result := intunique.GetMap(&input)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"len": len(*result),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "GetMap returns empty map -- empty slice", actual)
}

func Test_GetMap_WithDuplicates(t *testing.T) {
	// Arrange
	input := []int{1, 2, 2, 3, 3, 3}
	result := intunique.GetMap(&input)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"len":    len(*result),
		"has1":   (*result)[1],
		"has2":   (*result)[2],
		"has3":   (*result)[3],
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len":    3,
		"has1":   true,
		"has2":   true,
		"has3":   true,
	}
	expected.ShouldBeEqual(t, 0, "GetMap returns unique map -- with duplicates", actual)
}

// ═══════════════════════════════════════════
// Get — additional branch coverage
// ═══════════════════════════════════════════

func Test_Get_NilSlice(t *testing.T) {
	// Arrange
	result := intunique.Get(nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Get returns nil -- nil input", actual)
}

func Test_Get_SingleElement(t *testing.T) {
	// Arrange
	input := []int{42}
	result := intunique.Get(&input)

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": 42,
	}
	expected.ShouldBeEqual(t, 0, "Get returns same -- single element", actual)
}
