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
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S08 — Collection Part 1: Core methods (lines 1–700)
// ══════════════════════════════════════════════════════════════

// ── JsonString / JsonStringMust ─────────────────────────────

func Test_Collection_01_Collection_JsonString_FromS08(t *testing.T) {
	safeTest(t, "Test_01_Collection_JsonString", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		result := col.JsonString()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty JSON string", actual)
	})
}

func Test_Collection_02_Collection_JsonStringMust_FromS08(t *testing.T) {
	safeTest(t, "Test_02_Collection_JsonStringMust", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x"})

		// Act
		result := col.JsonStringMust()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty JSON string", actual)
	})
}

// ── HasAnyItem / LastIndex / HasIndex ────────────────────────

func Test_Collection_03_Collection_HasAnyItem_FromS08(t *testing.T) {
	safeTest(t, "Test_03_Collection_HasAnyItem", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		empty := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasAnyItem true", actual)
		actual = args.Map{"result": empty.HasAnyItem()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasAnyItem false for empty", actual)
	})
}

func Test_Collection_04_Collection_LastIndex_FromS08(t *testing.T) {
	safeTest(t, "Test_04_Collection_LastIndex", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		li := col.LastIndex()

		// Assert
		actual := args.Map{"result": li != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_05_Collection_HasIndex_FromS08(t *testing.T) {
	safeTest(t, "Test_05_Collection_HasIndex", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(0) true", actual)
		actual = args.Map{"result": col.HasIndex(1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(1) true", actual)
		actual = args.Map{"result": col.HasIndex(2)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(2) false", actual)
		actual = args.Map{"result": col.HasIndex(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasIndex(-1) false", actual)
	})
}

// ── ListStringsPtr / ListStrings / StringJSON ────────────────

func Test_Collection_06_Collection_ListStringsPtr_FromS08(t *testing.T) {
	safeTest(t, "Test_06_Collection_ListStringsPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act
		items := col.ListStringsPtr()

		// Assert
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
	})
}

func Test_Collection_07_Collection_ListStrings_FromS08(t *testing.T) {
	safeTest(t, "Test_07_Collection_ListStrings", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		items := col.ListStrings()

		// Assert
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 item", actual)
	})
}

func Test_Collection_08_Collection_StringJSON_FromS08(t *testing.T) {
	safeTest(t, "Test_08_Collection_StringJSON", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		s := col.StringJSON()

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

// ── RemoveAt ────────────────────────────────────────────────

func Test_Collection_09_Collection_RemoveAt_Valid_FromS08(t *testing.T) {
	safeTest(t, "Test_09_Collection_RemoveAt_Valid", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		ok := col.RemoveAt(1)

		// Assert
		actual := args.Map{"result": ok}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_10_Collection_RemoveAt_Invalid_FromS08(t *testing.T) {
	safeTest(t, "Test_10_Collection_RemoveAt_Invalid", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": col.RemoveAt(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for negative index", actual)
		actual = args.Map{"result": col.RemoveAt(5)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for out of range", actual)
	})
}

// ── Count / Capacity / Length / LengthLock ───────────────────

func Test_Collection_11_Collection_Count_FromS08(t *testing.T) {
	safeTest(t, "Test_11_Collection_Count", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.Count() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_12_Collection_Capacity_FromS08(t *testing.T) {
	safeTest(t, "Test_12_Collection_Capacity", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)
		empty := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.Capacity() < 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 10", actual)
		_ = empty.Capacity() // just exercise nil items path
	})
}

func Test_Collection_13_Collection_Length_Nil_FromS08(t *testing.T) {
	safeTest(t, "Test_13_Collection_Length_Nil", func() {
		// Arrange
		var col *corestr.Collection

		// Act & Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil", actual)
	})
}

func Test_Collection_14_Collection_LengthLock_FromS08(t *testing.T) {
	safeTest(t, "Test_14_Collection_LengthLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		l := col.LengthLock()

		// Assert
		actual := args.Map{"result": l != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

// ── IsEquals / IsEqualsWithSensitive ─────────────────────────

func Test_Collection_15_Collection_IsEquals_SameContent_FromS08(t *testing.T) {
	safeTest(t, "Test_15_Collection_IsEquals_SameContent", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x", "y"})
		b := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Collection_16_Collection_IsEquals_DiffContent_FromS08(t *testing.T) {
	safeTest(t, "Test_16_Collection_IsEquals_DiffContent", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x"})
		b := corestr.New.Collection.Strings([]string{"y"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Collection_17_Collection_IsEquals_DiffLength_FromS08(t *testing.T) {
	safeTest(t, "Test_17_Collection_IsEquals_DiffLength", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x"})
		b := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal on diff length", actual)
	})
}

func Test_Collection_18_Collection_IsEquals_BothEmpty_FromS08(t *testing.T) {
	safeTest(t, "Test_18_Collection_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.Collection()
		b := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for both empty", actual)
	})
}

func Test_Collection_19_Collection_IsEquals_OneEmpty_FromS08(t *testing.T) {
	safeTest(t, "Test_19_Collection_IsEquals_OneEmpty", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x"})
		b := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal when one empty", actual)
	})
}

func Test_Collection_20_Collection_IsEquals_BothNil_FromS08(t *testing.T) {
	safeTest(t, "Test_20_Collection_IsEquals_BothNil", func() {
		// Arrange
		var a *corestr.Collection
		var b *corestr.Collection

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for both nil", actual)
	})
}

func Test_Collection_21_Collection_IsEquals_OneNil_FromS08(t *testing.T) {
	safeTest(t, "Test_21_Collection_IsEquals_OneNil", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x"})
		var b *corestr.Collection

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal when one nil", actual)
	})
}

func Test_Collection_22_Collection_IsEquals_SamePtr_FromS08(t *testing.T) {
	safeTest(t, "Test_22_Collection_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for same pointer", actual)
	})
}

func Test_Collection_23_Collection_IsEqualsWithSensitive_CaseInsensitive_FromS08(t *testing.T) {
	safeTest(t, "Test_23_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello", "WORLD"})
		b := corestr.New.Collection.Strings([]string{"hello", "world"})

		// Act & Assert
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected case-insensitive equal", actual)
		actual = args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected case-sensitive not equal", actual)
	})
}

func Test_Collection_24_Collection_IsEqualsWithSensitive_CaseInsensitiveNotEqual_FromS08(t *testing.T) {
	safeTest(t, "Test_24_Collection_IsEqualsWithSensitive_CaseInsensitiveNotEqual", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"Goodbye"})

		// Act & Assert
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal even case-insensitive", actual)
	})
}

// ── IsEmpty / HasItems / IsEmptyLock ─────────────────────────

func Test_Collection_25_Collection_IsEmpty_FromS08(t *testing.T) {
	safeTest(t, "Test_25_Collection_IsEmpty", func() {
		// Arrange
		empty := corestr.Empty.Collection()
		nonEmpty := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": nonEmpty.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_Collection_26_Collection_HasItems_FromS08(t *testing.T) {
	safeTest(t, "Test_26_Collection_HasItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": col.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
	})
}

func Test_Collection_27_Collection_IsEmptyLock_FromS08(t *testing.T) {
	safeTest(t, "Test_27_Collection_IsEmptyLock", func() {
		// Arrange
		empty := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": empty.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty lock", actual)
	})
}

// ── Add / AddLock / AddNonEmpty / AddNonEmptyWhitespace ──────

func Test_Collection_28_Collection_Add_FromS08(t *testing.T) {
	safeTest(t, "Test_28_Collection_Add", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.Add("hello")

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_29_Collection_AddLock_FromS08(t *testing.T) {
	safeTest(t, "Test_29_Collection_AddLock", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddLock("a")
		col.AddLock("b")

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_30_Collection_AddNonEmpty_FromS08(t *testing.T) {
	safeTest(t, "Test_30_Collection_AddNonEmpty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddNonEmpty("")
		col.AddNonEmpty("hello")

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1, empty should be skipped", actual)
	})
}

func Test_Collection_31_Collection_AddNonEmptyWhitespace_FromS08(t *testing.T) {
	safeTest(t, "Test_31_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddNonEmptyWhitespace("  ")
		col.AddNonEmptyWhitespace("")
		col.AddNonEmptyWhitespace("ok")

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AddError ─────────────────────────────────────────────────

func Test_Collection_32_Collection_AddError_FromS08(t *testing.T) {
	safeTest(t, "Test_32_Collection_AddError", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddError(nil)
		col.AddError(errors.New("test err"))

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1, nil error should be skipped", actual)
		actual = args.Map{"result": col.First() != "test err"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'test err', got ''", actual)
	})
}

// ── AsDefaultError / AsError ─────────────────────────────────

func Test_Collection_33_Collection_AsDefaultError_FromS08(t *testing.T) {
	safeTest(t, "Test_33_Collection_AsDefaultError", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"err1", "err2"})

		// Act
		err := col.AsDefaultError()

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Collection_34_Collection_AsError_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_34_Collection_AsError_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		err := col.AsError(",")

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil error for empty collection", actual)
	})
}

// ── AddIf / AddIfMany ────────────────────────────────────────

func Test_Collection_35_Collection_AddIf_True_FromS08(t *testing.T) {
	safeTest(t, "Test_35_Collection_AddIf_True", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddIf(true, "yes")
		col.AddIf(false, "no")

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_36_Collection_AddIfMany_FromS08(t *testing.T) {
	safeTest(t, "Test_36_Collection_AddIfMany", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddIfMany(true, "a", "b")
		col.AddIfMany(false, "c", "d")

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── EachItemSplitBy ──────────────────────────────────────────

func Test_Collection_37_Collection_EachItemSplitBy_FromS08(t *testing.T) {
	safeTest(t, "Test_37_Collection_EachItemSplitBy", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a,b", "c,d"})

		// Act
		result := col.EachItemSplitBy(",")

		// Assert
		actual := args.Map{"result": len(result) != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

// ── ConcatNew ────────────────────────────────────────────────

func Test_Collection_38_Collection_ConcatNew_WithItems_FromS08(t *testing.T) {
	safeTest(t, "Test_38_Collection_ConcatNew_WithItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		newCol := col.ConcatNew(0, "b", "c")

		// Assert
		actual := args.Map{"result": newCol.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_39_Collection_ConcatNew_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_39_Collection_ConcatNew_Empty", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		newCol := col.ConcatNew(0)

		// Assert
		actual := args.Map{"result": newCol.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── ToError / ToDefaultError ─────────────────────────────────

func Test_Collection_40_Collection_ToError_FromS08(t *testing.T) {
	safeTest(t, "Test_40_Collection_ToError", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"err1"})

		// Act
		err := col.ToError(",")

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Collection_41_Collection_ToDefaultError_FromS08(t *testing.T) {
	safeTest(t, "Test_41_Collection_ToDefaultError", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"e1", "e2"})

		// Act
		err := col.ToDefaultError()

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

// ── AddFunc / AddFuncErr ─────────────────────────────────────

func Test_Collection_42_Collection_AddFunc_FromS08(t *testing.T) {
	safeTest(t, "Test_42_Collection_AddFunc", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddFunc(func() string { return "generated" })

		// Assert
		actual := args.Map{"result": col.Length() != 1 || col.First() != "generated"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'generated'", actual)
	})
}

func Test_Collection_43_Collection_AddFuncErr_Success_FromS08(t *testing.T) {
	safeTest(t, "Test_43_Collection_AddFuncErr_Success", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_44_Collection_AddFuncErr_Error_FromS08(t *testing.T) {
	safeTest(t, "Test_44_Collection_AddFuncErr_Error", func() {
		// Arrange
		col := corestr.Empty.Collection()
		called := false

		// Act
		col.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { called = true },
		)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": called}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected error handler called", actual)
	})
}

// ── AddsLock / Adds / AddStrings ─────────────────────────────

func Test_Collection_45_Collection_AddsLock_FromS08(t *testing.T) {
	safeTest(t, "Test_45_Collection_AddsLock", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddsLock("a", "b")

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_46_Collection_Adds_FromS08(t *testing.T) {
	safeTest(t, "Test_46_Collection_Adds", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.Adds("x", "y", "z")

		// Assert
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_47_Collection_AddStrings_FromS08(t *testing.T) {
	safeTest(t, "Test_47_Collection_AddStrings", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddStrings([]string{"a", "b"})

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── AddCollection / AddCollections ───────────────────────────

func Test_Collection_48_Collection_AddCollection_FromS08(t *testing.T) {
	safeTest(t, "Test_48_Collection_AddCollection", func() {
		// Arrange
		col := corestr.Empty.Collection()
		other := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		col.AddCollection(other)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_49_Collection_AddCollection_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_49_Collection_AddCollection_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		other := corestr.Empty.Collection()

		// Act
		col.AddCollection(other)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_50_Collection_AddCollections_FromS08(t *testing.T) {
	safeTest(t, "Test_50_Collection_AddCollections", func() {
		// Arrange
		col := corestr.Empty.Collection()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b", "c"})
		empty := corestr.Empty.Collection()

		// Act
		col.AddCollections(c1, empty, c2)

		// Assert
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

// ── AddPointerCollectionsLock ────────────────────────────────

func Test_Collection_51_Collection_AddPointerCollectionsLock_FromS08(t *testing.T) {
	safeTest(t, "Test_51_Collection_AddPointerCollectionsLock", func() {
		// Arrange
		col := corestr.Empty.Collection()
		c1 := corestr.New.Collection.Strings([]string{"x"})

		// Act
		col.AddPointerCollectionsLock(c1)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AddHashmapsValues / AddHashmapsKeys / AddHashmapsKeysValues ──

func Test_Collection_52_Collection_AddHashmapsValues_FromS08(t *testing.T) {
	safeTest(t, "Test_52_Collection_AddHashmapsValues", func() {
		// Arrange
		col := corestr.Empty.Collection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdate("k2", "v2")

		// Act
		col.AddHashmapsValues(hm)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_53_Collection_AddHashmapsValues_NilAndEmpty_FromS08(t *testing.T) {
	safeTest(t, "Test_53_Collection_AddHashmapsValues_NilAndEmpty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddHashmapsValues(nil)
		col.AddHashmapsValues(corestr.Empty.Hashmap())

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_54_Collection_AddHashmapsKeys_FromS08(t *testing.T) {
	safeTest(t, "Test_54_Collection_AddHashmapsKeys", func() {
		// Arrange
		col := corestr.Empty.Collection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")

		// Act
		col.AddHashmapsKeys(hm)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_55_Collection_AddHashmapsKeys_Nil_FromS08(t *testing.T) {
	safeTest(t, "Test_55_Collection_AddHashmapsKeys_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddHashmapsKeys(nil)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_56_Collection_AddHashmapsKeysValues_FromS08(t *testing.T) {
	safeTest(t, "Test_56_Collection_AddHashmapsKeysValues", func() {
		// Arrange
		col := corestr.Empty.Collection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")

		// Act
		col.AddHashmapsKeysValues(hm)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_57_Collection_AddHashmapsKeysValues_Nil_FromS08(t *testing.T) {
	safeTest(t, "Test_57_Collection_AddHashmapsKeysValues_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddHashmapsKeysValues(nil)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddHashmapsKeysValuesUsingFilter ─────────────────────────

func Test_Collection_58_Collection_AddHashmapsKeysValuesUsingFilter_FromS08(t *testing.T) {
	safeTest(t, "Test_58_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		// Arrange
		col := corestr.Empty.Collection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdate("k2", "v2")

		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key + "=" + pair.Value, true, false
		}

		// Act
		col.AddHashmapsKeysValuesUsingFilter(filter, hm)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_59_Collection_AddHashmapsKeysValuesUsingFilter_Nil_FromS08(t *testing.T) {
	safeTest(t, "Test_59_Collection_AddHashmapsKeysValuesUsingFilter_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddHashmapsKeysValuesUsingFilter(nil, nil)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_60_Collection_AddHashmapsKeysValuesUsingFilter_Break_FromS08(t *testing.T) {
	safeTest(t, "Test_60_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		// Arrange
		col := corestr.Empty.Collection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdate("k2", "v2")
		hm.AddOrUpdate("k3", "v3")

		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key, true, true // break immediately after first
		}

		// Act
		col.AddHashmapsKeysValuesUsingFilter(filter, hm)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (break after first)", actual)
	})
}

func Test_Collection_61_Collection_AddHashmapsKeysValuesUsingFilter_Skip_FromS08(t *testing.T) {
	safeTest(t, "Test_61_Collection_AddHashmapsKeysValuesUsingFilter_Skip", func() {
		// Arrange
		col := corestr.Empty.Collection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")

		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return "", false, false // skip all
		}

		// Act
		col.AddHashmapsKeysValuesUsingFilter(filter, hm)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddWithWgLock ────────────────────────────────────────────

func Test_Collection_62_Collection_AddWithWgLock_FromS08(t *testing.T) {
	safeTest(t, "Test_62_Collection_AddWithWgLock", func() {
		// Arrange
		col := corestr.Empty.Collection()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		col.AddWithWgLock(wg, "item")
		wg.Wait()

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── IndexAt / SafeIndexAtUsingLength ─────────────────────────

func Test_Collection_63_Collection_IndexAt_FromS08(t *testing.T) {
	safeTest(t, "Test_63_Collection_IndexAt", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": col.IndexAt(0) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
		actual = args.Map{"result": col.IndexAt(2) != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'c'", actual)
	})
}

func Test_Collection_64_Collection_SafeIndexAtUsingLength_FromS08(t *testing.T) {
	safeTest(t, "Test_64_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		v1 := col.SafeIndexAtUsingLength("default", 2, 0)
		v2 := col.SafeIndexAtUsingLength("default", 2, 5)

		// Assert
		actual := args.Map{"result": v1 != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
		actual = args.Map{"result": v2 != "default"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'default', got ''", actual)
	})
}

// ── First / Last / FirstOrDefault / LastOrDefault / Single ───

func Test_Collection_65_Collection_First_FromS08(t *testing.T) {
	safeTest(t, "Test_65_Collection_First", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act & Assert
		actual := args.Map{"result": col.First() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
	})
}

func Test_Collection_66_Collection_Last_FromS08(t *testing.T) {
	safeTest(t, "Test_66_Collection_Last", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act & Assert
		actual := args.Map{"result": col.Last() != "y"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'y'", actual)
	})
}

func Test_Collection_67_Collection_FirstOrDefault_NonEmpty_FromS08(t *testing.T) {
	safeTest(t, "Test_67_Collection_FirstOrDefault_NonEmpty", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": col.FirstOrDefault() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_Collection_68_Collection_FirstOrDefault_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_68_Collection_FirstOrDefault_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	})
}

func Test_Collection_69_Collection_LastOrDefault_NonEmpty_FromS08(t *testing.T) {
	safeTest(t, "Test_69_Collection_LastOrDefault_NonEmpty", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.LastOrDefault() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b'", actual)
	})
}

func Test_Collection_70_Collection_LastOrDefault_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_70_Collection_LastOrDefault_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	})
}

func Test_Collection_71_Collection_Single_FromS08(t *testing.T) {
	safeTest(t, "Test_71_Collection_Single", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"only"})

		// Act
		s := col.Single()

		// Assert
		actual := args.Map{"result": s != "only"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'only'", actual)
	})
}

func Test_Collection_72_Collection_Single_Panic_FromS08(t *testing.T) {
	safeTest(t, "Test_72_Collection_Single_Panic", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		defer func() {
			r := recover()
			actual := args.Map{"result": r == nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected panic for non-single collection", actual)
		}()
		col.Single()
	})
}

// ── Take / Skip ──────────────────────────────────────────────

func Test_Collection_73_Collection_Take_FromS08(t *testing.T) {
	safeTest(t, "Test_73_Collection_Take", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})

		// Act
		taken := col.Take(2)

		// Assert
		actual := args.Map{"result": taken.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_74_Collection_Take_MoreThanLength_FromS08(t *testing.T) {
	safeTest(t, "Test_74_Collection_Take_MoreThanLength", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		taken := col.Take(5)

		// Assert
		actual := args.Map{"result": taken.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected original collection", actual)
	})
}

func Test_Collection_75_Collection_Take_Zero_FromS08(t *testing.T) {
	safeTest(t, "Test_75_Collection_Take_Zero", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		taken := col.Take(0)

		// Assert
		actual := args.Map{"result": taken.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_76_Collection_Skip_FromS08(t *testing.T) {
	safeTest(t, "Test_76_Collection_Skip", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		skipped := col.Skip(1)

		// Assert
		actual := args.Map{"result": skipped.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_77_Collection_Skip_Zero_FromS08(t *testing.T) {
	safeTest(t, "Test_77_Collection_Skip_Zero", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		skipped := col.Skip(0)

		// Assert
		actual := args.Map{"result": skipped.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same collection", actual)
	})
}

// ── Reverse ──────────────────────────────────────────────────

func Test_Collection_78_Collection_Reverse_FromS08(t *testing.T) {
	safeTest(t, "Test_78_Collection_Reverse", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		col.Reverse()

		// Assert
		actual := args.Map{"result": col.First() != "c" || col.Last() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	})
}

func Test_Collection_79_Collection_Reverse_Two_FromS08(t *testing.T) {
	safeTest(t, "Test_79_Collection_Reverse_Two", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		col.Reverse()

		// Assert
		actual := args.Map{"result": col.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b' first", actual)
	})
}

func Test_Collection_80_Collection_Reverse_Single_FromS08(t *testing.T) {
	safeTest(t, "Test_80_Collection_Reverse_Single", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		col.Reverse()

		// Assert
		actual := args.Map{"result": col.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' unchanged", actual)
	})
}

// ── GetPagesSize / GetPagedCollection / GetSinglePageCollection ──

func Test_Collection_81_Collection_GetPagesSize_FromS08(t *testing.T) {
	safeTest(t, "Test_81_Collection_GetPagesSize", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act & Assert
		actual := args.Map{"result": col.GetPagesSize(2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
		actual = args.Map{"result": col.GetPagesSize(0) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for zero page size", actual)
		actual = args.Map{"result": col.GetPagesSize(-1) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for negative", actual)
	})
}

func Test_Collection_82_Collection_GetSinglePageCollection_FromS08(t *testing.T) {
	safeTest(t, "Test_82_Collection_GetSinglePageCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act
		page := col.GetSinglePageCollection(2, 2) // page 2 of 2-item pages

		// Assert
		actual := args.Map{"result": page.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_83_Collection_GetSinglePageCollection_LastPage_FromS08(t *testing.T) {
	safeTest(t, "Test_83_Collection_GetSinglePageCollection_LastPage", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act
		page := col.GetSinglePageCollection(2, 3) // page 3 of 2-item pages = 1 item

		// Assert
		actual := args.Map{"result": page.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_84_Collection_GetSinglePageCollection_SmallerThanPageSize_FromS08(t *testing.T) {
	safeTest(t, "Test_84_Collection_GetSinglePageCollection_SmallerThanPageSize", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		page := col.GetSinglePageCollection(10, 1)

		// Assert
		actual := args.Map{"result": page.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_85_Collection_GetPagedCollection_FromS08(t *testing.T) {
	safeTest(t, "Test_85_Collection_GetPagedCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act
		paged := col.GetPagedCollection(2)

		// Assert
		actual := args.Map{"result": paged.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
	})
}

func Test_Collection_86_Collection_GetPagedCollection_SmallerThanPageSize_FromS08(t *testing.T) {
	safeTest(t, "Test_86_Collection_GetPagedCollection_SmallerThanPageSize", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		paged := col.GetPagedCollection(10)

		// Assert
		actual := args.Map{"result": paged.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── resizeForItems / resizeForAnys / isResizeRequired ────────
// These are private methods exercised indirectly through AppendAnys etc.

func Test_Collection_87_Collection_ResizeForItems_IndirectViaAppendAnys_FromS08(t *testing.T) {
	safeTest(t, "Test_87_Collection_ResizeForItems_IndirectViaAppendAnys", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)

		// Build a large slice of anys to trigger resize
		items := make([]any, 250)
		for i := range items {
			items[i] = "item"
		}

		// Act
		col.AppendAnys(items...)

		// Assert
		actual := args.Map{"result": col.Length() != 250}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 250", actual)
	})
}

// ── AddStringsAsync ──────────────────────────────────────────

func Test_Collection_88_Collection_AddStringsAsync_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_88_Collection_AddStringsAsync_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		wg := &sync.WaitGroup{}

		// Act
		col.AddStringsAsync(wg, []string{})

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── InsertAt ─────────────────────────────────────────────────

func Test_Collection_89_Collection_InsertAt_AtEnd_FromS08(t *testing.T) {
	safeTest(t, "Test_89_Collection_InsertAt_AtEnd", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		col.InsertAt(1, "c") // at last index

		// Assert
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_90_Collection_InsertAt_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_90_Collection_InsertAt_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.InsertAt(0, "a")

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── ChainRemoveAt ────────────────────────────────────────────

func Test_Collection_91_Collection_ChainRemoveAt_FromS08(t *testing.T) {
	safeTest(t, "Test_91_Collection_ChainRemoveAt", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		result := col.ChainRemoveAt(1)

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── RemoveItemsIndexes / RemoveItemsIndexesPtr ───────────────

func Test_Collection_92_Collection_RemoveItemsIndexes_FromS08(t *testing.T) {
	safeTest(t, "Test_92_Collection_RemoveItemsIndexes", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})

		// Act
		col.RemoveItemsIndexes(true, 1, 3)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_93_Collection_RemoveItemsIndexes_NilIndexes_FromS08(t *testing.T) {
	safeTest(t, "Test_93_Collection_RemoveItemsIndexes_NilIndexes", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		col.RemoveItemsIndexes(true)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_94_Collection_RemoveItemsIndexesPtr_NilIndexes_FromS08(t *testing.T) {
	safeTest(t, "Test_94_Collection_RemoveItemsIndexesPtr_NilIndexes", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		col.RemoveItemsIndexesPtr(false, nil)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_95_Collection_RemoveItemsIndexesPtr_EmptyCollValidation_FromS08(t *testing.T) {
	safeTest(t, "Test_95_Collection_RemoveItemsIndexesPtr_EmptyCollValidation", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert — should panic with validation on
		defer func() {
			r := recover()
			actual := args.Map{"result": r == nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected panic for removing from empty with validation", actual)
		}()
		col.RemoveItemsIndexesPtr(false, []int{0})
	})
}

func Test_Collection_96_Collection_RemoveItemsIndexesPtr_EmptyCollIgnore_FromS08(t *testing.T) {
	safeTest(t, "Test_96_Collection_RemoveItemsIndexesPtr_EmptyCollIgnore", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.RemoveItemsIndexesPtr(true, []int{0})

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AppendCollectionPtr / AppendCollections ───────────────────

func Test_Collection_97_Collection_AppendCollectionPtr_FromS08(t *testing.T) {
	safeTest(t, "Test_97_Collection_AppendCollectionPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b", "c"})

		// Act
		col.AppendCollectionPtr(other)

		// Assert
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_98_Collection_AppendCollections_FromS08(t *testing.T) {
	safeTest(t, "Test_98_Collection_AppendCollections", func() {
		// Arrange
		col := corestr.Empty.Collection()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		// Act
		col.AppendCollections(c1, c2)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_99_Collection_AppendCollections_Empty_FromS08(t *testing.T) {
	safeTest(t, "Test_99_Collection_AppendCollections_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendCollections()

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
