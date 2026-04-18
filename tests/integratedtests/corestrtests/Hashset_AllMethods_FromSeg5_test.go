package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Segment 5b
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashset_IsEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

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

func Test_Hashset_Add_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Add", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a").Add("b")

		// Act
		actual := args.Map{
			"len": h.Length(),
			"has": h.Has("a"),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 items", actual)
	})
}

func Test_Hashset_AddBool_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddBool", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		existed1 := h.AddBool("a")
		existed2 := h.AddBool("a")

		// Act
		actual := args.Map{
			"existed1": existed1,
			"existed2": existed2,
		}

		// Assert
		expected := args.Map{
			"existed1": false,
			"existed2": true,
		}
		expected.ShouldBeEqual(t, 0, "AddBool -- new then existing", actual)
	})
}

func Test_Hashset_AddNonEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddNonEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddNonEmpty("a").AddNonEmpty("")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty -- skips empty", actual)
	})
}

func Test_Hashset_AddNonEmptyWhitespace_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddNonEmptyWhitespace", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("  ")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace -- skips whitespace", actual)
	})
}

func Test_Hashset_AddIf_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddIf", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddIf(true, "a").AddIf(false, "b")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf -- only true", actual)
	})
}

func Test_Hashset_AddIfMany_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddIfMany", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h.AddIfMany(true, "a", "b").AddIfMany(false, "c")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddIfMany -- only true batch", actual)
	})
}

func Test_Hashset_AddFunc_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddFunc", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddFunc(func() string { return "a" })

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddFunc -- added", actual)
	})
}

func Test_Hashset_AddFuncErr_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddFuncErr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddFuncErr(func() (string, error) { return "a", nil }, func(err error) {})

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddFuncErr -- added", actual)
	})
}

func Test_Hashset_AddFuncErr_Error_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddFuncErr_Error", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		var handledErr error
		h.AddFuncErr(
			func() (string, error) { return "", &testErr{} },
			func(err error) { handledErr = err },
		)

		// Act
		actual := args.Map{
			"handled": handledErr != nil,
			"len": h.Length(),
		}

		// Assert
		expected := args.Map{
			"handled": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "AddFuncErr error -- handled", actual)
	})
}

// testErr is defined in shared_compat_helpers.go

func Test_Hashset_Adds_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Adds", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h.Adds("a", "b", "c")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Adds -- 3 items", actual)
	})
}

func Test_Hashset_Adds_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Adds_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Adds(nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_Hashset_AddStrings_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 2 items", actual)
	})
}

func Test_Hashset_AddStrings_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStrings_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddStrings(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings nil -- no change", actual)
	})
}

func Test_Hashset_AddPtr_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddPtr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		s := "test"
		h.AddPtr(&s)

		// Act
		actual := args.Map{"has": h.Has("test")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddPtr -- added", actual)
	})
}

func Test_Hashset_AddLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddLock("a")

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddLock -- added", actual)
	})
}

func Test_Hashset_AddPtrLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddPtrLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		s := "test"
		h.AddPtrLock(&s)

		// Act
		actual := args.Map{"has": h.Has("test")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddPtrLock -- added", actual)
	})
}

func Test_Hashset_AddWithWgLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddWithWgLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddWithWgLock("a", &wg)
		wg.Wait()

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddWithWgLock -- added", actual)
	})
}

func Test_Hashset_AddHashsetItems_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetItems", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h2 := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.AddHashsetItems(h2)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems -- 2 items", actual)
	})
}

func Test_Hashset_AddHashsetItems_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetItems_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddHashsetItems(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems nil -- no change", actual)
	})
}

func Test_Hashset_AddItemsMap_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMap", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddItemsMap -- only true values", actual)
	})
}

func Test_Hashset_AddItemsMap_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMap_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddItemsMap(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMap nil -- no change", actual)
	})
}

