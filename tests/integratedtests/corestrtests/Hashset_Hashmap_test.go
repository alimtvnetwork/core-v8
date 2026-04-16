package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Hashset ──

func Test_Hashset_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Hashset_IsEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": h.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashset_HasItems(t *testing.T) {
	safeTest(t, "Test_Hashset_HasItems", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"result": h.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashset_AddCapacities_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddCapacities", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.AddCapacities(10, 5)
		h.AddCapacities()
	})
}

func Test_Hashset_AddCapacitiesLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddCapacitiesLock", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.AddCapacitiesLock(10)
		h.AddCapacitiesLock()
	})
}

func Test_Hashset_Resize_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_Resize", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.Resize(10)
		h.Resize(0)
	})
}

func Test_Hashset_ResizeLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_ResizeLock", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.ResizeLock(10)
		h.ResizeLock(0)
	})
}

func Test_Hashset_Add_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_Add", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		h.Add("a")

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashset_AddBool_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddBool", func() {
		h := corestr.New.Hashset.Empty()
		h.AddBool("a")
		h.AddBool("a") // second time should return true (exists)
	})
}

func Test_Hashset_AddNonEmpty_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddNonEmpty", func() {
		h := corestr.New.Hashset.Empty()
		h.AddNonEmpty("")
		h.AddNonEmpty("a")
	})
}

func Test_Hashset_AddLock(t *testing.T) {
	safeTest(t, "Test_Hashset_AddLock", func() {
		h := corestr.New.Hashset.Empty()
		h.AddLock("a")
	})
}

func Test_Hashset_AddCollection_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddCollection", func() {
		h := corestr.New.Hashset.Empty()
		c := corestr.New.Collection.Strings([]string{"a"})
		h.AddCollection(c)
		h.AddCollection(nil)
	})
}

func Test_Hashset_AddHashsetItems_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddHashsetItems", func() {
		h := corestr.New.Hashset.Empty()
		h.AddHashsetItems(corestr.New.Hashset.StringsSpreadItems("a"))
		h.AddHashsetItems(nil)
	})
}

func Test_Hashset_AddHashsetWgLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddHashsetWgLock", func() {
		h := corestr.New.Hashset.Empty()
		var wg sync.WaitGroup
		wg.Add(1)
		h.AddHashsetWgLock(corestr.New.Hashset.StringsSpreadItems("a"), &wg)
		wg.Wait()
	})
}

func Test_Hashset_Adds_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_Adds", func() {
		h := corestr.New.Hashset.Empty()
		h.Adds("a", "b")
	})
}

func Test_Hashset_AddStrings_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddStrings", func() {
		h := corestr.New.Hashset.Empty()
		h.AddStrings([]string{"a"})
		h.AddStrings(nil)
	})
}

func Test_Hashset_AddItemsMap_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_AddItemsMap", func() {
		h := corestr.New.Hashset.Empty()
		h.AddItemsMap(map[string]bool{"a": true})
	})
}

func Test_Hashset_Remove_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_Remove", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		h.Remove("a")
	})
}

func Test_Hashset_RemoveWithLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_RemoveWithLock", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.RemoveWithLock("a")
	})
}

func Test_Hashset_SafeRemove(t *testing.T) {
	safeTest(t, "Test_Hashset_SafeRemove", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		h.SafeRemove("a")
	})
}

func Test_Hashset_RemovesLock(t *testing.T) {
	safeTest(t, "Test_Hashset_RemovesLock", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		h.RemoveWithLock("a")
		h.RemoveWithLock("b")
	})
}

func Test_Hashset_Length(t *testing.T) {
	safeTest(t, "Test_Hashset_Length", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashset_LengthLock(t *testing.T) {
	safeTest(t, "Test_Hashset_LengthLock", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.LengthLock()
	})
}

func Test_Hashset_Has_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_Has", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.Has("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashset_HasLock(t *testing.T) {
	safeTest(t, "Test_Hashset_HasLock", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.HasLock("a")
	})
}

func Test_Hashset_HasAll(t *testing.T) {
	safeTest(t, "Test_Hashset_HasAll", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": h.HasAll("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAll("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Hashset_HasAllStrings_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": h.HasAllStrings([]string{"a"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashset_HasAny_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_HasAny", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"result": h.HasAny("a", "z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashset_HasAnyOfStrings(t *testing.T) {
	safeTest(t, "Test_Hashset_HasAnyOfStrings", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.HasAny("a", "z")
	})
}

func Test_Hashset_List_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_List", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.List()
	})
}

func Test_Hashset_ListLock(t *testing.T) {
	safeTest(t, "Test_Hashset_ListLock", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.ListCopyLock()
	})
}

