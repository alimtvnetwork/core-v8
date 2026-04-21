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

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Segment 5a
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_IsEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEmpty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{
			"empty": h.IsEmpty(),
			"hasItems": h.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Hashmap_AddOrUpdate_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdate", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.AddOrUpdate("k", "v")
		isNew2 := h.AddOrUpdate("k", "v2")

		// Act
		actual := args.Map{
			"isNew": isNew,
			"isNew2": isNew2,
			"len": h.Length(),
		}

		// Assert
		expected := args.Map{
			"isNew": true,
			"isNew2": false,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "AddOrUpdate -- new then update", actual)
	})
}

func Test_Hashmap_Set_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Set", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.Set("a", "1")

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "Set -- new key", actual)
	})
}

func Test_Hashmap_SetTrim_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SetTrim", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.SetTrim(" a ", " 1 ")
		val, found := h.Get("a")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "1",
		}
		expected.ShouldBeEqual(t, 0, "SetTrim -- trimmed key and val", actual)
	})
}

func Test_Hashmap_SetBySplitter_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SetBySplitter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.SetBySplitter("=", "key=value")
		val, found := h.Get("key")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "value",
		}
		expected.ShouldBeEqual(t, 0, "SetBySplitter -- split key=value", actual)
	})
}

func Test_Hashmap_SetBySplitter_NoValue_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SetBySplitter_NoValue", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.SetBySplitter("=", "keyonly")
		val, found := h.Get("keyonly")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "SetBySplitter no value -- empty val", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValInt_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValInt("n", 42)
		val, _ := h.Get("n")

		// Act
		actual := args.Map{"val": val}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValInt -- int to string", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat("f", 1.5)
		_, found := h.Get("f")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValFloat -- stored", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat64_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat64("f", 2.5)
		_, found := h.Get("f")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValFloat64 -- stored", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValAny("a", 123)
		_, found := h.Get("a")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValAny -- stored", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValueAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyValueAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 99})
		_, found := h.Get("k")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValueAny -- stored", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyVal_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyVal", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyVal -- new", actual)
	})
}

func Test_Hashmap_AddOrUpdateWithWgLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateWithWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("k", "v", &wg)
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateWithWgLock -- added", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateHashmap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("b", "2")
		h.AddOrUpdateHashmap(h2)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateHashmap -- merged", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.AddOrUpdateHashmap(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateHashmap nil -- no change", actual)
	})
}

func Test_Hashmap_AddOrUpdateMap_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateMap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap -- 2 items", actual)
	})
}

func Test_Hashmap_AddOrUpdateMap_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateMap_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap empty -- no change", actual)
	})
}

func Test_Hashmap_AddsOrUpdates_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdates", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdates -- 2 items", actual)
	})
}

func Test_Hashmap_AddsOrUpdates_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdates_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdates(nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdates nil -- no change", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: 1})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyAnyValues -- 1 item", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyAnyValues_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyAnyValues()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyAnyValues empty -- no change", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValues -- 1 item", actual)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyValues_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValues()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValues empty -- no change", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection -- 2 items", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_NilKeys_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateCollection_NilKeys", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateCollection(nil, nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection nil -- no change", actual)
	})
}

func Test_Hashmap_AddOrUpdateCollection_LenMismatch_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateCollection_LenMismatch", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection mismatch -- no change", actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(&wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateStringsPtrWgLock -- 2 items", actual)
	})
}

func Test_Hashmap_AddOrUpdateStringsPtrWgLock_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateStringsPtrWgLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(&wg, []string{}, []string{})
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateStringsPtrWgLock empty -- no change", actual)
	})
}

// ── Has / Contains / Missing ────────────────────────────────────────────────

func Test_Hashmap_Has_Contains_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Has_Contains", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"has":      h.Has("a"),
			"miss":     h.Has("z"),
			"contains": h.Contains("a"),
			"missing":  h.IsKeyMissing("z"),
		}

		// Assert
		expected := args.Map{
			"has":      true,
			"miss":     false,
			"contains": true,
			"missing":  true,
		}
		expected.ShouldBeEqual(t, 0, "Has/Contains/Missing -- correct", actual)
	})
}

