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
	"sync"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── Collection CRUD ──

func Test_Collection_Add_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Add", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.Add("a")

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddNonEmpty_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddNonEmptyWhitespace_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("a")

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddIf_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddIf(false, "skip")
		c.AddIf(true, "keep")

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddIfMany_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddIfMany", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddIfMany(false, "a", "b")
		c.AddIfMany(true, "c", "d")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_AddError_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddError", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddError(nil)

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		c.AddError(errForTest)
		actual = args.Map{"result": c.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_Adds_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Adds", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": c.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_AddStrings_FromC37CollectionCollect(t *testing.T) {
	safeTest(t, "Test_Collection_AddStrings", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"x", "y"})

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_AddLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddLock("a")

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddsLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddsLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_AddFunc_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddFunc", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddFuncErr_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddFuncErr", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddFuncErr(func() (string, error) { return "", errForTest }, func(e error) {})
		actual = args.Map{"result": c.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 still", actual)
	})
}

func Test_Collection_AddCollection_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddCollection", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c1.AddCollection(c2)

		// Act
		actual := args.Map{"result": c1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c1.AddCollection(corestr.Empty.Collection())
		actual = args.Map{"result": c1.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 still", actual)
	})
}

func Test_Collection_AddCollections_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddCollections", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AddCollections(c1, c2, corestr.Empty.Collection())

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_AddPointerCollectionsLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddPointerCollectionsLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(c1)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddHashmapsValues_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsValues", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("k", "v")
		c.AddHashmapsValues(h)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddHashmapsValues(nil)
		actual = args.Map{"result": c.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 after nil", actual)
	})
}

func Test_Collection_AddHashmapsKeys_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsKeys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("mykey", "myval")
		c.AddHashmapsKeys(h)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddHashmapsKeys(nil)
	})
}

func Test_Collection_AddHashmapsKeysValues_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsKeysValues", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(h)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AddHashmapsKeysValues(nil)
	})
}

// ── Collection query methods ──

func Test_Collection_HasAnyItem_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HasAnyItem", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Collection_HasIndex_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HasIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.HasIndex(1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.HasIndex(2)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": c.HasIndex(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_LastIndex_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_LastIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_Capacity_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Capacity", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"result": c.Capacity() < 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 10", actual)
		empty := &corestr.Collection{}
		_ = empty.Capacity()
	})
}

func Test_Collection_Count_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Count", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_Length_Nil_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Length_Nil", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_LengthLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_LengthLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_IsEmpty_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsEmpty", func() {
		// Act
		actual := args.Map{"result": corestr.Empty.Collection().IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_IsEmptyLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsEmptyLock", func() {
		// Act
		actual := args.Map{"result": corestr.Empty.Collection().IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_HasItems_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HasItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
	})
}

func Test_Collection_RemoveAt_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.RemoveAt(1)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": c.RemoveAt(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": c.RemoveAt(100)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── Collection equality ──

func Test_Collection_IsEquals_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsEquals", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Collection_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Collection_IsEquals_Nil", func() {
		// Arrange
		var a *corestr.Collection
		var b *corestr.Collection

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true nil==nil", actual)
		c := corestr.New.Collection.Strings([]string{"a"})
		actual = args.Map{"result": c.IsEquals(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_IsEquals_Self(t *testing.T) {
	safeTest(t, "Test_Collection_IsEquals_Self", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsEquals(a)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for self", actual)
	})
}

func Test_Collection_IsEquals_BothEmpty_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.Collection()
		b := corestr.Empty.Collection()

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Collection_IsEquals_DiffLen_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsEquals_DiffLen", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_IsEqualsWithSensitive_Insensitive(t *testing.T) {
	safeTest(t, "Test_Collection_IsEqualsWithSensitive_Insensitive", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal insensitive", actual)
		actual = args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal sensitive", actual)
	})
}

func Test_Collection_IsEqualsWithSensitive_Mismatch(t *testing.T) {
	safeTest(t, "Test_Collection_IsEqualsWithSensitive_Mismatch", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"World"})

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

// ── Collection accessors ──

