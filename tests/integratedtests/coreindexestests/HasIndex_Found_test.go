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

// ═══════════════════════════════════════════
// HasIndex
// ═══════════════════════════════════════════

func Test_HasIndex_Found(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.HasIndex([]int{1, 2, 3}, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasIndex returns true -- found", actual)
}

func Test_HasIndex_NotFound(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.HasIndex([]int{1, 2, 3}, 99)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasIndex returns false -- not found", actual)
}

func Test_HasIndex_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.HasIndex([]int{}, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasIndex returns false -- empty slice", actual)
}

// ═══════════════════════════════════════════
// IsWithinIndexRange
// ═══════════════════════════════════════════

func Test_IsWithinIndexRange_Within(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(2, 5)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsWithinIndexRange returns true -- within", actual)
}

func Test_IsWithinIndexRange_Exact(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(4, 5)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsWithinIndexRange returns true -- exact last index", actual)
}

func Test_IsWithinIndexRange_Beyond(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(5, 5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsWithinIndexRange returns false -- beyond", actual)
}

// ═══════════════════════════════════════════
// LastIndex
// ═══════════════════════════════════════════

func Test_LastIndex_Normal(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.LastIndex(5)}

	// Assert
	expected := args.Map{"result": 4}
	expected.ShouldBeEqual(t, 0, "LastIndex returns 4 -- length 5", actual)
}

func Test_LastIndex_One(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.LastIndex(1)}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "LastIndex returns 0 -- length 1", actual)
}

// ═══════════════════════════════════════════
// NameByIndex — remaining indexes
// ═══════════════════════════════════════════

func Test_NameByIndex_Second(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.NameByIndex(1)}

	// Assert
	expected := args.Map{"result": "Second"}
	expected.ShouldBeEqual(t, 0, "NameByIndex returns Second -- index 1", actual)
}

func Test_NameByIndex_Fifth(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.NameByIndex(4)}

	// Assert
	expected := args.Map{"result": "Fifth"}
	expected.ShouldBeEqual(t, 0, "NameByIndex returns Fifth -- index 4", actual)
}