func Test_Hashmap_HasLock_ContainsLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasLock_ContainsLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"hasLock":      h.HasLock("a"),
			"containsLock": h.ContainsLock("a"),
			"missingLock":  h.IsKeyMissingLock("z"),
			"hasWithLock":  h.HasWithLock("a"),
		}

		// Assert
		expected := args.Map{
			"hasLock":      true,
			"containsLock": true,
			"missingLock":  true,
			"hasWithLock":  true,
		}
		expected.ShouldBeEqual(t, 0, "HasLock/ContainsLock -- correct", actual)
	})
}

func Test_Hashmap_HasAllStrings_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Set("b", "2")

		// Act
		actual := args.Map{
			"all":  h.HasAllStrings("a", "b"),
			"miss": h.HasAllStrings("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllStrings -- all and missing", actual)
	})
}

func Test_Hashmap_HasAll_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAll", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"has": h.HasAll("a"),
			"miss": h.HasAll("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll -- found and missing", actual)
	})
}

func Test_Hashmap_HasAllCollectionItems_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has": h.HasAllCollectionItems(c),
			"nil": h.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllCollectionItems -- found and nil", actual)
	})
}

func Test_Hashmap_HasAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"any": h.HasAny("z", "a"),
			"none": h.HasAny("x", "y"),
		}

		// Assert
		expected := args.Map{
			"any": true,
			"none": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAny -- found and none", actual)
	})
}

func Test_Hashmap_HasAnyItem_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAnyItem", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"has": h.HasAnyItem()}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem -- true", actual)
	})
}

// ── Get / GetValue ──────────────────────────────────────────────────────────

func Test_Hashmap_Get_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Get", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		val, found := h.Get("a")
		val2, found2 := h.GetValue("a")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
			"val2": val2,
			"found2": found2,
		}

		// Assert
		expected := args.Map{
			"val": "1",
			"found": true,
			"val2": "1",
			"found2": true,
		}
		expected.ShouldBeEqual(t, 0, "Get/GetValue -- found", actual)
	})
}

// ── Items / Keys / Values ───────────────────────────────────────────────────

func Test_Hashmap_Items_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Items", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1 item", actual)
	})
}

func Test_Hashmap_SafeItems_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SafeItems_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"nil": h.SafeItems() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafeItems nil -- nil", actual)
	})
}

func Test_Hashmap_Keys_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Keys", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"len": len(h.Keys()),
			"allLen": len(h.AllKeys()),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"allLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Keys/AllKeys -- 1 key", actual)
	})
}

func Test_Hashmap_KeysCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.KeysCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysCollection -- 1 key", actual)
	})
}

func Test_Hashmap_ValuesList_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.ValuesList())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesList -- 1 value", actual)
	})
}

func Test_Hashmap_ValuesCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesCollection -- 1 value", actual)
	})
}

func Test_Hashmap_ValuesHashset_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesHashset().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesHashset -- 1 value", actual)
	})
}

func Test_Hashmap_KeysValuesList_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		keys, values := h.KeysValuesList()

		// Act
		actual := args.Map{
			"kLen": len(keys),
			"vLen": len(values),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesList -- 1 each", actual)
	})
}

func Test_Hashmap_KeysValuesCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		keys, values := h.KeysValuesCollection()

		// Act
		actual := args.Map{
			"kLen": keys.Length(),
			"vLen": values.Length(),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesCollection -- 1 each", actual)
	})
}

func Test_Hashmap_KeysValuePairs_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuePairs", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.KeysValuePairs())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysValuePairs -- 1 pair", actual)
	})
}

func Test_Hashmap_KeysValuePairsCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuePairsCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.KeysValuePairsCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysValuePairsCollection -- 1 pair", actual)
	})
}

func Test_Hashmap_KeysLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.KeysLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysLock -- 1 key", actual)
	})
}