func Test_Hashset_AddItemsMapWgLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMapWgLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		m := map[string]bool{"a": true, "b": false}
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddItemsMapWgLock(&m, &wg)
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddItemsMapWgLock -- only true", actual)
	})
}

func Test_Hashset_AddItemsMapWgLock_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMapWgLock_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddItemsMapWgLock(nil, nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMapWgLock nil -- no change", actual)
	})
}

func Test_Hashset_AddHashsetWgLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetWgLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddHashsetWgLock(h2, &wg)
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetWgLock -- 1 item", actual)
	})
}

func Test_Hashset_AddHashsetWgLock_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetWgLock_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddHashsetWgLock(nil, nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetWgLock nil -- no change", actual)
	})
}

func Test_Hashset_AddStringsPtrWgLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStringsPtrWgLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddStringsPtrWgLock([]string{"a", "b"}, &wg)
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsPtrWgLock -- 2 items", actual)
	})
}

func Test_Hashset_AddStringsLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStringsLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h.AddStringsLock([]string{"a", "b"})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsLock -- 2 items", actual)
	})
}

func Test_Hashset_AddStringsLock_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStringsLock_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddStringsLock(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock nil -- no change", actual)
	})
}

func Test_Hashset_AddCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollection", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h.AddCollection(c)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollection -- 2 items", actual)
	})
}

func Test_Hashset_AddCollection_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollection_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddCollection(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollection nil -- no change", actual)
	})
}

func Test_Hashset_AddCollections_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollections", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		h.AddCollections(c1, c2)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollections -- 2 items", actual)
	})
}

func Test_Hashset_AddCollections_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollections_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddCollections(nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollections nil -- no change", actual)
	})
}

func Test_Hashset_AddSimpleSlice_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddSimpleSlice", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		ss := corestr.SimpleSlice{"a", "b"}
		h.AddSimpleSlice(&ss)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSimpleSlice -- 2 items", actual)
	})
}

// ── Has / Contains / Missing ────────────────────────────────────────────────

func Test_Hashset_Has_Contains_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Has_Contains", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has":      h.Has("a"),
			"contains": h.Contains("a"),
			"missing":  h.IsMissing("z"),
			"hasLock":  h.HasLock("a"),
		}

		// Assert
		expected := args.Map{
			"has":      true,
			"contains": true,
			"missing":  true,
			"hasLock":  true,
		}
		expected.ShouldBeEqual(t, 0, "Has/Contains/Missing -- correct", actual)
	})
}

func Test_Hashset_IsMissingLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsMissingLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"missing": h.IsMissingLock("z")}

		// Assert
		expected := args.Map{"missing": true}
		expected.ShouldBeEqual(t, 0, "IsMissingLock -- true", actual)
	})
}

func Test_Hashset_HasWithLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasWithLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"has": h.HasWithLock("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasWithLock -- true", actual)
	})
}

func Test_Hashset_HasAllStrings_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"all": h.HasAllStrings([]string{"a", "b"}),
			"miss": h.HasAllStrings([]string{"a", "z"}),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllStrings -- all and missing", actual)
	})
}

func Test_Hashset_HasAll_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAll", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"all": h.HasAll("a", "b"),
			"miss": h.HasAll("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll -- all and missing", actual)
	})
}

func Test_Hashset_HasAllCollectionItems_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
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

func Test_Hashset_HasAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAny", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

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

func Test_Hashset_HasAnyItem_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAnyItem", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"has": h.HasAnyItem()}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem -- true", actual)
	})
}

func Test_Hashset_IsAllMissing_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsAllMissing", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"miss": h.IsAllMissing("x", "y"),
			"notMiss": h.IsAllMissing("a", "y"),
		}

		// Assert
		expected := args.Map{
			"miss": true,
			"notMiss": false,
		}
		expected.ShouldBeEqual(t, 0, "IsAllMissing -- all missing and partial", actual)
	})
}

// ── List / Items / Lines / Sorted ───────────────────────────────────────────

func Test_Hashset_List_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_List", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.List())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "List -- 1 item", actual)
	})
}