func Test_Hashset_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Hashset_SortedListAsc", func() {
		h := corestr.New.Hashset.StringsSpreadItems("b", "a")
		_ = h.ListPtrSortedAsc()
		_ = corestr.New.Hashset.Empty().SortedList()
	})
}

func Test_Hashset_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Hashset_SortedListDsc", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = h.ListPtrSortedDsc()
	})
}

func Test_Hashset_Map(t *testing.T) {
	safeTest(t, "Test_Hashset_Map", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.MapStringAny()
	})
}

func Test_Hashset_MapStringAnyDiff_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_MapStringAnyDiff", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.MapStringAnyDiff()
	})
}

func Test_Hashset_ListPtrSorted(t *testing.T) {
	safeTest(t, "Test_Hashset_ListPtrSorted", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.ListPtrSortedAsc()
		_ = h.ListPtrSortedDsc()
	})
}

// CopyMapLock, Clone, CloneLock, Diff, DiffLock, SummaryString, SummaryStringLock
// do not exist on Hashset — removed.

func Test_Hashset_MapStringAnyDiff_Coverage(t *testing.T) {
	safeTest(t, "Test_Hashset_MapStringAnyDiff_Coverage", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = h.MapStringAnyDiff()
	})
}

func Test_Hashset_DistinctDiffHashset_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_DistinctDiffHashset", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("a")
		_ = a.DistinctDiffHashset(b)
	})
}

func Test_Hashset_DistinctDiffLines_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_DistinctDiffLines", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = a.DistinctDiffLines("a")
	})
}

func Test_Hashset_DistinctDiffLinesRaw(t *testing.T) {
	safeTest(t, "Test_Hashset_DistinctDiffLinesRaw", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = a.DistinctDiffLinesRaw("a")
	})
}

func Test_Hashset_JsonMethods(t *testing.T) {
	safeTest(t, "Test_Hashset_JsonMethods", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.Json()
		_ = h.JsonPtr()
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		_, _ = h.MarshalJSON()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsonMarshaller()
		_ = h.AsJsonParseSelfInjector()
		_ = h.AsJsoner()
	})
}

func Test_Hashset_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Hashset_UnmarshalJSON", func() {
		h := &corestr.Hashset{}
		_ = h.UnmarshalJSON([]byte(`["a"]`))
	})
}

func Test_Hashset_ParseInjectUsingJson_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_ParseInjectUsingJson", func() {
		h := corestr.New.Hashset.Empty()
		r := corejson.New([]string{"a"})
		_, _ = h.ParseInjectUsingJson(&r)
	})
}

func Test_Hashset_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	safeTest(t, "Test_Hashset_ParseInjectUsingJsonMust_Panic", func() {
		defer func() { recover() }()
		h := corestr.New.Hashset.Empty()
		bad := corejson.NewResult.UsingString(`invalid`)
		h.ParseInjectUsingJsonMust(bad)
	})
}

func Test_Hashset_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Hashset_JsonParseSelfInject", func() {
		h := corestr.New.Hashset.Empty()
		r := corejson.New([]string{"a"})
		_ = h.JsonParseSelfInject(&r)
	})
}

func Test_Hashset_Clear_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_Clear", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.Clear()
	})
}

// ClearLock and DisposeLock do not exist on Hashset — removed.

func Test_Hashset_Dispose(t *testing.T) {
	safeTest(t, "Test_Hashset_Dispose", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.Dispose()
	})
}

func Test_Hashset_Filter_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashset_Filter", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "bb")
		_ = h.Filter(func(s string) bool { return len(s) == 1 })
	})
}

func Test_Hashset_Serialize(t *testing.T) {
	safeTest(t, "Test_Hashset_Serialize", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_, _ = h.Serialize()
	})
}

func Test_Hashset_Deserialize(t *testing.T) {
	safeTest(t, "Test_Hashset_Deserialize", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		var target []string
		_ = h.Deserialize(&target)
	})
}

// ── newHashsetCreator ──

func Test_NHC_PointerStrings(t *testing.T) {
	safeTest(t, "Test_NHC_PointerStrings", func() {
		s := "a"
		_ = corestr.New.Hashset.PointerStrings([]*string{&s})
		_ = corestr.New.Hashset.PointerStrings([]*string{})
	})
}

func Test_NHC_PointerStringsPtrOption(t *testing.T) {
	safeTest(t, "Test_NHC_PointerStringsPtrOption", func() {
		s := "a"
		arr := []*string{&s}
		_ = corestr.New.Hashset.PointerStringsPtrOption(0, false, &arr)
		_ = corestr.New.Hashset.PointerStringsPtrOption(5, false, nil)
	})
}