func Test_Hashmap_KeysLock_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"len": len(h.KeysLock())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "KeysLock empty -- 0", actual)
	})
}

func Test_Hashmap_KeysValuesListLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuesListLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		keys, vals := h.KeysValuesListLock()

		// Act
		actual := args.Map{
			"kLen": len(keys),
			"vLen": len(vals),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesListLock -- 1 each", actual)
	})
}

func Test_Hashmap_ValuesListCopyLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesListCopyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.ValuesListCopyLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesListCopyLock -- 1", actual)
	})
}

func Test_Hashmap_ValuesCollectionLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesCollectionLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesCollectionLock().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesCollectionLock -- 1", actual)
	})
}

func Test_Hashmap_ValuesHashsetLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesHashsetLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesHashsetLock().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesHashsetLock -- 1", actual)
	})
}

func Test_Hashmap_ItemsCopyLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ItemsCopyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		copied := h.ItemsCopyLock()

		// Act
		actual := args.Map{"len": len(*copied)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ItemsCopyLock -- 1", actual)
	})
}

// ── Filter ──────────────────────────────────────────────────────────────────

func Test_Hashmap_GetKeysFilteredItems_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("bb", "2")
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 1, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems -- 1 match", actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredItems_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems empty -- 0", actual)
	})
}

func Test_Hashmap_GetKeysFilteredItems_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredItems_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		h.Set("c", "3")
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"hasItems": len(result) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems break -- stops early", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("bb", "2")
		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 1, false
		})

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection -- 1 match", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredCollection_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection empty -- empty", actual)
	})
}

func Test_Hashmap_GetKeysFilteredCollection_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredCollection_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"hasItems": result.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection break -- stops", actual)
	})
}

// ── AddsOrUpdatesAnyUsingFilter ─────────────────────────────────────────────

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilter(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, false
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter -- 1 kept", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilter(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, true
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
			corestr.KeyAnyValuePair{Key: "b", Value: 2},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter break -- only 1", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesAnyUsingFilter(nil, nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter nil -- no change", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, false
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilterLock -- 1 kept", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilterLock_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, true
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
			corestr.KeyAnyValuePair{Key: "b", Value: 2},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilterLock break -- only 1", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, false
			},
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter -- 1 kept", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesUsingFilter_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesUsingFilter(nil, nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter nil -- no change", actual)
	})
}

func Test_Hashmap_AddsOrUpdatesUsingFilter_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesUsingFilter_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, true
			},
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter break -- only 1", actual)
	})
}

// ── Concat ──────────────────────────────────────────────────────────────────

func Test_Hashmap_ConcatNew_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("b", "2")
		result := h.ConcatNew(true, h2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- merged", actual)
	})
}

func Test_Hashmap_ConcatNew_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNew_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ConcatNew(true)

		// Act
		actual := args.Map{"hasItems": result.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty -- cloned", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNewUsingMaps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingMaps -- merged", actual)
	})
}

func Test_Hashmap_ConcatNewUsingMaps_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNewUsingMaps_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ConcatNewUsingMaps(true)

		// Act
		actual := args.Map{"hasItems": result.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingMaps empty -- cloned", actual)
	})
}

// ── Lock variants ───────────────────────────────────────────────────────────

func Test_Hashmap_AddOrUpdateLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateLock("k", "v")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateLock -- 1 item", actual)
	})
}

func Test_Hashmap_IsEmptyLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEmptyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"empty": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_Hashmap_LengthLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_LengthLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Hashmap_Length_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Length_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

// ── IsEqual ─────────────────────────────────────────────────────────────────

func Test_Hashmap_IsEqualPtr_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "1")
		h3 := corestr.New.Hashmap.Cap(2)
		h3.Set("a", "2")

		// Act
		actual := args.Map{
			"eq":      h1.IsEqualPtr(h2),
			"neq":     h1.IsEqualPtr(h3),
			"same":    h1.IsEqualPtr(h1),
			"nilBoth": (*corestr.Hashmap)(nil).IsEqualPtr(nil),
			"nilOne":  h1.IsEqualPtr(nil),
		}

		// Assert
		expected := args.Map{
			"eq":      true,
			"neq":     false,
			"same":    true,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr -- various", actual)
	})
}