func Test_Hashset_Items_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Items", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1 item", actual)
	})
}

func Test_Hashset_SafeStrings_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SafeStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.SafeStrings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SafeStrings -- 1 item", actual)
	})
}

func Test_Hashset_SafeStrings_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SafeStrings_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"len": len(h.SafeStrings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeStrings empty -- 0", actual)
	})
}

func Test_Hashset_Lines_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Lines", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.Lines())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Lines -- 1 item", actual)
	})
}

func Test_Hashset_Lines_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Lines_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"len": len(h.Lines())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Lines empty -- 0", actual)
	})
}

func Test_Hashset_OrderedList_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_OrderedList", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		result := h.OrderedList()

		// Act
		actual := args.Map{"first": result[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "OrderedList -- sorted asc", actual)
	})
}

func Test_Hashset_OrderedList_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_OrderedList_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"len": len(h.OrderedList())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "OrderedList empty -- 0", actual)
	})
}

func Test_Hashset_SortedList_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SortedList", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		result := h.SortedList()

		// Act
		actual := args.Map{"first": result[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedList -- sorted", actual)
	})
}

func Test_Hashset_ListPtrSortedAsc_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ListPtrSortedAsc", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"c", "a"})
		result := h.ListPtrSortedAsc()

		// Act
		actual := args.Map{"first": result[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "ListPtrSortedAsc -- sorted", actual)
	})
}

func Test_Hashset_ListPtrSortedDsc_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ListPtrSortedDsc", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "c"})
		result := h.ListPtrSortedDsc()

		// Act
		actual := args.Map{"first": result[0]}

		// Assert
		expected := args.Map{"first": "c"}
		expected.ShouldBeEqual(t, 0, "ListPtrSortedDsc -- descending", actual)
	})
}

func Test_Hashset_ListCopyLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ListCopyLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.ListCopyLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListCopyLock -- 1", actual)
	})
}

func Test_Hashset_SimpleSlice_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SimpleSlice", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": h.SimpleSlice().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice -- 1 item", actual)
	})
}

func Test_Hashset_SimpleSlice_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SimpleSlice_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"empty": h.SimpleSlice().IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice empty -- empty", actual)
	})
}

func Test_Hashset_Collection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Collection", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": h.Collection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection -- 1 item", actual)
	})
}

func Test_Hashset_MapStringAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MapStringAny", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.MapStringAny())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "MapStringAny -- 1 item", actual)
	})
}

func Test_Hashset_MapStringAny_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MapStringAny_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"len": len(h.MapStringAny())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "MapStringAny empty -- 0", actual)
	})
}

func Test_Hashset_MapStringAnyDiff_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MapStringAnyDiff", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.MapStringAnyDiff())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "MapStringAnyDiff -- 1 item", actual)
	})
}

// ── Filter ──────────────────────────────────────────────────────────────────

func Test_Hashset_Filter_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Filter", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"aa", "b", "cc"})
		result := h.Filter(func(s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter -- 2 match", actual)
	})
}

func Test_Hashset_GetFilteredItems_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"aa", "b"})
		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems -- 1 match", actual)
	})
}

func Test_Hashset_GetFilteredItems_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredItems_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems empty -- 0", actual)
	})
}

func Test_Hashset_GetFilteredItems_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredItems_Break", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"hasItems": len(result) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems break -- stops early", actual)
	})
}

func Test_Hashset_GetFilteredCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredCollection", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"aa", "b"})
		result := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection -- 1 match", actual)
	})
}

func Test_Hashset_GetFilteredCollection_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredCollection_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		result := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection empty -- empty", actual)
	})
}

func Test_Hashset_GetFilteredCollection_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredCollection_Break", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"hasItems": result.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection break -- stops", actual)
	})
}

func Test_Hashset_AddsUsingFilter_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddsUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h.AddsUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false },
			"a", "bb", "c",
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter -- 1 kept", actual)
	})
}

