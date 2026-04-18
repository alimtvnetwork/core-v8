package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S08b — Collection Part 2: Lines 700–1600
// ══════════════════════════════════════════════════════════════

// ── AppendAnysLock / AppendAnys ──────────────────────────────

func Test_Collection_AppendAnysLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_01_Collection_AppendAnysLock", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysLock("a", 42, nil)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_02_Collection_AppendAnysLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysLock()

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_AppendAnys_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_03_Collection_AppendAnys", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnys("hello", 123)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b_v2(t *testing.T) {
	safeTest(t, "Test_S08b_04_Collection_AppendAnys_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnys()

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_WithNil_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_05_Collection_AppendAnys_WithNil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnys(nil, "ok", nil)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AppendAnysUsingFilter / AppendAnysUsingFilterLock ────────

func Test_Collection_AppendAnysUsingFilter_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_06_Collection_AppendAnysUsingFilter", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return strings.ToUpper(str), true, false
		}

		// Act
		col.AppendAnysUsingFilter(filter, "hello", nil)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_Empty_FromS08b_v3(t *testing.T) {
	safeTest(t, "Test_S08b_07_Collection_AppendAnysUsingFilter_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysUsingFilter(nil)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Skip_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_08_Collection_AppendAnysUsingFilter_Skip", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return "", false, false // skip
		}

		// Act
		col.AppendAnysUsingFilter(filter, "a")

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Break_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_09_Collection_AppendAnysUsingFilter_Break", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true // keep and break
		}

		// Act
		col.AppendAnysUsingFilter(filter, "a", "b", "c")

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_10_Collection_AppendAnysUsingFilterLock", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		col.AppendAnysUsingFilterLock(filter, "x", nil, "y")

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Nil_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_11_Collection_AppendAnysUsingFilterLock_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysUsingFilterLock(nil, nil)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_SkipAndBreak_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_12_Collection_AppendAnysUsingFilterLock_SkipAndBreak", func() {
		// Arrange
		col := corestr.Empty.Collection()
		callCount := 0
		filter := func(str string, index int) (string, bool, bool) {
			callCount++
			if callCount == 1 {
				return "", false, false // skip
			}
			return str, true, true // keep and break
		}

		// Act
		col.AppendAnysUsingFilterLock(filter, "a", "b", "c")

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AppendNonEmptyAnys ───────────────────────────────────────

func Test_Collection_AppendNonEmptyAnys_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_13_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendNonEmptyAnys("hello", nil, "", "world")

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Nil_FromS08b_v2(t *testing.T) {
	safeTest(t, "Test_S08b_14_Collection_AppendNonEmptyAnys_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendNonEmptyAnys(nil)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddsAsync ────────────────────────────────────────────────

func Test_Collection_AddsAsync_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_15_Collection_AddsAsync", func() {
		// Arrange
		col := corestr.Empty.Collection()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		col.AddsAsync(wg, "a", "b")
		wg.Wait()

		// Assert
		actual := args.Map{"result": col.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected items added", actual)
	})
}