func Test_Hashmap_IsEqualPtr_DiffLen_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr_DiffLen", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "1")
		h2.Set("b", "2")

		// Act
		actual := args.Map{"eq": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr diff len -- false", actual)
	})
}

func Test_Hashmap_IsEqualPtr_BothEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr_BothEmpty", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(0)
		h2 := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"eq": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr both empty -- true", actual)
	})
}

func Test_Hashmap_IsEqualPtr_OneEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr_OneEmpty", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"eq": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr one empty -- false", actual)
	})
}

func Test_Hashmap_IsEqualPtrLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtrLock", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "1")

		// Act
		actual := args.Map{"eq": h1.IsEqualPtrLock(h2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtrLock -- true", actual)
	})
}

// ── Remove ──────────────────────────────────────────────────────────────────

func Test_Hashmap_Remove_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Remove", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Remove("a")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Remove -- removed", actual)
	})
}

func Test_Hashmap_RemoveWithLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_RemoveWithLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.RemoveWithLock("a")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveWithLock -- removed", actual)
	})
}

// ── String / Join ───────────────────────────────────────────────────────────

func Test_Hashmap_String_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_String", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"nonEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Hashmap_String_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_String_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"nonEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- has NoElements text", actual)
	})
}

func Test_Hashmap_StringLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_StringLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"nonEmpty": h.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Hashmap_StringLock_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_StringLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"nonEmpty": h.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty -- has NoElements text", actual)
	})
}

func Test_Hashmap_Join_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Join", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"nonEmpty": h.Join(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Join -- non-empty", actual)
	})
}

func Test_Hashmap_JoinKeys_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JoinKeys", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"val": h.JoinKeys(",")}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "JoinKeys -- a", actual)
	})
}

// ── Except ──────────────────────────────────────────────────────────────────

func Test_Hashmap_GetValuesExceptKeysInHashset_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesExceptKeysInHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		except := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetValuesExceptKeysInHashset(except)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesExceptKeysInHashset -- 1 remaining", actual)
	})
}

func Test_Hashmap_GetValuesExceptKeysInHashset_NilExcept_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesExceptKeysInHashset_NilExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.GetValuesExceptKeysInHashset(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesExceptKeysInHashset nil -- all values", actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesKeysExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		result := h.GetValuesKeysExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesKeysExcept -- 1 remaining", actual)
	})
}

func Test_Hashmap_GetValuesKeysExcept_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesKeysExcept_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.GetValuesKeysExcept(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesKeysExcept nil -- all values", actual)
	})
}

func Test_Hashmap_GetAllExceptCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetAllExceptCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		c := corestr.New.Collection.Strings([]string{"a"})
		result := h.GetAllExceptCollection(c)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection -- 1 remaining", actual)
	})
}

func Test_Hashmap_GetAllExceptCollection_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetAllExceptCollection_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil -- all values", actual)
	})
}

// ── KeysToLower / ValuesToLower ─────────────────────────────────────────────

func Test_Hashmap_KeysToLower_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("ABC", "1")
		result := h.KeysToLower()
		_, found := result.Get("abc")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "KeysToLower -- lowered", actual)
	})
}

func Test_Hashmap_ValuesToLower_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("ABC", "1")
		result := h.ValuesToLower()
		_, found := result.Get("abc")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "ValuesToLower -- delegates to KeysToLower", actual)
	})
}

// ── Diff ────────────────────────────────────────────────────────────────────

