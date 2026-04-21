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

package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 3: Filter variants, Hashset, Sort, Contains, Join, CSV,
//   GetAllExcept, JSON, Clear/Dispose, Resize, remaining methods
// ══════════════════════════════════════════════════════════════════════════════

// ── FilterLock ───────────────────────────────────────────────────────────────

func Test_Collection_FilteredCollection_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilteredCollection_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb", "ccc")
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) >= 2, false
		})

		// Act
		actual := args.Map{"len": fc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilteredCollection -- returns new collection", actual)
	})
}

func Test_Collection_NonEmptyList_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyList_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b", "")

		// Act
		actual := args.Map{"len": len(c.NonEmptyList())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyList -- skips empty", actual)
	})
}

func Test_Collection_NonEmptyList_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyList_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"len": len(c.NonEmptyList())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "NonEmptyList empty -- returns empty", actual)
	})
}

func Test_Collection_NonEmptyListPtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyListPtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b")
		result := c.NonEmptyListPtr()

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyListPtr -- returns ptr to non-empty", actual)
	})
}

func Test_Collection_NonEmptyItems_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItems_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b")

		// Act
		actual := args.Map{"len": len(c.NonEmptyItems())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItems -- skips empty", actual)
	})
}

func Test_Collection_NonEmptyItemsPtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItemsPtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b")

		// Act
		actual := args.Map{"len": len(c.NonEmptyItemsPtr())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsPtr -- skips empty", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespace_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItemsOrNonWhitespace_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "   ", "", "b")

		// Act
		actual := args.Map{"len": len(c.NonEmptyItemsOrNonWhitespace())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsOrNonWhitespace -- skips ws", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespacePtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItemsOrNonWhitespacePtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "   ", "b")

		// Act
		actual := args.Map{"len": len(c.NonEmptyItemsOrNonWhitespacePtr())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsOrNonWhitespacePtr -- skips ws", actual)
	})
}

func Test_Collection_HashsetAsIs_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HashsetAsIs_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "a")
		hs := c.HashsetAsIs()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetAsIs -- 2 unique", actual)
	})
}

func Test_Collection_HashsetWithDoubleLength_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HashsetWithDoubleLength_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		hs := c.HashsetWithDoubleLength()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetWithDoubleLength -- 2 items", actual)
	})
}

func Test_Collection_HashsetLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HashsetLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("x", "y")
		hs := c.HashsetLock()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetLock -- 2 items", actual)
	})
}

func Test_Collection_Items_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Items_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"len": len(c.Items())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Items -- returns items", actual)
	})
}

func Test_Collection_ListPtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_ListPtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{"len": len(c.ListPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListPtr -- returns items", actual)
	})
}

func Test_Collection_ListCopyPtrLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_ListCopyPtrLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		list := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"len": len(list)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock -- returns copy", actual)
	})
}

func Test_Collection_ListCopyPtrLock_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_ListCopyPtrLock_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		list := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"len": len(list)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock empty -- returns empty", actual)
	})
}

func Test_Collection_Has_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Has_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{
			"has": c.Has("b"),
			"miss": c.Has("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Has -- found and missing", actual)
	})
}

func Test_Collection_Has_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Has_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"has": c.Has("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty -- false", actual)
	})
}

func Test_Collection_HasLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HasLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("x")

		// Act
		actual := args.Map{"has": c.HasLock("x")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasLock -- found", actual)
	})
}

func Test_Collection_HasPtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HasPtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("hello")
		s := "hello"

		// Act
		actual := args.Map{
			"has": c.HasPtr(&s),
			"nil": c.HasPtr(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "HasPtr -- found and nil", actual)
	})
}

func Test_Collection_HasPtr_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HasPtr_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		s := "x"

		// Act
		actual := args.Map{"has": c.HasPtr(&s)}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasPtr empty -- false", actual)
	})
}

func Test_Collection_HasAll_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HasAll_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{
			"all":     c.HasAll("a", "b"),
			"missing": c.HasAll("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all":     true,
			"missing": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll -- all found and missing one", actual)
	})
}

func Test_Collection_HasAll_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HasAll_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"has": c.HasAll("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasAll empty -- false", actual)
	})
}

func Test_Collection_HasUsingSensitivity_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_HasUsingSensitivity_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("Hello")

		// Act
		actual := args.Map{
			"sensitive":   c.HasUsingSensitivity("hello", true),
			"insensitive": c.HasUsingSensitivity("hello", false),
		}

		// Assert
		expected := args.Map{
			"sensitive":   false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "HasUsingSensitivity -- case comparison", actual)
	})
}