func Test_Hashset_AddsUsingFilter_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddsUsingFilter_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.AddsUsingFilter(nil, nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter nil -- no change", actual)
	})
}

func Test_Hashset_AddsUsingFilter_Break_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddsUsingFilter_Break", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(4)
		h.AddsUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, true },
			"a", "b",
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter break -- 1 item", actual)
	})
}

// ── Except ──────────────────────────────────────────────────────────────────

func Test_Hashset_GetAllExceptHashset_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptHashset", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		except := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptHashset(except)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset -- 1 remaining", actual)
	})
}

func Test_Hashset_GetAllExceptHashset_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptHashset_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptHashset(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset nil -- all items", actual)
	})
}

func Test_Hashset_GetAllExcept_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExcept", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := h.GetAllExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExcept -- 1 remaining", actual)
	})
}

func Test_Hashset_GetAllExcept_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExcept_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExcept(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExcept nil -- all items", actual)
	})
}

func Test_Hashset_GetAllExceptSpread_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptSpread", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := h.GetAllExceptSpread("a")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptSpread -- 1 remaining", actual)
	})
}

func Test_Hashset_GetAllExceptSpread_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptSpread_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptSpread(nil...)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptSpread nil -- all items", actual)
	})
}

func Test_Hashset_GetAllExceptCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptCollection", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a"})
		result := h.GetAllExceptCollection(c)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection -- 1 remaining", actual)
	})
}

func Test_Hashset_GetAllExceptCollection_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptCollection_Nil", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil -- all items", actual)
	})
}

// ── Resize / AddCapacities ──────────────────────────────────────────────────

func Test_Hashset_Resize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Resize", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Resize(10)

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Resize -- preserved items", actual)
	})
}

func Test_Hashset_Resize_SmallerThanLen_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Resize_SmallerThanLen", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Resize(1)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Resize smaller -- no change", actual)
	})
}

func Test_Hashset_ResizeLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ResizeLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.ResizeLock(10)

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ResizeLock -- preserved items", actual)
	})
}

func Test_Hashset_AddCapacities_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacities", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacities(10, 20)

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCapacities -- preserved items", actual)
	})
}

func Test_Hashset_AddCapacities_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacities_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacities()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCapacities empty -- no change", actual)
	})
}

func Test_Hashset_AddCapacitiesLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacitiesLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacitiesLock(10)

		// Act
		actual := args.Map{"has": h.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCapacitiesLock -- preserved", actual)
	})
}

func Test_Hashset_AddCapacitiesLock_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacitiesLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacitiesLock()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCapacitiesLock empty -- no change", actual)
	})
}

// ── Concat ──────────────────────────────────────────────────────────────────

func Test_Hashset_ConcatNewHashsets_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewHashsets", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		result := h.ConcatNewHashsets(true, h2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets -- merged", actual)
	})
}

func Test_Hashset_ConcatNewHashsets_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewHashsets_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewHashsets(true)

		// Act
		actual := args.Map{"has": result.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets empty -- cloned", actual)
	})
}

func Test_Hashset_ConcatNewStrings_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true, []string{"b", "c"})

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings -- merged", actual)
	})
}

func Test_Hashset_ConcatNewStrings_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewStrings_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true)

		// Act
		actual := args.Map{"has": result.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings empty -- cloned", actual)
	})
}

// ── IsEquals / IsEqual ──────────────────────────────────────────────────────

func Test_Hashset_IsEquals_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		h3 := corestr.New.Hashset.Strings([]string{"b"})

		// Act
		actual := args.Map{
			"eq":      h1.IsEquals(h2),
			"neq":     h1.IsEquals(h3),
			"same":    h1.IsEquals(h1),
			"nilBoth": (*corestr.Hashset)(nil).IsEquals(nil),
			"nilOne":  h1.IsEquals(nil),
		}

		// Assert
		expected := args.Map{
			"eq":      true,
			"neq":     false,
			"same":    true,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEquals -- various", actual)
	})
}

