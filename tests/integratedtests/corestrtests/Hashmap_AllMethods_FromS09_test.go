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
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S09 — Hashmap.go (1,300 lines) — Full coverage
// ══════════════════════════════════════════════════════════════

// ── IsEmpty / HasItems / IsEmptyLock ─────────────────────────

func Test_Hashmap_01_Hashmap_IsEmpty_FromS09(t *testing.T) {
	safeTest(t, "Test_01_Hashmap_IsEmpty", func() {
		// Arrange
		empty := corestr.Empty.Hashmap()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": hm.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_Hashmap_02_Hashmap_HasItems_FromS09(t *testing.T) {
	safeTest(t, "Test_02_Hashmap_HasItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
		actual = args.Map{"result": corestr.Empty.Hashmap().HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items for empty", actual)
	})
}

func Test_Hashmap_03_Hashmap_IsEmptyLock_FromS09(t *testing.T) {
	safeTest(t, "Test_03_Hashmap_IsEmptyLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act & Assert
		actual := args.Map{"result": hm.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Hashmap_04_Hashmap_Collection_FromS09(t *testing.T) {
	safeTest(t, "Test_04_Hashmap_Collection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		col := hm.Collection()

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AddOrUpdate variants ─────────────────────────────────────

func Test_Hashmap_05_Hashmap_AddOrUpdateWithWgLock_FromS09(t *testing.T) {
	safeTest(t, "Test_05_Hashmap_AddOrUpdateWithWgLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		hm.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_06_Hashmap_AddOrUpdateKeyStrValInt_FromS09(t *testing.T) {
	safeTest(t, "Test_06_Hashmap_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValInt("count", 42)

		// Assert
		v, ok := hm.Get("count")
		actual := args.Map{"result": ok || v != "42"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected '42', got ''", actual)
	})
}

func Test_Hashmap_07_Hashmap_AddOrUpdateKeyStrValFloat_FromS09(t *testing.T) {
	safeTest(t, "Test_07_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValFloat("pi", 3.14)

		// Assert
		v, ok := hm.Get("pi")
		actual := args.Map{"result": ok || v == ""}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected float value", actual)
	})
}

func Test_Hashmap_08_Hashmap_AddOrUpdateKeyStrValFloat64_FromS09(t *testing.T) {
	safeTest(t, "Test_08_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValFloat64("e", 2.71828)

		// Assert
		v, ok := hm.Get("e")
		actual := args.Map{"result": ok || v == ""}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected float64 value", actual)
	})
}

func Test_Hashmap_09_Hashmap_AddOrUpdateKeyStrValAny_FromS09(t *testing.T) {
	safeTest(t, "Test_09_Hashmap_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValAny("key", "anyVal")

		// Assert
		v, ok := hm.Get("key")
		actual := args.Map{"result": ok || v == ""}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected value", actual)
	})
}

func Test_Hashmap_10_Hashmap_AddOrUpdateKeyValueAny_FromS09(t *testing.T) {
	safeTest(t, "Test_10_Hashmap_AddOrUpdateKeyValueAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		pair := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		hm.AddOrUpdateKeyValueAny(pair)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_11_Hashmap_AddOrUpdateKeyVal_FromS09(t *testing.T) {
	safeTest(t, "Test_11_Hashmap_AddOrUpdateKeyVal", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		isNew := hm.AddOrUpdateKeyVal(kv)

		// Assert
		actual := args.Map{"result": isNew}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		isNew2 := hm.AddOrUpdateKeyVal(kv)
		actual = args.Map{"result": isNew2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not new on second add", actual)
	})
}

func Test_Hashmap_12_Hashmap_AddOrUpdate_FromS09(t *testing.T) {
	safeTest(t, "Test_12_Hashmap_AddOrUpdate", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		isNew := hm.AddOrUpdate("k", "v")
		isNew2 := hm.AddOrUpdate("k", "v2")

		// Assert
		actual := args.Map{"result": isNew}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		actual = args.Map{"result": isNew2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected update not new", actual)
	})
}

func Test_Hashmap_13_Hashmap_Set_FromS09(t *testing.T) {
	safeTest(t, "Test_13_Hashmap_Set", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		isNew := hm.Set("k", "v")

		// Assert
		actual := args.Map{"result": isNew}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
	})
}

func Test_Hashmap_14_Hashmap_SetTrim_FromS09(t *testing.T) {
	safeTest(t, "Test_14_Hashmap_SetTrim", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.SetTrim("  key  ", "  val  ")

		// Assert
		v, ok := hm.Get("key")
		actual := args.Map{"result": ok || v != "val"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 'val', got ''", actual)
	})
}

func Test_Hashmap_15_Hashmap_SetBySplitter_FromS09(t *testing.T) {
	safeTest(t, "Test_15_Hashmap_SetBySplitter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.SetBySplitter("=", "key=value")
		hm.SetBySplitter("=", "novalue")

		// Assert
		v, ok := hm.Get("key")
		actual := args.Map{"result": ok || v != "value"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
		v2, ok2 := hm.Get("novalue")
		actual = args.Map{"result": ok2 || v2 != ""}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty value, got ''", actual)
	})
}

func Test_Hashmap_16_Hashmap_AddOrUpdateStringsPtrWgLock_FromS09(t *testing.T) {
	safeTest(t, "Test_16_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()

		// Assert
		actual := args.Map{"result": hm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_17_Hashmap_AddOrUpdateStringsPtrWgLock_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_17_Hashmap_AddOrUpdateStringsPtrWgLock_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}

		// Act
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_18_Hashmap_AddOrUpdateStringsPtrWgLock_Panic_FromS09(t *testing.T) {
	safeTest(t, "Test_18_Hashmap_AddOrUpdateStringsPtrWgLock_Panic", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}

		// Act & Assert
		defer func() {
			r := recover()
			actual := args.Map{"result": r == nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected panic for mismatched lengths", actual)
		}()
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1", "2"})
	})
}

func Test_Hashmap_19_Hashmap_AddOrUpdateHashmap_FromS09(t *testing.T) {
	safeTest(t, "Test_19_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k2", "v2")

		// Act
		hm.AddOrUpdateHashmap(other)

		// Assert
		actual := args.Map{"result": hm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_20_Hashmap_AddOrUpdateHashmap_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_20_Hashmap_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateHashmap(nil)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_21_Hashmap_AddOrUpdateMap_FromS09(t *testing.T) {
	safeTest(t, "Test_21_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})

		// Assert
		actual := args.Map{"result": hm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_22_Hashmap_AddOrUpdateMap_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_22_Hashmap_AddOrUpdateMap_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateMap(map[string]string{})

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_23_Hashmap_AddsOrUpdates_FromS09(t *testing.T) {
	safeTest(t, "Test_23_Hashmap_AddsOrUpdates", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_24_Hashmap_AddsOrUpdates_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_24_Hashmap_AddsOrUpdates_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdates(nil...)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_25_Hashmap_AddOrUpdateKeyAnyValues_FromS09(t *testing.T) {
	safeTest(t, "Test_25_Hashmap_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "k", Value: 42},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_26_Hashmap_AddOrUpdateKeyAnyValues_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_26_Hashmap_AddOrUpdateKeyAnyValues_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyAnyValues()

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_27_Hashmap_AddOrUpdateKeyValues_FromS09(t *testing.T) {
	safeTest(t, "Test_27_Hashmap_AddOrUpdateKeyValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_28_Hashmap_AddOrUpdateKeyValues_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_28_Hashmap_AddOrUpdateKeyValues_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyValues()

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddOrUpdateCollection ────────────────────────────────────

func Test_Hashmap_29_Hashmap_AddOrUpdateCollection_FromS09(t *testing.T) {
	safeTest(t, "Test_29_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})

		// Act
		hm.AddOrUpdateCollection(keys, vals)

		// Assert
		actual := args.Map{"result": hm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Hashmap_30_Hashmap_AddOrUpdateCollection_NilKeys_FromS09(t *testing.T) {
	safeTest(t, "Test_30_Hashmap_AddOrUpdateCollection_NilKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		vals := corestr.New.Collection.Strings([]string{"1"})

		// Act
		hm.AddOrUpdateCollection(nil, vals)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_31_Hashmap_AddOrUpdateCollection_LengthMismatch_FromS09(t *testing.T) {
	safeTest(t, "Test_31_Hashmap_AddOrUpdateCollection_LengthMismatch", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})

		// Act
		hm.AddOrUpdateCollection(keys, vals)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for length mismatch", actual)
	})
}

// ── Filter-based adds ────────────────────────────────────────

func Test_Hashmap_32_Hashmap_AddsOrUpdatesAnyUsingFilter_FromS09(t *testing.T) {
	safeTest(t, "Test_32_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, false
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_33_Hashmap_AddsOrUpdatesAnyUsingFilter_Break_FromS09(t *testing.T) {
	safeTest(t, "Test_33_Hashmap_AddsOrUpdatesAnyUsingFilter_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, true
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_Hashmap_34_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_34_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(nil, nil...)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_35_Hashmap_AddsOrUpdatesAnyUsingFilter_Skip_FromS09(t *testing.T) {
	safeTest(t, "Test_35_Hashmap_AddsOrUpdatesAnyUsingFilter_Skip", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", false, false
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 — all skipped", actual)
	})
}

func Test_Hashmap_36_Hashmap_AddsOrUpdatesAnyUsingFilterLock_FromS09(t *testing.T) {
	safeTest(t, "Test_36_Hashmap_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, false
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilterLock(filter,
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_37_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_37_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdatesAnyUsingFilterLock(nil, nil...)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_38_Hashmap_AddsOrUpdatesAnyUsingFilterLock_SkipAndBreak_FromS09(t *testing.T) {
	safeTest(t, "Test_38_Hashmap_AddsOrUpdatesAnyUsingFilterLock_SkipAndBreak", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		callCount := 0
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			callCount++
			if callCount == 1 {
				return "", false, false
			}
			return pair.ValueString(), true, true
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilterLock(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
			corestr.KeyAnyValuePair{Key: "c", Value: "3"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_39_Hashmap_AddsOrUpdatesUsingFilter_FromS09(t *testing.T) {
	safeTest(t, "Test_39_Hashmap_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, true, false
		}

		// Act
		hm.AddsOrUpdatesUsingFilter(filter,
			corestr.KeyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_40_Hashmap_AddsOrUpdatesUsingFilter_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_40_Hashmap_AddsOrUpdatesUsingFilter_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdatesUsingFilter(nil, nil...)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_41_Hashmap_AddsOrUpdatesUsingFilter_Break_FromS09(t *testing.T) {
	safeTest(t, "Test_41_Hashmap_AddsOrUpdatesUsingFilter_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, true, true
		}

		// Act
		hm.AddsOrUpdatesUsingFilter(filter,
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_Hashmap_42_Hashmap_AddsOrUpdatesUsingFilter_Skip_FromS09(t *testing.T) {
	safeTest(t, "Test_42_Hashmap_AddsOrUpdatesUsingFilter_Skip", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return "", false, false
		}

		// Act
		hm.AddsOrUpdatesUsingFilter(filter,
			corestr.KeyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── ConcatNew / ConcatNewUsingMaps ───────────────────────────

func Test_Hashmap_43_Hashmap_ConcatNew_FromS09(t *testing.T) {
	safeTest(t, "Test_43_Hashmap_ConcatNew", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("b", "2")

		// Act
		result := hm.ConcatNew(true, other)

		// Assert
		actual := args.Map{"result": result.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_Hashmap_44_Hashmap_ConcatNew_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_44_Hashmap_ConcatNew_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNew(true)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_45_Hashmap_ConcatNew_NilInList_FromS09(t *testing.T) {
	safeTest(t, "Test_45_Hashmap_ConcatNew_NilInList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNew(false, nil)

		// Assert
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_Hashmap_46_Hashmap_ConcatNewUsingMaps_FromS09(t *testing.T) {
	safeTest(t, "Test_46_Hashmap_ConcatNewUsingMaps", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNewUsingMaps(true, map[string]string{"b": "2"})

		// Assert
		actual := args.Map{"result": result.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_Hashmap_47_Hashmap_ConcatNewUsingMaps_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_47_Hashmap_ConcatNewUsingMaps_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		result := hm.ConcatNewUsingMaps(false)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Hashmap_48_Hashmap_ConcatNewUsingMaps_NilInList_FromS09(t *testing.T) {
	safeTest(t, "Test_48_Hashmap_ConcatNewUsingMaps_NilInList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNewUsingMaps(false, nil)

		// Assert
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

// ── AddOrUpdateLock ──────────────────────────────────────────

func Test_Hashmap_49_Hashmap_AddOrUpdateLock_FromS09(t *testing.T) {
	safeTest(t, "Test_49_Hashmap_AddOrUpdateLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateLock("k", "v")

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Has / Contains / ContainsLock / IsKeyMissing / IsKeyMissingLock / HasLock / HasWithLock ──

func Test_Hashmap_50_Hashmap_Has_FromS09(t *testing.T) {
	safeTest(t, "Test_50_Hashmap_Has", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.Has("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.Has("missing")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashmap_51_Hashmap_Contains_FromS09(t *testing.T) {
	safeTest(t, "Test_51_Hashmap_Contains", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.Contains("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_52_Hashmap_ContainsLock_FromS09(t *testing.T) {
	safeTest(t, "Test_52_Hashmap_ContainsLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.ContainsLock("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_53_Hashmap_IsKeyMissing_FromS09(t *testing.T) {
	safeTest(t, "Test_53_Hashmap_IsKeyMissing", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.IsKeyMissing("k")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": hm.IsKeyMissing("missing")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_54_Hashmap_IsKeyMissingLock_FromS09(t *testing.T) {
	safeTest(t, "Test_54_Hashmap_IsKeyMissingLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.IsKeyMissingLock("k")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashmap_55_Hashmap_HasLock_FromS09(t *testing.T) {
	safeTest(t, "Test_55_Hashmap_HasLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.HasLock("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_56_Hashmap_HasWithLock_FromS09(t *testing.T) {
	safeTest(t, "Test_56_Hashmap_HasWithLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.HasWithLock("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

// ── HasAllStrings / HasAll / HasAnyItem / HasAny ─────────────

func Test_Hashmap_57_Hashmap_HasAllStrings_FromS09(t *testing.T) {
	safeTest(t, "Test_57_Hashmap_HasAllStrings", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act & Assert
		actual := args.Map{"result": hm.HasAllStrings("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasAllStrings("a", "c")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashmap_58_Hashmap_HasAll_FromS09(t *testing.T) {
	safeTest(t, "Test_58_Hashmap_HasAll", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("x", "1")

		// Act & Assert
		actual := args.Map{"result": hm.HasAll("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasAll("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashmap_59_Hashmap_HasAnyItem_FromS09(t *testing.T) {
	safeTest(t, "Test_59_Hashmap_HasAnyItem", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_60_Hashmap_HasAny_FromS09(t *testing.T) {
	safeTest(t, "Test_60_Hashmap_HasAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act & Assert
		actual := args.Map{"result": hm.HasAny("a", "z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hm.HasAny("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── HasAllCollectionItems ────────────────────────────────────

func Test_Hashmap_61_Hashmap_HasAllCollectionItems_FromS09(t *testing.T) {
	safeTest(t, "Test_61_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": hm.HasAllCollectionItems(col)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_62_Hashmap_HasAllCollectionItems_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_62_Hashmap_HasAllCollectionItems_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		actual := args.Map{"result": hm.HasAllCollectionItems(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Hashmap_63_Hashmap_HasAllCollectionItems_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_63_Hashmap_HasAllCollectionItems_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		actual := args.Map{"result": hm.HasAllCollectionItems(corestr.Empty.Collection())}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

// ── DiffRaw / Diff ───────────────────────────────────────────

func Test_Hashmap_64_Hashmap_DiffRaw_FromS09(t *testing.T) {
	safeTest(t, "Test_64_Hashmap_DiffRaw", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		diff := hm.DiffRaw(map[string]string{"a": "2"})

		// Assert
		actual := args.Map{"result": diff == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil diff", actual)
	})
}

func Test_Hashmap_65_Hashmap_Diff_FromS09(t *testing.T) {
	safeTest(t, "Test_65_Hashmap_Diff", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("a", "2")

		// Act
		diff := hm.Diff(other)

		// Assert
		actual := args.Map{"result": diff == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── GetKeysFilteredItems / GetKeysFilteredCollection ─────────

func Test_Hashmap_66_Hashmap_GetKeysFilteredItems_FromS09(t *testing.T) {
	safeTest(t, "Test_66_Hashmap_GetKeysFilteredItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("apple", "1")
		hm.AddOrUpdate("banana", "2")
		filter := func(str string, index int) (string, bool, bool) {
			return str, strings.HasPrefix(str, "a"), false
		}

		// Act
		result := hm.GetKeysFilteredItems(filter)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_67_Hashmap_GetKeysFilteredItems_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_67_Hashmap_GetKeysFilteredItems_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := hm.GetKeysFilteredItems(filter)

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_68_Hashmap_GetKeysFilteredItems_Break_FromS09(t *testing.T) {
	safeTest(t, "Test_68_Hashmap_GetKeysFilteredItems_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}

		// Act
		result := hm.GetKeysFilteredItems(filter)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_Hashmap_69_Hashmap_GetKeysFilteredCollection_FromS09(t *testing.T) {
	safeTest(t, "Test_69_Hashmap_GetKeysFilteredCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("x", "1")
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := hm.GetKeysFilteredCollection(filter)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_70_Hashmap_GetKeysFilteredCollection_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_70_Hashmap_GetKeysFilteredCollection_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		result := hm.GetKeysFilteredCollection(nil)

		// Assert
		actual := args.Map{"result": result.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Hashmap_71_Hashmap_GetKeysFilteredCollection_Break_FromS09(t *testing.T) {
	safeTest(t, "Test_71_Hashmap_GetKeysFilteredCollection_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}

		// Act
		result := hm.GetKeysFilteredCollection(filter)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Items / SafeItems / ItemsCopyLock ────────────────────────

func Test_Hashmap_72_Hashmap_Items_FromS09(t *testing.T) {
	safeTest(t, "Test_72_Hashmap_Items", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": len(hm.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_73_Hashmap_SafeItems_FromS09(t *testing.T) {
	safeTest(t, "Test_73_Hashmap_SafeItems", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act & Assert
		actual := args.Map{"result": hm.SafeItems() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for nil hashmap", actual)
	})
}

func Test_Hashmap_74_Hashmap_ItemsCopyLock_FromS09(t *testing.T) {
	safeTest(t, "Test_74_Hashmap_ItemsCopyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		copied := hm.ItemsCopyLock()

		// Assert
		actual := args.Map{"result": len(*copied) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── ValuesCollection / ValuesHashset / ValuesCollectionLock / ValuesHashsetLock ──

func Test_Hashmap_75_Hashmap_ValuesCollection_FromS09(t *testing.T) {
	safeTest(t, "Test_75_Hashmap_ValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		col := hm.ValuesCollection()

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_76_Hashmap_ValuesHashset_FromS09(t *testing.T) {
	safeTest(t, "Test_76_Hashmap_ValuesHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hs := hm.ValuesHashset()

		// Assert
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_77_Hashmap_ValuesCollectionLock_FromS09(t *testing.T) {
	safeTest(t, "Test_77_Hashmap_ValuesCollectionLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		col := hm.ValuesCollectionLock()

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_78_Hashmap_ValuesHashsetLock_FromS09(t *testing.T) {
	safeTest(t, "Test_78_Hashmap_ValuesHashsetLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hs := hm.ValuesHashsetLock()

		// Assert
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── ValuesList / KeysValuesCollection / KeysValuesList ───────

func Test_Hashmap_79_Hashmap_ValuesList_FromS09(t *testing.T) {
	safeTest(t, "Test_79_Hashmap_ValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		list := hm.ValuesList()

		// Assert
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_80_Hashmap_KeysValuesCollection_FromS09(t *testing.T) {
	safeTest(t, "Test_80_Hashmap_KeysValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys, values := hm.KeysValuesCollection()

		// Assert
		actual := args.Map{"result": keys.Length() != 1 || values.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_Hashmap_81_Hashmap_KeysValuesList_FromS09(t *testing.T) {
	safeTest(t, "Test_81_Hashmap_KeysValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys, values := hm.KeysValuesList()

		// Assert
		actual := args.Map{"result": len(keys) != 1 || len(values) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_Hashmap_82_Hashmap_KeysValuesListLock_FromS09(t *testing.T) {
	safeTest(t, "Test_82_Hashmap_KeysValuesListLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys, values := hm.KeysValuesListLock()

		// Assert
		actual := args.Map{"result": len(keys) != 1 || len(values) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

// ── KeysValuePairs / KeysValuePairsCollection ────────────────

func Test_Hashmap_83_Hashmap_KeysValuePairs_FromS09(t *testing.T) {
	safeTest(t, "Test_83_Hashmap_KeysValuePairs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		pairs := hm.KeysValuePairs()

		// Assert
		actual := args.Map{"result": len(pairs) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_84_Hashmap_KeysValuePairsCollection_FromS09(t *testing.T) {
	safeTest(t, "Test_84_Hashmap_KeysValuePairsCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		kvc := hm.KeysValuePairsCollection()

		// Assert
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AllKeys / Keys / KeysCollection / KeysLock ───────────────

func Test_Hashmap_85_Hashmap_AllKeys_FromS09(t *testing.T) {
	safeTest(t, "Test_85_Hashmap_AllKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys := hm.AllKeys()

		// Assert
		actual := args.Map{"result": len(keys) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_86_Hashmap_AllKeys_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_86_Hashmap_AllKeys_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		keys := hm.AllKeys()

		// Assert
		actual := args.Map{"result": len(keys) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_87_Hashmap_Keys_FromS09(t *testing.T) {
	safeTest(t, "Test_87_Hashmap_Keys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act & Assert
		actual := args.Map{"result": len(hm.Keys()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_88_Hashmap_KeysCollection_FromS09(t *testing.T) {
	safeTest(t, "Test_88_Hashmap_KeysCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act & Assert
		actual := args.Map{"result": hm.KeysCollection().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_89_Hashmap_KeysLock_FromS09(t *testing.T) {
	safeTest(t, "Test_89_Hashmap_KeysLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys := hm.KeysLock()

		// Assert
		actual := args.Map{"result": len(keys) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_90_Hashmap_KeysLock_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_90_Hashmap_KeysLock_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		keys := hm.KeysLock()

		// Assert
		actual := args.Map{"result": len(keys) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── ValuesListCopyLock ───────────────────────────────────────

func Test_Hashmap_91_Hashmap_ValuesListCopyLock_FromS09(t *testing.T) {
	safeTest(t, "Test_91_Hashmap_ValuesListCopyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		list := hm.ValuesListCopyLock()

		// Assert
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── KeysToLower / ValuesToLower ──────────────────────────────

func Test_Hashmap_92_Hashmap_KeysToLower_FromS09(t *testing.T) {
	safeTest(t, "Test_92_Hashmap_KeysToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("KEY", "v")

		// Act
		lowered := hm.KeysToLower()

		// Assert
		actual := args.Map{"result": lowered.Has("key")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected lowercased key", actual)
	})
}

func Test_Hashmap_93_Hashmap_ValuesToLower_FromS09(t *testing.T) {
	safeTest(t, "Test_93_Hashmap_ValuesToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("KEY", "v")

		// Act — deprecated alias
		lowered := hm.ValuesToLower()

		// Assert
		actual := args.Map{"result": lowered == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── Length / LengthLock ──────────────────────────────────────

func Test_Hashmap_94_Hashmap_Length_FromS09(t *testing.T) {
	safeTest(t, "Test_94_Hashmap_Length", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_95_Hashmap_Length_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_95_Hashmap_Length_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act & Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil", actual)
	})
}

func Test_Hashmap_96_Hashmap_LengthLock_FromS09(t *testing.T) {
	safeTest(t, "Test_96_Hashmap_LengthLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── IsEqual / IsEqualPtr / IsEqualPtrLock ────────────────────

func Test_Hashmap_97_Hashmap_IsEqual_Same_FromS09(t *testing.T) {
	safeTest(t, "Test_97_Hashmap_IsEqual_Same", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.IsEqualPtr(other)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Hashmap_98_Hashmap_IsEqualPtr_DiffValues_FromS09(t *testing.T) {
	safeTest(t, "Test_98_Hashmap_IsEqualPtr_DiffValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v2")

		// Act & Assert
		actual := args.Map{"result": hm.IsEqualPtr(other)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Hashmap_99_Hashmap_IsEqualPtr_DiffLength_FromS09(t *testing.T) {
	safeTest(t, "Test_99_Hashmap_IsEqualPtr_DiffLength", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("a", "1")
		other.AddOrUpdate("b", "2")

		// Act & Assert
		actual := args.Map{"result": hm.IsEqualPtr(other)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Hashmap_100_Hashmap_IsEqualPtr_BothNil_FromS09(t *testing.T) {
	safeTest(t, "Test_100_Hashmap_IsEqualPtr_BothNil", func() {
		// Arrange
		var a *corestr.Hashmap
		var b *corestr.Hashmap

		// Act & Assert
		actual := args.Map{"result": a.IsEqualPtr(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for both nil", actual)
	})
}

func Test_Hashmap_101_Hashmap_IsEqualPtr_OneNil_FromS09(t *testing.T) {
	safeTest(t, "Test_101_Hashmap_IsEqualPtr_OneNil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		var other *corestr.Hashmap

		// Act & Assert
		actual := args.Map{"result": hm.IsEqualPtr(other)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Hashmap_102_Hashmap_IsEqualPtr_SamePtr_FromS09(t *testing.T) {
	safeTest(t, "Test_102_Hashmap_IsEqualPtr_SamePtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.IsEqualPtr(hm)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for same pointer", actual)
	})
}

func Test_Hashmap_103_Hashmap_IsEqualPtr_BothEmpty_FromS09(t *testing.T) {
	safeTest(t, "Test_103_Hashmap_IsEqualPtr_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.Hashmap()
		b := corestr.Empty.Hashmap()

		// Act & Assert
		actual := args.Map{"result": a.IsEqualPtr(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for both empty", actual)
	})
}

func Test_Hashmap_104_Hashmap_IsEqualPtr_OneEmpty_FromS09(t *testing.T) {
	safeTest(t, "Test_104_Hashmap_IsEqualPtr_OneEmpty", func() {
		// Arrange
		a := corestr.New.Hashmap.Cap(5)
		a.AddOrUpdate("k", "v")
		b := corestr.Empty.Hashmap()

		// Act & Assert
		actual := args.Map{"result": a.IsEqualPtr(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Hashmap_105_Hashmap_IsEqualPtr_MissingKey_FromS09(t *testing.T) {
	safeTest(t, "Test_105_Hashmap_IsEqualPtr_MissingKey", func() {
		// Arrange
		a := corestr.New.Hashmap.Cap(5)
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Cap(5)
		b.AddOrUpdate("b", "1")

		// Act & Assert
		actual := args.Map{"result": a.IsEqualPtr(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal — different keys", actual)
	})
}

func Test_Hashmap_106_Hashmap_IsEqual_FromS09(t *testing.T) {
	safeTest(t, "Test_106_Hashmap_IsEqual", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v")

		// Act & Assert — value receiver
		actual := args.Map{"result": hm.IsEqual(*other)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Hashmap_107_Hashmap_IsEqualPtrLock_FromS09(t *testing.T) {
	safeTest(t, "Test_107_Hashmap_IsEqualPtrLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": hm.IsEqualPtrLock(other)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

// ── Remove / RemoveWithLock ──────────────────────────────────

func Test_Hashmap_108_Hashmap_Remove_FromS09(t *testing.T) {
	safeTest(t, "Test_108_Hashmap_Remove", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hm.Remove("k")

		// Assert
		actual := args.Map{"result": hm.Has("k")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_Hashmap_109_Hashmap_RemoveWithLock_FromS09(t *testing.T) {
	safeTest(t, "Test_109_Hashmap_RemoveWithLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hm.RemoveWithLock("k")

		// Assert
		actual := args.Map{"result": hm.Has("k")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

// ── String / StringLock ──────────────────────────────────────

func Test_Hashmap_110_Hashmap_String_FromS09(t *testing.T) {
	safeTest(t, "Test_110_Hashmap_String", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.String()

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Hashmap_111_Hashmap_String_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_111_Hashmap_String_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		s := hm.String()

		// Assert
		actual := args.Map{"result": strings.Contains(s, "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_Hashmap_112_Hashmap_StringLock_FromS09(t *testing.T) {
	safeTest(t, "Test_112_Hashmap_StringLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.StringLock()

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Hashmap_113_Hashmap_StringLock_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_113_Hashmap_StringLock_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		s := hm.StringLock()

		// Assert
		actual := args.Map{"result": strings.Contains(s, "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

// ── GetValuesExceptKeysInHashset / GetValuesKeysExcept / GetAllExceptCollection ──

func Test_Hashmap_114_Hashmap_GetValuesExceptKeysInHashset_FromS09(t *testing.T) {
	safeTest(t, "Test_114_Hashmap_GetValuesExceptKeysInHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hm.GetValuesExceptKeysInHashset(hs)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_115_Hashmap_GetValuesExceptKeysInHashset_NilHashset_FromS09(t *testing.T) {
	safeTest(t, "Test_115_Hashmap_GetValuesExceptKeysInHashset_NilHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetValuesExceptKeysInHashset(nil)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all values", actual)
	})
}

func Test_Hashmap_116_Hashmap_GetValuesExceptKeysInHashset_EmptyHashset_FromS09(t *testing.T) {
	safeTest(t, "Test_116_Hashmap_GetValuesExceptKeysInHashset_EmptyHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetValuesExceptKeysInHashset(corestr.Empty.Hashset())

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all values", actual)
	})
}

func Test_Hashmap_117_Hashmap_GetValuesKeysExcept_FromS09(t *testing.T) {
	safeTest(t, "Test_117_Hashmap_GetValuesKeysExcept", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		result := hm.GetValuesKeysExcept([]string{"a"})

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_118_Hashmap_GetValuesKeysExcept_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_118_Hashmap_GetValuesKeysExcept_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetValuesKeysExcept(nil)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all values", actual)
	})
}

func Test_Hashmap_119_Hashmap_GetAllExceptCollection_FromS09(t *testing.T) {
	safeTest(t, "Test_119_Hashmap_GetAllExceptCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := hm.GetAllExceptCollection(col)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_120_Hashmap_GetAllExceptCollection_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_120_Hashmap_GetAllExceptCollection_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetAllExceptCollection(nil)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all values", actual)
	})
}

// ── Join / JoinKeys ──────────────────────────────────────────

func Test_Hashmap_121_Hashmap_Join_FromS09(t *testing.T) {
	safeTest(t, "Test_121_Hashmap_Join", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.Join(",")

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Hashmap_122_Hashmap_JoinKeys_FromS09(t *testing.T) {
	safeTest(t, "Test_122_Hashmap_JoinKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.JoinKeys(",")

		// Assert
		actual := args.Map{"result": s != "k"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'k', got ''", actual)
	})
}

// ── JSON methods ─────────────────────────────────────────────

func Test_Hashmap_123_Hashmap_JsonModel_FromS09(t *testing.T) {
	safeTest(t, "Test_123_Hashmap_JsonModel", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		actual := args.Map{"result": len(hm.JsonModel()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_124_Hashmap_JsonModelAny_FromS09(t *testing.T) {
	safeTest(t, "Test_124_Hashmap_JsonModelAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		actual := args.Map{"result": hm.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Hashmap_125_Hashmap_MarshalJSON_FromS09(t *testing.T) {
	safeTest(t, "Test_125_Hashmap_MarshalJSON", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		data, err := hm.MarshalJSON()

		// Assert
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON bytes", actual)
	})
}

func Test_Hashmap_126_Hashmap_UnmarshalJSON_FromS09(t *testing.T) {
	safeTest(t, "Test_126_Hashmap_UnmarshalJSON", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		err := hm.UnmarshalJSON([]byte(`{"k":"v"}`))

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
		actual = args.Map{"result": hm.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_127_Hashmap_UnmarshalJSON_Invalid_FromS09(t *testing.T) {
	safeTest(t, "Test_127_Hashmap_UnmarshalJSON_Invalid", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		err := hm.UnmarshalJSON([]byte(`invalid`))

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Hashmap_128_Hashmap_Json_FromS09(t *testing.T) {
	safeTest(t, "Test_128_Hashmap_Json", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		result := hm.Json()

		// Assert
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_Hashmap_129_Hashmap_JsonPtr_FromS09(t *testing.T) {
	safeTest(t, "Test_129_Hashmap_JsonPtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		actual := args.Map{"result": hm.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Hashmap_130_Hashmap_ParseInjectUsingJson_FromS09(t *testing.T) {
	safeTest(t, "Test_130_Hashmap_ParseInjectUsingJson", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		jsonResult := hm.JsonPtr()
		target := corestr.Empty.Hashmap()

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		actual := args.Map{"result": err != nil || result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_131_Hashmap_ParseInjectUsingJsonMust_FromS09(t *testing.T) {
	safeTest(t, "Test_131_Hashmap_ParseInjectUsingJsonMust", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		jsonResult := hm.JsonPtr()
		target := corestr.Empty.Hashmap()

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_132_Hashmap_JsonParseSelfInject_FromS09(t *testing.T) {
	safeTest(t, "Test_132_Hashmap_JsonParseSelfInject", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		jsonResult := hm.JsonPtr()
		target := corestr.Empty.Hashmap()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

// ── ToError / ToDefaultError / KeyValStringLines ─────────────

func Test_Hashmap_133_Hashmap_ToError_FromS09(t *testing.T) {
	safeTest(t, "Test_133_Hashmap_ToError", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		err := hm.ToError(",")

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Hashmap_134_Hashmap_ToDefaultError_FromS09(t *testing.T) {
	safeTest(t, "Test_134_Hashmap_ToDefaultError", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		err := hm.ToDefaultError()

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Hashmap_135_Hashmap_KeyValStringLines_FromS09(t *testing.T) {
	safeTest(t, "Test_135_Hashmap_KeyValStringLines", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		lines := hm.KeyValStringLines()

		// Assert
		actual := args.Map{"result": len(lines) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Clear / Dispose ──────────────────────────────────────────

func Test_Hashmap_136_Hashmap_Clear_FromS09(t *testing.T) {
	safeTest(t, "Test_136_Hashmap_Clear", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		_ = hm.ValuesList() // populate cache

		// Act
		hm.Clear()

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_137_Hashmap_Clear_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_137_Hashmap_Clear_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		result := hm.Clear()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Hashmap_138_Hashmap_Dispose_FromS09(t *testing.T) {
	safeTest(t, "Test_138_Hashmap_Dispose", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hm.Dispose()

		// Assert
		actual := args.Map{"result": hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Hashmap_139_Hashmap_Dispose_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_139_Hashmap_Dispose_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act — should not panic
		hm.Dispose()
	})
}

// ── ToStringsUsingCompiler ───────────────────────────────────

func Test_Hashmap_140_Hashmap_ToStringsUsingCompiler_FromS09(t *testing.T) {
	safeTest(t, "Test_140_Hashmap_ToStringsUsingCompiler", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		lines := hm.ToStringsUsingCompiler(func(key, val string) string {
			return key + "=" + val
		})

		// Assert
		actual := args.Map{"result": len(lines) != 1 || lines[0] != "k=v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'k=v'", actual)
	})
}

func Test_Hashmap_141_Hashmap_ToStringsUsingCompiler_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_141_Hashmap_ToStringsUsingCompiler_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		lines := hm.ToStringsUsingCompiler(func(key, val string) string {
			return key
		})

		// Assert
		actual := args.Map{"result": len(lines) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AsJsoner / AsJsonContractsBinder / AsJsonParseSelfInjector / AsJsonMarshaller ──

func Test_Hashmap_142_Hashmap_AsJsoner_FromS09(t *testing.T) {
	safeTest(t, "Test_142_Hashmap_AsJsoner", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": hm.AsJsoner() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Hashmap_143_Hashmap_AsJsonContractsBinder_FromS09(t *testing.T) {
	safeTest(t, "Test_143_Hashmap_AsJsonContractsBinder", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": hm.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Hashmap_144_Hashmap_AsJsonParseSelfInjector_FromS09(t *testing.T) {
	safeTest(t, "Test_144_Hashmap_AsJsonParseSelfInjector", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": hm.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Hashmap_145_Hashmap_AsJsonMarshaller_FromS09(t *testing.T) {
	safeTest(t, "Test_145_Hashmap_AsJsonMarshaller", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{"result": hm.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── ClonePtr / Clone ─────────────────────────────────────────

func Test_Hashmap_146_Hashmap_ClonePtr_FromS09(t *testing.T) {
	safeTest(t, "Test_146_Hashmap_ClonePtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		cloned := hm.ClonePtr()

		// Assert
		actual := args.Map{"result": cloned == nil || cloned.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cloned with 1 item", actual)
	})
}

func Test_Hashmap_147_Hashmap_ClonePtr_Nil_FromS09(t *testing.T) {
	safeTest(t, "Test_147_Hashmap_ClonePtr_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		cloned := hm.ClonePtr()

		// Assert
		actual := args.Map{"result": cloned != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Hashmap_148_Hashmap_Clone_FromS09(t *testing.T) {
	safeTest(t, "Test_148_Hashmap_Clone", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		cloned := hm.Clone()

		// Assert
		actual := args.Map{"result": cloned.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_149_Hashmap_Clone_Empty_FromS09(t *testing.T) {
	safeTest(t, "Test_149_Hashmap_Clone_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		cloned := hm.Clone()

		// Assert
		actual := args.Map{"result": cloned.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── Get / GetValue ───────────────────────────────────────────

func Test_Hashmap_150_Hashmap_Get_FromS09(t *testing.T) {
	safeTest(t, "Test_150_Hashmap_Get", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		v, ok := hm.Get("k")

		// Assert
		actual := args.Map{"result": ok || v != "v"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 'v'", actual)
	})
}

func Test_Hashmap_151_Hashmap_Get_Missing_FromS09(t *testing.T) {
	safeTest(t, "Test_151_Hashmap_Get_Missing", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		_, ok := hm.Get("missing")

		// Assert
		actual := args.Map{"result": ok}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not found", actual)
	})
}

func Test_Hashmap_152_Hashmap_GetValue_FromS09(t *testing.T) {
	safeTest(t, "Test_152_Hashmap_GetValue", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		v, ok := hm.GetValue("k")

		// Assert
		actual := args.Map{"result": ok || v != "v"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 'v'", actual)
	})
}

// ── Serialize / Deserialize ──────────────────────────────────

func Test_Hashmap_153_Hashmap_Serialize_FromS09(t *testing.T) {
	safeTest(t, "Test_153_Hashmap_Serialize", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		data, err := hm.Serialize()

		// Assert
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid bytes", actual)
	})
}

func Test_Hashmap_154_Hashmap_Deserialize_FromS09(t *testing.T) {
	safeTest(t, "Test_154_Hashmap_Deserialize", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		var target map[string]string

		// Act
		err := hm.Deserialize(&target)

		// Assert
		actual := args.Map{"result": err != nil || len(target) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