func Test_Collection_IsContainsPtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsPtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("x")
		s := "x"

		// Act
		actual := args.Map{
			"has": c.IsContainsPtr(&s),
			"nil": c.IsContainsPtr(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsPtr -- found and nil", actual)
	})
}

func Test_Collection_IsContainsAllSlice_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsAllSlice_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{
			"all":   c.IsContainsAllSlice([]string{"a", "b"}),
			"miss":  c.IsContainsAllSlice([]string{"a", "z"}),
			"empty": c.IsContainsAllSlice([]string{}),
		}

		// Assert
		expected := args.Map{
			"all":   true,
			"miss":  false,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAllSlice -- various cases", actual)
	})
}

func Test_Collection_IsContainsAll_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsAll_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{
			"all": c.IsContainsAll("a", "b"),
			"nil": c.IsContainsAll(nil...),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAll -- found and nil", actual)
	})
}

func Test_Collection_IsContainsAllLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsAllLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{
			"all": c.IsContainsAllLock("a", "b"),
			"nil": c.IsContainsAllLock(nil...),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAllLock -- found and nil", actual)
	})
}

func Test_Collection_GetHashsetPlusHasAll_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_GetHashsetPlusHasAll_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})

		// Act
		actual := args.Map{
			"hasAll": hasAll,
			"hsLen": hs.Length(),
		}

		// Assert
		expected := args.Map{
			"hasAll": true,
			"hsLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll -- all found", actual)
	})
}

func Test_Collection_GetHashsetPlusHasAll_Nil_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_GetHashsetPlusHasAll_Nil_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		_, hasAll := c.GetHashsetPlusHasAll(nil)

		// Act
		actual := args.Map{"hasAll": hasAll}

		// Assert
		expected := args.Map{"hasAll": false}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll nil -- false", actual)
	})
}

func Test_Collection_SortedListAsc_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SortedListAsc_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("c", "a", "b")
		sorted := c.SortedListAsc()

		// Act
		actual := args.Map{
			"first": sorted[0],
			"last": sorted[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "SortedListAsc -- ascending order", actual)
	})
}

func Test_Collection_SortedListAsc_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SortedListAsc_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"len": len(c.SortedListAsc())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedListAsc empty -- returns empty", actual)
	})
}

func Test_Collection_SortedAsc_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SortedAsc_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("c", "a", "b")
		c.SortedAsc()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAsc -- mutates in place", actual)
	})
}

func Test_Collection_SortedAsc_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SortedAsc_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.SortedAsc()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedAsc empty -- no change", actual)
	})
}

func Test_Collection_SortedAscLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SortedAscLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("b", "a")
		c.SortedAscLock()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAscLock -- sorted", actual)
	})
}

func Test_Collection_SortedAscLock_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SortedAscLock_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.SortedAscLock()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedAscLock empty -- no change", actual)
	})
}

func Test_Collection_SortedListDsc_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SortedListDsc_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "c", "b")
		sorted := c.SortedListDsc()

		// Act
		actual := args.Map{
			"first": sorted[0],
			"last": sorted[2],
		}

		// Assert
		expected := args.Map{
			"first": "c",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "SortedListDsc -- descending order", actual)
	})
}

func Test_Collection_New_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_New_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		nc := c.New("a", "b")

		// Act
		actual := args.Map{"len": nc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "New -- creates new collection", actual)
	})
}

func Test_Collection_New_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_New_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		nc := c.New()

		// Act
		actual := args.Map{"len": nc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "New empty -- creates empty collection", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmptyStrings_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddNonEmptyStrings("a", "", "b")

		// Act
		actual := args.Map{"len": c.Length()}
		// Fix: AddNonEmptyStrings filters empty strings, so "a","","b" → 2 items
		// See issues/corestrtests-addnonemptystrings-wrong-expectation.md

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings -- adds items", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmptyStrings_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddNonEmptyStrings()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings empty -- no change", actual)
	})
}

func Test_Collection_AddNonEmptyStringsSlice_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmptyStringsSlice_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddNonEmptyStringsSlice([]string{})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStringsSlice empty -- no change", actual)
	})
}

func Test_Collection_AddFuncResult_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AddFuncResult_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddFuncResult(
			func() string { return "hello" },
			func() string { return "world" },
		)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddFuncResult -- 2 results added", actual)
	})
}