func Test_Hashset_IsEquals_DiffLen_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals_DiffLen", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"eq": h1.IsEquals(h2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len -- false", actual)
	})
}

func Test_Hashset_IsEquals_BothEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals_BothEmpty", func() {
		// Arrange
		h1 := corestr.New.Hashset.Empty()
		h2 := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"eq": h1.IsEquals(h2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty -- true", actual)
	})
}

func Test_Hashset_IsEquals_OneEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals_OneEmpty", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"eq": h1.IsEquals(h2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty -- false", actual)
	})
}

func Test_Hashset_IsEqual_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"eq": h1.IsEqual(h2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual -- delegates to IsEquals", actual)
	})
}

func Test_Hashset_IsEqualsLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEqualsLock", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"eq": h1.IsEqualsLock(h2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock -- true", actual)
	})
}

// ── Remove ──────────────────────────────────────────────────────────────────

func Test_Hashset_Remove_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Remove", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Remove("a")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Remove -- removed", actual)
	})
}

func Test_Hashset_SafeRemove_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SafeRemove", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.SafeRemove("a").SafeRemove("z")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeRemove -- safe", actual)
	})
}

func Test_Hashset_RemoveWithLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_RemoveWithLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.RemoveWithLock("a")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveWithLock -- removed", actual)
	})
}

// ── String / Join ───────────────────────────────────────────────────────────

func Test_Hashset_String_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_String", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"nonEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Hashset_String_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_String_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"nonEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- has NoElements text", actual)
	})
}

func Test_Hashset_StringLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_StringLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"nonEmpty": h.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Hashset_StringLock_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_StringLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"nonEmpty": h.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty -- NoElements", actual)
	})
}

func Test_Hashset_Join_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Join", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"val": h.Join(",")}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Join -- value", actual)
	})
}

func Test_Hashset_JoinSorted_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JoinSorted", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"b", "a"})

		// Act
		actual := args.Map{"val": h.JoinSorted(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "JoinSorted -- sorted", actual)
	})
}

func Test_Hashset_JoinSorted_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JoinSorted_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"val": h.JoinSorted(",")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinSorted empty -- empty", actual)
	})
}

func Test_Hashset_JoinLine_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JoinLine", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"nonEmpty": h.JoinLine() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine -- non-empty", actual)
	})
}

func Test_Hashset_NonEmptyJoins_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_NonEmptyJoins", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"nonEmpty": h.NonEmptyJoins(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonEmptyJoins -- non-empty", actual)
	})
}

func Test_Hashset_NonWhitespaceJoins_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_NonWhitespaceJoins", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"nonEmpty": h.NonWhitespaceJoins(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonWhitespaceJoins -- non-empty", actual)
	})
}

// ── ToLowerSet ──────────────────────────────────────────────────────────────

func Test_Hashset_ToLowerSet_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ToLowerSet", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"ABC"})
		result := h.ToLowerSet()

		// Act
		actual := args.Map{"has": result.Has("abc")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ToLowerSet -- lowered", actual)
	})
}

// ── Length / IsEmpty Lock ───────────────────────────────────────────────────

func Test_Hashset_Length_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Length_Nil", func() {
		// Arrange
		var h *corestr.Hashset

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_Hashset_LengthLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_LengthLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": h.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Hashset_IsEmptyLock_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEmptyLock", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"empty": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_Hashset_Clear_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Clear", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Clear()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Hashset_Clear_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Clear_Nil", func() {
		// Arrange
		var h *corestr.Hashset

		// Act
		actual := args.Map{"nil": h.Clear() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

func Test_Hashset_Dispose_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Dispose", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Dispose()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_Hashset_Dispose_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Dispose_Nil", func() {
		var h *corestr.Hashset
		h.Dispose() // should not panic
	})
}

// ── Wrap / Transpile ────────────────────────────────────────────────────────

func Test_Hashset_WrapDoubleQuote_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapDoubleQuote", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapDoubleQuote()

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuote -- 1 item", actual)
	})
}