func Test_Collection_Nil_FromS08b_v3(t *testing.T) {
	safeTest(t, "Test_S08b_16_Collection_AddsAsync_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()
		wg := &sync.WaitGroup{}

		// Act
		col.AddsAsync(wg, nil...)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddsNonEmpty ─────────────────────────────────────────────

func Test_Collection_AddsNonEmpty_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_17_Collection_AddsNonEmpty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddsNonEmpty("a", "", "b", "")

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Nil_FromS08b_v4(t *testing.T) {
	safeTest(t, "Test_S08b_18_Collection_AddsNonEmpty_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddsNonEmpty(nil...)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddsNonEmptyPtrLock ──────────────────────────────────────

func Test_Collection_AddsNonEmptyPtrLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_19_Collection_AddsNonEmptyPtrLock", func() {
		// Arrange
		col := corestr.Empty.Collection()
		s1 := "hello"
		s2 := ""

		// Act
		col.AddsNonEmptyPtrLock(&s1, nil, &s2)

		// Assert
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_Nil_FromS08b_v5(t *testing.T) {
	safeTest(t, "Test_S08b_20_Collection_AddsNonEmptyPtrLock_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddsNonEmptyPtrLock(nil)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── UniqueBoolMap / UniqueList / UniqueBoolMapLock / UniqueListLock ──

func Test_Collection_UniqueBoolMap_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_21_Collection_UniqueBoolMap", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		m := col.UniqueBoolMap()

		// Assert
		actual := args.Map{"result": len(m) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_UniqueBoolMapLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_22_Collection_UniqueBoolMapLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "x"})

		// Act
		m := col.UniqueBoolMapLock()

		// Assert
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_UniqueList_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_23_Collection_UniqueList", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		list := col.UniqueList()

		// Assert
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_UniqueListLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_24_Collection_UniqueListLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "a"})

		// Act
		list := col.UniqueListLock()

		// Assert
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── List ─────────────────────────────────────────────────────

func Test_Collection_List_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_25_Collection_List", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		list := col.List()

		// Assert
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── Filter / FilterLock / FilteredCollection / FilteredCollectionLock ──

func Test_Collection_Filter_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_26_Collection_Filter", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"apple", "banana", "avocado"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, strings.HasPrefix(str, "a"), false
		}

		// Act
		result := col.Filter(filter)

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b_v4(t *testing.T) {
	safeTest(t, "Test_S08b_27_Collection_Filter_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		result := col.Filter(func(str string, index int) (string, bool, bool) {
			return str, true, false
		})

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Break_FromS08b_v2(t *testing.T) {
	safeTest(t, "Test_S08b_28_Collection_Filter_Break", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, index == 0 // break after first
		}

		// Act
		result := col.Filter(filter)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_FilterLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_29_Collection_FilterLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := col.FilterLock(filter)

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Break_FromS08b_v3(t *testing.T) {
	safeTest(t, "Test_S08b_30_Collection_FilterLock_Break", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true // keep and break
		}

		// Act
		result := col.FilterLock(filter)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_FilteredCollection_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_31_Collection_FilteredCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y", "z"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, str != "y", false
		}

		// Act
		result := col.FilteredCollection(filter)

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_FilteredCollectionLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_32_Collection_FilteredCollectionLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := col.FilteredCollectionLock(filter)

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── FilterPtrLock / FilterPtr ────────────────────────────────

func Test_Collection_FilterPtrLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_33_Collection_FilterPtrLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtrLock(filter)

		// Assert
		actual := args.Map{"result": len(*result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b_v5(t *testing.T) {
	safeTest(t, "Test_S08b_34_Collection_FilterPtrLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtrLock(filter)

		// Assert
		actual := args.Map{"result": len(*result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Break_FromS08b_v4(t *testing.T) {
	safeTest(t, "Test_S08b_35_Collection_FilterPtrLock_Break", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, true
		}

		// Act
		result := col.FilterPtrLock(filter)

		// Assert
		actual := args.Map{"result": len(*result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_FilterPtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_36_Collection_FilterPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtr(filter)

		// Assert
		actual := args.Map{"result": len(*result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b_v6(t *testing.T) {
	safeTest(t, "Test_S08b_37_Collection_FilterPtr_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtr(filter)

		// Assert
		actual := args.Map{"result": len(*result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Break_FromS08b_v5(t *testing.T) {
	safeTest(t, "Test_S08b_38_Collection_FilterPtr_Break", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, true
		}

		// Act
		result := col.FilterPtr(filter)

		// Assert
		actual := args.Map{"result": len(*result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── NonEmptyListPtr / NonEmptyList ───────────────────────────

func Test_Collection_NonEmptyList_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_39_Collection_NonEmptyList", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "", "b", ""})

		// Act
		list := col.NonEmptyList()

		// Assert
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b_v7(t *testing.T) {
	safeTest(t, "Test_S08b_40_Collection_NonEmptyList_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		list := col.NonEmptyList()

		// Assert
		actual := args.Map{"result": len(list) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_NonEmptyListPtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_41_Collection_NonEmptyListPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", ""})

		// Act
		listPtr := col.NonEmptyListPtr()

		// Assert
		actual := args.Map{"result": len(*listPtr) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── HashsetAsIs / HashsetWithDoubleLength / HashsetLock ──────

func Test_Collection_HashsetAsIs_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_42_Collection_HashsetAsIs", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		hs := col.HashsetAsIs()

		// Assert
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 unique", actual)
	})
}

func Test_Collection_HashsetWithDoubleLength_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_43_Collection_HashsetWithDoubleLength", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		hs := col.HashsetWithDoubleLength()

		// Assert
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Collection_HashsetLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_44_Collection_HashsetLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x"})

		// Act
		hs := col.HashsetLock()

		// Assert
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── NonEmptyItems / NonEmptyItemsPtr / NonEmptyItemsOrNonWhitespace / NonEmptyItemsOrNonWhitespacePtr ──

func Test_Collection_NonEmptyItems_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_45_Collection_NonEmptyItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		items := col.NonEmptyItems()

		// Assert
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_NonEmptyItemsPtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_46_Collection_NonEmptyItemsPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", ""})

		// Act
		items := col.NonEmptyItemsPtr()

		// Assert
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespace_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_47_Collection_NonEmptyItemsOrNonWhitespace", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "  ", "", "b"})

		// Act
		items := col.NonEmptyItemsOrNonWhitespace()

		// Assert
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespacePtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_48_Collection_NonEmptyItemsOrNonWhitespacePtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", " "})

		// Act
		items := col.NonEmptyItemsOrNonWhitespacePtr()

		// Assert
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Items / ListPtr / ListCopyPtrLock ────────────────────────

func Test_Collection_Items_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_49_Collection_Items", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": len(col.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_ListPtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_50_Collection_ListPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": len(col.ListPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_ListCopyPtrLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_51_Collection_ListCopyPtrLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		items := col.ListCopyPtrLock()

		// Assert
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_Empty_FromS08b_v8(t *testing.T) {
	safeTest(t, "Test_S08b_52_Collection_ListCopyPtrLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		items := col.ListCopyPtrLock()

		// Assert
		actual := args.Map{"result": len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── Has / HasLock / HasPtr / HasAll / HasUsingSensitivity ────

func Test_Collection_Has_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_53_Collection_Has", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.Has("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_Empty_FromS08b_v9(t *testing.T) {
	safeTest(t, "Test_S08b_54_Collection_Has_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_Collection_HasLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_55_Collection_HasLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": col.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Collection_HasPtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_56_Collection_HasPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x"})
		s := "x"
		missing := "z"

		// Act & Assert
		actual := args.Map{"result": col.HasPtr(&s)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.HasPtr(&missing)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": col.HasPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Collection_Empty_FromS08b_v10(t *testing.T) {
	safeTest(t, "Test_S08b_57_Collection_HasPtr_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		s := "a"

		// Act & Assert
		actual := args.Map{"result": col.HasPtr(&s)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_Collection_HasAll_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_58_Collection_HasAll", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": col.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.HasAll("a", "z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Collection_Empty_FromS08b_v11(t *testing.T) {
	safeTest(t, "Test_S08b_59_Collection_HasAll_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.HasAll("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_Collection_HasUsingSensitivity_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_60_Collection_HasUsingSensitivity", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"Hello", "World"})

		// Act & Assert
		actual := args.Map{"result": col.HasUsingSensitivity("hello", true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for case-sensitive", actual)
		actual = args.Map{"result": col.HasUsingSensitivity("hello", false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for case-insensitive", actual)
		actual = args.Map{"result": col.HasUsingSensitivity("missing", false)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for missing item", actual)
	})
}

// ── IsContainsPtr / IsContainsAll / IsContainsAllSlice / IsContainsAllLock ──

func Test_Collection_IsContainsPtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_61_Collection_IsContainsPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := "a"

		// Act & Assert
		actual := args.Map{"result": col.IsContainsPtr(&s)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.IsContainsPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Collection_IsContainsAllSlice_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_62_Collection_IsContainsAllSlice", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": col.IsContainsAllSlice([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.IsContainsAllSlice([]string{"a", "z"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": col.IsContainsAllSlice([]string{})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_Collection_IsContainsAll_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_63_Collection_IsContainsAll", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.IsContainsAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.IsContainsAll(nil...)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Collection_IsContainsAllLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_64_Collection_IsContainsAllLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": col.IsContainsAllLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": col.IsContainsAllLock(nil...)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Collection_EmptyCollection_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_65_Collection_IsContainsAllSlice_EmptyCollection", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		actual := args.Map{"result": col.IsContainsAllSlice([]string{"a"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty collection", actual)
	})
}

// ── GetHashsetPlusHasAll ─────────────────────────────────────

func Test_Collection_GetHashsetPlusHasAll_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_66_Collection_GetHashsetPlusHasAll", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		hs, hasAll := col.GetHashsetPlusHasAll([]string{"a", "b"})

		// Assert
		actual := args.Map{"result": hasAll}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hs == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashset", actual)
	})
}

func Test_Collection_NilItems_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_67_Collection_GetHashsetPlusHasAll_NilItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		_, hasAll := col.GetHashsetPlusHasAll(nil)

		// Assert
		actual := args.Map{"result": hasAll}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil items", actual)
	})
}

func Test_Collection_EmptyCollection_FromS08b_v2(t *testing.T) {
	safeTest(t, "Test_S08b_68_Collection_GetHashsetPlusHasAll_EmptyCollection", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		_, hasAll := col.GetHashsetPlusHasAll([]string{"a"})

		// Assert
		actual := args.Map{"result": hasAll}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty collection", actual)
	})
}

// ── SortedListAsc / SortedAsc / SortedAscLock / SortedListDsc ──

func Test_Collection_SortedListAsc_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_69_Collection_SortedListAsc", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"c", "a", "b"})

		// Act
		sorted := col.SortedListAsc()

		// Assert
		actual := args.Map{"result": sorted[0] != "a" || sorted[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted ascending", actual)
	})
}

func Test_Collection_Empty_FromS08b_v12(t *testing.T) {
	safeTest(t, "Test_S08b_70_Collection_SortedListAsc_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		sorted := col.SortedListAsc()

		// Assert
		actual := args.Map{"result": len(sorted) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_SortedAsc_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_71_Collection_SortedAsc", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"b", "a"})

		// Act
		col.SortedAsc()

		// Assert
		actual := args.Map{"result": col.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
	})
}

func Test_Collection_Empty_FromS08b_v13(t *testing.T) {
	safeTest(t, "Test_S08b_72_Collection_SortedAsc_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		result := col.SortedAsc()

		// Assert
		actual := args.Map{"result": result != col}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer for empty", actual)
	})
}

func Test_Collection_SortedAscLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_73_Collection_SortedAscLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"b", "a"})

		// Act
		col.SortedAscLock()

		// Assert
		actual := args.Map{"result": col.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
	})
}

func Test_Collection_Empty_FromS08b_v14(t *testing.T) {
	safeTest(t, "Test_S08b_74_Collection_SortedAscLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		result := col.SortedAscLock()

		// Assert
		actual := args.Map{"result": result != col}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer for empty", actual)
	})
}

func Test_Collection_SortedListDsc_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_75_Collection_SortedListDsc", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "c", "b"})

		// Act
		sorted := col.SortedListDsc()

		// Assert
		actual := args.Map{"result": sorted[0] != "c" || sorted[2] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted descending", actual)
	})
}

// ── New ──────────────────────────────────────────────────────

func Test_Collection_New_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_76_Collection_New", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		newCol := col.New("a", "b")

		// Assert
		actual := args.Map{"result": newCol.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b_v15(t *testing.T) {
	safeTest(t, "Test_S08b_77_Collection_New_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		newCol := col.New()

		// Assert
		actual := args.Map{"result": newCol.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddNonEmptyStrings / AddNonEmptyStringsSlice ─────────────

func Test_Collection_AddNonEmptyStrings_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_78_Collection_AddNonEmptyStrings", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddNonEmptyStrings("a", "", "b")

		// Assert — AddNonEmptyStrings filters empty strings, so only "a" and "b" are added
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Empty_FromS08b_v16(t *testing.T) {
	safeTest(t, "Test_S08b_79_Collection_AddNonEmptyStrings_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddNonEmptyStrings()

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddFuncResult ────────────────────────────────────────────

func Test_Collection_AddFuncResult_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_80_Collection_AddFuncResult", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddFuncResult(
			func() string { return "a" },
			func() string { return "b" },
		)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Nil_FromS08b_v6(t *testing.T) {
	safeTest(t, "Test_S08b_81_Collection_AddFuncResult_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddFuncResult(nil...)

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddStringsByFuncChecking ─────────────────────────────────

func Test_Collection_AddStringsByFuncChecking_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_82_Collection_AddStringsByFuncChecking", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddStringsByFuncChecking(
			[]string{"hello", "", "world"},
			func(line string) bool { return line != "" },
		)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── ExpandSlicePlusAdd ───────────────────────────────────────

func Test_Collection_ExpandSlicePlusAdd_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_83_Collection_ExpandSlicePlusAdd", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.ExpandSlicePlusAdd(
			[]string{"a,b", "c,d"},
			func(line string) []string { return strings.Split(line, ",") },
		)

		// Assert
		actual := args.Map{"result": col.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

// ── MergeSlicesOfSlice ───────────────────────────────────────

func Test_Collection_MergeSlicesOfSlice_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_84_Collection_MergeSlicesOfSlice", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.MergeSlicesOfSlice([]string{"a"}, []string{"b", "c"})

		// Assert
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

// ── GetAllExcept / GetAllExceptCollection ────────────────────

func Test_Collection_GetAllExceptCollection_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_85_Collection_GetAllExceptCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})

		// Act
		result := col.GetAllExceptCollection(except)

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_NilExcept_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_86_Collection_GetAllExceptCollection_NilExcept", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		result := col.GetAllExceptCollection(nil)

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_EmptyExcept_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_87_Collection_GetAllExceptCollection_EmptyExcept", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		except := corestr.Empty.Collection()

		// Act
		result := col.GetAllExceptCollection(except)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_GetAllExcept_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_88_Collection_GetAllExcept", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		result := col.GetAllExcept([]string{"a"})

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_NilSlice_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_89_Collection_GetAllExcept_NilSlice", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		result := col.GetAllExcept(nil)

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── CharCollectionMap ────────────────────────────────────────

func Test_Collection_CharCollectionMap_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_90_Collection_CharCollectionMap", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"apple", "banana", "avocado"})

		// Act
		ccm := col.CharCollectionMap()

		// Assert
		actual := args.Map{"result": ccm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── SummaryString / SummaryStringWithHeader / String ─────────

func Test_Collection_SummaryString_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_91_Collection_SummaryString", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.SummaryString(1)

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_SummaryStringWithHeader_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_92_Collection_SummaryStringWithHeader", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		s := col.SummaryStringWithHeader("Header:")

		// Assert
		actual := args.Map{"result": strings.HasPrefix(s, "Header:")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected header prefix", actual)
	})
}

func Test_Collection_Empty_FromS08b_v17(t *testing.T) {
	safeTest(t, "Test_S08b_93_Collection_SummaryStringWithHeader_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.SummaryStringWithHeader("Header:")

		// Assert
		actual := args.Map{"result": strings.Contains(s, "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_Collection_String_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_94_Collection_String", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.String()

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_Empty_FromS08b_v18(t *testing.T) {
	safeTest(t, "Test_S08b_95_Collection_String_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.String()

		// Assert
		actual := args.Map{"result": strings.Contains(s, "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_Collection_StringLock_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_96_Collection_StringLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		s := col.StringLock()

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_Empty_FromS08b_v19(t *testing.T) {
	safeTest(t, "Test_S08b_97_Collection_StringLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.StringLock()

		// Assert
		actual := args.Map{"result": strings.Contains(s, "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

// ── Csv / CsvOptions / CsvLines / CsvLinesOptions ───────────

func Test_Collection_Csv_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_98_Collection_Csv", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		csv := col.Csv()

		// Assert
		actual := args.Map{"result": csv == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_Empty_FromS08b_v20(t *testing.T) {
	safeTest(t, "Test_S08b_99_Collection_Csv_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		csv := col.Csv()

		// Assert
		actual := args.Map{"result": csv != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_CsvOptions_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_100_Collection_CsvOptions", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		csv := col.CsvOptions(true)

		// Assert
		actual := args.Map{"result": csv == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_Empty_FromS08b_v21(t *testing.T) {
	safeTest(t, "Test_S08b_101_Collection_CsvOptions_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		csv := col.CsvOptions(true)

		// Assert
		actual := args.Map{"result": csv != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_CsvLines_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_102_Collection_CsvLines", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		lines := col.CsvLines()

		// Assert
		actual := args.Map{"result": len(lines) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_CsvLinesOptions_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_103_Collection_CsvLinesOptions", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		lines := col.CsvLinesOptions(true)

		// Assert
		actual := args.Map{"result": len(lines) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AddCapacity / Resize ─────────────────────────────────────

func Test_Collection_AddCapacity_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_104_Collection_AddCapacity", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)

		// Act
		col.AddCapacity(10)

		// Assert
		actual := args.Map{"result": col.Capacity() < 15}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 15", actual)
	})
}

func Test_Collection_Nil_FromS08b_v7(t *testing.T) {
	safeTest(t, "Test_S08b_105_Collection_AddCapacity_Nil", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)

		// Act
		col.AddCapacity(nil...)

		// Assert — no panic
	})
}

func Test_Collection_Resize_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_106_Collection_Resize", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)

		// Act
		col.Resize(20)

		// Assert
		actual := args.Map{"result": col.Capacity() < 20}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 20", actual)
	})
}

func Test_Collection_SmallerThanExisting_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_107_Collection_Resize_SmallerThanExisting", func() {
		// Arrange
		col := corestr.New.Collection.Cap(20)

		// Act
		col.Resize(5)

		// Assert — should not shrink
		actual := args.Map{"result": col.Capacity() < 20}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity unchanged", actual)
	})
}

// ── Joins / NonEmptyJoins / NonWhitespaceJoins ───────────────

func Test_Collection_Joins_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_108_Collection_Joins", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.Joins(",")

		// Assert
		actual := args.Map{"result": s != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_Collection_WithExtra_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_109_Collection_Joins_WithExtra", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		s := col.Joins(",", "b", "c")

		// Assert
		actual := args.Map{"result": strings.Contains(s, "b") || !strings.Contains(s, "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected b and c in result, got ''", actual)
	})
}

func Test_Collection_NonEmptyJoins_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_110_Collection_NonEmptyJoins", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		s := col.NonEmptyJoins(",")

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_NonWhitespaceJoins_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_111_Collection_NonWhitespaceJoins", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "  ", "b"})

		// Act
		s := col.NonWhitespaceJoins(",")

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

// ── JSON methods ─────────────────────────────────────────────

func Test_Collection_JsonModel_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_112_Collection_JsonModel", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		model := col.JsonModel()

		// Assert
		actual := args.Map{"result": len(model) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_JsonModelAny_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_113_Collection_JsonModelAny", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		model := col.JsonModelAny()

		// Assert
		actual := args.Map{"result": model == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Collection_MarshalJSON_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_114_Collection_MarshalJSON", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		data, err := col.MarshalJSON()

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual = args.Map{"result": len(data) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
	})
}

func Test_Collection_UnmarshalJSON_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_115_Collection_UnmarshalJSON", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		err := col.UnmarshalJSON([]byte(`["a","b"]`))

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual = args.Map{"result": col.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Invalid_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_116_Collection_UnmarshalJSON_Invalid", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		err := col.UnmarshalJSON([]byte(`invalid`))

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Collection_Json_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_117_Collection_Json", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := col.Json()

		// Assert
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_Collection_JsonPtr_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_118_Collection_JsonPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := col.JsonPtr()

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── ParseInjectUsingJson / ParseInjectUsingJsonMust / JsonParseSelfInject ──

func Test_Collection_ParseInjectUsingJson_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_119_Collection_ParseInjectUsingJson", func() {
		// Arrange
		original := corestr.New.Collection.Strings([]string{"a", "b"})
		jsonResult := original.JsonPtr()
		target := corestr.Empty.Collection()

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual = args.Map{"result": result.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_ParseInjectUsingJsonMust_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_120_Collection_ParseInjectUsingJsonMust", func() {
		// Arrange
		original := corestr.New.Collection.Strings([]string{"a"})
		jsonResult := original.JsonPtr()
		target := corestr.Empty.Collection()

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Collection_JsonParseSelfInject_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_121_Collection_JsonParseSelfInject", func() {
		// Arrange
		original := corestr.New.Collection.Strings([]string{"x"})
		jsonResult := original.JsonPtr()
		target := corestr.Empty.Collection()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	})
}

// ── Clear / Dispose ──────────────────────────────────────────

func Test_Collection_Clear_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_122_Collection_Clear", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		col.Clear()

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Collection_Nil_FromS08b_v8(t *testing.T) {
	safeTest(t, "Test_S08b_123_Collection_Clear_Nil", func() {
		// Arrange
		var col *corestr.Collection

		// Act
		result := col.Clear()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Collection_Dispose_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_124_Collection_Dispose", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		col.Dispose()

		// Assert
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 after dispose", actual)
	})
}

func Test_Collection_Nil_FromS08b_v9(t *testing.T) {
	safeTest(t, "Test_S08b_125_Collection_Dispose_Nil", func() {
		// Arrange
		var col *corestr.Collection

		// Act — should not panic
		col.Dispose()
	})
}

// ── AsJsonMarshaller / AsJsonContractsBinder ─────────────────

func Test_Collection_AsJsonMarshaller_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_126_Collection_AsJsonMarshaller", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		m := col.AsJsonMarshaller()

		// Assert
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Collection_AsJsonContractsBinder_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_127_Collection_AsJsonContractsBinder", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		b := col.AsJsonContractsBinder()

		// Assert
		actual := args.Map{"result": b == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── Serialize / Deserialize ──────────────────────────────────

func Test_Collection_Serialize_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_128_Collection_Serialize", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		data, err := col.Serialize()

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual = args.Map{"result": len(data) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_Deserialize_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_129_Collection_Deserialize", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		var target []string

		// Act
		err := col.Deserialize(&target)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual = args.Map{"result": len(target) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Join / JoinLine ──────────────────────────────────────────

func Test_Collection_Join_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_130_Collection_Join", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.Join(",")

		// Assert
		actual := args.Map{"result": s != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_Collection_Empty_FromS08b_v22(t *testing.T) {
	safeTest(t, "Test_S08b_131_Collection_Join_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.Join(",")

		// Assert
		actual := args.Map{"result": s != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Collection_JoinLine_FromS08b(t *testing.T) {
	safeTest(t, "Test_S08b_132_Collection_JoinLine", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.JoinLine()

		// Assert
		actual := args.Map{"result": strings.Contains(s, "\n")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected newline", actual)
	})
}

func Test_Collection_Empty_FromS08b_v23(t *testing.T) {
	safeTest(t, "Test_S08b_133_Collection_JoinLine_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.JoinLine()

		// Assert
		actual := args.Map{"result": s != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}