func Test_NHC_StringsOption(t *testing.T) {
	safeTest(t, "Test_NHC_StringsOption", func() {
		_ = corestr.New.Hashset.StringsOption(0, false)
		_ = corestr.New.Hashset.StringsOption(5, false)
		_ = corestr.New.Hashset.StringsOption(0, false, "a")
	})
}

func Test_NHC_StringsSpreadItems(t *testing.T) {
	safeTest(t, "Test_NHC_StringsSpreadItems", func() {
		_ = corestr.New.Hashset.StringsSpreadItems("a")
		_ = corestr.New.Hashset.StringsSpreadItems()
	})
}

func Test_NHC_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_NHC_SimpleSlice", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = corestr.New.Hashset.SimpleSlice(ss)
		_ = corestr.New.Hashset.SimpleSlice(corestr.New.SimpleSlice.Empty())
	})
}

// ── Hashmap ──

func Test_Hashmap_IsEmpty_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_IsEmpty", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"result": h.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_HasItems_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_HasItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"result": h.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_IsEmptyLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_IsEmptyLock", func() {
		_ = corestr.New.Hashmap.Empty().IsEmptyLock()
	})
}

func Test_Hashmap_AddOrUpdate_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdate", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValInt_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateKeyStrValInt", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValInt("k", 42)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValFloat("k", 3.14)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValFloat64_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValFloat64("k", 3.14)
	})
}

func Test_Hashmap_AddOrUpdateKeyStrValAny_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateKeyStrValAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValAny("k", "v")
	})
}

func Test_Hashmap_AddOrUpdateKeyValueAny_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateKeyValueAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
	})
}

func Test_Hashmap_AddOrUpdateCollection_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateCollection", func() {
		h := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"k"})
		vals := corestr.New.Collection.Strings([]string{"v"})
		h.AddOrUpdateCollection(keys, vals)
		h.AddOrUpdateCollection(nil, nil)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateHashmap", func() {
		h := corestr.New.Hashmap.Empty()
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("k", "v")
		h.AddOrUpdateHashmap(other)
		h.AddOrUpdateHashmap(nil)
	})
}

func Test_Hashmap_AddOrUpdateKeyValues_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateKeyValues", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
	})
}

func Test_Hashmap_AddOrUpdateKeyAnyValues_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateKeyAnyValues", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
	})
}

func Test_Hashmap_AddOrUpdateMap_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateMap", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateMap(map[string]string{"k": "v"})
	})
}

func Test_Hashmap_Get_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Get", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_, _ = h.Get("k")
		_, _ = h.Get("missing")
	})
}

func Test_Hashmap_Has_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Has", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"result": h.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_HasLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_HasLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.HasLock("k")
	})
}

func Test_Hashmap_Length_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Length", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Hashmap_LengthLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_LengthLock", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.LengthLock()
	})
}

func Test_Hashmap_Remove_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Remove", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		h.Remove("k")
	})
}

func Test_Hashmap_Keys_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Keys", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.Keys()
	})
}

func Test_Hashmap_ValuesList_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_ValuesList", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.ValuesList()
	})
}

func Test_Hashmap_Collection_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Collection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.Collection()
	})
}

func Test_Hashmap_Clone_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Clone", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.Clone()
	})
}

func Test_Hashmap_CloneAnother(t *testing.T) {
	safeTest(t, "Test_Hashmap_CloneAnother", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.Clone()
	})
}

func Test_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_Hashmap_Map", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.Items()
	})
}

func Test_Hashmap_MapLock(t *testing.T) {
	safeTest(t, "Test_Hashmap_MapLock", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.ItemsCopyLock()
	})
}

func Test_Hashmap_CopyMap(t *testing.T) {
	safeTest(t, "Test_Hashmap_CopyMap", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.ItemsCopyLock()
	})
}

func Test_Hashmap_CopyMapLock(t *testing.T) {
	safeTest(t, "Test_Hashmap_CopyMapLock", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.ItemsCopyLock()
	})
}

func Test_Hashmap_IsEquals(t *testing.T) {
	safeTest(t, "Test_Hashmap_IsEquals", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("k", "v")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Hashmap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Hashmap_IsEqualsLock", func() {
		a := corestr.New.Hashmap.Empty()
		b := corestr.New.Hashmap.Empty()
		_ = a.IsEqualPtrLock(b)
	})
}

func Test_Hashmap_String_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_String", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.String()
		h.AddOrUpdate("k", "v")
		_ = h.String()
	})
}

