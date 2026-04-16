package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S10b — Hashset.go Lines 658-1469 — Query, Filter, JSON, etc
// ══════════════════════════════════════════════════════════════

// ── SortedList ───────────────────────────────────────────────

func Test_88_Hashset_SortedList(t *testing.T) {
	safeTest(t, "Test_88_Hashset_SortedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		// Act
		sorted := hs.SortedList()

		// Assert
		actual := args.Map{"result": len(sorted) != 3 || sorted[0] != "a" || sorted[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted asc", actual)
	})
}

// ── Filter ───────────────────────────────────────────────────

func Test_89_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_89_Hashset_Filter", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana", "avocado"})

		// Act
		result := hs.Filter(func(s string) bool {
			return strings.HasPrefix(s, "a")
		})

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── OrderedList / SafeStrings / Lines ────────────────────────

func Test_90_Hashset_OrderedList(t *testing.T) {
	safeTest(t, "Test_90_Hashset_OrderedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		// Act
		list := hs.OrderedList()

		// Assert
		actual := args.Map{"result": len(list) != 3 || list[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted", actual)
	})
}

func Test_91_Hashset_OrderedList_Empty(t *testing.T) {
	safeTest(t, "Test_91_Hashset_OrderedList_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		list := hs.OrderedList()

		// Assert
		actual := args.Map{"result": len(list) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_92_Hashset_SafeStrings(t *testing.T) {
	safeTest(t, "Test_92_Hashset_SafeStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.SafeStrings()

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_93_Hashset_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_93_Hashset_SafeStrings_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		result := hs.SafeStrings()

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_94_Hashset_Lines(t *testing.T) {
	safeTest(t, "Test_94_Hashset_Lines", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": len(hs.Lines()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_95_Hashset_Lines_Empty(t *testing.T) {
	safeTest(t, "Test_95_Hashset_Lines_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act & Assert
		actual := args.Map{"result": len(hs.Lines()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── SimpleSlice ──────────────────────────────────────────────

func Test_96_Hashset_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_96_Hashset_SimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		ss := hs.SimpleSlice()

		// Assert
		actual := args.Map{"result": ss.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_97_Hashset_SimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_97_Hashset_SimpleSlice_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		ss := hs.SimpleSlice()

		// Assert
		actual := args.Map{"result": ss.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ── GetFilteredItems ─────────────────────────────────────────

func Test_98_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_98_Hashset_GetFilteredItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, strings.HasPrefix(str, "a"), false
		}

		// Act
		result := hs.GetFilteredItems(filter)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_99_Hashset_GetFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_99_Hashset_GetFilteredItems_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := hs.GetFilteredItems(filter)

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_100_Hashset_GetFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_100_Hashset_GetFilteredItems_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}

		// Act
		result := hs.GetFilteredItems(filter)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_101_Hashset_GetFilteredItems_Skip(t *testing.T) {
	safeTest(t, "Test_101_Hashset_GetFilteredItems_Skip", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, false, false
		}

		// Act
		result := hs.GetFilteredItems(filter)

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── GetFilteredCollection ────────────────────────────────────

func Test_102_Hashset_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_102_Hashset_GetFilteredCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := hs.GetFilteredCollection(filter)

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_103_Hashset_GetFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_103_Hashset_GetFilteredCollection_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		result := hs.GetFilteredCollection(nil)

		// Assert
		actual := args.Map{"result": result.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_104_Hashset_GetFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_104_Hashset_GetFilteredCollection_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}

		// Act
		result := hs.GetFilteredCollection(filter)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_105_Hashset_GetFilteredCollection_Skip(t *testing.T) {
	safeTest(t, "Test_105_Hashset_GetFilteredCollection_Skip", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, false, false
		}

		// Act
		result := hs.GetFilteredCollection(filter)

		// Assert
		actual := args.Map{"result": result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── GetAllExcept variants ────────────────────────────────────

func Test_106_Hashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_106_Hashset_GetAllExceptHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		except := corestr.New.Hashset.Strings([]string{"b"})

		// Act
		result := hs.GetAllExceptHashset(except)

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_107_Hashset_GetAllExceptHashset_Nil(t *testing.T) {
	safeTest(t, "Test_107_Hashset_GetAllExceptHashset_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.GetAllExceptHashset(nil)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all items", actual)
	})
}

func Test_108_Hashset_GetAllExceptHashset_Empty(t *testing.T) {
	safeTest(t, "Test_108_Hashset_GetAllExceptHashset_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.GetAllExceptHashset(corestr.Empty.Hashset())

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all items", actual)
	})
}

func Test_109_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_109_Hashset_GetAllExcept", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		result := hs.GetAllExcept([]string{"a"})

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_110_Hashset_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_110_Hashset_GetAllExcept_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.GetAllExcept(nil)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all items", actual)
	})
}

func Test_111_Hashset_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_111_Hashset_GetAllExceptSpread", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		result := hs.GetAllExceptSpread("a")

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_112_Hashset_GetAllExceptSpread_Nil(t *testing.T) {
	safeTest(t, "Test_112_Hashset_GetAllExceptSpread_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.GetAllExceptSpread(nil...)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all items", actual)
	})
}

func Test_113_Hashset_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_113_Hashset_GetAllExceptCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := hs.GetAllExceptCollection(col)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_114_Hashset_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_114_Hashset_GetAllExceptCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.GetAllExceptCollection(nil)

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected all items", actual)
	})
}

// ── Items / List / MapStringAny / MapStringAnyDiff ───────────

func Test_115_Hashset_Items(t *testing.T) {
	safeTest(t, "Test_115_Hashset_Items", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": len(hs.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_116_Hashset_List(t *testing.T) {
	safeTest(t, "Test_116_Hashset_List", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		list := hs.List()

		// Assert
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// call again for cache path
		list2 := hs.List()
		actual = args.Map{"result": len(list2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 cached", actual)
	})
}

func Test_117_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_117_Hashset_MapStringAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		m := hs.MapStringAny()

		// Assert
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_118_Hashset_MapStringAny_Empty(t *testing.T) {
	safeTest(t, "Test_118_Hashset_MapStringAny_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		m := hs.MapStringAny()

		// Assert
		actual := args.Map{"result": len(m) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_119_Hashset_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_119_Hashset_MapStringAnyDiff", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		d := hs.MapStringAnyDiff()

		// Assert
		actual := args.Map{"result": d == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── JoinSorted / ListPtrSortedAsc / ListPtrSortedDsc / ListPtr ──

func Test_120_Hashset_JoinSorted(t *testing.T) {
	safeTest(t, "Test_120_Hashset_JoinSorted", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})

		// Act
		s := hs.JoinSorted(",")

		// Assert
		actual := args.Map{"result": s != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_121_Hashset_JoinSorted_Empty(t *testing.T) {
	safeTest(t, "Test_121_Hashset_JoinSorted_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		s := hs.JoinSorted(",")

		// Assert
		actual := args.Map{"result": s != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_122_Hashset_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_122_Hashset_ListPtrSortedAsc", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		// Act
		list := hs.ListPtrSortedAsc()

		// Assert
		actual := args.Map{"result": list[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
	})
}

func Test_123_Hashset_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_123_Hashset_ListPtrSortedDsc", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		// Act
		list := hs.ListPtrSortedDsc()

		// Assert
		actual := args.Map{"result": list[0] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_124_Hashset_ListPtr(t *testing.T) {
	safeTest(t, "Test_124_Hashset_ListPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		list := hs.ListPtr()

		// Assert
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Clear / Dispose ──────────────────────────────────────────

func Test_125_Hashset_Clear(t *testing.T) {
	safeTest(t, "Test_125_Hashset_Clear", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hs.Clear()

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_126_Hashset_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_126_Hashset_Clear_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act
		result := hs.Clear()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_127_Hashset_Dispose(t *testing.T) {
	safeTest(t, "Test_127_Hashset_Dispose", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hs.Dispose()

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_128_Hashset_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_128_Hashset_Dispose_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act — should not panic
		hs.Dispose()
	})
}

// ── ListCopyLock ─────────────────────────────────────────────

func Test_129_Hashset_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_129_Hashset_ListCopyLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		list := hs.ListCopyLock()

		// Assert
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── ToLowerSet ───────────────────────────────────────────────

func Test_130_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_130_Hashset_ToLowerSet", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"ABC", "Def"})

		// Act
		lowered := hs.ToLowerSet()

		// Assert
		actual := args.Map{"result": lowered.Has("abc") || !lowered.Has("def")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected lowered keys", actual)
	})
}

// ── Length / LengthLock ──────────────────────────────────────

func Test_131_Hashset_Length(t *testing.T) {
	safeTest(t, "Test_131_Hashset_Length", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_132_Hashset_Length_Nil(t *testing.T) {
	safeTest(t, "Test_132_Hashset_Length_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act & Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_133_Hashset_LengthLock(t *testing.T) {
	safeTest(t, "Test_133_Hashset_LengthLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Remove / SafeRemove / RemoveWithLock ─────────────────────

func Test_134_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_134_Hashset_Remove", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		hs.Remove("a")

		// Assert
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_135_Hashset_SafeRemove(t *testing.T) {
	safeTest(t, "Test_135_Hashset_SafeRemove", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hs.SafeRemove("a")
		hs.SafeRemove("missing")

		// Assert
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_136_Hashset_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_136_Hashset_RemoveWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hs.RemoveWithLock("a")

		// Assert
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

// ── String / StringLock ──────────────────────────────────────

func Test_137_Hashset_String(t *testing.T) {
	safeTest(t, "Test_137_Hashset_String", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		s := hs.String()

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_138_Hashset_String_Empty(t *testing.T) {
	safeTest(t, "Test_138_Hashset_String_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		s := hs.String()

		// Assert
		actual := args.Map{"result": strings.Contains(s, "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_139_Hashset_StringLock(t *testing.T) {
	safeTest(t, "Test_139_Hashset_StringLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		s := hs.StringLock()

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_140_Hashset_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_140_Hashset_StringLock_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		s := hs.StringLock()

		// Assert
		actual := args.Map{"result": strings.Contains(s, "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

// ── Join / NonEmptyJoins / NonWhitespaceJoins / JoinLine ─────

func Test_141_Hashset_Join(t *testing.T) {
	safeTest(t, "Test_141_Hashset_Join", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		s := hs.Join(",")

		// Assert
		actual := args.Map{"result": s != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
	})
}

func Test_142_Hashset_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_142_Hashset_NonEmptyJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		s := hs.NonEmptyJoins(",")

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_143_Hashset_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_143_Hashset_NonWhitespaceJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		s := hs.NonWhitespaceJoins(",")

		// Assert
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_144_Hashset_JoinLine(t *testing.T) {
	safeTest(t, "Test_144_Hashset_JoinLine", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		s := hs.JoinLine()

		// Assert
		actual := args.Map{"result": s != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
	})
}

// ── JSON methods ─────────────────────────────────────────────

func Test_145_Hashset_JsonModel(t *testing.T) {
	safeTest(t, "Test_145_Hashset_JsonModel", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": len(hs.JsonModel()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_146_Hashset_JsonModel_Empty(t *testing.T) {
	safeTest(t, "Test_146_Hashset_JsonModel_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act & Assert
		actual := args.Map{"result": len(hs.JsonModel()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_147_Hashset_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_147_Hashset_JsonModelAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_148_Hashset_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_148_Hashset_MarshalJSON", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		data, err := hs.MarshalJSON()

		// Assert
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON", actual)
	})
}

func Test_149_Hashset_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_149_Hashset_UnmarshalJSON", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		err := hs.UnmarshalJSON([]byte(`{"a":true,"b":true}`))

		// Assert
		actual := args.Map{"result": err != nil || hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_150_Hashset_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_150_Hashset_UnmarshalJSON_Invalid", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		err := hs.UnmarshalJSON([]byte(`invalid`))

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_151_Hashset_Json(t *testing.T) {
	safeTest(t, "Test_151_Hashset_Json", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.Json()

		// Assert
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_152_Hashset_JsonPtr(t *testing.T) {
	safeTest(t, "Test_152_Hashset_JsonPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_153_Hashset_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_153_Hashset_ParseInjectUsingJson", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jsonResult := hs.JsonPtr()
		target := corestr.Empty.Hashset()

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		actual := args.Map{"result": err != nil || result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_154_Hashset_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_154_Hashset_ParseInjectUsingJsonMust", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jsonResult := hs.JsonPtr()
		target := corestr.Empty.Hashset()

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_155_Hashset_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_155_Hashset_JsonParseSelfInject", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jsonResult := hs.JsonPtr()
		target := corestr.Empty.Hashset()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_156_Hashset_AsJsoner(t *testing.T) {
	safeTest(t, "Test_156_Hashset_AsJsoner", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": hs.AsJsoner() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_157_Hashset_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_157_Hashset_AsJsonContractsBinder", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": hs.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_158_Hashset_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_158_Hashset_AsJsonParseSelfInjector", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": hs.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_159_Hashset_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_159_Hashset_AsJsonMarshaller", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": hs.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── DistinctDiffLinesRaw ─────────────────────────────────────

func Test_160_Hashset_DistinctDiffLinesRaw(t *testing.T) {
	safeTest(t, "Test_160_Hashset_DistinctDiffLinesRaw", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		diff := hs.DistinctDiffLinesRaw("b", "c")

		// Assert
		actual := args.Map{"result": len(diff) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_161_Hashset_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_161_Hashset_DistinctDiffLinesRaw_BothEmpty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		diff := hs.DistinctDiffLinesRaw()

		// Assert
		actual := args.Map{"result": len(diff) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_162_Hashset_DistinctDiffLinesRaw_LeftOnly(t *testing.T) {
	safeTest(t, "Test_162_Hashset_DistinctDiffLinesRaw_LeftOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		diff := hs.DistinctDiffLinesRaw()

		// Assert
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_163_Hashset_DistinctDiffLinesRaw_RightOnly(t *testing.T) {
	safeTest(t, "Test_163_Hashset_DistinctDiffLinesRaw_RightOnly", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		diff := hs.DistinctDiffLinesRaw("a")

		// Assert
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── DistinctDiffHashset / DistinctDiffLines ──────────────────

func Test_164_Hashset_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_164_Hashset_DistinctDiffHashset", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a", "b"})
		b := corestr.New.Hashset.Strings([]string{"b", "c"})

		// Act
		diff := a.DistinctDiffHashset(b)

		// Assert
		actual := args.Map{"result": len(diff) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_165_Hashset_DistinctDiffLines(t *testing.T) {
	safeTest(t, "Test_165_Hashset_DistinctDiffLines", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		diff := hs.DistinctDiffLines("b", "c")

		// Assert
		actual := args.Map{"result": len(diff) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_166_Hashset_DistinctDiffLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_166_Hashset_DistinctDiffLines_BothEmpty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		diff := hs.DistinctDiffLines()

		// Assert
		actual := args.Map{"result": len(diff) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_167_Hashset_DistinctDiffLines_LeftOnly(t *testing.T) {
	safeTest(t, "Test_167_Hashset_DistinctDiffLines_LeftOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		diff := hs.DistinctDiffLines()

		// Assert
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_168_Hashset_DistinctDiffLines_RightOnly(t *testing.T) {
	safeTest(t, "Test_168_Hashset_DistinctDiffLines_RightOnly", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		diff := hs.DistinctDiffLines("x")

		// Assert
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Serialize / Deserialize ──────────────────────────────────

func Test_169_Hashset_Serialize(t *testing.T) {
	safeTest(t, "Test_169_Hashset_Serialize", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		data, err := hs.Serialize()

		// Assert
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid bytes", actual)
	})
}

func Test_170_Hashset_Deserialize(t *testing.T) {
	safeTest(t, "Test_170_Hashset_Deserialize", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		var target map[string]bool

		// Act
		err := hs.Deserialize(&target)

		// Assert
		actual := args.Map{"result": err != nil || len(target) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── Wrap methods / Transpile ─────────────────────────────────

func Test_171_Hashset_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_171_Hashset_WrapDoubleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.WrapDoubleQuote()

		// Assert
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_172_Hashset_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_172_Hashset_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.WrapDoubleQuoteIfMissing()

		// Assert
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_173_Hashset_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_173_Hashset_WrapSingleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.WrapSingleQuote()

		// Assert
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_174_Hashset_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_174_Hashset_WrapSingleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.WrapSingleQuoteIfMissing()

		// Assert
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_175_Hashset_Transpile(t *testing.T) {
	safeTest(t, "Test_175_Hashset_Transpile", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.Transpile(strings.ToUpper)

		// Assert
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_176_Hashset_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_176_Hashset_Transpile_Empty", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		result := hs.Transpile(strings.ToUpper)

		// Assert
		actual := args.Map{"result": result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