func Test_Hashset_WrapSingleQuote_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapSingleQuote", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapSingleQuote()

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuote -- 1 item", actual)
	})
}

func Test_Hashset_WrapDoubleQuoteIfMissing_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapDoubleQuoteIfMissing()

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuoteIfMissing -- 1 item", actual)
	})
}

func Test_Hashset_WrapSingleQuoteIfMissing_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapSingleQuoteIfMissing", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapSingleQuoteIfMissing()

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuoteIfMissing -- 1 item", actual)
	})
}

func Test_Hashset_Transpile_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Transpile", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Transpile -- 1 item", actual)
	})
}

func Test_Hashset_Transpile_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Transpile_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		result := h.Transpile(func(s string) string { return s })

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Transpile empty -- empty", actual)
	})
}

// ── DistinctDiff ────────────────────────────────────────────────────────────

func Test_Hashset_DistinctDiffLinesRaw_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		diff := h.DistinctDiffLinesRaw("b", "c")

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw -- 2 diff items", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_BothEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw_BothEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLinesRaw()

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw both empty -- 0", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_LeftEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw_LeftEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLinesRaw("a")

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw left empty -- right items", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_RightEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw_RightEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		diff := h.DistinctDiffLinesRaw()

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw right empty -- left items", actual)
	})
}

func Test_Hashset_DistinctDiffLines_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		diff := h.DistinctDiffLines("b", "c")

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines -- 2 diff items", actual)
	})
}

func Test_Hashset_DistinctDiffLines_BothEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines_BothEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLines()

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines both empty -- 0", actual)
	})
}

func Test_Hashset_DistinctDiffLines_LeftNotEmpty_RightEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines_LeftNotEmpty_RightEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		diff := h.DistinctDiffLines()

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines left only -- left items", actual)
	})
}

func Test_Hashset_DistinctDiffLines_LeftEmpty_RightNotEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines_LeftEmpty_RightNotEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLines("a")

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines right only -- right items", actual)
	})
}

func Test_Hashset_DistinctDiffHashset_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffHashset", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h2 := corestr.New.Hashset.Strings([]string{"b", "c"})
		diff := h.DistinctDiffHashset(h2)

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffHashset -- 2 diff items", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Hashset_Json_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Json", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		j := h.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Hashset_MarshalJSON_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MarshalJSON", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
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

func Test_Hashset_UnmarshalJSON_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_UnmarshalJSON", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		err := h.UnmarshalJSON([]byte(`{"a":true}`))

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

func Test_Hashset_UnmarshalJSON_Invalid_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_UnmarshalJSON_Invalid", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		err := h.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Hashset_JsonModel_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonModel", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(h.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_Hashset_JsonModel_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonModel_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"len": len(h.JsonModel())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "JsonModel empty -- 0", actual)
	})
}

func Test_Hashset_JsonModelAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonModelAny", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"notNil": h.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Hashset_ParseInjectUsingJson_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ParseInjectUsingJson", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Empty()
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

func Test_Hashset_ParseInjectUsingJsonMust_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ParseInjectUsingJsonMust", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Empty()
		result := h2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Hashset_JsonParseSelfInject_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonParseSelfInject", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Empty()
		err := h2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Hashset_Serialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Serialize", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
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

func Test_Hashset_Deserialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Deserialize", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		var dest map[string]bool
		err := h.Deserialize(&dest)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Hashset_InterfaceCasts_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_InterfaceCasts", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

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

// ══════════════════════════════════════════════════════════════════════════════
// HashsetDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashset_DataModel_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DataModel", func() {
		// Arrange
		dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
		h := corestr.NewHashsetUsingDataModel(dm)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetUsingDataModel -- 1 item", actual)
	})
}

func Test_Hashset_DataModel_Reverse_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DataModel_Reverse", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		dm := corestr.NewHashsetsDataModelUsing(h)

		// Act
		actual := args.Map{"len": len(dm.Items)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetsDataModelUsing -- 1 item", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashsetsCollection
// ══════════════════════════════════════════════════════════════════════════════