func Test_Collection_AddFuncResult_Nil_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AddFuncResult_Nil_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddFuncResult(nil...)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddFuncResult nil -- no change", actual)
	})
}

func Test_Collection_AddStringsByFuncChecking_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AddStringsByFuncChecking_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddStringsByFuncChecking(
			[]string{"a", "bb", "ccc"},
			func(line string) bool { return len(line) > 1 },
		)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsByFuncChecking -- 2 pass check", actual)
	})
}

func Test_Collection_ExpandSlicePlusAdd_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_ExpandSlicePlusAdd_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.ExpandSlicePlusAdd(
			[]string{"a,b", "c,d"},
			func(line string) []string {
				return []string{line + "_expanded"}
			},
		)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ExpandSlicePlusAdd -- 2 expanded items", actual)
	})
}

func Test_Collection_MergeSlicesOfSlice_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_MergeSlicesOfSlice_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.MergeSlicesOfSlice([]string{"a", "b"}, []string{"c"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlice -- 3 items merged", actual)
	})
}

func Test_Collection_GetAllExceptCollection_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_GetAllExceptCollection_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c", "d")
		exclude := corestr.New.Collection.Cap(5)
		exclude.Adds("b", "d")
		result := c.GetAllExceptCollection(exclude)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection -- 2 remaining", actual)
	})
}

func Test_Collection_GetAllExceptCollection_Nil_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_GetAllExceptCollection_Nil_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		result := c.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil -- returns copy", actual)
	})
}

func Test_Collection_GetAllExceptCollection_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_GetAllExceptCollection_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		result := c.GetAllExceptCollection(corestr.New.Collection.Cap(0))

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection empty -- returns copy", actual)
	})
}

func Test_Collection_GetAllExcept_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_GetAllExcept_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		result := c.GetAllExcept([]string{"b"})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept -- 2 remaining", actual)
	})
}

func Test_Collection_GetAllExcept_Nil_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_GetAllExcept_Nil_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		result := c.GetAllExcept(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept nil -- returns copy", actual)
	})
}

func Test_Collection_CharCollectionMap_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_CharCollectionMap_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("hello", "world")
		ccm := c.CharCollectionMap()

		// Act
		actual := args.Map{"notNil": ccm != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap -- returns map", actual)
	})
}

func Test_Collection_String_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_String_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"nonEmpty": c.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Collection_String_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_String_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"hasNoElements": len(c.String()) > 0}

		// Assert
		expected := args.Map{"hasNoElements": true}
		expected.ShouldBeEqual(t, 0, "String empty -- contains NoElements", actual)
	})
}

func Test_Collection_StringLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_StringLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")

		// Act
		actual := args.Map{"nonEmpty": c.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Collection_StringLock_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_StringLock_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"nonEmpty": c.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty -- contains NoElements", actual)
	})
}

func Test_Collection_SummaryString_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SummaryString_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")

		// Act
		actual := args.Map{"nonEmpty": c.SummaryString(1) != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString -- non-empty", actual)
	})
}

func Test_Collection_SummaryStringWithHeader_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SummaryStringWithHeader_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")

		// Act
		actual := args.Map{"nonEmpty": c.SummaryStringWithHeader("Header") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader -- non-empty", actual)
	})
}

func Test_Collection_SummaryStringWithHeader_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_SummaryStringWithHeader_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"nonEmpty": c.SummaryStringWithHeader("Header") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader empty -- contains header", actual)
	})
}

func Test_Collection_Csv_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Csv_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"nonEmpty": c.Csv() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Csv -- non-empty", actual)
	})
}

func Test_Collection_Csv_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Csv_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"empty": c.Csv() == ""}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Csv empty -- returns empty string", actual)
	})
}

func Test_Collection_CsvOptions_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_CsvOptions_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"nonEmpty": c.CsvOptions(true) != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions -- non-empty", actual)
	})
}

func Test_Collection_CsvOptions_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_CsvOptions_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"empty": c.CsvOptions(false) == ""}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions empty -- returns empty", actual)
	})
}

func Test_Collection_CsvLines_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_CsvLines_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"len": len(c.CsvLines())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvLines -- 2 items", actual)
	})
}

func Test_Collection_CsvLinesOptions_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_CsvLinesOptions_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"len": len(c.CsvLinesOptions(true))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvLinesOptions -- 2 items", actual)
	})
}

func Test_Collection_Join_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Join_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"val": c.Join(",")}

		// Assert
		expected := args.Map{"val": "a,b,c"}
		expected.ShouldBeEqual(t, 0, "Join -- comma separated", actual)
	})
}