func Test_Collection_First_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_First", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Collection_Last_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Last", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Last() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Collection_FirstOrDefault_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_FirstOrDefault", func() {
		// Act
		actual := args.Map{"result": corestr.Empty.Collection().FirstOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		c := corestr.New.Collection.Strings([]string{"a"})
		actual = args.Map{"result": c.FirstOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Collection_LastOrDefault_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_LastOrDefault", func() {
		// Act
		actual := args.Map{"result": corestr.Empty.Collection().LastOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		c := corestr.New.Collection.Strings([]string{"a"})
		actual = args.Map{"result": c.LastOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Collection_IndexAt_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IndexAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IndexAt(1) != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Collection_SafeIndexAtUsingLength_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.SafeIndexAtUsingLength("def", 1, 0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": c.SafeIndexAtUsingLength("def", 1, 5) != "def"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
	})
}

func Test_Collection_List_Items_ListStrings(t *testing.T) {
	safeTest(t, "Test_Collection_List_Items_ListStrings", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(c.List()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(c.Items()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(c.ListStrings()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(c.ListStringsPtr()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(c.ListPtr()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

// ── Collection transform ──

func Test_Collection_Take_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Take", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		taken := c.Take(2)

		// Act
		actual := args.Map{"result": taken.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		same := c.Take(10)
		actual = args.Map{"result": same != c}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ptr", actual)
		empty := c.Take(0)
		actual = args.Map{"result": empty.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Skip_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		skipped := c.Skip(1)

		// Act
		actual := args.Map{"result": skipped.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		same := c.Skip(0)
		actual = args.Map{"result": same != c}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ptr", actual)
	})
}

func Test_Collection_Reverse_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Reverse", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()

		// Act
		actual := args.Map{"result": c.First() != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
		// Two items
		c2 := corestr.New.Collection.Strings([]string{"x", "y"})
		c2.Reverse()
		actual = args.Map{"result": c2.First() != "y"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected y", actual)
		// Single
		c3 := corestr.New.Collection.Strings([]string{"z"})
		c3.Reverse()
		actual = args.Map{"result": c3.First() != "z"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected z", actual)
	})
}

func Test_Collection_SortedListAsc_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_SortedListAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		sorted := c.SortedListAsc()

		// Act
		actual := args.Map{"result": sorted[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
		actual = args.Map{"result": corestr.Empty.Collection().SortedListAsc() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected nil", actual)
	})
}

func Test_Collection_SortedAsc_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_SortedAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAsc()

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		corestr.Empty.Collection().SortedAsc()
	})
}

func Test_Collection_SortedAscLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_SortedAscLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Collection_SortedListDsc_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_SortedListDsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		dsc := c.SortedListDsc()

		// Act
		actual := args.Map{"result": dsc[0] != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
	})
}

func Test_Collection_UniqueList_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_UniqueList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		u := c.UniqueList()

		// Act
		actual := args.Map{"result": len(u) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_UniqueListLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_UniqueListLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		u := c.UniqueListLock()

		// Act
		actual := args.Map{"result": len(u) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_UniqueBoolMap_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_UniqueBoolMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		m := c.UniqueBoolMap()

		// Act
		actual := args.Map{"result": len(m) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_UniqueBoolMapLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_UniqueBoolMapLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		m := c.UniqueBoolMapLock()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Collection filter ──

func Test_Collection_Filter_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		result2 := corestr.Empty.Collection().Filter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual = args.Map{"result": len(result2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_Collection_Filter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_FilterLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_FilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		result2 := corestr.Empty.Collection().FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual = args.Map{"result": len(result2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_FilteredCollection_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_FilteredCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, s == "a", false
		})

		// Act
		actual := args.Map{"result": fc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_FilteredCollectionLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_FilteredCollectionLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": fc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_FilterPtr_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})

		// Act
		actual := args.Map{"result": len(*result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		result2 := corestr.Empty.Collection().FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		actual = args.Map{"result": len(*result2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_FilterPtrLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})

		// Act
		actual := args.Map{"result": len(*result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		result2 := corestr.Empty.Collection().FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		actual = args.Map{"result": len(*result2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── Collection search ──

func Test_Collection_Has_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Has", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.Has("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": corestr.Empty.Collection().Has("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_HasLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HasLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.HasLock("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Collection_HasPtr_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HasPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"

		// Act
		actual := args.Map{"result": c.HasPtr(&s)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.HasPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": corestr.Empty.Collection().HasPtr(&s)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_HasAll_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.HasAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.HasAll("a", "z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": corestr.Empty.Collection().HasAll("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_HasUsingSensitivity_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HasUsingSensitivity", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"Hello"})

		// Act
		actual := args.Map{"result": c.HasUsingSensitivity("hello", false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.HasUsingSensitivity("hello", true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": c.HasUsingSensitivity("world", false)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_IsContainsPtr_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"

		// Act
		actual := args.Map{"result": c.IsContainsPtr(&s)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.IsContainsPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		miss := "z"
		actual = args.Map{"result": c.IsContainsPtr(&miss)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_IsContainsAll_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IsContainsAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.IsContainsAll(nil...)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_IsContainsAllSlice_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsAllSlice", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IsContainsAllSlice([]string{"a"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.IsContainsAllSlice([]string{})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false on empty", actual)
		actual = args.Map{"result": corestr.Empty.Collection().IsContainsAllSlice([]string{"a"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_IsContainsAllLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_IsContainsAllLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.IsContainsAllLock("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.IsContainsAllLock(nil...)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false nil", actual)
	})
}

func Test_Collection_GetHashsetPlusHasAll_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_GetHashsetPlusHasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, ok := c.GetHashsetPlusHasAll([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ok}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hs.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_, ok2 := c.GetHashsetPlusHasAll(nil)
		actual = args.Map{"result": ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── Collection string ops ──

func Test_Collection_String_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_String", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": corestr.Empty.Collection().String() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have no elements text", actual)
	})
}

func Test_Collection_StringLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_StringLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.StringLock() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_StringJSON_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_StringJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.StringJSON() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_JsonString_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.JsonString() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": c.JsonStringMust() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_Join_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Join", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		actual = args.Map{"result": corestr.Empty.Collection().Join(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_JoinLine_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_JoinLine", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.JoinLine()

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": corestr.Empty.Collection().JoinLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_Joins_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Joins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.Joins(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": c.Joins(",", "b") == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_NonEmptyJoins_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyJoins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyJoins(",")

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_NonWhitespaceJoins_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_NonWhitespaceJoins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		r := c.NonWhitespaceJoins(",")

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

// ── Collection CSV ──

func Test_Collection_Csv_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Csv", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Csv() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.Collection().Csv() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_Collection_CsvOptions", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.CsvOptions(true) == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.Collection().CsvOptions(false) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_CsvLines_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_CsvLines", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		lines := c.CsvLines()

		// Act
		actual := args.Map{"result": len(lines) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_CsvLinesOptions_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_CsvLinesOptions", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		lines := c.CsvLinesOptions(true)

		// Act
		actual := args.Map{"result": len(lines) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Collection pages ──

func Test_Collection_GetPagesSize_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_GetPagesSize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act
		actual := args.Map{"result": c.GetPagesSize(2) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": c.GetPagesSize(0) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": c.GetPagesSize(-1) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_GetPagedCollection_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_GetPagedCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		paged := c.GetPagedCollection(2)

		// Act
		actual := args.Map{"result": paged.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1 page", actual)
	})
}

func Test_Collection_GetSinglePageCollection_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_GetSinglePageCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		page := c.GetSinglePageCollection(2, 1)

		// Act
		actual := args.Map{"result": page.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		page2 := c.GetSinglePageCollection(2, 2)
		actual = args.Map{"result": page2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// small collection
		small := corestr.New.Collection.Strings([]string{"a"})
		same := small.GetSinglePageCollection(10, 1)
		actual = args.Map{"result": same != small}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ptr", actual)
	})
}

// ── Collection hashset/map ──

func Test_Collection_HashsetAsIs_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HashsetAsIs", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_HashsetWithDoubleLength_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.HashsetWithDoubleLength()
	})
}

func Test_Collection_HashsetLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_HashsetLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_CharCollectionMap_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_CharCollectionMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"abc", "axy", "bcd"})
		ccm := c.CharCollectionMap()

		// Act
		actual := args.Map{"result": ccm.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 2 groups", actual)
	})
}

// ── Collection non-empty ──

func Test_Collection_NonEmptyList_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		ne := c.NonEmptyList()

		// Act
		actual := args.Map{"result": len(ne) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": corestr.Empty.Collection().NonEmptyList() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_NonEmptyListPtr_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyListPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", ""})
		ne := c.NonEmptyListPtr()

		// Act
		actual := args.Map{"result": len(*ne) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_NonEmptyItems_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", ""})

		// Act
		actual := args.Map{"result": len(c.NonEmptyItems()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_NonEmptyItemsPtr(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItemsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", ""})
		_ = c.NonEmptyItemsPtr()
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespace_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItemsOrNonWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "  "})
		r := c.NonEmptyItemsOrNonWhitespace()

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespacePtr(t *testing.T) {
	safeTest(t, "Test_Collection_NonEmptyItemsOrNonWhitespacePtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.NonEmptyItemsOrNonWhitespacePtr()
	})
}

// ── Collection add non-empty ──

func Test_Collection_AddsNonEmpty_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddsNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AddsNonEmpty()
	})
}

func Test_Collection_AddsNonEmptyPtrLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddsNonEmptyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := "hello"
		c.AddsNonEmptyPtrLock(&s, nil)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmptyStrings", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "b")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AddNonEmptyStrings()
	})
}

func Test_Collection_AddNonEmptyStringsSlice_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmptyStringsSlice", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a"})

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddNonEmptyStringsSlice(nil)
	})
}

// ── Collection misc ──

func Test_Collection_EachItemSplitBy_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a,b", "c"})
		split := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"result": len(split) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_ConcatNew_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		cn := c.ConcatNew(0, "b")

		// Act
		actual := args.Map{"result": cn.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		cn2 := c.ConcatNew(0)
		actual = args.Map{"result": cn2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_ToError_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ToError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"err1"})

		// Act
		actual := args.Map{"result": c.ToError(",") == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Collection_ToDefaultError_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ToDefaultError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"err1"})

		// Act
		actual := args.Map{"result": c.ToDefaultError() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Collection_AsError_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AsError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"err"})

		// Act
		actual := args.Map{"result": c.AsError(",") == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		actual = args.Map{"result": corestr.Empty.Collection().AsError(",") != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Collection_AsDefaultError_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AsDefaultError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"e"})

		// Act
		actual := args.Map{"result": c.AsDefaultError() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Collection_InsertAt_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_InsertAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		c.InsertAt(0, "b")

		// Act
		actual := args.Map{"result": c.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 3", actual)
	})
}

func Test_Collection_ChainRemoveAt_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ChainRemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_RemoveItemsIndexes_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_RemoveItemsIndexes", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil indexes
		c.RemoveItemsIndexes(true)
	})
}