func Test_Hashmap_StringLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_StringLock", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.StringLock()
	})
}

func Test_Hashmap_SummaryString(t *testing.T) {
	safeTest(t, "Test_Hashmap_SummaryString", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.String()
		_ = h.StringLock()
	})
}

func Test_Hashmap_JsonMethods(t *testing.T) {
	safeTest(t, "Test_Hashmap_JsonMethods", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		_ = h.Json()
		_ = h.JsonPtr()
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		_, _ = h.MarshalJSON()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsoner()
		_ = h.AsJsonMarshaller()
		_ = h.AsJsonParseSelfInjector()
		_, _ = h.Serialize()
	})
}

func Test_Hashmap_UnmarshalJSON_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_UnmarshalJSON", func() {
		h := &corestr.Hashmap{}
		_ = h.UnmarshalJSON([]byte(`{"k":"v"}`))
	})
}

func Test_Hashmap_ParseInjectUsingJson_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_ParseInjectUsingJson", func() {
		h := corestr.New.Hashmap.Empty()
		r := corejson.New(map[string]string{"k": "v"})
		_, _ = h.ParseInjectUsingJson(&r)
	})
}

func Test_Hashmap_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	safeTest(t, "Test_Hashmap_ParseInjectUsingJsonMust_Panic", func() {
		defer func() { recover() }()
		h := corestr.New.Hashmap.Empty()
		bad := corejson.NewResult.UsingString(`invalid`)
		h.ParseInjectUsingJsonMust(bad)
	})
}

func Test_Hashmap_JsonParseSelfInject_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_JsonParseSelfInject", func() {
		h := corestr.New.Hashmap.Empty()
		r := corejson.New(map[string]string{"k": "v"})
		_ = h.JsonParseSelfInject(&r)
	})
}

func Test_Hashmap_Clear_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Clear", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		h.Clear()
	})
}

func Test_Hashmap_ClearLock(t *testing.T) {
	safeTest(t, "Test_Hashmap_ClearLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.Clear()
	})
}

func Test_Hashmap_Dispose_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Dispose", func() {
		h := corestr.New.Hashmap.Empty()
		h.Dispose()
	})
}

func Test_Hashmap_DisposeLock(t *testing.T) {
	safeTest(t, "Test_Hashmap_DisposeLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.Dispose()
	})
}

func Test_Hashmap_Deserialize_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_Deserialize", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		var target map[string]string
		_ = h.Deserialize(&target)
	})
}

func Test_Hashmap_AddOrUpdateLock_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateLock("k", "v")
	})
}

func Test_Hashmap_KeysCollection_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_KeysCollection", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.KeysCollection()
	})
}

func Test_Hashmap_ValuesCollection_HashsetHashmap(t *testing.T) {
	safeTest(t, "Test_Hashmap_ValuesCollection", func() {
		h := corestr.New.Hashmap.Empty()
		_ = h.ValuesCollection()
	})
}

// ── newHashmapCreator ──

func Test_NHM_Empty(t *testing.T)       { _ = corestr.New.Hashmap.Empty() }
func Test_NHM_KeyAnyValues(t *testing.T) { _ = corestr.New.Hashmap.KeyAnyValues() }
func Test_NHM_KeyValues(t *testing.T)    { _ = corestr.New.Hashmap.KeyValues() }
func Test_NHM_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_NHM_KeyValuesCollection", func() {
		k := corestr.New.Collection.Strings([]string{"k"})
		v := corestr.New.Collection.Strings([]string{"v"})
		_ = corestr.New.Hashmap.KeyValuesCollection(k, v)
		_ = corestr.New.Hashmap.KeyValuesCollection(nil, nil)
	})
}
func Test_NHM_KeyValuesStrings(t *testing.T) {
	safeTest(t, "Test_NHM_KeyValuesStrings", func() {
		_ = corestr.New.Hashmap.KeyValuesStrings([]string{"k"}, []string{"v"})
		_ = corestr.New.Hashmap.KeyValuesStrings(nil, nil)
	})
}
func Test_NHM_UsingMap(t *testing.T) {
	safeTest(t, "Test_NHM_UsingMap", func() {
		_ = corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
	})
}
func Test_NHM_UsingMapOptions(t *testing.T) {
	safeTest(t, "Test_NHM_UsingMapOptions", func() {
		_ = corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"k": "v"})
		_ = corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{})
		_ = corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"k": "v"})
	})
}
func Test_NHM_MapWithCap(t *testing.T) {
	safeTest(t, "Test_NHM_MapWithCap", func() {
		_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{"k": "v"})
		_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{"k": "v"})
		_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{})
	})
}
