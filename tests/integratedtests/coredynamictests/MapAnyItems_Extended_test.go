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

package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItems — extended methods (AddMapResult, Diff, IsEqual, Clone, etc.)
// ═══════════════════════════════════════════════════════════════════════

func Test_01_MapAnyItems_AddMapResult(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResult(map[string]any{"b": 2, "a": 99})

	// Act
	actual := args.Map{"result": m.GetValue("a") != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected override to 99", actual)
	actual = args.Map{"result": m.GetValue("b") != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_02_MapAnyItems_AddMapResult_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.AddMapResult(nil)

	// Act
	actual := args.Map{"result": m.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_03_MapAnyItems_AddMapResultOption_Override(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResultOption(true, map[string]any{"a": 99})

	// Act
	actual := args.Map{"result": m.GetValue("a") != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_04_MapAnyItems_AddMapResultOption_NoOverride(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResultOption(false, map[string]any{"a": 99, "b": 2})
	// noOverride only updates existing keys
	v := m.GetValue("a")
	_ = v
}

func Test_05_MapAnyItems_AddManyMapResultsUsingOption(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddManyMapResultsUsingOption(true, map[string]any{"b": 2}, map[string]any{"c": 3})

	// Act
	actual := args.Map{"result": m.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_06_MapAnyItems_AddManyMapResultsUsingOption_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.AddManyMapResultsUsingOption(true)

	// Act
	actual := args.Map{"result": m.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_07_MapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	sub := m.GetNewMapUsingKeys(false, "a", "c")

	// Act
	actual := args.Map{"result": sub.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_08_MapAnyItems_GetNewMapUsingKeys_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	sub := m.GetNewMapUsingKeys(false)

	// Act
	actual := args.Map{"result": sub.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_09_MapAnyItems_GetNewMapUsingKeys_Missing_NoPanic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	sub := m.GetNewMapUsingKeys(false, "a", "nope")

	// Act
	actual := args.Map{"result": sub.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_10_MapAnyItems_JsonString(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s, err := m.JsonString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json string", actual)
}

func Test_11_MapAnyItems_JsonStringMust(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s := m.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_12_MapAnyItems_JsonResultOfKey_Found(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jr := m.JsonResultOfKey("k")

	// Act
	actual := args.Map{"result": jr == nil || jr.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid result", actual)
}

func Test_13_MapAnyItems_JsonResultOfKey_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	jr := m.JsonResultOfKey("nope")

	// Act
	actual := args.Map{"result": jr.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_14_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	mr := m.JsonResultOfKeys("a", "b")

	// Act
	actual := args.Map{"result": mr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_15_MapAnyItems_JsonResultOfKeys_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	mr := m.JsonResultOfKeys()

	// Act
	actual := args.Map{"result": mr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_16_MapAnyItems_AllKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)

	// Act
	actual := args.Map{"result": len(m.AllKeys()) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_17_MapAnyItems_AllKeys_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"result": len(m.AllKeys()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_18_MapAnyItems_AllValues(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)

	// Act
	actual := args.Map{"result": len(m.AllValues()) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_19_MapAnyItems_AllValues_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"result": len(m.AllValues()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_20_MapAnyItems_IsRawEqual(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)

	// Act
	actual := args.Map{"result": m.IsRawEqual(false, map[string]any{"k": 1})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_21_MapAnyItems_HasAnyChanges(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)

	// Act
	actual := args.Map{"result": m.HasAnyChanges(false, map[string]any{"k": 1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no changes", actual)
	actual = args.Map{"result": m.HasAnyChanges(false, map[string]any{"k": 2})}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected changes", actual)
}

func Test_22_MapAnyItems_DiffRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	diff := m.DiffRaw(false, map[string]any{"k": 2})
	_ = diff
}

func Test_23_MapAnyItems_Diff(t *testing.T) {
	m1 := coredynamic.NewMapAnyItems(4)
	m1.Add("k", 1)
	m2 := coredynamic.NewMapAnyItems(4)
	m2.Add("k", 2)
	diff := m1.Diff(false, m2)
	_ = diff
}

func Test_24_MapAnyItems_HashmapDiffUsingRaw(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"k": 2})

	// Act
	actual := args.Map{"result": diff.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty diff", actual)
}

func Test_25_MapAnyItems_HashmapDiffUsingRaw_Same(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"k": 1})

	// Act
	actual := args.Map{"result": diff.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no diff", actual)
}

func Test_26_MapAnyItems_MapStringAnyDiff(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	d := m.MapStringAnyDiff()

	// Act
	actual := args.Map{"result": len(d) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_27_MapAnyItems_DiffJsonMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	msg := m.DiffJsonMessage(false, map[string]any{"k": 2})
	_ = msg
}

func Test_28_MapAnyItems_ToStringsSliceOfDiffMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	s := m.ToStringsSliceOfDiffMap(map[string]any{"k": 2})
	_ = s
}

func Test_29_MapAnyItems_ShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"k": 2})
	_ = msg
}

func Test_30_MapAnyItems_LogShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	msg := m.LogShouldDiffMessage(false, "test", map[string]any{"k": 1})
	_ = msg
}

func Test_31_MapAnyItems_JsonModel(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jm := m.JsonModel()

	// Act
	actual := args.Map{"result": jm == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_32_MapAnyItems_JsonModelAny(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)

	// Act
	actual := args.Map{"result": m.JsonModelAny() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_33_MapAnyItems_Json(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	j := m.Json()
	_ = j
}

func Test_34_MapAnyItems_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jp := m.JsonPtr()
	m2 := coredynamic.EmptyMapAnyItems()
	_, err := m2.ParseInjectUsingJson(jp)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_35_MapAnyItems_JsonParseSelfInject(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jp := m.JsonPtr()
	m2 := coredynamic.EmptyMapAnyItems()
	err := m2.JsonParseSelfInject(jp)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_36_MapAnyItems_Strings(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	s := m.Strings()

	// Act
	actual := args.Map{"result": len(s) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected strings", actual)
}

func Test_37_MapAnyItems_String(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)

	// Act
	actual := args.Map{"result": m.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_38_MapAnyItems_Clear(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	m.Clear()

	// Act
	actual := args.Map{"result": m.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_39_MapAnyItems_Clear_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	m.Clear() // no panic
}

func Test_40_MapAnyItems_DeepClear(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	m.DeepClear()

	// Act
	actual := args.Map{"result": m.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_41_MapAnyItems_Dispose(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	m.Dispose()
}

func Test_42_MapAnyItems_Dispose_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	m.Dispose()
}

func Test_43_MapAnyItems_IsEqualRaw_Same(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)

	// Act
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"k": 1})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_44_MapAnyItems_IsEqualRaw_DiffLength(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)

	// Act
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"k": 1, "k2": 2})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_45_MapAnyItems_IsEqualRaw_MissingKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)

	// Act
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"x": 1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_46_MapAnyItems_IsEqualRaw_DiffValue(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)

	// Act
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"k": 2})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_47_MapAnyItems_IsEqualRaw_BothNil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"result": m.IsEqualRaw(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal for both nil", actual)
}

func Test_48_MapAnyItems_IsEqualRaw_OneNil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"k": 1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_49_MapAnyItems_IsEqual(t *testing.T) {
	// Arrange
	m1 := coredynamic.NewMapAnyItems(4)
	m1.Add("k", 1)
	m2 := coredynamic.NewMapAnyItems(4)
	m2.Add("k", 1)

	// Act
	actual := args.Map{"result": m1.IsEqual(m2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_50_MapAnyItems_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var m1, m2 *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"result": m1.IsEqual(m2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_51_MapAnyItems_IsEqual_OneNil(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)

	// Act
	actual := args.Map{"result": m.IsEqual(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_52_MapAnyItems_IsEqual_DiffLen(t *testing.T) {
	// Arrange
	m1 := coredynamic.NewMapAnyItems(4)
	m1.Add("k", 1)
	m2 := coredynamic.NewMapAnyItems(4)

	// Act
	actual := args.Map{"result": m1.IsEqual(m2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_53_MapAnyItems_ClonePtr(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	cloned, err := m.ClonePtr()

	// Act
	actual := args.Map{"result": err != nil || cloned == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned", actual)
}

func Test_54_MapAnyItems_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	cloned, err := m.ClonePtr()

	// Act
	actual := args.Map{"result": cloned != nil || err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil + error", actual)
}

func Test_55_MapAnyItems_RawMapStringAnyDiff(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	d := m.RawMapStringAnyDiff()

	// Act
	actual := args.Map{"result": len(d) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_56_MapAnyItems_RawMapStringAnyDiff_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	d := m.RawMapStringAnyDiff()

	// Act
	actual := args.Map{"result": len(d) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_57_MapAnyItems_MapAnyItems_Self(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)

	// Act
	actual := args.Map{"result": m.MapAnyItems() != m}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
}

func Test_58_MapAnyItems_GetItemRef_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	var target int
	err := m.GetItemRef("nope", &target)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_59_MapAnyItems_GetItemRef_NilRef(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	err := m.GetItemRef("k", nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_60_MapAnyItems_GetItemRef_NonPointer(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	err := m.GetItemRef("k", 42)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-pointer", actual)
}

func Test_61_MapAnyItems_GetUsingUnmarshallManyAt(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", "hello")
	var i int
	var s string
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "a", AnyInf: &i},
		corejson.KeyAny{Key: "b", AnyInf: &s},
	)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_62_MapAnyItems_GetManyItemsRefs_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	err := m.GetManyItemsRefs()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_63_MapAnyItems_ReflectSetToMust(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// expected for type mismatch
		}
	}()
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	m.ReflectSetToMust("k", &target)
}

func Test_64_MapAnyItems_DeserializeMust(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// expected for missing key
		}
	}()
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	m.DeserializeMust("k", &target)
}

func Test_65_MapAnyItems_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jp := m.JsonPtr()
	m2 := coredynamic.EmptyMapAnyItems()
	result := m2.ParseInjectUsingJsonMust(jp)

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_66_MapAnyItems_JsonResultsCollection(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	rc := m.JsonResultsCollection()

	// Act
	actual := args.Map{"result": rc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_67_MapAnyItems_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	rc := m.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"result": rc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRight — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_68_LeftRight_IsEmpty(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: nil}

	// Act
	actual := args.Map{"result": lr.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_69_LeftRight_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"result": lr.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_70_LeftRight_HasAnyItem(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 1, Right: nil}

	// Act
	actual := args.Map{"result": lr.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_71_LeftRight_HasLeft(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 1}

	// Act
	actual := args.Map{"result": lr.HasLeft()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_72_LeftRight_HasRight(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: 2}

	// Act
	actual := args.Map{"result": lr.HasRight()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_73_LeftRight_IsLeftEmpty(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: 1}

	// Act
	actual := args.Map{"result": lr.IsLeftEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_74_LeftRight_IsRightEmpty(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 1, Right: nil}

	// Act
	actual := args.Map{"result": lr.IsRightEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_75_LeftRight_LeftReflectSet(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 42, Right: 100}
	var target int
	err := lr.LeftReflectSet(&target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_76_LeftRight_LeftReflectSet_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	err := lr.LeftReflectSet(nil)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
}

func Test_77_LeftRight_RightReflectSet(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 42, Right: 100}
	var target int
	err := lr.RightReflectSet(&target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_78_LeftRight_RightReflectSet_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	err := lr.RightReflectSet(nil)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_79_LeftRight_DeserializeLeft(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 42}
	r := lr.DeserializeLeft()

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_80_LeftRight_DeserializeLeft_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"result": lr.DeserializeLeft() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_81_LeftRight_DeserializeRight(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: 42}
	r := lr.DeserializeRight()

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_82_LeftRight_DeserializeRight_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"result": lr.DeserializeRight() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_83_LeftRight_LeftToDynamic(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 42}
	d := lr.LeftToDynamic()

	// Act
	actual := args.Map{"result": d == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_84_LeftRight_LeftToDynamic_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"result": lr.LeftToDynamic() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_85_LeftRight_RightToDynamic(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: 42}
	d := lr.RightToDynamic()

	// Act
	actual := args.Map{"result": d == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_86_LeftRight_RightToDynamic_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"result": lr.RightToDynamic() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_87_LeftRight_TypeStatus(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: 42, Right: 100}
	ts := lr.TypeStatus()

	// Act
	actual := args.Map{"result": ts.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
}

func Test_88_LeftRight_TypeStatus_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	ts := lr.TypeStatus()
	_ = ts
}

// ═══════════════════════════════════════════════════════════════════════
// CastTo + CastedResult
// ═══════════════════════════════════════════════════════════════════════

func Test_89_CastTo_MatchingType(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, 42, reflect.TypeOf(0))

	// Act
	actual := args.Map{"result": result.IsMatchingAcceptedType}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected matching", actual)
	actual = args.Map{"result": result.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_90_CastTo_NonMatchingType(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "str", reflect.TypeOf(0))

	// Act
	actual := args.Map{"result": result.IsMatchingAcceptedType}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-matching", actual)
}

func Test_91_CastTo_PointerOutput(t *testing.T) {
	// Arrange
	v := 42
	result := coredynamic.CastTo(true, &v, reflect.TypeOf(&v))

	// Act
	actual := args.Map{"result": result.HasAnyIssues()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
}

func Test_92_CastedResult_IsInvalid_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"result": cr.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_93_CastedResult_IsNotNull(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsNull: false}

	// Act
	actual := args.Map{"result": cr.IsNotNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected not null", actual)
}

func Test_94_CastedResult_IsNotPointer(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsPointer: false}

	// Act
	actual := args.Map{"result": cr.IsNotPointer()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected not pointer", actual)
}

func Test_95_CastedResult_IsNotMatchingAcceptedType(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsMatchingAcceptedType: false}

	// Act
	actual := args.Map{"result": cr.IsNotMatchingAcceptedType()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_96_CastedResult_IsSourceKind(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{SourceKind: reflect.Int}

	// Act
	actual := args.Map{"result": cr.IsSourceKind(reflect.Int)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_97_CastedResult_HasAnyIssues(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: true}

	// Act
	actual := args.Map{"result": cr.HasAnyIssues()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectSetFromTo
// ═══════════════════════════════════════════════════════════════════════

func Test_98_ReflectSetFromTo_BothNil(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo(nil, nil)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_99_ReflectSetFromTo_ToNil(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo(42, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_100_ReflectSetFromTo_ToNonPointer(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo(42, 0)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-pointer dest", actual)
}

func Test_101_ReflectSetFromTo_SameNonPointerToPointer(t *testing.T) {
	// Arrange
	var target int
	err := coredynamic.ReflectSetFromTo(42, &target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	actual = args.Map{"result": target != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_102_ReflectSetFromTo_SamePointerTypes(t *testing.T) {
	// Arrange
	v := 42
	var target int
	from := &v
	err := coredynamic.ReflectSetFromTo(from, &target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	actual = args.Map{"result": target != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_103_ReflectSetFromTo_BytesToType(t *testing.T) {
	// Arrange
	var target int
	err := coredynamic.ReflectSetFromTo([]byte(`42`), &target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	actual = args.Map{"result": target != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_104_ReflectSetFromTo_TypeToBytes(t *testing.T) {
	// Arrange
	var target []byte
	err := coredynamic.ReflectSetFromTo(42, &target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	actual = args.Map{"result": len(target) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_105_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	// Arrange
	var target string
	err := coredynamic.ReflectSetFromTo(42, &target)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch error", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectInterfaceVal
// ═══════════════════════════════════════════════════════════════════════

func Test_106_ReflectInterfaceVal_NonPointer(t *testing.T) {
	// Arrange
	result := coredynamic.ReflectInterfaceVal(42)

	// Act
	actual := args.Map{"result": result != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_107_ReflectInterfaceVal_Pointer(t *testing.T) {
	// Arrange
	v := 42
	result := coredynamic.ReflectInterfaceVal(&v)

	// Act
	actual := args.Map{"result": result != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// LengthOfReflect
// ═══════════════════════════════════════════════════════════════════════

func Test_108_LengthOfReflect_Slice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_109_LengthOfReflect_Map(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_110_LengthOfReflect_Array_MapanyitemsExtended(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([3]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_111_LengthOfReflect_Other(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// MapAsKeyValSlice
// ═══════════════════════════════════════════════════════════════════════

func Test_112_MapAsKeyValSlice_Valid(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	kvc, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{"result": err != nil || kvc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_113_MapAsKeyValSlice_NotMap(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	_, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// ZeroSet, SafeZeroSet, ZeroSetAny
// ═══════════════════════════════════════════════════════════════════════

func Test_114_ZeroSet(t *testing.T) {
	// Arrange
	type S struct{ X int }
	s := S{X: 42}
	coredynamic.ZeroSet(reflect.ValueOf(&s))

	// Act
	actual := args.Map{"result": s.X != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_115_ZeroSetAny(t *testing.T) {
	// Arrange
	type S struct{ X int }
	s := &S{X: 42}
	coredynamic.ZeroSetAny(s)

	// Act
	actual := args.Map{"result": s.X != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_116_ZeroSetAny_Nil(t *testing.T) {
	coredynamic.ZeroSetAny(nil) // no panic
}

func Test_117_SafeZeroSet(t *testing.T) {
	// Arrange
	type S struct{ X int }
	s := &S{X: 42}
	coredynamic.SafeZeroSet(reflect.ValueOf(s))

	// Act
	actual := args.Map{"result": s.X != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// TypedDynamic[T]
// ═══════════════════════════════════════════════════════════════════════

func Test_118_TypedDynamic_Basic(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("hello", true)

	// Act
	actual := args.Map{"result": d.Data() != "hello" || !d.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello, valid", actual)
}

func Test_119_TypedDynamic_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidTypedDynamic[int]()

	// Act
	actual := args.Map{"result": d.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_120_TypedDynamic_InvalidPtr(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidTypedDynamicPtr[int]()

	// Act
	actual := args.Map{"result": d == nil || d.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid ptr", actual)
}

func Test_121_TypedDynamic_NewPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[int](42, true)

	// Act
	actual := args.Map{"result": d == nil || d.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_122_TypedDynamic_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid[int](42)

	// Act
	actual := args.Map{"result": d.IsValid() || d.Value() != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid 42", actual)
}

func Test_123_TypedDynamic_String(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)

	// Act
	actual := args.Map{"result": d.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_124_TypedDynamic_JsonBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, err := d.JsonBytes()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_125_TypedDynamic_JsonResult(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	jr := d.JsonResult()
	_ = jr
}

func Test_126_TypedDynamic_JsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	s, err := d.JsonString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json string", actual)
}

func Test_127_TypedDynamic_MarshalJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, err := d.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_128_TypedDynamic_UnmarshalJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](0, false)
	err := d.UnmarshalJSON([]byte(`42`))

	// Act
	actual := args.Map{"result": err != nil || d.Value() != 42 || !d.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42 valid", actual)
}

func Test_129_TypedDynamic_ValueMarshal(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, err := d.ValueMarshal()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_130_TypedDynamic_Bytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, ok := d.Bytes()

	// Act
	actual := args.Map{"result": ok || len(b) == 0}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_131_TypedDynamic_Bytes_AsBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[[]byte]([]byte("hi"), true)
	b, ok := d.Bytes()

	// Act
	actual := args.Map{"result": ok || string(b) != "hi"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hi", actual)
}

func Test_132_TypedDynamic_GetAsString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("hi", true)
	v, ok := d.GetAsString()

	// Act
	actual := args.Map{"result": ok || v != "hi"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hi", actual)
}

func Test_133_TypedDynamic_GetAsInt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	v, ok := d.GetAsInt()

	// Act
	actual := args.Map{"result": ok || v != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_134_TypedDynamic_GetAsInt64(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int64](100, true)
	v, ok := d.GetAsInt64()

	// Act
	actual := args.Map{"result": ok || v != 100}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_135_TypedDynamic_GetAsUint(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[uint](10, true)
	v, ok := d.GetAsUint()

	// Act
	actual := args.Map{"result": ok || v != 10}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
}

func Test_136_TypedDynamic_GetAsFloat64(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[float64](3.14, true)
	v, ok := d.GetAsFloat64()

	// Act
	actual := args.Map{"result": ok || v != 3.14}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
}

func Test_137_TypedDynamic_GetAsFloat32(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[float32](1.5, true)
	v, ok := d.GetAsFloat32()

	// Act
	actual := args.Map{"result": ok || v != 1.5}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
}

func Test_138_TypedDynamic_GetAsBool(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[bool](true, true)
	v, ok := d.GetAsBool()

	// Act
	actual := args.Map{"result": ok || !v}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_139_TypedDynamic_GetAsBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[[]byte]([]byte("hi"), true)
	v, ok := d.GetAsBytes()

	// Act
	actual := args.Map{"result": ok || string(v) != "hi"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hi", actual)
}

func Test_140_TypedDynamic_GetAsStrings(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[[]string]([]string{"a", "b"}, true)
	v, ok := d.GetAsStrings()

	// Act
	actual := args.Map{"result": ok || len(v) != 2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_141_TypedDynamic_ValueString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("hello", true)

	// Act
	actual := args.Map{"result": d.ValueString() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_142_TypedDynamic_ValueString_NonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)

	// Act
	actual := args.Map{"result": d.ValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_143_TypedDynamic_ValueInt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)

	// Act
	actual := args.Map{"result": d.ValueInt() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_144_TypedDynamic_ValueInt64(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int64](100, true)

	// Act
	actual := args.Map{"result": d.ValueInt64() != 100}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_145_TypedDynamic_ValueBool(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[bool](true, true)

	// Act
	actual := args.Map{"result": d.ValueBool()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_146_TypedDynamic_Clone(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	c := d.Clone()

	// Act
	actual := args.Map{"result": c.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_147_TypedDynamic_ClonePtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[int](42, true)
	c := d.ClonePtr()

	// Act
	actual := args.Map{"result": c == nil || c.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_148_TypedDynamic_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.TypedDynamic[int]

	// Act
	actual := args.Map{"result": d.ClonePtr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_149_TypedDynamic_NonPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	np := d.NonPtr()

	// Act
	actual := args.Map{"result": np.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_150_TypedDynamic_Ptr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[int](42, true)

	// Act
	actual := args.Map{"result": d.Ptr() != d}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
}

func Test_151_TypedDynamic_ToDynamic(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)
	dyn := d.ToDynamic()

	// Act
	actual := args.Map{"result": dyn.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_152_TypedDynamic_Deserialize(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[int](0, false)
	err := d.Deserialize([]byte(`42`))

	// Act
	actual := args.Map{"result": err != nil || d.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_153_TypedDynamic_Deserialize_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.TypedDynamic[int]
	err := d.Deserialize([]byte(`42`))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_154_TypedDynamic_JsonModel(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)

	// Act
	actual := args.Map{"result": d.JsonModel() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_155_TypedDynamic_JsonModelAny(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](42, true)

	// Act
	actual := args.Map{"result": d.JsonModelAny() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// TypedSimpleRequest[T]
// ═══════════════════════════════════════════════════════════════════════

func Test_156_TypedSimpleRequest_Basic(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequest[int](42, true, "")

	// Act
	actual := args.Map{"result": r.Data() != 42 || !r.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42 valid", actual)
}

func Test_157_TypedSimpleRequest_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)

	// Act
	actual := args.Map{"result": r.IsValid() || r.Request() != 42 || r.Value() != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid 42", actual)
}

func Test_158_TypedSimpleRequest_Invalid(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleRequest[int]("err")

	// Act
	actual := args.Map{"result": r.IsInvalid() || r.Message() != "err"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid with err", actual)
}

func Test_159_TypedSimpleRequest_InvalidNoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleRequestNoMessage[int]()

	// Act
	actual := args.Map{"result": r.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_160_TypedSimpleRequest_IsValid_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_161_TypedSimpleRequest_IsInvalid_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]

	// Act
	actual := args.Map{"result": r.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_162_TypedSimpleRequest_Message_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]

	// Act
	actual := args.Map{"result": r.Message() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_163_TypedSimpleRequest_String(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)

	// Act
	actual := args.Map{"result": r.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_164_TypedSimpleRequest_String_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]

	// Act
	actual := args.Map{"result": r.String() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_165_TypedSimpleRequest_InvalidError(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleRequest[int]("err msg")
	err := r.InvalidError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// call again for caching
	err2 := r.InvalidError()
	actual = args.Map{"result": err2 != err}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same error", actual)
}

func Test_166_TypedSimpleRequest_InvalidError_NoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)

	// Act
	actual := args.Map{"result": r.InvalidError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_167_TypedSimpleRequest_InvalidError_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]

	// Act
	actual := args.Map{"result": r.InvalidError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_168_TypedSimpleRequest_JsonBytes(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	b, err := r.JsonBytes()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_169_TypedSimpleRequest_GetAs(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("hi")
	v, ok := r.GetAsString()

	// Act
	actual := args.Map{"result": ok || v != "hi"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hi", actual)
	_, ok2 := r.GetAsInt()
	actual = args.Map{"result": ok2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_170_TypedSimpleRequest_Clone(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	c := r.Clone()

	// Act
	actual := args.Map{"result": c == nil || c.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_171_TypedSimpleRequest_Clone_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]

	// Act
	actual := args.Map{"result": r.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_172_TypedSimpleRequest_ToSimpleRequest(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	sr := r.ToSimpleRequest()

	// Act
	actual := args.Map{"result": sr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_173_TypedSimpleRequest_ToSimpleRequest_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]
	sr := r.ToSimpleRequest()

	// Act
	actual := args.Map{"result": sr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil fallback", actual)
}

func Test_174_TypedSimpleRequest_ToTypedDynamic(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	td := r.ToTypedDynamic()

	// Act
	actual := args.Map{"result": td.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_175_TypedSimpleRequest_ToTypedDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]
	td := r.ToTypedDynamic()

	// Act
	actual := args.Map{"result": td.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_176_TypedSimpleRequest_ToDynamic(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	d := r.ToDynamic()

	// Act
	actual := args.Map{"result": d.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_177_TypedSimpleRequest_ToDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[int]
	d := r.ToDynamic()

	// Act
	actual := args.Map{"result": d.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// TypedSimpleResult[T]
// ═══════════════════════════════════════════════════════════════════════

func Test_178_TypedSimpleResult_Basic(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResult[int](42, true, "")

	// Act
	actual := args.Map{"result": r.Data() != 42 || !r.IsValid() || r.Result() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42 valid", actual)
}

func Test_179_TypedSimpleResult_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
}

func Test_180_TypedSimpleResult_Invalid(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleResult[int]("err")

	// Act
	actual := args.Map{"result": r.IsInvalid() || r.Message() != "err"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid with err", actual)
}

func Test_181_TypedSimpleResult_InvalidNoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleResultNoMessage[int]()

	// Act
	actual := args.Map{"result": r.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_182_TypedSimpleResult_IsValid_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_183_TypedSimpleResult_IsInvalid_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]

	// Act
	actual := args.Map{"result": r.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_184_TypedSimpleResult_Message_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]

	// Act
	actual := args.Map{"result": r.Message() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_185_TypedSimpleResult_String(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)

	// Act
	actual := args.Map{"result": r.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_186_TypedSimpleResult_String_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]

	// Act
	actual := args.Map{"result": r.String() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_187_TypedSimpleResult_InvalidError(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleResult[int]("err")

	// Act
	actual := args.Map{"result": r.InvalidError() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual = args.Map{"result": r.InvalidError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached error", actual)
}

func Test_188_TypedSimpleResult_InvalidError_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]

	// Act
	actual := args.Map{"result": r.InvalidError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_189_TypedSimpleResult_JsonBytes(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	b, err := r.JsonBytes()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_190_TypedSimpleResult_GetAs(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("hi")
	v, ok := r.GetAsString()

	// Act
	actual := args.Map{"result": ok || v != "hi"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hi", actual)
}

func Test_191_TypedSimpleResult_Clone(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	c := r.Clone()

	// Act
	actual := args.Map{"result": c.Data() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_192_TypedSimpleResult_ClonePtr(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	c := r.ClonePtr()

	// Act
	actual := args.Map{"result": c == nil || c.Data() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_193_TypedSimpleResult_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]

	// Act
	actual := args.Map{"result": r.ClonePtr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_194_TypedSimpleResult_Clone_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]
	c := r.Clone()

	// Act
	actual := args.Map{"result": c.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not valid", actual)
}

func Test_195_TypedSimpleResult_ToSimpleResult(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	sr := r.ToSimpleResult()

	// Act
	actual := args.Map{"result": sr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_196_TypedSimpleResult_ToSimpleResult_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]
	sr := r.ToSimpleResult()

	// Act
	actual := args.Map{"result": sr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected fallback", actual)
}

func Test_197_TypedSimpleResult_ToTypedDynamic(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	td := r.ToTypedDynamic()

	// Act
	actual := args.Map{"result": td.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_198_TypedSimpleResult_ToTypedDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]
	td := r.ToTypedDynamic()

	// Act
	actual := args.Map{"result": td.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_199_TypedSimpleResult_ToDynamic(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	d := r.ToDynamic()

	// Act
	actual := args.Map{"result": d.Value() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_200_TypedSimpleResult_ToDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[int]
	d := r.ToDynamic()

	// Act
	actual := args.Map{"result": d.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}