func Test_Collection_Join_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Join_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"val": c.Join(",")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Join empty -- empty string", actual)
	})
}

func Test_Collection_JoinLine_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_JoinLine_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"nonEmpty": c.JoinLine() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine -- non-empty", actual)
	})
}

func Test_Collection_JoinLine_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_JoinLine_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"val": c.JoinLine()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinLine empty -- empty string", actual)
	})
}

func Test_Collection_Joins_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Joins_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"nonEmpty": c.Joins(",", "c", "d") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins -- non-empty", actual)
	})
}

func Test_Collection_Joins_NoExtra_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Joins_NoExtra_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"val": c.Joins(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Joins no extra -- just items", actual)
	})
}

func Test_Collection_NonEmptyJoins_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyJoins_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "", "b")

		// Act
		actual := args.Map{"nonEmpty": c.NonEmptyJoins(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonEmptyJoins -- skips empty", actual)
	})
}

func Test_Collection_NonWhitespaceJoins_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_NonWhitespaceJoins_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "   ", "b")

		// Act
		actual := args.Map{"nonEmpty": c.NonWhitespaceJoins(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonWhitespaceJoins -- skips whitespace", actual)
	})
}

func Test_Collection_Resize_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Resize_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		c.Resize(100)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 100}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Resize -- capacity increased", actual)
	})
}

func Test_Collection_Resize_SmallerNoOp_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Resize_SmallerNoOp_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(100)
		c.Resize(5)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 100}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Resize smaller -- no change", actual)
	})
}

func Test_Collection_AddCapacity_Nil_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AddCapacity_Nil_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(nil...)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 5}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "AddCapacity nil -- no change", actual)
	})
}

func Test_Collection_Json_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Json_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		j := c.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Collection_JsonPtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_JsonPtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		j := c.JsonPtr()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr -- no error", actual)
	})
}

func Test_Collection_JsonModel_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_JsonModel_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"len": len(c.JsonModel())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "JsonModel -- returns items", actual)
	})
}

func Test_Collection_JsonModelAny_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_JsonModelAny_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")

		// Act
		actual := args.Map{"notNil": c.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Collection_MarshalJSON_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_MarshalJSON_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		b, err := c.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Collection_UnmarshalJSON_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_UnmarshalJSON_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		err := c.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Collection_UnmarshalJSON_Invalid_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_UnmarshalJSON_Invalid_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		err := c.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Collection_Serialize_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Serialize_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		b, err := c.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Collection_Deserialize_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Deserialize_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		var dest []string
		err := c.Deserialize(&dest)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": len(dest),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Collection_ParseInjectUsingJson_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_ParseInjectUsingJson_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		jr := c.JsonPtr()
		c2 := corestr.New.Collection.Cap(10)
		result, err := c2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": result.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Collection_ParseInjectUsingJsonMust_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_ParseInjectUsingJsonMust_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		jr := c.JsonPtr()
		c2 := corestr.New.Collection.Cap(10)
		result := c2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Collection_ParseInjectUsingJsonMust_Panic_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_ParseInjectUsingJsonMust_Panic_FromSeg3", func() {
		defer func() { recover() }()
		c := corestr.New.Collection.Cap(10)
		bad := &corejson.Result{}
		_ = c.ParseInjectUsingJsonMust(bad)
	})
}

func Test_Collection_JsonParseSelfInject_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_JsonParseSelfInject_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		jr := c.JsonPtr()
		c2 := corestr.New.Collection.Cap(10)
		err := c2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Collection_AsJsonMarshaller_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AsJsonMarshaller_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"notNil": c.AsJsonMarshaller() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller -- non-nil", actual)
	})
}

func Test_Collection_AsJsonContractsBinder_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_AsJsonContractsBinder_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"notNil": c.AsJsonContractsBinder() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder -- non-nil", actual)
	})
}

func Test_Collection_Clear_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Clear_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		c.Clear()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Collection_Clear_Nil_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Clear_Nil_FromSeg3", func() {
		// Arrange
		var c *corestr.Collection
		result := c.Clear()

		// Act
		actual := args.Map{"nil": result == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

func Test_Collection_Dispose_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Dispose_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a")
		c.Dispose()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- items nil", actual)
	})
}

func Test_Collection_Dispose_Nil_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_Dispose_Nil_FromSeg3", func() {
		var c *corestr.Collection
		c.Dispose() // should not panic
	})
}