func Test_Collection_AppendCollectionPtr_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AppendCollectionPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(c2)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_AppendCollections_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AppendCollections", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c.AppendCollections(c1, corestr.Empty.Collection())

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AppendCollections()
	})
}

func Test_Collection_AppendAnys_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AppendAnys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnys(42, "hello", nil)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c.AppendAnys()
	})
}

func Test_Collection_AppendAnysLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AppendAnysLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnysLock(42)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AppendAnysLock()
	})
}

func Test_Collection_AppendNonEmptyAnys_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys(42, nil)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AppendNonEmptyAnys(nil)
	})
}

func Test_Collection_AppendAnysUsingFilter_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AppendAnysUsingFilter", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, false },
			42,
		)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AppendAnysUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, false },
		)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AppendAnysUsingFilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilterLock(
			func(s string, i int) (string, bool, bool) { return s, true, false },
			42,
		)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddsAsync_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddsAsync", func() {
		c := corestr.New.Collection.Cap(10)
		wg := sync.WaitGroup{}
		wg.Add(1)
		c.AddsAsync(&wg, "a", "b")
		wg.Wait()
		// may or may not be 2 due to race, but should not panic
	})
}

func Test_Collection_AddWithWgLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddWithWgLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(&wg, "a")
		wg.Wait()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddStringsByFuncChecking_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddStringsByFuncChecking", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking(
			[]string{"a", "", "b"},
			func(s string) bool { return s != "" },
		)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_ExpandSlicePlusAdd_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ExpandSlicePlusAdd", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd(
			[]string{"a,b", "c"},
			func(s string) []string { return []string{s} },
		)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_MergeSlicesOfSlice_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_MergeSlicesOfSlice", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_GetAllExcept_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_GetAllExcept", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.GetAllExcept([]string{"b"})

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		r2 := c.GetAllExcept(nil)
		actual = args.Map{"result": len(r2) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_GetAllExceptCollection_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_GetAllExceptCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		exc := corestr.New.Collection.Strings([]string{"a"})
		r2 := c.GetAllExceptCollection(exc)
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AddFuncResult_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddFuncResult", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "x" })

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddFuncResult()
	})
}

