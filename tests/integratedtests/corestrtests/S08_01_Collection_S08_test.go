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

func Test_01_Collection_JsonString(t *testing.T) {
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

func Test_02_Collection_JsonStringMust(t *testing.T) {
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

func Test_03_Collection_HasAnyItem(t *testing.T) {
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

func Test_04_Collection_LastIndex(t *testing.T) {
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

func Test_05_Collection_HasIndex(t *testing.T) {
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

func Test_06_Collection_ListStringsPtr(t *testing.T) {
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

func Test_07_Collection_ListStrings(t *testing.T) {
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

func Test_08_Collection_StringJSON(t *testing.T) {
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

func Test_09_Collection_RemoveAt_Valid(t *testing.T) {
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

func Test_10_Collection_RemoveAt_Invalid(t *testing.T) {
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

func Test_11_Collection_Count(t *testing.T) {
	safeTest(t, "Test_11_Collection_Count", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.Count() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_12_Collection_Capacity(t *testing.T) {
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

func Test_13_Collection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_13_Collection_Length_Nil", func() {
		// Arrange
		var col *corestr.Collection

		// Act & Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil", actual)
	})
}

func Test_14_Collection_LengthLock(t *testing.T) {
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

func Test_15_Collection_IsEquals_SameContent(t *testing.T) {
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

func Test_16_Collection_IsEquals_DiffContent(t *testing.T) {
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

func Test_17_Collection_IsEquals_DiffLength(t *testing.T) {
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

func Test_18_Collection_IsEquals_BothEmpty(t *testing.T) {
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

func Test_19_Collection_IsEquals_OneEmpty(t *testing.T) {
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

func Test_20_Collection_IsEquals_BothNil(t *testing.T) {
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

func Test_21_Collection_IsEquals_OneNil(t *testing.T) {
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

func Test_22_Collection_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_22_Collection_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"x"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for same pointer", actual)
	})
}

func Test_23_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
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

func Test_24_Collection_IsEqualsWithSensitive_CaseInsensitiveNotEqual(t *testing.T) {
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

func Test_25_Collection_IsEmpty(t *testing.T) {
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

func Test_26_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_26_Collection_HasItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": col.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
	})
}

func Test_27_Collection_IsEmptyLock(t *testing.T) {
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

func Test_28_Collection_Add(t *testing.T) {
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

func Test_29_Collection_AddLock(t *testing.T) {
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

func Test_30_Collection_AddNonEmpty(t *testing.T) {
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

func Test_31_Collection_AddNonEmptyWhitespace(t *testing.T) {
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

func Test_32_Collection_AddError(t *testing.T) {
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

func Test_33_Collection_AsDefaultError(t *testing.T) {
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

func Test_34_Collection_AsError_Empty(t *testing.T) {
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

func Test_35_Collection_AddIf_True(t *testing.T) {
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

func Test_36_Collection_AddIfMany(t *testing.T) {
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

func Test_37_Collection_EachItemSplitBy(t *testing.T) {
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

func Test_38_Collection_ConcatNew_WithItems(t *testing.T) {
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

func Test_39_Collection_ConcatNew_Empty(t *testing.T) {
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

func Test_40_Collection_ToError(t *testing.T) {
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

func Test_41_Collection_ToDefaultError(t *testing.T) {
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

func Test_42_Collection_AddFunc(t *testing.T) {
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

func Test_43_Collection_AddFuncErr_Success(t *testing.T) {
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

func Test_44_Collection_AddFuncErr_Error(t *testing.T) {
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

func Test_45_Collection_AddsLock(t *testing.T) {
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

func Test_46_Collection_Adds(t *testing.T) {
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

func Test_47_Collection_AddStrings(t *testing.T) {
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

func Test_48_Collection_AddCollection(t *testing.T) {
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

func Test_49_Collection_AddCollection_Empty(t *testing.T) {
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

func Test_50_Collection_AddCollections(t *testing.T) {
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

func Test_51_Collection_AddPointerCollectionsLock(t *testing.T) {
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

func Test_52_Collection_AddHashmapsValues(t *testing.T) {
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

func Test_53_Collection_AddHashmapsValues_NilAndEmpty(t *testing.T) {
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

func Test_54_Collection_AddHashmapsKeys(t *testing.T) {
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

func Test_55_Collection_AddHashmapsKeys_Nil(t *testing.T) {
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

func Test_56_Collection_AddHashmapsKeysValues(t *testing.T) {
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

func Test_57_Collection_AddHashmapsKeysValues_Nil(t *testing.T) {
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

func Test_58_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
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

func Test_59_Collection_AddHashmapsKeysValuesUsingFilter_Nil(t *testing.T) {
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

func Test_60_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
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

func Test_61_Collection_AddHashmapsKeysValuesUsingFilter_Skip(t *testing.T) {
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

func Test_62_Collection_AddWithWgLock(t *testing.T) {
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

func Test_63_Collection_IndexAt(t *testing.T) {
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

func Test_64_Collection_SafeIndexAtUsingLength(t *testing.T) {
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

func Test_65_Collection_First(t *testing.T) {
	safeTest(t, "Test_65_Collection_First", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act & Assert
		actual := args.Map{"result": col.First() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
	})
}

func Test_66_Collection_Last(t *testing.T) {
	safeTest(t, "Test_66_Collection_Last", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act & Assert
		actual := args.Map{"result": col.Last() != "y"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'y'", actual)
	})
}

func Test_67_Collection_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_67_Collection_FirstOrDefault_NonEmpty", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": col.FirstOrDefault() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_68_Collection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_68_Collection_FirstOrDefault_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	})
}

func Test_69_Collection_LastOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_69_Collection_LastOrDefault_NonEmpty", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.LastOrDefault() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b'", actual)
	})
}

func Test_70_Collection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_70_Collection_LastOrDefault_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	})
}

func Test_71_Collection_Single(t *testing.T) {
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

func Test_72_Collection_Single_Panic(t *testing.T) {
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

func Test_73_Collection_Take(t *testing.T) {
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

func Test_74_Collection_Take_MoreThanLength(t *testing.T) {
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

func Test_75_Collection_Take_Zero(t *testing.T) {
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

func Test_76_Collection_Skip(t *testing.T) {
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

func Test_77_Collection_Skip_Zero(t *testing.T) {
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

func Test_78_Collection_Reverse(t *testing.T) {
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

func Test_79_Collection_Reverse_Two(t *testing.T) {
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

func Test_80_Collection_Reverse_Single(t *testing.T) {
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

func Test_81_Collection_GetPagesSize(t *testing.T) {
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

func Test_82_Collection_GetSinglePageCollection(t *testing.T) {
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

func Test_83_Collection_GetSinglePageCollection_LastPage(t *testing.T) {
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

func Test_84_Collection_GetSinglePageCollection_SmallerThanPageSize(t *testing.T) {
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

func Test_85_Collection_GetPagedCollection(t *testing.T) {
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

func Test_86_Collection_GetPagedCollection_SmallerThanPageSize(t *testing.T) {
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

func Test_87_Collection_ResizeForItems_IndirectViaAppendAnys(t *testing.T) {
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

func Test_88_Collection_AddStringsAsync_Empty(t *testing.T) {
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

func Test_89_Collection_InsertAt_AtEnd(t *testing.T) {
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

func Test_90_Collection_InsertAt_Empty(t *testing.T) {
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

func Test_91_Collection_ChainRemoveAt(t *testing.T) {
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

func Test_92_Collection_RemoveItemsIndexes(t *testing.T) {
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

func Test_93_Collection_RemoveItemsIndexes_NilIndexes(t *testing.T) {
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

func Test_94_Collection_RemoveItemsIndexesPtr_NilIndexes(t *testing.T) {
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

func Test_95_Collection_RemoveItemsIndexesPtr_EmptyCollValidation(t *testing.T) {
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

func Test_96_Collection_RemoveItemsIndexesPtr_EmptyCollIgnore(t *testing.T) {
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

func Test_97_Collection_AppendCollectionPtr(t *testing.T) {
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

func Test_98_Collection_AppendCollections(t *testing.T) {
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

func Test_99_Collection_AppendCollections_Empty(t *testing.T) {
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