func Test_Hashmap_Diff_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Diff", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "2")
		diff := h.Diff(h2)

		// Act
		actual := args.Map{"notNil": diff != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Diff -- non-nil", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Hashmap_Json_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Json", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		j := h.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Hashmap_MarshalJSON_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_MarshalJSON", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		b, err := h.MarshalJSON()

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

func Test_Hashmap_UnmarshalJSON_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_UnmarshalJSON", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		err := h.UnmarshalJSON([]byte(`{"a":"1"}`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": h.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Hashmap_UnmarshalJSON_Invalid_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_UnmarshalJSON_Invalid", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		err := h.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJson_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ParseInjectUsingJson", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(0)
		result, err := h2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": result.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Hashmap_ParseInjectUsingJsonMust_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ParseInjectUsingJsonMust", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(0)
		result := h2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Hashmap_Serialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Serialize", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		b, err := h.Serialize()

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

func Test_Hashmap_Deserialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Deserialize", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		var dest map[string]string
		err := h.Deserialize(&dest)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": len(dest),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Hashmap_JsonModel_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JsonModel", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_Hashmap_JsonModelAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JsonModelAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"notNil": h.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Hashmap_InterfaceCasts_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_InterfaceCasts", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)

		// Act
		actual := args.Map{
			"jsoner":   h.AsJsoner() != nil,
			"binder":   h.AsJsonContractsBinder() != nil,
			"injector": h.AsJsonParseSelfInjector() != nil,
			"marsh":    h.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner":   true,
			"binder":   true,
			"injector": true,
			"marsh":    true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Hashmap_JsonParseSelfInject_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JsonParseSelfInject", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(0)
		err := h2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Error / KeyValStringLines ───────────────────────────────────────────────

func Test_Hashmap_ToError_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToError", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		err := h.ToError("; ")

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ToError -- non-nil", actual)
	})
}

func Test_Hashmap_ToDefaultError_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToDefaultError", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		err := h.ToDefaultError()

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ToDefaultError -- non-nil", actual)
	})
}

func Test_Hashmap_KeyValStringLines_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeyValStringLines", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.KeyValStringLines())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeyValStringLines -- 1 line", actual)
	})
}

func Test_Hashmap_ToStringsUsingCompiler_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToStringsUsingCompiler", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })

		// Act
		actual := args.Map{
			"len": len(result),
			"val": result[0],
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"val": "a=1",
		}
		expected.ShouldBeEqual(t, 0, "ToStringsUsingCompiler -- formatted", actual)
	})
}

func Test_Hashmap_ToStringsUsingCompiler_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToStringsUsingCompiler_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		result := h.ToStringsUsingCompiler(func(k, v string) string { return k })

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ToStringsUsingCompiler empty -- empty", actual)
	})
}

// ── Clone / Clear / Dispose ─────────────────────────────────────────────────

func Test_Hashmap_Clone_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clone", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		c := h.Clone()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Clone -- same items", actual)
	})
}

func Test_Hashmap_Clone_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clone_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		c := h.Clone()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clone empty -- empty", actual)
	})
}

func Test_Hashmap_ClonePtr_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ClonePtr", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		c := h.ClonePtr()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ClonePtr -- same items", actual)
	})
}

func Test_Hashmap_ClonePtr_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ClonePtr_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"nil": h.ClonePtr() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "ClonePtr nil -- returns nil", actual)
	})
}

func Test_Hashmap_Clear_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clear", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Clear()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Hashmap_Clear_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clear_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"nil": h.Clear() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

func Test_Hashmap_Dispose_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Dispose", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Dispose()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_Hashmap_Dispose_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Dispose_Nil", func() {
		var h *corestr.Hashmap
		h.Dispose() // should not panic
	})
}

// ── Collection ──────────────────────────────────────────────────────────────

func Test_Hashmap_Collection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Collection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.Collection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection -- 1 item", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashmap_DataModel_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_DataModel", func() {
		// Arrange
		dm := &corestr.HashmapDataModel{Items: map[string]string{"a": "1"}}
		h := corestr.NewHashmapUsingDataModel(dm)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashmapUsingDataModel -- 1 item", actual)
	})
}

func Test_Hashmap_DataModel_Reverse_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_DataModel_Reverse", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		dm := corestr.NewHashmapsDataModelUsing(h)

		// Act
		actual := args.Map{"len": len(dm.Items)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashmapsDataModelUsing -- 1 item", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff
// ══════════════════════════════════════════════════════════════════════════════