func Test_Collection_New_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_New", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.New("x", "y")

		// Act
		actual := args.Map{"result": n.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		n2 := c.New()
		actual = args.Map{"result": n2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_ListCopyPtrLock_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ListCopyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		l := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"result": len(l) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		l2 := corestr.Empty.Collection().ListCopyPtrLock()
		actual = args.Map{"result": len(l2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_SummaryString_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_SummaryString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryString(1)

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_SummaryStringWithHeader_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_SummaryStringWithHeader", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.SummaryStringWithHeader("H") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.Collection().SummaryStringWithHeader("H") == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

// ── Collection JSON ──

func Test_Collection_Json_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Json", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		j := c.Json()

		// Act
		actual := args.Map{"hasError": j.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
	})
}

func Test_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Collection_JsonPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		jp := c.JsonPtr()

		// Act
		actual := args.Map{"hasError": jp.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns no error", actual)
	})
}

func Test_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_Collection_JsonModel", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(c.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Collection_JsonModelAny", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Collection_MarshalJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.MarshalJSON()

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": len(b) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_Collection_UnmarshalJSON_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_UnmarshalJSON", func() {
		// Arrange
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": c.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// invalid
		err2 := c.UnmarshalJSON([]byte(`{invalid`))
		actual = args.Map{"result": err2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Collection_ParseInjectUsingJson_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ParseInjectUsingJson", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := &corestr.Collection{}
		result, err := c2.ParseInjectUsingJson(jr)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": result.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_ParseInjectUsingJsonMust_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_ParseInjectUsingJsonMust", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := &corestr.Collection{}
		result := c2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_JsonParseSelfInject_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_JsonParseSelfInject", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := &corestr.Collection{}
		err := c2.JsonParseSelfInject(jr)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Collection_Serialize_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.Serialize()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": len(b) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_Collection_Deserialize_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Deserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		var target []string
		err := c.Deserialize(&target)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

	// ── Collection interface casts ──

func Test_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Collection_AsJsonMarshaller", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

	// ── Collection resize/capacity ──

func Test_Collection_AddCapacity_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(10)
		actual := args.Map{"result": c.Capacity() < 12}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 12", actual)
		c.AddCapacity()
	})
}

func Test_Collection_Resize_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Resize", func() {
		c := corestr.New.Collection.Cap(2)
		c.Resize(100)
		actual := args.Map{"result": c.Capacity() < 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		c.Resize(1) // no-op
	})
}

	// ── Collection clear/dispose ──

func Test_Collection_Clear_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Clear()
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		var nilC *corestr.Collection
		actual = args.Map{"result": nilC.Clear() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Collection_Dispose_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		var nilC *corestr.Collection
		nilC.Dispose() // should not panic
	})
}

	// ── Collection Hashmap filter ──

func Test_Collection_AddHashmapsKeysValuesUsingFilter_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValuesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) {
				return p.Key + "=" + p.Value, true, false
			},
			h,
		)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		c.AddHashmapsKeysValuesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) { return "", true, false },
			nil,
		)
	})
}

	// ── newCollectionCreator ──

func Test_NewCollectionCreator_Methods(t *testing.T) {
	safeTest(t, "Test_NewCollectionCreator_Methods", func() {
		// Cap
		c1 := corestr.New.Collection.Cap(5)
		actual := args.Map{"result": c1.Capacity() < 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// CloneStrings
		c2 := corestr.New.Collection.CloneStrings([]string{"a"})
		actual = args.Map{"result": c2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// Create
		c3 := corestr.New.Collection.Create([]string{"a"})
		actual = args.Map{"result": c3.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// StringsOptions clone
		c4 := corestr.New.Collection.StringsOptions(true, []string{"a"})
		actual = args.Map{"result": c4.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// StringsOptions no clone empty
		c5 := corestr.New.Collection.StringsOptions(false, nil)
		actual = args.Map{"result": c5.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// LineUsingSep
		c6 := corestr.New.Collection.LineUsingSep(",", "a,b,c")
		actual = args.Map{"result": c6.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// LineDefault
		c7 := corestr.New.Collection.LineDefault("a\nb")
		actual = args.Map{"result": c7.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// StringsPlusCap
		c8 := corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		actual = args.Map{"result": c8.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		c9 := corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		actual = args.Map{"result": c9.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// CapStrings
		c10 := corestr.New.Collection.CapStrings(5, []string{"a"})
		actual = args.Map{"result": c10.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		c11 := corestr.New.Collection.CapStrings(0, []string{"a"})
		actual = args.Map{"result": c11.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// LenCap
		c12 := corestr.New.Collection.LenCap(3, 10)
		actual = args.Map{"result": c12.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// Empty
		c13 := corestr.New.Collection.Empty()
		actual = args.Map{"result": c13.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

	// ── emptyCreator ──

func Test_EmptyCreator_All_CollectionDeep(t *testing.T) {
	safeTest(t, "Test_EmptyCreator_All", func() {
		actual := args.Map{"result": corestr.Empty.Collection().Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.LinkedList().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.KeyAnyValuePair() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.KeyValuePair() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.KeyValueCollection().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.LinkedCollections().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.LeftRight() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = corestr.Empty.SimpleStringOnce()
		actual = args.Map{"result": corestr.Empty.SimpleStringOncePtr() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.Hashset().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.HashsetsCollection().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.Hashmap().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.CharCollectionMap().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.KeyValuesCollection().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.CollectionsOfCollection().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.CharHashsetMap().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}
